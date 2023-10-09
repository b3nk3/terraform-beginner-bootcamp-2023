output "terrahomes" {
  description = "A list of Terra Homes and their details (url, bucket name, bucket url)"
  value       = { for terrahome, th_outputs in module.terrahome_aws : terrahome => th_outputs }
}
