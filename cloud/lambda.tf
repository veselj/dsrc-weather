
resource "aws_lambda_function" "weather_data" {
  function_name = "weather-data"
  handler       = "bootstrap"
  runtime       = "provided.al2023"
  architectures = ["x86_64"]
  role          = aws_iam_role.lambda_exec.arn
  filename      = data.archive_file.lambda_zip.output_path
  # ...other config...
  depends_on = [data.archive_file.lambda_zip]
}
resource "aws_iam_role" "lambda_exec" {
  name = "lambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "${path.module}/../weather-data/bin/bootstrap"
  output_path = "${path.module}/weather-data.zip"
}

resource "aws_lambda_function_url" "weather_data_url" {
  function_name      = aws_lambda_function.weather_data.function_name
  authorization_type = "NONE"

  cors {
    allow_origins  = ["*"]
    allow_methods  = ["GET"]
    allow_headers  = ["Content-Type"]
    expose_headers    = ["keep-alive", "date"]
    max_age           = 86400
  }
}
