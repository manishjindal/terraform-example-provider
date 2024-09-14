terraform {
  required_providers {
    example = {
      source  = "terraform-example.com/exampleprovider/example"
    #   source  = "example"
      version = "~> 1.0.0"
    }
  }
}


provider "example" {}

# resource "example_example" "example" {
#   name = "example"
# }


resource "example_hello" "my-server" {
    address = "ssdsad"
}

resource "example_server" "my-server" {
    address = "ssdsad"
}