// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallZone describes the resource data model.
type FirewallZone struct {
	SelfIdentifier types.String `tfsdk:"zone_id" vyos:",self-id"`

	// LeafNodes
	LeafFirewallZoneDescrIPtion      types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallZoneEnableDefaultLog types.Bool   `tfsdk:"enable_default_log" vyos:"enable-default-log,omitempty"`
	LeafFirewallZoneDefaultAction    types.String `tfsdk:"default_action" vyos:"default-action,omitempty"`
	LeafFirewallZoneInterface        types.List   `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafFirewallZoneLocalZone        types.Bool   `tfsdk:"local_zone" vyos:"local-zone,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagFirewallZoneFrom bool `tfsdk:"from" vyos:"from,child"`

	// Nodes
	NodeFirewallZoneIntraZoneFiltering *FirewallZoneIntraZoneFiltering `tfsdk:"intra_zone_filtering" vyos:"intra-zone-filtering,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallZone) GetVyosPath() []string {
	return []string{
		"firewall",

		"zone",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallZone) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"zone_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Zone-policy

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Zone name  |

`,
		},

		// LeafNodes

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"enable_default_log": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Log packets hitting default-action

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"default_action": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default-action for traffic coming into this zone

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  drop  &emsp; |  Drop silently  |
    |  reject  &emsp; |  Drop and notify source  |

`,

			// Default:          stringdefault.StaticString(`drop`),
			Computed: true,
		},

		"interface": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface associated with zone

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface associated with zone  |

`,
		},

		"local_zone": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Zone to be local-zone

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"intra_zone_filtering": schema.SingleNestedAttribute{
			Attributes: FirewallZoneIntraZoneFiltering{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Intra-zone filtering

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallZone) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *FirewallZone) UnmarshalJSON(_ []byte) error {
	return nil
}
