resource "random_string" "bucket_suffix" {
  length  = 6
  special = false
  upper   = false
}

resource "aws_s3_bucket" "kops_bucket" {
  bucket = "${local.clusterNameSanitized}-${random_string.bucket_suffix.id}"

  force_destroy = true

  tags = {
    Name        = "kOps State bucket"
    Environment = local.clusterName
  }
}