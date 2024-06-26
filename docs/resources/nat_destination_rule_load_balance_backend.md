---
page_title: "vyos_nat_destination_rule_load_balance_backend Resource - vyos"

subcategory: "Nat"

description: |- 
  Network Address Translation (NAT) parameters⯯Destination NAT settings⯯Rule number for NAT⯯Apply NAT load balance⯯Translated IP address
---

# vyos_nat_destination_rule_load_balance_backend (Resource)
<center>

Network Address Translation (NAT) parameters  
⯯  
Destination NAT settings  
⯯  
Rule number for NAT  
⯯  
Apply NAT load balance  
⯯  
**Translated IP address**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `backend_id` (String) Translated IP address

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address to match  |
- `rule_id` (Number) Rule number for NAT

    &emsp;|Format    &emsp;|Description         |
    |------------|----------------------|
    &emsp;|1-999999  &emsp;|Number of NAT rule  |

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `weight` (Number) Set probability for this output value

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|1-100   &emsp;|Set probability for this output value  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
