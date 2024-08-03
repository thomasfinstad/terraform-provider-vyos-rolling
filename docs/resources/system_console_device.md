---
page_title: "vyos_system_console_device Resource - vyos"

subcategory: "System"

description: |- 
  system⯯Serial console configuration⯯Serial console device name
---

# vyos_system_console_device (Resource)
<center>

*system*  
⯯  
Serial console configuration  
⯯  
**Serial console device name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `speed` (String) Console baud rate

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1200    &emsp;|1200 bps     |
    &emsp;|2400    &emsp;|2400 bps     |
    &emsp;|4800    &emsp;|4800 bps     |
    &emsp;|9600    &emsp;|9600 bps     |
    &emsp;|19200   &emsp;|19200 bps    |
    &emsp;|38400   &emsp;|38400 bps    |
    &emsp;|57600   &emsp;|57600 bps    |
    &emsp;|115200  &emsp;|115200 bps   |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `device` (String) Serial console device name

    &emsp;|Format    &emsp;|Description                           |
    |------------|----------------------------------------|
    &emsp;|ttySN     &emsp;|TTY device name, regular serial port  |
    &emsp;|usbNbXpY  &emsp;|TTY device name, USB based            |
    &emsp;|hvcN      &emsp;|Xen console                           |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
