---
page_title: "vyos_service_lldp_interface Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯LLDP settings⯯Location data for interface
---

# vyos_service_lldp_interface (Resource)
<center>

*service*  
⯯  
LLDP settings  
⯯  
**Location data for interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `disable` (Boolean) Disable instance
- `location` (Attributes) LLDP-MED location data (see [below for nested schema](#nestedatt--location))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface` (String) Location data for interface

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|all     &emsp;|Location data all interfaces            |
    &emsp;|txt     &emsp;|Location data for a specific interface  |


&lt;a id=&#34;nestedatt--location&#34;&gt;&lt;/a&gt;
### Nested Schema for `location`

Optional:

- `coordinate_based` (Attributes) Coordinate based location (see [below for nested schema](#nestedatt--location--coordinate_based))
- `elin` (Number) ECS ELIN (Emergency location identifier number)

    &emsp;|Format        &emsp;|Description                                                 |
    |----------------|--------------------------------------------------------------|
    &emsp;|0-9999999999  &emsp;|Emergency Call Service ELIN number (between 10-25 numbers)  |

&lt;a id=&#34;nestedatt--location--coordinate_based&#34;&gt;&lt;/a&gt;
### Nested Schema for `location.coordinate_based`

Optional:

- `altitude` (String) Altitude in meters

    &emsp;|Format        &emsp;|Description         |
    |----------------|----------------------|
    &emsp;|0             &emsp;|No altitude         |
    &emsp;|[+-]&lt;meters&gt;  &emsp;|Altitude in meters  |
- `datum` (String) Coordinate datum type

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|WGS84   &emsp;|WGS84        |
    &emsp;|NAD83   &emsp;|NAD83        |
    &emsp;|MLLW    &emsp;|NAD83/MLLW   |
- `latitude` (String) Latitude

    &emsp;|Format      &emsp;|Description                      |
    |--------------|-----------------------------------|
    &emsp;|&lt;latitude&gt;  &emsp;|Latitude (example &#34;37.524449N&#34;)  |
- `longitude` (String) Longitude

    &emsp;|Format       &emsp;|Description                        |
    |---------------|-------------------------------------|
    &emsp;|&lt;longitude&gt;  &emsp;|Longitude (example &#34;122.267255W&#34;)  |



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
