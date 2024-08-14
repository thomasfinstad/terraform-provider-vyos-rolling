---
page_title: "vyos_interfaces_wireless_security_wpa_radius_server Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Wireless (WiFi/WLAN) Network Interface⯯Wireless security settings⯯Wifi Protected Access (WPA) parameters⯯RADIUS based user authentication⯯RADIUS server configuration
---

# vyos_interfaces_wireless_security_wpa_radius_server (Resource)
<center>

*interfaces*  
⯯  
Wireless (WiFi/WLAN) Network Interface  
⯯  
Wireless security settings  
⯯  
Wifi Protected Access (WPA) parameters  
⯯  
RADIUS based user authentication  
⯯  
**RADIUS server configuration**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `accounting` (Boolean) Enable RADIUS server to receive accounting info
- `disable` (Boolean) Disable instance
- `key` (String) Shared secret key

    |Format  &emsp;|Description            |
    |----------|-------------------------|
    |txt     &emsp;|Password string (key)  |
- `port` (Number) Port number used by connection

    |Format   &emsp;|Description      |
    |-----------|-------------------|
    |1-65535  &emsp;|Numeric IP port  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `server` (String) RADIUS server configuration

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |ipv4    &emsp;|RADIUS server IPv4 address  |
- `wireless` (String) Wireless (WiFi/WLAN) Network Interface

    |Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    |wlanN   &emsp;|Wireless (WiFi/WLAN) interface name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
