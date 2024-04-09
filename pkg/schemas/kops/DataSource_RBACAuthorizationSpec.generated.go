package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceRBACAuthorizationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}

	return res
}

func ExpandDataSourceRBACAuthorizationSpec(in map[string]interface{}) kops.RBACAuthorizationSpec {
	if in == nil {
		panic("expand RBACAuthorizationSpec failure, in is nil")
	}
	return kops.RBACAuthorizationSpec{}
}

func FlattenDataSourceRBACAuthorizationSpecInto(in kops.RBACAuthorizationSpec, out map[string]interface{}) {
}

func FlattenDataSourceRBACAuthorizationSpec(in kops.RBACAuthorizationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceRBACAuthorizationSpecInto(in, out)
	return out
}
