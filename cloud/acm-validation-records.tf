resource "aws_route53_record" "cert_validation" {
  for_each = {
    for dvo in aws_acm_certificate.dsrc_weather.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      type   = dvo.resource_record_type
      value  = dvo.resource_record_value
    }
  }

  # Route 53 is global, so you don't need provider aliasing for zone
  zone_id = aws_route53_zone.main.zone_id
  name    = each.value.name
  type    = each.value.type
  records = [each.value.value]
  ttl     = 60
}

resource "aws_acm_certificate_validation" "dsrc_weather" {
  provider                = aws.virginia
  certificate_arn         = aws_acm_certificate.dsrc_weather.arn
  validation_record_fqdns = [for record in aws_route53_record.cert_validation : record.fqdn]
}