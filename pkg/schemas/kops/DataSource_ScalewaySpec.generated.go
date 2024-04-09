package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceScalewaySpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceScalewaySpec(in map[string]interface{}) kops.ScalewaySpec {
	if in == nil {
		panic("expand ScalewaySpec failure, in is nil")
	}
	return kops.ScalewaySpec{}
}

func FlattenDataSourceScalewaySpecInto(in kops.ScalewaySpec, out map[string]interface{}) {
}

func FlattenDataSourceScalewaySpec(in kops.ScalewaySpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceScalewaySpecInto(in, out)
	return out
}
