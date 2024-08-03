---
page_title: "vyos_vpn_ipsec_esp_group Resource - vyos"

subcategory: "Vpn"

description: |- 
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯Encapsulating Security Payload (ESP) group name
---

# vyos_vpn_ipsec_esp_group (Resource)
<center>

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
**Encapsulating Security Payload (ESP) group name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `compression` (Boolean) Enable ESP compression
- `disable_rekey` (Boolean) Do not locally initiate a re-key of the SA, remote peer must re-key before expiration
- `life_bytes` (Number) Security Association byte count to expire

    &emsp;|Format               &emsp;|Description       |
    |-----------------------|--------------------|
    &emsp;|1024-26843545600000  &emsp;|SA life in bytes  |
- `life_packets` (Number) Security Association packet count to expire

    &emsp;|Format               &emsp;|Description         |
    |-----------------------|----------------------|
    &emsp;|1000-26843545600000  &emsp;|SA life in packets  |
- `lifetime` (Number) Security Association time to expire

    &emsp;|Format    &emsp;|Description             |
    |------------|--------------------------|
    &emsp;|30-86400  &emsp;|SA lifetime in seconds  |
- `mode` (String) ESP mode

    &emsp;|Format     &emsp;|Description     |
    |-------------|------------------|
    &emsp;|tunnel     &emsp;|Tunnel mode     |
    &emsp;|transport  &emsp;|Transport mode  |
- `pfs` (String) ESP Perfect Forward Secrecy

    &emsp;|Format      &emsp;|Description                                      |
    |--------------|---------------------------------------------------|
    &emsp;|enable      &emsp;|Inherit Diffie-Hellman group from the IKE group  |
    &emsp;|dh-group1   &emsp;|Use Diffie-Hellman group 1 (modp768)             |
    &emsp;|dh-group2   &emsp;|Use Diffie-Hellman group 2 (modp1024)            |
    &emsp;|dh-group5   &emsp;|Use Diffie-Hellman group 5 (modp1536)            |
    &emsp;|dh-group14  &emsp;|Use Diffie-Hellman group 14 (modp2048)           |
    &emsp;|dh-group15  &emsp;|Use Diffie-Hellman group 15 (modp3072)           |
    &emsp;|dh-group16  &emsp;|Use Diffie-Hellman group 16 (modp4096)           |
    &emsp;|dh-group17  &emsp;|Use Diffie-Hellman group 17 (modp6144)           |
    &emsp;|dh-group18  &emsp;|Use Diffie-Hellman group 18 (modp8192)           |
    &emsp;|dh-group19  &emsp;|Use Diffie-Hellman group 19 (ecp256)             |
    &emsp;|dh-group20  &emsp;|Use Diffie-Hellman group 20 (ecp384)             |
    &emsp;|dh-group21  &emsp;|Use Diffie-Hellman group 21 (ecp521)             |
    &emsp;|dh-group22  &emsp;|Use Diffie-Hellman group 22 (modp1024s160)       |
    &emsp;|dh-group23  &emsp;|Use Diffie-Hellman group 23 (modp2048s224)       |
    &emsp;|dh-group24  &emsp;|Use Diffie-Hellman group 24 (modp2048s256)       |
    &emsp;|dh-group25  &emsp;|Use Diffie-Hellman group 25 (ecp192)             |
    &emsp;|dh-group26  &emsp;|Use Diffie-Hellman group 26 (ecp224)             |
    &emsp;|dh-group27  &emsp;|Use Diffie-Hellman group 27 (ecp224bp)           |
    &emsp;|dh-group28  &emsp;|Use Diffie-Hellman group 28 (ecp256bp)           |
    &emsp;|dh-group29  &emsp;|Use Diffie-Hellman group 29 (ecp384bp)           |
    &emsp;|dh-group30  &emsp;|Use Diffie-Hellman group 30 (ecp512bp)           |
    &emsp;|dh-group31  &emsp;|Use Diffie-Hellman group 31 (curve25519)         |
    &emsp;|dh-group32  &emsp;|Use Diffie-Hellman group 32 (curve448)           |
    &emsp;|disable     &emsp;|Disable PFS                                      |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `esp_group` (String) Encapsulating Security Payload (ESP) group name


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
