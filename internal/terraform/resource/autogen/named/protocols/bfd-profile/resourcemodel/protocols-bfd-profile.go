// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsBfdProfile{}

// ProtocolsBfdProfile describes the resource data model.
type ProtocolsBfdProfile struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBfdProfileEchoMode   types.Bool   `tfsdk:"echo_mode" vyos:"echo-mode,omitempty"`
	LeafProtocolsBfdProfileMinimumTTL types.Number `tfsdk:"minimum_ttl" vyos:"minimum-ttl,omitempty"`
	LeafProtocolsBfdProfilePassive    types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafProtocolsBfdProfileShutdown   types.Bool   `tfsdk:"shutdown" vyos:"shutdown,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeProtocolsBfdProfileInterval *ProtocolsBfdProfileInterval `tfsdk:"interval" vyos:"interval,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBfdProfile) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBfdProfile) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBfdProfile) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBfdProfile) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"profile",
		o.SelfIdentifier.Attributes()["profile"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBfdProfile) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"bfd",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsBfdProfile) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBfdProfile) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"profile": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `Configure BFD profile used by individual peer

    |  Format  |  Description          |
    |----------|-----------------------|
    |  txt     |  Name of BFD profile  |
`,
						Description: `Configure BFD profile used by individual peer

    |  Format  |  Description          |
    |----------|-----------------------|
    |  txt     |  Name of BFD profile  |
`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in profile, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  profile, value must match: ^[a-zA-Z0-9-_]*$",
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

		"echo_mode": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enables the echo transmission mode

`,
			Description: `Enables the echo transmission mode

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"minimum_ttl": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Expect packets with at least this TTL

    |  Format  |  Description           |
    |----------|------------------------|
    |  1-254   |  Minimum TTL expected  |
`,
			Description: `Expect packets with at least this TTL

    |  Format  |  Description           |
    |----------|------------------------|
    |  1-254   |  Minimum TTL expected  |
`,
		},

		"passive": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not attempt to start sessions

`,
			Description: `Do not attempt to start sessions

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"shutdown": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable this peer

`,
			Description: `Disable this peer

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"interval": schema.SingleNestedAttribute{
			Attributes: ProtocolsBfdProfileInterval{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Configure timer intervals

`,
			Description: `Configure timer intervals

`,
		},
	}
}
