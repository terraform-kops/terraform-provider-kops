package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceDNSControllerGossipConfig(t *testing.T) {
	_default := kops.DNSControllerGossipConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.DNSControllerGossipConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"protocol":  nil,
					"listen":    nil,
					"secret":    nil,
					"secondary": nil,
					"seed":      nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceDNSControllerGossipConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceDNSControllerGossipConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceDNSControllerGossipConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"protocol":  nil,
		"listen":    nil,
		"secret":    nil,
		"secondary": nil,
		"seed":      nil,
	}
	type args struct {
		in kops.DNSControllerGossipConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.DNSControllerGossipConfig{},
			},
			want: _default,
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secondary - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Secondary = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Seed - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Seed = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceDNSControllerGossipConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceDNSControllerGossipConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceDNSControllerGossipConfig(t *testing.T) {
	_default := map[string]interface{}{
		"protocol":  nil,
		"listen":    nil,
		"secret":    nil,
		"secondary": nil,
		"seed":      nil,
	}
	type args struct {
		in kops.DNSControllerGossipConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.DNSControllerGossipConfig{},
			},
			want: _default,
		},
		{
			name: "Protocol - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Protocol = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Listen - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Listen = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secret - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Secret = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Secondary - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Secondary = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Seed - default",
			args: args{
				in: func() kops.DNSControllerGossipConfig {
					subject := kops.DNSControllerGossipConfig{}
					subject.Seed = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceDNSControllerGossipConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceDNSControllerGossipConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
