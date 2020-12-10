# christmas-tree

This resource writes a chirstmas tree to specified file.

## Example Usage

```hcl
resource "christmas-tree" "example" {
  path        = "/tmp/tree.txt"
  ball_color  = "red"
  star_color  = "yellow"
  light_color = "white"
}
```

## Argument References

* `path`: (Required) The PATH for provisioning christmas tree.
* `ball_color`: (Optional) The color of decorationg balls. Available colors: `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`, `default` (Default: `default`)
* `star_color`: (Optional) The color of decorationg stars. Available colors: `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`, `default` (Default: `default`)
* `light_color`: (Optional) The color of decorationg lights. Available colors: `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`, `default` (Default: `default`)

## Attributes References

* `path`: The PATH for provisioning christmas tree.
* `id`: `path` is stored as `id`.
* `ball_color`: The color of decorationg balls.
* `star_color`: The color of decorationg stars.
* `light_color`: The color of decorationg lights.
