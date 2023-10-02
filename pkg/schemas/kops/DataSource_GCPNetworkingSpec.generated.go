package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceGCPNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceGCPNetworkingSpec(in map[string]interface{}) kops.GCPNetworkingSpec {
	if in == nil {
		panic("expand GCPNetworkingSpec failure, in is nil")
	}
	return kops.GCPNetworkingSpec{}
}

func FlattenDataSourceGCPNetworkingSpecInto(in kops.GCPNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceGCPNetworkingSpec(in kops.GCPNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceGCPNetworkingSpecInto(in, out)
	return out
}
