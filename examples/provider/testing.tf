
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

# resource "vyos_firewall_ipv4_name" "example2" {
#   count = 1

#   name_id = "TF-Example2-${count.index}-${replace(plantimestamp(), ":", "-")}"

#   default_action = "accept"
#   description    = "Another terraform test"
# }

# resource "vyos_firewall_ipv4_name_rule" "example2" {
#   count = length(vyos_firewall_ipv4_name.example2)

#   name_id = vyos_firewall_ipv4_name.example2[count.index].name_id
#   rule_id = 99

#   action = "accept"

#   icmp = {
#     code = 0
#     type = 0
#   }
# }


// Empty named resource
resource "vyos_policy_access_list" "name" {
  access_list_id = 42
}

// Child of empty
resource "vyos_policy_access_list_rule" "name" {
  access_list_id = 42
  rule_id        = 69

  description = plantimestamp()
  action      = "permit"
  source = {
    host = "55.55.55.55"
  }
}

# vyos_policy_access_list_rule.name: Modifications complete after 10s [id=policy__access-list__42__rule__69]
