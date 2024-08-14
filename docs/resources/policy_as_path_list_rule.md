---
page_title: "vyos_policy_as_path_list_rule Resource - vyos"

subcategory: "Policy"

description: |- 
  Routing policy⯯Add a BGP autonomous system path filter⯯Rule for this as-path-list
---

# vyos_policy_as_path_list_rule (Resource)
<center>

Routing policy  
⯯  
Add a BGP autonomous system path filter  
⯯  
**Rule for this as-path-list**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `action` (String) Action to take on entries matching this rule

    |Format  &emsp;|Description              |
    |----------|---------------------------|
    |permit  &emsp;|Permit matching entries  |
    |deny    &emsp;|Deny matching entries    |
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `regex` (String) Regular expression to match against an AS path

    |Format  &emsp;|Description                                     |
    |----------|--------------------------------------------------|
    |txt     &emsp;|AS path regular expression (ex: &#34;64501 64502&#34;)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `as_path_list` (String) Add a BGP autonomous system path filter

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|AS path list name  |
- `rule` (Number) Rule for this as-path-list

    |Format   &emsp;|Description               |
    |-----------|----------------------------|
    |1-65535  &emsp;|AS path list rule number  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
