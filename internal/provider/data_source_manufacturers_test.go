package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceManufacturers(t *testing.T) {
	t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceManufacturer,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.nautobot_manufacturers.juniper", "name", "Juniper"),
				),
			},
			{
				Config: testAccDataSourceManufacturer,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.nautobot_manufacturers.juniper", "id", regexp.MustCompile("^4873d752")),
				),
			},
		},
	})
}

const testAccDataSourceManufacturer = `
resource "nautobot_manufacturer" "juniper" {
	id = "4873d752-5dbe-4006-8345-8279a0dfbbda"
	url = "https://develop.demo.nautobot.com/api/dcim/manufacturers/4873d752-5dbe-4006-8345-8279a0dfbbda/"
	name = "Juniper"
	slug = "juniper"
	description = "Juniper Networks"
	devicetype_count = 0
	platform_count = 1
	custom_fields = "{}"
	created = "2022-03-08"
	last_updated = "2022-03-08T14:50:48.492203Z"
	display = "Juniper"
}
`
