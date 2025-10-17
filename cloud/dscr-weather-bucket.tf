resource "aws_s3_bucket" "dsrc_weather" {
  bucket = "dsrc-weather-laetus-uk-site"
  force_destroy = true
  tags = {
    Name = "dsrc-weather Bucket"
  }
}

resource "aws_s3_bucket_website_configuration" "dsrc_weather" {
  bucket = aws_s3_bucket.dsrc_weather.id

  index_document {
    suffix = "index.html"
  }
  error_document {
    key = "error.html"
  }
}

resource "aws_s3_bucket_policy" "dsrc_weather_policy" {
  bucket = aws_s3_bucket.dsrc_weather.id

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
        Resource = "${aws_s3_bucket.dsrc_weather.arn}/*"
        Condition = {
          StringEquals = {
            "AWS:SourceArn" = "arn:aws:cloudfront::${data.aws_caller_identity.current.account_id}:distribution/${aws_cloudfront_distribution.dsrc_weather.id}"
          }
        }
      }
    ]
  })
}

data "aws_caller_identity" "current" {}

resource "null_resource" "sync_www_folder" {
  provisioner "local-exec" {
    command = "aws s3 sync ../www s3://${aws_s3_bucket.dsrc_weather.bucket} --delete"
  }

  triggers = {
    always_run = random_pet.always_run.id
  }
}