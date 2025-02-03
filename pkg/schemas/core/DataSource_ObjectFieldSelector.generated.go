package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourceObjectFieldSelector() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_version": ComputedString(),
			"field_path":  RequiredString(),
		},
	}

	return res
}

func ExpandDataSourceObjectFieldSelector(in map[string]interface{}) core.ObjectFieldSelector {
	if in == nil {
		panic("expand ObjectFieldSelector failure, in is nil")
	}
	return core.ObjectFieldSelector{
		APIVersion: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["api_version"]),
		FieldPath: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["field_path"]),
	}
}

func FlattenDataSourceObjectFieldSelectorInto(in core.ObjectFieldSelector, out map[string]interface{}) {
	out["api_version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.APIVersion)
	out["field_path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.FieldPath)
}

func FlattenDataSourceObjectFieldSelector(in core.ObjectFieldSelector) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceObjectFieldSelectorInto(in, out)
	return out
}
