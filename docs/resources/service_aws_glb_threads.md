---
page_title: "vyos_service_aws_glb_threads Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Amazon Web Service⯯Gateway load-balancer tunnel handler⯯Threads settings
---

# vyos_service_aws_glb_threads (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Amazon Web Service  
⯯  
Gateway load-balancer tunnel handler  
⯯  
**Threads settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `tunnel` (Number) Number of threads for each tunnel processor

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |1-256   &emsp;|Number of threads  |
- `tunnel_affinity` (String) List of cores worker threads

    |Format       &emsp;|Description                               |
    |---------------|--------------------------------------------|
    |&lt;idN&gt;-&lt;idM&gt;  &emsp;|CPU core id range (use &#39;-&#39; as delimiter)  |
- `udp` (Number) Number of threads for UDP receiver

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |1-256   &emsp;|Number of threads  |
- `udp_affinity` (String) List of cores worker threads

    |Format       &emsp;|Description                               |
    |---------------|--------------------------------------------|
    |&lt;idN&gt;-&lt;idM&gt;  &emsp;|CPU core id range (use &#39;-&#39; as delimiter)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
