package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	core "k8s.io/api/core/v1"
)

var _ = Schema

func ResourceSecretKeySelector() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":     OptionalString(),
			"key":      RequiredString(),
			"optional": OptionalBool(),
		},
	}

	return res
}

func ExpandResourceSecretKeySelector(in map[string]interface{}) core.SecretKeySelector {
	if in == nil {
		panic("expand SecretKeySelector failure, in is nil")
	}
	return core.SecretKeySelector{
		LocalObjectReference: func(in interface{}) core.LocalObjectReference {
			return ExpandResourceLocalObjectReference(in.(map[string]interface{}))
		}(in),
		Key: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["key"]),
		Optional: func(in interface{}) *bool {
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
		}(in["optional"]),
	}
}

func FlattenResourceSecretKeySelectorInto(in core.SecretKeySelector, out map[string]interface{}) {
	FlattenResourceLocalObjectReferenceInto(in.LocalObjectReference, out)
	out["key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Key)
	out["optional"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Optional)
}

func FlattenResourceSecretKeySelector(in core.SecretKeySelector) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceSecretKeySelectorInto(in, out)
	return out
}
