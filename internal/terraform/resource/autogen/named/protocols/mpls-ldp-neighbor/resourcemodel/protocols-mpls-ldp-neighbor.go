// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsMplsLdpNeighbor{}

// ProtocolsMplsLdpNeighbor describes the resource data model.
type ProtocolsMplsLdpNeighbor struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsMplsLdpNeighborPassword        types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafProtocolsMplsLdpNeighborTTLSecURIty     types.String `tfsdk:"ttl_security" vyos:"ttl-security,omitempty"`
	LeafProtocolsMplsLdpNeighborSessionHoldtime types.Number `tfsdk:"session_holdtime" vyos:"session-holdtime,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsMplsLdpNeighbor) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsMplsLdpNeighbor) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsMplsLdpNeighbor) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsMplsLdpNeighbor) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"neighbor",
		o.SelfIdentifier.Attributes()["neighbor"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsMplsLdpNeighbor) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"mpls",

		"ldp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsMplsLdpNeighbor) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMplsLdpNeighbor) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"neighbor": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `LDP neighbor parameters

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  Neighbor IPv4 address  |
`,
						Description: `LDP neighbor parameters

    |  Format  |  Description            |
    |----------|-------------------------|
    |  ipv4    |  Neighbor IPv4 address  |
`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in neighbor, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  neighbor, value must match: ^[a-zA-Z0-9-_]*$",
								),
							),
						},
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor password

`,
			Description: `Neighbor password

`,
		},

		"ttl_security": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Neighbor TTL security

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-254    |  TTL                            |
    |  disable  |  Disable neighbor TTL security  |
`,
			Description: `Neighbor TTL security

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-254    |  TTL                            |
    |  disable  |  Disable neighbor TTL security  |
`,
		},

		"session_holdtime": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Session IPv4 hold time

    |  Format    |  Description      |
    |------------|-------------------|
    |  15-65535  |  Time in seconds  |
`,
			Description: `Session IPv4 hold time

    |  Format    |  Description      |
    |------------|-------------------|
    |  15-65535  |  Time in seconds  |
`,
		},

		// Nodes

	}
}
