---
page_title: "vyos_policy_extcommunity_list_rule Resource - vyos"

subcategory: "Policy"

description: |- 
  Routing policy⯯Add a BGP extended community list entry⯯Rule for this BGP extended community list
---

# vyos_policy_extcommunity_list_rule (Resource)
<center>

Routing policy  
⯯  
Add a BGP extended community list entry  
⯯  
**Rule for this BGP extended community list**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `action` (String) Action to take on entries matching this rule

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|permit  &emsp;|Permit matching entries  |
    &emsp;|deny    &emsp;|Deny matching entries    |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `regex` (String) Regular expression to match against an extended community list

    &emsp;|Format          &emsp;|Description                                 |
    |------------------|----------------------------------------------|
    &emsp;|&lt;aa:nn:nn&gt;      &emsp;|Extended community list regular expression  |
    &emsp;|&lt;rt aa:nn:nn&gt;   &emsp;|Route Target regular expression             |
    &emsp;|&lt;soo aa:nn:nn&gt;  &emsp;|Site of Origin regular expression           |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `extcommunity_list` (String) Add a BGP extended community list entry

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|txt     &emsp;|BGP extended community-list name  |
- `rule` (Number) Rule for this BGP extended community list

    &emsp;|Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    &emsp;|1-65535  &emsp;|Extended community-list rule number  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
