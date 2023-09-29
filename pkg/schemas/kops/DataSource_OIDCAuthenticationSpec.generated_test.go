package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceOIDCAuthenticationSpec(t *testing.T) {
	_default := kops.OIDCAuthenticationSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OIDCAuthenticationSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"username_claim":  nil,
					"username_prefix": nil,
					"groups_claims":   func() []interface{} { return nil }(),
					"groups_prefix":   nil,
					"issuer_url":      nil,
					"client_id":       nil,
					"required_claims": func() map[string]interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceOIDCAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceOIDCAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOIDCAuthenticationSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"username_claim":  nil,
		"username_prefix": nil,
		"groups_claims":   func() []interface{} { return nil }(),
		"groups_prefix":   nil,
		"issuer_url":      nil,
		"client_id":       nil,
		"required_claims": func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kops.OIDCAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OIDCAuthenticationSpec{},
			},
			want: _default,
		},
		{
			name: "UsernameClaim - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.UsernameClaim = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UsernamePrefix - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.UsernamePrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GroupsClaims - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.GroupsClaims = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GroupsPrefix - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.GroupsPrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IssuerURL - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.IssuerURL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClientId - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.ClientID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequiredClaims - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.RequiredClaims = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceOIDCAuthenticationSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOIDCAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOIDCAuthenticationSpec(t *testing.T) {
	_default := map[string]interface{}{
		"username_claim":  nil,
		"username_prefix": nil,
		"groups_claims":   func() []interface{} { return nil }(),
		"groups_prefix":   nil,
		"issuer_url":      nil,
		"client_id":       nil,
		"required_claims": func() map[string]interface{} { return nil }(),
	}
	type args struct {
		in kops.OIDCAuthenticationSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OIDCAuthenticationSpec{},
			},
			want: _default,
		},
		{
			name: "UsernameClaim - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.UsernameClaim = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UsernamePrefix - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.UsernamePrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GroupsClaims - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.GroupsClaims = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GroupsPrefix - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.GroupsPrefix = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IssuerURL - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.IssuerURL = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClientId - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.ClientID = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RequiredClaims - default",
			args: args{
				in: func() kops.OIDCAuthenticationSpec {
					subject := kops.OIDCAuthenticationSpec{}
					subject.RequiredClaims = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceOIDCAuthenticationSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOIDCAuthenticationSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
