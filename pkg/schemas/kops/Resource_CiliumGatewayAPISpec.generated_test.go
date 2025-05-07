package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCiliumGatewayAPISpec(t *testing.T) {
	_default := kops.CiliumGatewayAPISpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CiliumGatewayAPISpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":             nil,
					"enable_secrets_sync": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCiliumGatewayAPISpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCiliumGatewayAPISpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCiliumGatewayAPISpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"enable_secrets_sync": nil,
	}
	type args struct {
		in kops.CiliumGatewayAPISpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CiliumGatewayAPISpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.CiliumGatewayAPISpec {
					subject := kops.CiliumGatewayAPISpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSecretsSync - default",
			args: args{
				in: func() kops.CiliumGatewayAPISpec {
					subject := kops.CiliumGatewayAPISpec{}
					subject.EnableSecretsSync = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceCiliumGatewayAPISpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCiliumGatewayAPISpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCiliumGatewayAPISpec(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":             nil,
		"enable_secrets_sync": nil,
	}
	type args struct {
		in kops.CiliumGatewayAPISpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CiliumGatewayAPISpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.CiliumGatewayAPISpec {
					subject := kops.CiliumGatewayAPISpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSecretsSync - default",
			args: args{
				in: func() kops.CiliumGatewayAPISpec {
					subject := kops.CiliumGatewayAPISpec{}
					subject.EnableSecretsSync = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCiliumGatewayAPISpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCiliumGatewayAPISpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
