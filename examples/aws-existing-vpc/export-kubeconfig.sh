export KOPS_STATE_STORE=`terraform output -raw kops_state_bucket`
export KOPS_CLUSTER_NAME=`terraform output -raw kops_cluster_name`
kops export kubeconfig --admin --auth-plugin

