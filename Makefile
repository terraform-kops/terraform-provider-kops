PROVIDER_VERSION          := "0.0.1"
OS                        := $(shell echo `uname` | tr '[:upper:]' '[:lower:]')
TOOLS_DIR                 := $(PWD)/.tools
GOIMPORTS                 := $(TOOLS_DIR)/goimports
GOIMPORTS_VERSION         := latest
GOARCH                    ?= $(shell uname -m)

ifeq (${GOARCH},x86_64)
	GOARCH := amd64
endif

# IMPROVEMENT 1: Fix dependency chain - remove circular dependencies
.PHONY: all
all: gen build test vet verify-gen

.PHONY: clean
clean:
	@rm -f terraform-provider-kops
	@rm -f .gen-timestamp
	@rm -rf ./pkg/schemas/config
	@rm -rf ./pkg/schemas/datasources
	@rm -rf ./pkg/schemas/kops
	@rm -rf ./pkg/schemas/kube
	@rm -rf ./pkg/schemas/resources
	@rm -rf ./pkg/schemas/utils
	@rm -rf ./docs/data-sources/*.md
	@rm -rf ./docs/provider-config/*.md
	@rm -rf ./docs/resources/*.md

$(GOIMPORTS):
	@echo Install goimports... >&2
	@GOBIN=$(TOOLS_DIR) go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)

# IMPROVEMENT 2: Incremental code generation - only regenerate if source changed
# Timestamp file tracks when generation last succeeded
.PHONY: gen-tf-code
gen-tf-code: $(GOIMPORTS) .gen-timestamp

.gen-timestamp: $(shell find pkg/api hack/gen-tf-code -name "*.go" -type f 2>/dev/null)
	@echo "Running code generation..."
	@rm -rf ./pkg/schemas/config
	@rm -rf ./pkg/schemas/datasources
	@rm -rf ./pkg/schemas/kops
	@rm -rf ./pkg/schemas/kube
	@rm -rf ./pkg/schemas/resources
	@rm -rf ./pkg/schemas/utils
	@go run ./hack/gen-tf-code/...
	@$(GOIMPORTS) -w ./pkg/schemas
	@touch .gen-timestamp

# IMPROVEMENT 3: Separate gen from formatting
.PHONY: gen
gen: gen-tf-code

.PHONY: build
build:
	@CGO_ENABLED=0 go build -ldflags="-s -w -X 'github.com/terraform-kops/terraform-provider-kops/pkg/version.BuildVersion=v${PROVIDER_VERSION}'" ./cmd/terraform-provider-kops

# IMPROVEMENT 4: Consolidated formatting - do it once
.PHONY: fmt
fmt:
	@goimports -w ./pkg ./cmd

.PHONY: verify-gen
verify-gen: gen
	@git --no-pager diff --exit-code ./pkg/schemas ./docs || \
		(echo 'Generated code is out of date. Run "make gen" and commit changes.' >&2 && exit 1)

# IMPROVEMENT 5: Parallel test execution
.PHONY: test
test:
	@go test -parallel=4 ./...

.PHONY: vet
vet:
	@go vet ./...

.PHONY: install
install: build
	@mkdir -p ${HOME}/.terraform.d/plugins/github/terraform-kops/kops/${PROVIDER_VERSION}/${OS}_${GOARCH}
	@cp terraform-provider-kops $(HOME)/.terraform.d/plugins/github/terraform-kops/kops/${PROVIDER_VERSION}/${OS}_${GOARCH}/terraform-provider-kops

# Quick build without regeneration (for iterative development)
.PHONY: quick
quick: build test vet

# EXAMPLES FOR TERRAFORM >= 0.15

.PHONY: examples
examples: example-basic example-aws-profile example-aws-assume-role example-aws-existing-vpc example-aws-kops-vpc example-bastion example-klog

.PHONY: example-basic
example-basic: install
	@terraform -chdir=./examples/basic init
	@terraform -chdir=./examples/basic validate
	@terraform -chdir=./examples/basic plan

.PHONY: example-aws-profile
example-aws-profile: install
	@terraform -chdir=./examples/aws-profile init
	@terraform -chdir=./examples/aws-profile validate
	@terraform -chdir=./examples/aws-profile plan

.PHONY: example-aws-assume-role
example-aws-assume-role: install
	@terraform -chdir=./examples/aws-assume-role init
	@terraform -chdir=./examples/aws-assume-role validate

.PHONY: example-aws-existing-vpc
example-aws-existing-vpc: install
	@terraform -chdir=./examples/aws-existing-vpc init
	@terraform -chdir=./examples/aws-existing-vpc validate

.PHONY: example-aws-kops-vpc
example-aws-kops-vpc: install
	@terraform -chdir=./examples/aws-kops-vpc init
	@terraform -chdir=./examples/aws-kops-vpc validate

.PHONY: example-bastion
example-bastion: install
	@terraform -chdir=./examples/bastion init
	@terraform -chdir=./examples/bastion validate
	@terraform -chdir=./examples/bastion plan

.PHONY: example-klog
example-klog: install
	@terraform -chdir=./examples/klog init
	@terraform -chdir=./examples/klog validate
	@terraform -chdir=./examples/klog plan

# INTEGRATION TESTS

.PHONY: integration
integration: integration-basic integration-external-policies

.PHONY: integration-reset
integration-reset:
	@rm -rf ./store
	@rm -f 	./terraform.tfstate

.PHONY: integration-basic
integration-basic: integration-reset
	@terraform -chdir=./tests/basic init
	@terraform -chdir=./tests/basic validate
	@terraform -chdir=./tests/basic plan
	@terraform -chdir=./tests/basic apply -auto-approve

.PHONY: integration-external-policies
integration-external-policies: integration-reset
	@terraform -chdir=./tests/external-policies init
	@terraform -chdir=./tests/external-policies validate
	@terraform -chdir=./tests/external-policies plan
	@terraform -chdir=./tests/external-policies apply -auto-approve
