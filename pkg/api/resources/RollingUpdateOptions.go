package resources

import "github.com/terraform-kops/terraform-provider-kops/pkg/api/utils"

type RollingUpdateOptions struct {
	// Skip allows skipping cluster rolling update
	Skip bool
	// ExcludeInstanceGroups is a list of instance group names to exclude from rolling updates.
	// Excluded instance groups will have their cloud resources updated (ASG launch template)
	// but existing instances will not be terminated or replaced.
	ExcludeInstanceGroups []string
	utils.RollingUpdateOptions
}
