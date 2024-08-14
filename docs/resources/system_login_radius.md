---
page_title: "vyos_system_login_radius Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯System User Login Configuration⯯RADIUS based user authentication
---

# vyos_system_login_radius (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
System User Login Configuration  
⯯  
**RADIUS based user authentication**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `security_mode` (String) Security mode for RADIUS authentication

    |Format     &emsp;|Description                                                                  |
    |-------------|-------------------------------------------------------------------------------|
    |mandatory  &emsp;|Deny access immediately if RADIUS answers with Access-Reject                 |
    |optional   &emsp;|Pass to the next authentication method if RADIUS answers with Access-Reject  |
- `source_address` (List of String) Source IP address used to initiate connection

    |Format  &emsp;|Description          |
    |----------|-----------------------|
    |ipv4    &emsp;|IPv4 source address  |
    |ipv6    &emsp;|IPv6 source address  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (String) VRF instance name

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
