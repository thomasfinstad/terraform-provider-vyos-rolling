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

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|1       &emsp;|Diffie-Hellman group 1 (modp768)        |
    &emsp;|2       &emsp;|Diffie-Hellman group 2 (modp1024)       |
    &emsp;|5       &emsp;|Diffie-Hellman group 5 (modp1536)       |
    &emsp;|14      &emsp;|Diffie-Hellman group 14 (modp2048)      |
    &emsp;|15      &emsp;|Diffie-Hellman group 15 (modp3072)      |
    &emsp;|16      &emsp;|Diffie-Hellman group 16 (modp4096)      |
    &emsp;|17      &emsp;|Diffie-Hellman group 17 (modp6144)      |
    &emsp;|18      &emsp;|Diffie-Hellman group 18 (modp8192)      |
    &emsp;|19      &emsp;|Diffie-Hellman group 19 (ecp256)        |
    &emsp;|20      &emsp;|Diffie-Hellman group 20 (ecp384)        |
    &emsp;|21      &emsp;|Diffie-Hellman group 21 (ecp521)        |
    &emsp;|22      &emsp;|Diffie-Hellman group 22 (modp1024s160)  |
    &emsp;|23      &emsp;|Diffie-Hellman group 23 (modp2048s224)  |
    &emsp;|24      &emsp;|Diffie-Hellman group 24 (modp2048s256)  |
    &emsp;|25      &emsp;|Diffie-Hellman group 25 (ecp192)        |
    &emsp;|26      &emsp;|Diffie-Hellman group 26 (ecp224)        |
    &emsp;|27      &emsp;|Diffie-Hellman group 27 (ecp224bp)      |
    &emsp;|28      &emsp;|Diffie-Hellman group 28 (ecp256bp)      |
    &emsp;|29      &emsp;|Diffie-Hellman group 29 (ecp384bp)      |
    &emsp;|30      &emsp;|Diffie-Hellman group 30 (ecp512bp)      |
    &emsp;|31      &emsp;|Diffie-Hellman group 31 (curve25519)    |
    &emsp;|32      &emsp;|Diffie-Hellman group 32 (curve448)      |
- `encryption` (String) Encryption algorithm

    &emsp;|Format             &emsp;|Description                                 |
    |---------------------|----------------------------------------------|
    &emsp;|null               &emsp;|Null encryption                             |
    &emsp;|aes128             &emsp;|128 bit AES-CBC                             |
    &emsp;|aes192             &emsp;|192 bit AES-CBC                             |
    &emsp;|aes256             &emsp;|256 bit AES-CBC                             |
    &emsp;|aes128ctr          &emsp;|128 bit AES-COUNTER                         |
    &emsp;|aes192ctr          &emsp;|192 bit AES-COUNTER                         |
    &emsp;|aes256ctr          &emsp;|256 bit AES-COUNTER                         |
    &emsp;|aes128ccm64        &emsp;|128 bit AES-CCM with 64 bit ICV             |
    &emsp;|aes192ccm64        &emsp;|192 bit AES-CCM with 64 bit ICV             |
    &emsp;|aes256ccm64        &emsp;|256 bit AES-CCM with 64 bit ICV             |
    &emsp;|aes128ccm96        &emsp;|128 bit AES-CCM with 96 bit ICV             |
    &emsp;|aes192ccm96        &emsp;|192 bit AES-CCM with 96 bit ICV             |
    &emsp;|aes256ccm96        &emsp;|256 bit AES-CCM with 96 bit ICV             |
    &emsp;|aes128ccm128       &emsp;|128 bit AES-CCM with 128 bit ICV            |
    &emsp;|aes192ccm128       &emsp;|192 bit AES-CCM with 128 bit IC             |
    &emsp;|aes256ccm128       &emsp;|256 bit AES-CCM with 128 bit ICV            |
    &emsp;|aes128gcm64        &emsp;|128 bit AES-GCM with 64 bit ICV             |
    &emsp;|aes192gcm64        &emsp;|192 bit AES-GCM with 64 bit ICV             |
    &emsp;|aes256gcm64        &emsp;|256 bit AES-GCM with 64 bit ICV             |
    &emsp;|aes128gcm96        &emsp;|128 bit AES-GCM with 96 bit ICV             |
    &emsp;|aes192gcm96        &emsp;|192 bit AES-GCM with 96 bit ICV             |
    &emsp;|aes256gcm96        &emsp;|256 bit AES-GCM with 96 bit ICV             |
    &emsp;|aes128gcm128       &emsp;|128 bit AES-GCM with 128 bit ICV            |
    &emsp;|aes192gcm128       &emsp;|192 bit AES-GCM with 128 bit ICV            |
    &emsp;|aes256gcm128       &emsp;|256 bit AES-GCM with 128 bit ICV            |
    &emsp;|aes128gmac         &emsp;|Null encryption with 128 bit AES-GMAC       |
    &emsp;|aes192gmac         &emsp;|Null encryption with 192 bit AES-GMAC       |
    &emsp;|aes256gmac         &emsp;|Null encryption with 256 bit AES-GMAC       |
    &emsp;|3des               &emsp;|168 bit 3DES-EDE-CBC                        |
    &emsp;|blowfish128        &emsp;|128 bit Blowfish-CBC                        |
    &emsp;|blowfish192        &emsp;|192 bit Blowfish-CBC                        |
    &emsp;|blowfish256        &emsp;|256 bit Blowfish-CBC                        |
    &emsp;|camellia128        &emsp;|128 bit Camellia-CBC                        |
    &emsp;|camellia192        &emsp;|192 bit Camellia-CBC                        |
    &emsp;|camellia256        &emsp;|256 bit Camellia-CBC                        |
    &emsp;|camellia128ctr     &emsp;|128 bit Camellia-COUNTER                    |
    &emsp;|camellia192ctr     &emsp;|192 bit Camellia-COUNTER                    |
    &emsp;|camellia256ctr     &emsp;|256 bit Camellia-COUNTER                    |
    &emsp;|camellia128ccm64   &emsp;|128 bit Camellia-CCM with 64 bit ICV        |
    &emsp;|camellia192ccm64   &emsp;|192 bit Camellia-CCM with 64 bit ICV        |
    &emsp;|camellia256ccm64   &emsp;|256 bit Camellia-CCM with 64 bit ICV        |
    &emsp;|camellia128ccm96   &emsp;|128 bit Camellia-CCM with 96 bit ICV        |
    &emsp;|camellia192ccm96   &emsp;|192 bit Camellia-CCM with 96 bit ICV        |
    &emsp;|camellia256ccm96   &emsp;|256 bit Camellia-CCM with 96 bit ICV        |
    &emsp;|camellia128ccm128  &emsp;|128 bit Camellia-CCM with 128 bit ICV       |
    &emsp;|camellia192ccm128  &emsp;|192 bit Camellia-CCM with 128 bit ICV       |
    &emsp;|camellia256ccm128  &emsp;|256 bit Camellia-CCM with 128 bit ICV       |
    &emsp;|serpent128         &emsp;|128 bit Serpent-CBC                         |
    &emsp;|serpent192         &emsp;|192 bit Serpent-CBC                         |
    &emsp;|serpent256         &emsp;|256 bit Serpent-CBC                         |
    &emsp;|twofish128         &emsp;|128 bit Twofish-CBC                         |
    &emsp;|twofish192         &emsp;|192 bit Twofish-CBC                         |
    &emsp;|twofish256         &emsp;|256 bit Twofish-CBC                         |
    &emsp;|cast128            &emsp;|128 bit CAST-CBC                            |
    &emsp;|chacha20poly1305   &emsp;|256 bit ChaCha20/Poly1305 with 128 bit ICV  |
- `hash` (String) Hash algorithm

    &emsp;|Format      &emsp;|Description        |
    |--------------|---------------------|
    &emsp;|md5         &emsp;|MD5 HMAC           |
    &emsp;|md5_128     &emsp;|MD5_128 HMAC       |
    &emsp;|sha1        &emsp;|SHA1 HMAC          |
    &emsp;|sha1_160    &emsp;|SHA1_160 HMAC      |
    &emsp;|sha256      &emsp;|SHA2_256_128 HMAC  |
    &emsp;|sha256_96   &emsp;|SHA2_256_96 HMAC   |
    &emsp;|sha384      &emsp;|SHA2_384_192 HMAC  |
    &emsp;|sha512      &emsp;|SHA2_512_256 HMAC  |
    &emsp;|aesxcbc     &emsp;|AES XCBC           |
    &emsp;|aescmac     &emsp;|AES CMAC           |
    &emsp;|aes128gmac  &emsp;|128-bit AES-GMAC   |
    &emsp;|aes192gmac  &emsp;|192-bit AES-GMAC   |
    &emsp;|aes256gmac  &emsp;|256-bit AES-GMAC   |
- `prf` (String) Pseudo-Random Functions

    &emsp;|Format      &emsp;|Description   |
    |--------------|----------------|
    &emsp;|prfmd5      &emsp;|MD5 PRF       |
    &emsp;|prfsha1     &emsp;|SHA1 PRF      |
    &emsp;|prfaesxcbc  &emsp;|AES XCBC PRF  |
    &emsp;|prfaescmac  &emsp;|AES CMAC PRF  |
    &emsp;|prfsha256   &emsp;|SHA2_256 PRF  |
    &emsp;|prfsha384   &emsp;|SHA2_384 PRF  |
    &emsp;|prfsha512   &emsp;|SHA2_512 PRF  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `ike_group` (String) Internet Key Exchange (IKE) group name
- `proposal` (Number) IKE proposal

    &emsp;|Format   &emsp;|Description         |
    |-----------|----------------------|
    &emsp;|1-65535  &emsp;|IKE group proposal  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
