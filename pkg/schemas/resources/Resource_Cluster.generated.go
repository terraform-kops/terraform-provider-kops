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

func ResourceCluster() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"channel":                           OptionalString(),
			"addons":                            OptionalList(kopsschemas.ResourceAddonSpec()),
			"config_store":                      OptionalStruct(kopsschemas.ResourceConfigStoreSpec()),
			"cloud_provider":                    RequiredStruct(kopsschemas.ResourceCloudProviderSpec()),
			"gossip_config":                     OptionalStruct(kopsschemas.ResourceGossipConfig()),
			"container_runtime":                 OptionalString(),
			"kubernetes_version":                OptionalString(),
			"dns_zone":                          OptionalString(),
			"dns_controller_gossip_config":      OptionalStruct(kopsschemas.ResourceDNSControllerGossipConfig()),
			"cluster_dns_domain":                OptionalString(),
			"ssh_access":                        OptionalList(String()),
			"node_port_access":                  OptionalList(String()),
			"ssh_key_name":                      OptionalString(),
			"update_policy":                     OptionalString(),
			"external_policies":                 OptionalComplexMap(List(String())),
			"additional_policies":               OptionalMap(String()),
			"file_assets":                       OptionalList(kopsschemas.ResourceFileAssetSpec()),
			"etcd_cluster":                      RequiredList(kopsschemas.ResourceEtcdClusterSpec()),
			"docker":                            OptionalStruct(kopsschemas.ResourceDockerConfig()),
			"containerd":                        OptionalStruct(kopsschemas.ResourceContainerdConfig()),
			"kube_dns":                          OptionalStruct(kopsschemas.ResourceKubeDNSConfig()),
			"kube_api_server":                   OptionalStruct(kopsschemas.ResourceKubeAPIServerConfig()),
			"kube_controller_manager":           OptionalStruct(kopsschemas.ResourceKubeControllerManagerConfig()),
			"external_cloud_controller_manager": OptionalStruct(kopsschemas.ResourceCloudControllerManagerConfig()),
			"kube_scheduler":                    OptionalStruct(kopsschemas.ResourceKubeSchedulerConfig()),
			"kube_proxy":                        OptionalStruct(kopsschemas.ResourceKubeProxyConfig()),
			"kubelet":                           OptionalStruct(kopsschemas.ResourceKubeletConfigSpec()),
			"control_plane_kubelet":             OptionalStruct(kopsschemas.ResourceKubeletConfigSpec()),
			"cloud_config":                      OptionalStruct(kopsschemas.ResourceCloudConfiguration()),
			"external_dns":                      OptionalStruct(kopsschemas.ResourceExternalDNSConfig()),
			"ntp":                               OptionalStruct(kopsschemas.ResourceNTPConfig()),
			"packages":                          OptionalList(String()),
			"node_problem_detector":             OptionalStruct(kopsschemas.ResourceNodeProblemDetectorConfig()),
			"metrics_server":                    OptionalStruct(kopsschemas.ResourceMetricsServerConfig()),
			"cert_manager":                      OptionalStruct(kopsschemas.ResourceCertManagerConfig()),
			"networking":                        RequiredStruct(kopsschemas.ResourceNetworkingSpec()),
			"api":                               OptionalStruct(kopsschemas.ResourceAPISpec()),
			"authentication":                    OptionalStruct(kopsschemas.ResourceAuthenticationSpec()),
			"authorization":                     OptionalStruct(kopsschemas.ResourceAuthorizationSpec()),
			"node_authorization":                OptionalStruct(kopsschemas.ResourceNodeAuthorizationSpec()),
			"cloud_labels":                      OptionalMap(String()),
			"hooks":                             OptionalList(kopsschemas.ResourceHookSpec()),
			"assets":                            OptionalStruct(kopsschemas.ResourceAssetsSpec()),
			"iam":                               OptionalComputedStruct(kopsschemas.ResourceIAMSpec()),
			"encryption_config":                 OptionalBool(),
			"use_host_certificates":             OptionalBool(),
			"sysctl_parameters":                 OptionalList(String()),
			"rolling_update":                    OptionalStruct(kopsschemas.ResourceRollingUpdate()),
			"cluster_autoscaler":                OptionalStruct(kopsschemas.ResourceClusterAutoscalerConfig()),
			"service_account_issuer_discovery":  OptionalStruct(kopsschemas.ResourceServiceAccountIssuerDiscoveryConfig()),
			"snapshot_controller":               OptionalStruct(kopsschemas.ResourceSnapshotControllerConfig()),
			"karpenter":                         OptionalStruct(kopsschemas.ResourceKarpenterConfig()),
			"labels":                            OptionalMap(String()),
			"annotations":                       OptionalMap(String()),
			"name":                              ForceNew(RequiredString()),
			"admin_ssh_key":                     Sensitive(OptionalString()),
			"secrets":                           OptionalStruct(ResourceClusterSecrets()),
			"revision":                          ComputedInt(),
			"delete":                            OptionalStruct(ResourceDeleteOptions()),
		},
	}
	res.SchemaVersion = 5
	res.StateUpgraders = []schema.StateUpgrader{
		{
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenResourceCluster(ExpandResourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 0,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenResourceCluster(ExpandResourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 1,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenResourceCluster(ExpandResourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 2,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenResourceCluster(ExpandResourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 3,
		}, {
			Type: res.CoreConfigSchema().ImpliedType(),
			Upgrade: func(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
				ret := FlattenResourceCluster(ExpandResourceCluster(rawState))
				ret["id"] = rawState["id"]
				return ret, nil
			},
			Version: 4,
		},
	}
	return res
}

func ExpandResourceCluster(in map[string]interface{}) resources.Cluster {
	if in == nil {
		panic("expand Cluster failure, in is nil")
	}
	return resources.Cluster{
		ClusterSpec: func(in interface{}) kops.ClusterSpec {
			return kopsschemas.ExpandResourceClusterSpec(in.(map[string]interface{}))
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
						return ExpandResourceClusterSecrets(in[0].(map[string]interface{}))
					}
					return resources.ClusterSecrets{}
				}(in))
			}(in)
		}(in["secrets"]),
		Revision: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["revision"]),
		Delete: func(in interface{}) resources.DeleteOptions {
			return func(in interface{}) resources.DeleteOptions {
				if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
					return ExpandResourceDeleteOptions(in[0].(map[string]interface{}))
				}
				return resources.DeleteOptions{}
			}(in)
		}(in["delete"]),
	}
}

func FlattenResourceClusterInto(in resources.Cluster, out map[string]interface{}) {
	kopsschemas.FlattenResourceClusterSpecInto(in.ClusterSpec, out)
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
					return []interface{}{FlattenResourceClusterSecrets(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Secrets)
	out["revision"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Revision)
	out["delete"] = func(in resources.DeleteOptions) interface{} {
		return func(in resources.DeleteOptions) []interface{} {
			return []interface{}{FlattenResourceDeleteOptions(in)}
		}(in)
	}(in.Delete)
}

func FlattenResourceCluster(in resources.Cluster) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceClusterInto(in, out)
	return out
}
