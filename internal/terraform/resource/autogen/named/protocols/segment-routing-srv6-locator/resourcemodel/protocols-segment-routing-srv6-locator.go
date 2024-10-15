// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
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

var _ helpers.VyosTopResourceDataModel = &ProtocolsSegmentRoutingSrvsixLocator{}

// ProtocolsSegmentRoutingSrvsixLocator describes the resource data model.
type ProtocolsSegmentRoutingSrvsixLocator struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsSegmentRoutingSrvsixLocatorBehaviorUsID types.Bool   `tfsdk:"behavior_usid" vyos:"behavior-usid,omitempty"`
	LeafProtocolsSegmentRoutingSrvsixLocatorPrefix       types.String `tfsdk:"prefix" vyos:"prefix,omitempty"`
	LeafProtocolsSegmentRoutingSrvsixLocatorBlockLen     types.Number `tfsdk:"block_len" vyos:"block-len,omitempty"`
	LeafProtocolsSegmentRoutingSrvsixLocatorFuncBits     types.Number `tfsdk:"func_bits" vyos:"func-bits,omitempty"`
	LeafProtocolsSegmentRoutingSrvsixLocatorNodeLen      types.Number `tfsdk:"node_len" vyos:"node-len,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsSegmentRoutingSrvsixLocator) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsSegmentRoutingSrvsixLocator) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsSegmentRoutingSrvsixLocator) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsSegmentRoutingSrvsixLocator) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"locator",
		o.SelfIdentifier.Attributes()["locator"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsSegmentRoutingSrvsixLocator) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"segment-routing",

		"srv6",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsSegmentRoutingSrvsixLocator) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsSegmentRoutingSrvsixLocator) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"locator": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Segment Routing SRv6 locator

`,
					Description: `Segment Routing SRv6 locator

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in locator, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  locator, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"behavior_usid": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Set SRv6 behavior uSID

`,
			Description: `Set SRv6 behavior uSID

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"prefix": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `SRv6 locator prefix

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  ipv6net  |  SRv6 locator prefix  |
`,
			Description: `SRv6 locator prefix

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  ipv6net  |  SRv6 locator prefix  |
`,
		},

		"block_len": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Configure SRv6 locator block length in bits

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  16-64   |  Specify SRv6 locator block length in bits  |
`,
			Description: `Configure SRv6 locator block length in bits

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  16-64   |  Specify SRv6 locator block length in bits  |
`,

			// Default:          stringdefault.StaticString(`40`),
			Computed: true,
		},

		"func_bits": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Configure SRv6 locator function length in bits

    |  Format  |  Description                                   |
    |----------|------------------------------------------------|
    |  0-64    |  Specify SRv6 locator function length in bits  |
`,
			Description: `Configure SRv6 locator function length in bits

    |  Format  |  Description                                   |
    |----------|------------------------------------------------|
    |  0-64    |  Specify SRv6 locator function length in bits  |
`,

			// Default:          stringdefault.StaticString(`16`),
			Computed: true,
		},

		"node_len": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Configure SRv6 locator node length in bits

    |  Format  |  Description                                 |
    |----------|----------------------------------------------|
    |  16-64   |  Configure SRv6 locator node length in bits  |
`,
			Description: `Configure SRv6 locator node length in bits

    |  Format  |  Description                                 |
    |----------|----------------------------------------------|
    |  16-64   |  Configure SRv6 locator node length in bits  |
`,

			// Default:          stringdefault.StaticString(`24`),
			Computed: true,
		},

		// Nodes

	}
}
