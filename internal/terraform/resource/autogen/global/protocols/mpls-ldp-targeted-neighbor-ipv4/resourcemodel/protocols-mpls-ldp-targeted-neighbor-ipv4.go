// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &ProtocolsMplsLdpTargetedNeighborIPvfour{}

// ProtocolsMplsLdpTargetedNeighborIPvfour describes the resource data model.
type ProtocolsMplsLdpTargetedNeighborIPvfour struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsMplsLdpTargetedNeighborIPvfourAddress       types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafProtocolsMplsLdpTargetedNeighborIPvfourEnable        types.Bool   `tfsdk:"enable" vyos:"enable,omitempty"`
	LeafProtocolsMplsLdpTargetedNeighborIPvfourHelloInterval types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafProtocolsMplsLdpTargetedNeighborIPvfourHelloHoldtime types.Number `tfsdk:"hello_holdtime" vyos:"hello-holdtime,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsMplsLdpTargetedNeighborIPvfour) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsMplsLdpTargetedNeighborIPvfour) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsMplsLdpTargetedNeighborIPvfour) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsMplsLdpTargetedNeighborIPvfour) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"ipv4",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsMplsLdpTargetedNeighborIPvfour) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"mpls",

		"ldp",

		"targeted-neighbor",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsMplsLdpTargetedNeighborIPvfour) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMplsLdpTargetedNeighborIPvfour) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Neighbor/session address

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  Neighbor/session address  |
`,
			Description: `Neighbor/session address

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  Neighbor/session address  |
`,
		},

		"enable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Accept and respond to targeted hellos

`,
			Description: `Accept and respond to targeted hellos

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"hello_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello interval

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
			Description: `Hello interval

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
		},

		"hello_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello hold time

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
			Description: `Hello hold time

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Time in seconds  |
`,
		},
	}
}
