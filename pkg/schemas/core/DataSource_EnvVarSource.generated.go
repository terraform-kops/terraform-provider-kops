package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourceEnvVarSource() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field_ref":          ComputedStruct(DataSourceObjectFieldSelector()),
			"resource_field_ref": ComputedStruct(DataSourceResourceFieldSelector()),
			"config_map_key_ref": ComputedStruct(DataSourceConfigMapKeySelector()),
			"secret_key_ref":     ComputedStruct(DataSourceSecretKeySelector()),
		},
	}

	return res
}

func ExpandDataSourceEnvVarSource(in map[string]interface{}) core.EnvVarSource {
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
						return ExpandDataSourceObjectFieldSelector(in[0].(map[string]interface{}))
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
						return ExpandDataSourceResourceFieldSelector(in[0].(map[string]interface{}))
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
						return ExpandDataSourceConfigMapKeySelector(in[0].(map[string]interface{}))
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
						return ExpandDataSourceSecretKeySelector(in[0].(map[string]interface{}))
					}
					return core.SecretKeySelector{}
				}(in))
			}(in)
		}(in["secret_key_ref"]),
	}
}

func FlattenDataSourceEnvVarSourceInto(in core.EnvVarSource, out map[string]interface{}) {
	out["field_ref"] = func(in *core.ObjectFieldSelector) interface{} {
		return func(in *core.ObjectFieldSelector) interface{} {
			if in == nil {
				return nil
			}
			return func(in core.ObjectFieldSelector) interface{} {
				return func(in core.ObjectFieldSelector) []interface{} {
					return []interface{}{FlattenDataSourceObjectFieldSelector(in)}
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
					return []interface{}{FlattenDataSourceResourceFieldSelector(in)}
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
					return []interface{}{FlattenDataSourceConfigMapKeySelector(in)}
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
					return []interface{}{FlattenDataSourceSecretKeySelector(in)}
				}(in)
			}(*in)
		}(in)
	}(in.SecretKeyRef)
}

func FlattenDataSourceEnvVarSource(in core.EnvVarSource) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceEnvVarSourceInto(in, out)
	return out
}
