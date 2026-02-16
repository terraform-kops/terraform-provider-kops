package utils

import (
	"testing"

	"k8s.io/kops/pkg/apis/kops"
)

func makeIGWithName(name string, role kops.InstanceGroupRole) kops.InstanceGroup {
	ig := kops.InstanceGroup{
		Spec: kops.InstanceGroupSpec{
			Role: role,
		},
	}
	ig.Name = name
	return ig
}

func TestFilterInstanceGroups(t *testing.T) {
	items := []kops.InstanceGroup{
		makeIGWithName("master-0", kops.InstanceGroupRoleControlPlane),
		makeIGWithName("master-1", kops.InstanceGroupRoleControlPlane),
		makeIGWithName("nodes", kops.InstanceGroupRoleNode),
		makeIGWithName("runner", kops.InstanceGroupRoleNode),
		makeIGWithName("api-server", kops.InstanceGroupRoleAPIServer),
	}

	tests := []struct {
		name             string
		controlPlaneOnly bool
		excludeNames     []string
		wantNames        []string
	}{
		{
			name:             "no exclusions, all instance groups",
			controlPlaneOnly: false,
			excludeNames:     nil,
			wantNames:        []string{"master-0", "master-1", "nodes", "runner", "api-server"},
		},
		{
			name:             "no exclusions, control plane only",
			controlPlaneOnly: true,
			excludeNames:     nil,
			wantNames:        []string{"master-0", "master-1", "api-server"},
		},
		{
			name:             "exclude a node instance group",
			controlPlaneOnly: false,
			excludeNames:     []string{"runner"},
			wantNames:        []string{"master-0", "master-1", "nodes", "api-server"},
		},
		{
			name:             "exclude a control plane instance group",
			controlPlaneOnly: false,
			excludeNames:     []string{"master-0"},
			wantNames:        []string{"master-1", "nodes", "runner", "api-server"},
		},
		{
			name:             "exclude control plane IG from control plane only phase",
			controlPlaneOnly: true,
			excludeNames:     []string{"master-0"},
			wantNames:        []string{"master-1", "api-server"},
		},
		{
			name:             "exclude non-existent name has no effect",
			controlPlaneOnly: false,
			excludeNames:     []string{"does-not-exist"},
			wantNames:        []string{"master-0", "master-1", "nodes", "runner", "api-server"},
		},
		{
			name:             "empty exclude list same as nil",
			controlPlaneOnly: false,
			excludeNames:     []string{},
			wantNames:        []string{"master-0", "master-1", "nodes", "runner", "api-server"},
		},
		{
			name:             "exclude multiple instance groups",
			controlPlaneOnly: false,
			excludeNames:     []string{"runner", "nodes"},
			wantNames:        []string{"master-0", "master-1", "api-server"},
		},
		{
			name:             "exclude node IG from control plane only phase has no effect",
			controlPlaneOnly: true,
			excludeNames:     []string{"runner"},
			wantNames:        []string{"master-0", "master-1", "api-server"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterInstanceGroups(items, tt.controlPlaneOnly, tt.excludeNames)

			if len(got) != len(tt.wantNames) {
				t.Fatalf("got %d instance groups, want %d", len(got), len(tt.wantNames))
			}
			for i, ig := range got {
				if ig.Name != tt.wantNames[i] {
					t.Errorf("instance group %d: got name %q, want %q", i, ig.Name, tt.wantNames[i])
				}
			}
		})
	}
}

func TestFilterInstanceGroupsEmptyInput(t *testing.T) {
	got := filterInstanceGroups(nil, false, nil)
	if got != nil {
		t.Errorf("expected nil for empty input, got %v", got)
	}

	got = filterInstanceGroups([]kops.InstanceGroup{}, false, []string{"runner"})
	if got != nil {
		t.Errorf("expected nil for empty items, got %v", got)
	}
}
