resource "aws_route53_zone" "main" {
  name = "laetus.uk"
  comment = "Hosted zone for laetus.uk managed by Terraform"

  tags = {
    Environment = "dev"
    ManagedBy   = "Terraform"
  }
}