package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCiliumIngressSpec(t *testing.T) {
	_default := kops.CiliumIngressSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CiliumIngressSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"enabled":                           nil,
					"enforce_https":                     nil,
					"enable_secrets_sync":               nil,
					"load_balancer_annotation_prefixes": "",
					"default_load_balancer_mode":        "",
					"shared_load_balancer_service_name": "",
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCiliumIngressSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCiliumIngressSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCiliumIngressSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":                           nil,
		"enforce_https":                     nil,
		"enable_secrets_sync":               nil,
		"load_balancer_annotation_prefixes": "",
		"default_load_balancer_mode":        "",
		"shared_load_balancer_service_name": "",
	}
	type args struct {
		in kops.CiliumIngressSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CiliumIngressSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnforceHttps - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.EnforceHttps = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSecretsSync - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.EnableSecretsSync = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancerAnnotationPrefixes - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.LoadBalancerAnnotationPrefixes = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultLoadBalancerMode - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.DefaultLoadBalancerMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SharedLoadBalancerServiceName - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.SharedLoadBalancerServiceName = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceCiliumIngressSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCiliumIngressSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCiliumIngressSpec(t *testing.T) {
	_default := map[string]interface{}{
		"enabled":                           nil,
		"enforce_https":                     nil,
		"enable_secrets_sync":               nil,
		"load_balancer_annotation_prefixes": "",
		"default_load_balancer_mode":        "",
		"shared_load_balancer_service_name": "",
	}
	type args struct {
		in kops.CiliumIngressSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CiliumIngressSpec{},
			},
			want: _default,
		},
		{
			name: "Enabled - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.Enabled = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnforceHttps - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.EnforceHttps = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableSecretsSync - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.EnableSecretsSync = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LoadBalancerAnnotationPrefixes - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.LoadBalancerAnnotationPrefixes = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DefaultLoadBalancerMode - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.DefaultLoadBalancerMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SharedLoadBalancerServiceName - default",
			args: args{
				in: func() kops.CiliumIngressSpec {
					subject := kops.CiliumIngressSpec{}
					subject.SharedLoadBalancerServiceName = ""
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCiliumIngressSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCiliumIngressSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
