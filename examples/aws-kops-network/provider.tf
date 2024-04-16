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

provider "helm" {
  kubernetes {
    host                   = data.kops_kube_config.kube_config.server
    username               = data.kops_kube_config.kube_config.kube_user
    password               = data.kops_kube_config.kube_config.kube_password
    client_certificate     = data.kops_kube_config.kube_config.client_cert
    client_key             = data.kops_kube_config.kube_config.client_key
    cluster_ca_certificate = data.kops_kube_config.kube_config.ca_certs
  }
}
