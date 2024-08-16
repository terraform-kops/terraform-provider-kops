package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceGCESpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"project":              OptionalString(),
			"service_account":      OptionalString(),
			"multizone":            OptionalBool(),
			"node_tags":            OptionalString(),
			"node_instance_prefix": OptionalString(),
			"pd_csi_driver":        OptionalStruct(ResourcePDCSIDriver()),
			"binaries_location":    OptionalString(),
		},
	}

	return res
}

func ExpandResourceGCESpec(in map[string]interface{}) kops.GCESpec {
	if in == nil {
		panic("expand GCESpec failure, in is nil")
	}
	return kops.GCESpec{
		Project: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["project"]),
		ServiceAccount: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["service_account"]),
		Multizone: func(in interface{}) *bool {
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
		}(in["multizone"]),
		NodeTags: func(in interface{}) *string {
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
		}(in["node_tags"]),
		NodeInstancePrefix: func(in interface{}) *string {
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
		}(in["node_instance_prefix"]),
		PDCSIDriver: func(in interface{}) *kops.PDCSIDriver {
			return func(in interface{}) *kops.PDCSIDriver {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.PDCSIDriver) *kops.PDCSIDriver {
					return &in
				}(func(in interface{}) kops.PDCSIDriver {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourcePDCSIDriver(in[0].(map[string]interface{}))
					}
					return kops.PDCSIDriver{}
				}(in))
			}(in)
		}(in["pd_csi_driver"]),
		BinariesLocation: func(in interface{}) *string {
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
		}(in["binaries_location"]),
	}
}

func FlattenResourceGCESpecInto(in kops.GCESpec, out map[string]interface{}) {
	out["project"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Project)
	out["service_account"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ServiceAccount)
	out["multizone"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Multizone)
	out["node_tags"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.NodeTags)
	out["node_instance_prefix"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.NodeInstancePrefix)
	out["pd_csi_driver"] = func(in *kops.PDCSIDriver) interface{} {
		return func(in *kops.PDCSIDriver) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.PDCSIDriver) interface{} {
				return func(in kops.PDCSIDriver) []interface{} {
					return []interface{}{FlattenResourcePDCSIDriver(in)}
				}(in)
			}(*in)
		}(in)
	}(in.PDCSIDriver)
	out["binaries_location"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.BinariesLocation)
}

func FlattenResourceGCESpec(in kops.GCESpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceGCESpecInto(in, out)
	return out
}
