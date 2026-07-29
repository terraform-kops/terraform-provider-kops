package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceLinodeSpec(t *testing.T) {
	_default := kops.LinodeSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.LinodeSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceLinodeSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceLinodeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLinodeSpecInto(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kops.LinodeSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LinodeSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceLinodeSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLinodeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceLinodeSpec(t *testing.T) {
	_default := map[string]interface{}{}
	type args struct {
		in kops.LinodeSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LinodeSpec{},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceLinodeSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceLinodeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
