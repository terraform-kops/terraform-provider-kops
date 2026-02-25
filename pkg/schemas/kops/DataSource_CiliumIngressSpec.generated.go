package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCiliumIngressSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":                           ComputedBool(),
			"enforce_https":                     Nullable(ComputedBool()),
			"enable_secrets_sync":               Nullable(ComputedBool()),
			"load_balancer_annotation_prefixes": ComputedString(),
			"default_load_balancer_mode":        ComputedString(),
			"shared_load_balancer_service_name": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceCiliumIngressSpec(in map[string]interface{}) kops.CiliumIngressSpec {
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
			if in, ok := in.([]interface{}); ok && len(in) == 1 {
				return func(in interface{}) *bool {
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
				}(in[0].(map[string]interface{})["value"])
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
			if in, ok := in.([]interface{}); ok && len(in) == 1 {
				return func(in interface{}) *bool {
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
				}(in[0].(map[string]interface{})["value"])
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

func FlattenDataSourceCiliumIngressSpecInto(in kops.CiliumIngressSpec, out map[string]interface{}) {
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
		if in == nil {
			return nil
		}
		return []interface{}{map[string]interface{}{"value": func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)}}
	}(in.EnforceHttps)
	out["enable_secrets_sync"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return []interface{}{map[string]interface{}{"value": func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)}}
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

func FlattenDataSourceCiliumIngressSpec(in kops.CiliumIngressSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCiliumIngressSpecInto(in, out)
	return out
}
