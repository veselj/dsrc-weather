resource "aws_s3_object" "index" {
  bucket = aws_s3_bucket.dsrc_weather.id
  key    = "index.html"
  source = "${path.module}/../www/index.html"  # Local file path
  content_type = "text/html"
}
