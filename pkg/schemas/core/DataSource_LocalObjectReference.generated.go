package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func DataSourceLocalObjectReference() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceLocalObjectReference(in map[string]interface{}) core.LocalObjectReference {
	if in == nil {
		panic("expand LocalObjectReference failure, in is nil")
	}
	return core.LocalObjectReference{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
	}
}

func FlattenDataSourceLocalObjectReferenceInto(in core.LocalObjectReference, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
}

func FlattenDataSourceLocalObjectReference(in core.LocalObjectReference) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLocalObjectReferenceInto(in, out)
	return out
}
