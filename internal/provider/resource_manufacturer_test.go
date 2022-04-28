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
						"nautobot_manufacturer.foo", "sample_attribute", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccResourceManufacturer = `
resource "nautobot_manufacturer" "foo" {
  sample_attribute = "bar"
}
`
