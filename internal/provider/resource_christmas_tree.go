package provider

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// StarChar describes a character of star
	StarChar = "*"
	// BallChar describes a character of ball
	BallChar = "o"
	// LightChar describes a character of light
	LightChar = "'"
	// LeafChar describes a character of leaf
	LeafChar = "."

	colorDescription = `available colors: "black", "red", "green", "yellow", "blue", "magenta", "cyan", "white", "default"`
)

func resourceChristmasTree() *schema.Resource {
	return &schema.Resource{
		Description: "This resource provides virtual christmas tree to your machine.",

		CreateContext: resourceChristmasTreeCreate,
		ReadContext:   resourceChristmasTreeRead,
		UpdateContext: resourceChristmasTreeUpdate,
		DeleteContext: resourceChristmasTreeDelete,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ball_color": {
				Type:        schema.TypeString,
				Default:     "default",
				Description: colorDescription,
				Optional:    true,
			},
			"star_color": {
				Type:        schema.TypeString,
				Default:     "default",
				Description: colorDescription,
				Optional:    true,
			},
			"light_color": {
				Type:        schema.TypeString,
				Default:     "default",
				Description: colorDescription,
				Optional:    true,
			},
		},
	}
}

func resourceChristmasTreeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return createOrUpdate(ctx, d)
}

func resourceChristmasTreeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	p := d.Get("path").(string)
	f, err := os.Open(p)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
	}
	defer f.Close()
	d.SetId(p)

	return diags
}

func resourceChristmasTreeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if d.HasChange("path") {
		before, _ := d.GetChange("path")
		if err := os.RemoveAll(before.(string)); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Error(),
			})
		}
	}
	return append(diags, createOrUpdate(ctx, d)...)
}

func resourceChristmasTreeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	p := d.Get("path").(string)
	if err := os.RemoveAll(p); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
	}
	d.SetId("")
	return diags
}

func createOrUpdate(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics

	p := d.Get("path").(string)
	starColor := d.Get("star_color").(string)
	ballColor := d.Get("ball_color").(string)
	lightColor := d.Get("light_color").(string)

	m := map[string]string{
		StarChar:  starColor,
		BallChar:  ballColor,
		LightChar: lightColor,
		LeafChar:  "green",
	}

	tree := ChristmasTreeText

	var err error
	for char, color := range m {
		tree, err = Colorize(tree, char, color)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Error(),
			})
		}
	}

	f, err := os.Create(p)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
		})
	}
	fmt.Fprint(f, tree)
	d.SetId(p)
	return diags
}
