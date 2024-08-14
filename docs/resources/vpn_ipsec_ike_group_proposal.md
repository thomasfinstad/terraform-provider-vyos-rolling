---
page_title: "vyos_vpn_ipsec_ike_group_proposal Resource - vyos"

subcategory: "Vpn"

description: |- 
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯Internet Key Exchange (IKE) group name⯯IKE proposal
---

# vyos_vpn_ipsec_ike_group_proposal (Resource)
<center>

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
Internet Key Exchange (IKE) group name  
⯯  
**IKE proposal**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `dh_group` (String) dh-grouphelp

    |Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    |1       &emsp;|Diffie-Hellman group 1 (modp768)        |
    |2       &emsp;|Diffie-Hellman group 2 (modp1024)       |
    |5       &emsp;|Diffie-Hellman group 5 (modp1536)       |
    |14      &emsp;|Diffie-Hellman group 14 (modp2048)      |
    |15      &emsp;|Diffie-Hellman group 15 (modp3072)      |
    |16      &emsp;|Diffie-Hellman group 16 (modp4096)      |
    |17      &emsp;|Diffie-Hellman group 17 (modp6144)      |
    |18      &emsp;|Diffie-Hellman group 18 (modp8192)      |
    |19      &emsp;|Diffie-Hellman group 19 (ecp256)        |
    |20      &emsp;|Diffie-Hellman group 20 (ecp384)        |
    |21      &emsp;|Diffie-Hellman group 21 (ecp521)        |
    |22      &emsp;|Diffie-Hellman group 22 (modp1024s160)  |
    |23      &emsp;|Diffie-Hellman group 23 (modp2048s224)  |
    |24      &emsp;|Diffie-Hellman group 24 (modp2048s256)  |
    |25      &emsp;|Diffie-Hellman group 25 (ecp192)        |
    |26      &emsp;|Diffie-Hellman group 26 (ecp224)        |
    |27      &emsp;|Diffie-Hellman group 27 (ecp224bp)      |
    |28      &emsp;|Diffie-Hellman group 28 (ecp256bp)      |
    |29      &emsp;|Diffie-Hellman group 29 (ecp384bp)      |
    |30      &emsp;|Diffie-Hellman group 30 (ecp512bp)      |
    |31      &emsp;|Diffie-Hellman group 31 (curve25519)    |
    |32      &emsp;|Diffie-Hellman group 32 (curve448)      |
- `encryption` (String) Encryption algorithm

    |Format             &emsp;|Description                                 |
    |---------------------|----------------------------------------------|
    |null               &emsp;|Null encryption                             |
    |aes128             &emsp;|128 bit AES-CBC                             |
    |aes192             &emsp;|192 bit AES-CBC                             |
    |aes256             &emsp;|256 bit AES-CBC                             |
    |aes128ctr          &emsp;|128 bit AES-COUNTER                         |
    |aes192ctr          &emsp;|192 bit AES-COUNTER                         |
    |aes256ctr          &emsp;|256 bit AES-COUNTER                         |
    |aes128ccm64        &emsp;|128 bit AES-CCM with 64 bit ICV             |
    |aes192ccm64        &emsp;|192 bit AES-CCM with 64 bit ICV             |
    |aes256ccm64        &emsp;|256 bit AES-CCM with 64 bit ICV             |
    |aes128ccm96        &emsp;|128 bit AES-CCM with 96 bit ICV             |
    |aes192ccm96        &emsp;|192 bit AES-CCM with 96 bit ICV             |
    |aes256ccm96        &emsp;|256 bit AES-CCM with 96 bit ICV             |
    |aes128ccm128       &emsp;|128 bit AES-CCM with 128 bit ICV            |
    |aes192ccm128       &emsp;|192 bit AES-CCM with 128 bit IC             |
    |aes256ccm128       &emsp;|256 bit AES-CCM with 128 bit ICV            |
    |aes128gcm64        &emsp;|128 bit AES-GCM with 64 bit ICV             |
    |aes192gcm64        &emsp;|192 bit AES-GCM with 64 bit ICV             |
    |aes256gcm64        &emsp;|256 bit AES-GCM with 64 bit ICV             |
    |aes128gcm96        &emsp;|128 bit AES-GCM with 96 bit ICV             |
    |aes192gcm96        &emsp;|192 bit AES-GCM with 96 bit ICV             |
    |aes256gcm96        &emsp;|256 bit AES-GCM with 96 bit ICV             |
    |aes128gcm128       &emsp;|128 bit AES-GCM with 128 bit ICV            |
    |aes192gcm128       &emsp;|192 bit AES-GCM with 128 bit ICV            |
    |aes256gcm128       &emsp;|256 bit AES-GCM with 128 bit ICV            |
    |aes128gmac         &emsp;|Null encryption with 128 bit AES-GMAC       |
    |aes192gmac         &emsp;|Null encryption with 192 bit AES-GMAC       |
    |aes256gmac         &emsp;|Null encryption with 256 bit AES-GMAC       |
    |3des               &emsp;|168 bit 3DES-EDE-CBC                        |
    |blowfish128        &emsp;|128 bit Blowfish-CBC                        |
    |blowfish192        &emsp;|192 bit Blowfish-CBC                        |
    |blowfish256        &emsp;|256 bit Blowfish-CBC                        |
    |camellia128        &emsp;|128 bit Camellia-CBC                        |
    |camellia192        &emsp;|192 bit Camellia-CBC                        |
    |camellia256        &emsp;|256 bit Camellia-CBC                        |
    |camellia128ctr     &emsp;|128 bit Camellia-COUNTER                    |
    |camellia192ctr     &emsp;|192 bit Camellia-COUNTER                    |
    |camellia256ctr     &emsp;|256 bit Camellia-COUNTER                    |
    |camellia128ccm64   &emsp;|128 bit Camellia-CCM with 64 bit ICV        |
    |camellia192ccm64   &emsp;|192 bit Camellia-CCM with 64 bit ICV        |
    |camellia256ccm64   &emsp;|256 bit Camellia-CCM with 64 bit ICV        |
    |camellia128ccm96   &emsp;|128 bit Camellia-CCM with 96 bit ICV        |
    |camellia192ccm96   &emsp;|192 bit Camellia-CCM with 96 bit ICV        |
    |camellia256ccm96   &emsp;|256 bit Camellia-CCM with 96 bit ICV        |
    |camellia128ccm128  &emsp;|128 bit Camellia-CCM with 128 bit ICV       |
    |camellia192ccm128  &emsp;|192 bit Camellia-CCM with 128 bit ICV       |
    |camellia256ccm128  &emsp;|256 bit Camellia-CCM with 128 bit ICV       |
    |serpent128         &emsp;|128 bit Serpent-CBC                         |
    |serpent192         &emsp;|192 bit Serpent-CBC                         |
    |serpent256         &emsp;|256 bit Serpent-CBC                         |
    |twofish128         &emsp;|128 bit Twofish-CBC                         |
    |twofish192         &emsp;|192 bit Twofish-CBC                         |
    |twofish256         &emsp;|256 bit Twofish-CBC                         |
    |cast128            &emsp;|128 bit CAST-CBC                            |
    |chacha20poly1305   &emsp;|256 bit ChaCha20/Poly1305 with 128 bit ICV  |
- `hash` (String) Hash algorithm

    |Format      &emsp;|Description        |
    |--------------|---------------------|
    |md5         &emsp;|MD5 HMAC           |
    |md5_128     &emsp;|MD5_128 HMAC       |
    |sha1        &emsp;|SHA1 HMAC          |
    |sha1_160    &emsp;|SHA1_160 HMAC      |
    |sha256      &emsp;|SHA2_256_128 HMAC  |
    |sha256_96   &emsp;|SHA2_256_96 HMAC   |
    |sha384      &emsp;|SHA2_384_192 HMAC  |
    |sha512      &emsp;|SHA2_512_256 HMAC  |
    |aesxcbc     &emsp;|AES XCBC           |
    |aescmac     &emsp;|AES CMAC           |
    |aes128gmac  &emsp;|128-bit AES-GMAC   |
    |aes192gmac  &emsp;|192-bit AES-GMAC   |
    |aes256gmac  &emsp;|256-bit AES-GMAC   |
- `prf` (String) Pseudo-Random Functions

    |Format      &emsp;|Description   |
    |--------------|----------------|
    |prfmd5      &emsp;|MD5 PRF       |
    |prfsha1     &emsp;|SHA1 PRF      |
    |prfaesxcbc  &emsp;|AES XCBC PRF  |
    |prfaescmac  &emsp;|AES CMAC PRF  |
    |prfsha256   &emsp;|SHA2_256 PRF  |
    |prfsha384   &emsp;|SHA2_384 PRF  |
    |prfsha512   &emsp;|SHA2_512 PRF  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `ike_group` (String) Internet Key Exchange (IKE) group name
- `proposal` (Number) IKE proposal

    |Format   &emsp;|Description         |
    |-----------|----------------------|
    |1-65535  &emsp;|IKE group proposal  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
