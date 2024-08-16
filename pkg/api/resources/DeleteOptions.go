package resources

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type DeleteOptions struct {
	// Interval is the time to wait between deletion attempts (default 10s)
	Interval *metav1.Duration
	// Wait is the amount of time to wait for the cluster resources to de deleted (default 10m0s)
	Wait *metav1.Duration
	// Count is the number of deletion retries to make before giving up (default 0)
	Count *int
}
