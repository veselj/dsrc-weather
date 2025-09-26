resource "aws_s3_object" "index" {
  bucket = aws_s3_bucket.testsite.id
  key    = "index.html"
  source = "${path.module}/../testsite/index.html"  # Local file path
  content_type = "text/html"
}
