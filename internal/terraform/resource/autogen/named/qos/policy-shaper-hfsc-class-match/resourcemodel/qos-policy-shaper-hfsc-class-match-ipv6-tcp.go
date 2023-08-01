// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyShaperHfscClassMatchIPvsixTCP describes the resource data model.
type QosPolicyShaperHfscClassMatchIPvsixTCP struct {
	// LeafNodes
	LeafQosPolicyShaperHfscClassMatchIPvsixTCPAck types.Bool `tfsdk:"ack" vyos:"ack,omitempty"`
	LeafQosPolicyShaperHfscClassMatchIPvsixTCPSyn types.Bool `tfsdk:"syn" vyos:"syn,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClassMatchIPvsixTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ack": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Match TCP ACK

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"syn": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Match TCP SYN

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyShaperHfscClassMatchIPvsixTCP) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyShaperHfscClassMatchIPvsixTCP) UnmarshalJSON(_ []byte) error {
	return nil
}