package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceAuthenticationSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kopeio": OptionalStruct(ResourceKopeioAuthenticationSpec()),
			"aws":    OptionalStruct(ResourceAWSAuthenticationSpec()),
			"oidc":   OptionalStruct(ResourceOIDCAuthenticationSpec()),
		},
	}

	return res
}

func ExpandResourceAuthenticationSpec(in map[string]interface{}) kops.AuthenticationSpec {
	if in == nil {
		panic("expand AuthenticationSpec failure, in is nil")
	}
	return kops.AuthenticationSpec{
		Kopeio: func(in interface{}) *kops.KopeioAuthenticationSpec {
			return func(in interface{}) *kops.KopeioAuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KopeioAuthenticationSpec) *kops.KopeioAuthenticationSpec {
					return &in
				}(func(in interface{}) kops.KopeioAuthenticationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceKopeioAuthenticationSpec(in[0].(map[string]interface{}))
					}
					return kops.KopeioAuthenticationSpec{}
				}(in))
			}(in)
		}(in["kopeio"]),
		AWS: func(in interface{}) *kops.AWSAuthenticationSpec {
			return func(in interface{}) *kops.AWSAuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AWSAuthenticationSpec) *kops.AWSAuthenticationSpec {
					return &in
				}(func(in interface{}) kops.AWSAuthenticationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceAWSAuthenticationSpec(in[0].(map[string]interface{}))
					}
					return kops.AWSAuthenticationSpec{}
				}(in))
			}(in)
		}(in["aws"]),
		OIDC: func(in interface{}) *kops.OIDCAuthenticationSpec {
			return func(in interface{}) *kops.OIDCAuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.OIDCAuthenticationSpec) *kops.OIDCAuthenticationSpec {
					return &in
				}(func(in interface{}) kops.OIDCAuthenticationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandResourceOIDCAuthenticationSpec(in[0].(map[string]interface{}))
					}
					return kops.OIDCAuthenticationSpec{}
				}(in))
			}(in)
		}(in["oidc"]),
	}
}

func FlattenResourceAuthenticationSpecInto(in kops.AuthenticationSpec, out map[string]interface{}) {
	out["kopeio"] = func(in *kops.KopeioAuthenticationSpec) interface{} {
		return func(in *kops.KopeioAuthenticationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KopeioAuthenticationSpec) interface{} {
				return func(in kops.KopeioAuthenticationSpec) []interface{} {
					return []interface{}{FlattenResourceKopeioAuthenticationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kopeio)
	out["aws"] = func(in *kops.AWSAuthenticationSpec) interface{} {
		return func(in *kops.AWSAuthenticationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AWSAuthenticationSpec) interface{} {
				return func(in kops.AWSAuthenticationSpec) []interface{} {
					return []interface{}{FlattenResourceAWSAuthenticationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.AWS)
	out["oidc"] = func(in *kops.OIDCAuthenticationSpec) interface{} {
		return func(in *kops.OIDCAuthenticationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.OIDCAuthenticationSpec) interface{} {
				return func(in kops.OIDCAuthenticationSpec) []interface{} {
					return []interface{}{FlattenResourceOIDCAuthenticationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.OIDC)
}

func FlattenResourceAuthenticationSpec(in kops.AuthenticationSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceAuthenticationSpecInto(in, out)
	return out
}
