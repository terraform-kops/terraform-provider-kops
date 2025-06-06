package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceCiliumNetworkingSpec(t *testing.T) {
	_default := kops.CiliumNetworkingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.CiliumNetworkingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"registry":                          "",
					"version":                           "",
					"memory_request":                    nil,
					"cpu_request":                       nil,
					"agent_prometheus_port":             0,
					"metrics":                           func() []interface{} { return nil }(),
					"chaining_mode":                     "",
					"debug":                             false,
					"disable_endpoint_crd":              false,
					"enable_policy":                     "",
					"enable_l7_proxy":                   nil,
					"enable_local_redirect_policy":      nil,
					"enable_bpf_masquerade":             nil,
					"enable_endpoint_health_checking":   nil,
					"enable_prometheus_metrics":         false,
					"enable_encryption":                 false,
					"encryption_type":                   "",
					"node_encryption":                   false,
					"identity_allocation_mode":          "",
					"identity_change_grace_period":      "",
					"masquerade":                        nil,
					"agent_pod_annotations":             func() map[string]interface{} { return nil }(),
					"operator_pod_annotations":          func() map[string]interface{} { return nil }(),
					"tunnel":                            "",
					"monitor_aggregation":               "",
					"bpfct_global_tcp_max":              0,
					"bpfct_global_any_max":              0,
					"bpflb_algorithm":                   "",
					"bpflb_maglev_table_size":           "",
					"bpfnat_global_max":                 0,
					"bpf_neigh_global_max":              0,
					"bpf_policy_map_max":                0,
					"bpflb_map_max":                     0,
					"bpflb_sock_host_ns_only":           false,
					"preallocate_bpf_maps":              false,
					"sidecar_istio_proxy_image":         "",
					"cluster_name":                      "",
					"cluster_id":                        0,
					"to_fqdns_dns_reject_response_code": "",
					"to_fqdns_enable_poller":            false,
					"ipam":                              "",
					"install_iptables_rules":            nil,
					"auto_direct_node_routes":           false,
					"enable_host_reachable_services":    false,
					"enable_node_port":                  false,
					"etcd_managed":                      false,
					"enable_remote_node_identity":       nil,
					"enable_unreachable_routes":         nil,
					"cni_exclusive":                     nil,
					"hubble":                            nil,
					"disable_cnp_status_updates":        nil,
					"enable_service_topology":           false,
					"ingress":                           nil,
					"gateway_api":                       nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceCiliumNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceCiliumNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCiliumNetworkingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"registry":                          "",
		"version":                           "",
		"memory_request":                    nil,
		"cpu_request":                       nil,
		"agent_prometheus_port":             0,
		"metrics":                           func() []interface{} { return nil }(),
		"chaining_mode":                     "",
		"debug":                             false,
		"disable_endpoint_crd":              false,
		"enable_policy":                     "",
		"enable_l7_proxy":                   nil,
		"enable_local_redirect_policy":      nil,
		"enable_bpf_masquerade":             nil,
		"enable_endpoint_health_checking":   nil,
		"enable_prometheus_metrics":         false,
		"enable_encryption":                 false,
		"encryption_type":                   "",
		"node_encryption":                   false,
		"identity_allocation_mode":          "",
		"identity_change_grace_period":      "",
		"masquerade":                        nil,
		"agent_pod_annotations":             func() map[string]interface{} { return nil }(),
		"operator_pod_annotations":          func() map[string]interface{} { return nil }(),
		"tunnel":                            "",
		"monitor_aggregation":               "",
		"bpfct_global_tcp_max":              0,
		"bpfct_global_any_max":              0,
		"bpflb_algorithm":                   "",
		"bpflb_maglev_table_size":           "",
		"bpfnat_global_max":                 0,
		"bpf_neigh_global_max":              0,
		"bpf_policy_map_max":                0,
		"bpflb_map_max":                     0,
		"bpflb_sock_host_ns_only":           false,
		"preallocate_bpf_maps":              false,
		"sidecar_istio_proxy_image":         "",
		"cluster_name":                      "",
		"cluster_id":                        0,
		"to_fqdns_dns_reject_response_code": "",
		"to_fqdns_enable_poller":            false,
		"ipam":                              "",
		"install_iptables_rules":            nil,
		"auto_direct_node_routes":           false,
		"enable_host_reachable_services":    false,
		"enable_node_port":                  false,
		"etcd_managed":                      false,
		"enable_remote_node_identity":       nil,
		"enable_unreachable_routes":         nil,
		"cni_exclusive":                     nil,
		"hubble":                            nil,
		"disable_cnp_status_updates":        nil,
		"enable_service_topology":           false,
		"ingress":                           nil,
		"gateway_api":                       nil,
	}
	type args struct {
		in kops.CiliumNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CiliumNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Registry - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Registry = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AgentPrometheusPort - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AgentPrometheusPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metrics - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Metrics = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ChainingMode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ChainingMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Debug - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Debug = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableEndpointCRD - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableEndpointCRD = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePolicy - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnablePolicy = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableL7Proxy - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableL7Proxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableLocalRedirectPolicy - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableLocalRedirectPolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableBPFMasquerade - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableBPFMasquerade = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableEndpointHealthChecking - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableEndpointHealthChecking = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePrometheusMetrics - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnablePrometheusMetrics = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableEncryption - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableEncryption = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionType - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EncryptionType = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeEncryption - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.NodeEncryption = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityAllocationMode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IdentityAllocationMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityChangeGracePeriod - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IdentityChangeGracePeriod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Masquerade - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Masquerade = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AgentPodAnnotations - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AgentPodAnnotations = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OperatorPodAnnotations - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.OperatorPodAnnotations = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Tunnel - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Tunnel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MonitorAggregation - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.MonitorAggregation = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFCTGlobalTCPMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFCTGlobalTCPMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFCTGlobalAnyMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFCTGlobalAnyMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBAlgorithm - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBAlgorithm = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBMaglevTableSize - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBMaglevTableSize = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFNATGlobalMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFNATGlobalMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFNeighGlobalMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFNeighGlobalMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFPolicyMapMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFPolicyMapMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBMapMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBMapMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBSockHostNSOnly - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBSockHostNSOnly = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreallocateBPFMaps - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.PreallocateBPFMaps = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SidecarIstioProxyImage - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.SidecarIstioProxyImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterId - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ClusterID = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ToFqdnsDnsRejectResponseCode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFQDNsDNSRejectResponseCode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ToFqdnsEnablePoller - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFQDNsEnablePoller = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipam - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IPAM = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstallIptablesRules - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.InstallIptablesRules = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AutoDirectNodeRoutes - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AutoDirectNodeRoutes = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableHostReachableServices - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableHostReachableServices = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableNodePort - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableNodePort = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdManaged - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EtcdManaged = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRemoteNodeIdentity - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableRemoteNodeIdentity = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableUnreachableRoutes - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableUnreachableRoutes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CniExclusive - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.CniExclusive = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hubble - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Hubble = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableCNPStatusUpdates - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableCNPStatusUpdates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableServiceTopology - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableServiceTopology = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ingress - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ingress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GatewayApi - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.GatewayAPI = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceCiliumNetworkingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCiliumNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceCiliumNetworkingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"registry":                          "",
		"version":                           "",
		"memory_request":                    nil,
		"cpu_request":                       nil,
		"agent_prometheus_port":             0,
		"metrics":                           func() []interface{} { return nil }(),
		"chaining_mode":                     "",
		"debug":                             false,
		"disable_endpoint_crd":              false,
		"enable_policy":                     "",
		"enable_l7_proxy":                   nil,
		"enable_local_redirect_policy":      nil,
		"enable_bpf_masquerade":             nil,
		"enable_endpoint_health_checking":   nil,
		"enable_prometheus_metrics":         false,
		"enable_encryption":                 false,
		"encryption_type":                   "",
		"node_encryption":                   false,
		"identity_allocation_mode":          "",
		"identity_change_grace_period":      "",
		"masquerade":                        nil,
		"agent_pod_annotations":             func() map[string]interface{} { return nil }(),
		"operator_pod_annotations":          func() map[string]interface{} { return nil }(),
		"tunnel":                            "",
		"monitor_aggregation":               "",
		"bpfct_global_tcp_max":              0,
		"bpfct_global_any_max":              0,
		"bpflb_algorithm":                   "",
		"bpflb_maglev_table_size":           "",
		"bpfnat_global_max":                 0,
		"bpf_neigh_global_max":              0,
		"bpf_policy_map_max":                0,
		"bpflb_map_max":                     0,
		"bpflb_sock_host_ns_only":           false,
		"preallocate_bpf_maps":              false,
		"sidecar_istio_proxy_image":         "",
		"cluster_name":                      "",
		"cluster_id":                        0,
		"to_fqdns_dns_reject_response_code": "",
		"to_fqdns_enable_poller":            false,
		"ipam":                              "",
		"install_iptables_rules":            nil,
		"auto_direct_node_routes":           false,
		"enable_host_reachable_services":    false,
		"enable_node_port":                  false,
		"etcd_managed":                      false,
		"enable_remote_node_identity":       nil,
		"enable_unreachable_routes":         nil,
		"cni_exclusive":                     nil,
		"hubble":                            nil,
		"disable_cnp_status_updates":        nil,
		"enable_service_topology":           false,
		"ingress":                           nil,
		"gateway_api":                       nil,
	}
	type args struct {
		in kops.CiliumNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.CiliumNetworkingSpec{},
			},
			want: _default,
		},
		{
			name: "Registry - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Registry = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AgentPrometheusPort - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AgentPrometheusPort = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Metrics - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Metrics = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ChainingMode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ChainingMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Debug - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Debug = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableEndpointCRD - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableEndpointCRD = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePolicy - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnablePolicy = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableL7Proxy - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableL7Proxy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableLocalRedirectPolicy - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableLocalRedirectPolicy = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableBPFMasquerade - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableBPFMasquerade = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableEndpointHealthChecking - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableEndpointHealthChecking = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnablePrometheusMetrics - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnablePrometheusMetrics = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableEncryption - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableEncryption = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EncryptionType - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EncryptionType = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "NodeEncryption - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.NodeEncryption = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityAllocationMode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IdentityAllocationMode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IdentityChangeGracePeriod - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IdentityChangeGracePeriod = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Masquerade - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Masquerade = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AgentPodAnnotations - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AgentPodAnnotations = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OperatorPodAnnotations - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.OperatorPodAnnotations = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Tunnel - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Tunnel = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "MonitorAggregation - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.MonitorAggregation = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFCTGlobalTCPMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFCTGlobalTCPMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFCTGlobalAnyMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFCTGlobalAnyMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBAlgorithm - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBAlgorithm = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBMaglevTableSize - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBMaglevTableSize = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFNATGlobalMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFNATGlobalMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFNeighGlobalMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFNeighGlobalMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFPolicyMapMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFPolicyMapMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBMapMax - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBMapMax = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "BPFLBSockHostNSOnly - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.BPFLBSockHostNSOnly = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "PreallocateBPFMaps - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.PreallocateBPFMaps = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SidecarIstioProxyImage - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.SidecarIstioProxyImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterName - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ClusterName = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ClusterId - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ClusterID = 0
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ToFqdnsDnsRejectResponseCode - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFQDNsDNSRejectResponseCode = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "ToFqdnsEnablePoller - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.ToFQDNsEnablePoller = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ipam - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.IPAM = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "InstallIptablesRules - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.InstallIptablesRules = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AutoDirectNodeRoutes - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.AutoDirectNodeRoutes = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableHostReachableServices - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableHostReachableServices = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableNodePort - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableNodePort = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EtcdManaged - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EtcdManaged = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableRemoteNodeIdentity - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableRemoteNodeIdentity = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableUnreachableRoutes - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableUnreachableRoutes = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CniExclusive - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.CniExclusive = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Hubble - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Hubble = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "DisableCNPStatusUpdates - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.DisableCNPStatusUpdates = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "EnableServiceTopology - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.EnableServiceTopology = false
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Ingress - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.Ingress = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "GatewayApi - default",
			args: args{
				in: func() kops.CiliumNetworkingSpec {
					subject := kops.CiliumNetworkingSpec{}
					subject.GatewayAPI = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceCiliumNetworkingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceCiliumNetworkingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
