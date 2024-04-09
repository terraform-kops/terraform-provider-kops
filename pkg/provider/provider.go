package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-kops/terraform-provider-kops/pkg/config"
	"github.com/terraform-kops/terraform-provider-kops/pkg/datasources"
	"github.com/terraform-kops/terraform-provider-kops/pkg/resources"
	configschemas "github.com/terraform-kops/terraform-provider-kops/pkg/schemas/config"
)

func NewProvider() *schema.Provider {
	return &schema.Provider{
		Schema: configschemas.ConfigProvider().Schema,
		DataSourcesMap: map[string]*schema.Resource{
			"kops_cluster":        datasources.Cluster(),
			"kops_cluster_status": datasources.ClusterStatus(),
			"kops_instance_group": datasources.InstanceGroup(),
			"kops_kube_config":    datasources.KubeConfig(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"kops_cluster":         resources.Cluster(),
			"kops_cluster_updater": resources.ClusterUpdater(),
			"kops_instance_group":  resources.InstanceGroup(),
		},
		ConfigureContextFunc: config.ConfigureProvider,
	}
}
