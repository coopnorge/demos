# https://www.terraform.io/docs/language/settings/index.html
# https://www.terraform.io/docs/language/expressions/version-constraints.html
terraform {
  required_providers {
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "~> 4.32"
    }
    google = {
      source  = "hashicorp/google"
      version = "~> 4.32"
    }
  }
  required_version = ">= 1.0.0, < 2.0.0"
}


