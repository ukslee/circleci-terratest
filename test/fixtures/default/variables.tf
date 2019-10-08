variable "aws_region" {
  description = "The AWS region name for the VPC (e.g. ap-northeast-2)"
  type        = string
  default     = "ap-northeast-2"
}

variable "vpc_name" {
  description = "The VPC name to create"
  type        = string
}

variable "cidr_block" {
  description = "The CIDR block for the VPC (e.g. 10.10.0.0/16)"
  type        = string
  default     = "10.10.0.0/16"
}
