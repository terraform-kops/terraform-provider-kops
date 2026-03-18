package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceDiscoveryServiceOptions() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceDiscoveryServiceOptions(in map[string]interface{}) kops.DiscoveryServiceOptions {
	if in == nil {
		panic("expand DiscoveryServiceOptions failure, in is nil")
	}
	return kops.DiscoveryServiceOptions{
		URL: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["url"]),
	}
}

func FlattenDataSourceDiscoveryServiceOptionsInto(in kops.DiscoveryServiceOptions, out map[string]interface{}) {
	out["url"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.URL)
}

func FlattenDataSourceDiscoveryServiceOptions(in kops.DiscoveryServiceOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceDiscoveryServiceOptionsInto(in, out)
	return out
}
