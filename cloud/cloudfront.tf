resource "aws_cloudfront_origin_access_control" "testsite_oac" {
  name                              = "testsite-oac"
  description                       = "OAC for testsite S3 bucket"
  origin_access_control_origin_type  = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

resource "aws_cloudfront_distribution" "testsite" {
  enabled             = true
  comment             = "CloudFront for example.com/testsite"
  default_root_object = "index.html"
  # CloudFront is a global service, so region doesn't matter (but you can use eu-west-1 provider)
  depends_on = [aws_acm_certificate_validation.testsite]

  origin {
    domain_name              = aws_s3_bucket.testsite.bucket_regional_domain_name
    origin_id                = "testsiteOrigin"
    origin_access_control_id = aws_cloudfront_origin_access_control.testsite_oac.id
  }

  aliases = ["testsite.laetus.uk", "laetus.uk", "www.laetus.uk"]

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
    acm_certificate_arn            = aws_acm_certificate_validation.testsite.certificate_arn
    ssl_support_method             = "sni-only"
    minimum_protocol_version       = "TLSv1.2_2021"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  tags = {
    Name = "Testsite CloudFront"
  }
}