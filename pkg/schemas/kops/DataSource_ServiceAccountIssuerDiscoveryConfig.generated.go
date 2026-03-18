package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceServiceAccountIssuerDiscoveryConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"discovery_store":          ComputedString(),
			"discovery_service":        ComputedStruct(DataSourceDiscoveryServiceOptions()),
			"enable_aws_oidc_provider": ComputedBool(),
			"additional_audiences":     ComputedList(String()),
		},
	}

	return res
}

func ExpandDataSourceServiceAccountIssuerDiscoveryConfig(in map[string]interface{}) kops.ServiceAccountIssuerDiscoveryConfig {
	if in == nil {
		panic("expand ServiceAccountIssuerDiscoveryConfig failure, in is nil")
	}
	return kops.ServiceAccountIssuerDiscoveryConfig{
		DiscoveryStore: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["discovery_store"]),
		DiscoveryService: func(in interface{}) *kops.DiscoveryServiceOptions {
			return func(in interface{}) *kops.DiscoveryServiceOptions {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.DiscoveryServiceOptions) *kops.DiscoveryServiceOptions {
					return &in
				}(func(in interface{}) kops.DiscoveryServiceOptions {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceDiscoveryServiceOptions(in[0].(map[string]interface{}))
					}
					return kops.DiscoveryServiceOptions{}
				}(in))
			}(in)
		}(in["discovery_service"]),
		EnableAWSOIDCProvider: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_aws_oidc_provider"]),
		AdditionalAudiences: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_audiences"]),
	}
}

func FlattenDataSourceServiceAccountIssuerDiscoveryConfigInto(in kops.ServiceAccountIssuerDiscoveryConfig, out map[string]interface{}) {
	out["discovery_store"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DiscoveryStore)
	out["discovery_service"] = func(in *kops.DiscoveryServiceOptions) interface{} {
		return func(in *kops.DiscoveryServiceOptions) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.DiscoveryServiceOptions) interface{} {
				return func(in kops.DiscoveryServiceOptions) []interface{} {
					return []interface{}{FlattenDataSourceDiscoveryServiceOptions(in)}
				}(in)
			}(*in)
		}(in)
	}(in.DiscoveryService)
	out["enable_aws_oidc_provider"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableAWSOIDCProvider)
	out["additional_audiences"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalAudiences)
}

func FlattenDataSourceServiceAccountIssuerDiscoveryConfig(in kops.ServiceAccountIssuerDiscoveryConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceServiceAccountIssuerDiscoveryConfigInto(in, out)
	return out
}
