output "bucket_name" {
  description = "Bucket name for our static hosting"
  value       = module.terrahouse.generated_bucket_name
}
