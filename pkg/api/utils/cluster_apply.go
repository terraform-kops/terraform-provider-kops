package utils

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

func ParseLifecycle(lifecycle string) (fi.Lifecycle, error) {
	if v, ok := fi.LifecycleNameMap[lifecycle]; ok {
		return v, nil
	}
	return "", fmt.Errorf("unknown lifecycle %q, available lifecycle: %s", lifecycle, strings.Join(fi.Lifecycles.List(), ","))
}

func ParseLifecycleOverrides(lifecycleOverrides map[string]string) (map[string]fi.Lifecycle, error) {
	lifecycleOverrideMap := map[string]fi.Lifecycle{}
	for k, v := range lifecycleOverrides {
		if lifecycleOverride, err := ParseLifecycle(v); err != nil {
			return nil, err
		} else {
			lifecycleOverrideMap[k] = lifecycleOverride
		}
	}
	if len(lifecycleOverrideMap) == 0 {
		return nil, nil
	} else {
		return lifecycleOverrideMap, nil
	}
}

func ClusterApply(clientset simple.Clientset, clusterName string, allowKopsDowngrade bool, lifecycleOverrides map[string]fi.Lifecycle, prune bool, controlPlaneOnly bool) (*cloudup.ApplyResults, error) {
	kc, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return nil, err
	}
	deletionProcessing := fi.DeletionProcessingModeDeleteIfNotDeferrred
	if prune {
		deletionProcessing = fi.DeletionProcessingModeDeleteIncludingDeferred
	}
	apply := &cloudup.ApplyClusterCmd{
		Cloud:              cloud,
		Cluster:            kc,
		Clientset:          clientset,
		TargetName:         cloudup.TargetDirect,
		AllowKopsDowngrade: allowKopsDowngrade,
		LifecycleOverrides: lifecycleOverrides,
		DeletionProcessing: deletionProcessing,
	}
	if controlPlaneOnly {
		apply.InstanceGroupFilter = func(group *kops.InstanceGroup) bool {
			for _, role := range []kops.InstanceGroupRole{kops.InstanceGroupRoleControlPlane, kops.InstanceGroupRoleAPIServer} {
				if group.Spec.Role == role {
					return true
				}
			}
			return false
		}
	}
	return apply.Run(context.Background())
}
