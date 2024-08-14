---
page_title: "vyos_service_ssh_access_control_allow Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  System services⯯Secure Shell (SSH)⯯SSH user/group access controls⯯Allow user/group SSH access
---

# vyos_service_ssh_access_control_allow (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

System services  
⯯  
Secure Shell (SSH)  
⯯  
SSH user/group access controls  
⯯  
**Allow user/group SSH access**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `group` (List of String) Allow members of a group to login
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `user` (List of String) Allow specific users to login

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  