---
page_title: "vyos_service_webproxy_url_filtering_squidguard_rule Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Webproxy service settings⯯URL filtering settings⯯URL filtering via squidGuard redirector⯯URL filter rule for a source-group
---

# vyos_service_webproxy_url_filtering_squidguard_rule (Resource)
<center>

*service*  
⯯  
Webproxy service settings  
⯯  
URL filtering settings  
⯯  
URL filtering via squidGuard redirector  
⯯  
**URL filter rule for a source-group**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

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
- `source_group` (String) Source-group for this rule

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|group   &emsp;|Source group identifier for this rule  |
- `time_period` (String) Time-period for this rule

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|period  &emsp;|Time period for this rule  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `rule` (Number) URL filter rule for a source-group

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1-1024  &emsp;|Rule Number  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
