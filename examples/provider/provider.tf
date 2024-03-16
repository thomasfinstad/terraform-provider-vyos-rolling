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

/* resource "vyos_firewall_group_port_group" "example" {
  port_group_id = "TF-Examples"

  description = "Example of terraform created resource ${plantimestamp()}"
  port        = [8080, "27015-27020", 44]
}

resource "vyos_firewall_ipv4_name" "example" {
  name_id = "TF-Example"

  default_action = "accept"
  description    = "Another terraform t ${plantimestamp()}"
} */

/* resource "vyos_firewall_ipv4_name" "example2" {
  count = 1

  name_id = "TF-Example2-${count.index}-${replace(plantimestamp(), ":", "-")}"

  default_action = "accept"
  description    = "Another terraform test"
}

resource "vyos_firewall_ipv4_name_rule" "example2" {
  count = length(vyos_firewall_ipv4_name.example2)

  name_id = vyos_firewall_ipv4_name.example2[count.index].name_id
  rule_id = 99

  action = "accept"

  icmp = {
    code = 0
    type = 0
  }
}
 */

# resource "vyos_policy_access_list" "name" {
#   access_list_id = 2

# }

resource "vyos_policy_access_list_rule" "name" {
  access_list_id = 2
  rule_id        = 5

  description = plantimestamp()
  action      = "permit"
  source = {
    host = "55.55.55.55"
  }
}
