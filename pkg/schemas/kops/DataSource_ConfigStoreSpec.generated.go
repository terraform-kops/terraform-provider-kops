package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceConfigStoreSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"base":     ComputedString(),
			"keypairs": ComputedString(),
			"secrets":  ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceConfigStoreSpec(in map[string]interface{}) kops.ConfigStoreSpec {
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

func FlattenDataSourceConfigStoreSpecInto(in kops.ConfigStoreSpec, out map[string]interface{}) {
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

func FlattenDataSourceConfigStoreSpec(in kops.ConfigStoreSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceConfigStoreSpecInto(in, out)
	return out
}
