// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// FirewallGroupPortGroup describes the resource data model.
type FirewallGroupPortGroup struct {
	SelfIdentifier types.String `tfsdk:"port_group_id" vyos:",self-id"`

	// LeafNodes
	LeafFirewallGroupPortGroupDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafFirewallGroupPortGroupPort        types.List   `tfsdk:"port" vyos:"port,omitempty"`
	LeafFirewallGroupPortGroupInclude     types.List   `tfsdk:"include" vyos:"include,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *FirewallGroupPortGroup) GetVyosPath() []string {
	return []string{
		"firewall",

		"group",

		"port-group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallGroupPortGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"port_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Firewall port-group

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

		"port": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Port-group member

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Named port (any name in /etc/services, e.g., http)  |
    |  number: 1-65535  &emsp; |  Numbered port  |
    |  start-end  &emsp; |  Numbered port range (e.g. 1001-1050)  |

`,
		},

		"include": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Include another port-group

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *FirewallGroupPortGroup) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *FirewallGroupPortGroup) UnmarshalJSON(_ []byte) error {
	return nil
}