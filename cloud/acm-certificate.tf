resource "aws_acm_certificate" "testsite" {
  domain_name               = "laetus.uk"
  subject_alternative_names = ["testsite.laetus.uk", "www.laetus.uk"]
  validation_method         = "DNS"
  provider          = aws.virginia

  lifecycle {
    create_before_destroy = true
  }

  tags = {
    Name = "Testsite ACM Certificate"
  }
}
