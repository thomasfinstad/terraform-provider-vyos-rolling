// There are in general 2 types of resources
//
// Named: can have multiple different resources of the same kind, contains an "identifier" attribute.
// Global: can have one resource of this kind, having multiple will result in configuration conflicts between the two.

// Empty named parent resource
resource "vyos_policy_access_list" "this" {
  identifier = {
    access_list = tonumber(regex("[0-9]{2}", md5(plantimestamp()))) + 1
  }
}

// Named child resource
resource "vyos_policy_access_list_rule" "this" {
  identifier = {
    access_list = vyos_policy_access_list.this.identifier.access_list
    rule        = 69
  }

  description = tonumber(regex("[0-9]{2}", md5(plantimestamp())))

  action = "permit"
  source = {
    host = "55.55.55.55"
  }
}

// Global resource
resource "vyos_system_conntrack_tcp" "this" {
  half_open_connections = tonumber(regex("[0-9]{2}", md5(plantimestamp()))) + 1
  loose                 = "enable"
}
