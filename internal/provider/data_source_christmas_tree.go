package provider

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChristmasTree() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceChristmasTreeRead,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tree": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceChristmasTreeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	p := d.Get("path").(string)
	f, err := os.Open(p)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
	}
	defer f.Close()
	tree, err := ioutil.ReadAll(f)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
	}
	d.SetId(p)
	d.Set("tree", string(tree))
	return diags
}
