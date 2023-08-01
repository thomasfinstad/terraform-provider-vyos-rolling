// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesLoopback describes the resource data model.
type InterfacesLoopback struct {
	SelfIdentifier types.String `tfsdk:"loopback_id" vyos:",self-id"`

	// LeafNodes
	LeafInterfacesLoopbackAddress     types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesLoopbackDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesLoopbackRedirect    types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesLoopbackIP     *InterfacesLoopbackIP     `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesLoopbackMirror *InterfacesLoopbackMirror `tfsdk:"mirror" vyos:"mirror,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesLoopback) GetVyosPath() []string {
	return []string{
		"interfaces",

		"loopback",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesLoopback) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"loopback_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Loopback Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  lo  &emsp; |  Loopback interface  |

`,
		},

		// LeafNodes

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 address and prefix length  |
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

`,
		},

		"redirect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Destination interface name  |

`,
		},

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesLoopbackIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesLoopbackMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesLoopback) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesLoopback) UnmarshalJSON(_ []byte) error {
	return nil
}
