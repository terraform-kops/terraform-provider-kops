package schemas

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-kops/terraform-provider-kops/pkg/api/resources"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	kopsschemas "github.com/terraform-kops/terraform-provider-kops/pkg/schemas/kops"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCluster() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"channel":                           ComputedString(),
			"addons":                            ComputedList(kopsschemas.DataSourceAddonSpec()),
			"config_store":                      ComputedStruct(kopsschemas.DataSourceConfigStoreSpec()),
			"cloud_provider":                    ComputedStruct(kopsschemas.DataSourceCloudProviderSpec()),
			"gossip_config":                     ComputedStruct(kopsschemas.DataSourceGossipConfig()),
			"container_runtime":                 ComputedString(),
			"kubernetes_version":                ComputedString(),
			"dns_zone":                          ComputedString(),
			"dns_controller_gossip_config":      ComputedStruct(kopsschemas.DataSourceDNSControllerGossipConfig()),
			"cluster_dns_domain":                ComputedString(),
			"ssh_access":                        ComputedList(String()),
			"node_port_access":                  ComputedList(String()),
			"ssh_key_name":                      ComputedString(),
			"update_policy":                     ComputedString(),
			"external_policies":                 ComputedComplexMap(List(String())),
			"additional_policies":               ComputedMap(String()),
			"file_assets":                       ComputedList(kopsschemas.DataSourceFileAssetSpec()),
			"etcd_cluster":                      ComputedList(kopsschemas.DataSourceEtcdClusterSpec()),
			"docker":                            ComputedStruct(kopsschemas.DataSourceDockerConfig()),
			"containerd":                        ComputedStruct(kopsschemas.DataSourceContainerdConfig()),
			"kube_dns":                          ComputedStruct(kopsschemas.DataSourceKubeDNSConfig()),
			"kube_api_server":                   ComputedStruct(kopsschemas.DataSourceKubeAPIServerConfig()),
			"kube_controller_manager":           ComputedStruct(kopsschemas.DataSourceKubeControllerManagerConfig()),
			"external_cloud_controller_manager": ComputedStruct(kopsschemas.DataSourceCloudControllerManagerConfig()),
			"kube_scheduler":                    ComputedStruct(kopsschemas.DataSourceKubeSchedulerConfig()),
			"kube_proxy":                        ComputedStruct(kopsschemas.DataSourceKubeProxyConfig()),
			"kubelet":                           ComputedStruct(kopsschemas.DataSourceKubeletConfigSpec()),
			"control_plane_kubelet":             ComputedStruct(kopsschemas.DataSourceKubeletConfigSpec()),
			"cloud_config":                      ComputedStruct(kopsschemas.DataSourceCloudConfiguration()),
			"external_dns":                      ComputedStruct(kopsschemas.DataSourceExternalDNSConfig()),
			"ntp":                               ComputedStruct(kopsschemas.DataSourceNTPConfig()),
			"packages":                          ComputedList(String()),
			"node_problem_detector":             ComputedStruct(kopsschemas.DataSourceNodeProblemDetectorConfig()),
			"metrics_server":                    ComputedStruct(kopsschemas.DataSourceMetricsServerConfig()),
			"cert_manager":                      ComputedStruct(kopsschemas.DataSourceCertManagerConfig()),
			"networking":                        ComputedStruct(kopsschemas.DataSourceNetworkingSpec()),
			"api":                               ComputedStruct(kopsschemas.DataSourceAPISpec()),
			"authentication":                    ComputedStruct(kopsschemas.DataSourceAuthenticationSpec()),
			"authorization":                     ComputedStruct(kopsschemas.DataSourceAuthorizationSpec()),
			"node_authorization":                ComputedStruct(kopsschemas.DataSourceNodeAuthorizationSpec()),
			"cloud_labels":                      ComputedMap(String()),
			"hooks":                             ComputedList(kopsschemas.DataSourceHookSpec()),
			"assets":                            ComputedStruct(kopsschemas.DataSourceAssetsSpec()),
			"iam":                               ComputedStruct(kopsschemas.DataSourceIAMSpec()),
			"encryption_config":                 ComputedBool(),
			"use_host_certificates":             ComputedBool(),
			"sysctl_parameters":                 ComputedList(String()),
			"rolling_update":                    ComputedStruct(kopsschemas.DataSourceRollingUpdate()),
			"cluster_autoscaler":                ComputedStruct(kopsschemas.DataSourceClusterAutoscalerConfig()),
			"service_account_issuer_discovery":  ComputedStruct(kopsschemas.DataSourceServiceAccountIssuerDiscoveryConfig()),
			"snapshot_controller":               ComputedStruct(kopsschemas.DataSourceSnapshotControllerConfig()),
			"karpenter":                         ComputedStruct(kopsschemas.DataSourceKarpenterConfig()),
			"labels":                            ComputedMap(String()),
			"annotations":                       ComputedMap(String()),
			"name":                              RequiredString(),
			"admin_ssh_key":                     ComputedString(),
			"secrets":                           ComputedStruct(DataSourceClusterSecrets()),
			"delete":                            ComputedStruct(DataSourceDeleteOptions()),
		},
	}
	res.SchemaVersion = 5
	res.StateUpgraders = []schema.StateUpgrader{
		{
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenDataSourceCluster(ExpandDataSourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 0,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenDataSourceCluster(ExpandDataSourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 1,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenDataSourceCluster(ExpandDataSourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 2,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenDataSourceCluster(ExpandDataSourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 3,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenDataSourceCluster(ExpandDataSourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 4,
		},
	}
	return res
}

func ExpandDataSourceCluster(in map[string]interface{}) resources.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return resources.Cluster{
		ClusterSpec: func(in interface{}) kops.ClusterSpec {
			return kopsschemas.ExpandDataSourceClusterSpec(in.(map[string]interface{}))
		}(in),
		Labels: func(in interface{}) map[string]string {
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
		}(in["labels"]),
		Annotations: func(in interface{}) map[string]string {
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
		}(in["annotations"]),
		Name: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["name"]),
		AdminSshKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["admin_ssh_key"]),
		Secrets: func(in interface{}) *resources.ClusterSecrets {
			return func(in interface{}) *resources.ClusterSecrets {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in resources.ClusterSecrets) *resources.ClusterSecrets {
					return &in
				}(func(in interface{}) resources.ClusterSecrets {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceClusterSecrets(in[0].(map[string]interface{}))
					}
					return resources.ClusterSecrets{}
				}(in))
			}(in)
		}(in["secrets"]),
		Delete: func(in interface{}) resources.DeleteOptions {
			return func(in interface{}) resources.DeleteOptions {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandDataSourceDeleteOptions(in[0].(map[string]interface{}))
				}
				return resources.DeleteOptions{}
			}(in)
		}(in["delete"]),
	}
}

func FlattenDataSourceClusterInto(in resources.Cluster, out map[string]interface{}) {
	kopsschemas.FlattenDataSourceClusterSpecInto(in.ClusterSpec, out)
	out["labels"] = func(in map[string]string) interface{} {
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
	}(in.Labels)
	out["annotations"] = func(in map[string]string) interface{} {
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
	}(in.Annotations)
	out["name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Name)
	out["admin_ssh_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AdminSshKey)
	out["secrets"] = func(in *resources.ClusterSecrets) interface{} {
		return func(in *resources.ClusterSecrets) interface{} {
			if in == nil {
				return nil
			}
			return func(in resources.ClusterSecrets) interface{} {
				return func(in resources.ClusterSecrets) []interface{} {
					return []interface{}{FlattenDataSourceClusterSecrets(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Secrets)
	out["delete"] = func(in resources.DeleteOptions) interface{} {
		return func(in resources.DeleteOptions) []interface{} {
			return []interface{}{FlattenDataSourceDeleteOptions(in)}
		}(in)
	}(in.Delete)
}

func FlattenDataSourceCluster(in resources.Cluster) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceClusterInto(in, out)
	return out
}
