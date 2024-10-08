package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceKubeSchedulerConfig(t *testing.T) {
	_default := kops.KubeSchedulerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KubeSchedulerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"master":                           "",
					"log_format":                       "",
					"log_level":                        0,
					"image":                            "",
					"leader_election":                  nil,
					"use_policy_config_map":            nil,
					"feature_gates":                    func() map[string]interface{} { return nil }(),
					"max_persistent_volumes":           nil,
					"qps":                              nil,
					"burst":                            0,
					"kube_api_qps":                     nil,
					"kube_api_burst":                   nil,
					"authentication_kubeconfig":        "",
					"authorization_kubeconfig":         "",
					"authorization_always_allow_paths": func() []interface{} { return nil }(),
					"enable_profiling":                 nil,
					"enable_contention_profiling":      nil,
					"tls_cert_file":                    nil,
					"tls_private_key_file":             "",
					"cpu_request":                      nil,
					"cpu_limit":                        nil,
					"memory_request":                   nil,
					"memory_limit":                     nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceKubeSchedulerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceKubeSchedulerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKubeSchedulerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"master":                           "",
		"log_format":                       "",
		"log_level":                        0,
		"image":                            "",
		"leader_election":                  nil,
		"use_policy_config_map":            nil,
		"feature_gates":                    func() map[string]interface{} { return nil }(),
		"max_persistent_volumes":           nil,
		"qps":                              nil,
		"burst":                            0,
		"kube_api_qps":                     nil,
		"kube_api_burst":                   nil,
		"authentication_kubeconfig":        "",
		"authorization_kubeconfig":         "",
		"authorization_always_allow_paths": func() []interface{} { return nil }(),
		"enable_profiling":                 nil,
		"enable_contention_profiling":      nil,
		"tls_cert_file":                    nil,
		"tls_private_key_file":             "",
		"cpu_request":                      nil,
		"cpu_limit":                        nil,
		"memory_request":                   nil,
		"memory_limit":                     nil,
	}
	type args struct {
		in kops.KubeSchedulerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KubeSchedulerConfig{},
			},
			want: _default,
		},
		{
			name: "Master - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogFormat - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.LogFormat = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LeaderElection - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.LeaderElection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UsePolicyConfigMap - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.UsePolicyConfigMap = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxPersistentVolumes - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.MaxPersistentVolumes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Qps - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Qps = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Burst - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Burst = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiQPS - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.KubeAPIQPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiBurst - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.KubeAPIBurst = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationKubeconfig - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.AuthenticationKubeconfig = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationKubeconfig - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.AuthorizationKubeconfig = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationAlwaysAllowPaths - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.AuthorizationAlwaysAllowPaths = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableProfiling - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.EnableProfiling = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableContentionProfiling - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.EnableContentionProfiling = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCertFile - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.TLSCertFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSPrivateKeyFile - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.TLSPrivateKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceKubeSchedulerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeSchedulerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKubeSchedulerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"master":                           "",
		"log_format":                       "",
		"log_level":                        0,
		"image":                            "",
		"leader_election":                  nil,
		"use_policy_config_map":            nil,
		"feature_gates":                    func() map[string]interface{} { return nil }(),
		"max_persistent_volumes":           nil,
		"qps":                              nil,
		"burst":                            0,
		"kube_api_qps":                     nil,
		"kube_api_burst":                   nil,
		"authentication_kubeconfig":        "",
		"authorization_kubeconfig":         "",
		"authorization_always_allow_paths": func() []interface{} { return nil }(),
		"enable_profiling":                 nil,
		"enable_contention_profiling":      nil,
		"tls_cert_file":                    nil,
		"tls_private_key_file":             "",
		"cpu_request":                      nil,
		"cpu_limit":                        nil,
		"memory_request":                   nil,
		"memory_limit":                     nil,
	}
	type args struct {
		in kops.KubeSchedulerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KubeSchedulerConfig{},
			},
			want: _default,
		},
		{
			name: "Master - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogFormat - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.LogFormat = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LeaderElection - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.LeaderElection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UsePolicyConfigMap - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.UsePolicyConfigMap = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MaxPersistentVolumes - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.MaxPersistentVolumes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Qps - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Qps = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Burst - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.Burst = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiQPS - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.KubeAPIQPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiBurst - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.KubeAPIBurst = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationKubeconfig - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.AuthenticationKubeconfig = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationKubeconfig - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.AuthorizationKubeconfig = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationAlwaysAllowPaths - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.AuthorizationAlwaysAllowPaths = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableProfiling - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.EnableProfiling = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableContentionProfiling - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.EnableContentionProfiling = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCertFile - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.TLSCertFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSPrivateKeyFile - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.TLSPrivateKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.KubeSchedulerConfig {
					subject := kops.KubeSchedulerConfig{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceKubeSchedulerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeSchedulerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
