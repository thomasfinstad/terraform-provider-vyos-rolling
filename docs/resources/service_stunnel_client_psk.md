---
page_title: "vyos_service_stunnel_client_psk Resource - vyos"

subcategory: "Service"

description: |- 
  System services⯯Stunnel TLS Proxy⯯Stunnel client config⯯Pre-shared key name
---

# vyos_service_stunnel_client_psk (Resource)
<center>

System services  
⯯  
Stunnel TLS Proxy  
⯯  
Stunnel client config  
⯯  
**Pre-shared key name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `id_param` (String) ID for authentication

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|txt     &emsp;|ID used for authentication  |
- `secret` (String) pre-shared secret key

    &emsp;|Format  &emsp;|Description                                                                                                                |
    |----------|-----------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt     &emsp;|pre-shared secret key are required to be at least 16 bytes long, which implies at least 32 characters for hexadecimal key  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `client` (String) Stunnel client config
- `psk` (String) Pre-shared key name


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  