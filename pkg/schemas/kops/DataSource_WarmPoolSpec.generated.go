package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceWarmPoolSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"min_size":               ComputedInt(),
			"max_size":               ComputedInt(),
			"enable_lifecycle_hook":  ComputedBool(),
			"lifecycle_hook_timeout": ComputedInt(),
			"additional_images":      ComputedList(String()),
		},
	}

	return res
}

func ExpandDataSourceWarmPoolSpec(in map[string]interface{}) kops.WarmPoolSpec {
	if in == nil {
		panic("expand WarmPoolSpec failure, in is nil")
	}
	return kops.WarmPoolSpec{
		MinSize: func(in interface{}) int64 {
			return int64(ExpandInt(in))
		}(in["min_size"]),
		MaxSize: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["max_size"]),
		EnableLifecycleHook: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_lifecycle_hook"]),
		LifecycleHookTimeout: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["lifecycle_hook_timeout"]),
		AdditionalImages: func(in interface{}) []string {
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
		}(in["additional_images"]),
	}
}

func FlattenDataSourceWarmPoolSpecInto(in kops.WarmPoolSpec, out map[string]interface{}) {
	out["min_size"] = func(in int64) interface{} {
		return FlattenInt(int(in))
	}(in.MinSize)
	out["max_size"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MaxSize)
	out["enable_lifecycle_hook"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableLifecycleHook)
	out["lifecycle_hook_timeout"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.LifecycleHookTimeout)
	out["additional_images"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalImages)
}

func FlattenDataSourceWarmPoolSpec(in kops.WarmPoolSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceWarmPoolSpecInto(in, out)
	return out
}
