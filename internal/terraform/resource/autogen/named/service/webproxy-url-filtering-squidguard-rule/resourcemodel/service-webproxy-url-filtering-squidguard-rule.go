// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceWebproxyURLFilteringSquIDguardRule describes the resource data model.
type ServiceWebproxyURLFilteringSquIDguardRule struct {
	SelfIdentifier types.Number `tfsdk:"rule_id" vyos:",self-id"`

	// LeafNodes
	LeafServiceWebproxyURLFilteringSquIDguardRuleAllowCategory     types.List   `tfsdk:"allow_category" vyos:"allow-category,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleAllowIPaddrURL    types.Bool   `tfsdk:"allow_ipaddr_url" vyos:"allow-ipaddr-url,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleBlockCategory     types.List   `tfsdk:"block_category" vyos:"block-category,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleDefaultAction     types.String `tfsdk:"default_action" vyos:"default-action,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleEnableSafeSearch  types.Bool   `tfsdk:"enable_safe_search" vyos:"enable-safe-search,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockKeyword types.List   `tfsdk:"local_block_keyword" vyos:"local-block-keyword,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlockURL     types.List   `tfsdk:"local_block_url" vyos:"local-block-url,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLocalBlock        types.List   `tfsdk:"local_block" vyos:"local-block,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOkURL        types.List   `tfsdk:"local_ok_url" vyos:"local-ok-url,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLocalOk           types.List   `tfsdk:"local_ok" vyos:"local-ok,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleLog               types.List   `tfsdk:"log" vyos:"log,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleRedirectURL       types.String `tfsdk:"redirect_url" vyos:"redirect-url,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleSourceGroup       types.String `tfsdk:"source_group" vyos:"source-group,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardRuleTimePeriod        types.String `tfsdk:"time_period" vyos:"time-period,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceWebproxyURLFilteringSquIDguardRule) GetVyosPath() []string {
	return []string{
		"service",

		"webproxy",

		"url-filtering",

		"squidguard",

		"rule",
		o.SelfIdentifier.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceWebproxyURLFilteringSquIDguardRule) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"rule_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `URL filter rule for a source-group

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-1024  &emsp; |  Rule Number  |

`,
		},

		// LeafNodes

		"allow_category": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Category to allow

`,
		},

		"allow_ipaddr_url": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Allow IP address URLs

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"block_category": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Category to block

`,
		},

		"default_action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default action (default: allow)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  allow  &emsp; |  Default filter action is allow)  |
    |  block  &emsp; |  Default filter action is block  |

`,
		},

		"enable_safe_search": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable safe-mode search on popular search engines

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"local_block_keyword": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local keyword to block

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  keyword  &emsp; |  Keyword (or regex) to block  |

`,
		},

		"local_block_url": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local URL to block

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  url  &emsp; |  Local URL to block (without "http://")  |

`,
		},

		"local_block": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local site to block

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IP address of site to block  |

`,
		},

		"local_ok_url": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local URL to allow

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  url  &emsp; |  Local URL to allow (without "http://")  |

`,
		},

		"local_ok": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local site to allow

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IP address of site to allow  |

`,
		},

		"log": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Log block category

`,
		},

		"redirect_url": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect URL for filtered websites

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  url  &emsp; |  URL for redirect  |

`,
		},

		"source_group": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source-group for this rule

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  group  &emsp; |  Source group identifier for this rule  |

`,
		},

		"time_period": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time-period for this rule

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  period  &emsp; |  Time period for this rule  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceWebproxyURLFilteringSquIDguardRule) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceWebproxyURLFilteringSquIDguardRule) UnmarshalJSON(_ []byte) error {
	return nil
}