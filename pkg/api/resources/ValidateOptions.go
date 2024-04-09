package resources

import "github.com/terraform-kops/terraform-provider-kops/pkg/api/utils"

type ValidateOptions struct {
	// Skip allows skipping cluster validation
	Skip bool
	utils.ValidateOptions
}
