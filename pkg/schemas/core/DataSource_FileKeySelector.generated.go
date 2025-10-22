package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourceFileKeySelector() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"volume_name": ComputedString(),
			"path":        ComputedString(),
			"key":         RequiredString(),
			"optional":    ComputedBool(),
		},
	}

	return res
}

func ExpandDataSourceFileKeySelector(in map[string]interface{}) core.FileKeySelector {
	if in == nil {
		panic("expand FileKeySelector failure, in is nil")
	}
	return core.FileKeySelector{
		VolumeName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["volume_name"]),
		Path: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["path"]),
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

func FlattenDataSourceFileKeySelectorInto(in core.FileKeySelector, out map[string]interface{}) {
	out["volume_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.VolumeName)
	out["path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Path)
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

func FlattenDataSourceFileKeySelector(in core.FileKeySelector) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceFileKeySelectorInto(in, out)
	return out
}
