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

    |Format  &emsp;|Description                                                             |
    |----------|--------------------------------------------------------------------------|
    |0       &emsp;|Very basic auditing logs (e.g., SA up/SA down)                          |
    |1       &emsp;|Generic control flow with errors, a good default to see whats going on  |
    |2       &emsp;|More detailed debugging control flow                                    |
- `subsystem` (List of String) Subsystem logging levels

    |Format  &emsp;|Description                                                 |
    |----------|--------------------------------------------------------------|
    |dmn     &emsp;|Main daemon setup/cleanup/signal handling                   |
    |mgr     &emsp;|IKE_SA manager, handling synchronization for IKE_SA access  |
    |ike     &emsp;|IKE_SA/ISAKMP SA                                            |
    |chd     &emsp;|CHILD_SA/IPsec SA                                           |
    |job     &emsp;|Jobs queuing/processing and thread pool management          |
    |cfg     &emsp;|Configuration management and plugins                        |
    |knl     &emsp;|IPsec/Networking kernel interface                           |
    |net     &emsp;|IKE network communication                                   |
    |asn     &emsp;|Low-level encoding/decoding (ASN.1, X.509 etc.)             |
    |enc     &emsp;|Packet encoding/decoding encryption/decryption operations   |
    |lib     &emsp;|libstrongswan library messages                              |
    |esp     &emsp;|libipsec library messages                                   |
    |tls     &emsp;| libtls library messages                                    |
    |tnc     &emsp;|Trusted Network Connect                                     |
    |imc     &emsp;|Integrity Measurement Collector                             |
    |imv     &emsp;|Integrity Measurement Verifier                              |
    |pts     &emsp;| Platform Trust Service                                     |
    |any     &emsp;|Any subsystem                                               |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
