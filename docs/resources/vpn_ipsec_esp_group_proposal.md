---
page_title: "vyos_vpn_ipsec_esp_group_proposal Resource - vyos"

subcategory: "Vpn"

description: |- 
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯Encapsulating Security Payload (ESP) group name⯯ESP group proposal
---

# vyos_vpn_ipsec_esp_group_proposal (Resource)
<center>

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
Encapsulating Security Payload (ESP) group name  
⯯  
**ESP group proposal**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

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
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `esp_group` (String) Encapsulating Security Payload (ESP) group name
- `proposal` (Number) ESP group proposal

    &emsp;|Format   &emsp;|Description                |
    |-----------|-----------------------------|
    &emsp;|1-65535  &emsp;|ESP group proposal number  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
