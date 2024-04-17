resource "kops_cluster" "cluster" {
  name               = local.clusterName
  admin_ssh_key      = file("${path.module}/../dummy_ssh.pub")
  kubernetes_version = "1.27.10"
    dns_zone         = local.dnsZone

  api {
    access = [
      "0.0.0.0/0"
    ]

    load_balancer {
      additional_security_groups = []
      class                      = "Network"
      cross_zone_load_balancing  = true
      idle_timeout_seconds       = 0
      type                       = "Public"
      use_for_internal_api       = true
    }
  }

  cloud_provider {
    aws {}
  }

  config_store {
    base = "s3://${aws_s3_bucket.kops_bucket.bucket}/${local.clusterName}"
  }

  iam {
    allow_container_registry = true
  }

  networking {
    network_cidr = local.network_cidr

    # private subnets
    subnet {
      name = "private-0"
      type = "Private"
      zone = local.privateSubnets[0].zone
    }
    subnet {
      name = "private-1"
      type = "Private"
      zone = local.privateSubnets[1].zone
    }
    subnet {
      name = "private-2"
      type = "Private"
      zone = local.privateSubnets[2].zone
    }
    subnet {
      name = "utility-0"
      type = "Utility"
      zone = local.utilitySubnets[0].zone
    }
    subnet {
      name = "utility-1"
      type = "Utility"
      zone = local.utilitySubnets[1].zone
    }
    subnet {
      name = "utility-2"
      type = "Utility"
      zone = local.utilitySubnets[2].zone
    }

    cilium {
      enable_remote_node_identity = true
      preallocate_bpf_maps        = false
    }

    topology {
      dns = "Public"
    }
  }

  # etcd clusters
  etcd_cluster {
    name = "main"
    member {
      name           = "control-plane-0"
      instance_group = "control-plane-0"
    }
    member {
      name           = "control-plane-1"
      instance_group = "control-plane-1"
    }
    member {
      name           = "control-plane-2"
      instance_group = "control-plane-2"
    }
  }
  etcd_cluster {
    name = "events"
    member {
      name           = "control-plane-0"
      instance_group = "control-plane-0"
    }
    member {
      name           = "control-plane-1"
      instance_group = "control-plane-1"
    }
    member {
      name           = "control-plane-2"
      instance_group = "control-plane-2"
    }
  }
}

resource "kops_instance_group" "control-plane-0" {
  cluster_name = kops_cluster.cluster.id
  name         = "control-plane-0"
  role         = "ControlPlane"
  min_size     = 1
  max_size     = 1
  machine_type = local.masterType
  subnets      = ["private-0"]
}

resource "kops_instance_group" "control-plane-1" {
  cluster_name = kops_cluster.cluster.id
  name         = "control-plane-1"
  role         = "ControlPlane"
  min_size     = 1
  max_size     = 1
  machine_type = local.masterType
  subnets      = ["private-1"]
}

resource "kops_instance_group" "control-plane-2" {
  cluster_name = kops_cluster.cluster.id
  name         = "control-plane-2"
  role         = "ControlPlane"
  min_size     = 1
  max_size     = 1
  machine_type = local.masterType
  subnets      = ["private-2"]
}

resource "kops_instance_group" "node-0" {
  cluster_name = kops_cluster.cluster.id
  name         = "node-0"
  role         = "Node"
  min_size     = 1
  max_size     = 2
  machine_type = local.nodeType
  subnets      = ["private-0"]
}

resource "kops_instance_group" "node-1" {
  cluster_name = kops_cluster.cluster.id
  name         = "node-1"
  role         = "Node"
  min_size     = 1
  max_size     = 2
  machine_type = local.nodeType
  subnets      = ["private-1"]
}

resource "kops_instance_group" "node-2" {
  cluster_name = kops_cluster.cluster.id
  name         = "node-2"
  role         = "Node"
  min_size     = 1
  max_size     = 2
  machine_type = local.nodeType
  subnets      = ["private-2"]
}

resource "kops_cluster_updater" "updater" {
  cluster_name = kops_cluster.cluster.id

  keepers = {
    cluster         = kops_cluster.cluster.revision
    control-plane-0 = kops_instance_group.control-plane-0.revision
    control-plane-1 = kops_instance_group.control-plane-1.revision
    control-plane-2 = kops_instance_group.control-plane-2.revision
    node-0          = kops_instance_group.node-0.revision
    node-1          = kops_instance_group.node-1.revision
    node-2          = kops_instance_group.node-2.revision
  }

  rolling_update {
    skip                = false
    fail_on_drain_error = true
    fail_on_validate    = true
    validate_count      = 1
  }

  validate {
    skip = false
  }
}

data "kops_kube_config" "kube_config" {
  cluster_name = kops_cluster.cluster.id
}
