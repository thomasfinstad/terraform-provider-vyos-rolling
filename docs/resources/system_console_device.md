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

<!--TOC-->

- [vyos_system_console_device (Resource)](#vyos_system_console_device-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [speed](#speed)
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

#### speed
- `speed` (String) Console baud rate

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  1200    &emsp;|  1200 bps     |
    |  2400    &emsp;|  2400 bps     |
    |  4800    &emsp;|  4800 bps     |
    |  9600    &emsp;|  9600 bps     |
    |  19200   &emsp;|  19200 bps    |
    |  38400   &emsp;|  38400 bps    |
    |  57600   &emsp;|  57600 bps    |
    |  115200  &emsp;|  115200 bps   |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `device` (String) Serial console device name

    |  Format    &emsp;|  Description                           |
    |------------|----------------------------------------|
    |  ttySN     &emsp;|  TTY device name, regular serial port  |
    |  usbNbXpY  &emsp;|  TTY device name, USB based            |
    |  hvcN      &emsp;|  Xen console                           |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_system_console_device.example "system__console__device__<device>"
```
