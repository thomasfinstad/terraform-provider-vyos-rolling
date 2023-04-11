# TODO Improve provider example

terraform {
  required_providers {
    vyos = {
      source = "ex.c/thomasfinstad/vyos"
    }
  }
}

provider "vyos" {}

resource "vyos_example" "name" {
  tag_node = {
    name = {
      enable = false
      ak     = "a"
    }
  }
}
