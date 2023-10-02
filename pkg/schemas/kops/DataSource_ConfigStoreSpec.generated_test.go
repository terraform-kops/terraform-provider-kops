package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceConfigStoreSpec(t *testing.T) {
	_default := kops.ConfigStoreSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.ConfigStoreSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"base":     "",
					"keypairs": "",
					"secrets":  "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceConfigStoreSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceConfigStoreSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceConfigStoreSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"base":     "",
		"keypairs": "",
		"secrets":  "",
	}
	type args struct {
		in kops.ConfigStoreSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ConfigStoreSpec{},
			},
			want: _default,
		},
		{
			name: "Base - default",
			args: args{
				in: func() kops.ConfigStoreSpec {
					subject := kops.ConfigStoreSpec{}
					subject.Base = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Keypairs - default",
			args: args{
				in: func() kops.ConfigStoreSpec {
					subject := kops.ConfigStoreSpec{}
					subject.Keypairs = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secrets - default",
			args: args{
				in: func() kops.ConfigStoreSpec {
					subject := kops.ConfigStoreSpec{}
					subject.Secrets = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceConfigStoreSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceConfigStoreSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceConfigStoreSpec(t *testing.T) {
	_default := map[string]interface{}{
		"base":     "",
		"keypairs": "",
		"secrets":  "",
	}
	type args struct {
		in kops.ConfigStoreSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.ConfigStoreSpec{},
			},
			want: _default,
		},
		{
			name: "Base - default",
			args: args{
				in: func() kops.ConfigStoreSpec {
					subject := kops.ConfigStoreSpec{}
					subject.Base = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Keypairs - default",
			args: args{
				in: func() kops.ConfigStoreSpec {
					subject := kops.ConfigStoreSpec{}
					subject.Keypairs = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secrets - default",
			args: args{
				in: func() kops.ConfigStoreSpec {
					subject := kops.ConfigStoreSpec{}
					subject.Secrets = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceConfigStoreSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceConfigStoreSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
