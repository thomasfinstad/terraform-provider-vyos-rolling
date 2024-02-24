// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ProtocolsBgpAddressFamilyLtwovpnEvpnFlooding describes the resource data model.
type ProtocolsBgpAddressFamilyLtwovpnEvpnFlooding struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnFloodingDisable            types.Bool `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafProtocolsBgpAddressFamilyLtwovpnEvpnFloodingHeadEndReplication types.Bool `tfsdk:"head_end_replication" vyos:"head-end-replication,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnFlooding) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyLtwovpnEvpnFlooding) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"protocols",

		"bgp",

		"address-family",

		"l2vpn-evpn",

		"flooding",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyLtwovpnEvpnFlooding) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not flood any BUM packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"head_end_replication": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Flood BUM packets using head-end replication

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}