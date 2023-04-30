// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/customtypes"
)

// QosPolicyShaperHfscClassLinkshare describes the resource data model.
type QosPolicyShaperHfscClassLinkshare struct {
	// LeafNodes
	QosPolicyShaperHfscClassLinkshareD    customtypes.CustomStringValue `tfsdk:"d" json:"d,omitempty"`
	QosPolicyShaperHfscClassLinkshareMone customtypes.CustomStringValue `tfsdk:"m1" json:"m1,omitempty"`
	QosPolicyShaperHfscClassLinkshareMtwo customtypes.CustomStringValue `tfsdk:"m2" json:"m2,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceAttributes generates the attributes for the resource at this level
func (o QosPolicyShaperHfscClassLinkshare) ResourceAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"d": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Service curve delay

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Time in milliseconds  |

`,
		},

		"m1": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Linkshare m1 parameter for class traffic

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Rate in kbit (kilobit per second)  |
|  <number>%%  |  Percentage of overall rate  |
|  <number>bit  |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit  |
|  <number>ibit  |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
|  <number>ibps  |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
|  <number>bps  |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |

`,

			// Default:          stringdefault.StaticString(`100%%`),
			Computed: true,
		},

		"m2": schema.StringAttribute{
			CustomType: customtypes.CustomStringType{},
			Optional:   true,
			MarkdownDescription: `Linkshare m2 parameter for class traffic

|  Format  |  Description  |
|----------|---------------|
|  <number>  |  Rate in kbit (kilobit per second)  |
|  <number>%%  |  Percentage of overall rate  |
|  <number>bit  |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit  |
|  <number>ibit  |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
|  <number>ibps  |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
|  <number>bps  |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |

`,

			// Default:          stringdefault.StaticString(`100%%`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
