---
page_title: "vyos_service_webproxy_url_filtering_squidguard Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Webproxy service settings⯯URL filtering settings⯯URL filtering via squidGuard redirector
---

# vyos_service_webproxy_url_filtering_squidguard (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Webproxy service settings  
⯯  
URL filtering settings  
⯯  
**URL filtering via squidGuard redirector**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `allow_category` (List of String) Category to allow
- `allow_ipaddr_url` (Boolean) Allow IP address URLs
- `block_category` (List of String) Category to block
- `default_action` (String) Default action (default: allow)

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|allow   &emsp;|Default filter action is allow)  |
    &emsp;|block   &emsp;|Default filter action is block   |
- `enable_safe_search` (Boolean) Enable safe-mode search on popular search engines
- `local_block` (List of String) Local site to block

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|ipv4    &emsp;|IP address of site to block  |
- `local_block_keyword` (List of String) Local keyword to block

    &emsp;|Format   &emsp;|Description                  |
    |-----------|-------------------------------|
    &emsp;|keyword  &emsp;|Keyword (or regex) to block  |
- `local_block_url` (List of String) Local URL to block

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|url     &emsp;|Local URL to block (without &#34;http://&#34;)  |
- `local_ok` (List of String) Local site to allow

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|ipv4    &emsp;|IP address of site to allow  |
- `local_ok_url` (List of String) Local URL to allow

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|url     &emsp;|Local URL to allow (without &#34;http://&#34;)  |
- `log` (List of String) Log block category
- `redirect_url` (String) Redirect URL for filtered websites

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|url     &emsp;|URL for redirect  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  