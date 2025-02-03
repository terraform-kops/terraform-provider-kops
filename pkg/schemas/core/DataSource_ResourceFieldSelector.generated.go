package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourceResourceFieldSelector() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"container_name": ComputedString(),
			"resource":       RequiredString(),
		},
	}

	return res
}

func ExpandDataSourceResourceFieldSelector(in map[string]interface{}) core.ResourceFieldSelector {
	if in == nil {
		panic("expand ResourceFieldSelector failure, in is nil")
	}
	return core.ResourceFieldSelector{
		ContainerName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["container_name"]),
		Resource: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["resource"]),
	}
}

func FlattenDataSourceResourceFieldSelectorInto(in core.ResourceFieldSelector, out map[string]interface{}) {
	out["container_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ContainerName)
	out["resource"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Resource)
}

func FlattenDataSourceResourceFieldSelector(in core.ResourceFieldSelector) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceResourceFieldSelectorInto(in, out)
	return out
}
