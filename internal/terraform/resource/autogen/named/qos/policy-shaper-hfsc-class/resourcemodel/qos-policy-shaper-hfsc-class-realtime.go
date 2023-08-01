// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// QosPolicyShaperHfscClassRealtime describes the resource data model.
type QosPolicyShaperHfscClassRealtime struct {
	// LeafNodes
	LeafQosPolicyShaperHfscClassRealtimeD    types.String `tfsdk:"d" vyos:"d,omitempty"`
	LeafQosPolicyShaperHfscClassRealtimeMone types.String `tfsdk:"m1" vyos:"m1,omitempty"`
	LeafQosPolicyShaperHfscClassRealtimeMtwo types.String `tfsdk:"m2" vyos:"m2,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscClassRealtime) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"d": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Service curve delay

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Time in milliseconds  |

`,
		},

		"m1": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Linkshare m1 parameter for class traffic

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Rate in kbit (kilobit per second)  |
    |  <number>%%  &emsp; |  Percentage of overall rate  |
    |  <number>bit  &emsp; |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit  |
    |  <number>ibit  &emsp; |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    |  <number>ibps  &emsp; |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
    |  <number>bps  &emsp; |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |

`,

			// Default:          stringdefault.StaticString(`100%%`),
			Computed: true,
		},

		"m2": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Linkshare m2 parameter for class traffic

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <number>  &emsp; |  Rate in kbit (kilobit per second)  |
    |  <number>%%  &emsp; |  Percentage of overall rate  |
    |  <number>bit  &emsp; |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit  |
    |  <number>ibit  &emsp; |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    |  <number>ibps  &emsp; |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
    |  <number>bps  &emsp; |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |

`,

			// Default:          stringdefault.StaticString(`100%%`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *QosPolicyShaperHfscClassRealtime) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *QosPolicyShaperHfscClassRealtime) UnmarshalJSON(_ []byte) error {
	return nil
}
