module "vpc" {
  source          = "terraform-aws-modules/vpc/aws"
  version         = "5.7.1"
  name            = local.cluster_name
  cidr            = "10.0.0.0/16"
  azs             = ["${local.region}a", "${local.region}b", "${local.region}c"]
  private_subnets = ["10.0.0.0/19", "10.0.32.0/19", "10.0.64.0/19"]
  public_subnets = concat(
    ["10.0.192.0/20", "10.0.176.0/20", "10.0.160.0/20"],
    ["10.0.240.0/20", "10.0.224.0/20", "10.0.208.0/20"]
  )
  enable_nat_gateway                   = true
  single_nat_gateway                   = false
  one_nat_gateway_per_az               = false
  enable_dns_support                   = true
  enable_dns_hostnames                 = true
  enable_flow_log                      = false
  create_flow_log_cloudwatch_iam_role  = false
  create_flow_log_cloudwatch_log_group = false
  private_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    "karpenter.sh/discovery"                      = local.cluster_name
    Type                                          = "private"
  }
  public_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    "karpenter.sh/discovery"                      = local.cluster_name
    Type                                          = "public"
  }
  tags = merge(local.tags, {
    Name = local.cluster_name
  })
}

data "aws_subnet" "private" {
  for_each = {
    for k, v in module.vpc.private_subnets : k => v
  }

  #   count  = length(module.vpc.private_subnets)
  id     = each.value
  vpc_id = module.vpc.vpc_id
}

data "aws_subnet" "public" {
  count  = 3
  id     = module.vpc.public_subnets[count.index]
  vpc_id = module.vpc.vpc_id
}

data "aws_subnet" "utility" {
  count  = 3
  id     = module.vpc.public_subnets[count.index + 3]
  vpc_id = module.vpc.vpc_id
}
