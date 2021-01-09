package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceChristmasTree(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:                    testAccDefaultResourceChristmasTree,
				PreventPostDestroyRefresh: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"christmas-tree.foo", "id", "/tmp/test-tree",
					),
					resource.TestCheckResourceAttr(
						"christmas-tree.foo", "star_color", "default",
					),
				),
			},
			{
				Config: testAccYellowStarResourceChristmasTree,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"christmas-tree.foo", "id", "/tmp/test-tree",
					),
					resource.TestCheckResourceAttr(
						"christmas-tree.foo", "star_color", "yellow",
					),
				),
			},
		},
	})
}

const (
	testAccDefaultResourceChristmasTree = `
resource "christmas-tree" "foo" {
  path = "/tmp/test-tree"
}
  `
	testAccYellowStarResourceChristmasTree = `
resource "christmas-tree" "foo" {
  path = "/tmp/test-tree"
  star_color = "yellow"
}
  `
)
