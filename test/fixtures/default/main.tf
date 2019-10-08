module "this" {
  source = "../../../"

  cidr_block = var.cidr_block
  vpc_name   = var.vpc_name
}
