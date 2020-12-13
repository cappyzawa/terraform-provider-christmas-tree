---
page_title: "christmas-tree Resource - terraform-provider-christmas-tree"
subcategory: ""
description: |-
  This resource provides virtual christmas tree to your machine.
---

# Resource `christmas-tree`

This resource provides virtual christmas tree to your machine.

## Example Usage

```terraform
resource "christmas-tree" "example" {
  path        = "/tmp/example"
  ball_color  = "red"
  star_color  = "yellow"
  light_color = "white"
}
```

## Schema

### Required

- **path** (String, Required)

### Optional

- **ball_color** (String, Optional) available colors: "black", "red", "green", "yellow", "blue", "magenta", "cyan", "white", "default"
- **id** (String, Optional) The ID of this resource.
- **light_color** (String, Optional) available colors: "black", "red", "green", "yellow", "blue", "magenta", "cyan", "white", "default"
- **star_color** (String, Optional) available colors: "black", "red", "green", "yellow", "blue", "magenta", "cyan", "white", "default"


