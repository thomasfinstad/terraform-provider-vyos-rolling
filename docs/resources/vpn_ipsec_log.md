---
page_title: "vyos_vpn_ipsec_log Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯IPsec logging
---

# vyos_vpn_ipsec_log (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
**IPsec logging**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `level` (String) Global IPsec logging Level

    &emsp;|Format  &emsp;|Description                                                             |
    |----------|--------------------------------------------------------------------------|
    &emsp;|0       &emsp;|Very basic auditing logs (e.g., SA up/SA down)                          |
    &emsp;|1       &emsp;|Generic control flow with errors, a good default to see whats going on  |
    &emsp;|2       &emsp;|More detailed debugging control flow                                    |
- `subsystem` (List of String) Subsystem logging levels

    &emsp;|Format  &emsp;|Description                                                 |
    |----------|--------------------------------------------------------------|
    &emsp;|dmn     &emsp;|Main daemon setup/cleanup/signal handling                   |
    &emsp;|mgr     &emsp;|IKE_SA manager, handling synchronization for IKE_SA access  |
    &emsp;|ike     &emsp;|IKE_SA/ISAKMP SA                                            |
    &emsp;|chd     &emsp;|CHILD_SA/IPsec SA                                           |
    &emsp;|job     &emsp;|Jobs queuing/processing and thread pool management          |
    &emsp;|cfg     &emsp;|Configuration management and plugins                        |
    &emsp;|knl     &emsp;|IPsec/Networking kernel interface                           |
    &emsp;|net     &emsp;|IKE network communication                                   |
    &emsp;|asn     &emsp;|Low-level encoding/decoding (ASN.1, X.509 etc.)             |
    &emsp;|enc     &emsp;|Packet encoding/decoding encryption/decryption operations   |
    &emsp;|lib     &emsp;|libstrongswan library messages                              |
    &emsp;|esp     &emsp;|libipsec library messages                                   |
    &emsp;|tls     &emsp;| libtls library messages                                    |
    &emsp;|tnc     &emsp;|Trusted Network Connect                                     |
    &emsp;|imc     &emsp;|Integrity Measurement Collector                             |
    &emsp;|imv     &emsp;|Integrity Measurement Verifier                              |
    &emsp;|pts     &emsp;| Platform Trust Service                                     |
    &emsp;|any     &emsp;|Any subsystem                                               |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
