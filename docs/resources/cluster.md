# kops_cluster

Provides a kOps cluster resource.

## Example usage

```hcl
resource "kops_cluster" "cluster" {
  name                 = "cluster.example.com"
  admin_ssh_key        = file("path to ssh public key file")
  kubernetes_version   = "stable"
  dns_zone             = "example.com"
  network_id           = "net-0"

  cloud_provider {
    aws {}
  }

  iam {
    allow_container_registry = true
  }

  networking {
    calico {}
  }

  topology {
    masters = "private"
    nodes   = "private"

    dns {
      type = "Private"
    }
  }

  # cluster subnets
  subnet {
    name        = "private-0"
    provider_id = "subnet-0"
    type        = "Private"
    zone        = "zone-0"
  }

  subnet {
    name        = "private-1"
    provider_id = "subnet-1"
    type        = "Private"
    zone        = "zone-1"
  }

  subnet {
    name        = "private-2"
    provider_id = "subnet-2"
    type        = "Private"
    zone        = "zone-2"
  }

  # etcd clusters
  etcd_cluster {
    name            = "main"

    member {
      name             = "master-0"
      instance_group   = "master-0"
    }

    member {
      name             = "master-1"
      instance_group   = "master-1"
    }

    member {
      name             = "master-2"
      instance_group   = "master-2"
    }
  }

  etcd_cluster {
    name            = "events"

    member {
      name             = "master-0"
      instance_group   = "master-0"
    }

    member {
      name             = "master-1"
      instance_group   = "master-1"
    }

    member {
      name             = "master-2"
      instance_group   = "master-2"
    }
  }
}
```

## Nullable arguments

Because kOps sometimes uses pointers to hold data and terraform doesn't offer a way to
differentiate between unset arguments and their default value in a configuration, it can
be necessary wrap those arguments in a nested resource to account for the `null` value.

An example of this is the `anonymous_auth` argument in the `kube_api_server_config` or `kubelet_config_spec`
resources. The `null` value cannot be considered equivalent to `false` on the kOps side, but terraform won't
let us know when it is set or not in the configuration and therefore will provide `false` in case it is unset
in the configuration.

To workaround this limitation, the nullable type is a resource with a single `value` argument. It wraps the
actual value of the argument and makes it possible to account for `null`. When using a nullable argument,
you should assign it this way in the configuration:

```hcl
resource "kops_cluster" "cluster" {
  // ...

  kubelet {
    anonymous_auth {
      value = false
    }
  }

  kube_api_server {
    anonymous_auth {
      value = false
    }
  }

  // ...
}
```


## Argument Reference

The following arguments are supported:
- `channel` - (Optional) - String - The Channel we are following.
- `addons` - (Optional) - List([addon_spec](#addon_spec)) - Additional addons that should be installed on the cluster.
- `config_store` - (Optional) - [config_store_spec](#config_store_spec) - ConfigStore configures the stores that nodes use to get their configuration.
- `cloud_provider` - (Required) - [cloud_provider_spec](#cloud_provider_spec) - CloudProvider configures the cloud provider to use.
- `gossip_config` - (Optional) - [gossip_config](#gossip_config) - GossipConfig for the cluster assuming the use of gossip DNS.
- `container_runtime` - (Optional) - String - ContainerRuntime was removed.
- `kubernetes_version` - (Optional) - String - The version of kubernetes to install (optional, and can be a "spec" like stable).
- `dns_zone` - (Optional) - String - DNSZone is the DNS zone we should use when configuring DNS<br />This is because some clouds let us define a managed zone foo.bar, and then have<br />kubernetes.dev.foo.bar, without needing to define dev.foo.bar as a hosted zone.<br />DNSZone will probably be a suffix of the MasterPublicName.<br />Note that DNSZone can either by the host name of the zone (containing dots),<br />or can be an identifier for the zone.
- `dns_controller_gossip_config` - (Optional) - [dns_controller_gossip_config](#dns_controller_gossip_config) - DNSControllerGossipConfig for the cluster assuming the use of gossip DNS.
- `cluster_dns_domain` - (Optional) - String - ClusterDNSDomain is the suffix we use for internal DNS names (normally cluster.local).
- `ssh_access` - (Optional) - List(String) - SSHAccess is a list of the CIDRs that can access SSH.
- `node_port_access` - (Optional) - List(String) - NodePortAccess is a list of the CIDRs that can access the node ports range (30000-32767).
- `ssh_key_name` - (Optional) - String - SSHKeyName specifies a preexisting SSH key to use.
- `update_policy` - (Optional) - String - UpdatePolicy determines the policy for applying upgrades automatically.<br />Valid values:<br />  'automatic' (default): apply updates automatically (apply OS security upgrades, avoiding rebooting when possible)<br />  'external': do not apply updates automatically; they are applied manually or by an external system.
- `external_policies` - (Optional) - Map(List(String)) - ExternalPolicies allows the insertion of pre-existing managed policies on IG Roles.
- `additional_policies` - (Optional) - Map(String) - Additional policies to add for roles.
- `file_assets` - (Optional) - List([file_asset_spec](#file_asset_spec)) - A collection of files assets for deployed cluster wide.
- `etcd_cluster` - (Required) - List([etcd_cluster_spec](#etcd_cluster_spec)) - EtcdClusters stores the configuration for each cluster.
- `docker` - (Optional) - [docker_config](#docker_config) - Docker was removed.
- `containerd` - (Optional) - [containerd_config](#containerd_config) - Component configurations.
- `kube_dns` - (Optional) - [kube_dns_config](#kube_dns_config)
- `kube_api_server` - (Optional) - [kube_api_server_config](#kube_api_server_config)
- `kube_controller_manager` - (Optional) - [kube_controller_manager_config](#kube_controller_manager_config)
- `external_cloud_controller_manager` - (Optional) - [cloud_controller_manager_config](#cloud_controller_manager_config)
- `kube_scheduler` - (Optional) - [kube_scheduler_config](#kube_scheduler_config)
- `kube_proxy` - (Optional) - [kube_proxy_config](#kube_proxy_config)
- `kubelet` - (Optional) - [kubelet_config_spec](#kubelet_config_spec) - Kubelet is the kubelet configuration for nodes not belonging to the control plane.<br />It can be overridden by the kubelet configuration specified in the instance group.
- `control_plane_kubelet` - (Optional) - [kubelet_config_spec](#kubelet_config_spec) - ControlPlaneKubelet is the kubelet configuration for nodes belonging to the control plane<br />It can be overridden by the kubelet configuration specified in the instance group.
- `cloud_config` - (Optional) - [cloud_configuration](#cloud_configuration)
- `external_dns` - (Optional) - [external_dns_config](#external_dns_config)
- `ntp` - (Optional) - [ntp_config](#ntp_config)
- `packages` - (Optional) - List(String) - Packages specifies additional packages to be installed.
- `node_problem_detector` - (Optional) - [node_problem_detector_config](#node_problem_detector_config) - NodeProblemDetector determines the node problem detector configuration.
- `metrics_server` - (Optional) - [metrics_server_config](#metrics_server_config) - MetricsServer determines the metrics server configuration.
- `cert_manager` - (Optional) - [cert_manager_config](#cert_manager_config) - CertManager determines the metrics server configuration.
- `networking` - (Required) - [networking_spec](#networking_spec) - Networking configures networking.
- `api` - (Optional) - [api_spec](#api_spec) - API controls how the Kubernetes API is exposed.
- `authentication` - (Optional) - [authentication_spec](#authentication_spec) - Authentication field controls how the cluster is configured for authentication.
- `authorization` - (Optional) - [authorization_spec](#authorization_spec) - Authorization field controls how the cluster is configured for authorization.
- `node_authorization` - (Optional) - [node_authorization_spec](#node_authorization_spec) - NodeAuthorization defined the custom node authorization configuration.
- `cloud_labels` - (Optional) - Map(String) - CloudLabels defines additional tags or labels on cloud provider resources.
- `hooks` - (Optional) - List([hook_spec](#hook_spec)) - Hooks for custom actions e.g. on first installation.
- `assets` - (Optional) - [assets_spec](#assets_spec) - Assets is alternative locations for files and containers; the API under construction, will remove this comment once this API is fully functional.
- `iam` - (Optional) - (Computed) - [iam_spec](#iam_spec) - IAM field adds control over the IAM security policies applied to resources.
- `encryption_config` - (Optional) - Bool - EncryptionConfig controls if encryption is enabled.
- `use_host_certificates` - (Optional) - Bool - UseHostCertificates will mount /etc/ssl/certs to inside needed containers.<br />This is needed if some APIs do have self-signed certs.
- `sysctl_parameters` - (Optional) - List(String) - SysctlParameters will configure kernel parameters using sysctl(8). When<br />specified, each parameter must follow the form variable=value, the way<br />it would appear in sysctl.conf.
- `rolling_update` - (Optional) - [rolling_update](#rolling_update) - RollingUpdate defines the default rolling-update settings for instance groups.
- `cluster_autoscaler` - (Optional) - [cluster_autoscaler_config](#cluster_autoscaler_config) - ClusterAutoscaler defines the cluster autoscaler configuration.
- `service_account_issuer_discovery` - (Optional) - [service_account_issuer_discovery_config](#service_account_issuer_discovery_config) - ServiceAccountIssuerDiscovery configures the OIDC Issuer for ServiceAccounts.
- `snapshot_controller` - (Optional) - [snapshot_controller_config](#snapshot_controller_config) - SnapshotController defines the CSI Snapshot Controller configuration.
- `karpenter` - (Optional) - [karpenter_config](#karpenter_config) - Karpenter defines the Karpenter configuration.
- `labels` - (Optional) - Map(String) - Map of string keys and values that can be used to organize and categorize<br />(scope and select) objects. May match selectors of replication controllers<br />and services.
- `annotations` - (Optional) - Map(String) - Annotations is an unstructured key value map stored with a resource that may be<br />set by external tools to store and retrieve arbitrary metadata. They are not<br />queryable and should be preserved when modifying objects.
- `name` - (Required) - (Force new) - String - Name defines the cluster name.
- `admin_ssh_key` - (Optional) - (Sensitive) - String - AdminSshKey defines the cluster admin ssh key.
- `secrets` - (Optional) - [cluster_secrets](#cluster_secrets) - Secrets defines the cluster secret.
- `revision` - (Computed) - Int - Revision is incremented every time the resource changes, this is useful for triggering cluster updater.
- `delete` - (Optional) - [delete_options](#delete_options) - Delete holds cluster delete options.

## Nested resources

### addon_spec

AddonSpec defines an addon that we want to install in the cluster.

#### Argument Reference

The following arguments are supported:

- `manifest` - (Required) - String - Manifest is a path to the manifest that defines the addon.

### config_store_spec

ConfigStoreSpec configures the stores that nodes use to get their configuration.

#### Argument Reference

The following arguments are supported:

- `base` - (Optional) - (Computed) - String - Base is the VFS path where we store configuration for the cluster<br />This might be different than the location where the cluster spec itself is stored,<br />both because this must be accessible to the cluster,<br />and because it might be on a different cloud or storage system (etcd vs S3).
- `keypairs` - (Optional) - String - Keypairs is the VFS path to where certificates and corresponding private keys are stored.
- `secrets` - (Optional) - String - Secrets is the VFS path to where secrets are stored.

### cloud_provider_spec

CloudProviderSpec configures the cloud provider to use.

#### Argument Reference

The following arguments are supported:

- `aws` - (Optional) - [aws_spec](#aws_spec) - AWS configures the AWS cloud provider.
- `azure` - (Optional) - [azure_spec](#azure_spec) - Azure configures the Azure cloud provider.
- `do` - (Optional) - [do_spec](#do_spec) - DO configures the Digital Ocean cloud provider.
- `gce` - (Optional) - [gce_spec](#gce_spec) - GCE configures the GCE cloud provider.
- `hetzner` - (Optional) - [hetzner_spec](#hetzner_spec) - Hetzner configures the Hetzner cloud provider.
- `openstack` - (Optional) - [openstack_spec](#openstack_spec) - Openstack configures the Openstack cloud provider.
- `scaleway` - (Optional) - [scaleway_spec](#scaleway_spec) - Scaleway configures the Scaleway cloud provider.

### aws_spec

AWSSpec configures the AWS cloud provider.

#### Argument Reference

The following arguments are supported:

- `ebs_csi_driver` - (Optional) - [ebs_csi_driver_spec](#ebs_csi_driver_spec) - EBSCSIDriverSpec is the config for the EBS CSI driver.
- `node_termination_handler` - (Optional) - [node_termination_handler_spec](#node_termination_handler_spec) - NodeTerminationHandler determines the node termination handler configuration.
- `load_balancer_controller` - (Optional) - [load_balancer_controller_spec](#load_balancer_controller_spec) - LoadbalancerController determines the Load Balancer Controller configuration.
- `pod_identity_webhook` - (Optional) - [pod_identity_webhook_spec](#pod_identity_webhook_spec) - PodIdentityWebhook determines the EKS Pod Identity Webhook configuration.
- `warm_pool` - (Optional) - [warm_pool_spec](#warm_pool_spec) - WarmPool defines the default warm pool settings for instance groups.
- `node_ip_families` - (Optional) - List(String) - NodeIPFamilies control the IP families reported for each node.
- `disable_security_group_ingress` - (Optional) - Bool - DisableSecurityGroupIngress disables the Cloud Controller Manager's creation<br />of an AWS Security Group for each load balancer provisioned for a Service.
- `elb_security_group` - (Optional) - String - ElbSecurityGroup specifies an existing AWS Security group for the Cloud Controller<br />Manager to assign to each ELB provisioned for a Service, instead of creating<br />one per ELB.
- `spotinst_product` - (Optional) - String - Spotinst cloud-config specs.
- `spotinst_orientation` - (Optional) - String
- `binaries_location` - (Optional) - String - BinariesLocation is the location of the AWS cloud provider binaries.

### ebs_csi_driver_spec

EBSCSIDriverSpec is the config for the AWS EBS CSI driver.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Required) - Bool - Enabled enables the AWS EBS CSI driver. Can only be set to true.<br />Default: true.
- `managed` - (Required) - Bool - Managed controls if aws-ebs-csi-driver is manged and deployed by kOps.<br />The deployment of aws-ebs-csi-driver is skipped if this is set to false.
- `version` - (Optional) - String - Version is the container image tag used.<br />Default: The latest stable release which is compatible with your Kubernetes version.
- `kube_api_qps` - (Optional) - Quantity - KubeAPIQPS QPS to use while talking with Kubernetes API server. (default 20).
- `kube_api_burst` - (Optional) - Int - KubeAPIBurst Burst to use while talking with Kubernetes API server. (default 100).
- `host_network` - (Optional) - Bool - HostNetwork can be used for large clusters for faster access to node info via instance metadata.<br />Default: false.
- `volume_attach_limit` - (Optional) - Int - VolumeAttachLimit is the maximum number of volumes attachable per node.<br />If specified, the limit applies to all nodes.<br />If not specified, the value is approximated from the instance type.<br />Default: -.
- `pod_annotations` - (Optional) - Map(String) - PodAnnotations are the annotations added to AWS EBS CSI node and controller Pods.<br />Default: none.

### node_termination_handler_spec

NodeTerminationHandlerSpec determines the node termination handler configuration.

#### Argument Reference

The following arguments are supported:

- `delete_sqs_msg_if_node_not_found` - (Optional) - Bool - DeleteSQSMsgIfNodeNotFound makes node termination handler delete the SQS Message from the SQS Queue if the targeted node is not found.<br />Only used in Queue Processor mode.<br />Default: false.
- `enabled` - (Required) - Bool - Enabled enables the node termination handler.<br />Default: true.
- `enable_spot_interruption_draining` - (Required) - Bool - EnableSpotInterruptionDraining makes node termination handler drain nodes when spot interruption termination notice is received.<br />Cannot be disabled in queue-processor mode.<br />Default: true.
- `enable_scheduled_event_draining` - (Optional) - Bool - EnableScheduledEventDraining makes node termination handler drain nodes before the maintenance window starts for an EC2 instance scheduled event.<br />Cannot be disabled in queue-processor mode.<br />Default: true.
- `enable_rebalance_monitoring` - (Optional) - Bool - EnableRebalanceMonitoring makes node termination handler cordon nodes when the rebalance recommendation notice is received.<br />In queue-processor mode, cannot be enabled without rebalance draining.<br />Default: false.
- `enable_rebalance_draining` - (Optional) - Bool - EnableRebalanceDraining makes node termination handler drain nodes when the rebalance recommendation notice is received.<br />Default: false.
- `enable_prometheus_metrics` - (Optional) - Bool - EnablePrometheusMetrics enables the "/metrics" endpoint.<br />Default: false.
- `enable_sqs_termination_draining` - (Optional) - Bool - EnableSQSTerminationDraining enables queue-processor mode which drains nodes when an SQS termination event is received.<br />Default: true.
- `exclude_from_load_balancers` - (Optional) - Bool - ExcludeFromLoadBalancers makes node termination handler will mark for exclusion from load balancers before node are cordoned.<br />Default: true.
- `managed_asg_tag` - (Optional) - String - ManagedASGTag is the tag used to determine which nodes NTH can take action on<br />This field has kept its name even though it now maps to the --managed-tag flag due to keeping the API stable.<br />Node termination handler does no longer check the ASG for this tag, but the actual EC2 instances.
- `pod_termination_grace_period` - (Optional) - Int - PodTerminationGracePeriod is the time in seconds given to each pod to terminate gracefully.<br />If negative, the default value specified in the pod will be used, which defaults to 30 seconds if not specified for the pod.<br />Default: -1.
- `taint_node` - (Optional) - Bool - TaintNode makes node termination handler taint nodes when an interruption event occurs.<br />Default: false.
- `memory_limit` - (Optional) - Quantity - MemoryLimit of NodeTerminationHandler container.<br />Default: none.
- `memory_request` - (Optional) - Quantity - MemoryRequest of NodeTerminationHandler container.<br />Default: 64Mi.
- `cpu_request` - (Optional) - Quantity - CPURequest of NodeTerminationHandler container.<br />Default: 50m.
- `version` - (Optional) - String - Version is the container image tag used.
- `webhook_template` - (Optional) - String - Replaces the default webhook message template.
- `webhook_url` - (Optional) - String - If specified, posts event data to URL upon instance interruption action.

### load_balancer_controller_spec

LoadBalancerControllerSpec determines the AWS LB controller configuration.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enabled enables the loadbalancer controller.<br />Default: false.
- `version` - (Optional) - String - Version is the container image tag used.
- `enable_waf` - (Optional) - Bool - EnableWAF specifies whether the controller can use WAFs (Classic Regional).<br />Default: false.
- `enable_wa_fv2` - (Optional) - Bool - EnableWAFv2 specifies whether the controller can use WAFs (V2).<br />Default: false.
- `enable_shield` - (Optional) - Bool - EnableShield specifies whether the controller can enable Shield Advanced.<br />Default: false.

### pod_identity_webhook_spec

PodIdentityWebhookSpec configures an EKS Pod Identity Webhook.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool
- `replicas` - (Optional) - Int

### warm_pool_spec

#### Argument Reference

The following arguments are supported:

- `min_size` - (Optional) - Int - MinSize is the minimum size of the warm pool.
- `max_size` - (Optional) - Int - MaxSize is the maximum size of the warm pool. The desired size of the instance group<br />is subtracted from this number to determine the desired size of the warm pool<br />(unless the resulting number is smaller than MinSize).<br />The default is the instance group's MaxSize.
- `enable_lifecycle_hook` - (Optional) - Bool - EnableLifecyleHook determines if an ASG lifecycle hook will be added ensuring that nodeup runs to completion.<br />Note that the metadata API must be protected from arbitrary Pods when this is enabled.

### azure_spec

AzureSpec defines Azure specific cluster configuration.

#### Argument Reference

The following arguments are supported:

- `subscription_id` - (Optional) - String - SubscriptionID specifies the subscription used for the cluster installation.
- `storage_account_id` - (Optional) - String - StorageAccountID specifies the storage account used for the cluster installation.
- `tenant_id` - (Optional) - String - TenantID is the ID of the tenant that the cluster is deployed in.
- `resource_group_name` - (Optional) - String - ResourceGroupName specifies the name of the resource group<br />where the cluster is built.<br />If this is empty, kops will create a new resource group<br />whose name is same as the cluster name. If this is not<br />empty, kops will not create a new resource group, and<br />it will just reuse the existing resource group of the name.<br />This follows the model that kops takes for AWS VPC.
- `route_table_name` - (Optional) - String - RouteTableName is the name of the route table attached to the subnet that the cluster is deployed in.
- `admin_user` - (Optional) - String - AdminUser specifies the admin user of VMs.

### do_spec

DOSpec configures the Digital Ocean cloud provider.


This resource has no attributes.

### gce_spec

GCESpec configures the GCE cloud provider.

#### Argument Reference

The following arguments are supported:

- `project` - (Optional) - String - Project is the cloud project we should use.
- `service_account` - (Optional) - String - ServiceAccount specifies the service account with which the GCE VM runs.
- `multizone` - (Optional) - Bool
- `node_tags` - (Optional) - String
- `node_instance_prefix` - (Optional) - String
- `pd_csi_driver` - (Optional) - [pd_csi_driver](#pd_csi_driver) - PDCSIDriver is the config for the PD CSI driver.
- `use_startup_script` - (Optional) - Bool - UseStartupScript specifies enables using startup-script instead of user-data metadata.
- `binaries_location` - (Optional) - String - BinariesLocation is the location of the GCE cloud provider binaries.

### pd_csi_driver

PDCSIDriver is the config for the GCP PD CSI driver.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enabled enables the GCP PD CSI driver.

### hetzner_spec

HetznerSpec configures the Hetzner cloud provider.


This resource has no attributes.

### openstack_spec

OpenstackSpec defines cloud config elements for the openstack cloud provider.

#### Argument Reference

The following arguments are supported:

- `loadbalancer` - (Optional) - [openstack_loadbalancer_config](#openstack_loadbalancer_config)
- `monitor` - (Optional) - [openstack_monitor](#openstack_monitor)
- `router` - (Optional) - [openstack_router](#openstack_router)
- `block_storage` - (Optional) - [openstack_block_storage_config](#openstack_block_storage_config)
- `insecure_skip_verify` - (Optional) - Bool
- `network` - (Optional) - [openstack_network](#openstack_network)
- `metadata` - (Optional) - [openstack_metadata](#openstack_metadata)

### openstack_loadbalancer_config

OpenstackLoadbalancerConfig defines the config for a neutron loadbalancer.

#### Argument Reference

The following arguments are supported:

- `method` - (Optional) - String
- `provider` - (Optional) - String
- `use_octavia` - (Optional) - Bool
- `floating_network` - (Optional) - String
- `floating_network_id` - (Optional) - String
- `floating_subnet` - (Optional) - String
- `subnet_id` - (Optional) - String
- `manage_sec_groups` - (Optional) - Bool
- `enable_ingress_hostname` - (Optional) - Bool
- `ingress_hostname_suffix` - (Optional) - String
- `flavor_id` - (Optional) - String

### openstack_monitor

OpenstackMonitor defines the config for a health monitor.

#### Argument Reference

The following arguments are supported:

- `delay` - (Optional) - String
- `timeout` - (Optional) - String
- `max_retries` - (Optional) - Int

### openstack_router

OpenstackRouter defines the config for a router.

#### Argument Reference

The following arguments are supported:

- `external_network` - (Optional) - String
- `dns_servers` - (Optional) - String
- `external_subnet` - (Optional) - String
- `availability_zone_hints` - (Optional) - List(String)

### openstack_block_storage_config

#### Argument Reference

The following arguments are supported:

- `version` - (Optional) - String
- `ignore_az` - (Optional) - Bool
- `override_az` - (Optional) - String
- `ignore_volume_micro_version` - (Optional) - Bool
- `metrics_enabled` - (Optional) - Bool
- `create_storage_class` - (Optional) - Bool - CreateStorageClass provisions a default class for the Cinder plugin.
- `csi_plugin_image` - (Optional) - String
- `csi_topology_support` - (Optional) - Bool
- `cluster_name` - (Optional) - String - ClusterName sets the --cluster flag for the cinder-csi-plugin to the provided name.

### openstack_network

OpenstackNetwork defines the config for a network.

#### Argument Reference

The following arguments are supported:

- `availability_zone_hints` - (Optional) - List(String)
- `ipv6_support_disabled` - (Optional) - Bool
- `public_network_names` - (Optional) - List(String)
- `internal_network_names` - (Optional) - List(String)
- `address_sort_order` - (Optional) - String

### openstack_metadata

OpenstackMetadata defines config for metadata service related settings.

#### Argument Reference

The following arguments are supported:

- `config_drive` - (Optional) - Bool - ConfigDrive specifies to use config drive for retrieving user data instead of the metadata service when launching instances.

### scaleway_spec

ScalewaySpec configures the Scaleway cloud provider.


This resource has no attributes.

### gossip_config

#### Argument Reference

The following arguments are supported:

- `protocol` - (Optional) - String
- `listen` - (Optional) - String
- `secret` - (Optional) - String
- `secondary` - (Optional) - [gossip_config_secondary](#gossip_config_secondary)

### gossip_config_secondary

#### Argument Reference

The following arguments are supported:

- `protocol` - (Optional) - String([Nullable](#nullable-arguments))
- `listen` - (Optional) - String
- `secret` - (Optional) - String

### dns_controller_gossip_config

#### Argument Reference

The following arguments are supported:

- `protocol` - (Optional) - String
- `listen` - (Optional) - String
- `secret` - (Optional) - String
- `secondary` - (Optional) - [dns_controller_gossip_config_secondary](#dns_controller_gossip_config_secondary)
- `seed` - (Optional) - String

### dns_controller_gossip_config_secondary

#### Argument Reference

The following arguments are supported:

- `protocol` - (Optional) - String([Nullable](#nullable-arguments))
- `listen` - (Optional) - String
- `secret` - (Optional) - String
- `seed` - (Optional) - String

### file_asset_spec

FileAssetSpec defines the structure for a file asset.

#### Argument Reference

The following arguments are supported:

- `name` - (Required) - String - Name is a shortened reference to the asset.
- `path` - (Required) - String - Path is the location this file should reside.
- `roles` - (Optional) - List(String) - Roles is a list of roles the file asset should be applied, defaults to all.
- `content` - (Required) - String - Content is the contents of the file.
- `is_base64` - (Optional) - Bool - IsBase64 indicates the contents is base64 encoded.
- `mode` - (Optional) - String - Mode is this file's mode and permission bits.

### etcd_cluster_spec

EtcdClusterSpec is the etcd cluster specification.

#### Argument Reference

The following arguments are supported:

- `name` - (Required) - String - Name is the name of the etcd cluster (main, events etc).
- `provider` - (Optional) - String - Provider is the provider used to run etcd: Manager, Legacy.<br />Defaults to Manager.
- `member` - (Required) - List([etcd_member_spec](#etcd_member_spec)) - Members stores the configurations for each member of the cluster (including the data volume).
- `version` - (Optional) - String - Version is the version of etcd to run.
- `leader_election_timeout` - (Optional) - Duration - LeaderElectionTimeout is the time (in milliseconds) for an etcd leader election timeout.
- `heartbeat_interval` - (Optional) - Duration - HeartbeatInterval is the time (in milliseconds) for an etcd heartbeat interval.
- `image` - (Optional) - String - Image is the etcd container image to use. Setting this will ignore the Version specified.
- `backups` - (Optional) - [etcd_backup_spec](#etcd_backup_spec) - Backups describes how we do backups of etcd.
- `manager` - (Optional) - [etcd_manager_spec](#etcd_manager_spec) - Manager describes the manager configuration.
- `memory_request` - (Optional) - Quantity - MemoryRequest specifies the memory requests of each etcd container in the cluster.
- `cpu_request` - (Optional) - Quantity - CPURequest specifies the cpu requests of each etcd container in the cluster.

### etcd_member_spec

EtcdMemberSpec is a specification for a etcd member.

#### Argument Reference

The following arguments are supported:

- `name` - (Required) - String - Name is the name of the member within the etcd cluster.
- `instance_group` - (Required) - String - InstanceGroup is the instanceGroup this volume is associated.
- `volume_type` - (Optional) - String - VolumeType is the underlying cloud storage class.
- `volume_iops` - (Optional) - Int - If volume type is io1, then we need to specify the number of IOPS.
- `volume_throughput` - (Optional) - Int - Parameter for disks that support provisioned throughput.
- `volume_size` - (Optional) - Int - VolumeSize is the underlying cloud volume size.
- `kms_key_id` - (Optional) - String - KmsKeyID is a AWS KMS ID used to encrypt the volume.
- `encrypted_volume` - (Optional) - Bool - EncryptedVolume indicates you want to encrypt the volume.

### etcd_backup_spec

EtcdBackupSpec describes how we want to do backups of etcd.

#### Argument Reference

The following arguments are supported:

- `backup_store` - (Required) - String - BackupStore is the VFS path where we will read/write backup data.
- `image` - (Required) - String - Image is the etcd backup manager image to use.  Setting this will create a sidecar container in the etcd pod with the specified image.

### etcd_manager_spec

EtcdManagerSpec describes how we configure the etcd manager.

#### Argument Reference

The following arguments are supported:

- `image` - (Optional) - String - Image is the etcd manager image to use.
- `env` - (Optional) - List([env_var](#env_var)) - Env allows users to pass in env variables to the etcd-manager container.<br />Variables starting with ETCD_ will be further passed down to the etcd process.<br />This allows etcd setting to be overwriten. No config validation is done.<br />A list of etcd config ENV vars can be found at https://github.com/etcd-io/etcd/blob/master/Documentation/op-guide/configuration.md.
- `backup_interval` - (Optional) - Duration - BackupInterval which is used for backups. The default is 15 minutes.
- `backup_retention_days` - (Optional) - Int - BackupRetentionDays which is used for backups. The default is 90 days.
- `discovery_poll_interval` - (Optional) - Duration - DiscoveryPollInterval which is used for discovering other cluster members. The default is 60 seconds.
- `listen_metrics_ur_ls` - (Optional) - List(String) - ListenMetricsURLs is the list of URLs to listen on that will respond to both the /metrics and /health endpoints.
- `log_level` - (Optional) - Int - LogLevel allows the klog library verbose log level to be set for etcd-manager. The default is 6.<br />https://github.com/google/glog#verbose-logging.

### env_var

EnvVar represents an environment variable present in a Container.

#### Argument Reference

The following arguments are supported:

- `name` - (Required) - String - Name of the environment variable. Must be a C_IDENTIFIER.
- `value` - (Optional) - String - Variable references $(VAR_NAME) are expanded<br />using the previous defined environment variables in the container and<br />any service environment variables. If a variable cannot be resolved,<br />the reference in the input string will be unchanged. The $(VAR_NAME)<br />syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped<br />references will never be expanded, regardless of whether the variable<br />exists or not.<br />Defaults to "".<br />+optional.

### docker_config

DockerConfig is the configuration for docker.

#### Argument Reference

The following arguments are supported:

- `authorization_plugins` - (Optional) - List(String) - AuthorizationPlugins is a list of authorization plugins.
- `bridge` - (Optional) - String - Bridge is the network interface containers should bind onto.
- `bridge_ip` - (Optional) - String - BridgeIP is a specific IP address and netmask for the docker0 bridge, using standard CIDR notation.
- `data_root` - (Optional) - String - DataRoot is the root directory of persistent docker state (default "/var/lib/docker").
- `default_ulimit` - (Optional) - List(String) - DefaultUlimit is the ulimits for containers.
- `default_runtime` - (Optional) - String - DefaultRuntime is the default OCI runtime for containers (default "runc").
- `dns` - (Optional) - List(String) - DNS is the IP address of the DNS server.
- `exec_opt` - (Optional) - List(String) - ExecOpt is a series of options passed to the runtime.
- `exec_root` - (Optional) - String - ExecRoot is the root directory for execution state files (default "/var/run/docker").
- `experimental` - (Optional) - Bool - Experimental features permits enabling new features such as dockerd metrics.
- `health_check` - (Optional) - Bool - HealthCheck enables the periodic health-check service.
- `hosts` - (Optional) - List(String) - Hosts enables you to configure the endpoints the docker daemon listens on i.e. tcp://0.0.0.0.2375 or unix:///var/run/docker.sock etc.
- `ip_masq` - (Optional) - Bool - IPMasq enables ip masquerading for containers.
- `ip_tables` - (Optional) - Bool - IPtables enables addition of iptables rules.
- `insecure_registry` - (Optional) - String - InsecureRegistry enable insecure registry communication @question according to dockers this a list??.
- `insecure_registries` - (Optional) - List(String) - InsecureRegistries enables multiple insecure docker registry communications.
- `live_restore` - (Optional) - Bool - LiveRestore enables live restore of docker when containers are still running.
- `log_driver` - (Optional) - String - LogDriver is the default driver for container logs (default "json-file").
- `log_level` - (Optional) - String - LogLevel is the logging level ("debug", "info", "warn", "error", "fatal") (default "info").
- `log_opt` - (Optional) - List(String) - Logopt is a series of options given to the log driver options for containers.
- `max_concurrent_downloads` - (Optional) - Int - MaxConcurrentDownloads sets the max concurrent downloads for each pull.
- `max_concurrent_uploads` - (Optional) - Int - MaxConcurrentUploads sets the max concurrent uploads for each push.
- `max_download_attempts` - (Optional) - Int - MaxDownloadAttempts sets the max download attempts for each pull.
- `metrics_address` - (Optional) - String - Metrics address is the endpoint to serve with Prometheus format metrics.
- `mtu` - (Optional) - Int - MTU is the containers network MTU.
- `packages` - (Optional) - [packages_config](#packages_config) - Packages overrides the URL and hash for the packages.
- `registry_mirrors` - (Optional) - List(String) - RegistryMirrors is a referred list of docker registry mirror.
- `runtimes` - (Optional) - List(String) - Runtimes registers an additional OCI compatible runtime (default []).
- `selinux_enabled` - (Optional) - Bool - SelinuxEnabled enables SELinux support.
- `skip_install` - (Optional) - Bool - SkipInstall when set to true will prevent kops from installing and modifying Docker in any way.
- `storage` - (Optional) - String - Storage is the docker storage driver to use.
- `storage_opts` - (Optional) - List(String) - StorageOpts is a series of options passed to the storage driver.
- `user_namespace_remap` - (Optional) - String - UserNamespaceRemap sets the user namespace remapping option for the docker daemon.
- `version` - (Optional) - String - Version is consumed by the nodeup and used to pick the docker version.

### packages_config

#### Argument Reference

The following arguments are supported:

- `hash_amd64` - (Optional) - String - HashAmd64 overrides the hash for the AMD64 package.
- `hash_arm64` - (Optional) - String - HashArm64 overrides the hash for the ARM64 package.
- `url_amd64` - (Optional) - String - UrlAmd64 overrides the URL for the AMD64 package.
- `url_arm64` - (Optional) - String - UrlArm64 overrides the URL for the ARM64 package.

### containerd_config

ContainerdConfig is the configuration for containerd.

#### Argument Reference

The following arguments are supported:

- `address` - (Optional) - String - Address of containerd's GRPC server (default "/run/containerd/containerd.sock").
- `config_additions` - (Optional) - Map(IntOrString) - ConfigAdditions adds additional config entries to the generated config file.
- `config_override` - (Optional) - String - ConfigOverride is the complete containerd config file provided by the user.
- `log_level` - (Optional) - String - LogLevel controls the logging details [trace, debug, info, warn, error, fatal, panic] (default "info").
- `packages` - (Optional) - [packages_config](#packages_config) - Packages overrides the URL and hash for the packages.
- `registry_mirrors` - (Optional) - Map(List(String)) - RegistryMirrors is list of image registries.
- `root` - (Optional) - String - Root directory for persistent data (default "/var/lib/containerd").
- `skip_install` - (Optional) - Bool - SkipInstall prevents kOps from installing and modifying containerd in any way (default "false").
- `state` - (Optional) - String - State directory for execution state files (default "/run/containerd").
- `version` - (Optional) - String - Version used to pick the containerd package.
- `nvidia_gpu` - (Optional) - [nvidia_gpu_config](#nvidia_gpu_config) - NvidiaGPU configures the Nvidia GPU runtime.
- `runc` - (Optional) - [runc](#runc) - Runc configures the runc runtime.
- `se_linux_enabled` - (Optional) - Bool - SelinuxEnabled enables SELinux support.
- `nri` - (Optional) - [nri_config](#nri_config) - NRI configures the Node Resource Interface.

### nvidia_gpu_config

#### Argument Reference

The following arguments are supported:

- `driver_package` - (Optional) - String - Package is the name of the nvidia driver package that will be installed.<br />Default is "nvidia-headless-510-server".
- `enabled` - (Optional) - Bool - Enabled determines if kOps will install the Nvidia GPU runtime and drivers.<br />They will only be installed on intances that has an Nvidia GPU.
- `dcgm_exporter` - (Optional) - [dcgm_exporter_config](#dcgm_exporter_config) - DCGMExporterConfig configures the DCGM exporter.

### dcgm_exporter_config

DCGMExporterConfig configures the DCGM exporter.<br />Only the DCGMExporterConfig in the cluster level takes effect. Configurations on the Instance Group are ignored.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enabled determines if kOps will install the DCGM exporter.

### runc

#### Argument Reference

The following arguments are supported:

- `version` - (Optional) - String - Version used to pick the runc package.
- `packages` - (Optional) - [packages_config](#packages_config) - Packages overrides the URL and hash for the packages.

### nri_config

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enable NRI support in containerd.
- `plugin_registration_timeout` - (Optional) - Duration - PluginRegistrationTimeout is the timeout for plugin registration.
- `plugin_request_timeout` - (Optional) - Duration - PluginRequestTimeout is the timeout for a plugin to handle a request.

### kube_dns_config

KubeDNSConfig defines the kube dns configuration.

#### Argument Reference

The following arguments are supported:

- `cache_max_size` - (Optional) - Int - CacheMaxSize is the maximum entries to keep in dnsmasq.
- `cache_max_concurrent` - (Optional) - Int - CacheMaxConcurrent is the maximum number of concurrent queries for dnsmasq.
- `tolerations` - (Optional) - List([toleration](#toleration)) - Tolerations	are tolerations to apply to the kube-dns deployment.
- `affinity` - (Optional) - [affinity](#affinity) - Affinity is the kube-dns affinity, uses the same syntax as kubectl's affinity.
- `core_dns_image` - (Optional) - String - CoreDNSImage is used to override the default image used for CoreDNS.
- `cpa_image` - (Optional) - String - CPAImage is used to override the default image used for Cluster Proportional Autoscaler.
- `domain` - (Optional) - String - Domain is the dns domain.
- `external_core_file` - (Optional) - String - ExternalCoreFile is used to provide a complete CoreDNS CoreFile by the user - ignores other provided flags which modify the CoreFile.
- `provider` - (Optional) - String - Provider indicates whether CoreDNS or kube-dns will be the default service discovery.
- `server_ip` - (Optional) - String - ServerIP is the server ip.
- `stub_domains` - (Optional) - Map(List(String)) - StubDomains redirects a domains to another DNS service.
- `upstream_nameservers` - (Optional) - List(String) - UpstreamNameservers sets the upstream nameservers for queries not on the cluster domain.
- `memory_request` - (Optional) - Quantity - MemoryRequest specifies the memory requests of each dns container in the cluster. Default 70m.
- `cpu_request` - (Optional) - Quantity - CPURequest specifies the cpu requests of each dns container in the cluster. Default 100m.
- `memory_limit` - (Optional) - Quantity - MemoryLimit specifies the memory limit of each dns container in the cluster. Default 170m.
- `node_local_dns` - (Optional) - [node_local_dns_config](#node_local_dns_config) - NodeLocalDNS specifies the configuration for the node-local-dns addon.

### toleration

The pod this Toleration is attached to tolerates any taint that matches<br />the triple <key,value,effect> using the matching operator <operator>.

#### Argument Reference

The following arguments are supported:

- `key` - (Optional) - String - Key is the taint key that the toleration applies to. Empty means match all taint keys.<br />If the key is empty, operator must be Exists; this combination means to match all values and all keys.<br />+optional.
- `operator` - (Optional) - String - Operator represents a key's relationship to the value.<br />Valid operators are Exists and Equal. Defaults to Equal.<br />Exists is equivalent to wildcard for value, so that a pod can<br />tolerate all taints of a particular category.<br />+optional.
- `value` - (Optional) - String - Value is the taint value the toleration matches to.<br />If the operator is Exists, the value should be empty, otherwise just a regular string.<br />+optional.
- `effect` - (Optional) - String - Effect indicates the taint effect to match. Empty means match all taint effects.<br />When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.<br />+optional.
- `toleration_seconds` - (Optional) - Int - TolerationSeconds represents the period of time the toleration (which must be<br />of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,<br />it is not set, which means tolerate the taint forever (do not evict). Zero and<br />negative values will be treated as 0 (evict immediately) by the system.<br />+optional.

### affinity

Affinity is a group of affinity scheduling rules.

#### Argument Reference

The following arguments are supported:

- `node_affinity` - (Optional) - [node_affinity](#node_affinity) - Describes node affinity scheduling rules for the pod.<br />+optional.
- `pod_affinity` - (Optional) - [pod_affinity](#pod_affinity) - Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).<br />+optional.
- `pod_anti_affinity` - (Optional) - [pod_anti_affinity](#pod_anti_affinity) - Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).<br />+optional.

### node_affinity

Node affinity is a group of node affinity scheduling rules.

#### Argument Reference

The following arguments are supported:

- `required_during_scheduling_ignored_during_execution` - (Optional) - [node_selector](#node_selector) - If the affinity requirements specified by this field are not met at<br />scheduling time, the pod will not be scheduled onto the node.<br />If the affinity requirements specified by this field cease to be met<br />at some point during pod execution (e.g. due to an update), the system<br />may or may not try to eventually evict the pod from its node.<br />+optional.
- `preferred_during_scheduling_ignored_during_execution` - (Optional) - List([preferred_scheduling_term](#preferred_scheduling_term)) - The scheduler will prefer to schedule pods to nodes that satisfy<br />the affinity expressions specified by this field, but it may choose<br />a node that violates one or more of the expressions. The node that is<br />most preferred is the one with the greatest sum of weights, i.e.<br />for each node that meets all of the scheduling requirements (resource<br />request, requiredDuringScheduling affinity expressions, etc.),<br />compute a sum by iterating through the elements of this field and adding<br />"weight" to the sum if the node matches the corresponding matchExpressions; the<br />node(s) with the highest sum are the most preferred.<br />+optional<br />+listType=atomic.

### node_selector

A node selector represents the union of the results of one or more label queries<br />over a set of nodes; that is, it represents the OR of the selectors represented<br />by the node selector terms.<br />+structType=atomic.

#### Argument Reference

The following arguments are supported:

- `node_selector_terms` - (Optional) - List([node_selector_term](#node_selector_term)) - Required. A list of node selector terms. The terms are ORed.<br />+listType=atomic.

### node_selector_term

A null or empty node selector term matches no objects. The requirements of<br />them are ANDed.<br />The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.<br />+structType=atomic.

#### Argument Reference

The following arguments are supported:

- `match_expressions` - (Optional) - List([node_selector_requirement](#node_selector_requirement)) - A list of node selector requirements by node's labels.<br />+optional<br />+listType=atomic.
- `match_fields` - (Optional) - List([node_selector_requirement](#node_selector_requirement)) - A list of node selector requirements by node's fields.<br />+optional<br />+listType=atomic.

### node_selector_requirement

A node selector requirement is a selector that contains values, a key, and an operator<br />that relates the key and values.

#### Argument Reference

The following arguments are supported:

- `key` - (Optional) - String - The label key that the selector applies to.
- `operator` - (Optional) - String - Represents a key's relationship to a set of values.<br />Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
- `values` - (Optional) - List(String) - An array of string values. If the operator is In or NotIn,<br />the values array must be non-empty. If the operator is Exists or DoesNotExist,<br />the values array must be empty. If the operator is Gt or Lt, the values<br />array must have a single element, which will be interpreted as an integer.<br />This array is replaced during a strategic merge patch.<br />+optional<br />+listType=atomic.

### preferred_scheduling_term

An empty preferred scheduling term matches all objects with implicit weight 0<br />(i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).

#### Argument Reference

The following arguments are supported:

- `weight` - (Optional) - Int - Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
- `preference` - (Optional) - [node_selector_term](#node_selector_term) - A node selector term, associated with the corresponding weight.

### pod_affinity

Pod affinity is a group of inter pod affinity scheduling rules.

#### Argument Reference

The following arguments are supported:

- `required_during_scheduling_ignored_during_execution` - (Optional) - List([pod_affinity_term](#pod_affinity_term)) - If the affinity requirements specified by this field are not met at<br />scheduling time, the pod will not be scheduled onto the node.<br />If the affinity requirements specified by this field cease to be met<br />at some point during pod execution (e.g. due to a pod label update), the<br />system may or may not try to eventually evict the pod from its node.<br />When there are multiple elements, the lists of nodes corresponding to each<br />podAffinityTerm are intersected, i.e. all terms must be satisfied.<br />+optional<br />+listType=atomic.
- `preferred_during_scheduling_ignored_during_execution` - (Optional) - List([weighted_pod_affinity_term](#weighted_pod_affinity_term)) - The scheduler will prefer to schedule pods to nodes that satisfy<br />the affinity expressions specified by this field, but it may choose<br />a node that violates one or more of the expressions. The node that is<br />most preferred is the one with the greatest sum of weights, i.e.<br />for each node that meets all of the scheduling requirements (resource<br />request, requiredDuringScheduling affinity expressions, etc.),<br />compute a sum by iterating through the elements of this field and adding<br />"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br />node(s) with the highest sum are the most preferred.<br />+optional<br />+listType=atomic.

### pod_affinity_term

Defines a set of pods (namely those matching the labelSelector<br />relative to the given namespace(s)) that this pod should be<br />co-located (affinity) or not co-located (anti-affinity) with,<br />where co-located is defined as running on a node whose value of<br />the label with key <topologyKey> matches that of any node on which<br />a pod of the set of pods is running.

#### Argument Reference

The following arguments are supported:

- `label_selector` - (Optional) - [label_selector](#label_selector) - A label query over a set of resources, in this case pods.<br />If it's null, this PodAffinityTerm matches with no Pods.<br />+optional.
- `namespaces` - (Optional) - List(String) - namespaces specifies a static list of namespace names that the term applies to.<br />The term is applied to the union of the namespaces listed in this field<br />and the ones selected by namespaceSelector.<br />null or empty namespaces list and null namespaceSelector means "this pod's namespace".<br />+optional<br />+listType=atomic.
- `topology_key` - (Optional) - String - This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br />the labelSelector in the specified namespaces, where co-located is defined as running on a node<br />whose value of the label with key topologyKey matches that of any node on which any of the<br />selected pods is running.<br />Empty topologyKey is not allowed.
- `namespace_selector` - (Optional) - [label_selector](#label_selector) - A label query over the set of namespaces that the term applies to.<br />The term is applied to the union of the namespaces selected by this field<br />and the ones listed in the namespaces field.<br />null selector and null or empty namespaces list means "this pod's namespace".<br />An empty selector ({}) matches all namespaces.<br />+optional.
- `match_label_keys` - (Optional) - List(String) - MatchLabelKeys is a set of pod label keys to select which pods will<br />be taken into consideration. The keys are used to lookup values from the<br />incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br />to select the group of existing pods which pods will be taken into consideration<br />for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br />pod labels will be ignored. The default value is empty.<br />The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br />Also, matchLabelKeys cannot be set when labelSelector isn't set.<br />This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).<br /><br />+listType=atomic<br />+optional.
- `mismatch_label_keys` - (Optional) - List(String) - MismatchLabelKeys is a set of pod label keys to select which pods will<br />be taken into consideration. The keys are used to lookup values from the<br />incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br />to select the group of existing pods which pods will be taken into consideration<br />for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br />pod labels will be ignored. The default value is empty.<br />The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br />Also, mismatchLabelKeys cannot be set when labelSelector isn't set.<br />This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).<br /><br />+listType=atomic<br />+optional.

### label_selector

A label selector is a label query over a set of resources. The result of matchLabels and<br />matchExpressions are ANDed. An empty label selector matches all objects. A null<br />label selector matches no objects.<br />+structType=atomic.

#### Argument Reference

The following arguments are supported:

- `match_labels` - (Optional) - Map(String) - matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels<br />map is equivalent to an element of matchExpressions, whose key field is "key", the<br />operator is "In", and the values array contains only "value". The requirements are ANDed.<br />+optional.
- `match_expressions` - (Optional) - List([label_selector_requirement](#label_selector_requirement)) - matchExpressions is a list of label selector requirements. The requirements are ANDed.<br />+optional<br />+listType=atomic.

### label_selector_requirement

A label selector requirement is a selector that contains values, a key, and an operator that<br />relates the key and values.

#### Argument Reference

The following arguments are supported:

- `key` - (Optional) - String - key is the label key that the selector applies to.
- `operator` - (Optional) - String - operator represents a key's relationship to a set of values.<br />Valid operators are In, NotIn, Exists and DoesNotExist.
- `values` - (Optional) - List(String) - values is an array of string values. If the operator is In or NotIn,<br />the values array must be non-empty. If the operator is Exists or DoesNotExist,<br />the values array must be empty. This array is replaced during a strategic<br />merge patch.<br />+optional<br />+listType=atomic.

### weighted_pod_affinity_term

The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s).

#### Argument Reference

The following arguments are supported:

- `weight` - (Optional) - Int - weight associated with matching the corresponding podAffinityTerm,<br />in the range 1-100.
- `pod_affinity_term` - (Optional) - [pod_affinity_term](#pod_affinity_term) - Required. A pod affinity term, associated with the corresponding weight.

### pod_anti_affinity

Pod anti affinity is a group of inter pod anti affinity scheduling rules.

#### Argument Reference

The following arguments are supported:

- `required_during_scheduling_ignored_during_execution` - (Optional) - List([pod_affinity_term](#pod_affinity_term)) - If the anti-affinity requirements specified by this field are not met at<br />scheduling time, the pod will not be scheduled onto the node.<br />If the anti-affinity requirements specified by this field cease to be met<br />at some point during pod execution (e.g. due to a pod label update), the<br />system may or may not try to eventually evict the pod from its node.<br />When there are multiple elements, the lists of nodes corresponding to each<br />podAffinityTerm are intersected, i.e. all terms must be satisfied.<br />+optional<br />+listType=atomic.
- `preferred_during_scheduling_ignored_during_execution` - (Optional) - List([weighted_pod_affinity_term](#weighted_pod_affinity_term)) - The scheduler will prefer to schedule pods to nodes that satisfy<br />the anti-affinity expressions specified by this field, but it may choose<br />a node that violates one or more of the expressions. The node that is<br />most preferred is the one with the greatest sum of weights, i.e.<br />for each node that meets all of the scheduling requirements (resource<br />request, requiredDuringScheduling anti-affinity expressions, etc.),<br />compute a sum by iterating through the elements of this field and adding<br />"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br />node(s) with the highest sum are the most preferred.<br />+optional<br />+listType=atomic.

### node_local_dns_config

NodeLocalDNSConfig are options of the node-local-dns.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enabled activates the node-local-dns addon.
- `external_core_file` - (Optional) - String - ExternalCoreFile is used to provide a complete NodeLocalDNS CoreFile by the user - ignores other provided flags which modify the CoreFile.
- `additional_config` - (Optional) - String - AdditionalConfig is used to provide additional config for node local dns by the user - it will include the original CoreFile made by kOps.
- `image` - (Optional) - String - Image overrides the default container image used for node-local-dns addon.
- `local_ip` - (Optional) - String - Local listen IP address. It can be any IP in the 169.254.20.0/16 space or any other IP address that can be guaranteed to not collide with any existing IP.
- `forward_to_kube_dns` - (Optional) - Bool - If enabled, nodelocal dns will use kubedns as a default upstream.
- `memory_request` - (Optional) - Quantity - MemoryRequest specifies the memory requests of each node-local-dns container in the daemonset. Default 5Mi.
- `cpu_request` - (Optional) - Quantity - CPURequest specifies the cpu requests of each node-local-dns container in the daemonset. Default 25m.
- `pod_annotations` - (Optional) - Map(String) - PodAnnotations makes possible to add additional annotations to node-local-dns.<br />Default: none.

### kube_api_server_config

KubeAPIServerConfig defines the configuration for the kube api.

#### Argument Reference

The following arguments are supported:

- `image` - (Optional) - String - Image is the container image used.
- `disable_basic_auth` - (Optional) - Bool - DisableBasicAuth removes the --basic-auth-file flag.
- `log_format` - (Optional) - String - LogFormat is the logging format of the api.<br />Supported values: text, json.<br />Default: text.
- `log_level` - (Optional) - Int - LogLevel is the logging level of the api.
- `cloud_provider` - (Optional) - String - CloudProvider is the name of the cloudProvider we are using, aws, gce etcd.
- `secure_port` - (Optional) - Int - SecurePort is the port the kube runs on.
- `insecure_port` - (Optional) - Int - InsecurePort is the port the insecure api runs.
- `address` - (Optional) - String - Address is the binding address for the kube api: Deprecated - use insecure-bind-address and bind-address.
- `advertise_address` - (Optional) - String - AdvertiseAddress is the IP address on which to advertise the apiserver to members of the cluster.
- `bind_address` - (Optional) - String - BindAddress is the binding address for the secure kubernetes API.
- `insecure_bind_address` - (Optional) - String - InsecureBindAddress is the binding address for the InsecurePort for the insecure kubernetes API.
- `enable_bootstrap_auth_token` - (Optional) - Bool - EnableBootstrapAuthToken enables 'bootstrap.kubernetes.io/token' in the 'kube-system' namespace to be used for TLS bootstrapping authentication.
- `enable_aggregator_routing` - (Optional) - Bool - EnableAggregatorRouting enables aggregator routing requests to endpoints IP rather than cluster IP.
- `admission_control` - (Optional) - List(String) - AdmissionControl is a list of admission controllers to use: Deprecated - use enable-admission-plugins instead.
- `append_admission_plugins` - (Optional) - List(String) - AppendAdmissionPlugins appends list of enabled admission plugins.
- `enable_admission_plugins` - (Optional) - List(String) - EnableAdmissionPlugins is a list of enabled admission plugins.
- `disable_admission_plugins` - (Optional) - List(String) - DisableAdmissionPlugins is a list of disabled admission plugins.
- `admission_control_config_file` - (Optional) - String - AdmissionControlConfigFile is the location of the admission-control-config-file.
- `service_cluster_ip_range` - (Optional) - String - ServiceClusterIPRange is the service address range.
- `service_node_port_range` - (Optional) - String - Passed as --service-node-port-range to kube-apiserver. Expects 'startPort-endPort' format e.g. 30000-33000.
- `etcd_servers` - (Optional) - List(String) - EtcdServers is a list of the etcd service to connect.
- `etcd_servers_overrides` - (Optional) - List(String) - EtcdServersOverrides is per-resource etcd servers overrides, comma separated. The individual override format: group/resource#servers, where servers are http://ip:port, semicolon separated.
- `etcd_ca_file` - (Optional) - String - EtcdCAFile is the path to a ca certificate.
- `etcd_cert_file` - (Optional) - String - EtcdCertFile is the path to a certificate.
- `etcd_key_file` - (Optional) - String - EtcdKeyFile is the path to a private key.
- `basic_auth_file` - (Optional) - String - TODO: Remove unused BasicAuthFile.
- `client_ca_file` - (Optional) - String - ClientCAFile is the file used by apisever that contains the client CA.
- `tls_cert_file` - (Optional) - String - TODO: Remove unused TLSCertFile.
- `tls_private_key_file` - (Optional) - String - TODO: Remove unused TLSPrivateKeyFile.
- `tls_cipher_suites` - (Optional) - List(String) - TLSCipherSuites indicates the allowed TLS cipher suite.
- `tls_min_version` - (Optional) - String - TLSMinVersion indicates the minimum TLS version allowed.
- `token_auth_file` - (Optional) - String - TODO: Remove unused TokenAuthFile.
- `allow_privileged` - (Optional) - Bool - AllowPrivileged indicates if we can run privileged containers.
- `api_server_count` - (Optional) - Int - APIServerCount is the number of api servers.
- `runtime_config` - (Optional) - Map(String) - RuntimeConfig is a series of keys/values are parsed into the `--runtime-config` parameters.
- `kubelet_client_certificate` - (Optional) - String - KubeletClientCertificate is the path of a certificate for secure communication between api and kubelet.
- `kubelet_certificate_authority` - (Optional) - String - KubeletCertificateAuthority is the path of a certificate authority for secure communication between api and kubelet.
- `kubelet_client_key` - (Optional) - String - KubeletClientKey is the path of a private to secure communication between api and kubelet.
- `anonymous_auth` - (Optional) - Bool([Nullable](#nullable-arguments)) - AnonymousAuth indicates if anonymous authentication is permitted.
- `kubelet_preferred_address_types` - (Optional) - List(String) - KubeletPreferredAddressTypes is a list of the preferred NodeAddressTypes to use for kubelet connections.
- `storage_backend` - (Optional) - String - StorageBackend is the backend storage.
- `oidc_username_claim` - (Optional) - String - OIDCUsernameClaim is the OpenID claim to use as the user name.<br />Note that claims other than the default ('sub') is not guaranteed to be<br />unique and immutable.
- `oidc_username_prefix` - (Optional) - String - OIDCUsernamePrefix is the prefix prepended to username claims to prevent<br />clashes with existing names (such as 'system:' users).
- `oidc_groups_claim` - (Optional) - String - OIDCGroupsClaim if provided, the name of a custom OpenID Connect claim for<br />specifying user groups.<br />The claim value is expected to be a string or array of strings.
- `oidc_groups_prefix` - (Optional) - String - OIDCGroupsPrefix is the prefix prepended to group claims to prevent<br />clashes with existing names (such as 'system:' groups).
- `oidc_issuer_url` - (Optional) - String - OIDCIssuerURL is the URL of the OpenID issuer, only HTTPS scheme will<br />be accepted.<br />If set, it will be used to verify the OIDC JSON Web Token (JWT).
- `oidc_client_id` - (Optional) - String - OIDCClientID is the client ID for the OpenID Connect client, must be set<br />if oidc-issuer-url is set.
- `oidc_required_claim` - (Optional) - List(String) - A key=value pair that describes a required claim in the ID Token.<br />If set, the claim is verified to be present in the ID Token with a matching value.<br />Repeat this flag to specify multiple claims.
- `oidc_ca_file` - (Optional) - String - OIDCCAFile if set, the OpenID server's certificate will be verified by one<br />of the authorities in the oidc-ca-file.
- `authentication_config_file` - (Optional) - String - AuthenticationConfigFile is the location of the authentication-config<br />this option is mutually exclusive with all OIDC options.
- `proxy_client_cert_file` - (Optional) - String - The apiserver's client certificate used for outbound requests.
- `proxy_client_key_file` - (Optional) - String - The apiserver's client key used for outbound requests.
- `audit_log_format` - (Optional) - String - AuditLogFormat flag specifies the format type for audit log files.
- `audit_log_path` - (Optional) - String - If set, all requests coming to the apiserver will be logged to this file.
- `audit_log_max_age` - (Optional) - Int - The maximum number of days to retain old audit log files based on the timestamp encoded in their filename.
- `audit_log_max_backups` - (Optional) - Int - The maximum number of old audit log files to retain.
- `audit_log_max_size` - (Optional) - Int - The maximum size in megabytes of the audit log file before it gets rotated. Defaults to 100MB.
- `audit_policy_file` - (Optional) - String - AuditPolicyFile is the full path to a advanced audit configuration file e.g. /srv/kubernetes/audit.conf.
- `audit_webhook_batch_buffer_size` - (Optional) - Int - AuditWebhookBatchBufferSize is The size of the buffer to store events before batching and writing. Only used in batch mode. (default 10000).
- `audit_webhook_batch_max_size` - (Optional) - Int - AuditWebhookBatchMaxSize is The maximum size of a batch. Only used in batch mode. (default 400).
- `audit_webhook_batch_max_wait` - (Optional) - Duration - AuditWebhookBatchMaxWait is The amount of time to wait before force writing the batch that hadn't reached the max size. Only used in batch mode. (default 30s).
- `audit_webhook_batch_throttle_burst` - (Optional) - Int - AuditWebhookBatchThrottleBurst is Maximum number of requests sent at the same moment if ThrottleQPS was not utilized before. Only used in batch mode. (default 15).
- `audit_webhook_batch_throttle_enable` - (Optional) - Bool - AuditWebhookBatchThrottleEnable is Whether batching throttling is enabled. Only used in batch mode. (default true).
- `audit_webhook_batch_throttle_qps` - (Optional) - Quantity - AuditWebhookBatchThrottleQps is Maximum average number of batches per second. Only used in batch mode. (default 10).
- `audit_webhook_config_file` - (Optional) - String - AuditWebhookConfigFile is Path to a kubeconfig formatted file that defines the audit webhook configuration. Requires the 'AdvancedAuditing' feature gate.
- `audit_webhook_initial_backoff` - (Optional) - Duration - AuditWebhookInitialBackoff is The amount of time to wait before retrying the first failed request. (default 10s).
- `audit_webhook_mode` - (Optional) - String - AuditWebhookMode is Strategy for sending audit events. Blocking indicates sending events should block server responses. Batch causes the backend to buffer and write events asynchronously. Known modes are batch,blocking. (default "batch").
- `authentication_token_webhook_config_file` - (Optional) - String - File with webhook configuration for token authentication in kubeconfig format. The API server will query the remote service to determine authentication for bearer tokens.
- `authentication_token_webhook_cache_ttl` - (Optional) - Duration - The duration to cache responses from the webhook token authenticator. Default is 2m. (default 2m0s).
- `authorization_mode` - (Optional) - String - AuthorizationMode is the authorization mode the kubeapi is running in.
- `authorization_webhook_config_file` - (Optional) - String - File with webhook configuration for authorization in kubeconfig format. The API server will query the remote service to determine whether to authorize the request.
- `authorization_webhook_cache_authorized_ttl` - (Optional) - Duration - The duration to cache authorized responses from the webhook token authorizer. Default is 5m. (default 5m0s).
- `authorization_webhook_cache_unauthorized_ttl` - (Optional) - Duration - The duration to cache authorized responses from the webhook token authorizer. Default is 30s. (default 30s).
- `authorization_rbac_super_user` - (Optional) - String - AuthorizationRBACSuperUser is the name of the superuser for default rbac.
- `encryption_provider_config` - (Optional) - String - EncryptionProviderConfig enables encryption at rest for secrets.
- `experimental_encryption_provider_config` - (Optional) - String - ExperimentalEncryptionProviderConfig enables encryption at rest for secrets.
- `requestheader_username_headers` - (Optional) - List(String) - List of request headers to inspect for usernames. X-Remote-User is common.
- `requestheader_group_headers` - (Optional) - List(String) - List of request headers to inspect for groups. X-Remote-Group is suggested.
- `requestheader_extra_header_prefixes` - (Optional) - List(String) - List of request header prefixes to inspect. X-Remote-Extra- is suggested.
- `requestheader_client_ca_file` - (Optional) - String - Root certificate bundle to use to verify client certificates on incoming requests before trusting usernames in headers specified by --requestheader-username-headers.
- `requestheader_allowed_names` - (Optional) - List(String) - List of client certificate common names to allow to provide usernames in headers specified by --requestheader-username-headers. If empty, any client certificate validated by the authorities in --requestheader-client-ca-file is allowed.
- `feature_gates` - (Optional) - Map(String) - FeatureGates is set of key=value pairs that describe feature gates for alpha/experimental features.
- `goaway_chance` - (Optional) - String - GoawayChance is the probability that send a GOAWAY to HTTP/2 clients. Default to 0, means never send GOAWAY. Max is 0.02 to prevent break the apiserver.
- `max_requests_inflight` - (Optional) - Int - MaxRequestsInflight The maximum number of non-mutating requests in flight at a given time.
- `max_mutating_requests_inflight` - (Optional) - Int - MaxMutatingRequestsInflight The maximum number of mutating requests in flight at a given time. Defaults to 200.
- `http2_max_streams_per_connection` - (Optional) - Int - HTTP2MaxStreamsPerConnection sets the limit that the server gives to clients for the maximum number of streams in an HTTP/2 connection. Zero means to use golang's default.
- `etcd_quorum_read` - (Optional) - Bool - EtcdQuorumRead configures the etcd-quorum-read flag, which forces consistent reads from etcd.
- `request_timeout` - (Optional) - Duration - RequestTimeout configures the duration a handler must keep a request open before timing it out. (default 1m0s).
- `min_request_timeout` - (Optional) - Int - MinRequestTimeout configures the minimum number of seconds a handler must keep a request open before timing it out.<br />Currently only honored by the watch request handler.
- `watch_cache` - (Optional) - Bool - Used to disable watch caching in the apiserver, defaults to enabling caching by omission.
- `watch_cache_sizes` - (Optional) - List(String) - Set the watch-cache-sizes parameter for the apiserver<br />The only meaningful value is setting to 0, which disable caches for specific object types.<br />Setting any values other than 0 for a resource will yield no effect since the caches are dynamic.
- `service_account_key_file` - (Optional) - List(String) - File containing PEM-encoded x509 RSA or ECDSA private or public keys, used to verify ServiceAccount tokens.<br />The specified file can contain multiple keys, and the flag can be specified multiple times with different files.<br />If unspecified, --tls-private-key-file is used.
- `service_account_signing_key_file` - (Optional) - String - Path to the file that contains the current private key of the service account token issuer.<br />The issuer will sign issued ID tokens with this private key. (Requires the 'TokenRequest' feature gate.).
- `service_account_issuer` - (Optional) - String - Identifier of the service account token issuer. The issuer will assert this identifier<br />in "iss" claim of issued tokens. This value is a string or URI.
- `additional_service_account_issuers` - (Optional) - List(String) - AdditionalServiceAccountIssuers can contain additional service account token issuers.
- `service_account_jwksuri` - (Optional) - String - ServiceAccountJWKSURI overrides the path for the jwks document; this is useful when we are republishing the service account discovery information elsewhere.
- `api_audiences` - (Optional) - List(String) - Identifiers of the API. The service account token authenticator will validate that<br />tokens used against the API are bound to at least one of these audiences. If the<br />--service-account-issuer flag is configured and this flag is not, this field<br />defaults to a single element list containing the issuer URL.
- `cpu_request` - (Optional) - Quantity - CPURequest, cpu request compute resource for api server. Defaults to "150m".
- `cpu_limit` - (Optional) - Quantity - CPULimit, cpu limit compute resource for api server e.g. "500m".
- `memory_request` - (Optional) - Quantity - MemoryRequest, memory request compute resource for api server e.g. "30Mi".
- `memory_limit` - (Optional) - Quantity - MemoryLimit, memory limit compute resource for api server e.g. "30Mi".
- `event_ttl` - (Optional) - Duration - Amount of time to retain Kubernetes events.
- `audit_dynamic_configuration` - (Optional) - Bool - AuditDynamicConfiguration enables dynamic audit configuration via AuditSinks.
- `enable_profiling` - (Optional) - Bool - EnableProfiling enables profiling via web interface host:port/debug/pprof/.
- `enable_contention_profiling` - (Optional) - Bool - EnableContentionProfiling enables block profiling, if profiling is enabled.
- `cors_allowed_origins` - (Optional) - List(String) - CorsAllowedOrigins is a list of origins for CORS. An allowed origin can be a regular<br />expression to support subdomain matching. If this list is empty CORS will not be enabled.
- `default_not_ready_toleration_seconds` - (Optional) - Int - DefaultNotReadyTolerationSeconds indicates the tolerationSeconds of the toleration for notReady:NoExecute that is added by default to every pod that does not already have such a toleration.
- `default_unreachable_toleration_seconds` - (Optional) - Int - DefaultUnreachableTolerationSeconds indicates the tolerationSeconds of the toleration for unreachable:NoExecute that is added by default to every pod that does not already have such a toleration.
- `env` - (Optional) - List([env_var](#env_var)) - Env allows users to pass in env variables to the apiserver container.<br />This can be useful to control some environment runtime settings, such as GOMEMLIMIT and GOCG to tweak the memory settings of the apiserver<br />This also allows the flexibility for adding any other variables for future use cases.

### env_var

EnvVar represents an environment variable present in a Container.

#### Argument Reference

The following arguments are supported:

- `name` - (Required) - String - Name of the environment variable. Must be a C_IDENTIFIER.
- `value` - (Optional) - String - Variable references $(VAR_NAME) are expanded<br />using the previously defined environment variables in the container and<br />any service environment variables. If a variable cannot be resolved,<br />the reference in the input string will be unchanged. Double $$ are reduced<br />to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.<br />"$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".<br />Escaped references will never be expanded, regardless of whether the variable<br />exists or not.<br />Defaults to "".<br />+optional.
- `value_from` - (Optional) - [env_var_source](#env_var_source) - Source for the environment variable's value. Cannot be used if value is not empty.<br />+optional.

### env_var_source

EnvVarSource represents a source for the value of an EnvVar.

#### Argument Reference

The following arguments are supported:

- `field_ref` - (Optional) - [object_field_selector](#object_field_selector) - Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`,<br />spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.<br />+optional.
- `resource_field_ref` - (Optional) - [resource_field_selector](#resource_field_selector) - Selects a resource of the container: only resources limits and requests<br />(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.<br />+optional.
- `config_map_key_ref` - (Optional) - [config_map_key_selector](#config_map_key_selector) - Selects a key of a ConfigMap.<br />+optional.
- `secret_key_ref` - (Optional) - [secret_key_selector](#secret_key_selector) - Selects a key of a secret in the pod's namespace<br />+optional.

### object_field_selector

ObjectFieldSelector selects an APIVersioned field of an object.<br />+structType=atomic.

#### Argument Reference

The following arguments are supported:

- `api_version` - (Optional) - String - Version of the schema the FieldPath is written in terms of, defaults to "v1".<br />+optional.
- `field_path` - (Required) - String - Path of the field to select in the specified API version.

### resource_field_selector

ResourceFieldSelector represents container resources (cpu, memory) and their output format<br />+structType=atomic.

#### Argument Reference

The following arguments are supported:

- `container_name` - (Optional) - String - Container name: required for volumes, optional for env vars<br />+optional.
- `resource` - (Required) - String - Required: resource to select.

### config_map_key_selector

Selects a key from a ConfigMap.<br />+structType=atomic.

#### Argument Reference

The following arguments are supported:

- `name` - (Optional) - String - Name of the referent.<br />This field is effectively required, but due to backwards compatibility is<br />allowed to be empty. Instances of this type with an empty value here are<br />almost certainly wrong.<br />More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br />+optional<br />+default=""<br />+kubebuilder:default=""<br />TODO: Drop `kubebuilder:default` when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.
- `key` - (Required) - String - The key to select.
- `optional` - (Optional) - Bool - Specify whether the ConfigMap or its key must be defined<br />+optional.

### secret_key_selector

SecretKeySelector selects a key of a Secret.<br />+structType=atomic.

#### Argument Reference

The following arguments are supported:

- `name` - (Optional) - String - Name of the referent.<br />This field is effectively required, but due to backwards compatibility is<br />allowed to be empty. Instances of this type with an empty value here are<br />almost certainly wrong.<br />More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br />+optional<br />+default=""<br />+kubebuilder:default=""<br />TODO: Drop `kubebuilder:default` when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.
- `key` - (Required) - String - The key of the secret to select from.  Must be a valid secret key.
- `optional` - (Optional) - Bool - Specify whether the Secret or its key must be defined<br />+optional.

### kube_controller_manager_config

KubeControllerManagerConfig is the configuration for the controller.

#### Argument Reference

The following arguments are supported:

- `master` - (Optional) - String - Master is the url for the kube api master.
- `log_format` - (Optional) - String - LogFormat is the logging format of the controler manager.<br />Supported values: text, json.<br />Default: text.
- `log_level` - (Optional) - Int - LogLevel is the defined logLevel.
- `service_account_private_key_file` - (Optional) - String - ServiceAccountPrivateKeyFile is the location of the private key for service account token signing.
- `image` - (Optional) - String - Image is the container image to use.
- `cloud_provider` - (Optional) - String - CloudProvider is the provider for cloud services.
- `cluster_name` - (Optional) - String - ClusterName is the instance prefix for the cluster.
- `cluster_cidr` - (Optional) - String - ClusterCIDR is CIDR Range for Pods in cluster.
- `allocate_node_cidrs` - (Optional) - Bool - AllocateNodeCIDRs enables CIDRs for Pods to be allocated and, if ConfigureCloudRoutes is true, to be set on the cloud provider.
- `node_cidr_mask_size` - (Optional) - Int - NodeCIDRMaskSize set the size for the mask of the nodes.
- `configure_cloud_routes` - (Optional) - Bool - ConfigureCloudRoutes enables CIDRs allocated with to be configured on the cloud provider.
- `controllers` - (Optional) - List(String) - Controllers is a list of controllers to enable on the controller-manager.
- `cidr_allocator_type` - (Optional) - String - CIDRAllocatorType specifies the type of CIDR allocator to use.
- `root_ca_file` - (Optional) - String - rootCAFile is the root certificate authority will be included in service account's token secret. This must be a valid PEM-encoded CA bundle.
- `leader_election` - (Optional) - [leader_election_configuration](#leader_election_configuration) - LeaderElection defines the configuration of leader election client.
- `attach_detach_reconcile_sync_period` - (Optional) - Duration - AttachDetachReconcileSyncPeriod is the amount of time the reconciler sync states loop<br />wait between successive executions. Is set to 1 min by kops by default.
- `disable_attach_detach_reconcile_sync` - (Optional) - Bool - DisableAttachDetachReconcileSync disables the reconcile sync loop in the attach-detach controller.<br />This can cause volumes to become mismatched with pods.
- `terminated_pod_gc_threshold` - (Optional) - Int - TerminatedPodGCThreshold is the number of terminated pods that can exist<br />before the terminated pod garbage collector starts deleting terminated pods.<br />If <= 0, the terminated pod garbage collector is disabled.
- `node_monitor_period` - (Optional) - Duration - NodeMonitorPeriod is the period for syncing NodeStatus in NodeController. (default 5s).
- `node_monitor_grace_period` - (Optional) - Duration - NodeMonitorGracePeriod is the amount of time which we allow running Node to be unresponsive before marking it unhealthy. (default 40s)<br />Must be N-1 times more than kubelet's nodeStatusUpdateFrequency, where N means number of retries allowed for kubelet to post node status.
- `pod_eviction_timeout` - (Optional) - Duration - PodEvictionTimeout is the grace period for deleting pods on failed nodes. (default 5m0s).
- `use_service_account_credentials` - (Optional) - Bool - UseServiceAccountCredentials controls whether we use individual service account credentials for each controller.
- `horizontal_pod_autoscaler_sync_period` - (Optional) - Duration - HorizontalPodAutoscalerSyncPeriod is the amount of time between syncs<br />During each period, the controller manager queries the resource utilization<br />against the metrics specified in each HorizontalPodAutoscaler definition.
- `horizontal_pod_autoscaler_downscale_delay` - (Optional) - Duration - HorizontalPodAutoscalerDownscaleDelay is a duration that specifies<br />how long the autoscaler has to wait before another downscale<br />operation can be performed after the current one has completed.
- `horizontal_pod_autoscaler_downscale_stabilization` - (Optional) - Duration - HorizontalPodAutoscalerDownscaleStabilization is the period for which<br />autoscaler will look backwards and not scale down below any<br />recommendation it made during that period.
- `horizontal_pod_autoscaler_upscale_delay` - (Optional) - Duration - HorizontalPodAutoscalerUpscaleDelay is a duration that specifies how<br />long the autoscaler has to wait before another upscale operation can<br />be performed after the current one has completed.
- `horizontal_pod_autoscaler_initial_readiness_delay` - (Optional) - Duration - HorizontalPodAutoscalerInitialReadinessDelay is the period after pod start<br />during which readiness changes will be treated as initial readiness. (default 30s).
- `horizontal_pod_autoscaler_cpu_initialization_period` - (Optional) - Duration - HorizontalPodAutoscalerCPUInitializationPeriod is the period after pod start<br />when CPU samples might be skipped. (default 5m).
- `horizontal_pod_autoscaler_tolerance` - (Optional) - Quantity - HorizontalPodAutoscalerTolerance is the minimum change (from 1.0) in the<br />desired-to-actual metrics ratio for the horizontal pod autoscaler to<br />consider scaling.
- `horizontal_pod_autoscaler_use_rest_clients` - (Optional) - Bool - HorizontalPodAutoscalerUseRestClients determines if the new-style clients<br />should be used if support for custom metrics is enabled.
- `experimental_cluster_signing_duration` - (Optional) - Duration - ExperimentalClusterSigningDuration is the max length of duration that the signed certificates will be given. (default 365*24h)<br />Deprecated - use cluster-signing-duration instead.
- `cluster_signing_duration` - (Optional) - Duration - ClusterSigningDuration is the max length of duration that the signed certificates will be given. (default 365*24h).
- `feature_gates` - (Optional) - Map(String) - FeatureGates is set of key=value pairs that describe feature gates for alpha/experimental features.
- `tls_cert_file` - (Optional) - String - TLSCertFile is the file containing the TLS server certificate.
- `tls_cipher_suites` - (Optional) - List(String) - TLSCipherSuites indicates the allowed TLS cipher suite.
- `tls_min_version` - (Optional) - String - TLSMinVersion indicates the minimum TLS version allowed.
- `tls_private_key_file` - (Optional) - String - TLSPrivateKeyFile is the file containing the private key for the TLS server certificate.
- `min_resync_period` - (Optional) - String - MinResyncPeriod indicates the resync period in reflectors.<br />The resync period will be random between MinResyncPeriod and 2*MinResyncPeriod. (default 12h0m0s).
- `kube_api_qps` - (Optional) - Quantity - KubeAPIQPS QPS to use while talking with kubernetes apiserver. (default 20).
- `kube_api_burst` - (Optional) - Int - KubeAPIBurst Burst to use while talking with kubernetes apiserver. (default 30).
- `concurrent_deployment_syncs` - (Optional) - Int - The number of deployment objects that are allowed to sync concurrently.
- `concurrent_endpoint_syncs` - (Optional) - Int - The number of endpoint objects that are allowed to sync concurrently.
- `concurrent_namespace_syncs` - (Optional) - Int - The number of namespace objects that are allowed to sync concurrently.
- `concurrent_replicaset_syncs` - (Optional) - Int - The number of replicaset objects that are allowed to sync concurrently.
- `concurrent_service_syncs` - (Optional) - Int - The number of service objects that are allowed to sync concurrently.
- `concurrent_resource_quota_syncs` - (Optional) - Int - The number of resourcequota objects that are allowed to sync concurrently.
- `concurrent_serviceaccount_token_syncs` - (Optional) - Int - The number of serviceaccount objects that are allowed to sync concurrently to create tokens.
- `concurrent_rc_syncs` - (Optional) - Int - The number of replicationcontroller objects that are allowed to sync concurrently.
- `concurrent_horizontal_pod_austoscaler_syncs` - (Optional) - Int - The number of horizontal pod autoscaler objects that are allowed to sync concurrently (default 5).
- `concurrent_job_syncs` - (Optional) - Int - The number of job objects that are allowed to sync concurrently (default 5).
- `authentication_kubeconfig` - (Optional) - String - AuthenticationKubeconfig is the path to an Authentication Kubeconfig.
- `authorization_kubeconfig` - (Optional) - String - AuthorizationKubeconfig is the path to an Authorization Kubeconfig.
- `authorization_always_allow_paths` - (Optional) - List(String) - AuthorizationAlwaysAllowPaths is the list of HTTP paths to skip during authorization.
- `external_cloud_volume_plugin` - (Optional) - String - ExternalCloudVolumePlugin is a fallback mechanism that allows a legacy, in-tree cloudprovider to be used for volume plugins<br />even when an external cloud controller manager is being used.  This can be used instead of installing CSI.  The value should<br />be the same as is used for the --cloud-provider flag, i.e. "aws".
- `endpoint_updates_batch_period` - (Optional) - Duration - The length of endpoint updates batching period. Processing of pod changes will be delayed by this duration<br />to join them with potential upcoming updates and reduce the overall number of endpoints updates.<br />Larger number = higher endpoint programming latency, but lower number of endpoints revision generated.
- `endpoint_slice_updates_batch_period` - (Optional) - Duration - The length of endpoint slice updates batching period. Processing of pod changes will be delayed by this duration<br />to join them with potential upcoming updates and reduce the overall number of endpoints updates.<br />Larger number = higher endpoint programming latency, but lower number of endpoints revision generated.
- `enable_profiling` - (Optional) - Bool - EnableProfiling enables profiling via web interface host:port/debug/pprof/.
- `enable_contention_profiling` - (Optional) - Bool - EnableContentionProfiling enables block profiling, if profiling is enabled.
- `enable_leader_migration` - (Optional) - Bool - EnableLeaderMigration enables controller leader migration.
- `cpu_request` - (Optional) - Quantity - CPURequest, cpu request compute resource for kube-controler-manager. Defaults to "100m".
- `cpu_limit` - (Optional) - Quantity - CPULimit, cpu limit compute resource for kube-controler-manager e.g. "500m".
- `memory_request` - (Optional) - Quantity - MemoryRequest, memory request compute resource for kube-controler-manager e.g. "30Mi".
- `memory_limit` - (Optional) - Quantity - MemoryLimit, memory limit compute resource for kube-controler-manager e.g. "30Mi".

### leader_election_configuration

LeaderElectionConfiguration defines the configuration of leader election<br />clients for components that can run with leader election enabled.

#### Argument Reference

The following arguments are supported:

- `leader_elect` - (Optional) - Bool - leaderElect enables a leader election client to gain leadership<br />before executing the main loop. Enable this when running replicated<br />components for high availability.
- `leader_elect_lease_duration` - (Optional) - Duration - leaderElectLeaseDuration is the length in time non-leader candidates<br />will wait after observing a leadership renewal until attempting to acquire<br />leadership of a led but unrenewed leader slot. This is effectively the<br />maximum duration that a leader can be stopped before it is replaced by another candidate.
- `leader_elect_renew_deadline_duration` - (Optional) - Duration - LeaderElectRenewDeadlineDuration is the interval between attempts by the acting master to<br />renew a leadership slot before it stops leading. This must be less than or equal to the lease duration.
- `leader_elect_resource_lock` - (Optional) - String - LeaderElectResourceLock is the type of resource object that is used for locking during<br />leader election. Supported options are endpoints (default) and `configmaps`.
- `leader_elect_resource_name` - (Optional) - String - LeaderElectResourceName is the name of resource object that is used for locking during leader election.
- `leader_elect_resource_namespace` - (Optional) - String - LeaderElectResourceNamespace is the namespace of resource object that is used for locking during leader election.
- `leader_elect_retry_period` - (Optional) - Duration - LeaderElectRetryPeriod is The duration the clients should wait between attempting acquisition<br />and renewal of a leadership. This is only applicable if leader election is enabled.

### cloud_controller_manager_config

CloudControllerManagerConfig is the configuration of the cloud controller.

#### Argument Reference

The following arguments are supported:

- `master` - (Optional) - String - Master is the url for the kube api master.
- `log_level` - (Optional) - Int - LogLevel is the verbosity of the logs.
- `image` - (Optional) - String - Image is the OCI image of the cloud controller manager.
- `cloud_provider` - (Optional) - String - CloudProvider is the provider for cloud services.
- `cluster_name` - (Optional) - String - ClusterName is the instance prefix for the cluster.
- `allow_untagged_cloud` - (Optional) - Bool - Allow the cluster to run without the cluster-id on cloud instances.
- `cluster_cidr` - (Optional) - String - ClusterCIDR is CIDR Range for Pods in cluster.
- `allocate_node_cidrs` - (Optional) - Bool - AllocateNodeCIDRs enables CIDRs for Pods to be allocated and, if<br />ConfigureCloudRoutes is true, to be set on the cloud provider.
- `configure_cloud_routes` - (Optional) - Bool - ConfigureCloudRoutes enables CIDRs allocated with to be configured on the cloud provider.
- `controllers` - (Optional) - List(String) - Controllers is a list of controllers to enable on the controller-manager.
- `cidr_allocator_type` - (Optional) - String - CIDRAllocatorType specifies the type of CIDR allocator to use.
- `leader_election` - (Optional) - [leader_election_configuration](#leader_election_configuration) - LeaderElection defines the configuration of leader election client.
- `use_service_account_credentials` - (Optional) - Bool - UseServiceAccountCredentials controls whether we use individual service account credentials for each controller.
- `enable_leader_migration` - (Optional) - Bool - EnableLeaderMigration enables controller leader migration.
- `cpu_request` - (Optional) - Quantity - CPURequest of NodeTerminationHandler container.<br />Default: 200m.
- `node_status_update_frequency` - (Optional) - Duration - NodeStatusUpdateFrequency is the duration between node status updates. (default: 5m).
- `concurrent_node_syncs` - (Optional) - Int - ConcurrentNodeSyncs is the number of workers concurrently synchronizing nodes. (default: 1).

### kube_scheduler_config

KubeSchedulerConfig is the configuration for the kube-scheduler.

#### Argument Reference

The following arguments are supported:

- `master` - (Optional) - String - Master is a url to the kube master.
- `log_format` - (Optional) - String - LogFormat is the logging format of the scheduler.<br />Supported values: text, json.<br />Default: text.
- `log_level` - (Optional) - Int - LogLevel is the logging level.
- `image` - (Optional) - String - Image is the container image to use.
- `leader_election` - (Optional) - [leader_election_configuration](#leader_election_configuration) - LeaderElection defines the configuration of leader election client.
- `use_policy_config_map` - (Optional) - Bool - UsePolicyConfigMap enable setting the scheduler policy from a configmap<br />Deprecated - use KubeSchedulerConfiguration instead.
- `feature_gates` - (Optional) - Map(String) - FeatureGates is set of key=value pairs that describe feature gates for alpha/experimental features.
- `max_persistent_volumes` - (Optional) - Int - MaxPersistentVolumes changes the maximum number of persistent volumes the scheduler will scheduler onto the same<br />node. Only takes effect if value is positive. This corresponds to the KUBE_MAX_PD_VOLS environment variable.<br />The default depends on the version and the cloud provider<br />as outlined: https://kubernetes.io/docs/concepts/storage/storage-limits/.
- `qps` - (Optional) - Quantity - Qps sets the maximum qps to send to apiserver after the burst quota is exhausted.
- `burst` - (Optional) - Int - Burst sets the maximum qps to send to apiserver after the burst quota is exhausted.
- `kube_api_qps` - (Optional) - Quantity - KubeAPIQPS QPS to use while talking with kubernetes apiserver. (default 20).
- `kube_api_burst` - (Optional) - Int - KubeAPIBurst Burst to use while talking with kubernetes apiserver. (default 30).
- `authentication_kubeconfig` - (Optional) - String - AuthenticationKubeconfig is the path to an Authentication Kubeconfig.
- `authorization_kubeconfig` - (Optional) - String - AuthorizationKubeconfig is the path to an Authorization Kubeconfig.
- `authorization_always_allow_paths` - (Optional) - List(String) - AuthorizationAlwaysAllowPaths is the list of HTTP paths to skip during authorization.
- `enable_profiling` - (Optional) - Bool - EnableProfiling enables profiling via web interface host:port/debug/pprof/.
- `enable_contention_profiling` - (Optional) - Bool - EnableContentionProfiling enables block profiling, if profiling is enabled.
- `tls_cert_file` - (Optional) - String - TLSCertFile is the file containing the TLS server certificate.
- `tls_private_key_file` - (Optional) - String - TLSPrivateKeyFile is the file containing the private key for the TLS server certificate.
- `cpu_request` - (Optional) - Quantity - CPURequest, cpu request compute resource for scheduler. Defaults to "100m".
- `cpu_limit` - (Optional) - Quantity - CPULimit, cpu limit compute resource for scheduler e.g. "500m".
- `memory_request` - (Optional) - Quantity - MemoryRequest, memory request compute resource for scheduler e.g. "30Mi".
- `memory_limit` - (Optional) - Quantity - MemoryLimit, memory limit compute resource for scheduler e.g. "30Mi".

### kube_proxy_config

KubeProxyConfig defines the configuration for a proxy.

#### Argument Reference

The following arguments are supported:

- `image` - (Optional) - String
- `cpu_request` - (Optional) - Quantity - CPURequest, cpu request compute resource for kube proxy e.g. "20m".
- `cpu_limit` - (Optional) - Quantity - CPULimit, cpu limit compute resource for kube proxy e.g. "30m".
- `memory_request` - (Optional) - Quantity - MemoryRequest, memory request compute resource for kube proxy e.g. "30Mi".
- `memory_limit` - (Optional) - Quantity - MemoryLimit, memory limit compute resource for kube proxy e.g. "30Mi".
- `log_level` - (Optional) - Int - LogLevel is the logging level of the proxy.
- `cluster_cidr` - (Optional) - String - ClusterCIDR is the CIDR range of the pods in the cluster.
- `hostname_override` - (Optional) - String - HostnameOverride, if non-empty, will be used as the identity instead of the actual hostname.
- `bind_address` - (Optional) - String - BindAddress is IP address for the proxy server to serve on.
- `master` - (Optional) - String - Master is the address of the Kubernetes API server (overrides any value in kubeconfig).
- `metrics_bind_address` - (Optional) - String - MetricsBindAddress is the IP address for the metrics server to serve on.
- `enabled` - (Required) - Bool - Enabled allows enabling or disabling kube-proxy.
- `proxy_mode` - (Optional) - String - Which proxy mode to use: (userspace, iptables(default), ipvs).
- `ip_vs_exclude_cidrs` - (Optional) - List(String) - IPVSExcludeCIDRs is comma-separated list of CIDR's which the ipvs proxier should not touch when cleaning up IPVS rules.
- `ip_vs_min_sync_period` - (Optional) - Duration - IPVSMinSyncPeriod is the minimum interval of how often the ipvs rules can be refreshed as endpoints and services change (e.g. '5s', '1m', '2h22m').
- `ip_vs_scheduler` - (Optional) - String - IPVSScheduler is the ipvs scheduler type when proxy mode is ipvs.
- `ip_vs_sync_period` - (Optional) - Duration - IPVSSyncPeriod duration is the maximum interval of how often ipvs rules are refreshed.
- `feature_gates` - (Optional) - Map(String) - FeatureGates is a series of key pairs used to switch on features for the proxy.
- `conntrack_max_per_core` - (Optional) - Int - Maximum number of NAT connections to track per CPU core (default: 131072).
- `conntrack_min` - (Optional) - Int - Minimum number of conntrack entries to allocate, regardless of conntrack-max-per-core.

### kubelet_config_spec

KubeletConfigSpec defines the kubelet configuration.

#### Argument Reference

The following arguments are supported:

- `api_servers` - (Optional) - String - APIServers is not used for clusters version 1.6 and later - flag removed.
- `anonymous_auth` - (Optional) - Bool([Nullable](#nullable-arguments)) - AnonymousAuth permits you to control auth to the kubelet api.
- `authorization_mode` - (Optional) - String - AuthorizationMode is the authorization mode the kubelet is running in.
- `bootstrap_kubeconfig` - (Optional) - String - BootstrapKubeconfig is the path to a kubeconfig file that will be used to get client certificate for kubelet.
- `client_ca_file` - (Optional) - String - ClientCAFile is the path to a CA certificate.
- `tls_cert_file` - (Optional) - String - TODO: Remove unused TLSCertFile.
- `tls_private_key_file` - (Optional) - String - TODO: Remove unused TLSPrivateKeyFile.
- `tls_cipher_suites` - (Optional) - List(String) - TLSCipherSuites indicates the allowed TLS cipher suite.
- `tls_min_version` - (Optional) - String - TLSMinVersion indicates the minimum TLS version allowed.
- `kubeconfig_path` - (Optional) - String - KubeconfigPath is the path of kubeconfig for the kubelet.
- `require_kubeconfig` - (Optional) - Bool - RequireKubeconfig indicates a kubeconfig is required.
- `log_format` - (Optional) - String - LogFormat is the logging format of the kubelet.<br />Supported values: text, json.<br />Default: text.
- `log_level` - (Optional) - Int - LogLevel is the logging level of the kubelet.
- `pod_manifest_path` - (Optional) - String - config is the path to the config file or directory of files.
- `hostname_override` - (Optional) - String - HostnameOverride is the hostname used to identify the kubelet instead of the actual hostname.
- `pod_infra_container_image` - (Optional) - String - PodInfraContainerImage is the image whose network/ipc containers in each pod will use.
- `seccomp_default` - (Optional) - Bool - SeccompDefault enables the use of `RuntimeDefault` as the default seccomp profile for all workloads.
- `seccomp_profile_root` - (Optional) - String - SeccompProfileRoot is the directory path for seccomp profiles.
- `allow_privileged` - (Optional) - Bool - AllowPrivileged enables containers to request privileged mode (defaults to false).
- `enable_debugging_handlers` - (Optional) - Bool - EnableDebuggingHandlers enables server endpoints for log collection and local running of containers and commands.
- `register_node` - (Optional) - Bool - RegisterNode enables automatic registration with the apiserver.
- `node_status_update_frequency` - (Optional) - Duration - NodeStatusUpdateFrequency Specifies how often kubelet posts node status to master (default 10s)<br />must work with nodeMonitorGracePeriod in KubeControllerManagerConfig.
- `cluster_domain` - (Optional) - String - ClusterDomain is the DNS domain for this cluster.
- `cluster_dns` - (Optional) - String - ClusterDNS is the IP address for a cluster DNS server.
- `network_plugin_name` - (Optional) - String - NetworkPluginName is the name of the network plugin to be invoked for various events in kubelet/pod lifecycle.
- `cloud_provider` - (Optional) - String - CloudProvider is the provider for cloud services.
- `kubelet_cgroups` - (Optional) - String - KubeletCgroups is the absolute name of cgroups to isolate the kubelet in.
- `runtime_cgroups` - (Optional) - String - Cgroups that container runtime is expected to be isolated in.
- `read_only_port` - (Optional) - Int - ReadOnlyPort is the port used by the kubelet api for read-only access (default 10255).
- `system_cgroups` - (Optional) - String - SystemCgroups is absolute name of cgroups in which to place<br />all non-kernel processes that are not already in a container. Empty<br />for no container. Rolling back the flag requires a reboot.
- `cgroup_root` - (Optional) - String - cgroupRoot is the root cgroup to use for pods. This is handled by the container runtime on a best effort basis.
- `configure_cbr0` - (Optional) - Bool - configureCBR0 enables the kubelet to configure cbr0 based on Node.Spec.PodCIDR.
- `hairpin_mode` - (Optional) - String - How should the kubelet configure the container bridge for hairpin packets.<br />Setting this flag allows endpoints in a Service to loadbalance back to<br />themselves if they should try to access their own Service. Values:<br />  "promiscuous-bridge": make the container bridge promiscuous.<br />  "hairpin-veth":       set the hairpin flag on container veth interfaces.<br />  "none":               do nothing.<br />Setting --configure-cbr0 to false implies that to achieve hairpin NAT<br />one must set --hairpin-mode=veth-flag, because bridge assumes the<br />existence of a container bridge named cbr0.
- `babysit_daemons` - (Optional) - Bool - The node has babysitter process monitoring docker and kubelet. Removed as of 1.7.
- `max_pods` - (Optional) - Int - MaxPods is the number of pods that can run on this Kubelet.
- `nvidia_gp_us` - (Optional) - Int - NvidiaGPUs is the number of NVIDIA GPU devices on this node.
- `pod_cidr` - (Optional) - String - PodCIDR is the CIDR to use for pod IP addresses, only used in standalone mode.<br />In cluster mode, this is obtained from the master.
- `resolver_config` - (Optional) - String - ResolverConfig is the resolver configuration file used as the basis for the container DNS resolution configuration."), [].
- `reconcile_cidr` - (Optional) - Bool - ReconcileCIDR is Reconcile node CIDR with the CIDR specified by the<br />API server. No-op if register-node or configure-cbr0 is false.
- `register_schedulable` - (Optional) - Bool - registerSchedulable tells the kubelet to register the node as schedulable. No-op if register-node is false.
- `serialize_image_pulls` - (Optional) - Bool - SerializeImagePulls when enabled, tells the Kubelet to pull images one at a time.
- `node_labels` - (Optional) - Map(String) - NodeLabels to add when registering the node in the cluster.
- `non_masquerade_cidr` - (Optional) - String - NonMasqueradeCIDR configures masquerading: traffic to IPs outside this range will use IP masquerade.
- `enable_custom_metrics` - (Optional) - Bool - Enable gathering custom metrics.
- `network_plugin_mtu` - (Optional) - Int - NetworkPluginMTU is the MTU to be passed to the network plugin,<br />and overrides the default MTU for cases where it cannot be automatically<br />computed (such as IPSEC).
- `image_minimum_gc_age` - (Optional) - String - imageMinimumGCAge is the minimum age for an unused image before it is garbage collected. Default: "2m".
- `image_maximum_gc_age` - (Optional) - String - imageMaximumGCAge is the maximum age an image can be unused before it is garbage collected.<br />The default of this field is "0s", which disables this field--meaning images won't be garbage<br />collected based on being unused for too long. Default: "0s" (disabled).
- `image_gc_high_threshold_percent` - (Optional) - Int - ImageGCHighThresholdPercent is the percent of disk usage after which<br />image garbage collection is always run.
- `image_gc_low_threshold_percent` - (Optional) - Int - ImageGCLowThresholdPercent is the percent of disk usage before which<br />image garbage collection is never run. Lowest disk usage to garbage<br />collect to.
- `image_pull_progress_deadline` - (Optional) - Duration - ImagePullProgressDeadline is the timeout for image pulls<br />If no pulling progress is made before this deadline, the image pulling will be cancelled. (default 1m0s).
- `eviction_hard` - (Optional) - String - Comma-delimited list of hard eviction expressions.  For example, 'memory.available<300Mi'.
- `eviction_soft` - (Optional) - String - Comma-delimited list of soft eviction expressions.  For example, 'memory.available<300Mi'.
- `eviction_soft_grace_period` - (Optional) - String - Comma-delimited list of grace periods for each soft eviction signal.  For example, 'memory.available=30s'.
- `eviction_pressure_transition_period` - (Optional) - Duration - Duration for which the kubelet has to wait before transitioning out of an eviction pressure condition.
- `eviction_max_pod_grace_period` - (Optional) - Int - Maximum allowed grace period (in seconds) to use when terminating pods in response to a soft eviction threshold being met.
- `eviction_minimum_reclaim` - (Optional) - String - Comma-delimited list of minimum reclaims (e.g. imagefs.available=2Gi) that describes the minimum amount of resource the kubelet will reclaim when performing a pod eviction if that resource is under pressure.
- `volume_plugin_directory` - (Optional) - String - The full path of the directory in which to search for additional third party volume plugins (this path must be writeable, dependent on your choice of OS).
- `taints` - (Optional) - List(String) - Taints to add when registering a node in the cluster.
- `feature_gates` - (Optional) - Map(String) - FeatureGates is set of key=value pairs that describe feature gates for alpha/experimental features.
- `kernel_memcg_notification` - (Optional) - Bool - Integrate with the kernel memcg notification to determine if memory eviction thresholds are crossed rather than polling.
- `kube_reserved` - (Optional) - Map(String) - Resource reservation for kubernetes system daemons like the kubelet, container runtime, node problem detector, etc.
- `kube_reserved_cgroup` - (Optional) - String - Control group for kube daemons.
- `system_reserved` - (Optional) - Map(String) - Capture resource reservation for OS system daemons like sshd, udev, etc.
- `system_reserved_cgroup` - (Optional) - String - Parent control group for OS system daemons.
- `enforce_node_allocatable` - (Optional) - String - Enforce Allocatable across pods whenever the overall usage across all pods exceeds Allocatable.
- `runtime_request_timeout` - (Optional) - Duration - RuntimeRequestTimeout is timeout for runtime requests on - pull, logs, exec and attach.
- `volume_stats_agg_period` - (Optional) - Duration - VolumeStatsAggPeriod is the interval for kubelet to calculate and cache the volume disk usage for all pods and volumes.
- `fail_swap_on` - (Optional) - Bool - Tells the Kubelet to fail to start if swap is enabled on the node.
- `experimental_allowed_unsafe_sysctls` - (Optional) - List(String) - ExperimentalAllowedUnsafeSysctls are passed to the kubelet config to whitelist allowable sysctls<br />Was promoted to beta and renamed. https://github.com/kubernetes/kubernetes/pull/63717.
- `allowed_unsafe_sysctls` - (Optional) - List(String) - AllowedUnsafeSysctls are passed to the kubelet config to whitelist allowable sysctls.
- `streaming_connection_idle_timeout` - (Optional) - Duration - StreamingConnectionIdleTimeout is the maximum time a streaming connection can be idle before the connection is automatically closed.
- `docker_disable_shared_pid` - (Optional) - Bool - DockerDisableSharedPID was removed.
- `root_dir` - (Optional) - String - RootDir is the directory path for managing kubelet files (volume mounts,etc).
- `authentication_token_webhook` - (Optional) - Bool - AuthenticationTokenWebhook uses the TokenReview API to determine authentication for bearer tokens.
- `authentication_token_webhook_cache_ttl` - (Optional) - Duration - AuthenticationTokenWebhook sets the duration to cache responses from the webhook token authenticator. Default is 2m. (default 2m0s).
- `cpu_cfs_quota` - (Optional) - Bool([Nullable](#nullable-arguments)) - CPUCFSQuota enables CPU CFS quota enforcement for containers that specify CPU limits.
- `cpu_cfs_quota_period` - (Optional) - Duration - CPUCFSQuotaPeriod sets CPU CFS quota period value, cpu.cfs_period_us, defaults to Linux Kernel default.
- `cpu_manager_policy` - (Optional) - String - CpuManagerPolicy allows for changing the default policy of None to static.
- `registry_pull_qps` - (Optional) - Int - RegistryPullQPS if > 0, limit registry pull QPS to this value.  If 0, unlimited. (default 5).
- `registry_burst` - (Optional) - Int - RegistryBurst Maximum size of a bursty pulls, temporarily allows pulls to burst to this number, while still not exceeding registry-qps. Only used if --registry-qps > 0 (default 10).
- `topology_manager_policy` - (Optional) - String - TopologyManagerPolicy determines the allocation policy for the topology manager.
- `rotate_certificates` - (Optional) - Bool - rotateCertificates enables client certificate rotation.
- `protect_kernel_defaults` - (Optional) - Bool - Default kubelet behaviour for kernel tuning. If set, kubelet errors if any of kernel tunables is different than kubelet defaults.<br />(DEPRECATED: This parameter should be set via the config file specified by the Kubelet's --config flag.
- `cgroup_driver` - (Optional) - String - CgroupDriver allows the explicit setting of the kubelet cgroup driver. If omitted, defaults to cgroupfs.
- `housekeeping_interval` - (Optional) - Duration - HousekeepingInterval allows to specify interval between container housekeepings.
- `event_qps` - (Optional) - Int - EventQPS if > 0, limit event creations per second to this value.  If 0, unlimited.
- `event_burst` - (Optional) - Int - EventBurst temporarily allows event records to burst to this number, while still not exceeding EventQPS. Only used if EventQPS > 0.
- `container_log_max_size` - (Optional) - String - ContainerLogMaxSize is the maximum size (e.g. 10Mi) of container log file before it is rotated.
- `container_log_max_files` - (Optional) - Int - ContainerLogMaxFiles is the maximum number of container log files that can be present for a container. The number must be >= 2.
- `enable_cadvisor_json_endpoints` - (Optional) - Bool - EnableCadvisorJsonEndpoints enables cAdvisor json `/spec` and `/stats/*` endpoints. Defaults to False.
- `pod_pids_limit` - (Optional) - Int - PodPidsLimit is the maximum number of pids in any pod.
- `experimental_allocatable_ignore_eviction` - (Optional) - Bool - ExperimentalAllocatableIgnoreEviction enables ignoring Hard Eviction Thresholds while calculating Node Allocatable.
- `shutdown_grace_period` - (Optional) - Duration - ShutdownGracePeriod specifies the total duration that the node should delay the shutdown by.<br />Default: 30s.
- `shutdown_grace_period_critical_pods` - (Optional) - Duration - ShutdownGracePeriodCriticalPods specifies the duration used to terminate critical pods during a node shutdown.<br />Default: 10s.
- `memory_swap_behavior` - (Optional) - String - MemorySwapBehavior defines how swap is used by container workloads.<br />Supported values: LimitedSwap, "UnlimitedSwap.

### cloud_configuration

CloudConfiguration defines the cloud provider configuration.

#### Argument Reference

The following arguments are supported:

- `manage_storage_classes` - (Optional) - Bool([Nullable](#nullable-arguments)) - ManageStorageClasses specifies whether kOps should create and maintain a set of<br />StorageClasses, one of which it nominates as the default class for the cluster.

### external_dns_config

ExternalDNSConfig are options of the dns-controller.

#### Argument Reference

The following arguments are supported:

- `watch_ingress` - (Optional) - Bool - WatchIngress indicates you want the dns-controller to watch and create dns entries for ingress resources.<br />Default: true if provider is 'external-dns', false otherwise.
- `watch_namespace` - (Optional) - String - WatchNamespace is namespace to watch, defaults to all (use to control whom can creates dns entries).
- `provider` - (Optional) - String - Provider determines which implementation of ExternalDNS to use.<br />'dns-controller' will use kOps DNS Controller.<br />'external-dns' will use kubernetes-sigs/external-dns.

### ntp_config

NTPConfig is the configuration for NTP.

#### Argument Reference

The following arguments are supported:

- `managed` - (Optional) - Bool - Managed controls if the NTP configuration is managed by kOps.<br />The NTP configuration task is skipped if this is set to false.

### node_problem_detector_config

NodeProblemDetector determines the node problem detector configuration.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enabled enables the NodeProblemDetector.<br />Default: false.
- `image` - (Optional) - String - Image is the NodeProblemDetector container image used.
- `memory_request` - (Optional) - Quantity - MemoryRequest of NodeProblemDetector container.<br />Default: 80Mi.
- `cpu_request` - (Optional) - Quantity - CPURequest of NodeProblemDetector container.<br />Default: 10m.
- `memory_limit` - (Optional) - Quantity - MemoryLimit of NodeProblemDetector container.<br />Default: 80Mi.
- `cpu_limit` - (Optional) - Quantity - CPULimit of NodeProblemDetector container.<br />Default: 10m.

### metrics_server_config

MetricsServerConfig determines the metrics server configuration.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enabled enables the metrics server.<br />Default: false.
- `image` - (Optional) - String - Image is the container image used.<br />Default: the latest supported image for the specified kubernetes version.
- `insecure` - (Required) - Bool - Insecure determines if API server will validate metrics server TLS cert.<br />Default: true.

### cert_manager_config

CertManagerConfig determines the cert manager configuration.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Required) - Bool - Enabled enables the cert manager.<br />Default: false.
- `managed` - (Required) - Bool - Managed controls if cert-manager is manged and deployed by kOps.<br />The deployment of cert-manager is skipped if this is set to false.
- `image` - (Optional) - String - Image is the container image used.<br />Default: the latest supported image for the specified kubernetes version.
- `default_issuer` - (Optional) - String - defaultIssuer sets a default clusterIssuer<br />Default: none.
- `nameservers` - (Optional) - List(String) - nameservers is a list of nameserver IP addresses to use instead of the pod defaults.<br />Default: none.
- `hosted_zone_ids` - (Optional) - List(String) - HostedZoneIDs is a list of route53 hostedzone IDs that cert-manager will be allowed to do dns-01 validation for.
- `feature_gates` - (Optional) - Map(Bool) - FeatureGates is a list of experimental features that can be enabled or disabled.

### networking_spec

NetworkingSpec configures networking.

#### Argument Reference

The following arguments are supported:

- `network_id` - (Optional) - String - NetworkID is the cloud provider's identifier of the existing network (for example, AWS VPC) the cluster should use.<br />If not specified, kOps will create a new network.
- `network_cidr` - (Optional) - (Computed) - String - NetworkCIDR is the primary IPv4 CIDR used for the cloud provider's network.<br />It is not required on GCE.<br />On DO, it maps to the VPC CIDR.
- `additional_network_cidrs` - (Optional) - List(String) - AdditionalNetworkCIDRs is a list of additional CIDR used for the AWS VPC<br />or otherwise allocated to k8s. This is a real CIDR, not the internal k8s network<br />On AWS, it maps to any additional CIDRs added to a VPC.
- `subnet` - (Required) - List([cluster_subnet_spec](#cluster_subnet_spec)) - Subnets are the subnets that the cluster can use.
- `tag_subnets` - (Optional) - Bool([Nullable](#nullable-arguments)) - TagSubnets controls if tags are added to subnets to enable use by load balancers (AWS only). Default: true.
- `topology` - (Required) - [topology_spec](#topology_spec) - Topology defines the type of network topology to use on the cluster - default public<br />This is heavily weighted towards AWS for the time being, but should also be agnostic enough<br />to port out to GCE later if needed.
- `egress_proxy` - (Optional) - [egress_proxy_spec](#egress_proxy_spec) - HTTPProxy defines connection information to support use of a private cluster behind an forward HTTP Proxy.
- `non_masquerade_cidr` - (Optional) - (Computed) - String - NonMasqueradeCIDR is the CIDR for the internal k8s network (for pod IPs)<br />It cannot overlap ServiceClusterIPRange.
- `pod_cidr` - (Optional) - String - PodCIDR is the CIDR from which we allocate IPs for pods.
- `service_cluster_ip_range` - (Optional) - String - ServiceClusterIPRange is the CIDR, from the internal network, where we allocate IPs for services.
- `isolate_control_plane` - (Optional) - Bool - IsolateControlPlane determines whether we should lock down masters so that they are not on the pod network.<br />true is the kube-up behaviour, but it is very surprising: it means that daemonsets only work on the master<br />if they have hostNetwork=true.<br />false is now the default, and it will:<br /> * give the master a normal PodCIDR<br /> * run kube-proxy on the master<br /> * enable debugging handlers on the master, so kubectl logs works.
- `classic` - (Optional) - [classic_networking_spec](#classic_networking_spec)
- `kubenet` - (Optional) - [kubenet_networking_spec](#kubenet_networking_spec)
- `external` - (Optional) - [external_networking_spec](#external_networking_spec)
- `cni` - (Optional) - [cni_networking_spec](#cni_networking_spec)
- `kopeio` - (Optional) - [kopeio_networking_spec](#kopeio_networking_spec)
- `weave` - (Optional) - [weave_networking_spec](#weave_networking_spec)
- `flannel` - (Optional) - [flannel_networking_spec](#flannel_networking_spec)
- `calico` - (Optional) - [calico_networking_spec](#calico_networking_spec)
- `canal` - (Optional) - [canal_networking_spec](#canal_networking_spec)
- `kube_router` - (Optional) - [kuberouter_networking_spec](#kuberouter_networking_spec)
- `romana` - (Optional) - [romana_networking_spec](#romana_networking_spec)
- `amazon_vpc` - (Optional) - [amazon_vpc_networking_spec](#amazon_vpc_networking_spec)
- `cilium` - (Optional) - [cilium_networking_spec](#cilium_networking_spec)
- `lyft_vpc` - (Optional) - [lyft_vpc_networking_spec](#lyft_vpc_networking_spec)
- `gcp` - (Optional) - [gcp_networking_spec](#gcp_networking_spec)
- `kindnet` - (Optional) - [kindnet_networking_spec](#kindnet_networking_spec)

### cluster_subnet_spec

ClusterSubnetSpec defines a subnet<br />TODO: move to networking.go.

#### Argument Reference

The following arguments are supported:

- `name` - (Required) - String - Name is the name of the subnet.
- `cidr` - (Optional) - (Computed) - String - CIDR is the IPv4 CIDR block assigned to the subnet.
- `ipv6_cidr` - (Optional) - String - IPv6CIDR is the IPv6 CIDR block assigned to the subnet.
- `zone` - (Required) - String - Zone is the zone the subnet is in, set for subnets that are zonally scoped.
- `region` - (Optional) - String - Region is the region the subnet is in, set for subnets that are regionally scoped.
- `id` - (Optional) - String - ID is the cloud provider ID for the objects associated with the zone (the subnet on AWS).
- `egress` - (Optional) - String - Egress defines the method of traffic egress for this subnet.
- `type` - (Required) - String - Type define which one if the internal types (public, utility, private) the network is.
- `public_ip` - (Optional) - String - PublicIP to attach to NatGateway.
- `additional_routes` - (Optional) - List([route_spec](#route_spec)) - AdditionalRoutes to attach to the subnet's route table.

### route_spec

#### Argument Reference

The following arguments are supported:

- `cidr` - (Optional) - String - CIDR destination of the route.
- `target` - (Optional) - String - Target of the route.

### topology_spec

#### Argument Reference

The following arguments are supported:

- `bastion` - (Optional) - [bastion_spec](#bastion_spec) - Bastion provide an external facing point of entry into a network<br />containing private network instances. This host can provide a single<br />point of fortification or audit and can be started and stopped to enable<br />or disable inbound SSH communication from the Internet. Some call the bastion<br />the "jump server".
- `dns` - (Required) - String - DNS specifies the environment for hosted DNS zones. (Public, Private, None).

### bastion_spec

#### Argument Reference

The following arguments are supported:

- `bastion_public_name` - (Required) - String - PublicName is the domain name for the bastion load balancer.
- `load_balancer` - (Optional) - [bastion_load_balancer_spec](#bastion_load_balancer_spec) - LoadBalancer contains settings for the load balancer fronting bastion instances.

### bastion_load_balancer_spec

#### Argument Reference

The following arguments are supported:

- `type` - (Required) - String - Type of load balancer to create, it can be Public or Internal.

### egress_proxy_spec

#### Argument Reference

The following arguments are supported:

- `http_proxy` - (Required) - [http_proxy](#http_proxy)
- `proxy_excludes` - (Optional) - String

### http_proxy

#### Argument Reference

The following arguments are supported:

- `host` - (Required) - String
- `port` - (Required) - Int

### classic_networking_spec

ClassicNetworkingSpec is the specification of classic networking mode, integrated into kubernetes.<br />Support been removed since Kubernetes 1.4.


This resource has no attributes.

### kubenet_networking_spec

KubenetNetworkingSpec is the specification for kubenet networking, largely integrated but intended to replace classic.


This resource has no attributes.

### external_networking_spec

ExternalNetworkingSpec is the specification for networking that is implemented by a user-provided Daemonset that uses the Kubenet kubelet networking plugin.


This resource has no attributes.

### cni_networking_spec

CNINetworkingSpec is the specification for networking that is implemented by a user-provided Daemonset, which uses the CNI kubelet networking plugin.

#### Argument Reference

The following arguments are supported:

- `uses_secondary_ip` - (Optional) - Bool

### kopeio_networking_spec

KopeioNetworkingSpec declares that we want Kopeio networking.


This resource has no attributes.

### weave_networking_spec

WeaveNetworkingSpec declares that we want Weave networking.

#### Argument Reference

The following arguments are supported:

- `mtu` - (Optional) - Int
- `conn_limit` - (Optional) - Int
- `no_masq_local` - (Optional) - Int
- `memory_request` - (Optional) - Quantity - MemoryRequest memory request of weave container. Default 200Mi.
- `cpu_request` - (Optional) - Quantity - CPURequest CPU request of weave container. Default 50m.
- `memory_limit` - (Optional) - Quantity - MemoryLimit memory limit of weave container. Default 200Mi.
- `cpu_limit` - (Optional) - Quantity - CPULimit CPU limit of weave container.
- `net_extra_args` - (Optional) - String - NetExtraArgs are extra arguments that are passed to weave-kube.
- `npc_memory_request` - (Optional) - Quantity - NPCMemoryRequest memory request of weave npc container. Default 200Mi.
- `npc_cpu_request` - (Optional) - Quantity - NPCCPURequest CPU request of weave npc container. Default 50m.
- `npc_memory_limit` - (Optional) - Quantity - NPCMemoryLimit memory limit of weave npc container. Default 200Mi.
- `npc_cpu_limit` - (Optional) - Quantity - NPCCPULimit CPU limit of weave npc container.
- `npc_extra_args` - (Optional) - String - NPCExtraArgs are extra arguments that are passed to weave-npc.
- `version` - (Optional) - String - Version specifies the Weave container image tag. The default depends on the kOps version.

### flannel_networking_spec

FlannelNetworkingSpec declares that we want Flannel networking.

#### Argument Reference

The following arguments are supported:

- `backend` - (Optional) - String - Backend is the backend overlay type we want to use (vxlan or udp).
- `iptables_resync_seconds` - (Optional) - Int - IptablesResyncSeconds sets resync period for iptables rules, in seconds.

### calico_networking_spec

CalicoNetworkingSpec declares that we want Calico networking.

#### Argument Reference

The following arguments are supported:

- `registry` - (Optional) - String - Registry overrides the Calico container image registry.
- `version` - (Optional) - String - Version overrides the Calico container image tag.
- `allow_ip_forwarding` - (Optional) - Bool - AllowIPForwarding enable ip_forwarding setting within the container namespace.<br />(default: false).
- `aws_src_dst_check` - (Optional) - String - AWSSrcDstCheck enables/disables ENI source/destination checks (AWS IPv4 only)<br />Options: Disable (default for IPv4), Enable, or DoNothing.
- `bpf_enabled` - (Optional) - Bool - BPFEnabled enables the eBPF dataplane mode.
- `bpf_external_service_mode` - (Optional) - String - BPFExternalServiceMode controls how traffic from outside the cluster to NodePorts and ClusterIPs is handled.<br />In Tunnel mode, packet is tunneled from the ingress host to the host with the backing pod and back again.<br />In DSR mode, traffic is tunneled to the host with the backing pod and then returned directly;<br />this requires a network that allows direct return.<br />Default: Tunnel (other options: DSR).
- `bpf_kube_proxy_iptables_cleanup_enabled` - (Optional) - Bool - BPFKubeProxyIptablesCleanupEnabled controls whether Felix will clean up the iptables rules<br />created by the Kubernetes kube-proxy; should only be enabled if kube-proxy is not running.
- `bpf_log_level` - (Optional) - String - BPFLogLevel controls the log level used by the BPF programs. The logs are emitted<br />to the BPF trace pipe, accessible with the command tc exec BPF debug.<br />Default: Off (other options: Info, Debug).
- `chain_insert_mode` - (Optional) - String - ChainInsertMode controls whether Felix inserts rules to the top of iptables chains, or<br />appends to the bottom. Leaving the default option is safest to prevent accidentally<br />breaking connectivity. Default: 'insert' (other options: 'append').
- `cpu_request` - (Optional) - Quantity - CPURequest CPU request of Calico container. Default: 100m.
- `cross_subnet` - (Optional) - Bool - CrossSubnet is deprecated as of kOps 1.22 and has no effect.
- `encapsulation_mode` - (Optional) - String - EncapsulationMode specifies the network packet encapsulation protocol for Calico to use,<br />employing such encapsulation at the necessary scope per the related CrossSubnet field. In<br />"ipip" mode, Calico will use IP-in-IP encapsulation as needed. In "vxlan" mode, Calico will<br />encapsulate packets as needed using the VXLAN scheme.<br />Options: ipip (default) or vxlan.
- `ip_ip_mode` - (Optional) - String - IPIPMode determines when to use IP-in-IP encapsulation for the default Calico IPv4 pool.<br />It is conveyed to the "calico-node" daemon container via the CALICO_IPV4POOL_IPIP<br />environment variable. EncapsulationMode must be set to "ipip".<br />Options: "CrossSubnet", "Always", or "Never".<br />Default: "CrossSubnet" if EncapsulationMode is "ipip", "Never" otherwise.
- `ipv4_auto_detection_method` - (Optional) - String - IPv4AutoDetectionMethod configures how Calico chooses the IP address used to route<br />between nodes.  This should be set when the host has multiple interfaces<br />and it is important to select the interface used.<br />Options: "first-found" (default), "can-reach=DESTINATION",<br />"interface=INTERFACE-REGEX", or "skip-interface=INTERFACE-REGEX".
- `ipv6_auto_detection_method` - (Optional) - String - IPv6AutoDetectionMethod configures how Calico chooses the IP address used to route<br />between nodes.  This should be set when the host has multiple interfaces<br />and it is important to select the interface used.<br />Options: "first-found" (default), "can-reach=DESTINATION",<br />"interface=INTERFACE-REGEX", or "skip-interface=INTERFACE-REGEX".
- `iptables_backend` - (Optional) - String - IptablesBackend controls which variant of iptables binary Felix uses<br />Default: Auto (other options: Legacy, NFT).
- `log_severity_screen` - (Optional) - String - LogSeverityScreen lets us set the desired log level. (Default: info).
- `mtu` - (Optional) - Int - MTU to be set in the cni-network-config for calico.
- `prometheus_metrics_enabled` - (Optional) - Bool - PrometheusMetricsEnabled can be set to enable the experimental Prometheus<br />metrics server (default: false).
- `prometheus_metrics_port` - (Optional) - Int - PrometheusMetricsPort is the TCP port that the experimental Prometheus<br />metrics server should bind to (default: 9091).
- `prometheus_go_metrics_enabled` - (Optional) - Bool - PrometheusGoMetricsEnabled enables Prometheus Go runtime metrics collection.
- `prometheus_process_metrics_enabled` - (Optional) - Bool - PrometheusProcessMetricsEnabled enables Prometheus process metrics collection.
- `typha_prometheus_metrics_enabled` - (Optional) - Bool - TyphaPrometheusMetricsEnabled enables Prometheus metrics collection from Typha<br />(default: false).
- `typha_prometheus_metrics_port` - (Optional) - Int - TyphaPrometheusMetricsPort is the TCP port the typha Prometheus metrics server<br />should bind to (default: 9093).
- `typha_replicas` - (Optional) - Int - TyphaReplicas is the number of replicas of Typha to deploy.
- `vxlan_mode` - (Optional) - String - VXLANMode determines when to use VXLAN encapsulation for the default Calico IPv4 pool.<br />It is conveyed to the "calico-node" daemon container via the CALICO_IPV4POOL_VXLAN<br />environment variable. EncapsulationMode must be set to "vxlan".<br />Options: "CrossSubnet", "Always", or "Never".<br />Default: "CrossSubnet" if EncapsulationMode is "vxlan", "Never" otherwise.
- `wireguard_enabled` - (Optional) - Bool - WireguardEnabled enables WireGuard encryption for all on-the-wire pod-to-pod traffic<br />(default: false).

### canal_networking_spec

CanalNetworkingSpec declares that we want Canal networking.

#### Argument Reference

The following arguments are supported:

- `chain_insert_mode` - (Optional) - String - ChainInsertMode controls whether Felix inserts rules to the top of iptables chains, or<br />appends to the bottom. Leaving the default option is safest to prevent accidentally<br />breaking connectivity. Default: 'insert' (other options: 'append').
- `cpu_request` - (Optional) - Quantity - CPURequest CPU request of Canal container. Default: 100m.
- `default_endpoint_to_host_action` - (Optional) - String - DefaultEndpointToHostAction allows users to configure the default behaviour<br />for traffic between pod to host after calico rules have been processed.<br />Default: ACCEPT (other options: DROP, RETURN).
- `flanneld_iptables_forward_rules` - (Optional) - Bool - FlanneldIptablesForwardRules configures Flannel to add the<br />default ACCEPT traffic rules to the iptables FORWARD chain. (default: true).
- `iptables_backend` - (Optional) - String - IptablesBackend controls which variant of iptables binary Felix uses<br />Default: Auto (other options: Legacy, NFT).
- `log_severity_sys` - (Optional) - String - LogSeveritySys the severity to set for logs which are sent to syslog<br />Default: INFO (other options: DEBUG, WARNING, ERROR, CRITICAL, NONE).
- `mtu` - (Optional) - Int - MTU to be set in the cni-network-config (default: 1500).
- `prometheus_go_metrics_enabled` - (Optional) - Bool - PrometheusGoMetricsEnabled enables Prometheus Go runtime metrics collection.
- `prometheus_metrics_enabled` - (Optional) - Bool - PrometheusMetricsEnabled can be set to enable the experimental Prometheus<br />metrics server (default: false).
- `prometheus_metrics_port` - (Optional) - Int - PrometheusMetricsPort is the TCP port that the experimental Prometheus<br />metrics server should bind to (default: 9091).
- `prometheus_process_metrics_enabled` - (Optional) - Bool - PrometheusProcessMetricsEnabled enables Prometheus process metrics collection.
- `typha_prometheus_metrics_enabled` - (Optional) - Bool - TyphaPrometheusMetricsEnabled enables Prometheus metrics collection from Typha<br />(default: false).
- `typha_prometheus_metrics_port` - (Optional) - Int - TyphaPrometheusMetricsPort is the TCP port the typha Prometheus metrics server<br />should bind to (default: 9093).
- `typha_replicas` - (Optional) - Int - TyphaReplicas is the number of replicas of Typha to deploy.

### kuberouter_networking_spec

KuberouterNetworkingSpec declares that we want Kube-router networking.


This resource has no attributes.

### romana_networking_spec

RomanaNetworkingSpec declares that we want Romana networking<br />Romana is deprecated as of kOps 1.18 and removed as of kOps 1.19.

#### Argument Reference

The following arguments are supported:

- `daemon_service_ip` - (Optional) - String - DaemonServiceIP is the Kubernetes Service IP for the romana-daemon pod.
- `etcd_service_ip` - (Optional) - String - EtcdServiceIP is the Kubernetes Service IP for the etcd backend used by Romana.

### amazon_vpc_networking_spec

AmazonVPCNetworkingSpec declares that we want Amazon VPC CNI networking.

#### Argument Reference

The following arguments are supported:

- `image` - (Optional) - String - Image is the container image name to use.
- `init_image` - (Optional) - String - InitImage is the init container image name to use.
- `network_policy_agent_image` - (Optional) - String - NetworkPolicyAgentImage is the container image to use for the network policy agent.
- `env` - (Optional) - List([env_var](#env_var)) - Env is a list of environment variables to set in the container.

### cilium_networking_spec

CiliumNetworkingSpec declares that we want Cilium networking.

#### Argument Reference

The following arguments are supported:

- `registry` - (Optional) - String - Registry overrides the default Cilium container registry (quay.io).
- `version` - (Optional) - String - Version is the version of the Cilium agent and the Cilium Operator.
- `memory_request` - (Optional) - Quantity - MemoryRequest memory request of Cilium agent + operator container. (default: 128Mi).
- `cpu_request` - (Optional) - Quantity - CPURequest CPU request of Cilium agent + operator container. (default: 25m).
- `agent_prometheus_port` - (Optional) - Int - AgentPrometheusPort is the port to listen to for Prometheus metrics.<br />Defaults to 9090.
- `metrics` - (Optional) - List(String) - Metrics is a list of metrics to add or remove from the default list of metrics the agent exposes.
- `chaining_mode` - (Optional) - String - ChainingMode allows using Cilium in combination with other CNI plugins.<br />With Cilium CNI chaining, the base network connectivity and IP address management is managed<br />by the non-Cilium CNI plugin, but Cilium attaches eBPF programs to the network devices created<br />by the non-Cilium plugin to provide L3/L4 network visibility, policy enforcement and other advanced features.<br />Default: none.
- `debug` - (Optional) - Bool - Debug runs Cilium in debug mode.
- `disable_endpoint_crd` - (Optional) - Bool - DisableEndpointCRD disables usage of CiliumEndpoint CRD.<br />Default: false.
- `enable_policy` - (Optional) - String - EnablePolicy specifies the policy enforcement mode.<br />"default": Follows Kubernetes policy enforcement.<br />"always": Cilium restricts all traffic if no policy is in place.<br />"never": Cilium allows all traffic regardless of policies in place.<br />If unspecified, "default" policy mode will be used.
- `enable_l7_proxy` - (Optional) - Bool - EnableL7Proxy enables L7 proxy for L7 policy enforcement.<br />Default: true.
- `enable_local_redirect_policy` - (Optional) - Bool - EnableLocalRedirectPolicy that enables pod traffic destined to an IP address and port/protocol<br />tuple or Kubernetes service to be redirected locally to backend pod(s) within a node, using eBPF.<br />https://docs.cilium.io/en/stable/network/kubernetes/local-redirect-policy/<br />Default: false.
- `enable_bpf_masquerade` - (Optional) - Bool - EnableBPFMasquerade enables masquerading packets from endpoints leaving the host with BPF instead of iptables.<br />Default: false.
- `enable_endpoint_health_checking` - (Optional) - Bool - EnableEndpointHealthChecking enables connectivity health checking between virtual endpoints.<br />Default: true.
- `enable_prometheus_metrics` - (Optional) - Bool - EnablePrometheusMetrics enables the Cilium "/metrics" endpoint for both the agent and the operator.
- `enable_encryption` - (Optional) - Bool - EnableEncryption enables Cilium Encryption.<br />Default: false.
- `encryption_type` - (Optional) - String - EncryptionType specifies Cilium Encryption method ("ipsec", "wireguard").<br />Default: ipsec.
- `node_encryption` - (Optional) - Bool - NodeEncryption enables encryption for pure node to node traffic.<br />Default: false.
- `identity_allocation_mode` - (Optional) - String - IdentityAllocationMode specifies in which backend identities are stored ("crd", "kvstore").<br />Default: crd.
- `identity_change_grace_period` - (Optional) - String - IdentityChangeGracePeriod specifies the duration to wait before using a changed identity.<br />Default: 5s.
- `masquerade` - (Optional) - Bool - Masquerade enables masquerading IPv4 traffic to external destinations behind the node IP.<br />Default: false if IPAM is "eni" or in IPv6 mode, otherwise true.
- `agent_pod_annotations` - (Optional) - Map(String) - AgentPodAnnotations makes possible to add additional annotations to cilium agent.<br />Default: none.
- `operator_pod_annotations` - (Optional) - Map(String) - OperatorPodAnnotations makes possible to add additional annotations to cilium operator.<br />Default: none.
- `tunnel` - (Optional) - String - Tunnel specifies the Cilium tunnelling mode. Possible values are "vxlan", "geneve", or "disabled".<br />Default: vxlan.
- `monitor_aggregation` - (Optional) - String - MonitorAggregation sets the level of packet monitoring. Possible values are "low", "medium", or "maximum".<br />Default: medium.
- `bpfct_global_tcp_max` - (Optional) - Int - BPFCTGlobalTCPMax is the maximum number of entries in the TCP CT table.<br />Default: 524288.
- `bpfct_global_any_max` - (Optional) - Int - BPFCTGlobalAnyMax is the maximum number of entries in the non-TCP CT table.<br />Default: 262144.
- `bpflb_algorithm` - (Optional) - String - BPFLBAlgorithm is the load balancing algorithm ("random", "maglev").<br />Default: random.
- `bpflb_maglev_table_size` - (Optional) - String - BPFLBMaglevTableSize is the per service backend table size when going with Maglev (parameter M).<br />Default: 16381.
- `bpfnat_global_max` - (Optional) - Int - BPFNATGlobalMax is the the maximum number of entries in the BPF NAT table.<br />Default: 524288.
- `bpf_neigh_global_max` - (Optional) - Int - BPFNeighGlobalMax is the the maximum number of entries in the BPF Neighbor table.<br />Default: 524288.
- `bpf_policy_map_max` - (Optional) - Int - BPFPolicyMapMax is the maximum number of entries in endpoint policy map.<br />Default: 16384.
- `bpflb_map_max` - (Optional) - Int - BPFLBMapMax is the maximum number of entries in bpf lb service, backend and affinity maps.<br />Default: 65536.
- `bpflb_sock_host_ns_only` - (Optional) - Bool - BPFLBSockHostNSOnly enables skipping socket LB for services when inside a pod namespace,<br />in favor of service LB at the pod interface. Socket LB is still used when in the host namespace.<br />Required by service mesh (e.g., Istio, Linkerd).<br />Default: false.
- `preallocate_bpf_maps` - (Required) - Bool - PreallocateBPFMaps reduces the per-packet latency at the expense of up-front memory allocation.<br />Default: true.
- `sidecar_istio_proxy_image` - (Optional) - String - SidecarIstioProxyImage is the regular expression matching compatible Istio sidecar istio-proxy<br />container image names.<br />Default: cilium/istio_proxy.
- `cluster_name` - (Optional) - String - ClusterName is the name of the cluster. It is only relevant when building a mesh of clusters.
- `cluster_id` - (Optional) - Int - ClusterID is the ID of the cluster. It is only relevant when building a mesh of clusters.<br />Must be a number between 1 and 255.
- `to_fqdns_dns_reject_response_code` - (Optional) - String - ToFQDNsDNSRejectResponseCode sets the DNS response code for rejecting DNS requests.<br />Possible values are "nameError" or "refused".<br />Default: refused.
- `to_fqdns_enable_poller` - (Optional) - Bool - ToFQDNsEnablePoller replaces the DNS proxy-based implementation of FQDN policies<br />with the less powerful legacy implementation.<br />Default: false.
- `ipam` - (Optional) - String - IPAM specifies the IP address allocation mode to use.<br />Possible values are "crd" and "eni".<br />"eni" will use AWS native networking for pods. Eni requires masquerade to be set to false.<br />"crd" will use CRDs for controlling IP address management.<br />"hostscope" will use hostscope IPAM mode.<br />"kubernetes" will use addersing based on node pod CIDR.<br />Default: "kubernetes".
- `install_iptables_rules` - (Optional) - Bool - InstallIptablesRules enables installing the base IPTables rules used for masquerading and kube-proxy.<br />Default: true.
- `auto_direct_node_routes` - (Optional) - Bool - AutoDirectNodeRoutes adds automatic L2 routing between nodes.<br />Default: false.
- `enable_host_reachable_services` - (Optional) - Bool - EnableHostReachableServices configures Cilium to enable services to be<br />reached from the host namespace in addition to pod namespaces.<br />https://docs.cilium.io/en/v1.9/gettingstarted/host-services/<br />Default: false.
- `enable_node_port` - (Optional) - Bool - EnableNodePort replaces kube-proxy with Cilium's BPF implementation.<br />Requires spec.kubeProxy.enabled be set to false.<br />Default: false.
- `etcd_managed` - (Optional) - Bool - EtcdManagd installs an additional etcd cluster that is used for Cilium state change.<br />The cluster is operated by cilium-etcd-operator.<br />Default: false.
- `enable_remote_node_identity` - (Required) - Bool - EnableRemoteNodeIdentity enables the remote-node-identity.<br />Default: true.
- `enable_unreachable_routes` - (Optional) - Bool - EnableUnreachableRoutes enables unreachable routes on pod deletion.<br />Default: false.
- `cni_exclusive` - (Optional) - Bool - CniExclusive configures whether to remove other CNI configuration files.<br />Default: true.
- `hubble` - (Optional) - [hubble_spec](#hubble_spec) - Hubble configures the Hubble service on the Cilium agent.
- `disable_cnp_status_updates` - (Optional) - Bool - DisableCNPStatusUpdates determines if CNP NodeStatus updates will be sent to the Kubernetes api-server.
- `enable_service_topology` - (Optional) - Bool - EnableServiceTopology determine if cilium should use topology aware hints.
- `ingress` - (Optional) - [cilium_ingress_spec](#cilium_ingress_spec) - Ingress specifies the configuration for Cilium Ingress settings.
- `gateway_api` - (Optional) - [cilium_gateway_api_spec](#cilium_gateway_api_spec) - GatewayAPI specifies the configuration for Cilium Gateway API settings.

### hubble_spec

HubbleSpec configures the Hubble service on the Cilium agent.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enabled decides if Hubble is enabled on the agent or not.
- `metrics` - (Optional) - List(String) - Metrics is a list of metrics to collect. If empty or null, metrics are disabled.<br />See https://docs.cilium.io/en/stable/observability/metrics/#hubble-exported-metrics.

### cilium_ingress_spec

CiliumIngressSpec configures Cilium Ingress settings.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enabled specifies whether Cilium Ingress is enabled.
- `enforce_https` - (Optional) - Bool - EnforceHttps specifies whether HTTPS enforcement is enabled for Ingress traffic.<br />Default: true.
- `enable_secrets_sync` - (Optional) - Bool - EnableSecretsSync specifies whether synchronization of secrets is enabled.<br />Default: true.
- `load_balancer_annotation_prefixes` - (Optional) - String - LoadBalancerAnnotationPrefixes specifies annotation prefixes for Load Balancer configuration.<br />Default: "service.beta.kubernetes.io service.kubernetes.io cloud.google.com".
- `default_load_balancer_mode` - (Optional) - String - DefaultLoadBalancerMode specifies the default load balancer mode.<br />Possible values: 'shared' or 'dedicated'<br />Default: dedicated.
- `shared_load_balancer_service_name` - (Optional) - String - SharedLoadBalancerServiceName specifies the name of the shared load balancer service.<br />Default: cilium-ingress.

### cilium_gateway_api_spec

CiliumGatewayAPISpec configures Cilium Gateway API settings.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enabled specifies whether Cilium Gateway API is enabled.
- `enable_secrets_sync` - (Optional) - Bool - EnableSecretsSync specifies whether synchronization of secrets is enabled.<br />Default: true.

### lyft_vpc_networking_spec

LyftVPCNetworkingSpec declares that we want to use the cni-ipvlan-vpc-k8s CNI networking.<br />Lyft VPC is deprecated as of kOps 1.22 and removed as of kOps 1.23.

#### Argument Reference

The following arguments are supported:

- `subnet_tags` - (Optional) - Map(String)

### gcp_networking_spec

GCPNetworkingSpec is the specification of GCP's native networking mode, using IP aliases.


This resource has no attributes.

### kindnet_networking_spec

KindnetNetworkingSpec configures Kindnet settings.

#### Argument Reference

The following arguments are supported:

- `version` - (Optional) - String - Version is the version of the kindnet agent.<br />Default: v1.8.0.
- `network_policies` - (Optional) - Bool - Enable network policies.
- `admin_network_policies` - (Optional) - Bool
- `baseline_admin_network_policies` - (Optional) - Bool
- `dns_caching` - (Optional) - Bool - enable dns caching.
- `nat64` - (Optional) - Bool - enable nat64 on ipv6 clusters.
- `fast_path_threshold` - (Optional) - Int - number of packets in a connection to offload it to the fast path.
- `masquerade` - (Optional) - [kindnet_masquerade_spec](#kindnet_masquerade_spec) - node agent masquerading rules.
- `log_level` - (Optional) - Int - log level.

### kindnet_masquerade_spec

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool
- `non_masquerade_cidrs` - (Optional) - List(String)

### api_spec

APISpec provides configuration details related to the Kubernetes API.

#### Argument Reference

The following arguments are supported:

- `dns` - (Optional) - [dns_access_spec](#dns_access_spec) - DNS will be used to provide configuration for the Kubernetes API's DNS server.
- `load_balancer` - (Optional) - [load_balancer_access_spec](#load_balancer_access_spec) - LoadBalancer is the configuration for the Kubernetes API load balancer.
- `public_name` - (Optional) - String - PublicName is the external DNS name for the Kubernetes API.
- `additional_sans` - (Optional) - List(String) - AdditionalSANs adds additional Subject Alternate Names to the Kubernetes API certificate.
- `access` - (Optional) - List(String) - Access is a list of the CIDRs that can access the Kubernetes API endpoint.

### dns_access_spec


This resource has no attributes.

### load_balancer_access_spec

LoadBalancerAccessSpec provides configuration details related to API LoadBalancer and its access.

#### Argument Reference

The following arguments are supported:

- `class` - (Optional) - String - LoadBalancerClass specifies the class of load balancer to create: Classic, Network.
- `type` - (Required) - String - Type of load balancer to create may Public or Internal.
- `idle_timeout_seconds` - (Optional) - Int - IdleTimeoutSeconds sets the timeout of the api loadbalancer.
- `security_group_override` - (Optional) - String - SecurityGroupOverride overrides the default Kops created SG for the load balancer.
- `additional_security_groups` - (Optional) - List(String) - AdditionalSecurityGroups attaches additional security groups (e.g. sg-123456).
- `use_for_internal_api` - (Optional) - Bool - UseForInternalAPI indicates whether the LB should be used by the kubelet.
- `ssl_certificate` - (Optional) - String - SSLCertificate allows you to specify the ACM cert to be used the LB.
- `ssl_policy` - (Optional) - String - SSLPolicy allows you to overwrite the LB listener's Security Policy.
- `cross_zone_load_balancing` - (Optional) - Bool - CrossZoneLoadBalancing allows you to enable the cross zone load balancing.
- `subnets` - (Optional) - List([load_balancer_subnet_spec](#load_balancer_subnet_spec)) - Subnets allows you to specify the subnets that must be used for the load balancer.
- `access_log` - (Optional) - [access_log_spec](#access_log_spec) - AccessLog is the configuration of access logs.

### load_balancer_subnet_spec

LoadBalancerSubnetSpec provides configuration for subnets used for a load balancer.

#### Argument Reference

The following arguments are supported:

- `name` - (Optional) - String - Name specifies the name of the cluster subnet.
- `private_ipv4_address` - (Optional) - String - PrivateIPv4Address specifies the private IPv4 address to use for a NLB.
- `allocation_id` - (Optional) - String - AllocationID specifies the Elastic IP Allocation ID for use by a NLB.

### access_log_spec

#### Argument Reference

The following arguments are supported:

- `interval` - (Optional) - Int - Interval is the publishing interval in minutes. This parameter is only used with classic load balancer.
- `bucket` - (Optional) - String - Bucket is the S3 bucket name to store the logs in.
- `bucket_prefix` - (Optional) - String - BucketPrefix is the S3 bucket prefix. Logs are stored in the root if not configured.

### authentication_spec

#### Argument Reference

The following arguments are supported:

- `kopeio` - (Optional) - [kopeio_authentication_spec](#kopeio_authentication_spec)
- `aws` - (Optional) - [aws_authentication_spec](#aws_authentication_spec)
- `oidc` - (Optional) - [oidc_authentication_spec](#oidc_authentication_spec)

### kopeio_authentication_spec


This resource has no attributes.

### aws_authentication_spec

#### Argument Reference

The following arguments are supported:

- `image` - (Optional) - String - Image is the AWS IAM Authenticator container image to use.
- `backend_mode` - (Optional) - String - BackendMode is the AWS IAM Authenticator backend to use. Default MountedFile.
- `cluster_id` - (Optional) - String - ClusterID identifies the cluster performing authentication to prevent certain replay attacks. Default master public DNS name.
- `memory_request` - (Optional) - Quantity - MemoryRequest memory request of AWS IAM Authenticator container. Default 20Mi.
- `cpu_request` - (Optional) - Quantity - CPURequest CPU request of AWS IAM Authenticator container. Default 10m.
- `memory_limit` - (Optional) - Quantity - MemoryLimit memory limit of AWS IAM Authenticator container. Default 20Mi.
- `cpu_limit` - (Optional) - Quantity - CPULimit CPU limit of AWS IAM Authenticator container. Default 10m.
- `identity_mappings` - (Optional) - List([aws_authentication_identity_mapping_spec](#aws_authentication_identity_mapping_spec)) - IdentityMappings maps IAM Identities to Kubernetes users/groups.

### aws_authentication_identity_mapping_spec

#### Argument Reference

The following arguments are supported:

- `arn` - (Optional) - String - Arn of the IAM User or IAM Role to be allowed to authenticate.
- `username` - (Optional) - String - Username that Kubernetes will see the user as.
- `groups` - (Optional) - List(String) - Groups to be attached to your users/roles.

### oidc_authentication_spec

#### Argument Reference

The following arguments are supported:

- `username_claim` - (Optional) - String - UsernameClaim is the OpenID claim to use as the username.<br />Note that claims other than the default ('sub') are not guaranteed to be<br />unique and immutable.
- `username_prefix` - (Optional) - String - UsernamePrefix is the prefix prepended to username claims to prevent<br />clashes with existing names (such as 'system:' users).
- `groups_claims` - (Optional) - List(String) - GroupsClaims are the names of the custom OpenID Connect claims for<br />specifying user groups (optional).
- `groups_prefix` - (Optional) - String - GroupsPrefix is the prefix prepended to group claims to prevent<br />clashes with existing names (such as 'system:' groups).
- `issuer_url` - (Optional) - String - IssuerURL is the URL of the OpenID issuer. Only the HTTPS scheme will<br />be accepted.<br />If set, will be used to verify the OIDC JSON Web Token (JWT).
- `client_id` - (Optional) - String - ClientID is the client ID for the OpenID Connect client. Must be set<br />if issuerURL is set.
- `required_claims` - (Optional) - Map(String) - RequiredClaims are key/value pairs that describe required claims in the ID Token.<br />If set, the claims are verified to be present in the ID Token with corresponding values.

### authorization_spec

#### Argument Reference

The following arguments are supported:

- `always_allow` - (Optional) - [always_allow_authorization_spec](#always_allow_authorization_spec)
- `rbac` - (Optional) - [rbac_authorization_spec](#rbac_authorization_spec)

### always_allow_authorization_spec


This resource has no attributes.

### rbac_authorization_spec


This resource has no attributes.

### node_authorization_spec

NodeAuthorizationSpec is used to node authorization.

#### Argument Reference

The following arguments are supported:

- `node_authorizer` - (Optional) - [node_authorizer_spec](#node_authorizer_spec) - NodeAuthorizer defined the configuration for the node authorizer.

### node_authorizer_spec

NodeAuthorizerSpec defines the configuration for a node authorizer.

#### Argument Reference

The following arguments are supported:

- `authorizer` - (Optional) - String - Authorizer is the authorizer to use.
- `features` - (Optional) - List(String) - Features is a series of authorizer features to enable or disable.
- `image` - (Optional) - String - Image is the location of container.
- `node_url` - (Optional) - String - NodeURL is the node authorization service url.
- `port` - (Optional) - Int - Port is the port the service is running on the master.
- `interval` - (Optional) - Duration - Interval the time between retires for authorization request.
- `timeout` - (Optional) - Duration - Timeout the max time for authorization request.
- `token_ttl` - (Optional) - Duration - TokenTTL is the max ttl for an issued token.

### hook_spec

HookSpec is a definition hook.

#### Argument Reference

The following arguments are supported:

- `name` - (Required) - String - Name is an optional name for the hook, otherwise the name is kops-hook-<index>.
- `enabled` - (Optional) - Bool - Enabled indicates if you want the unit switched on. Default: true.
- `roles` - (Optional) - List(String) - Roles is an optional list of roles the hook should be rolled out to, defaults to all.
- `requires` - (Optional) - List(String) - Requires is a series of systemd units the action requires.
- `before` - (Optional) - List(String) - Before is a series of systemd units which this hook must run before.
- `exec_container` - (Optional) - [exec_container_action](#exec_container_action) - ExecContainer is the image itself.
- `manifest` - (Optional) - String - Manifest is a raw systemd unit file.
- `use_raw_manifest` - (Optional) - Bool - UseRawManifest indicates that the contents of Manifest should be used as the contents<br />of the systemd unit, unmodified. Before and Requires are ignored when used together<br />with this value (and validation shouldn't allow them to be set).

### exec_container_action

ExecContainerAction defines an hood action.

#### Argument Reference

The following arguments are supported:

- `image` - (Required) - String - Image is the container image.
- `command` - (Optional) - List(String) - Command is the command supplied to the above image.
- `environment` - (Optional) - Map(String) - Environment is a map of environment variables added to the hook.

### assets_spec

AssetsSpec defines the privately hosted assets.

#### Argument Reference

The following arguments are supported:

- `container_registry` - (Optional) - String - ContainerRegistry is a url for to a container registry.
- `file_repository` - (Optional) - String - FileRepository is the url for a private file serving repository.
- `container_proxy` - (Optional) - String - ContainerProxy is a url for a pull-through proxy of a container registry.

### iam_spec

IAMSpec adds control over the IAM security policies applied to resources.

#### Argument Reference

The following arguments are supported:

- `legacy` - (Optional) - Bool
- `allow_container_registry` - (Optional) - Bool
- `permissions_boundary` - (Optional) - String
- `use_service_account_external_permissions` - (Optional) - Bool - UseServiceAccountExternalPermissions determines if managed ServiceAccounts will use external permissions directly.<br />If this is set to false, ServiceAccounts will assume external permissions from the instances they run on.
- `service_account_external_permissions` - (Optional) - List([service_account_external_permission](#service_account_external_permission)) - ServiceAccountExternalPermissions defines the relationship between Kubernetes ServiceAccounts and permissions with external resources.

### service_account_external_permission

ServiceAccountExternalPermissions grants a ServiceAccount permissions to external resources.

#### Argument Reference

The following arguments are supported:

- `name` - (Optional) - String - Name is the name of the Kubernetes ServiceAccount.
- `namespace` - (Optional) - String - Namespace is the namespace of the Kubernetes ServiceAccount.
- `aws` - (Optional) - [aws_permission](#aws_permission) - AWS grants permissions to AWS resources.

### aws_permission

AWSPermission grants permissions to AWS resources.

#### Argument Reference

The following arguments are supported:

- `policy_ar_ns` - (Optional) - List(String) - PolicyARNs is a list of existing IAM Policies.
- `inline_policy` - (Optional) - String - InlinePolicy is an IAM Policy that will be attached inline to the IAM Role.

### rolling_update

#### Argument Reference

The following arguments are supported:

- `drain_and_terminate` - (Optional) - Bool - DrainAndTerminate enables draining and terminating nodes during rolling updates.<br />Defaults to true.
- `max_unavailable` - (Optional) - IntOrString - MaxUnavailable is the maximum number of nodes that can be unavailable during the update.<br />The value can be an absolute number (for example 5) or a percentage of desired<br />nodes (for example 10%).<br />The absolute number is calculated from a percentage by rounding down.<br />Defaults to 1 if MaxSurge is 0, otherwise defaults to 0.<br />Example: when this is set to 30%, the InstanceGroup can be scaled<br />down to 70% of desired nodes immediately when the rolling update<br />starts. Once new nodes are ready, more old nodes can be drained,<br />ensuring that the total number of nodes available at all times<br />during the update is at least 70% of desired nodes.<br />+optional.
- `max_surge` - (Optional) - IntOrString - MaxSurge is the maximum number of extra nodes that can be created<br />during the update.<br />The value can be an absolute number (for example 5) or a percentage of<br />desired machines (for example 10%).<br />The absolute number is calculated from a percentage by rounding up.<br />Has no effect on instance groups with role "Master".<br />Defaults to 1 on AWS, 0 otherwise.<br />Example: when this is set to 30%, the InstanceGroup can be scaled<br />up immediately when the rolling update starts, such that the total<br />number of old and new nodes do not exceed 130% of desired<br />nodes.<br />+optional.

### cluster_autoscaler_config

ClusterAutoscalerConfig determines the cluster autoscaler configuration.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enabled enables the cluster autoscaler.<br />Default: false.
- `expander` - (Optional) - String - Expander determines the strategy for which instance group gets expanded.<br />Supported values: least-waste, most-pods, random, price, priority.<br />The price expander is only supported on GCE.<br />By default, kOps will generate the priority expander ConfigMap based on the `autoscale` and `autoscalePriority` fields in the InstanceGroup specs.<br />Default: least-waste.
- `balance_similar_node_groups` - (Optional) - Bool - BalanceSimilarNodeGroups makes the cluster autoscaler treat similar node groups as one.<br />Default: false.
- `emit_per_nodegroup_metrics` - (Optional) - Bool - EmitPerNodegroupMetrics If true, publishes the node groups min and max metrics count set on the cluster autoscaler.<br />Default: false.
- `aws_use_static_instance_list` - (Optional) - Bool - AWSUseStaticInstanceList makes cluster autoscaler to use statically defined set of AWS EC2 Instance List.<br />Default: false.
- `ignore_daemon_sets_utilization` - (Optional) - Bool - IgnoreDaemonSetsUtilization causes the cluster autoscaler to ignore DaemonSet-managed pods when calculating resource utilization for scaling down.<br />Default: false.
- `scale_down_utilization_threshold` - (Optional) - String - ScaleDownUtilizationThreshold determines the utilization threshold for node scale-down.<br />Default: 0.5.
- `skip_nodes_with_custom_controller_pods` - (Optional) - Bool - SkipNodesWithCustomControllerPods makes the cluster autoscaler skip scale-down of nodes with pods owned by custom controllers.<br />Default: true.
- `skip_nodes_with_system_pods` - (Required) - Bool - SkipNodesWithSystemPods makes the cluster autoscaler skip scale-down of nodes with non-DaemonSet pods in the kube-system namespace.<br />Default: true.
- `skip_nodes_with_local_storage` - (Required) - Bool - SkipNodesWithLocalStorage makes the cluster autoscaler skip scale-down of nodes with local storage.<br />Default: true.
- `new_pod_scale_up_delay` - (Optional) - String - NewPodScaleUpDelay causes the cluster autoscaler to ignore unschedulable pods until they are a certain "age", regardless of the scan-interval<br />Default: 0s.
- `scale_down_delay_after_add` - (Optional) - String - ScaleDownDelayAfterAdd determines the time after scale up that scale down evaluation resumes<br />Default: 10m0s.
- `scale_down_unneeded_time` - (Optional) - String - scaleDownUnneededTime determines the time a node should be unneeded before it is eligible for scale down<br />Default: 10m0s.
- `scale_down_unready_time` - (Optional) - String - ScaleDownUnreadyTime determines the time an unready node should be unneeded before it is eligible for scale down<br />Default: 20m0s.
- `cordon_node_before_terminating` - (Optional) - Bool - CordonNodeBeforeTerminating should CA cordon nodes before terminating during downscale process<br />Default: false.
- `image` - (Optional) - String - Image is the container image used.<br />Default: the latest supported image for the specified kubernetes version.
- `memory_request` - (Optional) - Quantity - MemoryRequest of cluster autoscaler container.<br />Default: 300Mi.
- `cpu_request` - (Optional) - Quantity - CPURequest of cluster autoscaler container.<br />Default: 100m.
- `max_node_provision_time` - (Optional) - String - MaxNodeProvisionTime determines how long CAS will wait for a node to join the cluster.
- `pod_annotations` - (Optional) - Map(String) - PodAnnotations are the annotations added to cluster autoscaler pods when they are created.<br />Default: none.
- `create_priority_expender_config` - (Optional) - Bool - CreatePriorityExpenderConfig makes kOps create the priority-expander ConfigMap<br />Default: true.
- `custom_priority_expander_config` - (Optional) - Map(List(String)) - CustomPriorityExpanderConfig overides the priority-expander ConfigMap with the provided configuration. Any InstanceGroup configuration will be ignored if this is set.<br />This could be useful in order to use regex on priorities configuration.

### service_account_issuer_discovery_config

ServiceAccountIssuerDiscoveryConfig configures an OIDC Issuer.

#### Argument Reference

The following arguments are supported:

- `discovery_store` - (Optional) - String - DiscoveryStore is the VFS path to where OIDC Issuer Discovery metadata is stored.
- `enable_aws_oidc_provider` - (Optional) - Bool - EnableAWSOIDCProvider will provision an AWS OIDC provider that trusts the ServiceAccount Issuer.
- `additional_audiences` - (Optional) - List(String) - AdditionalAudiences adds user defined audiences to the provisioned AWS OIDC provider.

### snapshot_controller_config

SnapshotControllerConfig is the config for the CSI Snapshot Controller.

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool - Enabled enables the CSI Snapshot Controller.
- `install_default_class` - (Optional) - Bool - InstallDefaultClass will install the default VolumeSnapshotClass.

### karpenter_config

#### Argument Reference

The following arguments are supported:

- `enabled` - (Optional) - Bool
- `log_encoding` - (Optional) - String
- `log_level` - (Optional) - String
- `image` - (Optional) - String
- `memory_limit` - (Optional) - Quantity
- `memory_request` - (Optional) - Quantity
- `cpu_request` - (Optional) - Quantity

### cluster_secrets

ClusterSecrets defines cluster secrets.

#### Argument Reference

The following arguments are supported:

- `docker_config` - (Optional) - (Sensitive) - String - DockerConfig holds a valid docker config.<br />After creating a dockerconfig secret, a /root/.docker/config.json file will be added to newly created nodes.<br />This file will be used by Kubernetes to authenticate to container registries and will also work when using containerd as container runtime.
- `cluster_ca_cert` - (Optional) - (Sensitive) - (Computed) - String
- `cluster_ca_key` - (Optional) - (Sensitive) - (Computed) - String

### delete_options

#### Argument Reference

The following arguments are supported:

- `interval` - (Optional) - Duration - Interval is the time to wait between deletion attempts (default 10s).
- `wait` - (Optional) - Duration - Wait is the amount of time to wait for the cluster resources to de deleted (default 10m0s).
- `count` - (Optional) - Int - Count is the number of deletion retries to make before giving up (default 0).


## Import

You can import an existing cluster by creating a `kops_cluster` configuration
and running the `terraform import` command:

1. Create a terraform configuration:

    ```hcl
    provider "kops" {
      state_store = "s3://cluster.example.com"
    }

    resource "kops_cluster" "cluster" {
      name        = "cluster.example.com"
      
      // ....
    }
    ```

1. Run `terraform import`:

    ```shell
    terraform import kops_cluster.cluster cluster.example.com
    ```

