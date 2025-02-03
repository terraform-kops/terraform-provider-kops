package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandDataSourceConfigMapKeySelector(t *testing.T) {
	_default := core.ConfigMapKeySelector{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.ConfigMapKeySelector
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
			got := ExpandDataSourceConfigMapKeySelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceConfigMapKeySelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceConfigMapKeySelectorInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":     "",
		"key":      "",
		"optional": nil,
	}
	type args struct {
		in core.ConfigMapKeySelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.ConfigMapKeySelector{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() core.ConfigMapKeySelector {
					subject := core.ConfigMapKeySelector{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() core.ConfigMapKeySelector {
					subject := core.ConfigMapKeySelector{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Optional - default",
			args: args{
				in: func() core.ConfigMapKeySelector {
					subject := core.ConfigMapKeySelector{}
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
			FlattenDataSourceConfigMapKeySelectorInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceConfigMapKeySelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceConfigMapKeySelector(t *testing.T) {
	_default := map[string]interface{}{
		"name":     "",
		"key":      "",
		"optional": nil,
	}
	type args struct {
		in core.ConfigMapKeySelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.ConfigMapKeySelector{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() core.ConfigMapKeySelector {
					subject := core.ConfigMapKeySelector{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Key - default",
			args: args{
				in: func() core.ConfigMapKeySelector {
					subject := core.ConfigMapKeySelector{}
					subject.Key = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Optional - default",
			args: args{
				in: func() core.ConfigMapKeySelector {
					subject := core.ConfigMapKeySelector{}
					subject.Optional = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceConfigMapKeySelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceConfigMapKeySelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
