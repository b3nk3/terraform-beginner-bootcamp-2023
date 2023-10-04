terraform {
  required_providers {
    terratowns = {
      source  = "local.providers/local/terratowns"
      version = "1.0.0"
    }
  }
  # cloud {
  #   organization = "bencodes"

  #   workspaces {
  #     name = "terra-house-3"
  #   }
  # }
}

provider "terratowns" {
  # mock deets
  # endpoint  = "http://localhost:4567/api/"
  # user_uuid = "e328f4ab-b99f-421c-84c9-4ccea042c7d1"
  # token     = "9b49b3fb-b8e9-483c-b703-97ba88eef8e0"

  # prod deets
  endpoint  = var.terratowns_endpoint
  user_uuid = var.teacherseat_user_uuid
  token     = var.terratowns_access_token
}

module "terrahome_aws" {
  source              = "./modules/terrahouse"
  user_uuid           = var.teacherseat_user_uuid
  index_html_filepath = var.index_html_filepath
  error_html_filepath = var.error_html_filepath
  content_version     = var.content_version
  assets_filepath     = var.assets_filepath
}

resource "terratowns_home" "home" {
  name            = "Diablo den"
  description     = "Diablo is my favourite game"
  domain_name     = module.terrahome_aws.cloudfront_url
  town            = "missingo"
  content_version = 1

}
