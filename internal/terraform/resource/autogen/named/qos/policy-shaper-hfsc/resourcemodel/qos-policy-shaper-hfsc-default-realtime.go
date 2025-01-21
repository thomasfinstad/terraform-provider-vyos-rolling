/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (realtime) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &QosPolicyShaperHfscDefaultRealtime{}

// QosPolicyShaperHfscDefaultRealtime describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type QosPolicyShaperHfscDefaultRealtime struct {
	// LeafNodes
	LeafQosPolicyShaperHfscDefaultRealtimeD    types.String `tfsdk:"d" vyos:"d,omitempty"`
	LeafQosPolicyShaperHfscDefaultRealtimeMone types.String `tfsdk:"m1" vyos:"m1,omitempty"`
	LeafQosPolicyShaperHfscDefaultRealtimeMtwo types.String `tfsdk:"m2" vyos:"m2,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscDefaultRealtime) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"d":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (d) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Service curve delay

    |  Format    |  Description           |
    |------------|------------------------|
    |  <number>  |  Time in milliseconds  |
`,
			Description: `Service curve delay

    |  Format    |  Description           |
    |------------|------------------------|
    |  <number>  |  Time in milliseconds  |
`,
		},

		"m1":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (m1) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Linkshare m1 parameter for class traffic

    |  Format        |  Description                                              |
    |----------------|-----------------------------------------------------------|
    |  <number>      |  Rate in kbit (kilobit per second)                        |
    |  <number>%%    |  Percentage of overall rate                               |
    |  <number>bit   |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    |  <number>ibit  |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    |  <number>ibps  |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
    |  <number>bps   |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |
`,
			Description: `Linkshare m1 parameter for class traffic

    |  Format        |  Description                                              |
    |----------------|-----------------------------------------------------------|
    |  <number>      |  Rate in kbit (kilobit per second)                        |
    |  <number>%%    |  Percentage of overall rate                               |
    |  <number>bit   |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    |  <number>ibit  |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    |  <number>ibps  |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
    |  <number>bps   |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |
`,
		},

		"m2":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (m2) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Linkshare m2 parameter for class traffic

    |  Format        |  Description                                              |
    |----------------|-----------------------------------------------------------|
    |  <number>      |  Rate in kbit (kilobit per second)                        |
    |  <number>%%    |  Percentage of overall rate                               |
    |  <number>bit   |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    |  <number>ibit  |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    |  <number>ibps  |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
    |  <number>bps   |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |
`,
			Description: `Linkshare m2 parameter for class traffic

    |  Format        |  Description                                              |
    |----------------|-----------------------------------------------------------|
    |  <number>      |  Rate in kbit (kilobit per second)                        |
    |  <number>%%    |  Percentage of overall rate                               |
    |  <number>bit   |  bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    |  <number>ibit  |  kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    |  <number>ibps  |  kibps(1024&8), mibps(1024^2&8), gibps, tibps - Byte/sec  |
    |  <number>bps   |  bps(8),kbps(8&10^3),mbps(8&10^6), gbps, tbps - Byte/sec  |
`,
		},

		// TagNodes

		// Nodes

	}
}
