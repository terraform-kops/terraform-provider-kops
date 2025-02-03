package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandDataSourceObjectFieldSelector(t *testing.T) {
	_default := core.ObjectFieldSelector{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.ObjectFieldSelector
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"api_version": "",
					"field_path":  "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceObjectFieldSelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceObjectFieldSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceObjectFieldSelectorInto(t *testing.T) {
	_default := map[string]interface{}{
		"api_version": "",
		"field_path":  "",
	}
	type args struct {
		in core.ObjectFieldSelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.ObjectFieldSelector{},
			},
			want: _default,
		},
		{
			name: "ApiVersion - default",
			args: args{
				in: func() core.ObjectFieldSelector {
					subject := core.ObjectFieldSelector{}
					subject.APIVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FieldPath - default",
			args: args{
				in: func() core.ObjectFieldSelector {
					subject := core.ObjectFieldSelector{}
					subject.FieldPath = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceObjectFieldSelectorInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceObjectFieldSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceObjectFieldSelector(t *testing.T) {
	_default := map[string]interface{}{
		"api_version": "",
		"field_path":  "",
	}
	type args struct {
		in core.ObjectFieldSelector
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.ObjectFieldSelector{},
			},
			want: _default,
		},
		{
			name: "ApiVersion - default",
			args: args{
				in: func() core.ObjectFieldSelector {
					subject := core.ObjectFieldSelector{}
					subject.APIVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FieldPath - default",
			args: args{
				in: func() core.ObjectFieldSelector {
					subject := core.ObjectFieldSelector{}
					subject.FieldPath = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceObjectFieldSelector(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceObjectFieldSelector() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
