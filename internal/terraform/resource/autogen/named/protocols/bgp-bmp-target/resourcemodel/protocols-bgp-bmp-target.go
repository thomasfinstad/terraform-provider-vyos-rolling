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

var _ helpers.VyosTopResourceDataModel = &ProtocolsBgpBmpTarget{}

// ProtocolsBgpBmpTarget describes the resource data model.
type ProtocolsBgpBmpTarget struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBgpBmpTargetAddress  types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafProtocolsBgpBmpTargetPort     types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafProtocolsBgpBmpTargetMinRetry types.Number `tfsdk:"min_retry" vyos:"min-retry,omitempty"`
	LeafProtocolsBgpBmpTargetMaxRetry types.Number `tfsdk:"max_retry" vyos:"max-retry,omitempty"`
	LeafProtocolsBgpBmpTargetMirror   types.Bool   `tfsdk:"mirror" vyos:"mirror,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeProtocolsBgpBmpTargetMonitor *ProtocolsBgpBmpTargetMonitor `tfsdk:"monitor" vyos:"monitor,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsBgpBmpTarget) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBgpBmpTarget) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBgpBmpTarget) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpBmpTarget) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"target",
		o.SelfIdentifier.Attributes()["target"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBgpBmpTarget) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"bgp",

		"bmp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsBgpBmpTarget) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpBmpTarget) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"target": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `BMP target

`,
						Description: `BMP target

`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in target, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  target, value must match: ^[a-zA-Z0-9-_]*$",
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

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
    |  ipv6    |  IPv6 address  |
`,
			Description: `IP address

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
    |  ipv6    |  IPv6 address  |
`,
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

			// Default:          stringdefault.StaticString(`5000`),
			Computed: true,
		},

		"min_retry": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Minimum connection retry interval (in milliseconds)

    |  Format        |  Description                        |
    |----------------|-------------------------------------|
    |  100-86400000  |  Minimum connection retry interval  |
`,
			Description: `Minimum connection retry interval (in milliseconds)

    |  Format        |  Description                        |
    |----------------|-------------------------------------|
    |  100-86400000  |  Minimum connection retry interval  |
`,

			// Default:          stringdefault.StaticString(`1000`),
			Computed: true,
		},

		"max_retry": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum connection retry interval

    |  Format          |  Description                        |
    |------------------|-------------------------------------|
    |  100-4294967295  |  Maximum connection retry interval  |
`,
			Description: `Maximum connection retry interval

    |  Format          |  Description                        |
    |------------------|-------------------------------------|
    |  100-4294967295  |  Maximum connection retry interval  |
`,

			// Default:          stringdefault.StaticString(`2000`),
			Computed: true,
		},

		"mirror": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send BMP route mirroring messages

`,
			Description: `Send BMP route mirroring messages

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"monitor": schema.SingleNestedAttribute{
			Attributes: ProtocolsBgpBmpTargetMonitor{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Send BMP route monitoring messages

`,
			Description: `Send BMP route monitoring messages

`,
		},
	}
}
