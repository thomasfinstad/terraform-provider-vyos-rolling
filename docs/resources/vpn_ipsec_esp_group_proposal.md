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

<!--TOC-->

- [vyos_vpn_ipsec_esp_group_proposal (Resource)](#vyos_vpn_ipsec_esp_group_proposal-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [encryption](#encryption)
      - [hash](#hash)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### encryption
- `encryption` (String) Encryption algorithm

    |  Format             &emsp;|  Description                                 |
    |---------------------|----------------------------------------------|
    |  null               &emsp;|  Null encryption                             |
    |  aes128             &emsp;|  128 bit AES-CBC                             |
    |  aes192             &emsp;|  192 bit AES-CBC                             |
    |  aes256             &emsp;|  256 bit AES-CBC                             |
    |  aes128ctr          &emsp;|  128 bit AES-COUNTER                         |
    |  aes192ctr          &emsp;|  192 bit AES-COUNTER                         |
    |  aes256ctr          &emsp;|  256 bit AES-COUNTER                         |
    |  aes128ccm64        &emsp;|  128 bit AES-CCM with 64 bit ICV             |
    |  aes192ccm64        &emsp;|  192 bit AES-CCM with 64 bit ICV             |
    |  aes256ccm64        &emsp;|  256 bit AES-CCM with 64 bit ICV             |
    |  aes128ccm96        &emsp;|  128 bit AES-CCM with 96 bit ICV             |
    |  aes192ccm96        &emsp;|  192 bit AES-CCM with 96 bit ICV             |
    |  aes256ccm96        &emsp;|  256 bit AES-CCM with 96 bit ICV             |
    |  aes128ccm128       &emsp;|  128 bit AES-CCM with 128 bit ICV            |
    |  aes192ccm128       &emsp;|  192 bit AES-CCM with 128 bit IC             |
    |  aes256ccm128       &emsp;|  256 bit AES-CCM with 128 bit ICV            |
    |  aes128gcm64        &emsp;|  128 bit AES-GCM with 64 bit ICV             |
    |  aes192gcm64        &emsp;|  192 bit AES-GCM with 64 bit ICV             |
    |  aes256gcm64        &emsp;|  256 bit AES-GCM with 64 bit ICV             |
    |  aes128gcm96        &emsp;|  128 bit AES-GCM with 96 bit ICV             |
    |  aes192gcm96        &emsp;|  192 bit AES-GCM with 96 bit ICV             |
    |  aes256gcm96        &emsp;|  256 bit AES-GCM with 96 bit ICV             |
    |  aes128gcm128       &emsp;|  128 bit AES-GCM with 128 bit ICV            |
    |  aes192gcm128       &emsp;|  192 bit AES-GCM with 128 bit ICV            |
    |  aes256gcm128       &emsp;|  256 bit AES-GCM with 128 bit ICV            |
    |  aes128gmac         &emsp;|  Null encryption with 128 bit AES-GMAC       |
    |  aes192gmac         &emsp;|  Null encryption with 192 bit AES-GMAC       |
    |  aes256gmac         &emsp;|  Null encryption with 256 bit AES-GMAC       |
    |  3des               &emsp;|  168 bit 3DES-EDE-CBC                        |
    |  blowfish128        &emsp;|  128 bit Blowfish-CBC                        |
    |  blowfish192        &emsp;|  192 bit Blowfish-CBC                        |
    |  blowfish256        &emsp;|  256 bit Blowfish-CBC                        |
    |  camellia128        &emsp;|  128 bit Camellia-CBC                        |
    |  camellia192        &emsp;|  192 bit Camellia-CBC                        |
    |  camellia256        &emsp;|  256 bit Camellia-CBC                        |
    |  camellia128ctr     &emsp;|  128 bit Camellia-COUNTER                    |
    |  camellia192ctr     &emsp;|  192 bit Camellia-COUNTER                    |
    |  camellia256ctr     &emsp;|  256 bit Camellia-COUNTER                    |
    |  camellia128ccm64   &emsp;|  128 bit Camellia-CCM with 64 bit ICV        |
    |  camellia192ccm64   &emsp;|  192 bit Camellia-CCM with 64 bit ICV        |
    |  camellia256ccm64   &emsp;|  256 bit Camellia-CCM with 64 bit ICV        |
    |  camellia128ccm96   &emsp;|  128 bit Camellia-CCM with 96 bit ICV        |
    |  camellia192ccm96   &emsp;|  192 bit Camellia-CCM with 96 bit ICV        |
    |  camellia256ccm96   &emsp;|  256 bit Camellia-CCM with 96 bit ICV        |
    |  camellia128ccm128  &emsp;|  128 bit Camellia-CCM with 128 bit ICV       |
    |  camellia192ccm128  &emsp;|  192 bit Camellia-CCM with 128 bit ICV       |
    |  camellia256ccm128  &emsp;|  256 bit Camellia-CCM with 128 bit ICV       |
    |  serpent128         &emsp;|  128 bit Serpent-CBC                         |
    |  serpent192         &emsp;|  192 bit Serpent-CBC                         |
    |  serpent256         &emsp;|  256 bit Serpent-CBC                         |
    |  twofish128         &emsp;|  128 bit Twofish-CBC                         |
    |  twofish192         &emsp;|  192 bit Twofish-CBC                         |
    |  twofish256         &emsp;|  256 bit Twofish-CBC                         |
    |  cast128            &emsp;|  128 bit CAST-CBC                            |
    |  chacha20poly1305   &emsp;|  256 bit ChaCha20/Poly1305 with 128 bit ICV  |
#### hash
- `hash` (String) Hash algorithm

    |  Format      &emsp;|  Description        |
    |--------------|---------------------|
    |  md5         &emsp;|  MD5 HMAC           |
    |  md5_128     &emsp;|  MD5_128 HMAC       |
    |  sha1        &emsp;|  SHA1 HMAC          |
    |  sha1_160    &emsp;|  SHA1_160 HMAC      |
    |  sha256      &emsp;|  SHA2_256_128 HMAC  |
    |  sha256_96   &emsp;|  SHA2_256_96 HMAC   |
    |  sha384      &emsp;|  SHA2_384_192 HMAC  |
    |  sha512      &emsp;|  SHA2_512_256 HMAC  |
    |  aesxcbc     &emsp;|  AES XCBC           |
    |  aescmac     &emsp;|  AES CMAC           |
    |  aes128gmac  &emsp;|  128-bit AES-GMAC   |
    |  aes192gmac  &emsp;|  192-bit AES-GMAC   |
    |  aes256gmac  &emsp;|  256-bit AES-GMAC   |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `esp_group` (String) Encapsulating Security Payload (ESP) group name
- `proposal` (Number) ESP group proposal

    |  Format   &emsp;|  Description                |
    |-----------|-----------------------------|
    |  1-65535  &emsp;|  ESP group proposal number  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_vpn_ipsec_esp_group_proposal.example "vpn__ipsec__esp-group__<esp-group>__proposal__<proposal>"
```
