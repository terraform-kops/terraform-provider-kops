locals {
  masterType           = "t3.medium"
  nodeType             = "t3.medium"
  clusterName          = "cluster.example.com"
  clusterNameSanitized = "cluster-example-com"
  dnsZone              = "cluster.example.com"
  network_cidr         = "10.23.0.0/16"
  privateSubnets = [
    { subnetId = "private-subnet-0", zone = "us-west-2a" },
    { subnetId = "private-subnet-1", zone = "us-west-2b" },
    { subnetId = "private-subnet-2", zone = "us-west-2c" }
  ]
  utilitySubnets = [
    { subnetId = "utility-subnet-0", zone = "us-west-2a" },
    { subnetId = "utility-subnet-1", zone = "us-west-2b" },
    { subnetId = "utility-subnet-2", zone = "us-west-2c" }
  ]
}
