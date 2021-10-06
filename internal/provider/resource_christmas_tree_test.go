package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
				ResourceName:            "christmas-tree.foo",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"light_color", "ball_color", "star_color"},
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) != 1 {
						t.Errorf("length of state should be 1, but it is %d", len(s))
					}
					if _, ok := s[0].Attributes["path"]; !ok {
						t.Error("path should exist as an attribute")
					}
					return nil
				},
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
