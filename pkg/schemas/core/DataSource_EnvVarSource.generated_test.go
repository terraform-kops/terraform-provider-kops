package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandDataSourceEnvVarSource(t *testing.T) {
	_default := core.EnvVarSource{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.EnvVarSource
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"field_ref":          nil,
					"resource_field_ref": nil,
					"config_map_key_ref": nil,
					"secret_key_ref":     nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceEnvVarSource(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceEnvVarSource() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEnvVarSourceInto(t *testing.T) {
	_default := map[string]interface{}{
		"field_ref":          nil,
		"resource_field_ref": nil,
		"config_map_key_ref": nil,
		"secret_key_ref":     nil,
	}
	type args struct {
		in core.EnvVarSource
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.EnvVarSource{},
			},
			want: _default,
		},
		{
			name: "FieldRef - default",
			args: args{
				in: func() core.EnvVarSource {
					subject := core.EnvVarSource{}
					subject.FieldRef = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ResourceFieldRef - default",
			args: args{
				in: func() core.EnvVarSource {
					subject := core.EnvVarSource{}
					subject.ResourceFieldRef = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigMapKeyRef - default",
			args: args{
				in: func() core.EnvVarSource {
					subject := core.EnvVarSource{}
					subject.ConfigMapKeyRef = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretKeyRef - default",
			args: args{
				in: func() core.EnvVarSource {
					subject := core.EnvVarSource{}
					subject.SecretKeyRef = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceEnvVarSourceInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEnvVarSource() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEnvVarSource(t *testing.T) {
	_default := map[string]interface{}{
		"field_ref":          nil,
		"resource_field_ref": nil,
		"config_map_key_ref": nil,
		"secret_key_ref":     nil,
	}
	type args struct {
		in core.EnvVarSource
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.EnvVarSource{},
			},
			want: _default,
		},
		{
			name: "FieldRef - default",
			args: args{
				in: func() core.EnvVarSource {
					subject := core.EnvVarSource{}
					subject.FieldRef = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ResourceFieldRef - default",
			args: args{
				in: func() core.EnvVarSource {
					subject := core.EnvVarSource{}
					subject.ResourceFieldRef = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigMapKeyRef - default",
			args: args{
				in: func() core.EnvVarSource {
					subject := core.EnvVarSource{}
					subject.ConfigMapKeyRef = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretKeyRef - default",
			args: args{
				in: func() core.EnvVarSource {
					subject := core.EnvVarSource{}
					subject.SecretKeyRef = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceEnvVarSource(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEnvVarSource() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
