resource "aws_cloudwatch_event_rule" "weather_collector_schedule" {
  name                = "weather-collector-schedule"
  schedule_expression = "rate(1 minute)"
}

resource "aws_cloudwatch_event_target" "weather_collector_target" {
  rule      = aws_cloudwatch_event_rule.weather_collector_schedule.name
  target_id = "weather-collector"
  arn       = aws_lambda_function.weather_collector.arn
}

resource "aws_lambda_permission" "allow_cloudwatch_to_invoke" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.weather_collector.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.weather_collector_schedule.arn
}

variable collector_lambda_function_name {
  type    = string
  default = "weather-collector"
}

resource "aws_lambda_function" "weather_collector" {
  function_name = var.collector_lambda_function_name
  handler       = "bootstrap"
  runtime       = "provided.al2023"
  architectures = ["x86_64"]
  role          = aws_iam_role.lambda_collector_exec.arn
  filename      = data.archive_file.lambda_collector_zip.output_path
  source_code_hash = data.archive_file.lambda_collector_zip.output_base64sha256
  timeout = 60 # Timeout in seconds

  # environment {
  #   variables = {
  #     AWS_LAMBDA_FUNCTION_NAME = var.collector_lambda_function_name
  #   }
  # }

  # ...other config...
  depends_on = [data.archive_file.lambda_collector_zip]
}
resource "aws_iam_role" "lambda_collector_exec" {
  name = "lambda_collector_exec_role"

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

resource "aws_iam_role_policy_attachment" "lambda_collector_logs" {
  role       = aws_iam_role.lambda_collector_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_cloudwatch_log_group" "weather_collector" {
  name              = "/aws/lambda/weather-collector"
  retention_in_days = 14
}

data "archive_file" "lambda_collector_zip" {
  type        = "zip"
  source_file = "${path.module}/../weather-collector/bin/bootstrap"
  output_path = "${path.module}/weather-collector.zip"
}

resource "aws_iam_policy" "lambda_collector_dynamodb_rw" {
  name        = "lambda_collector_dynamodb_rw"
  description = "Allow PutItem and GetItem on the weather table"
  policy      = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "dynamodb:PutItem",
          "dynamodb:GetItem"
        ]
        Resource = [
          aws_dynamodb_table.weather_samples.arn,
          aws_dynamodb_table.tide_times.arn,
          aws_dynamodb_table.weather.arn
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_collector_dynamodb_rw" {
  role       = aws_iam_role.lambda_collector_exec.name
  policy_arn = aws_iam_policy.lambda_collector_dynamodb_rw.arn
}

resource "aws_iam_policy" "lambda_collector_invoke_lambda" {
  name        = "lambda_collector_invoke_lambda"
  description = "Allow invoking Lambda functions"
  policy      = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "lambda:InvokeFunction"
        ]
        Resource = aws_lambda_function.weather_collector.arn
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_collector_invoke_lambda" {
  role       = aws_iam_role.lambda_collector_exec.name
  policy_arn = aws_iam_policy.lambda_collector_invoke_lambda.arn
}
