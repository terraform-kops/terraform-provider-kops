package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandResourceSecretKeySelector(t *testing.T) {
	_default := core.SecretKeySelector{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.SecretKeySelector
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"name":     "",
					"key":      "",
					"optional": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceSecretKeySelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceSecretKeySelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceSecretKeySelectorInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":     "",
		"key":      "",
		"optional": nil,
	}
	type args struct {
		in core.SecretKeySelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.SecretKeySelector{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() core.SecretKeySelector {
					subject := core.SecretKeySelector{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() core.SecretKeySelector {
					subject := core.SecretKeySelector{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Optional - default",
			args: args{
				in: func() core.SecretKeySelector {
					subject := core.SecretKeySelector{}
					subject.Optional = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceSecretKeySelectorInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceSecretKeySelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceSecretKeySelector(t *testing.T) {
	_default := map[string]interface{}{
		"name":     "",
		"key":      "",
		"optional": nil,
	}
	type args struct {
		in core.SecretKeySelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.SecretKeySelector{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() core.SecretKeySelector {
					subject := core.SecretKeySelector{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() core.SecretKeySelector {
					subject := core.SecretKeySelector{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Optional - default",
			args: args{
				in: func() core.SecretKeySelector {
					subject := core.SecretKeySelector{}
					subject.Optional = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceSecretKeySelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceSecretKeySelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
