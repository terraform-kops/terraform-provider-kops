package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-kops/terraform-provider-kops/pkg/api/resources"
	"github.com/terraform-kops/terraform-provider-kops/pkg/api/utils"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	utilsschemas "github.com/terraform-kops/terraform-provider-kops/pkg/schemas/utils"
)

var _ = Schema

func ResourceRollingUpdateOptions() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"skip":                    OptionalBool(),
			"exclude_instance_groups": OptionalList(String()),
			"master_interval":         OptionalDuration(),
			"node_interval":           OptionalDuration(),
			"bastion_interval":        OptionalDuration(),
			"fail_on_drain_error":     OptionalBool(),
			"fail_on_validate":        OptionalBool(),
			"post_drain_delay":        OptionalDuration(),
			"validation_timeout":      OptionalDuration(),
			"validate_count":          OptionalInt(),
			"cloud_only":              OptionalBool(),
			"force":                   OptionalBool(),
		},
	}

	return res
}

func ExpandResourceRollingUpdateOptions(in map[string]interface{}) resources.RollingUpdateOptions {
	if in == nil {
		panic("expand RollingUpdateOptions failure, in is nil")
	}
	return resources.RollingUpdateOptions{
		Skip: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["skip"]),
		ExcludeInstanceGroups: func(in interface{}) []string {
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
		}(in["exclude_instance_groups"]),
		RollingUpdateOptions: func(in interface{}) utils.RollingUpdateOptions {
			return utilsschemas.ExpandResourceRollingUpdateOptions(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenResourceRollingUpdateOptionsInto(in resources.RollingUpdateOptions, out map[string]interface{}) {
	out["skip"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Skip)
	out["exclude_instance_groups"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.ExcludeInstanceGroups)
	utilsschemas.FlattenResourceRollingUpdateOptionsInto(in.RollingUpdateOptions, out)
}

func FlattenResourceRollingUpdateOptions(in resources.RollingUpdateOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceRollingUpdateOptionsInto(in, out)
	return out
}
