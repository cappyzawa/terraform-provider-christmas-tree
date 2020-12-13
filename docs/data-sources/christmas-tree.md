# christmas-tree

This data source gets a christmas tree based on file path.

## Example Usage

```hcl
data "christmas-tree" "example" {
  path = "/tmp/tree.txt"
}

output "example" {
  value = file(data.christmas-tree.example.path)
}
```

## Argument References

* `path`: (Required) The PATH for getting christmas tree.

## Attributes References

* `path`: The PATH for getting a christmas tree.
* `tree`: A string representing a christmas tree.
