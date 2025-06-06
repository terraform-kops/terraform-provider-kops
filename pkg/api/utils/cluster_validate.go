package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/validation"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

type ValidateOptions struct {
	// Timeout defines the maximum time to wait until the cluster becomes valid
	Timeout *metav1.Duration
	// PollInterval defines the interval between validation attempts
	PollInterval *metav1.Duration
}

func makeValidator(clientset simple.Clientset, clusterName string) (validation.ClusterValidator, error) {
	kc, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return nil, err
	}
	list, err := clientset.InstanceGroupsFor(kc).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("cannot get InstanceGroups for %q: %v", kc.ObjectMeta.Name, err)
	}
	if len(list.Items) == 0 {
		return nil, fmt.Errorf("no InstanceGroup objects found")
	}
	configBuilder, err := GetKubeConfigBuilder(clientset, clusterName, nil, false)
	if err != nil {
		return nil, err
	}
	config, err := BuildRestConfig(configBuilder)
	if err != nil {
		return nil, err
	}
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("cannot build kubernetes api client for %q: %v", kc.Name, err)
	}
	allowAllInstanceGroups := func(ig *kops.InstanceGroup) bool {
		return true
	}
	allowAllPods := func(pod *v1.Pod) bool {
		return true
	}
	validator, err := validation.NewClusterValidator(kc, cloud, list, allowAllInstanceGroups, allowAllPods, config, k8sClient)
	if err != nil {
		return nil, fmt.Errorf("unexpected error creating validator: %v", err)
	}
	return validator, nil
}

func ClusterIsValid(clientset simple.Clientset, clusterName string) (bool, error) {
	if validator, err := makeValidator(clientset, clusterName); err != nil {
		return false, err
	} else {
		result, err := validator.Validate(context.Background())
		if err != nil {
			return false, err
		}
		if len(result.Failures) != 0 {
			return false, fmt.Errorf("Validation failures")
		}
		return true, nil
	}
}

func ClusterValidate(clientset simple.Clientset, clusterName string, options ValidateOptions) error {
	if validator, err := makeValidator(clientset, clusterName); err != nil {
		return err
	} else {
		timeout := time.Now()
		if options.Timeout != nil {
			timeout = timeout.Add(options.Timeout.Duration)
		} else {
			timeout = timeout.Add(15 * time.Minute)
		}
		pollInterval := 10 * time.Second
		if options.PollInterval != nil {
			pollInterval = options.PollInterval.Duration
		}
		consecutive := 0
		for {
			if time.Now().After(timeout) {
				return fmt.Errorf("wait time exceeded during validation")
			}
			result, err := validator.Validate(context.Background())
			if err != nil {
				consecutive = 0
				log.Printf("(will retry): unexpected error during validation: %v\n", err)
				time.Sleep(pollInterval)
				continue
			}
			if len(result.Failures) == 0 {
				consecutive++
				if consecutive < 0 {
					log.Printf("(will retry): cluster passed validation %d consecutive times\n", consecutive)
					time.Sleep(pollInterval)
					continue
				} else {
					return nil
				}
			} else {
				if consecutive == 0 {
					log.Println("(will retry): cluster not yet healthy")
					time.Sleep(pollInterval)
					continue
				}
			}
		}
	}
}
