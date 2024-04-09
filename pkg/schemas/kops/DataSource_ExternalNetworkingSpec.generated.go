package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceExternalNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceExternalNetworkingSpec(in map[string]interface{}) kops.ExternalNetworkingSpec {
	if in == nil {
		panic("expand ExternalNetworkingSpec failure, in is nil")
	}
	return kops.ExternalNetworkingSpec{}
}

func FlattenDataSourceExternalNetworkingSpecInto(in kops.ExternalNetworkingSpec, out map[string]interface{}) {
}

func FlattenDataSourceExternalNetworkingSpec(in kops.ExternalNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceExternalNetworkingSpecInto(in, out)
	return out
}
