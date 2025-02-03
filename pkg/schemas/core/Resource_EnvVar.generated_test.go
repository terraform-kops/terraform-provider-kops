package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandResourceEnvVar(t *testing.T) {
	_default := core.EnvVar{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.EnvVar
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"name":       "",
					"value":      "",
					"value_from": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceEnvVar(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceEnvVar() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceEnvVarInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":       "",
		"value":      "",
		"value_from": nil,
	}
	type args struct {
		in core.EnvVar
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.EnvVar{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() core.EnvVar {
					subject := core.EnvVar{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Value - default",
			args: args{
				in: func() core.EnvVar {
					subject := core.EnvVar{}
					subject.Value = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ValueFrom - default",
			args: args{
				in: func() core.EnvVar {
					subject := core.EnvVar{}
					subject.ValueFrom = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceEnvVarInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEnvVar() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceEnvVar(t *testing.T) {
	_default := map[string]interface{}{
		"name":       "",
		"value":      "",
		"value_from": nil,
	}
	type args struct {
		in core.EnvVar
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.EnvVar{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() core.EnvVar {
					subject := core.EnvVar{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Value - default",
			args: args{
				in: func() core.EnvVar {
					subject := core.EnvVar{}
					subject.Value = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ValueFrom - default",
			args: args{
				in: func() core.EnvVar {
					subject := core.EnvVar{}
					subject.ValueFrom = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceEnvVar(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceEnvVar() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
