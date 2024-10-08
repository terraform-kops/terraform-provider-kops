package schemas

import (
	"reflect"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceClusterSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"channel":                           ComputedString(),
			"addons":                            ComputedList(DataSourceAddonSpec()),
			"config_store":                      ComputedStruct(DataSourceConfigStoreSpec()),
			"cloud_provider":                    ComputedStruct(DataSourceCloudProviderSpec()),
			"gossip_config":                     ComputedStruct(DataSourceGossipConfig()),
			"container_runtime":                 ComputedString(),
			"kubernetes_version":                ComputedString(),
			"dns_zone":                          ComputedString(),
			"dns_controller_gossip_config":      ComputedStruct(DataSourceDNSControllerGossipConfig()),
			"cluster_dns_domain":                ComputedString(),
			"ssh_access":                        ComputedList(String()),
			"node_port_access":                  ComputedList(String()),
			"ssh_key_name":                      ComputedString(),
			"update_policy":                     ComputedString(),
			"external_policies":                 ComputedComplexMap(List(String())),
			"additional_policies":               ComputedMap(String()),
			"file_assets":                       ComputedList(DataSourceFileAssetSpec()),
			"etcd_cluster":                      ComputedList(DataSourceEtcdClusterSpec()),
			"docker":                            ComputedStruct(DataSourceDockerConfig()),
			"containerd":                        ComputedStruct(DataSourceContainerdConfig()),
			"kube_dns":                          ComputedStruct(DataSourceKubeDNSConfig()),
			"kube_api_server":                   ComputedStruct(DataSourceKubeAPIServerConfig()),
			"kube_controller_manager":           ComputedStruct(DataSourceKubeControllerManagerConfig()),
			"external_cloud_controller_manager": ComputedStruct(DataSourceCloudControllerManagerConfig()),
			"kube_scheduler":                    ComputedStruct(DataSourceKubeSchedulerConfig()),
			"kube_proxy":                        ComputedStruct(DataSourceKubeProxyConfig()),
			"kubelet":                           ComputedStruct(DataSourceKubeletConfigSpec()),
			"control_plane_kubelet":             ComputedStruct(DataSourceKubeletConfigSpec()),
			"cloud_config":                      ComputedStruct(DataSourceCloudConfiguration()),
			"external_dns":                      ComputedStruct(DataSourceExternalDNSConfig()),
			"ntp":                               ComputedStruct(DataSourceNTPConfig()),
			"packages":                          ComputedList(String()),
			"node_problem_detector":             ComputedStruct(DataSourceNodeProblemDetectorConfig()),
			"metrics_server":                    ComputedStruct(DataSourceMetricsServerConfig()),
			"cert_manager":                      ComputedStruct(DataSourceCertManagerConfig()),
			"networking":                        ComputedStruct(DataSourceNetworkingSpec()),
			"api":                               ComputedStruct(DataSourceAPISpec()),
			"authentication":                    ComputedStruct(DataSourceAuthenticationSpec()),
			"authorization":                     ComputedStruct(DataSourceAuthorizationSpec()),
			"node_authorization":                ComputedStruct(DataSourceNodeAuthorizationSpec()),
			"cloud_labels":                      ComputedMap(String()),
			"hooks":                             ComputedList(DataSourceHookSpec()),
			"assets":                            ComputedStruct(DataSourceAssetsSpec()),
			"iam":                               ComputedStruct(DataSourceIAMSpec()),
			"encryption_config":                 ComputedBool(),
			"use_host_certificates":             ComputedBool(),
			"sysctl_parameters":                 ComputedList(String()),
			"rolling_update":                    ComputedStruct(DataSourceRollingUpdate()),
			"cluster_autoscaler":                ComputedStruct(DataSourceClusterAutoscalerConfig()),
			"service_account_issuer_discovery":  ComputedStruct(DataSourceServiceAccountIssuerDiscoveryConfig()),
			"snapshot_controller":               ComputedStruct(DataSourceSnapshotControllerConfig()),
			"karpenter":                         ComputedStruct(DataSourceKarpenterConfig()),
		},
	}

	return res
}

func ExpandDataSourceClusterSpec(in map[string]interface{}) kops.ClusterSpec {
	if in == nil {
		panic("expand ClusterSpec failure, in is nil")
	}
	return kops.ClusterSpec{
		Channel: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["channel"]),
		Addons: func(in interface{}) []kops.AddonSpec {
			return func(in interface{}) []kops.AddonSpec {
				if in == nil {
					return nil
				}
				var out []kops.AddonSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.AddonSpec {
						if in == nil {
							return kops.AddonSpec{}
						}
						return ExpandDataSourceAddonSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["addons"]),
		ConfigStore: func(in interface{}) kops.ConfigStoreSpec {
			return func(in interface{}) kops.ConfigStoreSpec {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandDataSourceConfigStoreSpec(in[0].(map[string]interface{}))
				}
				return kops.ConfigStoreSpec{}
			}(in)
		}(in["config_store"]),
		CloudProvider: func(in interface{}) kops.CloudProviderSpec {
			return func(in interface{}) kops.CloudProviderSpec {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandDataSourceCloudProviderSpec(in[0].(map[string]interface{}))
				}
				return kops.CloudProviderSpec{}
			}(in)
		}(in["cloud_provider"]),
		GossipConfig: func(in interface{}) *kops.GossipConfig {
			return func(in interface{}) *kops.GossipConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.GossipConfig) *kops.GossipConfig {
					return &in
				}(func(in interface{}) kops.GossipConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceGossipConfig(in[0].(map[string]interface{}))
					}
					return kops.GossipConfig{}
				}(in))
			}(in)
		}(in["gossip_config"]),
		ContainerRuntime: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["container_runtime"]),
		KubernetesVersion: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kubernetes_version"]),
		DNSZone: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["dns_zone"]),
		DNSControllerGossipConfig: func(in interface{}) *kops.DNSControllerGossipConfig {
			return func(in interface{}) *kops.DNSControllerGossipConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.DNSControllerGossipConfig) *kops.DNSControllerGossipConfig {
					return &in
				}(func(in interface{}) kops.DNSControllerGossipConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceDNSControllerGossipConfig(in[0].(map[string]interface{}))
					}
					return kops.DNSControllerGossipConfig{}
				}(in))
			}(in)
		}(in["dns_controller_gossip_config"]),
		ClusterDNSDomain: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_dns_domain"]),
		SSHAccess: func(in interface{}) []string {
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
		}(in["ssh_access"]),
		NodePortAccess: func(in interface{}) []string {
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
		}(in["node_port_access"]),
		SSHKeyName: func(in interface{}) *string {
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
		}(in["ssh_key_name"]),
		UpdatePolicy: func(in interface{}) *string {
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
		}(in["update_policy"]),
		ExternalPolicies: func(in interface{}) map[string][]string {
			return func(in interface{}) map[string][]string {
				if in == nil {
					return nil
				}
				if in, ok := in.([]interface{}); ok {
					if len(in) > 0 {
						out := map[string][]string{}
						for _, in := range in {
							if in, ok := in.(map[string]interface{}); ok {
								key := ExpandString(in["key"])
								value := func(in interface{}) []string {
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
								}(in["value"])
								out[key] = value
							}
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["external_policies"]),
		AdditionalPolicies: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["additional_policies"]),
		FileAssets: func(in interface{}) []kops.FileAssetSpec {
			return func(in interface{}) []kops.FileAssetSpec {
				if in == nil {
					return nil
				}
				var out []kops.FileAssetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.FileAssetSpec {
						if in == nil {
							return kops.FileAssetSpec{}
						}
						return ExpandDataSourceFileAssetSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["file_assets"]),
		EtcdClusters: func(in interface{}) []kops.EtcdClusterSpec {
			return func(in interface{}) []kops.EtcdClusterSpec {
				if in == nil {
					return nil
				}
				var out []kops.EtcdClusterSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.EtcdClusterSpec {
						if in == nil {
							return kops.EtcdClusterSpec{}
						}
						return ExpandDataSourceEtcdClusterSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["etcd_cluster"]),
		Docker: func(in interface{}) *kops.DockerConfig {
			return func(in interface{}) *kops.DockerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.DockerConfig) *kops.DockerConfig {
					return &in
				}(func(in interface{}) kops.DockerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceDockerConfig(in[0].(map[string]interface{}))
					}
					return kops.DockerConfig{}
				}(in))
			}(in)
		}(in["docker"]),
		Containerd: func(in interface{}) *kops.ContainerdConfig {
			return func(in interface{}) *kops.ContainerdConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ContainerdConfig) *kops.ContainerdConfig {
					return &in
				}(func(in interface{}) kops.ContainerdConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceContainerdConfig(in[0].(map[string]interface{}))
					}
					return kops.ContainerdConfig{}
				}(in))
			}(in)
		}(in["containerd"]),
		KubeDNS: func(in interface{}) *kops.KubeDNSConfig {
			return func(in interface{}) *kops.KubeDNSConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeDNSConfig) *kops.KubeDNSConfig {
					return &in
				}(func(in interface{}) kops.KubeDNSConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceKubeDNSConfig(in[0].(map[string]interface{}))
					}
					return kops.KubeDNSConfig{}
				}(in))
			}(in)
		}(in["kube_dns"]),
		KubeAPIServer: func(in interface{}) *kops.KubeAPIServerConfig {
			return func(in interface{}) *kops.KubeAPIServerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeAPIServerConfig) *kops.KubeAPIServerConfig {
					return &in
				}(func(in interface{}) kops.KubeAPIServerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceKubeAPIServerConfig(in[0].(map[string]interface{}))
					}
					return kops.KubeAPIServerConfig{}
				}(in))
			}(in)
		}(in["kube_api_server"]),
		KubeControllerManager: func(in interface{}) *kops.KubeControllerManagerConfig {
			return func(in interface{}) *kops.KubeControllerManagerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeControllerManagerConfig) *kops.KubeControllerManagerConfig {
					return &in
				}(func(in interface{}) kops.KubeControllerManagerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceKubeControllerManagerConfig(in[0].(map[string]interface{}))
					}
					return kops.KubeControllerManagerConfig{}
				}(in))
			}(in)
		}(in["kube_controller_manager"]),
		ExternalCloudControllerManager: func(in interface{}) *kops.CloudControllerManagerConfig {
			return func(in interface{}) *kops.CloudControllerManagerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CloudControllerManagerConfig) *kops.CloudControllerManagerConfig {
					return &in
				}(func(in interface{}) kops.CloudControllerManagerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceCloudControllerManagerConfig(in[0].(map[string]interface{}))
					}
					return kops.CloudControllerManagerConfig{}
				}(in))
			}(in)
		}(in["external_cloud_controller_manager"]),
		KubeScheduler: func(in interface{}) *kops.KubeSchedulerConfig {
			return func(in interface{}) *kops.KubeSchedulerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeSchedulerConfig) *kops.KubeSchedulerConfig {
					return &in
				}(func(in interface{}) kops.KubeSchedulerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceKubeSchedulerConfig(in[0].(map[string]interface{}))
					}
					return kops.KubeSchedulerConfig{}
				}(in))
			}(in)
		}(in["kube_scheduler"]),
		KubeProxy: func(in interface{}) *kops.KubeProxyConfig {
			return func(in interface{}) *kops.KubeProxyConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeProxyConfig) *kops.KubeProxyConfig {
					return &in
				}(func(in interface{}) kops.KubeProxyConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceKubeProxyConfig(in[0].(map[string]interface{}))
					}
					return kops.KubeProxyConfig{}
				}(in))
			}(in)
		}(in["kube_proxy"]),
		Kubelet: func(in interface{}) *kops.KubeletConfigSpec {
			return func(in interface{}) *kops.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec {
					return &in
				}(func(in interface{}) kops.KubeletConfigSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceKubeletConfigSpec(in[0].(map[string]interface{}))
					}
					return kops.KubeletConfigSpec{}
				}(in))
			}(in)
		}(in["kubelet"]),
		ControlPlaneKubelet: func(in interface{}) *kops.KubeletConfigSpec {
			return func(in interface{}) *kops.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec {
					return &in
				}(func(in interface{}) kops.KubeletConfigSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceKubeletConfigSpec(in[0].(map[string]interface{}))
					}
					return kops.KubeletConfigSpec{}
				}(in))
			}(in)
		}(in["control_plane_kubelet"]),
		CloudConfig: func(in interface{}) *kops.CloudConfiguration {
			return func(in interface{}) *kops.CloudConfiguration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CloudConfiguration) *kops.CloudConfiguration {
					return &in
				}(func(in interface{}) kops.CloudConfiguration {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceCloudConfiguration(in[0].(map[string]interface{}))
					}
					return kops.CloudConfiguration{}
				}(in))
			}(in)
		}(in["cloud_config"]),
		ExternalDNS: func(in interface{}) *kops.ExternalDNSConfig {
			return func(in interface{}) *kops.ExternalDNSConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ExternalDNSConfig) *kops.ExternalDNSConfig {
					return &in
				}(func(in interface{}) kops.ExternalDNSConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceExternalDNSConfig(in[0].(map[string]interface{}))
					}
					return kops.ExternalDNSConfig{}
				}(in))
			}(in)
		}(in["external_dns"]),
		NTP: func(in interface{}) *kops.NTPConfig {
			return func(in interface{}) *kops.NTPConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NTPConfig) *kops.NTPConfig {
					return &in
				}(func(in interface{}) kops.NTPConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceNTPConfig(in[0].(map[string]interface{}))
					}
					return kops.NTPConfig{}
				}(in))
			}(in)
		}(in["ntp"]),
		Packages: func(in interface{}) []string {
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
		}(in["packages"]),
		NodeProblemDetector: func(in interface{}) *kops.NodeProblemDetectorConfig {
			return func(in interface{}) *kops.NodeProblemDetectorConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NodeProblemDetectorConfig) *kops.NodeProblemDetectorConfig {
					return &in
				}(func(in interface{}) kops.NodeProblemDetectorConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceNodeProblemDetectorConfig(in[0].(map[string]interface{}))
					}
					return kops.NodeProblemDetectorConfig{}
				}(in))
			}(in)
		}(in["node_problem_detector"]),
		MetricsServer: func(in interface{}) *kops.MetricsServerConfig {
			return func(in interface{}) *kops.MetricsServerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.MetricsServerConfig) *kops.MetricsServerConfig {
					return &in
				}(func(in interface{}) kops.MetricsServerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceMetricsServerConfig(in[0].(map[string]interface{}))
					}
					return kops.MetricsServerConfig{}
				}(in))
			}(in)
		}(in["metrics_server"]),
		CertManager: func(in interface{}) *kops.CertManagerConfig {
			return func(in interface{}) *kops.CertManagerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.CertManagerConfig) *kops.CertManagerConfig {
					return &in
				}(func(in interface{}) kops.CertManagerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceCertManagerConfig(in[0].(map[string]interface{}))
					}
					return kops.CertManagerConfig{}
				}(in))
			}(in)
		}(in["cert_manager"]),
		Networking: func(in interface{}) kops.NetworkingSpec {
			return func(in interface{}) kops.NetworkingSpec {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandDataSourceNetworkingSpec(in[0].(map[string]interface{}))
				}
				return kops.NetworkingSpec{}
			}(in)
		}(in["networking"]),
		API: func(in interface{}) kops.APISpec {
			return func(in interface{}) kops.APISpec {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandDataSourceAPISpec(in[0].(map[string]interface{}))
				}
				return kops.APISpec{}
			}(in)
		}(in["api"]),
		Authentication: func(in interface{}) *kops.AuthenticationSpec {
			return func(in interface{}) *kops.AuthenticationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AuthenticationSpec) *kops.AuthenticationSpec {
					return &in
				}(func(in interface{}) kops.AuthenticationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceAuthenticationSpec(in[0].(map[string]interface{}))
					}
					return kops.AuthenticationSpec{}
				}(in))
			}(in)
		}(in["authentication"]),
		Authorization: func(in interface{}) *kops.AuthorizationSpec {
			return func(in interface{}) *kops.AuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AuthorizationSpec) *kops.AuthorizationSpec {
					return &in
				}(func(in interface{}) kops.AuthorizationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceAuthorizationSpec(in[0].(map[string]interface{}))
					}
					return kops.AuthorizationSpec{}
				}(in))
			}(in)
		}(in["authorization"]),
		NodeAuthorization: func(in interface{}) *kops.NodeAuthorizationSpec {
			return func(in interface{}) *kops.NodeAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.NodeAuthorizationSpec) *kops.NodeAuthorizationSpec {
					return &in
				}(func(in interface{}) kops.NodeAuthorizationSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceNodeAuthorizationSpec(in[0].(map[string]interface{}))
					}
					return kops.NodeAuthorizationSpec{}
				}(in))
			}(in)
		}(in["node_authorization"]),
		CloudLabels: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["cloud_labels"]),
		Hooks: func(in interface{}) []kops.HookSpec {
			return func(in interface{}) []kops.HookSpec {
				if in == nil {
					return nil
				}
				var out []kops.HookSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.HookSpec {
						if in == nil {
							return kops.HookSpec{}
						}
						return ExpandDataSourceHookSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["hooks"]),
		Assets: func(in interface{}) *kops.AssetsSpec {
			return func(in interface{}) *kops.AssetsSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AssetsSpec) *kops.AssetsSpec {
					return &in
				}(func(in interface{}) kops.AssetsSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceAssetsSpec(in[0].(map[string]interface{}))
					}
					return kops.AssetsSpec{}
				}(in))
			}(in)
		}(in["assets"]),
		IAM: func(in interface{}) *kops.IAMSpec {
			return func(in interface{}) *kops.IAMSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.IAMSpec) *kops.IAMSpec {
					return &in
				}(func(in interface{}) kops.IAMSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceIAMSpec(in[0].(map[string]interface{}))
					}
					return kops.IAMSpec{}
				}(in))
			}(in)
		}(in["iam"]),
		EncryptionConfig: func(in interface{}) *bool {
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
		}(in["encryption_config"]),
		UseHostCertificates: func(in interface{}) *bool {
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
		}(in["use_host_certificates"]),
		SysctlParameters: func(in interface{}) []string {
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
		}(in["sysctl_parameters"]),
		RollingUpdate: func(in interface{}) *kops.RollingUpdate {
			return func(in interface{}) *kops.RollingUpdate {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.RollingUpdate) *kops.RollingUpdate {
					return &in
				}(func(in interface{}) kops.RollingUpdate {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceRollingUpdate(in[0].(map[string]interface{}))
					}
					return kops.RollingUpdate{}
				}(in))
			}(in)
		}(in["rolling_update"]),
		ClusterAutoscaler: func(in interface{}) *kops.ClusterAutoscalerConfig {
			return func(in interface{}) *kops.ClusterAutoscalerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ClusterAutoscalerConfig) *kops.ClusterAutoscalerConfig {
					return &in
				}(func(in interface{}) kops.ClusterAutoscalerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceClusterAutoscalerConfig(in[0].(map[string]interface{}))
					}
					return kops.ClusterAutoscalerConfig{}
				}(in))
			}(in)
		}(in["cluster_autoscaler"]),
		ServiceAccountIssuerDiscovery: func(in interface{}) *kops.ServiceAccountIssuerDiscoveryConfig {
			return func(in interface{}) *kops.ServiceAccountIssuerDiscoveryConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ServiceAccountIssuerDiscoveryConfig) *kops.ServiceAccountIssuerDiscoveryConfig {
					return &in
				}(func(in interface{}) kops.ServiceAccountIssuerDiscoveryConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceServiceAccountIssuerDiscoveryConfig(in[0].(map[string]interface{}))
					}
					return kops.ServiceAccountIssuerDiscoveryConfig{}
				}(in))
			}(in)
		}(in["service_account_issuer_discovery"]),
		SnapshotController: func(in interface{}) *kops.SnapshotControllerConfig {
			return func(in interface{}) *kops.SnapshotControllerConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.SnapshotControllerConfig) *kops.SnapshotControllerConfig {
					return &in
				}(func(in interface{}) kops.SnapshotControllerConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceSnapshotControllerConfig(in[0].(map[string]interface{}))
					}
					return kops.SnapshotControllerConfig{}
				}(in))
			}(in)
		}(in["snapshot_controller"]),
		Karpenter: func(in interface{}) *kops.KarpenterConfig {
			return func(in interface{}) *kops.KarpenterConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KarpenterConfig) *kops.KarpenterConfig {
					return &in
				}(func(in interface{}) kops.KarpenterConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceKarpenterConfig(in[0].(map[string]interface{}))
					}
					return kops.KarpenterConfig{}
				}(in))
			}(in)
		}(in["karpenter"]),
	}
}

func FlattenDataSourceClusterSpecInto(in kops.ClusterSpec, out map[string]interface{}) {
	out["channel"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Channel)
	out["addons"] = func(in []kops.AddonSpec) interface{} {
		return func(in []kops.AddonSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.AddonSpec) interface{} {
					return FlattenDataSourceAddonSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Addons)
	out["config_store"] = func(in kops.ConfigStoreSpec) interface{} {
		return func(in kops.ConfigStoreSpec) []interface{} {
			return []interface{}{FlattenDataSourceConfigStoreSpec(in)}
		}(in)
	}(in.ConfigStore)
	out["cloud_provider"] = func(in kops.CloudProviderSpec) interface{} {
		return func(in kops.CloudProviderSpec) []interface{} {
			return []interface{}{FlattenDataSourceCloudProviderSpec(in)}
		}(in)
	}(in.CloudProvider)
	out["gossip_config"] = func(in *kops.GossipConfig) interface{} {
		return func(in *kops.GossipConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.GossipConfig) interface{} {
				return func(in kops.GossipConfig) []interface{} {
					return []interface{}{FlattenDataSourceGossipConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.GossipConfig)
	out["container_runtime"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ContainerRuntime)
	out["kubernetes_version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KubernetesVersion)
	out["dns_zone"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.DNSZone)
	out["dns_controller_gossip_config"] = func(in *kops.DNSControllerGossipConfig) interface{} {
		return func(in *kops.DNSControllerGossipConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.DNSControllerGossipConfig) interface{} {
				return func(in kops.DNSControllerGossipConfig) []interface{} {
					return []interface{}{FlattenDataSourceDNSControllerGossipConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.DNSControllerGossipConfig)
	out["cluster_dns_domain"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterDNSDomain)
	out["ssh_access"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.SSHAccess)
	out["node_port_access"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.NodePortAccess)
	out["ssh_key_name"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.SSHKeyName)
	out["update_policy"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.UpdatePolicy)
	out["external_policies"] = func(in map[string][]string) interface{} {
		return func(in map[string][]string) []interface{} {
			if in == nil {
				return nil
			}
			keys := make([]string, 0, len(in))
			for key := range in {
				keys = append(keys, key)
			}
			sort.SliceStable(keys, func(i, j int) bool {
				return keys[i] < keys[j]
			})
			var out []interface{}
			for _, key := range keys {
				in := in[key]
				out = append(out, map[string]interface{}{
					"key": key,
					"value": func(in []string) []interface{} {
						var out []interface{}
						for _, in := range in {
							out = append(out, FlattenString(string(in)))
						}
						return out
					}(in),
				})
			}
			return out
		}(in)
	}(in.ExternalPolicies)
	out["additional_policies"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.AdditionalPolicies)
	out["file_assets"] = func(in []kops.FileAssetSpec) interface{} {
		return func(in []kops.FileAssetSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.FileAssetSpec) interface{} {
					return FlattenDataSourceFileAssetSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.FileAssets)
	out["etcd_cluster"] = func(in []kops.EtcdClusterSpec) interface{} {
		return func(in []kops.EtcdClusterSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.EtcdClusterSpec) interface{} {
					return FlattenDataSourceEtcdClusterSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.EtcdClusters)
	out["docker"] = func(in *kops.DockerConfig) interface{} {
		return func(in *kops.DockerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.DockerConfig) interface{} {
				return func(in kops.DockerConfig) []interface{} {
					return []interface{}{FlattenDataSourceDockerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Docker)
	out["containerd"] = func(in *kops.ContainerdConfig) interface{} {
		return func(in *kops.ContainerdConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ContainerdConfig) interface{} {
				return func(in kops.ContainerdConfig) []interface{} {
					return []interface{}{FlattenDataSourceContainerdConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Containerd)
	out["kube_dns"] = func(in *kops.KubeDNSConfig) interface{} {
		return func(in *kops.KubeDNSConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeDNSConfig) interface{} {
				return func(in kops.KubeDNSConfig) []interface{} {
					return []interface{}{FlattenDataSourceKubeDNSConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeDNS)
	out["kube_api_server"] = func(in *kops.KubeAPIServerConfig) interface{} {
		return func(in *kops.KubeAPIServerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeAPIServerConfig) interface{} {
				return func(in kops.KubeAPIServerConfig) []interface{} {
					return []interface{}{FlattenDataSourceKubeAPIServerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeAPIServer)
	out["kube_controller_manager"] = func(in *kops.KubeControllerManagerConfig) interface{} {
		return func(in *kops.KubeControllerManagerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeControllerManagerConfig) interface{} {
				return func(in kops.KubeControllerManagerConfig) []interface{} {
					return []interface{}{FlattenDataSourceKubeControllerManagerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeControllerManager)
	out["external_cloud_controller_manager"] = func(in *kops.CloudControllerManagerConfig) interface{} {
		return func(in *kops.CloudControllerManagerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.CloudControllerManagerConfig) interface{} {
				return func(in kops.CloudControllerManagerConfig) []interface{} {
					return []interface{}{FlattenDataSourceCloudControllerManagerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ExternalCloudControllerManager)
	out["kube_scheduler"] = func(in *kops.KubeSchedulerConfig) interface{} {
		return func(in *kops.KubeSchedulerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeSchedulerConfig) interface{} {
				return func(in kops.KubeSchedulerConfig) []interface{} {
					return []interface{}{FlattenDataSourceKubeSchedulerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeScheduler)
	out["kube_proxy"] = func(in *kops.KubeProxyConfig) interface{} {
		return func(in *kops.KubeProxyConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeProxyConfig) interface{} {
				return func(in kops.KubeProxyConfig) []interface{} {
					return []interface{}{FlattenDataSourceKubeProxyConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.KubeProxy)
	out["kubelet"] = func(in *kops.KubeletConfigSpec) interface{} {
		return func(in *kops.KubeletConfigSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeletConfigSpec) interface{} {
				return func(in kops.KubeletConfigSpec) []interface{} {
					return []interface{}{FlattenDataSourceKubeletConfigSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kubelet)
	out["control_plane_kubelet"] = func(in *kops.KubeletConfigSpec) interface{} {
		return func(in *kops.KubeletConfigSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeletConfigSpec) interface{} {
				return func(in kops.KubeletConfigSpec) []interface{} {
					return []interface{}{FlattenDataSourceKubeletConfigSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ControlPlaneKubelet)
	out["cloud_config"] = func(in *kops.CloudConfiguration) interface{} {
		return func(in *kops.CloudConfiguration) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.CloudConfiguration) interface{} {
				return func(in kops.CloudConfiguration) []interface{} {
					return []interface{}{FlattenDataSourceCloudConfiguration(in)}
				}(in)
			}(*in)
		}(in)
	}(in.CloudConfig)
	out["external_dns"] = func(in *kops.ExternalDNSConfig) interface{} {
		return func(in *kops.ExternalDNSConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ExternalDNSConfig) interface{} {
				return func(in kops.ExternalDNSConfig) []interface{} {
					return []interface{}{FlattenDataSourceExternalDNSConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ExternalDNS)
	out["ntp"] = func(in *kops.NTPConfig) interface{} {
		return func(in *kops.NTPConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.NTPConfig) interface{} {
				return func(in kops.NTPConfig) []interface{} {
					return []interface{}{FlattenDataSourceNTPConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NTP)
	out["packages"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Packages)
	out["node_problem_detector"] = func(in *kops.NodeProblemDetectorConfig) interface{} {
		return func(in *kops.NodeProblemDetectorConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.NodeProblemDetectorConfig) interface{} {
				return func(in kops.NodeProblemDetectorConfig) []interface{} {
					return []interface{}{FlattenDataSourceNodeProblemDetectorConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeProblemDetector)
	out["metrics_server"] = func(in *kops.MetricsServerConfig) interface{} {
		return func(in *kops.MetricsServerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.MetricsServerConfig) interface{} {
				return func(in kops.MetricsServerConfig) []interface{} {
					return []interface{}{FlattenDataSourceMetricsServerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.MetricsServer)
	out["cert_manager"] = func(in *kops.CertManagerConfig) interface{} {
		return func(in *kops.CertManagerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.CertManagerConfig) interface{} {
				return func(in kops.CertManagerConfig) []interface{} {
					return []interface{}{FlattenDataSourceCertManagerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.CertManager)
	out["networking"] = func(in kops.NetworkingSpec) interface{} {
		return func(in kops.NetworkingSpec) []interface{} {
			return []interface{}{FlattenDataSourceNetworkingSpec(in)}
		}(in)
	}(in.Networking)
	out["api"] = func(in kops.APISpec) interface{} {
		return func(in kops.APISpec) []interface{} {
			return []interface{}{FlattenDataSourceAPISpec(in)}
		}(in)
	}(in.API)
	out["authentication"] = func(in *kops.AuthenticationSpec) interface{} {
		return func(in *kops.AuthenticationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AuthenticationSpec) interface{} {
				return func(in kops.AuthenticationSpec) []interface{} {
					return []interface{}{FlattenDataSourceAuthenticationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Authentication)
	out["authorization"] = func(in *kops.AuthorizationSpec) interface{} {
		return func(in *kops.AuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AuthorizationSpec) interface{} {
				return func(in kops.AuthorizationSpec) []interface{} {
					return []interface{}{FlattenDataSourceAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Authorization)
	out["node_authorization"] = func(in *kops.NodeAuthorizationSpec) interface{} {
		return func(in *kops.NodeAuthorizationSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.NodeAuthorizationSpec) interface{} {
				return func(in kops.NodeAuthorizationSpec) []interface{} {
					return []interface{}{FlattenDataSourceNodeAuthorizationSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.NodeAuthorization)
	out["cloud_labels"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.CloudLabels)
	out["hooks"] = func(in []kops.HookSpec) interface{} {
		return func(in []kops.HookSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.HookSpec) interface{} {
					return FlattenDataSourceHookSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Hooks)
	out["assets"] = func(in *kops.AssetsSpec) interface{} {
		return func(in *kops.AssetsSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.AssetsSpec) interface{} {
				return func(in kops.AssetsSpec) []interface{} {
					return []interface{}{FlattenDataSourceAssetsSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Assets)
	out["iam"] = func(in *kops.IAMSpec) interface{} {
		return func(in *kops.IAMSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.IAMSpec) interface{} {
				return func(in kops.IAMSpec) []interface{} {
					return []interface{}{FlattenDataSourceIAMSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.IAM)
	out["encryption_config"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.EncryptionConfig)
	out["use_host_certificates"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.UseHostCertificates)
	out["sysctl_parameters"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.SysctlParameters)
	out["rolling_update"] = func(in *kops.RollingUpdate) interface{} {
		return func(in *kops.RollingUpdate) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.RollingUpdate) interface{} {
				return func(in kops.RollingUpdate) []interface{} {
					return []interface{}{FlattenDataSourceRollingUpdate(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RollingUpdate)
	out["cluster_autoscaler"] = func(in *kops.ClusterAutoscalerConfig) interface{} {
		return func(in *kops.ClusterAutoscalerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ClusterAutoscalerConfig) interface{} {
				return func(in kops.ClusterAutoscalerConfig) []interface{} {
					return []interface{}{FlattenDataSourceClusterAutoscalerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ClusterAutoscaler)
	out["service_account_issuer_discovery"] = func(in *kops.ServiceAccountIssuerDiscoveryConfig) interface{} {
		return func(in *kops.ServiceAccountIssuerDiscoveryConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ServiceAccountIssuerDiscoveryConfig) interface{} {
				return func(in kops.ServiceAccountIssuerDiscoveryConfig) []interface{} {
					return []interface{}{FlattenDataSourceServiceAccountIssuerDiscoveryConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.ServiceAccountIssuerDiscovery)
	out["snapshot_controller"] = func(in *kops.SnapshotControllerConfig) interface{} {
		return func(in *kops.SnapshotControllerConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.SnapshotControllerConfig) interface{} {
				return func(in kops.SnapshotControllerConfig) []interface{} {
					return []interface{}{FlattenDataSourceSnapshotControllerConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.SnapshotController)
	out["karpenter"] = func(in *kops.KarpenterConfig) interface{} {
		return func(in *kops.KarpenterConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KarpenterConfig) interface{} {
				return func(in kops.KarpenterConfig) []interface{} {
					return []interface{}{FlattenDataSourceKarpenterConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Karpenter)
}

func FlattenDataSourceClusterSpec(in kops.ClusterSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceClusterSpecInto(in, out)
	return out
}
