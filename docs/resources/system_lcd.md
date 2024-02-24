---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_system_lcd Resource - vyos"
subcategory: "system"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>system</i>

  <br>
  &darr;
  <br>
  <b>
  System LCD display
  </b>
  </div>
---

# vyos_system_lcd (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>system</i>

<br>
&darr;
<br>
<b>
System LCD display
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) Physical device used by LCD display

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ttySXX  &emsp; |  TTY device name, regular serial port  |
    |  usbNbXpY  &emsp; |  TTY device name, USB based  |
- `model` (String) Model of the display attached to this system

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  cfa-533  &emsp; |  Crystalfontz CFA-533  |
    |  cfa-631  &emsp; |  Crystalfontz CFA-631  |
    |  cfa-633  &emsp; |  Crystalfontz CFA-633  |
    |  cfa-635  &emsp; |  Crystalfontz CFA-635  |
    |  hd44780  &emsp; |  Hitachi HD44780, Caswell Appliances  |
    |  sdec  &emsp; |  Lanner, Watchguard, Nexcom NSA, Sophos UTM appliances  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).