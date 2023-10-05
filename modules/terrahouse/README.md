# Terraform Module: `terrahome_aws`

This Terraform module manages the deployment of AWS resources for Terraform-based websites created by TerraHouse. It allows you to deploy multiple TerraHouse websites, each with its own configuration.

## Usage

```hcl
module "terrahome_aws" {
  for_each        = var.terrahomes
  source          = "./modules/terrahouse"
  user_uuid       = var.teacherseat_user_uuid
  content_version = each.value.content_version

  index_html_filepath = "${var.public_path}/${each.key}/index.html"
  error_html_filepath = "${var.public_path}/${each.key}/error.html"
  assets_filepath     = "${var.public_path}/${each.key}/assets/"
}
```

Inputs
for_each (required): A map of TerraHouse configurations. Each entry in this map represents a different TerraHouse website to be deployed.

source (required): The path to the module source code.

user_uuid (required): The UUID of the user deploying TerraHouse.

content_version (required): The version of TerraHouse content to deploy.

index_html_filepath (required): The file path to the index.html file for the website. It is constructed based on the public_path and the key from the for_each loop.

error_html_filepath (required): The file path to the error.html file for the website. It is constructed based on the public_path and the key from the for_each loop.

assets_filepath (required): The file path to the assets directory for the website. It is constructed based on the public_path and the key from the for_each loop.
