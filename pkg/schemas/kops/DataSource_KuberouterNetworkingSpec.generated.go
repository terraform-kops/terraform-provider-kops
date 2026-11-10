package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceKuberouterNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"use_nf_tables_for_netpol": ComputedBool(),
			"external_ip_ranges":       ComputedList(String()),
			"load_balancer_ip_ranges":  ComputedList(String()),
		},
	}

	return res
}

func ExpandDataSourceKuberouterNetworkingSpec(in map[string]interface{}) kops.KuberouterNetworkingSpec {
	if in == nil {
		panic("expand KuberouterNetworkingSpec failure, in is nil")
	}
	return kops.KuberouterNetworkingSpec{
		UseNFTablesForNetpol: func(in interface{}) *bool {
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
		}(in["use_nf_tables_for_netpol"]),
		ExternalIPRanges: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["external_ip_ranges"]),
		LoadBalancerIPRanges: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["load_balancer_ip_ranges"]),
	}
}

func FlattenDataSourceKuberouterNetworkingSpecInto(in kops.KuberouterNetworkingSpec, out map[string]interface{}) {
	out["use_nf_tables_for_netpol"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.UseNFTablesForNetpol)
	out["external_ip_ranges"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.ExternalIPRanges)
	out["load_balancer_ip_ranges"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.LoadBalancerIPRanges)
}

func FlattenDataSourceKuberouterNetworkingSpec(in kops.KuberouterNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceKuberouterNetworkingSpecInto(in, out)
	return out
}
