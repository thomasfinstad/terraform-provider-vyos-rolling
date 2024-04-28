# TODO Improve provider example
#  milestone: 3

terraform {
  required_providers {
    vyos = {
      source = "thomasfinstad/vyos-rolling"
    }
  }
}

variable "vyos_host" {
  type    = string
  default = "http://vyos.local"
}

variable "vyos_key" {
  type    = string
  default = "one two three four five"
}

provider "vyos" {
  endpoint = var.vyos_host
  api_key  = var.vyos_key

  certificate = {
    disable_verify = true
  }

  default_timeouts = 2
}
