resource "aws_s3_object" "index" {
  bucket = aws_s3_bucket.dsrc_weather.id
  key    = "index.html"
  source = "${path.module}/../www/index.html"  # Local file path
  content_type = "text/html"
  cache_control = "no-cache, no-store, must-revalidate"
}

resource "aws_s3_object" "favicon" {
  bucket        = aws_s3_bucket.dsrc_weather.id
  key           = "favicon.ico"
  source        = "${path.module}/../www/favicon.ico"  # Local file path
  content_type  = "image/x-icon"
  cache_control = "no-cache, no-store, must-revalidate"
}