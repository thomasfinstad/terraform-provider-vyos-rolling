---
page_title: "vyos_service_stunnel_client_psk Resource - vyos"

subcategory: "Service"

description: |-
  service⯯Stunnel TLS Proxy⯯Stunnel client config⯯Pre-shared key name
---

# vyos_service_stunnel_client_psk (Resource)
<center>


*service*  
⯯  
Stunnel TLS Proxy  
⯯  
Stunnel client config  
⯯  
**Pre-shared key name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_service_stunnel_client_psk (Resource)](#vyos_service_stunnel_client_psk-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [id_param](#id_param)
      - [secret](#secret)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### id_param
- `id_param` (String) ID for authentication

    |  Format  &emsp;|  Description                 |
    |----------|------------------------------|
    |  txt     &emsp;|  ID used for authentication  |
#### secret
- `secret` (String) pre-shared secret key

    |  Format  &emsp;|  Description                                                  |
    |----------|---------------------------------------------------------------|
    |  txt     &emsp;|  16byte pre-shared-secret key (32 character hexadecimal key)  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `client` (String) Stunnel client config
- `psk` (String) Pre-shared key name


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_service_stunnel_client_psk.example "service__stunnel__client__<client>__psk__<psk>"
```
