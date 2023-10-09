terraform {
  required_providers {
    terratowns = {
      source  = "local.providers/local/terratowns"
      version = "1.0.0"
    }
  }
  cloud {
    organization = "bencodes"

    workspaces {
      name = "bens-terra-homes"
    }
  }
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
  for_each        = var.terrahomes
  source          = "./modules/terrahouse"
  user_uuid       = var.teacherseat_user_uuid
  content_version = each.value.content_version

  index_html_filepath = "${var.public_path}/${each.key}/index.html"
  error_html_filepath = "${var.public_path}/${each.key}/error.html"
  assets_filepath     = "${var.public_path}/${each.key}/assets/"
}

resource "terratowns_home" "home" {
  for_each        = var.terrahomes
  name            = each.value.name
  description     = each.value.description
  domain_name     = module.terrahome_aws[each.key].cloudfront_url
  town            = each.value.town
  content_version = each.value.content_version

}
