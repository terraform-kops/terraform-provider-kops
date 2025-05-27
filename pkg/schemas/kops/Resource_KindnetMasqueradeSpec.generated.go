package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceKindnetMasqueradeSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled":              OptionalBool(),
			"non_masquerade_cidrs": OptionalList(String()),
		},
	}

	return res
}

func ExpandResourceKindnetMasqueradeSpec(in map[string]interface{}) kops.KindnetMasqueradeSpec {
	if in == nil {
		panic("expand KindnetMasqueradeSpec failure, in is nil")
	}
	return kops.KindnetMasqueradeSpec{
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
		NonMasqueradeCIDRs: func(in interface{}) []string {
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
		}(in["non_masquerade_cidrs"]),
	}
}

func FlattenResourceKindnetMasqueradeSpecInto(in kops.KindnetMasqueradeSpec, out map[string]interface{}) {
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
	out["non_masquerade_cidrs"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.NonMasqueradeCIDRs)
}

func FlattenResourceKindnetMasqueradeSpec(in kops.KindnetMasqueradeSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceKindnetMasqueradeSpecInto(in, out)
	return out
}
