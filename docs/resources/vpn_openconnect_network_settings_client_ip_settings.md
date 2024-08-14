---
page_title: "vyos_vpn_openconnect_network_settings_client_ip_settings Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  vpn⯯SSL VPN OpenConnect, AnyConnect compatible server⯯Network settings⯯Client IP pools settings
---

# vyos_vpn_openconnect_network_settings_client_ip_settings (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*vpn*  
⯯  
SSL VPN OpenConnect, AnyConnect compatible server  
⯯  
Network settings  
⯯  
**Client IP pools settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `subnet` (String) Client IP subnet (CIDR notation)

    |Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    |ipv4net  &emsp;|IPv4 address and prefix length  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
