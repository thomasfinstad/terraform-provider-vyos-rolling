---
page_title: "vyos_service_console_server_device Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Serial Console Server⯯System serial interface name (ttyS or ttyUSB)
---

# vyos_service_console_server_device (Resource)
<center>

*service*  
⯯  
Serial Console Server  
⯯  
**System serial interface name (ttyS or ttyUSB)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `alias` (String) Human-readable name for this console
- `data_bits` (String) Serial port data bits
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `parity` (String) Parity setting
- `speed` (String) Serial port baud rate
- `ssh` (Attributes) SSH remote access to this console (see [below for nested schema](#nestedatt--ssh))
- `stop_bits` (String) Serial port stop bits
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `device` (String) System serial interface name (ttyS or ttyUSB)

    &emsp;|Format    &emsp;|Description                 |
    |------------|------------------------------|
    &emsp;|ttySxxx   &emsp;|Regular serial interface    |
    &emsp;|usbxbxpx  &emsp;|USB based serial interface  |


&lt;a id=&#34;nestedatt--ssh&#34;&gt;&lt;/a&gt;
### Nested Schema for `ssh`

Optional:

- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
