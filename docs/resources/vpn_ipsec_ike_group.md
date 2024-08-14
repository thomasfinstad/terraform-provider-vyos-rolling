---
page_title: "vyos_vpn_ipsec_ike_group Resource - vyos"

subcategory: "Vpn"

description: |- 
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯Internet Key Exchange (IKE) group name
---

# vyos_vpn_ipsec_ike_group (Resource)
<center>

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
**Internet Key Exchange (IKE) group name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `close_action` (String) Action to take if a child SA is unexpectedly closed

    |Format  &emsp;|Description                                            |
    |----------|---------------------------------------------------------|
    |none    &emsp;|Do nothing                                             |
    |trap    &emsp;|Attempt to re-negotiate when matching traffic is seen  |
    |start   &emsp;|Attempt to re-negotiate the connection immediately     |
- `dead_peer_detection` (Attributes) Dead Peer Detection (DPD) (see [below for nested schema](#nestedatt--dead_peer_detection))
- `disable_mobike` (Boolean) Disable MOBIKE Support (IKEv2 only)
- `ikev2_reauth` (Boolean) Re-authentication of the remote peer during an IKE re-key (IKEv2 only)
- `key_exchange` (String) IKE version

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |ikev1   &emsp;|Use IKEv1 for key exchange  |
    |ikev2   &emsp;|Use IKEv2 for key exchange  |
- `lifetime` (Number) IKE lifetime

    |Format   &emsp;|Description              |
    |-----------|---------------------------|
    |0-86400  &emsp;|IKE lifetime in seconds  |
- `mode` (String) IKEv1 phase 1 mode

    |Format      &emsp;|Description                                          |
    |--------------|-------------------------------------------------------|
    |main        &emsp;|Use the main mode (recommended)                      |
    |aggressive  &emsp;|Use the aggressive mode (insecure, not recommended)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `ike_group` (String) Internet Key Exchange (IKE) group name


<a id="nestedatt--dead_peer_detection"></a>
### Nested Schema for `dead_peer_detection`

Optional:

- `action` (String) Keep-alive failure action

    |Format   &emsp;|Description                                                           |
    |-----------|------------------------------------------------------------------------|
    |trap     &emsp;|Attempt to re-negotiate the connection when matching traffic is seen  |
    |clear    &emsp;|Remove the connection immediately                                     |
    |restart  &emsp;|Attempt to re-negotiate the connection immediately                    |
- `interval` (Number) Keep-alive interval

    |Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    |2-86400  &emsp;|Keep-alive interval in seconds  |
- `timeout` (Number) Dead Peer Detection keep-alive timeout (IKEv1 only)

    |Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    |2-86400  &emsp;|Keep-alive timeout in seconds  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
