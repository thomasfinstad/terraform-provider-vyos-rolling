---
page_title: "vyos_protocols_eigrp_metric Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Enhanced Interior Gateway Routing Protocol (EIGRP)⯯Modify metrics and parameters for advertisement
---

# vyos_protocols_eigrp_metric (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Enhanced Interior Gateway Routing Protocol (EIGRP)  
⯯  
**Modify metrics and parameters for advertisement**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `weights` (Number) Modify metric coefficients

    |Format  &emsp;|Description  |
    |----------|---------------|
    |0-255   &emsp;|K1           |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
