terraform {
  required_providers {
    kops = {
      source  = "github/terraform-kops/kops"
      version = "0.0.1"
    }
  }
}

provider "kops" {
  state_store = "s3://${aws_s3_bucket.kops_bucket.bucket}"
}
