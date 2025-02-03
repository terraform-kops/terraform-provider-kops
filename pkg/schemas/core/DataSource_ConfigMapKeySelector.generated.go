package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourceConfigMapKeySelector() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":     ComputedString(),
			"key":      RequiredString(),
			"optional": ComputedBool(),
		},
	}

	return res
}

func ExpandDataSourceConfigMapKeySelector(in map[string]interface{}) core.ConfigMapKeySelector {
	if in == nil {
		panic("expand ConfigMapKeySelector failure, in is nil")
	}
	return core.ConfigMapKeySelector{
		LocalObjectReference: func(in interface{}) core.LocalObjectReference {
			return ExpandDataSourceLocalObjectReference(in.(map[string]interface{}))
		}(in),
		Key: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["key"]),
		Optional: func(in interface{}) *bool {
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
		}(in["optional"]),
	}
}

func FlattenDataSourceConfigMapKeySelectorInto(in core.ConfigMapKeySelector, out map[string]interface{}) {
	FlattenDataSourceLocalObjectReferenceInto(in.LocalObjectReference, out)
	out["key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Key)
	out["optional"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Optional)
}

func FlattenDataSourceConfigMapKeySelector(in core.ConfigMapKeySelector) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceConfigMapKeySelectorInto(in, out)
	return out
}
