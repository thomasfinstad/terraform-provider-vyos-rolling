---
page_title: "vyos_load_balancing_reverse_proxy_service_rule Resource - vyos"

subcategory: "Load"

description: |- 
  load-balancing⯯Configure reverse-proxy⯯Frontend service name⯯Proxy rule number
---

# vyos_load_balancing_reverse_proxy_service_rule (Resource)
<center>

*load-balancing*  
⯯  
Configure reverse-proxy  
⯯  
Frontend service name  
⯯  
**Proxy rule number**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `domain_name` (List of String) Domain name to match

    |Format  &emsp;|Description              |
    |----------|---------------------------|
    |txt     &emsp;|Domain address to match  |
- `set` (Attributes) Proxy modifications (see [below for nested schema](#nestedatt--set))
- `ssl` (String) SSL match options

    |Format          &emsp;|Description                                                  |
    |------------------|---------------------------------------------------------------|
    |req-ssl-sni     &emsp;|SSL Server Name Indication (SNI) request match               |
    |ssl-fc-sni      &emsp;|SSL frontend connection Server Name Indication match         |
    |ssl-fc-sni-end  &emsp;|SSL frontend match end of connection Server Name Indication  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `url_path` (Attributes) URL path match (see [below for nested schema](#nestedatt--url_path))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `rule` (Number) Proxy rule number

    |Format   &emsp;|Description                 |
    |-----------|------------------------------|
    |1-10000  &emsp;|Number for this proxy rule  |
- `service` (String) Frontend service name


<a id="nestedatt--set"></a>
### Nested Schema for `set`

Optional:

- `backend` (String) Backend name
- `redirect_location` (String) Set URL location

    |Format  &emsp;|Description       |
    |----------|--------------------|
    |url     &emsp;|Set URL location  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


<a id="nestedatt--url_path"></a>
### Nested Schema for `url_path`

Optional:

- `begin` (List of String) Begin URL match

    |Format  &emsp;|Description  |
    |----------|---------------|
    |url     &emsp;|Begin URL    |
- `end` (List of String) End URL match

    |Format  &emsp;|Description  |
    |----------|---------------|
    |url     &emsp;|End URL      |
- `exact` (List of String) Exactly URL match

    |Format  &emsp;|Description  |
    |----------|---------------|
    |url     &emsp;|Exactly URL  |  
