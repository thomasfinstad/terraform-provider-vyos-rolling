---
page_title: "vyos_container_name_volume Resource - terraform-provider-vyos"
subcategory: "container"
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

## Schema

### Required

- `name_id` (String) Container name
- `volume_id` (String) Mount a volume into the container

### Optional

- `destination` (String) Destination container directory

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|txt     &emsp;|Destination container directory  |
- `mode` (String) Volume access mode ro/rw

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|ro      &emsp;|Volume mounted into the container as read-only   |
    &emsp;|rw      &emsp;|Volume mounted into the container as read-write  |
- `propagation` (String) Volume bind propagation

    &emsp;|Format    &emsp;|Description                                                                                                |
    |------------|-------------------------------------------------------------------------------------------------------------|
    &emsp;|shared    &emsp;|Sub-mounts of the original mount are exposed to replica mounts                                             |
    &emsp;|slave     &emsp;|Allow replica mount to see sub-mount from the original mount but not vice versa                            |
    &emsp;|private   &emsp;|Sub-mounts within a mount are not visible to replica mounts or the original mount                          |
    &emsp;|rshared   &emsp;|Allows sharing of mount points and their nested mount points between both the original and replica mounts  |
    &emsp;|rslave    &emsp;|Allows mount point and their nested mount points between original an replica mounts                        |
    &emsp;|rprivate  &emsp;|No mount points within original or replica mounts in any direction                                         |
- `source` (String) Source host directory

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|txt     &emsp;|Source host directory  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
