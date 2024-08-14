---
page_title: "vyos_protocols_bgp_parameters Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Border Gateway Protocol (BGP)⯯BGP parameters
---

# vyos_protocols_bgp_parameters (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
**BGP parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `allow_martian_nexthop` (Boolean) Allow Martian nexthops to be received in the NLRI from a peer
- `always_compare_med` (Boolean) Always compare MEDs from different neighbors
- `cluster_id` (String) Route-reflector cluster-id

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |ipv4    &emsp;|Route-reflector cluster-id  |
- `deterministic_med` (Boolean) Compare MEDs between different peers in the same AS
- `disable_ebgp_connected_route_check` (Boolean) Disable checking if nexthop is connected on eBGP session
- `ebgp_requires_policy` (Boolean) Require in and out policy for eBGP peers (RFC8212)
- `fast_convergence` (Boolean) Teardown sessions immediately whenever peer becomes unreachable
- `graceful_shutdown` (Boolean) Graceful shutdown
- `labeled_unicast` (String) BGP Labeled-unicast options

    |Format              &emsp;|Description                                                 |
    |----------------------|--------------------------------------------------------------|
    |explicit-null       &emsp;|Use explicit-null label values for all local prefixes       |
    |ipv4-explicit-null  &emsp;|Use IPv4 explicit-null label value for IPv4 local prefixes  |
    |ipv6-explicit-null  &emsp;|Use IPv6 explicit-null label value for IPv4 local prefixes  |
- `log_neighbor_changes` (Boolean) Log neighbor up/down changes and reset reason
- `minimum_holdtime` (Number) BGP minimum holdtime

    |Format   &emsp;|Description                  |
    |-----------|-------------------------------|
    |1-65535  &emsp;|Minimum holdtime in seconds  |
- `network_import_check` (Boolean) Enable IGP route check for network statements
- `no_client_to_client_reflection` (Boolean) Disable client to client route reflection
- `no_fast_external_failover` (Boolean) Disable immediate session reset on peer link down event
- `no_hard_administrative_reset` (Boolean) Do not send hard reset CEASE Notification for &#39;Administrative Reset&#39;
- `no_suppress_duplicates` (Boolean) Disable suppress duplicate updates if the route actually not changed
- `reject_as_sets` (Boolean) Reject routes with AS_SET or AS_CONFED_SET flag
- `route_reflector_allow_outbound_policy` (Boolean) Route reflector client allow policy outbound
- `router_id` (String) Override default router identifier

    |Format  &emsp;|Description                     |
    |----------|----------------------------------|
    |ipv4    &emsp;|Router-ID in IP address format  |
- `shutdown` (Boolean) Administrative shutdown of the BGP instance
- `suppress_fib_pending` (Boolean) Advertise only routes that are programmed in kernel to peers
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
