resource "aws_cloudfront_origin_access_control" "dsrc_weather_oac" {
  name                              = "dsrc_weather_oac"
  description                       = "OAC for dsrc-weather S3 bucket"
  origin_access_control_origin_type  = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

resource "aws_cloudfront_distribution" "dsrc_weather" {
  enabled             = true
  comment             = "dsrc_weather.laetus.uk distribution"
  default_root_object = "index.html"
  # CloudFront is a global service, so region doesn't matter (but you can use eu-west-1 provider)
  depends_on = [aws_acm_certificate_validation.dsrc_weather]

  origin {
    domain_name              = aws_s3_bucket.dsrc_weather.bucket_regional_domain_name
    origin_id                = "testsiteOrigin"
    origin_access_control_id = aws_cloudfront_origin_access_control.dsrc_weather_oac.id
  }

  aliases = ["dsrc-weather.laetus.uk"]

  default_cache_behavior {
    allowed_methods        = ["GET", "HEAD", "OPTIONS"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = "testsiteOrigin"
    viewer_protocol_policy = "redirect-to-https"

    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  price_class = "PriceClass_100"

  viewer_certificate {
    acm_certificate_arn            = aws_acm_certificate_validation.dsrc_weather.certificate_arn
    ssl_support_method             = "sni-only"
    minimum_protocol_version       = "TLSv1.2_2021"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  tags = {
    Name = "dsrc_weather CloudFront"
  }
}

resource "random_pet" "always_run" {
  length    = 3
}

resource "null_resource" "cloudfront_invalidate" {
  provisioner "local-exec" {
    command = "aws cloudfront create-invalidation --distribution-id ${aws_cloudfront_distribution.dsrc_weather.id} --paths '/*'"
  }
  triggers = {
    always_run = random_pet.always_run.id
  }
}
