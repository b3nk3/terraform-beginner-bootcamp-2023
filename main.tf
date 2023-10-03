terraform {
  # cloud {
  #   organization = "bencodes"

  #   workspaces {
  #     name = "terra-house-3"
  #   }
  # }
}

module "terrahouse" {
  source              = "./modules/terrahouse"
  bucket_name         = var.bucket_name
  user_uuid           = var.user_uuid
  index_html_filepath = var.index_html_filepath
  error_html_filepath = var.error_html_filepath
  content_version     = var.content_version
}
