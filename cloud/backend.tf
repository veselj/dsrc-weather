// Create resources with backend config commented out.
// Apply to create S3 and DynamoDB.
// Uncomment/configure the backend.
// Run terraform init to migrate state.

terraform {
  backend "s3" {
    bucket         = "dsrc-weather-tf-state"
    key            = "terraform/terraform.tfstate"
    region         = "eu-west-1"
    dynamodb_table = "dsrc-weather-tf-lock-table"
    encrypt        = true
  }
}