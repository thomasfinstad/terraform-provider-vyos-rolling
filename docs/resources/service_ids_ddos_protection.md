---
page_title: "vyos_service_ids_ddos_protection Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Intrusion Detection System⯯FastNetMon detection and protection parameters
---

# vyos_service_ids_ddos_protection (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Intrusion Detection System  
⯯  
**FastNetMon detection and protection parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `alert_script` (String) Path to fastnetmon alert script
- `ban_time` (Number) How long we should keep an IP in blocked state

    |Format        &emsp;|Description      |
    |----------------|-------------------|
    |1-4294967294  &emsp;|Time in seconds  |
- `direction` (List of String) Direction for processing traffic
- `excluded_network` (List of String) Specify IPv4 and IPv6 networks which are going to be excluded from protection

    |Format   &emsp;|Description                 |
    |-----------|------------------------------|
    |ipv4net  &emsp;|IPv4 prefix(es) to exclude  |
    |ipv6net  &emsp;|IPv6 prefix(es) to exclude  |
- `listen_interface` (List of String) Listen interface for mirroring traffic
- `mode` (String) Traffic capture mode

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |mirror  &emsp;|Listen to mirrored traffic  |
    |sflow   &emsp;|Capture sFlow flows         |
- `network` (List of String) Specify IPv4 and IPv6 networks which belong to you

    |Format   &emsp;|Description           |
    |-----------|------------------------|
    |ipv4net  &emsp;|Your IPv4 prefix(es)  |
    |ipv6net  &emsp;|Your IPv6 prefix(es)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
