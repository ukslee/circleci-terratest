output "vpc_id" {
  description = "The id of the VPC"
  value       = aws_vpc.global.id
}

output "cidr_block" {
  description = "The CIDR block of the VPC"
  value = aws_vpc.global.cidr_block
}
