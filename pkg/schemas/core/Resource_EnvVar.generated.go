package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func ResourceEnvVar() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":       RequiredString(),
			"value":      OptionalString(),
			"value_from": OptionalStruct(ResourceEnvVarSource()),
		},
	}

	return res
}

func ExpandResourceEnvVar(in map[string]interface{}) core.EnvVar {
	if in == nil {
		panic("expand EnvVar failure, in is nil")
	}
	return core.EnvVar{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		Value: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["value"]),
		ValueFrom: func(in interface{}) *core.EnvVarSource {
			return func(in interface{}) *core.EnvVarSource {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in core.EnvVarSource) *core.EnvVarSource {
					return &in
				}(func(in interface{}) core.EnvVarSource {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceEnvVarSource(in[0].(map[string]interface{}))
					}
					return core.EnvVarSource{}
				}(in))
			}(in)
		}(in["value_from"]),
	}
}

func FlattenResourceEnvVarInto(in core.EnvVar, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["value"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Value)
	out["value_from"] = func(in *core.EnvVarSource) interface{} {
		return func(in *core.EnvVarSource) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.EnvVarSource) interface{} {
				return func(in core.EnvVarSource) []interface{} {
					return []interface{}{FlattenResourceEnvVarSource(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ValueFrom)
}

func FlattenResourceEnvVar(in core.EnvVar) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceEnvVarInto(in, out)
	return out
}
