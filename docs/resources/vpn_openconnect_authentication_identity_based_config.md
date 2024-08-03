---
page_title: "vyos_vpn_openconnect_authentication_identity_based_config Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  vpn⯯SSL VPN OpenConnect, AnyConnect compatible server⯯Authentication for remote access SSL VPN Server⯯Include configuration file by username or RADIUS group attribute
---

# vyos_vpn_openconnect_authentication_identity_based_config (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*vpn*  
⯯  
SSL VPN OpenConnect, AnyConnect compatible server  
⯯  
Authentication for remote access SSL VPN Server  
⯯  
**Include configuration file by username or RADIUS group attribute**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `default_config` (String) Default configuration if discrete config could not be found

    &emsp;|Format    &emsp;|Description                                                 |
    |------------|--------------------------------------------------------------|
    &emsp;|filename  &emsp;|Default configuration filename, must be under /config/auth  |
- `directory` (String) Directory to containing configuration files

    &emsp;|Format  &emsp;|Description                                                  |
    |----------|---------------------------------------------------------------|
    &emsp;|path    &emsp;|Path to configuration directory, must be under /config/auth  |
- `disable` (Boolean) Disable instance
- `mode` (String) Select per user or per group configuration file - ignored if authentication group is configured

    &emsp;|Format  &emsp;|Description                                         |
    |----------|------------------------------------------------------|
    &emsp;|user    &emsp;|Match configuration file on username                |
    &emsp;|group   &emsp;|Match RADIUS response class attribute as file name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
