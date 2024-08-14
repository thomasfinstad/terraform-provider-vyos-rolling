---
page_title: "vyos_service_webproxy_url_filtering_squidguard_auto_update Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Webproxy service settings⯯URL filtering settings⯯URL filtering via squidGuard redirector⯯Auto update settings
---

# vyos_service_webproxy_url_filtering_squidguard_auto_update (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Webproxy service settings  
⯯  
URL filtering settings  
⯯  
URL filtering via squidGuard redirector  
⯯  
**Auto update settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `update_hour` (Number) Hour of day for database update

    |Format  &emsp;|Description               |
    |----------|----------------------------|
    |0-23    &emsp;|Hour for database update  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
