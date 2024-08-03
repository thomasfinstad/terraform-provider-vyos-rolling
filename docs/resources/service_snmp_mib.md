---
page_title: "vyos_service_snmp_mib Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Simple Network Management Protocol (SNMP)⯯Management information base (MIB)
---

# vyos_service_snmp_mib (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Simple Network Management Protocol (SNMP)  
⯯  
**Management information base (MIB)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `interface` (List of String) Sets the interface name prefix to include in the IF-MIB data collection

    &emsp;|Format  &emsp;|Description                              |
    |----------|-------------------------------------------|
    &emsp;|br      &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|bond    &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|dum     &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|eth     &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|gnv     &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|macsec  &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|peth    &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|sstpc   &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|tun     &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|veth    &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|vti     &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|vtun    &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|vxlan   &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|wg      &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|wlan    &emsp;|Allow prefix for IF-MIB data collection  |
    &emsp;|wwan    &emsp;|Allow prefix for IF-MIB data collection  |
- `interface_max` (Number) Sets the maximum number of interfaces included in IF-MIB data collection

    &emsp;|Format        &emsp;|Description                                                               |
    |----------------|----------------------------------------------------------------------------|
    &emsp;|1-4294967295  &emsp;|Sets the maximum number of interfaces included in IF-MIB data collection  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
