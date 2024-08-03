---
page_title: "vyos_service_ssh_dynamic_protection Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  System services⯯Secure Shell (SSH)⯯Allow dynamic protection
---

# vyos_service_ssh_dynamic_protection (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

System services  
⯯  
Secure Shell (SSH)  
⯯  
**Allow dynamic protection**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `allow_from` (List of String) Always allow inbound connections from these systems

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv4     &emsp;|Address to match against        |
    &emsp;|ipv4net  &emsp;|IPv4 address and prefix length  |
    &emsp;|ipv6     &emsp;|IPv6 address to match against   |
    &emsp;|ipv6net  &emsp;|IPv6 address and prefix length  |
- `block_time` (Number) Block source IP in seconds. Subsequent blocks increase by a factor of 1.5

    &emsp;|Format   &emsp;|Description                            |
    |-----------|-----------------------------------------|
    &emsp;|1-65535  &emsp;|Time interval in seconds for blocking  |
- `detect_time` (Number) Remember source IP in seconds before reset their score

    &emsp;|Format   &emsp;|Description               |
    |-----------|----------------------------|
    &emsp;|1-65535  &emsp;|Time interval in seconds  |
- `threshold` (Number) Block source IP when their cumulative attack score exceeds threshold

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Threshold score  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
