package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceClusterField(t *testing.T) {
	_default := kops.ClusterField{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ClusterField
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"path": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceClusterField(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceClusterField() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceClusterFieldInto(t *testing.T) {
	_default := map[string]interface{}{
		"path": "",
	}
	type args struct {
		in kops.ClusterField
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ClusterField{},
			},
			want: _default,
		},
		{
			name: "Path - default",
			args: args{
				in: func() kops.ClusterField {
					subject := kops.ClusterField{}
					subject.Path = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceClusterFieldInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceClusterField() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceClusterField(t *testing.T) {
	_default := map[string]interface{}{
		"path": "",
	}
	type args struct {
		in kops.ClusterField
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ClusterField{},
			},
			want: _default,
		},
		{
			name: "Path - default",
			args: args{
				in: func() kops.ClusterField {
					subject := kops.ClusterField{}
					subject.Path = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceClusterField(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceClusterField() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
