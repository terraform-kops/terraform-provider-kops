package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceGVisorConfig(t *testing.T) {
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
			got := ExpandResourceGVisorConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceGVisorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceGVisorConfigInto(t *testing.T) {
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
			FlattenResourceGVisorConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceGVisorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceGVisorConfig(t *testing.T) {
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
			got := FlattenResourceGVisorConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceGVisorConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
