package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandDataSourceResourceFieldSelector(t *testing.T) {
	_default := core.ResourceFieldSelector{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.ResourceFieldSelector
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"container_name": "",
					"resource":       "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceResourceFieldSelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceResourceFieldSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceResourceFieldSelectorInto(t *testing.T) {
	_default := map[string]interface{}{
		"container_name": "",
		"resource":       "",
	}
	type args struct {
		in core.ResourceFieldSelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.ResourceFieldSelector{},
			},
			want: _default,
		},
		{
			name: "ContainerName - default",
			args: args{
				in: func() core.ResourceFieldSelector {
					subject := core.ResourceFieldSelector{}
					subject.ContainerName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Resource - default",
			args: args{
				in: func() core.ResourceFieldSelector {
					subject := core.ResourceFieldSelector{}
					subject.Resource = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceResourceFieldSelectorInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceResourceFieldSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceResourceFieldSelector(t *testing.T) {
	_default := map[string]interface{}{
		"container_name": "",
		"resource":       "",
	}
	type args struct {
		in core.ResourceFieldSelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.ResourceFieldSelector{},
			},
			want: _default,
		},
		{
			name: "ContainerName - default",
			args: args{
				in: func() core.ResourceFieldSelector {
					subject := core.ResourceFieldSelector{}
					subject.ContainerName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Resource - default",
			args: args{
				in: func() core.ResourceFieldSelector {
					subject := core.ResourceFieldSelector{}
					subject.Resource = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceResourceFieldSelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceResourceFieldSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
