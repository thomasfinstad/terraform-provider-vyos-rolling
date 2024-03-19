---
page_title: "vyos_container_name Resource - terraform-provider-vyos"
subcategory: "container"
description: |-
  Container applications
  ⯯
  Container name
---

# vyos_container_name (Resource)
<center>

Container applications
⯯
**Container name**


</center>

## Schema

### Required

- `name_id` (String) Container name

### Optional

- `allow_host_networks` (Boolean) Allow host networks in container
- `arguments` (String) The command&#39;s arguments for this container
- `cap_add` (List of String) Container capabilities/permissions

    &emsp;|Format            &emsp;|Description                                                            |
    |--------------------|-------------------------------------------------------------------------|
    &emsp;|net-admin         &emsp;|Network operations (interface, firewall, routing tables)               |
    &emsp;|net-bind-service  &emsp;|Bind a socket to privileged ports (port numbers less than 1024)        |
    &emsp;|net-raw           &emsp;|Permission to create raw network sockets                               |
    &emsp;|setpcap           &emsp;|Capability sets (from bounded or inherited set)                        |
    &emsp;|sys-admin         &emsp;|Administation operations (quotactl, mount, sethostname, setdomainame)  |
    &emsp;|sys-module        &emsp;|Load, unload and delete kernel modules                                 |
    &emsp;|sys-time          &emsp;|Permission to set system clock                                         |
- `command` (String) Override the default CMD from the image
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `disable` (Boolean) Disable instance
- `entrypoint` (String) Override the default ENTRYPOINT from the image
- `gid` (Number) Group ID this container will run as

    &emsp;|Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    &emsp;|0-65535  &emsp;|Group ID this container will run as  |
- `host_name` (String) Container host name
- `image` (String) Image name in the hub-registry
- `memory` (Number) Memory (RAM) available to this container

    &emsp;|Format   &emsp;|Description                         |
    |-----------|--------------------------------------|
    &emsp;|0        &emsp;|Unlimited                           |
    &emsp;|1-16384  &emsp;|Container memory in megabytes (MB)  |
- `restart` (String) Restart options for container

    &emsp;|Format      &emsp;|Description                                                                         |
    |--------------|--------------------------------------------------------------------------------------|
    &emsp;|no          &emsp;|Do not restart containers on exit                                                   |
    &emsp;|on-failure  &emsp;|Restart containers when they exit with a non-zero exit code, retrying indefinitely  |
    &emsp;|always      &emsp;|Restart containers when they exit, regardless of status, retrying indefinitely      |
- `shared_memory` (Number) Shared memory available to this container

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|0       &emsp;|Unlimited                           |
    &emsp;|1-8192  &emsp;|Container memory in megabytes (MB)  |
- `uid` (Number) User ID this container will run as

    &emsp;|Format   &emsp;|Description                         |
    |-----------|--------------------------------------|
    &emsp;|0-65535  &emsp;|User ID this container will run as  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
