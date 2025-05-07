package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceKindnetNetworkingSpec(t *testing.T) {
	_default := kops.KindnetNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KindnetNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"version":                         "",
					"network_policies":                nil,
					"admin_network_policies":          nil,
					"baseline_admin_network_policies": nil,
					"dns_caching":                     nil,
					"nat64":                           nil,
					"fast_path_threshold":             nil,
					"masquerade":                      nil,
					"log_level":                       nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceKindnetNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceKindnetNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKindnetNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"version":                         "",
		"network_policies":                nil,
		"admin_network_policies":          nil,
		"baseline_admin_network_policies": nil,
		"dns_caching":                     nil,
		"nat64":                           nil,
		"fast_path_threshold":             nil,
		"masquerade":                      nil,
		"log_level":                       nil,
	}
	type args struct {
		in kops.KindnetNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KindnetNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkPolicies - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.NetworkPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdminNetworkPolicies - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.AdminNetworkPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BaselineAdminNetworkPolicies - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.BaselineAdminNetworkPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DnsCaching - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.DNSCaching = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NAT64 - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.NAT64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FastPathThreshold - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.FastPathThreshold = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Masquerade - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.Masquerade = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceKindnetNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKindnetNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceKindnetNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"version":                         "",
		"network_policies":                nil,
		"admin_network_policies":          nil,
		"baseline_admin_network_policies": nil,
		"dns_caching":                     nil,
		"nat64":                           nil,
		"fast_path_threshold":             nil,
		"masquerade":                      nil,
		"log_level":                       nil,
	}
	type args struct {
		in kops.KindnetNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KindnetNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NetworkPolicies - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.NetworkPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AdminNetworkPolicies - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.AdminNetworkPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BaselineAdminNetworkPolicies - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.BaselineAdminNetworkPolicies = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DnsCaching - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.DNSCaching = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NAT64 - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.NAT64 = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FastPathThreshold - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.FastPathThreshold = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Masquerade - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.Masquerade = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KindnetNetworkingSpec {
					subject := kops.KindnetNetworkingSpec{}
					subject.LogLevel = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceKindnetNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceKindnetNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
