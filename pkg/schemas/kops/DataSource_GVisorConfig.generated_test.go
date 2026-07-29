package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceGVisorConfig(t *testing.T) {
	_default := kops.GVisorConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.GVisorConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":  nil,
					"platform": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceGVisorConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceGVisorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceGVisorConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":  nil,
		"platform": "",
	}
	type args struct {
		in kops.GVisorConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.GVisorConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.GVisorConfig {
					subject := kops.GVisorConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Platform - default",
			args: args{
				in: func() kops.GVisorConfig {
					subject := kops.GVisorConfig{}
					subject.Platform = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceGVisorConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceGVisorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceGVisorConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":  nil,
		"platform": "",
	}
	type args struct {
		in kops.GVisorConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.GVisorConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.GVisorConfig {
					subject := kops.GVisorConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Platform - default",
			args: args{
				in: func() kops.GVisorConfig {
					subject := kops.GVisorConfig{}
					subject.Platform = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceGVisorConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceGVisorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
