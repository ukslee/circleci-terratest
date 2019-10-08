output "vpc_id" {
  description = "The id of the VPC"
  value       = module.this.vpc_id
}

output "vpc_cidr_block" {
  description = "The CIDR block of the VPC"
  value = module.this.cidr_block
}
