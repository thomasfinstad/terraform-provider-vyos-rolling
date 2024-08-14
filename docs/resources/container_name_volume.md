---
page_title: "vyos_container_name_volume Resource - vyos"

subcategory: "Container"

description: |- 
  Container applications⯯Container name⯯Mount a volume into the container
---

# vyos_container_name_volume (Resource)
<center>

Container applications  
⯯  
Container name  
⯯  
**Mount a volume into the container**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `destination` (String) Destination container directory

    |Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    |txt     &emsp;|Destination container directory  |
- `mode` (String) Volume access mode ro/rw

    |Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    |ro      &emsp;|Volume mounted into the container as read-only   |
    |rw      &emsp;|Volume mounted into the container as read-write  |
- `propagation` (String) Volume bind propagation

    |Format    &emsp;|Description                                                                                                |
    |------------|-------------------------------------------------------------------------------------------------------------|
    |shared    &emsp;|Sub-mounts of the original mount are exposed to replica mounts                                             |
    |slave     &emsp;|Allow replica mount to see sub-mount from the original mount but not vice versa                            |
    |private   &emsp;|Sub-mounts within a mount are not visible to replica mounts or the original mount                          |
    |rshared   &emsp;|Allows sharing of mount points and their nested mount points between both the original and replica mounts  |
    |rslave    &emsp;|Allows mount point and their nested mount points between original an replica mounts                        |
    |rprivate  &emsp;|No mount points within original or replica mounts in any direction                                         |
- `source` (String) Source host directory

    |Format  &emsp;|Description            |
    |----------|-------------------------|
    |txt     &emsp;|Source host directory  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `name` (String) Container name
- `volume` (String) Mount a volume into the container


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
