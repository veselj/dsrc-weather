resource "aws_acm_certificate" "dsrc_weather" {
  domain_name               = "dsrc-weather.laetus.uk"
  subject_alternative_names = ["rtyc-weather.laetus.uk"]
  validation_method         = "DNS"
  provider          = aws.virginia

  lifecycle {
    create_before_destroy = true
  }

  tags = {
    Name = "Dsrc/Rtyc Weather ACM Certificate"
  }
}
