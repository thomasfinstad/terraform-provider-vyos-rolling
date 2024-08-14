---
page_title: "vyos_container_name Resource - vyos"

subcategory: "Container"

description: |- 
  Container applications⯯Container name
---

# vyos_container_name (Resource)
<center>

Container applications  
⯯  
**Container name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `allow_host_networks` (Boolean) Allow sharing host networking with container
- `allow_host_pid` (Boolean) Allow sharing host process namespace with container
- `arguments` (String) The command&#39;s arguments for this container
- `capability` (List of String) Grant individual Linux capability to container instance

    |Format            &emsp;|Description                                                            |
    |--------------------|-------------------------------------------------------------------------|
    |net-admin         &emsp;|Network operations (interface, firewall, routing tables)               |
    |net-bind-service  &emsp;|Bind a socket to privileged ports (port numbers less than 1024)        |
    |net-raw           &emsp;|Permission to create raw network sockets                               |
    |setpcap           &emsp;|Capability sets (from bounded or inherited set)                        |
    |sys-admin         &emsp;|Administation operations (quotactl, mount, sethostname, setdomainame)  |
    |sys-module        &emsp;|Load, unload and delete kernel modules                                 |
    |sys-nice          &emsp;|Permission to set process nice value                                   |
    |sys-time          &emsp;|Permission to set system clock                                         |
- `command` (String) Override the default CMD from the image
- `cpu_quota` (String) This limits the number of CPU resources the container can use

    |Format  &emsp;|Description                                                                         |
    |----------|--------------------------------------------------------------------------------------|
    |0       &emsp;|Unlimited                                                                           |
    |txt     &emsp;|Amount of CPU time the container can use in amount of cores (up to three decimals)  |
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `disable` (Boolean) Disable instance
- `entrypoint` (String) Override the default ENTRYPOINT from the image
- `gid` (Number) Group ID this container will run as

    |Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    |0-65535  &emsp;|Group ID this container will run as  |
- `host_name` (String) Container host name
- `image` (String) Container image to use

    |Format  &emsp;|Description                     |
    |----------|----------------------------------|
    |txt     &emsp;|Image name in the hub-registry  |
- `memory` (Number) Memory (RAM) available to this container

    |Format   &emsp;|Description                         |
    |-----------|--------------------------------------|
    |0        &emsp;|Unlimited                           |
    |1-16384  &emsp;|Container memory in megabytes (MB)  |
- `restart` (String) Restart options for container

    |Format      &emsp;|Description                                                                         |
    |--------------|--------------------------------------------------------------------------------------|
    |no          &emsp;|Do not restart containers on exit                                                   |
    |on-failure  &emsp;|Restart containers when they exit with a non-zero exit code, retrying indefinitely  |
    |always      &emsp;|Restart containers when they exit, regardless of status, retrying indefinitely      |
- `shared_memory` (Number) Shared memory available to this container

    |Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    |0       &emsp;|Unlimited                           |
    |1-8192  &emsp;|Container memory in megabytes (MB)  |
- `sysctl` (Attributes) Configure namespaced kernel parameters of the container (see [below for nested schema](#nestedatt--sysctl))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `uid` (Number) User ID this container will run as

    |Format   &emsp;|Description                         |
    |-----------|--------------------------------------|
    |0-65535  &emsp;|User ID this container will run as  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `name` (String) Container name


<a id="nestedatt--sysctl"></a>
### Nested Schema for `sysctl`


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
