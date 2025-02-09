/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (interface) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsIsisInterface{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (interface) */
// VrfNameProtocolsIsisInterfaceSelfIdentifier used by TagNodes to keep the id info
type VrfNameProtocolsIsisInterfaceSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (interface) */

	VrfNameProtocolsIsisInterface types.String `tfsdk:"interface" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (isis) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (protocols) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (name) */

	VrfName types.String `tfsdk:"name" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (vrf) */
}

// VrfNameProtocolsIsisInterface describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type VrfNameProtocolsIsisInterface struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *VrfNameProtocolsIsisInterfaceSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafVrfNameProtocolsIsisInterfaceCircuitType         types.String `tfsdk:"circuit_type" vyos:"circuit-type,omitempty"`
	LeafVrfNameProtocolsIsisInterfaceHelloPadding        types.Bool   `tfsdk:"hello_padding" vyos:"hello-padding,omitempty"`
	LeafVrfNameProtocolsIsisInterfaceHelloInterval       types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafVrfNameProtocolsIsisInterfaceHelloMultIPlier     types.Number `tfsdk:"hello_multiplier" vyos:"hello-multiplier,omitempty"`
	LeafVrfNameProtocolsIsisInterfaceMetric              types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafVrfNameProtocolsIsisInterfacePassive             types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafVrfNameProtocolsIsisInterfacePriority            types.Number `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafVrfNameProtocolsIsisInterfacePsnpInterval        types.Number `tfsdk:"psnp_interval" vyos:"psnp-interval,omitempty"`
	LeafVrfNameProtocolsIsisInterfaceNoThreeWayHandshake types.Bool   `tfsdk:"no_three_way_handshake" vyos:"no-three-way-handshake,omitempty"`

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsIsisInterfaceBfd *VrfNameProtocolsIsisInterfaceBfd `tfsdk:"bfd" vyos:"bfd,omitempty"`

	NodeVrfNameProtocolsIsisInterfaceLdpSync *VrfNameProtocolsIsisInterfaceLdpSync `tfsdk:"ldp_sync" vyos:"ldp-sync,omitempty"`

	NodeVrfNameProtocolsIsisInterfaceNetwork *VrfNameProtocolsIsisInterfaceNetwork `tfsdk:"network" vyos:"network,omitempty"`

	NodeVrfNameProtocolsIsisInterfacePassword *VrfNameProtocolsIsisInterfacePassword `tfsdk:"password" vyos:"password,omitempty"`
}

// SetID configures the resource ID
func (o *VrfNameProtocolsIsisInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VrfNameProtocolsIsisInterface) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsIsisInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsIsisInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"interface",
		o.SelfIdentifier.VrfNameProtocolsIsisInterface.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsIsisInterface) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (isis) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (protocols) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (name) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vrf) */
		"vrf", // Node

		"name",
		o.SelfIdentifier.VrfName.ValueString(),

		"protocols", // Node

		"isis", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsIsisInterface) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (isis) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (protocols) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (name) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (name) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vrf) */
		"vrf", // Node

		"name",
		o.SelfIdentifier.VrfName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (isis) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (protocols) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (name) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (vrf) */

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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  name, value must match: ^[.:a-zA-Z0-9-_/]+$",
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

		"circuit_type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (circuit-type) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (hello-padding) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (hello-interval) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (hello-multiplier) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (metric) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (passive) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (priority) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (psnp-interval) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (no-three-way-handshake) */
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
			Attributes: VrfNameProtocolsIsisInterfaceBfd{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Enable Bidirectional Forwarding Detection (BFD)

`,
			Description: `Enable Bidirectional Forwarding Detection (BFD)

`,
		},

		"ldp_sync": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisInterfaceLdpSync{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `LDP-IGP synchronization configuration for interface

`,
			Description: `LDP-IGP synchronization configuration for interface

`,
		},

		"network": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisInterfaceNetwork{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Set network type

`,
			Description: `Set network type

`,
		},

		"password": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisInterfacePassword{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Configure the authentication password for a circuit

`,
			Description: `Configure the authentication password for a circuit

`,
		},
	}
}
