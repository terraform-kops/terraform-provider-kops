locals {
  master_type            = "t3.medium"
  node_type              = "t3.medium"
  cluster_name           = "cluster.example.com"
  cluster_name_sanitized = "cluster-example-com"
  dns_zone               = "cluster.example.com"
  region                 = "us-west-2"

  tags = {
    Cluster = local.cluster_name
  }

}
