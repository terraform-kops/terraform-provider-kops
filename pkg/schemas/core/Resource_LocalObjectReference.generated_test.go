package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	core "k8s.io/api/core/v1"
)

func TestExpandResourceLocalObjectReference(t *testing.T) {
	_default := core.LocalObjectReference{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want core.LocalObjectReference
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"name": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceLocalObjectReference(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceLocalObjectReference() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLocalObjectReferenceInto(t *testing.T) {
	_default := map[string]interface{}{
		"name": "",
	}
	type args struct {
		in core.LocalObjectReference
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.LocalObjectReference{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() core.LocalObjectReference {
					subject := core.LocalObjectReference{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceLocalObjectReferenceInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLocalObjectReference() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLocalObjectReference(t *testing.T) {
	_default := map[string]interface{}{
		"name": "",
	}
	type args struct {
		in core.LocalObjectReference
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: core.LocalObjectReference{},
			},
			want: _default,
		},
		{
			name: "Name - default",
			args: args{
				in: func() core.LocalObjectReference {
					subject := core.LocalObjectReference{}
					subject.Name = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceLocalObjectReference(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLocalObjectReference() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
