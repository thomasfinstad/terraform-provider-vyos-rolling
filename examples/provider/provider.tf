# TODO Improve provider example

terraform {
  required_providers {
    vyos = {
      source = "local/providers/vyos"
    }
  }
}

provider "vyos" {
  // TODO replace dev instance endpoint with var
  endpoint = "https://192.168.2.252"

  // TODO replace dev instance api key with var
  api_key = "FoFeMcws4XpbKk4TDceQWzUdDsHptAr7FcuApxiHUV"

  certificate = {
    disable_verify = true
  }
}

resource "vyos_firewall_group_port_group" "example" {
  port_group_id = "TF-Examples"

  description = "Example of terraform created resource ${timestamp()}"
  port        = [8080, "27015-27020", 44]
}

resource "vyos_high_availability_vrrp_global_parameters" "example" {
  startup_delay = 13
}

resource "vyos_firewall_ipv4_name" "example" {
  name_id = "TF-Example"

  default_action = "accept"
  description    = "Another terraform t ${timestamp()}"
}
