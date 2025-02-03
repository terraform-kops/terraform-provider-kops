package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceClusterField(t *testing.T) {
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
			got := ExpandResourceClusterField(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceClusterField() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceClusterFieldInto(t *testing.T) {
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
			FlattenResourceClusterFieldInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceClusterField() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceClusterField(t *testing.T) {
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
			got := FlattenResourceClusterField(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceClusterField() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
