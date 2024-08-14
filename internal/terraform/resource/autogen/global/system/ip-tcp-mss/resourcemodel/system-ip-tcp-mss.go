// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &SystemIPTCPMss{}

// SystemIPTCPMss describes the resource data model.
type SystemIPTCPMss struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemIPTCPMssProbing types.String `tfsdk:"probing" vyos:"probing,omitempty"`
	LeafSystemIPTCPMssBase    types.Number `tfsdk:"base" vyos:"base,omitempty"`
	LeafSystemIPTCPMssFloor   types.Number `tfsdk:"floor" vyos:"floor,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemIPTCPMss) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemIPTCPMss) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemIPTCPMss) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemIPTCPMss) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"mss",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemIPTCPMss) GetVyosParentPath() []string {
	return []string{
		"system",

		"ip",

		"tcp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *SystemIPTCPMss) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemIPTCPMss) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"probing": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Attempt to lower the MSS if TCP connections fail to establish

    |  Format              |  Description                                                  |
    |----------------------|---------------------------------------------------------------|
    |  on-icmp-black-hole  |  Attempt TCP MSS probing when an ICMP black hole is detected  |
    |  force               |  Attempt TCP MSS probing by default                           |
`,
			Description: `Attempt to lower the MSS if TCP connections fail to establish

    |  Format              |  Description                                                  |
    |----------------------|---------------------------------------------------------------|
    |  on-icmp-black-hole  |  Attempt TCP MSS probing when an ICMP black hole is detected  |
    |  force               |  Attempt TCP MSS probing by default                           |
`,
		},

		"base": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Base MSS to start probing from (applicable to "probing force")

    |  Format   |  Description                                 |
    |-----------|----------------------------------------------|
    |  48-1460  |  Base MSS value for probing (default: 1024)  |
`,
			Description: `Base MSS to start probing from (applicable to "probing force")

    |  Format   |  Description                                 |
    |-----------|----------------------------------------------|
    |  48-1460  |  Base MSS value for probing (default: 1024)  |
`,
		},

		"floor": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum MSS to stop probing at (default: 48)

    |  Format   |  Description                 |
    |-----------|------------------------------|
    |  48-1460  |  Minimum MSS value to probe  |
`,
			Description: `Minimum MSS to stop probing at (default: 48)

    |  Format   |  Description                 |
    |-----------|------------------------------|
    |  48-1460  |  Minimum MSS value to probe  |
`,
		},
	}
}