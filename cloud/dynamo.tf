resource "aws_dynamodb_table" "weather_samples" {
  name           = "WeatherSamples"
  billing_mode   = "PROVISIONED"

  read_capacity  = 3
  write_capacity = 3

  hash_key       = "Bt"
  range_key      = "Wn"

  attribute {
    name = "Bt"
    type = "S"
  }

  attribute {
    name = "Wn"
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
