---
page_title: "vyos_protocols_ospf_parameters Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Open Shortest Path First (OSPF)⯯OSPF specific parameters
---

# vyos_protocols_ospf_parameters (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Open Shortest Path First (OSPF)  
⯯  
**OSPF specific parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `abr_type` (String) OSPF ABR type

    &emsp;|Format    &emsp;|Description        |
    |------------|---------------------|
    &emsp;|cisco     &emsp;|Cisco ABR type     |
    &emsp;|ibm       &emsp;|IBM ABR type       |
    &emsp;|shortcut  &emsp;|Shortcut ABR type  |
    &emsp;|standard  &emsp;|Standard ABR type  |
- `opaque_lsa` (Boolean) Enable the Opaque-LSA capability (rfc2370)
- `rfc1583_compatibility` (Boolean) Enable RFC1583 criteria for handling AS external routes
- `router_id` (String) Override default router identifier

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv4    &emsp;|Router-ID in IP address format  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
