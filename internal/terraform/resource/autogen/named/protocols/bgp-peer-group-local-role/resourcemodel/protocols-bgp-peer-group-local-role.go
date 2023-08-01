// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsBgpPeerGroupLocalRole describes the resource data model.
type ProtocolsBgpPeerGroupLocalRole struct {
	SelfIdentifier types.String `tfsdk:"local_role_id" vyos:",self-id"`

	ParentIDProtocolsBgpPeerGroup types.String `tfsdk:"peer_group" vyos:"peer-group,parent-id"`

	// LeafNodes
	LeafProtocolsBgpPeerGroupLocalRoleStrict types.Bool `tfsdk:"strict" vyos:"strict,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpPeerGroupLocalRole) GetVyosPath() []string {
	return []string{
		"protocols",

		"bgp",

		"peer-group",
		o.ParentIDProtocolsBgpPeerGroup.ValueString(),

		"local-role",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpPeerGroupLocalRole) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"local_role_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Local role for BGP neighbor (RFC9234)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  customer  &emsp; |  Using Transit  |
    |  peer  &emsp; |  Public/Private Peering  |
    |  provider  &emsp; |  Providing Transit  |
    |  rs-client  &emsp; |  RS Client  |
    |  rs-server  &emsp; |  Route Server  |

`,
		},

		"peer_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Name of peer-group

`,
		},

		// LeafNodes

		"strict": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor must send this exact capability, otherwise a role missmatch notification will be sent

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsBgpPeerGroupLocalRole) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsBgpPeerGroupLocalRole) UnmarshalJSON(_ []byte) error {
	return nil
}
