resource "aws_route53_record" "dsrc_weather_alias" {
  zone_id = aws_route53_zone.main.zone_id
  name    = "dsrc-weather.laetus.uk"  # If you want example.com, use just "example.com"
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.dsrc_weather.domain_name
    zone_id                = aws_cloudfront_distribution.dsrc_weather.hosted_zone_id
    evaluate_target_health = true
  }
}