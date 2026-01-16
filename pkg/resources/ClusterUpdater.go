package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-kops/terraform-provider-kops/pkg/config"
	"github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	resourcesschema "github.com/terraform-kops/terraform-provider-kops/pkg/schemas/resources"
)

func ClusterUpdater() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: ClusterUpdaterCreateOrUpdate,
		ReadContext:          schema.NoopContext,
		UpdateWithoutTimeout: ClusterUpdaterCreateOrUpdate,
		DeleteWithoutTimeout: ClusterUpdaterDelete,
		CustomizeDiff: func(c context.Context, d *schema.ResourceDiff, m interface{}) error {
			if err := schemas.CustomizeDiffRevision(c, d, m); err != nil {
				return err
			}
			if err := schemas.CustomizeDiffVersion(c, d, m); err != nil {
				return err
			}
			return nil
		},
		Schema: resourcesschema.ResourceClusterUpdater().Schema,
	}
}

func ClusterUpdaterCreateOrUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	oldState, newState := d.GetChange("")
	hasNewResource := oldState == nil
	if oldState != nil {
		o := resourcesschema.ExpandResourceClusterUpdater(oldState.(map[string]interface{}))
		n := resourcesschema.ExpandResourceClusterUpdater(newState.(map[string]interface{}))
		for k := range n.Keepers {
			if _, exists := o.Keepers[k]; !exists {
				hasNewResource = true
				break
			}
		}
	}
	in := resourcesschema.ExpandResourceClusterUpdater(d.Get("").(map[string]interface{}))
	if err := in.UpdateCluster(config.Clientset(m), hasNewResource); err != nil {
		// On error, restore old computed values to prevent state drift
		if oldState != nil {
			oldMap := oldState.(map[string]interface{})
			if oldRevision, ok := oldMap["revision"]; ok {
				d.Set("revision", oldRevision)
			}
			if oldVersion, ok := oldMap["provider_version"]; ok {
				d.Set("provider_version", oldVersion)
			}
		}
		return diag.FromErr(err)
	}
	d.SetId(in.ClusterName)
	return nil
}

func ClusterUpdaterDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.FromErr(schema.RemoveFromState(d, m))
}
