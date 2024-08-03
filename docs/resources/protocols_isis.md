---
page_title: "vyos_protocols_isis Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Intermediate System to Intermediate System (IS-IS)
---

# vyos_protocols_isis (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
**Intermediate System to Intermediate System (IS-IS)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `advertise_high_metrics` (Boolean) Advertise high metric value on all interfaces
- `advertise_passive_only` (Boolean) Advertise prefixes of passive interfaces only
- `dynamic_hostname` (Boolean) Dynamic hostname for IS-IS
- `level` (String) IS-IS level number

    &emsp;|Format     &emsp;|Description                               |
    |-------------|--------------------------------------------|
    &emsp;|level-1    &emsp;|Act as a station router                   |
    &emsp;|level-1-2  &emsp;|Act as both a station and an area router  |
    &emsp;|level-2    &emsp;|Act as an area router                     |
- `log_adjacency_changes` (Boolean) Log adjacency state changes
- `lsp_gen_interval` (Number) Minimum interval between regenerating same LSP

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|1-120   &emsp;|Minimum interval in seconds  |
- `lsp_mtu` (Number) Configure the maximum size of generated LSPs

    &emsp;|Format    &emsp;|Description                     |
    |------------|----------------------------------|
    &emsp;|128-4352  &emsp;|Maximum size of generated LSPs  |
- `lsp_refresh_interval` (Number) LSP refresh interval

    &emsp;|Format   &emsp;|Description                      |
    |-----------|-----------------------------------|
    &emsp;|1-65235  &emsp;|LSP refresh interval in seconds  |
- `max_lsp_lifetime` (Number) Maximum LSP lifetime

    &emsp;|Format     &emsp;|Description              |
    |-------------|---------------------------|
    &emsp;|350-65535  &emsp;|LSP lifetime in seconds  |
- `metric_style` (String) Use old-style (ISO 10589) or new-style packet formats

    &emsp;|Format      &emsp;|Description                                            |
    |--------------|---------------------------------------------------------|
    &emsp;|narrow      &emsp;|Use old style of TLVs with narrow metric               |
    &emsp;|transition  &emsp;|Send and accept both styles of TLVs during transition  |
    &emsp;|wide        &emsp;|Use new style of TLVs to carry wider metric            |
- `net` (String) A Network Entity Title for this process (ISO only)

    &emsp;|Format                &emsp;|Description                 |
    |------------------------|------------------------------|
    &emsp;|XX.XXXX. ... .XXX.XX  &emsp;|Network entity title (NET)  |
- `purge_originator` (Boolean) Use the RFC 6232 purge-originator
- `set_attached_bit` (Boolean) Set attached bit to identify as L1/L2 router for inter-area traffic
- `set_overload_bit` (Boolean) Set overload bit to avoid any transit traffic
- `spf_interval` (Number) Minimum interval between SPF calculations

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|1-120   &emsp;|Interval in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `topology` (String) Configure IS-IS topologies

    &emsp;|Format          &emsp;|Description                   |
    |------------------|--------------------------------|
    &emsp;|ipv4-multicast  &emsp;|Use IPv4 multicast topology   |
    &emsp;|ipv4-mgmt       &emsp;|Use IPv4 management topology  |
    &emsp;|ipv6-unicast    &emsp;|Use IPv6 unicast topology     |
    &emsp;|ipv6-multicast  &emsp;|Use IPv6 multicast topology   |
    &emsp;|ipv6-mgmt       &emsp;|Use IPv6 management topology  |
    &emsp;|ipv6-dstsrc     &emsp;|Use IPv6 dst-src topology     |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
