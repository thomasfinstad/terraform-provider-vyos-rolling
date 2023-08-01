// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsNhrpTunnelDynamicMap describes the resource data model.
type ProtocolsNhrpTunnelDynamicMap struct {
	SelfIdentifier types.String `tfsdk:"dynamic_map_id" vyos:",self-id"`

	ParentIDProtocolsNhrpTunnel types.String `tfsdk:"tunnel" vyos:"tunnel,parent-id"`

	// LeafNodes
	LeafProtocolsNhrpTunnelDynamicMapNbmaDomainName types.String `tfsdk:"nbma_domain_name" vyos:"nbma-domain-name,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsNhrpTunnelDynamicMap) GetVyosPath() []string {
	return []string{
		"protocols",

		"nhrp",

		"tunnel",
		o.ParentIDProtocolsNhrpTunnel.ValueString(),

		"dynamic-map",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsNhrpTunnelDynamicMap) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `dynamic_map_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"dynamic_map_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Set an HUB tunnel address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Set the IP address and prefix length  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"tunnel_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Tunnel for NHRP

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  tunN  &emsp; |  NHRP tunnel name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"nbma_domain_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Set HUB fqdn (nbma-address - fqdn)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <fqdn>  &emsp; |  Set the external HUB fqdn  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsNhrpTunnelDynamicMap) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsNhrpTunnelDynamicMap) UnmarshalJSON(_ []byte) error {
	return nil
}
