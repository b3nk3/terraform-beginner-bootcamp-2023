terraform {
  # cloud {
  #   organization = "bencodes"

  #   workspaces {
  #     name = "terra-house-3"
  #   }
  # }
}

module "terrahouse" {
  source      = "./modules/terrahouse"
  bucket_name = var.bucket_name
  user_uuid   = var.user_uuid
}
