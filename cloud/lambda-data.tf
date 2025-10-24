
resource "aws_lambda_function" "weather_data" {
  function_name = "weather-data"
  handler       = "bootstrap"
  runtime       = "provided.al2023"
  architectures = ["x86_64"]
  role          = aws_iam_role.lambda_data_exec.arn
  filename      = data.archive_file.lambda_data_zip.output_path
  # ...other config...
  depends_on = [data.archive_file.lambda_data_zip]
}
resource "aws_iam_role" "lambda_data_exec" {
  name = "lambda_data_exec_role"

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

resource "aws_iam_role_policy_attachment" "lambda_data_logs" {
  role       = aws_iam_role.lambda_data_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_cloudwatch_log_group" "weather_data" {
  name              = "/aws/lambda/weather-data"
  retention_in_days = 14
}

data "archive_file" "lambda_data_zip" {
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
