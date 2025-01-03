---
page_title: "vyos_service_dns_forwarding_authoritative_domain Resource - vyos"

subcategory: "Service"

description: |-
  service⯯Domain Name System (DNS) related services⯯DNS forwarding⯯Domain to host authoritative records for
---

# vyos_service_dns_forwarding_authoritative_domain (Resource)
<center>


*service*  
⯯  
Domain Name System (DNS) related services  
⯯  
DNS forwarding  
⯯  
**Domain to host authoritative records for**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_service_dns_forwarding_authoritative_domain (Resource)](#vyos_service_dns_forwarding_authoritative_domain-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [disable](#disable)
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

#### disable
- `disable` (Boolean) Disable instance
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `authoritative_domain` (String) Domain to host authoritative records for

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  txt     &emsp;|  An absolute DNS domain name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_service_dns_forwarding_authoritative_domain.example "service__dns__forwarding__authoritative-domain__<authoritative-domain>"
```
