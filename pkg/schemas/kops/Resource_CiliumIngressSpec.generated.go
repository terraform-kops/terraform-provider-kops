package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceCiliumIngressSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":                           OptionalBool(),
			"enforce_https":                     OptionalBool(),
			"enable_secrets_sync":               OptionalBool(),
			"load_balancer_annotation_prefixes": OptionalString(),
			"default_load_balancer_mode":        OptionalString(),
			"shared_load_balancer_service_name": OptionalString(),
		},
	}

	return res
}

func ExpandResourceCiliumIngressSpec(in map[string]interface{}) kops.CiliumIngressSpec {
	if in == nil {
		panic("expand CiliumIngressSpec failure, in is nil")
	}
	return kops.CiliumIngressSpec{
		Enabled: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["enabled"]),
		EnforceHttps: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["enforce_https"]),
		EnableSecretsSync: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["enable_secrets_sync"]),
		LoadBalancerAnnotationPrefixes: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["load_balancer_annotation_prefixes"]),
		DefaultLoadBalancerMode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["default_load_balancer_mode"]),
		SharedLoadBalancerServiceName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["shared_load_balancer_service_name"]),
	}
}

func FlattenResourceCiliumIngressSpecInto(in kops.CiliumIngressSpec, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Enabled)
	out["enforce_https"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnforceHttps)
	out["enable_secrets_sync"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EnableSecretsSync)
	out["load_balancer_annotation_prefixes"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.LoadBalancerAnnotationPrefixes)
	out["default_load_balancer_mode"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DefaultLoadBalancerMode)
	out["shared_load_balancer_service_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SharedLoadBalancerServiceName)
}

func FlattenResourceCiliumIngressSpec(in kops.CiliumIngressSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceCiliumIngressSpecInto(in, out)
	return out
}
