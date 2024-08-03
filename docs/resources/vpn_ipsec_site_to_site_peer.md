---
page_title: "vyos_vpn_ipsec_site_to_site_peer Resource - vyos"

subcategory: "Vpn"

description: |- 
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯Site-to-site VPN⯯Connection name of the peer
---

# vyos_vpn_ipsec_site_to_site_peer (Resource)
<center>

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
Site-to-site VPN  
⯯  
**Connection name of the peer**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `authentication` (Attributes) Peer authentication (see [below for nested schema](#nestedatt--authentication))
- `connection_type` (String) Connection type

    &emsp;|Format    &emsp;|Description                                   |
    |------------|------------------------------------------------|
    &emsp;|initiate  &emsp;|Bring the connection up immediately           |
    &emsp;|respond   &emsp;|Wait for the peer to initiate the connection  |
    &emsp;|none      &emsp;|Load the connection only                      |
- `default_esp_group` (String) Defult ESP group name
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `dhcp_interface` (String) DHCP interface supplying next-hop IP address

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|txt     &emsp;|DHCP interface name  |
- `disable` (Boolean) Disable instance
- `force_udp_encapsulation` (Boolean) Force UDP encapsulation
- `ike_group` (String) Internet Key Exchange (IKE) group name
- `ikev2_reauth` (String) Re-authentication of the remote peer during an IKE re-key (IKEv2 only)

    &emsp;|Format   &emsp;|Description                                                                                          |
    |-----------|-------------------------------------------------------------------------------------------------------|
    &emsp;|yes      &emsp;|Enable remote host re-autentication during an IKE re-key. Currently broken due to a strong swan bug  |
    &emsp;|no       &emsp;|Disable remote host re-authenticaton during an IKE re-key.                                           |
    &emsp;|inherit  &emsp;|Inherit the reauth configuration form your IKE-group                                                 |
- `local_address` (String) IPv4 or IPv6 address of a local interface to use for VPN

    &emsp;|Format  &emsp;|Description                                                      |
    |----------|-------------------------------------------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address of a local interface for VPN                        |
    &emsp;|ipv6    &emsp;|IPv6 address of a local interface for VPN                        |
    &emsp;|any     &emsp;|Allow any IPv4 address present on the system to be used for VPN  |
- `remote_address` (List of String) IPv4 or IPv6 address of the remote peer

    &emsp;|Format    &emsp;|Description                                     |
    |------------|--------------------------------------------------|
    &emsp;|ipv4      &emsp;|IPv4 address of the remote peer                 |
    &emsp;|ipv6      &emsp;|IPv6 address of the remote peer                 |
    &emsp;|hostname  &emsp;|Fully qualified domain name of the remote peer  |
    &emsp;|any       &emsp;|Allow any IP address of the remote peer         |
- `replay_window` (Number) IPsec replay window to configure for this CHILD_SA

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|0       &emsp;|Disable IPsec replay protection  |
    &emsp;|1-2040  &emsp;|Replay window size in packets    |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `virtual_address` (List of String) Initiator request virtual-address from peer

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv4    &emsp;|Request IPv4 address from peer  |
    &emsp;|ipv6    &emsp;|Request IPv6 address from peer  |
- `vti` (Attributes) Virtual tunnel interface (see [below for nested schema](#nestedatt--vti))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `peer` (String) Connection name of the peer

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|txt     &emsp;|Connection name of the peer  |


&lt;a id=&#34;nestedatt--authentication&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication`

Optional:

- `local_id` (String) Local ID for peer authentication

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|txt     &emsp;|Local ID used for peer authentication  |
- `mode` (String) Authentication mode

    &emsp;|Format             &emsp;|Description                |
    |---------------------|-----------------------------|
    &emsp;|pre-shared-secret  &emsp;|Use pre-shared secret key  |
    &emsp;|rsa                &emsp;|Use RSA key                |
    &emsp;|x509               &emsp;|Use x.509 certificate      |
- `remote_id` (String) ID for remote authentication

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|txt     &emsp;|ID used for peer authentication  |
- `rsa` (Attributes) RSA keys (see [below for nested schema](#nestedatt--authentication--rsa))
- `use_x509_id` (Boolean) Use certificate common name as ID
- `x509` (Attributes) X.509 certificate (see [below for nested schema](#nestedatt--authentication--x509))

&lt;a id=&#34;nestedatt--authentication--rsa&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication.rsa`

Optional:

- `local_key` (String) Name of PKI key-pair with local private key
- `passphrase` (String) Local private key passphrase
- `remote_key` (String) Name of PKI key-pair with remote public key


&lt;a id=&#34;nestedatt--authentication--x509&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication.x509`

Optional:

- `ca_certificate` (List of String) Certificate Authority chain in PKI configuration

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|txt     &emsp;|Name of CA in PKI configuration  |
- `certificate` (String) Certificate in PKI configuration

    &emsp;|Format  &emsp;|Description                               |
    |----------|--------------------------------------------|
    &emsp;|txt     &emsp;|Name of certificate in PKI configuration  |
- `passphrase` (String) Private key passphrase

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|txt     &emsp;|Passphrase to decrypt the private key  |



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


&lt;a id=&#34;nestedatt--vti&#34;&gt;&lt;/a&gt;
### Nested Schema for `vti`

Optional:

- `bind` (String) VTI tunnel interface associated with this configuration
- `esp_group` (String) Encapsulating Security Payloads (ESP) group name  
