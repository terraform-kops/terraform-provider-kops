package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func ResourceEnvVarSource() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field_ref":          OptionalStruct(ResourceObjectFieldSelector()),
			"resource_field_ref": OptionalStruct(ResourceResourceFieldSelector()),
			"config_map_key_ref": OptionalStruct(ResourceConfigMapKeySelector()),
			"secret_key_ref":     OptionalStruct(ResourceSecretKeySelector()),
		},
	}

	return res
}

func ExpandResourceEnvVarSource(in map[string]interface{}) core.EnvVarSource {
	if in == nil {
		panic("expand EnvVarSource failure, in is nil")
	}
	return core.EnvVarSource{
		FieldRef: func(in interface{}) *core.ObjectFieldSelector {
			return func(in interface{}) *core.ObjectFieldSelector {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in core.ObjectFieldSelector) *core.ObjectFieldSelector {
					return &in
				}(func(in interface{}) core.ObjectFieldSelector {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceObjectFieldSelector(in[0].(map[string]interface{}))
					}
					return core.ObjectFieldSelector{}
				}(in))
			}(in)
		}(in["field_ref"]),
		ResourceFieldRef: func(in interface{}) *core.ResourceFieldSelector {
			return func(in interface{}) *core.ResourceFieldSelector {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in core.ResourceFieldSelector) *core.ResourceFieldSelector {
					return &in
				}(func(in interface{}) core.ResourceFieldSelector {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceResourceFieldSelector(in[0].(map[string]interface{}))
					}
					return core.ResourceFieldSelector{}
				}(in))
			}(in)
		}(in["resource_field_ref"]),
		ConfigMapKeyRef: func(in interface{}) *core.ConfigMapKeySelector {
			return func(in interface{}) *core.ConfigMapKeySelector {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in core.ConfigMapKeySelector) *core.ConfigMapKeySelector {
					return &in
				}(func(in interface{}) core.ConfigMapKeySelector {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceConfigMapKeySelector(in[0].(map[string]interface{}))
					}
					return core.ConfigMapKeySelector{}
				}(in))
			}(in)
		}(in["config_map_key_ref"]),
		SecretKeyRef: func(in interface{}) *core.SecretKeySelector {
			return func(in interface{}) *core.SecretKeySelector {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in core.SecretKeySelector) *core.SecretKeySelector {
					return &in
				}(func(in interface{}) core.SecretKeySelector {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceSecretKeySelector(in[0].(map[string]interface{}))
					}
					return core.SecretKeySelector{}
				}(in))
			}(in)
		}(in["secret_key_ref"]),
	}
}

func FlattenResourceEnvVarSourceInto(in core.EnvVarSource, out map[string]interface{}) {
	out["field_ref"] = func(in *core.ObjectFieldSelector) interface{} {
		return func(in *core.ObjectFieldSelector) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.ObjectFieldSelector) interface{} {
				return func(in core.ObjectFieldSelector) []interface{} {
					return []interface{}{FlattenResourceObjectFieldSelector(in)}
				}(in)
			}(*in)
		}(in)
	}(in.FieldRef)
	out["resource_field_ref"] = func(in *core.ResourceFieldSelector) interface{} {
		return func(in *core.ResourceFieldSelector) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.ResourceFieldSelector) interface{} {
				return func(in core.ResourceFieldSelector) []interface{} {
					return []interface{}{FlattenResourceResourceFieldSelector(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ResourceFieldRef)
	out["config_map_key_ref"] = func(in *core.ConfigMapKeySelector) interface{} {
		return func(in *core.ConfigMapKeySelector) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.ConfigMapKeySelector) interface{} {
				return func(in core.ConfigMapKeySelector) []interface{} {
					return []interface{}{FlattenResourceConfigMapKeySelector(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ConfigMapKeyRef)
	out["secret_key_ref"] = func(in *core.SecretKeySelector) interface{} {
		return func(in *core.SecretKeySelector) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.SecretKeySelector) interface{} {
				return func(in core.SecretKeySelector) []interface{} {
					return []interface{}{FlattenResourceSecretKeySelector(in)}
				}(in)
			}(*in)
		}(in)
	}(in.SecretKeyRef)
}

func FlattenResourceEnvVarSource(in core.EnvVarSource) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceEnvVarSourceInto(in, out)
	return out
}
