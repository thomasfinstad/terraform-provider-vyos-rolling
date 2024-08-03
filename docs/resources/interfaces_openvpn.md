---
page_title: "vyos_interfaces_openvpn Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯OpenVPN Tunnel Interface
---

# vyos_interfaces_openvpn (Resource)
<center>

*interfaces*  
⯯  
**OpenVPN Tunnel Interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `authentication` (Attributes) Authentication settings (see [below for nested schema](#nestedatt--authentication))
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `device_type` (String) OpenVPN interface device-type

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|tun     &emsp;|TUN device, required for OSI layer 3  |
    &emsp;|tap     &emsp;|TAP device, required for OSI layer 2  |
- `disable` (Boolean) Administratively disable interface
- `encryption` (Attributes) Data Encryption settings (see [below for nested schema](#nestedatt--encryption))
- `hash` (String) Hashing Algorithm

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|md5     &emsp;|MD5 algorithm      |
    &emsp;|sha1    &emsp;|SHA-1 algorithm    |
    &emsp;|sha256  &emsp;|SHA-256 algorithm  |
    &emsp;|sha384  &emsp;|SHA-384 algorithm  |
    &emsp;|sha512  &emsp;|SHA-512 algorithm  |
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
- `ipv6` (Attributes) IPv6 routing parameters (see [below for nested schema](#nestedatt--ipv6))
- `keep_alive` (Attributes) Keepalive helper options (see [below for nested schema](#nestedatt--keep_alive))
- `local_host` (String) Local IP address to accept connections (all if not set)

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|ipv4    &emsp;|Local IPv4 address  |
    &emsp;|ipv6    &emsp;|Local IPv6 address  |
- `local_port` (Number) Local port number to accept connections

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `mirror` (Attributes) Mirror ingress/egress packets (see [below for nested schema](#nestedatt--mirror))
- `mode` (String) OpenVPN mode of operation

    &emsp;|Format        &emsp;|Description                   |
    |----------------|--------------------------------|
    &emsp;|site-to-site  &emsp;|Site-to-site mode             |
    &emsp;|client        &emsp;|Client in client-server mode  |
    &emsp;|server        &emsp;|Server in client-server mode  |
- `offload` (Attributes) Configurable offload options (see [below for nested schema](#nestedatt--offload))
- `openvpn_option` (List of String) Additional OpenVPN options. You must use the syntax of openvpn.conf in this text-field. Using this without proper knowledge may result in a crashed OpenVPN server. Check system log to look for errors.
- `persistent_tunnel` (Boolean) Do not close and reopen interface (TUN/TAP device) on client restarts
- `protocol` (String) OpenVPN communication protocol

    &emsp;|Format       &emsp;|Description                             |
    |---------------|------------------------------------------|
    &emsp;|udp          &emsp;|UDP                                     |
    &emsp;|tcp-passive  &emsp;|TCP and accepts connections passively   |
    &emsp;|tcp-active   &emsp;|TCP and initiates connections actively  |
- `redirect` (String) Redirect incoming packet to destination

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|txt     &emsp;|Destination interface name  |
- `remote_address` (List of String) IP address of remote end of tunnel

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|ipv4    &emsp;|Remote end IPv4 address  |
    &emsp;|ipv6    &emsp;|Remote end IPv6 address  |
- `remote_host` (List of String) Remote host to connect to (dynamic if not set)

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address of remote host  |
    &emsp;|ipv6    &emsp;|IPv6 address of remote host  |
    &emsp;|txt     &emsp;|Hostname of remote host      |
- `remote_port` (Number) Remote port number to connect to

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `replace_default_route` (Attributes) OpenVPN tunnel to be used as the default route (see [below for nested schema](#nestedatt--replace_default_route))
- `server` (Attributes) Server-mode options (see [below for nested schema](#nestedatt--server))
- `shared_secret_key` (String) Secret key shared with remote end of tunnel
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `tls` (Attributes) Transport Layer Security (TLS) options (see [below for nested schema](#nestedatt--tls))
- `use_lzo_compression` (Boolean) Use fast LZO compression on this TUN/TAP interface
- `vrf` (String) VRF instance name

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `openvpn` (String) OpenVPN Tunnel Interface

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|vtunN   &emsp;|OpenVPN interface name  |


&lt;a id=&#34;nestedatt--authentication&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication`

Optional:

- `password` (String) Password used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Password     |
- `username` (String) Username used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Username     |


&lt;a id=&#34;nestedatt--encryption&#34;&gt;&lt;/a&gt;
### Nested Schema for `encryption`

Optional:

- `cipher` (String) Standard Data Encryption Algorithm

    &emsp;|Format     &emsp;|Description                           |
    |-------------|----------------------------------------|
    &emsp;|none       &emsp;|Disable encryption                    |
    &emsp;|3des       &emsp;|DES algorithm with triple encryption  |
    &emsp;|aes128     &emsp;|AES algorithm with 128-bit key CBC    |
    &emsp;|aes128gcm  &emsp;|AES algorithm with 128-bit key GCM    |
    &emsp;|aes192     &emsp;|AES algorithm with 192-bit key CBC    |
    &emsp;|aes192gcm  &emsp;|AES algorithm with 192-bit key GCM    |
    &emsp;|aes256     &emsp;|AES algorithm with 256-bit key CBC    |
    &emsp;|aes256gcm  &emsp;|AES algorithm with 256-bit key GCM    |
- `data_ciphers` (List of String) Cipher negotiation list for use in server or client mode

    &emsp;|Format     &emsp;|Description                           |
    |-------------|----------------------------------------|
    &emsp;|none       &emsp;|Disable encryption                    |
    &emsp;|3des       &emsp;|DES algorithm with triple encryption  |
    &emsp;|aes128     &emsp;|AES algorithm with 128-bit key CBC    |
    &emsp;|aes128gcm  &emsp;|AES algorithm with 128-bit key GCM    |
    &emsp;|aes192     &emsp;|AES algorithm with 192-bit key CBC    |
    &emsp;|aes192gcm  &emsp;|AES algorithm with 192-bit key GCM    |
    &emsp;|aes256     &emsp;|AES algorithm with 256-bit key CBC    |
    &emsp;|aes256gcm  &emsp;|AES algorithm with 256-bit key GCM    |


&lt;a id=&#34;nestedatt--ip&#34;&gt;&lt;/a&gt;
### Nested Schema for `ip`

Optional:

- `adjust_mss` (String) Adjust TCP MSS value

    &emsp;|Format             &emsp;|Description                                     |
    |---------------------|--------------------------------------------------|
    &emsp;|clamp-mss-to-pmtu  &emsp;|Automatically sets the MSS to the proper value  |
    &emsp;|536-65535          &emsp;|TCP Maximum segment size in bytes               |
- `arp_cache_timeout` (Number) ARP cache entry timeout in seconds

    &emsp;|Format   &emsp;|Description                        |
    |-----------|-------------------------------------|
    &emsp;|1-86400  &emsp;|ARP cache entry timout in seconds  |
- `disable_arp_filter` (Boolean) Disable ARP filter on this interface
- `disable_forwarding` (Boolean) Disable IP forwarding on this interface
- `enable_arp_accept` (Boolean) Enable ARP accept on this interface
- `enable_arp_announce` (Boolean) Enable ARP announce on this interface
- `enable_arp_ignore` (Boolean) Enable ARP ignore on this interface
- `enable_directed_broadcast` (Boolean) Enable directed broadcast forwarding on this interface
- `enable_proxy_arp` (Boolean) Enable proxy-arp on this interface
- `proxy_arp_pvlan` (Boolean) Enable private VLAN proxy ARP on this interface
- `source_validation` (String) Source validation by reversed path (RFC3704)

    &emsp;|Format   &emsp;|Description                                                  |
    |-----------|---------------------------------------------------------------|
    &emsp;|strict   &emsp;|Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    &emsp;|loose    &emsp;|Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    &emsp;|disable  &emsp;|No source validation                                         |


&lt;a id=&#34;nestedatt--ipv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6`

Optional:

- `accept_dad` (String) Accept Duplicate Address Detection

    &emsp;|Format  &emsp;|Description                                                                |
    |----------|-----------------------------------------------------------------------------|
    &emsp;|0       &emsp;|Disable DAD                                                                |
    &emsp;|1       &emsp;|Enable DAD                                                                 |
    &emsp;|2       &emsp;|Enable DAD - disable IPv6 if MAC-based duplicate link-local address found  |
- `address` (Attributes) IPv6 address configuration modes (see [below for nested schema](#nestedatt--ipv6--address))
- `adjust_mss` (String) Adjust TCP MSS value

    &emsp;|Format             &emsp;|Description                                     |
    |---------------------|--------------------------------------------------|
    &emsp;|clamp-mss-to-pmtu  &emsp;|Automatically sets the MSS to the proper value  |
    &emsp;|536-65535          &emsp;|TCP Maximum segment size in bytes               |
- `base_reachable_time` (Number) Base reachable time in seconds

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|1-86400  &emsp;|Base reachable time in seconds  |
- `disable_forwarding` (Boolean) Disable IP forwarding on this interface
- `dup_addr_detect_transmits` (Number) Number of NS messages to send while performing DAD

    &emsp;|Format  &emsp;|Description                                         |
    |----------|------------------------------------------------------|
    &emsp;|0       &emsp;|Disable Duplicate Address Dectection (DAD)          |
    &emsp;|1-n     &emsp;|Number of NS messages to send while performing DAD  |
- `source_validation` (String) Source validation by reversed path (RFC3704)

    &emsp;|Format   &emsp;|Description                                                  |
    |-----------|---------------------------------------------------------------|
    &emsp;|strict   &emsp;|Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    &emsp;|loose    &emsp;|Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    &emsp;|disable  &emsp;|No source validation                                         |

&lt;a id=&#34;nestedatt--ipv6--address&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6.address`

Optional:

- `autoconf` (Boolean) Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)
- `eui64` (List of String) Prefix for IPv6 address with MAC-based EUI-64

    &emsp;|Format                &emsp;|Description       |
    |------------------------|--------------------|
    &emsp;|&lt;h:h:h:h:h:h:h:h/64&gt;  &emsp;|IPv6 /64 network  |
- `no_default_link_local` (Boolean) Remove the default link-local address from the interface



&lt;a id=&#34;nestedatt--keep_alive&#34;&gt;&lt;/a&gt;
### Nested Schema for `keep_alive`

Optional:

- `failure_count` (Number) Maximum number of keepalive packet failures

    &emsp;|Format  &emsp;|Description                                  |
    |----------|-----------------------------------------------|
    &emsp;|0-1000  &emsp;|Maximum number of keepalive packet failures  |
- `interval` (Number) Keepalive packet interval in seconds

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|0-600   &emsp;|Keepalive packet interval (seconds)  |


&lt;a id=&#34;nestedatt--mirror&#34;&gt;&lt;/a&gt;
### Nested Schema for `mirror`

Optional:

- `egress` (String) Mirror egress traffic to destination interface

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|txt     &emsp;|Destination interface name  |
- `ingress` (String) Mirror ingress traffic to destination interface

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|txt     &emsp;|Destination interface name  |


&lt;a id=&#34;nestedatt--offload&#34;&gt;&lt;/a&gt;
### Nested Schema for `offload`

Optional:

- `dco` (Boolean) Enable data channel offload on this interface


&lt;a id=&#34;nestedatt--replace_default_route&#34;&gt;&lt;/a&gt;
### Nested Schema for `replace_default_route`

Optional:

- `local` (String) Tunnel endpoints are on the same subnet


&lt;a id=&#34;nestedatt--server&#34;&gt;&lt;/a&gt;
### Nested Schema for `server`

Optional:

- `client_ip_pool` (Attributes) Pool of client IPv4 addresses (see [below for nested schema](#nestedatt--server--client_ip_pool))
- `client_ipv6_pool` (Attributes) Pool of client IPv6 addresses (see [below for nested schema](#nestedatt--server--client_ipv6_pool))
- `domain_name` (String) DNS suffix to be pushed to all clients

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|txt     &emsp;|Domain Name Server suffix  |
- `max_connections` (Number) Number of maximum client connections

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|1-4096  &emsp;|Number of concurrent clients  |
- `mfa` (Attributes) multi-factor authentication (see [below for nested schema](#nestedatt--server--mfa))
- `name_server` (List of String) Domain Name Servers (DNS) addresses

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|ipv4    &emsp;|Domain Name Server (DNS) IPv4 address  |
    &emsp;|ipv6    &emsp;|Domain Name Server (DNS) IPv6 address  |
- `reject_unconfigured_clients` (Boolean) Reject connections from clients that are not explicitly configured
- `subnet` (List of String) Server-mode subnet (from which client IPs are allocated)

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 network and prefix length  |
    &emsp;|ipv6net  &emsp;|IPv6 network and prefix length  |
- `topology` (String) Topology for clients

    &emsp;|Format          &emsp;|Description                    |
    |------------------|---------------------------------|
    &emsp;|subnet          &emsp;|Subnet topology (recommended)  |
    &emsp;|point-to-point  &emsp;|Point-to-point topology        |
    &emsp;|net30           &emsp;|net30 topology (deprecated)    |

&lt;a id=&#34;nestedatt--server--client_ip_pool&#34;&gt;&lt;/a&gt;
### Nested Schema for `server.client_ip_pool`

Optional:

- `disable` (Boolean) Disable instance
- `start` (String) First IP address in the pool

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|ipv4    &emsp;|IPv4 address  |
- `stop` (String) Last IP address in the pool

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|ipv4    &emsp;|IPv4 address  |
- `subnet_mask` (String) Subnet mask pushed to dynamic clients. If not set the server subnet mask will be used. Only used with topology subnet or device type tap. Not used with bridged interfaces.

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|ipv4    &emsp;|IPv4 subnet mask  |


&lt;a id=&#34;nestedatt--server--client_ipv6_pool&#34;&gt;&lt;/a&gt;
### Nested Schema for `server.client_ipv6_pool`

Optional:

- `base` (String) Client IPv6 pool base address with optional prefix length

    &emsp;|Format   &emsp;|Description                                                                                                                                |
    |-----------|---------------------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|ipv6net  &emsp;|Client IPv6 pool base address with optional prefix length (defaults: base = server subnet + 0x1000, prefix length = server prefix length)  |
- `disable` (Boolean) Disable instance


&lt;a id=&#34;nestedatt--server--mfa&#34;&gt;&lt;/a&gt;
### Nested Schema for `server.mfa`

Optional:

- `totp` (Attributes) Time-based one-time passwords (see [below for nested schema](#nestedatt--server--mfa--totp))

&lt;a id=&#34;nestedatt--server--mfa--totp&#34;&gt;&lt;/a&gt;
### Nested Schema for `server.mfa.totp`

Optional:

- `challenge` (String) Expect password as result of a challenge response protocol

    &emsp;|Format   &emsp;|Description                 |
    |-----------|------------------------------|
    &emsp;|disable  &emsp;|Disable challenge-response  |
    &emsp;|enable   &emsp;|Enable chalenge-response    |
- `digits` (String) Number of digits to use for totp hash

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|1-65535  &emsp;|Digits       |
- `drift` (String) Time drift in seconds

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|1-65535  &emsp;|Seconds      |
- `slop` (String) Maximum allowed clock slop in seconds

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|1-65535  &emsp;|Seconds      |
- `step` (String) Step value for totp in seconds

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|1-65535  &emsp;|Seconds      |




&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


&lt;a id=&#34;nestedatt--tls&#34;&gt;&lt;/a&gt;
### Nested Schema for `tls`

Optional:

- `auth_key` (String) TLS shared secret key for tls-auth
- `ca_certificate` (List of String) Certificate Authority chain in PKI configuration

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|txt     &emsp;|Name of CA in PKI configuration  |
- `certificate` (String) Certificate in PKI configuration

    &emsp;|Format  &emsp;|Description                               |
    |----------|--------------------------------------------|
    &emsp;|txt     &emsp;|Name of certificate in PKI configuration  |
- `crypt_key` (String) Static key to use to authenticate control channel
- `dh_params` (String) Diffie Hellman parameters (server only)
- `peer_fingerprint` (List of String) Peer certificate SHA256 fingerprint
- `role` (String) TLS negotiation role

    &emsp;|Format   &emsp;|Description                        |
    |-----------|-------------------------------------|
    &emsp;|active   &emsp;|Initiate TLS negotiation actively  |
    &emsp;|passive  &emsp;|Wait for incoming TLS connection   |
- `tls_version_min` (String) Specify the minimum required TLS version

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1.0     &emsp;|TLS v1.0     |
    &emsp;|1.1     &emsp;|TLS v1.1     |
    &emsp;|1.2     &emsp;|TLS v1.2     |
    &emsp;|1.3     &emsp;|TLS v1.3     |  
