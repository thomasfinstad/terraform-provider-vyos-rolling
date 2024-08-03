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

    &emsp;|Format  &emsp;|Description                                            |
    |----------|---------------------------------------------------------|
    &emsp;|none    &emsp;|Do nothing                                             |
    &emsp;|trap    &emsp;|Attempt to re-negotiate when matching traffic is seen  |
    &emsp;|start   &emsp;|Attempt to re-negotiate the connection immediately     |
- `dead_peer_detection` (Attributes) Dead Peer Detection (DPD) (see [below for nested schema](#nestedatt--dead_peer_detection))
- `disable_mobike` (Boolean) Disable MOBIKE Support (IKEv2 only)
- `ikev2_reauth` (Boolean) Re-authentication of the remote peer during an IKE re-key (IKEv2 only)
- `key_exchange` (String) IKE version

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|ikev1   &emsp;|Use IKEv1 for key exchange  |
    &emsp;|ikev2   &emsp;|Use IKEv2 for key exchange  |
- `lifetime` (Number) IKE lifetime

    &emsp;|Format   &emsp;|Description              |
    |-----------|---------------------------|
    &emsp;|0-86400  &emsp;|IKE lifetime in seconds  |
- `mode` (String) IKEv1 phase 1 mode

    &emsp;|Format      &emsp;|Description                                          |
    |--------------|-------------------------------------------------------|
    &emsp;|main        &emsp;|Use the main mode (recommended)                      |
    &emsp;|aggressive  &emsp;|Use the aggressive mode (insecure, not recommended)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `ike_group` (String) Internet Key Exchange (IKE) group name


&lt;a id=&#34;nestedatt--dead_peer_detection&#34;&gt;&lt;/a&gt;
### Nested Schema for `dead_peer_detection`

Optional:

- `action` (String) Keep-alive failure action

    &emsp;|Format   &emsp;|Description                                                           |
    |-----------|------------------------------------------------------------------------|
    &emsp;|trap     &emsp;|Attempt to re-negotiate the connection when matching traffic is seen  |
    &emsp;|clear    &emsp;|Remove the connection immediately                                     |
    &emsp;|restart  &emsp;|Attempt to re-negotiate the connection immediately                    |
- `interval` (Number) Keep-alive interval

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|2-86400  &emsp;|Keep-alive interval in seconds  |
- `timeout` (Number) Dead Peer Detection keep-alive timeout (IKEv1 only)

    &emsp;|Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    &emsp;|2-86400  &emsp;|Keep-alive timeout in seconds  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
