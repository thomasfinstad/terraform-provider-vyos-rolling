# TODO Improve provider example
#  milestone: 3

terraform {
  required_version = "~> 1.0"
  required_providers {
    vyos = {
      source = "thomasfinstad/vyos-rolling"
    }
  }
}

variable "vyos_host" {
  type    = string
  default = "https://vyos.local"
}

variable "vyos_key" {
  type    = string
  default = "MySuperSecretKey"
}

provider "vyos" {
  endpoint = var.vyos_host
  api_key  = var.vyos_key

  certificate = {
    disable_verify = true
  }

  default_timeouts = 1
}
