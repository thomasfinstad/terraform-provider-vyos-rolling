// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesVti describes the resource data model.
type InterfacesVti struct {
	SelfIdentifier types.String `tfsdk:"vti_id" vyos:",self-id"`

	// LeafNodes
	LeafInterfacesVtiAddress     types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesVtiDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesVtiDisable     types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesVtiMtu         types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesVtiRedirect    types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesVtiVrf         types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeInterfacesVtiIP     *InterfacesVtiIP     `tfsdk:"ip" vyos:"ip,omitempty"`
	NodeInterfacesVtiIPvsix *InterfacesVtiIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`
	NodeInterfacesVtiMirror *InterfacesVtiMirror `tfsdk:"mirror" vyos:"mirror,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesVti) GetVyosPath() []string {
	return []string{
		"interfaces",

		"vti",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesVti) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"vti_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Virtual Tunnel Interface (XFRM)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  vtiN  &emsp; |  VTI interface name  |

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

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mtu": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 68-16000  &emsp; |  Maximum Transmission Unit in byte  |

`,

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
		},

		"redirect": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Destination interface name  |

`,
		},

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
		},

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesVtiIP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesVtiIPvsix{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesVtiMirror{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesVti) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesVti) UnmarshalJSON(_ []byte) error {
	return nil
}
