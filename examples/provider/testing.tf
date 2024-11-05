// There are in general 2 types of resources
//
// Named: can have multiple different resources of the same kind, contains an "identifier" attribute.
// Global: can have one resource of this kind, having multiple will result in configuration conflicts between the two.

// Named resource
resource "vyos_system_syslog_global_facility" "all" {
  identifier = { facility = "all" }
  level      = "info"
}

// Global resource with multi value list
resource "vyos_service_dns_forwarding" "this" {
  allow_from = [
    "192.168.0.0/16",
  ]
  ignore_hosts_file = false
  listen_address = [
    "192.168.2.253",
    "192.168.10.253",
    "192.168.30.253",
    "192.168.20.253",
    "192.168.240.253",
    "192.168.249.253",
  ]
  no_serve_rfc1918 = false
  system           = false
}

// Named child resource
# resource "vyos_policy_access_list_rule" "this" {
#   identifier = {
#     access_list = vyos_policy_access_list.this.identifier.access_list
#     rule        = 69
#   }

#   description = tonumber(regex("[0-9]{2}", md5(plantimestamp())))

#   action = "permit"
#   source = {
#     host = "55.55.55.55"
#   }
# }

// Global resource
resource "vyos_system_conntrack_tcp" "this" {
  half_open_connections = tonumber(regex("[0-9]{2}", md5(abspath(path.root)))) + 1
  loose                 = "enable"
}

// Global resource with an merged TagNode child
resource "vyos_service_ntp" "this" {
  server = {
    "no.pool.ntp.org" = {
      pool   = true
      prefer = true
    }
    "pool.ntp.org" = {
      pool   = true
      prefer = false
    }
    "time1.vyos.net" = {
      pool   = true
      prefer = false
    }
  }
}

//////
// Regression test for: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/issues/224
# resource "vyos_interfaces_dummy" "issue224" {
#   identifier = {
#     dummy = "dumb224"
#   }
#   address = ["10.0.224.1/24"]
# }

# resource "vyos_service_dhcp_server_shared_network_name" "issue224" {
#   identifier = {
#     shared_network_name = "LAN"
#   }
#   subnet = { for i in range(length(vyos_interfaces_dummy.issue224.address)) : vyos_interfaces_dummy.issue224.address[i] => { subnet_id = 224 + i } }
# }

# resource "vyos_service_dhcp_server_shared_network_name_subnet_range" "issue224" {
#   for_each = vyos_service_dhcp_server_shared_network_name.issue224.subnet

#   identifier = {
#     shared_network_name = vyos_service_dhcp_server_shared_network_name.issue224.identifier.shared_network_name
#     subnet              = each.key
#     range               = "224-${each.value.subnet_id}"
#   }
#   start = cidrhost(each.key, 100)
#   stop  = cidrhost(each.key, 150)
# }
