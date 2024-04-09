package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceCloudConfiguration() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"manage_storage_classes": Nullable(OptionalBool()),
		},
	}

	return res
}

func ExpandResourceCloudConfiguration(in map[string]interface{}) kops.CloudConfiguration {
	if in == nil {
		panic("expand CloudConfiguration failure, in is nil")
	}
	return kops.CloudConfiguration{
		ManageStorageClasses: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if in, ok := in.([]interface{}); ok && len(in) == 1 {
				return func(in interface{}) *bool {
					return func(in interface{}) *bool {
						if in == nil {
							return nil
						}
						if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
							return nil
						}
						return func(in bool) *bool {
							return &in
						}(bool(ExpandBool(in)))
					}(in)
				}(in[0].(map[string]interface{})["value"])
			}
			return nil
		}(in["manage_storage_classes"]),
	}
}

func FlattenResourceCloudConfigurationInto(in kops.CloudConfiguration, out map[string]interface{}) {
	out["manage_storage_classes"] = func(in *bool) interface{} {
		if in == nil {
			return nil
		}
		return []interface{}{map[string]interface{}{"value": func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)}}
	}(in.ManageStorageClasses)
}

func FlattenResourceCloudConfiguration(in kops.CloudConfiguration) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceCloudConfigurationInto(in, out)
	return out
}
