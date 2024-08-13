package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceNRIConfig(t *testing.T) {
	_default := kops.NRIConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.NRIConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":                     nil,
					"plugin_registration_timeout": nil,
					"plugin_request_timeout":      nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceNRIConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceNRIConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNRIConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":                     nil,
		"plugin_registration_timeout": nil,
		"plugin_request_timeout":      nil,
	}
	type args struct {
		in kops.NRIConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NRIConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NRIConfig {
					subject := kops.NRIConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PluginRegistrationTimeout - default",
			args: args{
				in: func() kops.NRIConfig {
					subject := kops.NRIConfig{}
					subject.PluginRegistrationTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PluginRequestTimeout - default",
			args: args{
				in: func() kops.NRIConfig {
					subject := kops.NRIConfig{}
					subject.PluginRequestTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceNRIConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNRIConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceNRIConfig(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":                     nil,
		"plugin_registration_timeout": nil,
		"plugin_request_timeout":      nil,
	}
	type args struct {
		in kops.NRIConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.NRIConfig{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.NRIConfig {
					subject := kops.NRIConfig{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PluginRegistrationTimeout - default",
			args: args{
				in: func() kops.NRIConfig {
					subject := kops.NRIConfig{}
					subject.PluginRegistrationTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PluginRequestTimeout - default",
			args: args{
				in: func() kops.NRIConfig {
					subject := kops.NRIConfig{}
					subject.PluginRequestTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceNRIConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceNRIConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
