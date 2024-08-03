---
page_title: "vyos_service_mdns_repeater Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Multicast DNS (mDNS) parameters⯯mDNS repeater configuration
---

# vyos_service_mdns_repeater (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Multicast DNS (mDNS) parameters  
⯯  
**mDNS repeater configuration**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `allow_service` (List of String) Allowed mDNS services to be repeated

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|txt     &emsp;|mDNS service  |
- `browse_domain` (List of String) mDNS browsing domains in addition to the default one

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|txt     &emsp;|mDNS browsing domain  |
- `disable` (Boolean) Disable instance
- `interface` (List of String) Interface to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `ip_version` (String) IP address version to use

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|_ipv4   &emsp;|Use only IPv4 address           |
    &emsp;|_ipv6   &emsp;|Use only IPv6 address           |
    &emsp;|both    &emsp;|Use both IPv4 and IPv6 address  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrrp_disable` (Boolean) Disables mDNS repeater on VRRP interfaces not in MASTER state

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
