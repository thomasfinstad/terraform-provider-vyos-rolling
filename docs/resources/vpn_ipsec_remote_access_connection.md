---
page_title: "vyos_vpn_ipsec_remote_access_connection Resource - vyos"

subcategory: "Vpn"

description: |- 
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯IKEv2 remote access VPN⯯IKEv2 VPN connection name
---

# vyos_vpn_ipsec_remote_access_connection (Resource)
<center>

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
IKEv2 remote access VPN  
⯯  
**IKEv2 VPN connection name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `authentication` (Attributes) Authentication for remote access (see [below for nested schema](#nestedatt--authentication))
- `bind` (String) VTI tunnel interface associated with this configuration
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `dhcp_interface` (String) DHCP interface supplying next-hop IP address

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|txt     &emsp;|DHCP interface name  |
- `disable` (Boolean) Disable instance
- `esp_group` (String) Encapsulating Security Payloads (ESP) group name
- `ike_group` (String) Internet Key Exchange (IKE) group name
- `local` (Attributes) Local parameters for interesting traffic (see [below for nested schema](#nestedatt--local))
- `local_address` (String) IPv4 or IPv6 address of a local interface to use for VPN

    &emsp;|Format  &emsp;|Description                                                      |
    |----------|-------------------------------------------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address of a local interface for VPN                        |
    &emsp;|ipv6    &emsp;|IPv6 address of a local interface for VPN                        |
    &emsp;|any     &emsp;|Allow any IPv4 address present on the system to be used for VPN  |
- `pool` (List of String) IP address pool

    &emsp;|Format  &emsp;|Description                                                   |
    |----------|----------------------------------------------------------------|
    &emsp;|txt     &emsp;|Predefined IP pool name                                       |
    &emsp;|dhcp    &emsp;|Forward requests for virtual IP addresses to a DHCP server    |
    &emsp;|radius  &emsp;|Forward requests for virtual IP addresses to a RADIUS server  |
- `replay_window` (Number) IPsec replay window to configure for this CHILD_SA

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|0       &emsp;|Disable IPsec replay protection  |
    &emsp;|1-2040  &emsp;|Replay window size in packets    |
- `timeout` (Number) Timeout to close connection if no data is transmitted

    &emsp;|Format   &emsp;|Description                |
    |-----------|-----------------------------|
    &emsp;|0        &emsp;|Disable inactivity checks  |
    &emsp;|1-86400  &emsp;|Timeout in seconds         |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `unique` (String) Connection uniqueness enforcement policy

    &emsp;|Format   &emsp;|Description                                                                       |
    |-----------|------------------------------------------------------------------------------------|
    &emsp;|never    &emsp;|Never enforce connection uniqueness                                               |
    &emsp;|keep     &emsp;|Reject new connection attempts if the same user already has an active connection  |
    &emsp;|replace  &emsp;|Delete any existing connection if a new one for the same user gets established    |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `connection` (String) IKEv2 VPN connection name

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|txt     &emsp;|Connection name  |


&lt;a id=&#34;nestedatt--authentication&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication`

Optional:

- `client_mode` (String) Client authentication mode

    &emsp;|Format        &emsp;|Description                                 |
    |----------------|----------------------------------------------|
    &emsp;|x509          &emsp;|Use IPsec x.509 certificate authentication  |
    &emsp;|eap-tls       &emsp;|Use EAP-TLS authentication                  |
    &emsp;|eap-mschapv2  &emsp;|Use EAP-MSCHAPv2 authentication             |
    &emsp;|eap-radius    &emsp;|Use EAP-RADIUS authentication               |
- `eap_id` (String) Remote EAP ID for client authentication

    &emsp;|Format  &emsp;|Description                              |
    |----------|-------------------------------------------|
    &emsp;|txt     &emsp;|Remote EAP ID for client authentication  |
    &emsp;|any     &emsp;|Allow any EAP ID                         |
- `local_id` (String) Local ID for peer authentication

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|txt     &emsp;|Local ID used for peer authentication  |
- `local_users` (Attributes) Local user authentication (see [below for nested schema](#nestedatt--authentication--local_users))
- `pre_shared_secret` (String) Pre-shared secret key

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|txt     &emsp;|Pre-shared secret key  |
- `server_mode` (String) Server authentication mode

    &emsp;|Format             &emsp;|Description                  |
    |---------------------|-------------------------------|
    &emsp;|pre-shared-secret  &emsp;|Use a pre-shared secret key  |
    &emsp;|x509               &emsp;|Use x.509 certificate        |
- `x509` (Attributes) X.509 certificate (see [below for nested schema](#nestedatt--authentication--x509))

&lt;a id=&#34;nestedatt--authentication--local_users&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication.local_users`


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



&lt;a id=&#34;nestedatt--local&#34;&gt;&lt;/a&gt;
### Nested Schema for `local`

Optional:

- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `prefix` (List of String) Local IPv4 or IPv6 prefix

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv4net  &emsp;|Local IPv4 prefix  |
    &emsp;|ipv6net  &emsp;|Local IPv6 prefix  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
