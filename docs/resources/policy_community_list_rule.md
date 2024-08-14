---
page_title: "vyos_policy_community_list_rule Resource - vyos"

subcategory: "Policy"

description: |- 
  Routing policy⯯Add a BGP community list entry⯯Rule for this BGP community list
---

# vyos_policy_community_list_rule (Resource)
<center>

Routing policy  
⯯  
Add a BGP community list entry  
⯯  
**Rule for this BGP community list**


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
- `regex` (String) Regular expression to match against a community-list

    |Format        &emsp;|Description                                                  |
    |----------------|---------------------------------------------------------------|
    |&lt;aa:nn&gt;       &emsp;|Community number in AA:NN format                             |
    |local-AS      &emsp;|Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03  |
    |no-advertise  &emsp;|Well-known communities value NO_ADVERTISE 0xFFFFFF02         |
    |no-export     &emsp;|Well-known communities value NO_EXPORT 0xFFFFFF01            |
    |internet      &emsp;|Well-known communities value 0                               |
    |additive      &emsp;|New value is appended to the existing value                  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `community_list` (String) Add a BGP community list entry

    |Format  &emsp;|Description              |
    |----------|---------------------------|
    |txt     &emsp;|BGP community-list name  |
- `rule` (Number) Rule for this BGP community list

    |Format   &emsp;|Description                 |
    |-----------|------------------------------|
    |1-65535  &emsp;|Community-list rule number  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
