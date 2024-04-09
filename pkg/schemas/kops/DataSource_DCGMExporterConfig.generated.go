package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceDCGMExporterConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": ComputedBool(),
		},
	}

	return res
}

func ExpandDataSourceDCGMExporterConfig(in map[string]interface{}) kops.DCGMExporterConfig {
	if in == nil {
		panic("expand DCGMExporterConfig failure, in is nil")
	}
	return kops.DCGMExporterConfig{
		Enabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enabled"]),
	}
}

func FlattenDataSourceDCGMExporterConfigInto(in kops.DCGMExporterConfig, out map[string]interface{}) {
	out["enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Enabled)
}

func FlattenDataSourceDCGMExporterConfig(in kops.DCGMExporterConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceDCGMExporterConfigInto(in, out)
	return out
}
