package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKindnetNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":                         OptionalString(),
			"network_policies":                OptionalBool(),
			"admin_network_policies":          OptionalBool(),
			"baseline_admin_network_policies": OptionalBool(),
			"dns_caching":                     OptionalBool(),
			"nat64":                           OptionalBool(),
			"fast_path_threshold":             OptionalInt(),
			"masquerade":                      OptionalStruct(ResourceKindnetMasqueradeSpec()),
			"log_level":                       OptionalInt(),
		},
	}

	return res
}

func ExpandResourceKindnetNetworkingSpec(in map[string]interface{}) kops.KindnetNetworkingSpec {
	if in == nil {
		panic("expand KindnetNetworkingSpec failure, in is nil")
	}
	return kops.KindnetNetworkingSpec{
		Version: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["version"]),
		NetworkPolicies: func(in interface{}) *bool {
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
		}(in["network_policies"]),
		AdminNetworkPolicies: func(in interface{}) *bool {
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
		}(in["admin_network_policies"]),
		BaselineAdminNetworkPolicies: func(in interface{}) *bool {
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
		}(in["baseline_admin_network_policies"]),
		DNSCaching: func(in interface{}) *bool {
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
		}(in["dns_caching"]),
		NAT64: func(in interface{}) *bool {
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
		}(in["nat64"]),
		FastPathThreshold: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["fast_path_threshold"]),
		Masquerade: func(in interface{}) *kops.KindnetMasqueradeSpec {
			return func(in interface{}) *kops.KindnetMasqueradeSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KindnetMasqueradeSpec) *kops.KindnetMasqueradeSpec {
					return &in
				}(func(in interface{}) kops.KindnetMasqueradeSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKindnetMasqueradeSpec(in[0].(map[string]interface{}))
					}
					return kops.KindnetMasqueradeSpec{}
				}(in))
			}(in)
		}(in["masquerade"]),
		LogLevel: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["log_level"]),
	}
}

func FlattenResourceKindnetNetworkingSpecInto(in kops.KindnetNetworkingSpec, out map[string]interface{}) {
	out["version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Version)
	out["network_policies"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.NetworkPolicies)
	out["admin_network_policies"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.AdminNetworkPolicies)
	out["baseline_admin_network_policies"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.BaselineAdminNetworkPolicies)
	out["dns_caching"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.DNSCaching)
	out["nat64"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.NAT64)
	out["fast_path_threshold"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.FastPathThreshold)
	out["masquerade"] = func(in *kops.KindnetMasqueradeSpec) interface{} {
		return func(in *kops.KindnetMasqueradeSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KindnetMasqueradeSpec) interface{} {
				return func(in kops.KindnetMasqueradeSpec) []interface{} {
					return []interface{}{FlattenResourceKindnetMasqueradeSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Masquerade)
	out["log_level"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.LogLevel)
}

func FlattenResourceKindnetNetworkingSpec(in kops.KindnetNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKindnetNetworkingSpecInto(in, out)
	return out
}
