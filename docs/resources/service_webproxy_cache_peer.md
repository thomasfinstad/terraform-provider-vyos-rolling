---
page_title: "vyos_service_webproxy_cache_peer Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Webproxy service settings⯯Specify other caches in a hierarchy
---

# vyos_service_webproxy_cache_peer (Resource)
<center>

*service*  
⯯  
Webproxy service settings  
⯯  
**Specify other caches in a hierarchy**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (String) Hostname or IP address of peer

    |Format    &emsp;|Description                    |
    |------------|---------------------------------|
    |ipv4      &emsp;|Squid cache-peer IPv4 address  |
    |hostname  &emsp;|Squid cache-peer hostname      |
- `http_port` (Number) Default Proxy Port

    |Format      &emsp;|Description          |
    |--------------|-----------------------|
    |1025-65535  &emsp;|Default port number  |
- `icp_port` (Number) Cache peer ICP port

    |Format   &emsp;|Description          |
    |-----------|-----------------------|
    |0        &emsp;|Cache peer disabled  |
    |1-65535  &emsp;|Cache peer ICP port  |
- `options` (String) Cache peer options

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |txt     &emsp;|Cache peer options  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `type` (String) Squid peer type (default parent)

    |Format     &emsp;|Description                            |
    |-------------|-----------------------------------------|
    |parent     &emsp;|Peer is a parent                       |
    |sibling    &emsp;|Peer is a sibling                      |
    |multicast  &emsp;|Peer is a member of a multicast group  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `cache_peer` (String) Specify other caches in a hierarchy

    |Format    &emsp;|Description       |
    |------------|--------------------|
    |hostname  &emsp;|Cache peers FQDN  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
