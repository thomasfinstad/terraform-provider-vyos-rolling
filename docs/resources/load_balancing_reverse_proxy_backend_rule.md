---
page_title: "vyos_load_balancing_reverse_proxy_backend_rule Resource - vyos"

subcategory: "Load"

description: |- 
  load-balancing⯯Configure reverse-proxy⯯Backend server name⯯Proxy rule number
---

# vyos_load_balancing_reverse_proxy_backend_rule (Resource)
<center>

*load-balancing*  
⯯  
Configure reverse-proxy  
⯯  
Backend server name  
⯯  
**Proxy rule number**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `domain_name` (List of String) Domain name to match

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|txt     &emsp;|Domain address to match  |
- `set` (Attributes) Proxy modifications (see [below for nested schema](#nestedatt--set))
- `ssl` (String) SSL match options

    &emsp;|Format          &emsp;|Description                                                  |
    |------------------|---------------------------------------------------------------|
    &emsp;|req-ssl-sni     &emsp;|SSL Server Name Indication (SNI) request match               |
    &emsp;|ssl-fc-sni      &emsp;|SSL frontend connection Server Name Indication match         |
    &emsp;|ssl-fc-sni-end  &emsp;|SSL frontend match end of connection Server Name Indication  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `url_path` (Attributes) URL path match (see [below for nested schema](#nestedatt--url_path))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `backend` (String) Backend server name
- `rule` (Number) Proxy rule number

    &emsp;|Format   &emsp;|Description                 |
    |-----------|------------------------------|
    &emsp;|1-10000  &emsp;|Number for this proxy rule  |


&lt;a id=&#34;nestedatt--set&#34;&gt;&lt;/a&gt;
### Nested Schema for `set`

Optional:

- `redirect_location` (String) Set URL location

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|url     &emsp;|Set URL location  |
- `server` (String) Server name


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


&lt;a id=&#34;nestedatt--url_path&#34;&gt;&lt;/a&gt;
### Nested Schema for `url_path`

Optional:

- `begin` (List of String) Begin URL match

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|url     &emsp;|Begin URL    |
- `end` (List of String) End URL match

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|url     &emsp;|End URL      |
- `exact` (List of String) Exactly URL match

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|url     &emsp;|Exactly URL  |  
