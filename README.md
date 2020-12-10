Terraform Provider Christmas Tree
==================

Using the provider
----------------------

[cappyzawa/christmas\-tree \| Terraform Registry](https://registry.terraform.io/providers/cappyzawa/christmas-tree/latest)

Developing the Provider
---------------------------

To compile the provider, run `make install`. This will build the provider and put the provider binary in the `$HOME/.terraform.d/plugins` directory.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
