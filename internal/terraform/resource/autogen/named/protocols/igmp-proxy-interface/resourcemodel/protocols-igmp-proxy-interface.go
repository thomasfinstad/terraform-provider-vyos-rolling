// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsIgmpProxyInterface describes the resource data model.
type ProtocolsIgmpProxyInterface struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	// LeafNodes
	LeafProtocolsIgmpProxyInterfaceAltSubnet types.List   `tfsdk:"alt_subnet" vyos:"alt-subnet,omitempty"`
	LeafProtocolsIgmpProxyInterfaceRole      types.String `tfsdk:"role" vyos:"role,omitempty"`
	LeafProtocolsIgmpProxyInterfaceThreshold types.Number `tfsdk:"threshold" vyos:"threshold,omitempty"`
	LeafProtocolsIgmpProxyInterfaceWhitelist types.List   `tfsdk:"whitelist" vyos:"whitelist,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetID returns the resource ID
func (o ProtocolsIgmpProxyInterface) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o ProtocolsIgmpProxyInterface) SetID(id types.String) {
	o.ID = id
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsIgmpProxyInterface) GetVyosPath() []string {
	return []string{
		"protocols",

		"igmp-proxy",

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsIgmpProxyInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface for IGMP proxy

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"alt_subnet": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Unicast source networks allowed for multicast traffic to be proxyed

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 network  |

`,
		},

		"role": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IGMP interface role

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  upstream  &emsp; |  Upstream interface (only 1 allowed)  |
    |  downstream  &emsp; |  Downstream interface(s)  |
    |  disabled  &emsp; |  Disabled interface  |

`,

			// Default:          stringdefault.StaticString(`downstream`),
			Computed: true,
		},

		"threshold": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `TTL threshold

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  TTL threshold for the interfaces  |

`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"whitelist": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Group to whitelist

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 network  |

`,
		},

		// Nodes

	}
}
