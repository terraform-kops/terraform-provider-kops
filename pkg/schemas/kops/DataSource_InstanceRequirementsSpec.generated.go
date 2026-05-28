package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceInstanceRequirementsSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu":                     ComputedStruct(DataSourceMinMaxSpec()),
			"memory":                  ComputedStruct(DataSourceMinMaxSpec()),
			"excluded_instance_types": ComputedList(String()),
		},
	}

	return res
}

func ExpandDataSourceInstanceRequirementsSpec(in map[string]interface{}) kops.InstanceRequirementsSpec {
	if in == nil {
		panic("expand InstanceRequirementsSpec failure, in is nil")
	}
	return kops.InstanceRequirementsSpec{
		CPU: func(in interface{}) *kops.MinMaxSpec {
			return func(in interface{}) *kops.MinMaxSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.MinMaxSpec) *kops.MinMaxSpec {
					return &in
				}(func(in interface{}) kops.MinMaxSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceMinMaxSpec(in[0].(map[string]interface{}))
					}
					return kops.MinMaxSpec{}
				}(in))
			}(in)
		}(in["cpu"]),
		Memory: func(in interface{}) *kops.MinMaxSpec {
			return func(in interface{}) *kops.MinMaxSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.MinMaxSpec) *kops.MinMaxSpec {
					return &in
				}(func(in interface{}) kops.MinMaxSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceMinMaxSpec(in[0].(map[string]interface{}))
					}
					return kops.MinMaxSpec{}
				}(in))
			}(in)
		}(in["memory"]),
		ExcludedInstanceTypes: func(in interface{}) []string {
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
		}(in["excluded_instance_types"]),
	}
}

func FlattenDataSourceInstanceRequirementsSpecInto(in kops.InstanceRequirementsSpec, out map[string]interface{}) {
	out["cpu"] = func(in *kops.MinMaxSpec) interface{} {
		return func(in *kops.MinMaxSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.MinMaxSpec) interface{} {
				return func(in kops.MinMaxSpec) []interface{} {
					return []interface{}{FlattenDataSourceMinMaxSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.CPU)
	out["memory"] = func(in *kops.MinMaxSpec) interface{} {
		return func(in *kops.MinMaxSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.MinMaxSpec) interface{} {
				return func(in kops.MinMaxSpec) []interface{} {
					return []interface{}{FlattenDataSourceMinMaxSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Memory)
	out["excluded_instance_types"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.ExcludedInstanceTypes)
}

func FlattenDataSourceInstanceRequirementsSpec(in kops.InstanceRequirementsSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceInstanceRequirementsSpecInto(in, out)
	return out
}
