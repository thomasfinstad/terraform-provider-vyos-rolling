// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsBgpPeerGroupLocalAs{}

// VrfNameProtocolsBgpPeerGroupLocalAs describes the resource data model.
type VrfNameProtocolsBgpPeerGroupLocalAs struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVrfNameProtocolsBgpPeerGroupLocalAsNoPrepend *VrfNameProtocolsBgpPeerGroupLocalAsNoPrepend `tfsdk:"no_prepend" vyos:"no-prepend,omitempty"`
}

// SetID configures the resource ID
func (o *VrfNameProtocolsBgpPeerGroupLocalAs) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VrfNameProtocolsBgpPeerGroupLocalAs) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsBgpPeerGroupLocalAs) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsBgpPeerGroupLocalAs) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"local-as",
		o.SelfIdentifier.Attributes()["local_as"].(types.Number).ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsBgpPeerGroupLocalAs) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",

		o.SelfIdentifier.Attributes()["name"].(types.String).ValueString(),

		"protocols",

		"bgp",

		"peer-group",

		o.SelfIdentifier.Attributes()["peer_group"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsBgpPeerGroupLocalAs) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",

		o.SelfIdentifier.Attributes()["name"].(types.String).ValueString(),

		"protocols",

		"bgp",

		"peer-group",

		o.SelfIdentifier.Attributes()["peer_group"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpPeerGroupLocalAs) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"local_as": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `Specify alternate ASN for this BGP process

    |  Format        |  Description                     |
    |----------------|----------------------------------|
    |  1-4294967294  |  Autonomous System Number (ASN)  |
`,
					Description: `Specify alternate ASN for this BGP process

    |  Format        |  Description                     |
    |----------------|----------------------------------|
    |  1-4294967294  |  Autonomous System Number (ASN)  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				"name": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Virtual Routing and Forwarding instance

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
					Description: `Virtual Routing and Forwarding instance

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in name, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  name, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},

				"peer_group": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Name of peer-group

`,
					Description: `Name of peer-group

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in peer_group, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  peer_group, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		// Nodes

		"no_prepend": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsBgpPeerGroupLocalAsNoPrepend{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Disable prepending local-as from/to updates for eBGP peers

`,
			Description: `Disable prepending local-as from/to updates for eBGP peers

`,
		},
	}
}
