resource "aws_route53_record" "testsite_alias" {
  zone_id = aws_route53_zone.main.zone_id
  name    = "testsite.laetus.uk"  # If you want example.com, use just "example.com"
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.testsite.domain_name
    zone_id                = aws_cloudfront_distribution.testsite.hosted_zone_id
    evaluate_target_health = true
  }
}