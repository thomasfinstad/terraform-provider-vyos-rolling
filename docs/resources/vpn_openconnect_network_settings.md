---
page_title: "vyos_vpn_openconnect_network_settings Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  vpn⯯SSL VPN OpenConnect, AnyConnect compatible server⯯Network settings
---

# vyos_vpn_openconnect_network_settings (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*vpn*  
⯯  
SSL VPN OpenConnect, AnyConnect compatible server  
⯯  
**Network settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `name_server` (List of String) Domain Name Servers (DNS) addresses

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|ipv4    &emsp;|Domain Name Server (DNS) IPv4 address  |
    &emsp;|ipv6    &emsp;|Domain Name Server (DNS) IPv6 address  |
- `push_route` (List of String) Route to be pushed to the client

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 network and prefix length  |
    &emsp;|ipv6net  &emsp;|IPv6 network and prefix length  |
- `split_dns` (List of String) Domains over which the provided DNS should be used

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|txt     &emsp;|Client prefix length  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `tunnel_all_dns` (String) If the tunnel-all-dns option is set to yes, tunnel all DNS queries via the VPN. This is the default when a default route is set.

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|yes     &emsp;|Enable tunneling of all DNS traffic   |
    &emsp;|no      &emsp;|Disable tunneling of all DNS traffic  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
