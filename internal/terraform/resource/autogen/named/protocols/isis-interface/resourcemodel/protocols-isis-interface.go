/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsIsisInterface{}

// ProtocolsIsisInterface describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ProtocolsIsisInterface struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl */
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsIsisInterfaceCircuitType         types.String `tfsdk:"circuit_type" vyos:"circuit-type,omitempty"`
	LeafProtocolsIsisInterfaceHelloPadding        types.Bool   `tfsdk:"hello_padding" vyos:"hello-padding,omitempty"`
	LeafProtocolsIsisInterfaceHelloInterval       types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafProtocolsIsisInterfaceHelloMultIPlier     types.Number `tfsdk:"hello_multiplier" vyos:"hello-multiplier,omitempty"`
	LeafProtocolsIsisInterfaceMetric              types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafProtocolsIsisInterfacePassive             types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafProtocolsIsisInterfacePriority            types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafProtocolsIsisInterfacePsnpInterval        types.Number `tfsdk:"psnp_interval" vyos:"psnp-interval,omitempty"`
	LeafProtocolsIsisInterfaceNoThreeWayHandshake types.Bool   `tfsdk:"no_three_way_handshake" vyos:"no-three-way-handshake,omitempty"`

	// TagNodes

	// Nodes

	NodeProtocolsIsisInterfaceBfd *ProtocolsIsisInterfaceBfd `tfsdk:"bfd" vyos:"bfd,omitempty"`

	NodeProtocolsIsisInterfaceLdpSync *ProtocolsIsisInterfaceLdpSync `tfsdk:"ldp_sync" vyos:"ldp-sync,omitempty"`

	NodeProtocolsIsisInterfaceNetwork *ProtocolsIsisInterfaceNetwork `tfsdk:"network" vyos:"network,omitempty"`

	NodeProtocolsIsisInterfacePassword *ProtocolsIsisInterfacePassword `tfsdk:"password" vyos:"password,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsIsisInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsIsisInterface) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsIsisInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsIsisInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"interface",
		o.SelfIdentifier.Attributes()["interface"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsIsisInterface) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"protocols", // Node

		"isis", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsIsisInterface) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsIsisInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"interface": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Interface params

`,
					Description: `Interface params

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in interface, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  interface, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"circuit_type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Configure circuit type for interface

    |  Format        |  Description                          |
    |----------------|---------------------------------------|
    |  level-1       |  Level-1 only adjacencies are formed  |
    |  level-1-2     |  Level-1-2 adjacencies are formed     |
    |  level-2-only  |  Level-2 only adjacencies are formed  |
`,
			Description: `Configure circuit type for interface

    |  Format        |  Description                          |
    |----------------|---------------------------------------|
    |  level-1       |  Level-1 only adjacencies are formed  |
    |  level-1-2     |  Level-1-2 adjacencies are formed     |
    |  level-2-only  |  Level-2 only adjacencies are formed  |
`,
		},

		"hello_padding":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Add padding to IS-IS hello packets

`,
			Description: `Add padding to IS-IS hello packets

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"hello_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set Hello interval

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-600   |  Set Hello interval  |
`,
			Description: `Set Hello interval

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-600   |  Set Hello interval  |
`,
		},

		"hello_multiplier":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set Hello interval

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  2-100   |  Set multiplier for Hello holding time  |
`,
			Description: `Set Hello interval

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  2-100   |  Set multiplier for Hello holding time  |
`,
		},

		"metric":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set default metric for circuit

    |  Format      |  Description           |
    |--------------|------------------------|
    |  0-16777215  |  Default metric value  |
`,
			Description: `Set default metric for circuit

    |  Format      |  Description           |
    |--------------|------------------------|
    |  0-16777215  |  Default metric value  |
`,
		},

		"passive":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Configure passive mode for interface

`,
			Description: `Configure passive mode for interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"priority":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set priority for Designated Router election

    |  Format  |  Description     |
    |----------|------------------|
    |  0-127   |  Priority value  |
`,
			Description: `Set priority for Designated Router election

    |  Format  |  Description     |
    |----------|------------------|
    |  0-127   |  Priority value  |
`,
		},

		"psnp_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Set PSNP interval

    |  Format  |  Description               |
    |----------|----------------------------|
    |  0-127   |  PSNP interval in seconds  |
`,
			Description: `Set PSNP interval

    |  Format  |  Description               |
    |----------|----------------------------|
    |  0-127   |  PSNP interval in seconds  |
`,
		},

		"no_three_way_handshake":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable three-way handshake

`,
			Description: `Disable three-way handshake

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"bfd": schema.SingleNestedAttribute{
			Attributes: ProtocolsIsisInterfaceBfd{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD)

`,
			Description: `Enable Bidirectional Forwarding Detection (BFD)

`,
		},

		"ldp_sync": schema.SingleNestedAttribute{
			Attributes: ProtocolsIsisInterfaceLdpSync{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `LDP-IGP synchronization configuration for interface

`,
			Description: `LDP-IGP synchronization configuration for interface

`,
		},

		"network": schema.SingleNestedAttribute{
			Attributes: ProtocolsIsisInterfaceNetwork{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Set network type

`,
			Description: `Set network type

`,
		},

		"password": schema.SingleNestedAttribute{
			Attributes: ProtocolsIsisInterfacePassword{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Configure the authentication password for a circuit

`,
			Description: `Configure the authentication password for a circuit

`,
		},
	}
}
