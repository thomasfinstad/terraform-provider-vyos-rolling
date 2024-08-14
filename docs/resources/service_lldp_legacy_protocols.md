---
page_title: "vyos_service_lldp_legacy_protocols Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯LLDP settings⯯Legacy (vendor specific) protocols
---

# vyos_service_lldp_legacy_protocols (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
LLDP settings  
⯯  
**Legacy (vendor specific) protocols**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `cdp` (Boolean) Listen for CDP for Cisco routers/switches
- `edp` (Boolean) Listen for EDP for Extreme routers/switches
- `fdp` (Boolean) Listen for FDP for Foundry routers/switches
- `sonmp` (Boolean) Listen for SONMP for Nortel routers/switches
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  