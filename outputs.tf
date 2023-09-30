output "bucket_name" {
  description = "Bucket name for our static hosting"
  value       = module.terrahouse.generated_bucket_name
}
output "s3_website_endpoint" {
  description = "Endpoint of our S3 static hosting"
  value       = module.terrahouse.website_endpoint
}
