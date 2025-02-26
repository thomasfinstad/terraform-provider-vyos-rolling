---
page_title: "vyos_protocols_isis_fast_reroute_lfa_local_load_sharing_disable Resource - vyos"

subcategory: "Protocols"

description: |-
  ~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.
  protocols⯯Intermediate System to Intermediate System (IS-IS)⯯IS-IS fast reroute configuration⯯Loop free alternate functionality⯯Local loop free alternate options⯯Load share prefixes across multiple backups⯯Disable load sharing
---

# vyos_protocols_isis_fast_reroute_lfa_local_load_sharing_disable (Resource)
<center>

~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*protocols*  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
IS-IS fast reroute configuration  
⯯  
Loop free alternate functionality  
⯯  
Local loop free alternate options  
⯯  
Load share prefixes across multiple backups  
⯯  
**Disable load sharing**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_protocols_isis_fast_reroute_lfa_local_load_sharing_disable (Resource)](#vyos_protocols_isis_fast_reroute_lfa_local_load_sharing_disable-resource)
  - [Schema](#schema)
    - [Optional](#optional)
      - [level_1](#level_1)
      - [level_2](#level_2)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

#### level_1
- `level_1` (Boolean) Match on IS-IS level-1 routes
#### level_2
- `level_2` (Boolean) Match on IS-IS level-2 routes
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_protocols_isis_fast_reroute_lfa_local_load_sharing_disable.example "protocols__isis__fast-reroute__lfa__local__load-sharing__disable"
```
