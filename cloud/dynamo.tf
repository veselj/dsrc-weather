resource "aws_dynamodb_table" "weather_samples" {
  name           = "WeatherSamples"
  billing_mode   = "PROVISIONED"

  read_capacity  = 3
  write_capacity = 3

  hash_key       = "Bucket"
  range_key      = "When"

  attribute {
    name = "Bucket"
    type = "S"
  }

  attribute {
    name = "When"
    type = "N"
  }

  ttl {
    attribute_name = "expires_at"
    enabled        = true
  }
  tags = {
    Environment = "dev"
    Name        = "WeatherSamples"
  }
}
