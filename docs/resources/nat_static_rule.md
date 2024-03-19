---
page_title: "vyos_nat_static_rule Resource - terraform-provider-vyos"
subcategory: "nat"
description: |-
  Network Address Translation (NAT) parameters
  ⯯
  Static NAT (one-to-one)
  ⯯
  Rule number for NAT
---

# vyos_nat_static_rule (Resource)
<center>

Network Address Translation (NAT) parameters
⯯
Static NAT (one-to-one)
⯯
**Rule number for NAT**


</center>

## Schema

### Required

- `rule_id` (String) Rule number for NAT

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `destination` (Attributes) NAT destination parameters (see [below for nested schema](#nestedatt--destination))
- `inbound_interface` (String) Inbound interface of NAT traffic
- `translation` (Attributes) Translation address or prefix (see [below for nested schema](#nestedatt--translation))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--destination&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination`

Optional:

- `address` (String) IP address, prefix

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|ipv4     &emsp;|IPv4 address to match  |
    &emsp;|ipv4net  &emsp;|IPv4 prefix to match   |


&lt;a id=&#34;nestedatt--translation&#34;&gt;&lt;/a&gt;
### Nested Schema for `translation`

Optional:

- `address` (String) IP address, prefix

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|ipv4     &emsp;|IPv4 address to match  |
    &emsp;|ipv4net  &emsp;|IPv4 prefix to match   |  &emsp;|
