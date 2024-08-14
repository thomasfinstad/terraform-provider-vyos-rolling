---
page_title: "vyos_system_option Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯System Options
---

# vyos_system_option (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
**System Options**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `ctrl_alt_delete` (String) System action on Ctrl-Alt-Delete keystroke

    |Format    &emsp;|Description          |
    |------------|-----------------------|
    |ignore    &emsp;|Ignore key sequence  |
    |reboot    &emsp;|Reboot system        |
    |poweroff  &emsp;|Poweroff system      |
- `disable_usb_autosuspend` (Boolean) Disable autosuspend for all USB devices
- `keyboard_layout` (String) System keyboard layout, type ISO2

    |Format     &emsp;|Description     |
    |-------------|------------------|
    |us         &emsp;|United States   |
    |uk         &emsp;|United Kingdom  |
    |fr         &emsp;|France          |
    |de         &emsp;|Germany         |
    |es         &emsp;|Spain           |
    |fi         &emsp;|Finland         |
    |jp106      &emsp;|Japan           |
    |no         &emsp;|Norway          |
    |dk         &emsp;|Denmark         |
    |se-latin1  &emsp;|Sweden          |
    |dvorak     &emsp;|Dvorak          |
- `performance` (String) Tune system performance

    |Format      &emsp;|Description                          |
    |--------------|---------------------------------------|
    |throughput  &emsp;|Tune for maximum network throughput  |
    |latency     &emsp;|Tune for low network latency         |
- `reboot_on_panic` (Boolean) Reboot system on kernel panic
- `root_partition_auto_resize` (Boolean) Enable root partition auto-extention on system boot
- `startup_beep` (Boolean) plays sound via system speaker when you can login
- `time_format` (String) System time-format

    |Format   &emsp;|Description          |
    |-----------|-----------------------|
    |12-hour  &emsp;|12 hour time format  |
    |24-hour  &emsp;|24 hour time format  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
