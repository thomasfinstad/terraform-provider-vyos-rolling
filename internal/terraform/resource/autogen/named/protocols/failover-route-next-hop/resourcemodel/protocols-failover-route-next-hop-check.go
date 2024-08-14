// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &ProtocolsFailoverRouteNextHopCheck{}

// ProtocolsFailoverRouteNextHopCheck describes the resource data model.
type ProtocolsFailoverRouteNextHopCheck struct {
	// LeafNodes
	LeafProtocolsFailoverRouteNextHopCheckPolicy  types.String `tfsdk:"policy" vyos:"policy,omitempty"`
	LeafProtocolsFailoverRouteNextHopCheckPort    types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafProtocolsFailoverRouteNextHopCheckTarget  types.List   `tfsdk:"target" vyos:"target,omitempty"`
	LeafProtocolsFailoverRouteNextHopCheckTimeout types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafProtocolsFailoverRouteNextHopCheckType    types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsFailoverRouteNextHopCheck) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"policy": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Policy for check targets

    |  Format         |  Description                |
    |-----------------|-----------------------------|
    |  all-available  |  All targets must be alive  |
    |  any-available  |  Any target must be alive   |
`,
			Description: `Policy for check targets

    |  Format         |  Description                |
    |-----------------|-----------------------------|
    |  all-available  |  All targets must be alive  |
    |  any-available  |  Any target must be alive   |
`,

			// Default:          stringdefault.StaticString(`any-available`),
			Computed: true,
		},

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
		},

		"target": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Check target address

    |  Format  |  Description       |
    |----------|--------------------|
    |  ipv4    |  Address to check  |
`,
			Description: `Check target address

    |  Format  |  Description       |
    |----------|--------------------|
    |  ipv4    |  Address to check  |
`,
		},

		"timeout": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout between checks

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  1-300   |  Timeout in seconds between checks  |
`,
			Description: `Timeout between checks

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  1-300   |  Timeout in seconds between checks  |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Check type

    |  Format  |  Description           |
    |----------|------------------------|
    |  arp     |  Check target by ARP   |
    |  icmp    |  Check target by ICMP  |
    |  tcp     |  Check target by TCP   |
`,
			Description: `Check type

    |  Format  |  Description           |
    |----------|------------------------|
    |  arp     |  Check target by ARP   |
    |  icmp    |  Check target by ICMP  |
    |  tcp     |  Check target by TCP   |
`,

			// Default:          stringdefault.StaticString(`icmp`),
			Computed: true,
		},

		// Nodes

	}
}