package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceKubeControllerManagerConfig(t *testing.T) {
	_default := kops.KubeControllerManagerConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.KubeControllerManagerConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"master":                                    "",
					"log_format":                                "",
					"log_level":                                 0,
					"service_account_private_key_file":          "",
					"image":                                     "",
					"cloud_provider":                            "",
					"cluster_name":                              "",
					"cluster_cidr":                              "",
					"allocate_node_cidrs":                       nil,
					"node_cidr_mask_size":                       nil,
					"configure_cloud_routes":                    nil,
					"controllers":                               func() []interface{} { return nil }(),
					"cidr_allocator_type":                       nil,
					"root_ca_file":                              "",
					"leader_election":                           nil,
					"attach_detach_reconcile_sync_period":       nil,
					"disable_attach_detach_reconcile_sync":      nil,
					"terminated_pod_gc_threshold":               nil,
					"node_monitor_period":                       nil,
					"node_monitor_grace_period":                 nil,
					"pod_eviction_timeout":                      nil,
					"use_service_account_credentials":           nil,
					"horizontal_pod_autoscaler_sync_period":     nil,
					"horizontal_pod_autoscaler_downscale_delay": nil,
					"horizontal_pod_autoscaler_downscale_stabilization":   nil,
					"horizontal_pod_autoscaler_upscale_delay":             nil,
					"horizontal_pod_autoscaler_initial_readiness_delay":   nil,
					"horizontal_pod_autoscaler_cpu_initialization_period": nil,
					"horizontal_pod_autoscaler_tolerance":                 nil,
					"horizontal_pod_autoscaler_use_rest_clients":          nil,
					"experimental_cluster_signing_duration":               nil,
					"cluster_signing_duration":                            nil,
					"feature_gates":                                       func() map[string]interface{} { return nil }(),
					"tls_cert_file":                                       nil,
					"tls_cipher_suites":                                   func() []interface{} { return nil }(),
					"tls_min_version":                                     "",
					"tls_private_key_file":                                "",
					"min_resync_period":                                   "",
					"kube_api_qps":                                        nil,
					"kube_api_burst":                                      nil,
					"concurrent_deployment_syncs":                         nil,
					"concurrent_endpoint_syncs":                           nil,
					"concurrent_namespace_syncs":                          nil,
					"concurrent_replicaset_syncs":                         nil,
					"concurrent_service_syncs":                            nil,
					"concurrent_resource_quota_syncs":                     nil,
					"concurrent_serviceaccount_token_syncs":               nil,
					"concurrent_rc_syncs":                                 nil,
					"concurrent_horizontal_pod_austoscaler_syncs":         nil,
					"concurrent_job_syncs":                                nil,
					"authentication_kubeconfig":                           "",
					"authorization_kubeconfig":                            "",
					"authorization_always_allow_paths":                    func() []interface{} { return nil }(),
					"external_cloud_volume_plugin":                        "",
					"endpoint_updates_batch_period":                       nil,
					"endpoint_slice_updates_batch_period":                 nil,
					"enable_profiling":                                    nil,
					"enable_contention_profiling":                         nil,
					"enable_leader_migration":                             nil,
					"cpu_request":                                         nil,
					"cpu_limit":                                           nil,
					"memory_request":                                      nil,
					"memory_limit":                                        nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceKubeControllerManagerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceKubeControllerManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKubeControllerManagerConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"master":                                    "",
		"log_format":                                "",
		"log_level":                                 0,
		"service_account_private_key_file":          "",
		"image":                                     "",
		"cloud_provider":                            "",
		"cluster_name":                              "",
		"cluster_cidr":                              "",
		"allocate_node_cidrs":                       nil,
		"node_cidr_mask_size":                       nil,
		"configure_cloud_routes":                    nil,
		"controllers":                               func() []interface{} { return nil }(),
		"cidr_allocator_type":                       nil,
		"root_ca_file":                              "",
		"leader_election":                           nil,
		"attach_detach_reconcile_sync_period":       nil,
		"disable_attach_detach_reconcile_sync":      nil,
		"terminated_pod_gc_threshold":               nil,
		"node_monitor_period":                       nil,
		"node_monitor_grace_period":                 nil,
		"pod_eviction_timeout":                      nil,
		"use_service_account_credentials":           nil,
		"horizontal_pod_autoscaler_sync_period":     nil,
		"horizontal_pod_autoscaler_downscale_delay": nil,
		"horizontal_pod_autoscaler_downscale_stabilization":   nil,
		"horizontal_pod_autoscaler_upscale_delay":             nil,
		"horizontal_pod_autoscaler_initial_readiness_delay":   nil,
		"horizontal_pod_autoscaler_cpu_initialization_period": nil,
		"horizontal_pod_autoscaler_tolerance":                 nil,
		"horizontal_pod_autoscaler_use_rest_clients":          nil,
		"experimental_cluster_signing_duration":               nil,
		"cluster_signing_duration":                            nil,
		"feature_gates":                                       func() map[string]interface{} { return nil }(),
		"tls_cert_file":                                       nil,
		"tls_cipher_suites":                                   func() []interface{} { return nil }(),
		"tls_min_version":                                     "",
		"tls_private_key_file":                                "",
		"min_resync_period":                                   "",
		"kube_api_qps":                                        nil,
		"kube_api_burst":                                      nil,
		"concurrent_deployment_syncs":                         nil,
		"concurrent_endpoint_syncs":                           nil,
		"concurrent_namespace_syncs":                          nil,
		"concurrent_replicaset_syncs":                         nil,
		"concurrent_service_syncs":                            nil,
		"concurrent_resource_quota_syncs":                     nil,
		"concurrent_serviceaccount_token_syncs":               nil,
		"concurrent_rc_syncs":                                 nil,
		"concurrent_horizontal_pod_austoscaler_syncs":         nil,
		"concurrent_job_syncs":                                nil,
		"authentication_kubeconfig":                           "",
		"authorization_kubeconfig":                            "",
		"authorization_always_allow_paths":                    func() []interface{} { return nil }(),
		"external_cloud_volume_plugin":                        "",
		"endpoint_updates_batch_period":                       nil,
		"endpoint_slice_updates_batch_period":                 nil,
		"enable_profiling":                                    nil,
		"enable_contention_profiling":                         nil,
		"enable_leader_migration":                             nil,
		"cpu_request":                                         nil,
		"cpu_limit":                                           nil,
		"memory_request":                                      nil,
		"memory_limit":                                        nil,
	}
	type args struct {
		in kops.KubeControllerManagerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KubeControllerManagerConfig{},
			},
			want: _default,
		},
		{
			name: "Master - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogFormat - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.LogFormat = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountPrivateKeyFile - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ServiceAccountPrivateKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.CloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterCidr - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ClusterCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllocateNodeCidrs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.AllocateNodeCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeCidrMaskSize - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.NodeCIDRMaskSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigureCloudRoutes - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConfigureCloudRoutes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Controllers - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.Controllers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CidrAllocatorType - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.CIDRAllocatorType = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootCaFile - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.RootCAFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LeaderElection - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.LeaderElection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AttachDetachReconcileSyncPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.AttachDetachReconcileSyncPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableAttachDetachReconcileSync - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.DisableAttachDetachReconcileSync = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TerminatedPodGCThreshold - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.TerminatedPodGCThreshold = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeMonitorPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.NodeMonitorPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeMonitorGracePeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.NodeMonitorGracePeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodEvictionTimeout - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.PodEvictionTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseServiceAccountCredentials - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.UseServiceAccountCredentials = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerSyncPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerSyncPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerDownscaleDelay - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerDownscaleDelay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerDownscaleStabilization - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerDownscaleStabilization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerUpscaleDelay - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerUpscaleDelay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerInitialReadinessDelay - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerInitialReadinessDelay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerCpuInitializationPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerCPUInitializationPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerTolerance - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerTolerance = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerUseRestClients - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerUseRestClients = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExperimentalClusterSigningDuration - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ExperimentalClusterSigningDuration = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterSigningDuration - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ClusterSigningDuration = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCertFile - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.TLSCertFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCipherSuites - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.TLSCipherSuites = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSMinVersion - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.TLSMinVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSPrivateKeyFile - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.TLSPrivateKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MinResyncPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.MinResyncPeriod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiQPS - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.KubeAPIQPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiBurst - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.KubeAPIBurst = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentDeploymentSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentDeploymentSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentEndpointSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentEndpointSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentNamespaceSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentNamespaceSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentReplicasetSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentReplicasetSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentServiceSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentServiceSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentResourceQuotaSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentResourceQuotaSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentServiceaccountTokenSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentServiceaccountTokenSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentRCSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentRCSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentHorizontalPodAustoscalerSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentHorizontalPodAustoscalerSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentJobSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentJobSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationKubeconfig - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.AuthenticationKubeconfig = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationKubeconfig - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.AuthorizationKubeconfig = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationAlwaysAllowPaths - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.AuthorizationAlwaysAllowPaths = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalCloudVolumePlugin - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ExternalCloudVolumePlugin = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EndpointUpdatesBatchPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.EndpointUpdatesBatchPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EndpointSliceUpdatesBatchPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.EndpointSliceUpdatesBatchPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableProfiling - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.EnableProfiling = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableContentionProfiling - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.EnableContentionProfiling = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableLeaderMigration - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.EnableLeaderMigration = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
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
			FlattenResourceKubeControllerManagerConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeControllerManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceKubeControllerManagerConfig(t *testing.T) {
	_default := map[string]interface{}{
		"master":                                    "",
		"log_format":                                "",
		"log_level":                                 0,
		"service_account_private_key_file":          "",
		"image":                                     "",
		"cloud_provider":                            "",
		"cluster_name":                              "",
		"cluster_cidr":                              "",
		"allocate_node_cidrs":                       nil,
		"node_cidr_mask_size":                       nil,
		"configure_cloud_routes":                    nil,
		"controllers":                               func() []interface{} { return nil }(),
		"cidr_allocator_type":                       nil,
		"root_ca_file":                              "",
		"leader_election":                           nil,
		"attach_detach_reconcile_sync_period":       nil,
		"disable_attach_detach_reconcile_sync":      nil,
		"terminated_pod_gc_threshold":               nil,
		"node_monitor_period":                       nil,
		"node_monitor_grace_period":                 nil,
		"pod_eviction_timeout":                      nil,
		"use_service_account_credentials":           nil,
		"horizontal_pod_autoscaler_sync_period":     nil,
		"horizontal_pod_autoscaler_downscale_delay": nil,
		"horizontal_pod_autoscaler_downscale_stabilization":   nil,
		"horizontal_pod_autoscaler_upscale_delay":             nil,
		"horizontal_pod_autoscaler_initial_readiness_delay":   nil,
		"horizontal_pod_autoscaler_cpu_initialization_period": nil,
		"horizontal_pod_autoscaler_tolerance":                 nil,
		"horizontal_pod_autoscaler_use_rest_clients":          nil,
		"experimental_cluster_signing_duration":               nil,
		"cluster_signing_duration":                            nil,
		"feature_gates":                                       func() map[string]interface{} { return nil }(),
		"tls_cert_file":                                       nil,
		"tls_cipher_suites":                                   func() []interface{} { return nil }(),
		"tls_min_version":                                     "",
		"tls_private_key_file":                                "",
		"min_resync_period":                                   "",
		"kube_api_qps":                                        nil,
		"kube_api_burst":                                      nil,
		"concurrent_deployment_syncs":                         nil,
		"concurrent_endpoint_syncs":                           nil,
		"concurrent_namespace_syncs":                          nil,
		"concurrent_replicaset_syncs":                         nil,
		"concurrent_service_syncs":                            nil,
		"concurrent_resource_quota_syncs":                     nil,
		"concurrent_serviceaccount_token_syncs":               nil,
		"concurrent_rc_syncs":                                 nil,
		"concurrent_horizontal_pod_austoscaler_syncs":         nil,
		"concurrent_job_syncs":                                nil,
		"authentication_kubeconfig":                           "",
		"authorization_kubeconfig":                            "",
		"authorization_always_allow_paths":                    func() []interface{} { return nil }(),
		"external_cloud_volume_plugin":                        "",
		"endpoint_updates_batch_period":                       nil,
		"endpoint_slice_updates_batch_period":                 nil,
		"enable_profiling":                                    nil,
		"enable_contention_profiling":                         nil,
		"enable_leader_migration":                             nil,
		"cpu_request":                                         nil,
		"cpu_limit":                                           nil,
		"memory_request":                                      nil,
		"memory_limit":                                        nil,
	}
	type args struct {
		in kops.KubeControllerManagerConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.KubeControllerManagerConfig{},
			},
			want: _default,
		},
		{
			name: "Master - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.Master = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogFormat - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.LogFormat = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LogLevel - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.LogLevel = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ServiceAccountPrivateKeyFile - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ServiceAccountPrivateKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.Image = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CloudProvider - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.CloudProvider = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterCidr - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ClusterCIDR = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AllocateNodeCidrs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.AllocateNodeCIDRs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeCidrMaskSize - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.NodeCIDRMaskSize = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConfigureCloudRoutes - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConfigureCloudRoutes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Controllers - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.Controllers = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CidrAllocatorType - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.CIDRAllocatorType = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "RootCaFile - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.RootCAFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "LeaderElection - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.LeaderElection = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AttachDetachReconcileSyncPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.AttachDetachReconcileSyncPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableAttachDetachReconcileSync - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.DisableAttachDetachReconcileSync = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TerminatedPodGCThreshold - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.TerminatedPodGCThreshold = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeMonitorPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.NodeMonitorPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeMonitorGracePeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.NodeMonitorGracePeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PodEvictionTimeout - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.PodEvictionTimeout = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "UseServiceAccountCredentials - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.UseServiceAccountCredentials = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerSyncPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerSyncPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerDownscaleDelay - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerDownscaleDelay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerDownscaleStabilization - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerDownscaleStabilization = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerUpscaleDelay - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerUpscaleDelay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerInitialReadinessDelay - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerInitialReadinessDelay = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerCpuInitializationPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerCPUInitializationPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerTolerance - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerTolerance = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "HorizontalPodAutoscalerUseRestClients - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.HorizontalPodAutoscalerUseRestClients = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExperimentalClusterSigningDuration - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ExperimentalClusterSigningDuration = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterSigningDuration - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ClusterSigningDuration = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "FeatureGates - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.FeatureGates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCertFile - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.TLSCertFile = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSCipherSuites - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.TLSCipherSuites = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSMinVersion - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.TLSMinVersion = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "TLSPrivateKeyFile - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.TLSPrivateKeyFile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MinResyncPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.MinResyncPeriod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiQPS - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.KubeAPIQPS = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "KubeApiBurst - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.KubeAPIBurst = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentDeploymentSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentDeploymentSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentEndpointSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentEndpointSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentNamespaceSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentNamespaceSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentReplicasetSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentReplicasetSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentServiceSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentServiceSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentResourceQuotaSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentResourceQuotaSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentServiceaccountTokenSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentServiceaccountTokenSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentRCSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentRCSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentHorizontalPodAustoscalerSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentHorizontalPodAustoscalerSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ConcurrentJobSyncs - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ConcurrentJobSyncs = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthenticationKubeconfig - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.AuthenticationKubeconfig = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationKubeconfig - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.AuthorizationKubeconfig = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AuthorizationAlwaysAllowPaths - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.AuthorizationAlwaysAllowPaths = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ExternalCloudVolumePlugin - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.ExternalCloudVolumePlugin = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EndpointUpdatesBatchPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.EndpointUpdatesBatchPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EndpointSliceUpdatesBatchPeriod - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.EndpointSliceUpdatesBatchPeriod = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableProfiling - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.EnableProfiling = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableContentionProfiling - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.EnableContentionProfiling = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableLeaderMigration - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.EnableLeaderMigration = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuLimit - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.KubeControllerManagerConfig {
					subject := kops.KubeControllerManagerConfig{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceKubeControllerManagerConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceKubeControllerManagerConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
