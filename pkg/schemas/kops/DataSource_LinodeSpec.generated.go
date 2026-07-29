package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceLinodeSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceLinodeSpec(in map[string]interface{}) kops.LinodeSpec {
	if in == nil {
		panic("expand LinodeSpec failure, in is nil")
	}
	return kops.LinodeSpec{}
}

func FlattenDataSourceLinodeSpecInto(in kops.LinodeSpec, out map[string]interface{}) {
}

func FlattenDataSourceLinodeSpec(in kops.LinodeSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceLinodeSpecInto(in, out)
	return out
}
