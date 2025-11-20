# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a Terraform provider for kOps (Kubernetes Operations) that enables managing Kubernetes clusters through Terraform without shell scripts or YAML templating. It integrates directly with the kOps API using pure Go code.

**Current kOps version**: v1.34.0
**Compatible Terraform versions**: 1.5.x and higher

## Build & Development Commands

### Building
```bash
# Full build with code generation, formatting, and tests
make all

# Quick iterative build (no code generation - for development)
make quick

# Just build the provider binary
make build

# Build and install locally for testing with Terraform
make install
```

**Build Performance**: Initial builds take ~3-4 minutes. Use `make quick` for faster iteration during development (~30 seconds) which skips code generation if API types haven't changed.

### Code Generation
```bash
# Generate Terraform schemas from kOps API types (incremental - only if needed)
make gen

# Force clean regeneration of all code
make clean gen
```

**IMPORTANT**:
- Schemas in `pkg/schemas/` are **auto-generated** (617 files) and should NOT be manually edited
- Code generation is **incremental** - only runs if `pkg/api/` or `hack/gen-tf-code/` changes
- All schema generation logic is in `hack/gen-tf-code/`
- Code generation takes ~90 seconds and is the main build bottleneck

### Testing
```bash
# Run unit tests
make test

# Run go vet
make vet

# Format code
make fmt

# Verify generated code is up to date (used in CI)
make verify-gen

# Run integration tests
make integration

# Run specific integration tests
make integration-basic
make integration-external-policies
```

### Example Validation
```bash
# Validate all examples
make examples

# Validate specific examples
make example-basic
make example-aws-profile
make example-bastion
```

## Architecture

### Code Organization

```
cmd/terraform-provider-kops/     # Provider entry point
pkg/
  ├── api/                       # API layer - core business logic
  │   ├── config/                # Provider configuration types
  │   ├── datasources/           # Data source API definitions
  │   ├── resources/             # Resource API definitions (Cluster, InstanceGroup, etc.)
  │   ├── utils/                 # Shared utilities (cluster operations)
  │   └── kube/                  # Kubernetes config handling
  ├── provider/                  # Terraform provider registration
  ├── resources/                 # Terraform resource implementations (CRUD operations)
  ├── datasources/               # Terraform data source implementations
  ├── schemas/                   # AUTO-GENERATED - Terraform schemas
  │   ├── config/                # Provider config schemas
  │   ├── datasources/           # Data source schemas
  │   ├── resources/             # Resource schemas
  │   ├── kops/                  # kOps API schemas
  │   ├── core/                  # Kubernetes core API schemas
  │   └── utils/                 # Schema utilities
  └── config/                    # Provider configuration handling
hack/gen-tf-code/                # Code generation tool
```

### Key Architectural Patterns

1. **Two-Layer Architecture**:
   - **API Layer** (`pkg/api/*`): Defines core types and business logic for interacting with kOps API
   - **Terraform Layer** (`pkg/resources/*`, `pkg/datasources/*`): Implements Terraform CRUD operations, delegates to API layer

2. **Code Generation**:
   - Run `hack/gen-tf-code/main.go` to generate Terraform schemas from Go API types
   - Uses reflection to parse types from `pkg/api/config`, `pkg/api/datasources`, `pkg/api/resources`
   - Generates schema definitions, expanders (TF state → Go types), and flatteners (Go types → TF state)
   - Templates located in `hack/gen-tf-code/templates/`

3. **Resource/Data Source Pattern**:
   - Each resource/data source has an API definition in `pkg/api/`
   - Auto-generated schemas in `pkg/schemas/` with Expand/Flatten functions
   - Terraform implementation in `pkg/resources/` or `pkg/datasources/` using the generated schemas

### Core Resources

- **kops_cluster**: Manages cluster configuration (API: `pkg/api/resources/Cluster.go`)
- **kops_instance_group**: Manages instance groups (API: `pkg/api/resources/InstanceGroup.go`)
- **kops_cluster_updater**: Handles cluster apply and rolling updates (API: `pkg/api/resources/ClusterUpdater.go`)

### Core Data Sources

- **kops_cluster**: Fetches cluster state
- **kops_instance_group**: Fetches instance group state
- **kops_cluster_status**: Fetches cluster validation status
- **kops_kube_config**: Retrieves kubeconfig

## Development Workflow

### Adding New Fields to Resources

1. Add field to the API type in `pkg/api/resources/` or `pkg/api/config/`
2. Run `make gen-tf-code` to regenerate schemas
3. Update business logic in `pkg/api/utils/` or resource handlers if needed
4. Run `make verify-gen` to ensure generated code matches source
5. Test with `make test` and validate examples with `make examples`

### Code Generation Details

The code generator (`hack/gen-tf-code/main.go`):
- Parses Go types using reflection and Go AST parsing
- Maps kOps API types to Terraform schema types
- Generates `Expand*` functions (converts Terraform data to Go structs)
- Generates `Flatten*` functions (converts Go structs to Terraform data)
- Creates both schema definitions and unit tests

### Common Gotchas

- **Generated files**: Never edit files in `pkg/schemas/` with `.generated.go` suffix - they're regenerated on every `make gen-tf-code`
- **Revision tracking**: Resources use a `revision` field that auto-increments on changes to trigger the cluster updater
- **kOps integration**: The provider uses kOps' VFS (Virtual File System) for state storage (S3, GCS, etc.)
- **Build performance**: Code generation is slow (~90s). For iterative development, use `make quick` to skip regeneration. See BUILD_OPTIMIZATION.md for details.

## CI/CD

GitHub Actions workflows:
- `build.yaml`: Builds provider and validates examples against multiple Terraform versions (optimized: pre-generates code, uses `-p 4` parallelism, removed disk cleanup)
- `verify-codegen.yaml`: Ensures generated code is up to date
- `release.yaml`: Handles releases via goreleaser (optimized: pre-generates code, uses `-p 4` parallelism)
- `tagpr.yml`: Manages version tagging

**GoReleaser Performance**: Optimized from ~25 minutes to ~8-12 minutes (52-68% faster). See `GORELEASER_OPTIMIZATION.md` for details.

## Testing Philosophy

- Unit tests for schema generation (auto-generated in `pkg/schemas/*_test.go`)
- Integration tests in `tests/` directory that create real Terraform plans
- Example validation ensures provider works with documented examples
