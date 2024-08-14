---
page_title: "vyos_service_webproxy_url_filtering_squidguard_source_group Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Webproxy service settings⯯URL filtering settings⯯URL filtering via squidGuard redirector⯯Source group name
---

# vyos_service_webproxy_url_filtering_squidguard_source_group (Resource)
<center>

*service*  
⯯  
Webproxy service settings  
⯯  
URL filtering settings  
⯯  
URL filtering via squidGuard redirector  
⯯  
**Source group name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (List of String) Address for source-group

    |Format     &emsp;|Description                  |
    |-------------|-------------------------------|
    |ipv4       &emsp;|IPv4 address to match        |
    |ipv4net    &emsp;|IPv4 prefix to match         |
    |ipv4range  &emsp;|IPv4 address range to match  |
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `domain` (List of String) Domain for source-group

    |Format  &emsp;|Description                       |
    |----------|------------------------------------|
    |domain  &emsp;|Domain name for the source-group  |
- `ldap_ip_search` (List of String) LDAP search expression for an IP address list
- `ldap_user_search` (List of String) LDAP search expression for a user group
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `user` (String) List of user names

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `source_group` (String) Source group name

    |Format  &emsp;|Description           |
    |----------|------------------------|
    |name    &emsp;|Name of source group  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
