package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func ResourceLocalObjectReference() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": OptionalString(),
		},
	}

	return res
}

func ExpandResourceLocalObjectReference(in map[string]interface{}) core.LocalObjectReference {
	if in == nil {
		panic("expand LocalObjectReference failure, in is nil")
	}
	return core.LocalObjectReference{
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
	}
}

func FlattenResourceLocalObjectReferenceInto(in core.LocalObjectReference, out map[string]interface{}) {
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
}

func FlattenResourceLocalObjectReference(in core.LocalObjectReference) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceLocalObjectReferenceInto(in, out)
	return out
}
