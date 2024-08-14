---
page_title: "vyos_vrf_name_protocols_isis_interface Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Intermediate System to Intermediate System (IS-IS)⯯Interface params
---

# vyos_vrf_name_protocols_isis_interface (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
**Interface params**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `bfd` (Attributes) Enable Bidirectional Forwarding Detection (BFD) (see [below for nested schema](#nestedatt--bfd))
- `circuit_type` (String) Configure circuit type for interface

    |Format        &emsp;|Description                          |
    |----------------|---------------------------------------|
    |level-1       &emsp;|Level-1 only adjacencies are formed  |
    |level-1-2     &emsp;|Level-1-2 adjacencies are formed     |
    |level-2-only  &emsp;|Level-2 only adjacencies are formed  |
- `hello_interval` (Number) Set Hello interval

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |1-600   &emsp;|Set Hello interval  |
- `hello_multiplier` (Number) Set Hello interval

    |Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    |2-100   &emsp;|Set multiplier for Hello holding time  |
- `hello_padding` (Boolean) Add padding to IS-IS hello packets
- `ldp_sync` (Attributes) LDP-IGP synchronization configuration for interface (see [below for nested schema](#nestedatt--ldp_sync))
- `metric` (Number) Set default metric for circuit

    |Format      &emsp;|Description           |
    |--------------|------------------------|
    |0-16777215  &emsp;|Default metric value  |
- `network` (Attributes) Set network type (see [below for nested schema](#nestedatt--network))
- `no_three_way_handshake` (Boolean) Disable three-way handshake
- `passive` (Boolean) Configure passive mode for interface
- `password` (Attributes) Configure the authentication password for a circuit (see [below for nested schema](#nestedatt--password))
- `priority` (Number) Set priority for Designated Router election

    |Format  &emsp;|Description     |
    |----------|------------------|
    |0-127   &emsp;|Priority value  |
- `psnp_interval` (Number) Set PSNP interval

    |Format  &emsp;|Description               |
    |----------|----------------------------|
    |0-127   &emsp;|PSNP interval in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface params
- `name` (String) Virtual Routing and Forwarding instance

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |


<a id="nestedatt--bfd"></a>
### Nested Schema for `bfd`

Optional:

- `profile` (String) Use settings from BFD profile

    |Format  &emsp;|Description       |
    |----------|--------------------|
    |txt     &emsp;|BFD profile name  |


<a id="nestedatt--ldp_sync"></a>
### Nested Schema for `ldp_sync`

Optional:

- `disable` (Boolean) Disable instance
- `holddown` (Number) Hold down timer for LDP-IGP cost restoration

    |Format   &emsp;|Description                                                                                   |
    |-----------|------------------------------------------------------------------------------------------------|
    |0-10000  &emsp;|Time to wait in seconds for LDP-IGP synchronization to occur before restoring interface cost  |


<a id="nestedatt--network"></a>
### Nested Schema for `network`

Optional:

- `point_to_point` (Boolean) point-to-point network type


<a id="nestedatt--password"></a>
### Nested Schema for `password`

Optional:

- `md5` (String) MD5 authentication type

    |Format  &emsp;|Description          |
    |----------|-----------------------|
    |txt     &emsp;|Level-wide password  |
- `plaintext_password` (String) Plain-text authentication type

    |Format  &emsp;|Description       |
    |----------|--------------------|
    |txt     &emsp;|Circuit password  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
