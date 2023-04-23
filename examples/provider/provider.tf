# TODO Improve provider example

terraform {
  required_providers {
    vyos = {
      source = "ex.c/thomasfinstad/vyos"
    }
  }
}

provider "vyos" {
  // TODO replace dev instance endpoint with generic 192.168.0.1 address
  endpoint = "https://192.168.2.252"

  // TODO replace dev instance api key with elipsis
  api_key = "FoFeMcws4XpbKk4TDceQWzUdDsHptAr7FcuApxiHUV"

  certificate = {
    disable_verify = true
  }
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
