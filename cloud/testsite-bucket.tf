resource "aws_s3_bucket" "testsite" {
  bucket = "laetus-uk-testsite-site"
  force_destroy = true
  tags = {
    Name = "Testsite Bucket"
  }
}

resource "aws_s3_bucket_website_configuration" "testsite" {
  bucket = aws_s3_bucket.testsite.id

  index_document {
    suffix = "index.html"
  }
  error_document {
    key = "error.html"
  }
}

resource "aws_s3_bucket_policy" "testsite_policy" {
  bucket = aws_s3_bucket.testsite.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid = "AllowCloudFrontServicePrincipalReadOnly"
        Effect = "Allow"
        Principal = {
          Service = "cloudfront.amazonaws.com"
        }
        Action = [
          "s3:GetObject"
        ]
        Resource = "${aws_s3_bucket.testsite.arn}/*"
        Condition = {
          StringEquals = {
            "AWS:SourceArn" = "arn:aws:cloudfront::${data.aws_caller_identity.current.account_id}:distribution/${aws_cloudfront_distribution.testsite.id}"
          }
        }
      }
    ]
  })
}

data "aws_caller_identity" "current" {}