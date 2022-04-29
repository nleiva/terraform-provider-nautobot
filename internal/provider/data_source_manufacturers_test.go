package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceManufacturers(t *testing.T) {
	// https://github.com/hashicorp/terraform-plugin-sdk/issues/952
	t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManufacturer,
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckResourceAttr("data.nautobot_manufacturers.list.manufacturers", "", ""),
					resource.TestCheckOutput("vendor", "juniper"),
				),
			},
		},
	})
}

const testAccDataSourceManufacturer = `
provider "nautobot" {
	url = "https://demo.nautobot.com/api/"
	token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  }

data "nautobot_manufacturers" "list" {}

variable "filter" {
	type    = string
	default = "Juniper"
}

output "vendor" {
	value = [
	  for manufacturer in data.nautobot_manufacturers.list.manufacturers :
	  manufacturer.slug
	  if manufacturer.name == var.filter
	]
}
`
