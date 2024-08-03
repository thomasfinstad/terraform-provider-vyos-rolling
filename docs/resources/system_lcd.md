---
page_title: "vyos_system_lcd Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯System LCD display
---

# vyos_system_lcd (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
**System LCD display**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `device` (String) Physical device used by LCD display

    &emsp;|Format    &emsp;|Description                           |
    |------------|----------------------------------------|
    &emsp;|ttySXX    &emsp;|TTY device name, regular serial port  |
    &emsp;|usbNbXpY  &emsp;|TTY device name, USB based            |
- `model` (String) Model of the display attached to this system

    &emsp;|Format   &emsp;|Description                                            |
    |-----------|---------------------------------------------------------|
    &emsp;|cfa-533  &emsp;|Crystalfontz CFA-533                                   |
    &emsp;|cfa-631  &emsp;|Crystalfontz CFA-631                                   |
    &emsp;|cfa-633  &emsp;|Crystalfontz CFA-633                                   |
    &emsp;|cfa-635  &emsp;|Crystalfontz CFA-635                                   |
    &emsp;|hd44780  &emsp;|Hitachi HD44780, Caswell Appliances                    |
    &emsp;|sdec     &emsp;|Lanner, Watchguard, Nexcom NSA, Sophos UTM appliances  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
