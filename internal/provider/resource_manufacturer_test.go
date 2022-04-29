package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceManufacturer(t *testing.T) {
	t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceManufacturer,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"nautobot_manufacturer.test", "slug", regexp.MustCompile("^juniper$")),
				),
			},
		},
	})

}

const testAccResourceManufacturer = `
provider "nautobot" {
	url = "https://demo.nautobot.com/api/"
	token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  }

resource "nautobot_manufacturer" "test" {
	name = "Juniper"
}
`
