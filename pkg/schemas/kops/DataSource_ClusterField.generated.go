package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceClusterField() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": ComputedString(),
		},
	}

	return res
}

func ExpandDataSourceClusterField(in map[string]interface{}) kops.ClusterField {
	if in == nil {
		panic("expand ClusterField failure, in is nil")
	}
	return kops.ClusterField{
		Path: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["path"]),
	}
}

func FlattenDataSourceClusterFieldInto(in kops.ClusterField, out map[string]interface{}) {
	out["path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Path)
}

func FlattenDataSourceClusterField(in kops.ClusterField) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceClusterFieldInto(in, out)
	return out
}
