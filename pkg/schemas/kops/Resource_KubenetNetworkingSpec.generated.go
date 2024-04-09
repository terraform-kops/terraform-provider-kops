package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKubenetNetworkingSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceKubenetNetworkingSpec(in map[string]interface{}) kops.KubenetNetworkingSpec {
	if in == nil {
		panic("expand KubenetNetworkingSpec failure, in is nil")
	}
	return kops.KubenetNetworkingSpec{}
}

func FlattenResourceKubenetNetworkingSpecInto(in kops.KubenetNetworkingSpec, out map[string]interface{}) {
}

func FlattenResourceKubenetNetworkingSpec(in kops.KubenetNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKubenetNetworkingSpecInto(in, out)
	return out
}
