package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceConfigStoreSpec(t *testing.T) {
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
			got := ExpandResourceConfigStoreSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceConfigStoreSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceConfigStoreSpecInto(t *testing.T) {
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
			FlattenResourceConfigStoreSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceConfigStoreSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceConfigStoreSpec(t *testing.T) {
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
			got := FlattenResourceConfigStoreSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceConfigStoreSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
