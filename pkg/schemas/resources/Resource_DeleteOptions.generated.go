package schemas

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-kops/terraform-provider-kops/pkg/api/resources"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Schema

func ResourceDeleteOptions() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"interval": OptionalDuration(),
			"wait":     OptionalDuration(),
			"count":    OptionalInt(),
		},
	}

	return res
}

func ExpandResourceDeleteOptions(in map[string]interface{}) resources.DeleteOptions {
	if in == nil {
		panic("expand DeleteOptions failure, in is nil")
	}
	return resources.DeleteOptions{
		Interval: func(in interface{}) *meta.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *meta.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.Duration) *meta.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["interval"]),
		Wait: func(in interface{}) *meta.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *meta.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.Duration) *meta.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["wait"]),
		Count: func(in interface{}) *int {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int) *int {
					return &in
				}(int(ExpandInt(in)))
			}(in)
		}(in["count"]),
	}
}

func FlattenResourceDeleteOptionsInto(in resources.DeleteOptions, out map[string]interface{}) {
	out["interval"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.Interval)
	out["wait"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.Wait)
	out["count"] = func(in *int) interface{} {
		return func(in *int) interface{} {
			if in == nil {
				return nil
			}
			return func(in int) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.Count)
}

func FlattenResourceDeleteOptions(in resources.DeleteOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceDeleteOptionsInto(in, out)
	return out
}
