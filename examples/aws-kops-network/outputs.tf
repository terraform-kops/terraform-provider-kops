output "kops_state_bucket" {
  value = "s3://${aws_s3_bucket.kops_bucket.bucket}"
}

output "kops_cluster_name" {
  value = local.clusterName
}