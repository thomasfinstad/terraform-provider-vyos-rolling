---
page_title: "vyos_firewall_global_options_timeout_tcp Resource - vyos"

subcategory: "Firewall"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Firewall⯯Global Options⯯Connection timeout options⯯TCP connection timeout options
---

# vyos_firewall_global_options_timeout_tcp (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

Firewall  
⯯  
Global Options  
⯯  
Connection timeout options  
⯯  
**TCP connection timeout options**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `close` (Number) TCP CLOSE timeout in seconds

    |Format      &emsp;|Description                   |
    |--------------|--------------------------------|
    |1-21474836  &emsp;|TCP CLOSE timeout in seconds  |
- `close_wait` (Number) TCP CLOSE-WAIT timeout in seconds

    |Format      &emsp;|Description                        |
    |--------------|-------------------------------------|
    |1-21474836  &emsp;|TCP CLOSE-WAIT timeout in seconds  |
- `established` (Number) TCP ESTABLISHED timeout in seconds

    |Format      &emsp;|Description                         |
    |--------------|--------------------------------------|
    |1-21474836  &emsp;|TCP ESTABLISHED timeout in seconds  |
- `fin_wait` (Number) TCP FIN-WAIT timeout in seconds

    |Format      &emsp;|Description                      |
    |--------------|-----------------------------------|
    |1-21474836  &emsp;|TCP FIN-WAIT timeout in seconds  |
- `last_ack` (Number) TCP LAST-ACK timeout in seconds

    |Format      &emsp;|Description                      |
    |--------------|-----------------------------------|
    |1-21474836  &emsp;|TCP LAST-ACK timeout in seconds  |
- `syn_recv` (Number) TCP SYN-RECEIVED timeout in seconds

    |Format      &emsp;|Description                          |
    |--------------|---------------------------------------|
    |1-21474836  &emsp;|TCP SYN-RECEIVED timeout in seconds  |
- `syn_sent` (Number) TCP SYN-SENT timeout in seconds

    |Format      &emsp;|Description                      |
    |--------------|-----------------------------------|
    |1-21474836  &emsp;|TCP SYN-SENT timeout in seconds  |
- `time_wait` (Number) TCP TIME-WAIT timeout in seconds

    |Format      &emsp;|Description                       |
    |--------------|------------------------------------|
    |1-21474836  &emsp;|TCP TIME-WAIT timeout in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
