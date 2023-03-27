terraform {
  required_providers {
    nautobot = {
      version = "0.3.0"
      source  = "github.com/nleiva/nautobot"
    }
  }
}

provider "nautobot" {
  url = "https://demo.nautobot.com/api/"
  token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}

resource "nautobot_manufacturer" "new" {
  description = "Created with Terraform"
  name    = "New Vendor"
}