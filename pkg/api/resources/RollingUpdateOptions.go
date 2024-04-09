package resources

import "github.com/terraform-kops/terraform-provider-kops/pkg/api/utils"

type RollingUpdateOptions struct {
	// Skip allows skipping cluster rolling update
	Skip bool
	utils.RollingUpdateOptions
}
