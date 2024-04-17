resource "kops_cluster" "cluster" {
  name               = local.cluster_name
  admin_ssh_key      = file("${path.module}/../dummy_ssh.pub")
  kubernetes_version = "1.27.10"
  dns_zone           = local.dns_zone

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
    base = "s3://${aws_s3_bucket.kops_bucket.bucket}/${local.cluster_name}"
  }

  iam {
    allow_container_registry = true
  }

  networking {
    network_id = module.vpc.vpc_id

    dynamic "subnet" {
      for_each = data.aws_subnet.private
      content {
        type = "Private"
        id   = subnet.value.id
        name = subnet.value.id
        zone = subnet.value.availability_zone
      }
    }
    dynamic "subnet" {
      for_each = data.aws_subnet.utility
      content {
        type = "Utility"
        id   = subnet.value.id
        name = subnet.value.id
        zone = subnet.value.availability_zone
      }
    }
    dynamic "subnet" {
      for_each = data.aws_subnet.public
      content {
        type = "Public"
        id   = subnet.value.id
        name = subnet.value.id
        zone = subnet.value.availability_zone
      }
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

resource "kops_instance_group" "control_plane" {
  for_each = data.aws_subnet.private

  cluster_name = kops_cluster.cluster.id
  name         = "control-plane-${each.key}"
  role         = "ControlPlane"
  min_size     = 1
  max_size     = 1
  machine_type = local.master_type
  subnets      = [each.value.id]
}

resource "kops_instance_group" "node" {
  for_each = data.aws_subnet.private

  cluster_name = kops_cluster.cluster.id
  name         = "node-${each.key}"
  role         = "Node"
  min_size     = 1
  max_size     = 2
  machine_type = local.node_type
  subnets      = [each.value.id]
}

resource "kops_cluster_updater" "updater" {
  cluster_name = kops_cluster.cluster.id

  keepers = {
    cluster       = kops_cluster.cluster.revision
    control-plane = format("%#v", { for k, v in kops_instance_group.control_plane : k => v.revision })
    node          = format("%#v", { for k, v in kops_instance_group.node : k => v.revision })
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

  depends_on = [kops_cluster_updater.updater]
}
