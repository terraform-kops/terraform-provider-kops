package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKarpenterConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":      OptionalBool(),
			"log_encoding": OptionalString(),
			"log_level":    OptionalString(),
			"image":        OptionalString(),
		},
	}

	return res
}

func ExpandResourceKarpenterConfig(in map[string]interface{}) kops.KarpenterConfig {
	if in == nil {
		panic("expand KarpenterConfig failure, in is nil")
	}
	return kops.KarpenterConfig{
		Enabled: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enabled"]),
		LogEncoding: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["log_encoding"]),
		LogLevel: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["log_level"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
	}
}

func FlattenResourceKarpenterConfigInto(in kops.KarpenterConfig, out map[string]interface{}) {
	out["enabled"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Enabled)
	out["log_encoding"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.LogEncoding)
	out["log_level"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.LogLevel)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
}

func FlattenResourceKarpenterConfig(in kops.KarpenterConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKarpenterConfigInto(in, out)
	return out
}
