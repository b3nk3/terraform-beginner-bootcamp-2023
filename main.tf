terraform {
  cloud {
    organization = "bencodes"

    workspaces {
      name = "terra-house-3"
    }
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.5.1"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region = "ca-central-1"
}
provider "random" {
  # Configuration options
}

resource "random_string" "bucket_name" {
  length  = 32
  lower   = true
  upper   = false
  special = false
}

resource "aws_s3_bucket" "example" {
  bucket = random_string.bucket_name.result

}

output "generated_bucket_name" {
  value = random_string.bucket_name.result
}
