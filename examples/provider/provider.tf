# TODO Improve provider example

terraform {
  required_providers {
    vyos = {
      source = "ex.c/thomasfinstad/vyos"
    }
  }
}

provider "vyos" {
  endpoint = "https://192.168.2.252"
  api_key  = "FoFeMcws4XpbKk4TDceQWzUdDsHptAr7FcuApxiHUV"
}

resource "vyos_firewall_name" "test" {
  identifier = "new_provider"
  rule = {
    33 = {
      action   = "accept"
      protocol = "udp"
    }
  }
}
