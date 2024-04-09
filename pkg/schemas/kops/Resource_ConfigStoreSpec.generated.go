package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceConfigStoreSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"base":     OptionalComputedString(),
			"keypairs": OptionalString(),
			"secrets":  OptionalString(),
		},
	}

	return res
}

func ExpandResourceConfigStoreSpec(in map[string]interface{}) kops.ConfigStoreSpec {
	if in == nil {
		panic("expand ConfigStoreSpec failure, in is nil")
	}
	return kops.ConfigStoreSpec{
		Base: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["base"]),
		Keypairs: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["keypairs"]),
		Secrets: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["secrets"]),
	}
}

func FlattenResourceConfigStoreSpecInto(in kops.ConfigStoreSpec, out map[string]interface{}) {
	out["base"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Base)
	out["keypairs"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Keypairs)
	out["secrets"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Secrets)
}

func FlattenResourceConfigStoreSpec(in kops.ConfigStoreSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceConfigStoreSpecInto(in, out)
	return out
}
