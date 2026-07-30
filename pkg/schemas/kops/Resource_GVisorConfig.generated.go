package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceGVisorConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":  OptionalBool(),
			"platform": OptionalString(),
		},
	}

	return res
}

func ExpandResourceGVisorConfig(in map[string]interface{}) kops.GVisorConfig {
	if in == nil {
		panic("expand GVisorConfig failure, in is nil")
	}
	return kops.GVisorConfig{
		Enabled: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
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
		}(in["enabled"]),
		Platform: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["platform"]),
	}
}

func FlattenResourceGVisorConfigInto(in kops.GVisorConfig, out map[string]interface{}) {
	out["enabled"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Enabled)
	out["platform"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Platform)
}

func FlattenResourceGVisorConfig(in kops.GVisorConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceGVisorConfigInto(in, out)
	return out
}
