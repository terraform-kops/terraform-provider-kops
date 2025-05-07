package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceKindnetMasqueradeSpec(t *testing.T) {
	_default := kops.KindnetMasqueradeSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KindnetMasqueradeSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":              nil,
					"non_masquerade_cidrs": func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceKindnetMasqueradeSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceKindnetMasqueradeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKindnetMasqueradeSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":              nil,
		"non_masquerade_cidrs": func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.KindnetMasqueradeSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KindnetMasqueradeSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.KindnetMasqueradeSpec {
					subject := kops.KindnetMasqueradeSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NonMasqueradeCidrs - default",
			args: args{
				in: func() kops.KindnetMasqueradeSpec {
					subject := kops.KindnetMasqueradeSpec{}
					subject.NonMasqueradeCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceKindnetMasqueradeSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKindnetMasqueradeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKindnetMasqueradeSpec(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":              nil,
		"non_masquerade_cidrs": func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.KindnetMasqueradeSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KindnetMasqueradeSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.KindnetMasqueradeSpec {
					subject := kops.KindnetMasqueradeSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NonMasqueradeCidrs - default",
			args: args{
				in: func() kops.KindnetMasqueradeSpec {
					subject := kops.KindnetMasqueradeSpec{}
					subject.NonMasqueradeCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceKindnetMasqueradeSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKindnetMasqueradeSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
