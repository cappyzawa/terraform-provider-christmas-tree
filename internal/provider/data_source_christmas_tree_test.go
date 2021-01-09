package provider

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceChristmasTree(t *testing.T) {
	f, _ := os.Open("../../testdata/tree.txt")
	defer f.Close()
	expectTree, _ := ioutil.ReadAll(f)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChristmasTree,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.christmas-tree.foo", "id", "../../testdata/tree.txt",
					),
					resource.TestCheckResourceAttr(
						"data.christmas-tree.foo", "path", "../../testdata/tree.txt",
					),
					resource.TestCheckResourceAttr(
						"data.christmas-tree.foo", "tree", string(expectTree),
					),
				),
			},
		},
	})
}

const (
	testAccDataSourceChristmasTree = `
data "christmas-tree" "foo" {
  path = "../../testdata/tree.txt"
}
  `
)
