package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceGCPNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceGCPNetworkingSpec(in map[string]interface{}) kops.GCPNetworkingSpec {
	if in == nil {
		panic("expand GCPNetworkingSpec failure, in is nil")
	}
	return kops.GCPNetworkingSpec{}
}

func FlattenResourceGCPNetworkingSpecInto(in kops.GCPNetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceGCPNetworkingSpec(in kops.GCPNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceGCPNetworkingSpecInto(in, out)
	return out
}
