// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyShaperClassMatchIPvsix describes the resource data model.
type QosPolicyShaperClassMatchIPvsix struct {
	// LeafNodes
	LeafQosPolicyShaperClassMatchIPvsixDscp      types.String `tfsdk:"dscp" vyos:"dscp,omitempty"`
	LeafQosPolicyShaperClassMatchIPvsixMaxLength types.Number `tfsdk:"max_length" vyos:"max-length,omitempty"`
	LeafQosPolicyShaperClassMatchIPvsixProtocol  types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeQosPolicyShaperClassMatchIPvsixDestination *QosPolicyShaperClassMatchIPvsixDestination `tfsdk:"destination" vyos:"destination,omitempty"`
	NodeQosPolicyShaperClassMatchIPvsixSource      *QosPolicyShaperClassMatchIPvsixSource      `tfsdk:"source" vyos:"source,omitempty"`
	NodeQosPolicyShaperClassMatchIPvsixTCP         *QosPolicyShaperClassMatchIPvsixTCP         `tfsdk:"tcp" vyos:"tcp,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperClassMatchIPvsix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"dscp": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Match on Differentiated Services Codepoint (DSCP)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-63  &emsp; |  Differentiated Services Codepoint (DSCP) value   |
    |  default  &emsp; |  match DSCP (000000)  |
    |  reliability  &emsp; |  match DSCP (000001)  |
    |  throughput  &emsp; |  match DSCP (000010)  |
    |  lowdelay  &emsp; |  match DSCP (000100)  |
    |  priority  &emsp; |  match DSCP (001000)  |
    |  immediate  &emsp; |  match DSCP (010000)  |
    |  flash  &emsp; |  match DSCP (011000)  |
    |  flash-override  &emsp; |  match DSCP (100000)  |
    |  critical  &emsp; |  match DSCP (101000)  |
    |  internet  &emsp; |  match DSCP (110000)  |
    |  network  &emsp; |  match DSCP (111000)  |
    |  AF11  &emsp; |  High-throughput data  |
    |  AF12  &emsp; |  High-throughput data  |
    |  AF13  &emsp; |  High-throughput data  |
    |  AF21  &emsp; |  Low-latency data  |
    |  AF22  &emsp; |  Low-latency data  |
    |  AF23  &emsp; |  Low-latency data  |
    |  AF31  &emsp; |  Multimedia streaming  |
    |  AF32  &emsp; |  Multimedia streaming  |
    |  AF33  &emsp; |  Multimedia streaming  |
    |  AF41  &emsp; |  Multimedia conferencing  |
    |  AF42  &emsp; |  Multimedia conferencing  |
    |  AF43  &emsp; |  Multimedia conferencing  |
    |  CS1  &emsp; |  Low-priority data  |
    |  CS2  &emsp; |  OAM  |
    |  CS3  &emsp; |  Broadcast video  |
    |  CS4  &emsp; |  Real-time interactive  |
    |  CS5  &emsp; |  Signaling  |
    |  CS6  &emsp; |  Network control  |
    |  CS7  &emsp; |    |
    |  EF  &emsp; |  Expedited Forwarding  |

`,
		},

		"max_length": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum packet length

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Maximum packet/payload length  |

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Protocol name  |

`,
		},

		// Nodes

		"destination": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPvsixDestination{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on destination port or address

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPvsixSource{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Match on source port or address

`,
		},

		"tcp": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperClassMatchIPvsixTCP{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `TCP Flags matching

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyShaperClassMatchIPvsix) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyShaperClassMatchIPvsix) UnmarshalJSON(_ []byte) error {
	return nil
}
