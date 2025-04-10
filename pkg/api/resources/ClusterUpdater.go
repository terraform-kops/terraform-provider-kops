package resources

import (
	"github.com/terraform-kops/terraform-provider-kops/pkg/api/utils"
	"k8s.io/kops/pkg/client/simple"
)

// ClusterUpdater takes care of applying changes and/or rolling update the cluster when needed
type ClusterUpdater struct {
	// Revision is incremented every time the resource changes, this is useful for triggering cluster updater
	Revision int
	// ProviderVersion is set to the currently running provider version, this will trigger cluster updater on version changes
	ProviderVersion string
	// ClusterName is the target cluster name
	ClusterName string
	// Keepers contains arbitrary strings used to update the resource when one changes
	Keepers map[string]string
	// Apply holds cluster apply options
	Apply ApplyOptions
	// RollingUpdate holds cluster rolling update options
	RollingUpdate RollingUpdateOptions
	// Validate holds cluster validation options
	Validate ValidateOptions
}

func (u *ClusterUpdater) UpdateCluster(clientset simple.Clientset, isNewResource bool) error {
	// Handle ControlPlanes
	if !u.Apply.Skip {
		if lifecycleOverrides, err := utils.ParseLifecycleOverrides(u.Apply.LifecycleOverrides); err != nil {
			return err
		} else {
			if _, err := utils.ClusterApply(clientset, u.ClusterName, u.Apply.AllowKopsDowngrade, lifecycleOverrides, false, !isNewResource); err != nil {
				return err
			}
		}
	}
	if !u.Validate.Skip {
		if err := utils.ClusterValidate(clientset, u.ClusterName, u.Validate.ValidateOptions); err != nil {
			return err
		}
	}
	if !u.RollingUpdate.Skip {
		if err := utils.ClusterRollingUpdate(clientset, u.ClusterName, u.RollingUpdate.RollingUpdateOptions, true); err != nil {
			return err
		}
	}

	// Nodes are only handled separately during updates
	if !isNewResource {
		// Handle nodes
		if !u.Apply.Skip {
			if lifecycleOverrides, err := utils.ParseLifecycleOverrides(u.Apply.LifecycleOverrides); err != nil {
				return err
			} else {
				if _, err := utils.ClusterApply(clientset, u.ClusterName, u.Apply.AllowKopsDowngrade, lifecycleOverrides, false, false); err != nil {
					return err
				}
			}
		}
	}
	if !u.Validate.Skip {
		if err := utils.ClusterValidate(clientset, u.ClusterName, u.Validate.ValidateOptions); err != nil {
			return err
		}
	}
	if !u.RollingUpdate.Skip {
		if err := utils.ClusterRollingUpdate(clientset, u.ClusterName, u.RollingUpdate.RollingUpdateOptions, false); err != nil {
			return err
		}
	}

	// Prune deferred deletions
	if !u.Apply.Skip {
		if lifecycleOverrides, err := utils.ParseLifecycleOverrides(u.Apply.LifecycleOverrides); err != nil {
			return err
		} else {
			if _, err := utils.ClusterApply(clientset, u.ClusterName, u.Apply.AllowKopsDowngrade, lifecycleOverrides, true, false); err != nil {
				return err
			}
		}
	}
	return nil
}
