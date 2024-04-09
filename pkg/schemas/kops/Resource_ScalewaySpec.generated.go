package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceScalewaySpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandResourceScalewaySpec(in map[string]interface{}) kops.ScalewaySpec {
	if in == nil {
		panic("expand ScalewaySpec failure, in is nil")
	}
	return kops.ScalewaySpec{}
}

func FlattenResourceScalewaySpecInto(in kops.ScalewaySpec, out map[string]interface{}) {
}

func FlattenResourceScalewaySpec(in kops.ScalewaySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceScalewaySpecInto(in, out)
	return out
}
