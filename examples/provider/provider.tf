# TODO Improve provider example

terraform {
  required_providers {
    vyos = {
      source = "github.com/thomasfinstad/vyos"
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

# resource "vyos_firewall_group_port_group" "example" {
#   port_group_id = "TF-Examples"

#   description = "Example of terraform created resource"
#   port        = [8080, "27015-27020", 443]
# }

# resource "vyos_firewall_zone" "example" {
#   zone_id = "TF-Examples"

#   intra_zone_filtering = {
#     action = "accept"
#     firewall = {
#       name = "test"
#     }
#   }
# }

resource "vyos_high_availability_vrrp_global_parameters" "example" {
  startup_delay = 12
}
