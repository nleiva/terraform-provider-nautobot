# Terraform Provider Nautobot 

Nautobot provider created for educational purposes. You can fork it for long term development :-)

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.17

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the `make` command: 
```sh
$ make install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider


The provide takes two arguments, `url` and `token`. For the data sources and resources supported, take a look at the [internal/provider](internal/provider) folder. In the next example, we capture the data of all manufacturers and create a new manufacturer "Vendor I".


```hcl
terraform {
  required_providers {
    nautobot = {
      version = "0.1"
      source  = "github.com/nleiva/nautobot"
    }
  }
}

provider "nautobot" {
  url = "https://demo.nautobot.com/api/"
  token = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}

data "nautobot_manufacturers" "all" {}

resource "nautobot_manufacturer" "new" {
  description = "Created with Terraform"
  name    = "Vendor I"
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `make install`.

To generate or update documentation, run `go generate ./...`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
