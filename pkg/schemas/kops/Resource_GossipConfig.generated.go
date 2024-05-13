package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceGossipConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol":  OptionalString(),
			"listen":    OptionalString(),
			"secret":    OptionalString(),
			"secondary": OptionalStruct(ResourceGossipConfigSecondary()),
		},
	}

	return res
}

func ExpandResourceGossipConfig(in map[string]interface{}) kops.GossipConfig {
	if in == nil {
		panic("expand GossipConfig failure, in is nil")
	}
	return kops.GossipConfig{
		Protocol: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["protocol"]),
		Listen: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["listen"]),
		Secret: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["secret"]),
		Secondary: func(in interface{}) *kops.GossipConfigSecondary {
			return func(in interface{}) *kops.GossipConfigSecondary {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.GossipConfigSecondary) *kops.GossipConfigSecondary {
					return &in
				}(func(in interface{}) kops.GossipConfigSecondary {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceGossipConfigSecondary(in[0].(map[string]interface{}))
					}
					return kops.GossipConfigSecondary{}
				}(in))
			}(in)
		}(in["secondary"]),
	}
}

func FlattenResourceGossipConfigInto(in kops.GossipConfig, out map[string]interface{}) {
	out["protocol"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Protocol)
	out["listen"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Listen)
	out["secret"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Secret)
	out["secondary"] = func(in *kops.GossipConfigSecondary) interface{} {
		return func(in *kops.GossipConfigSecondary) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.GossipConfigSecondary) interface{} {
				return func(in kops.GossipConfigSecondary) []interface{} {
					return []interface{}{FlattenResourceGossipConfigSecondary(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Secondary)
}

func FlattenResourceGossipConfig(in kops.GossipConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceGossipConfigInto(in, out)
	return out
}
