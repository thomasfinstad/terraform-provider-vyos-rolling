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

var _ helpers.VyosTopResourceDataModel = &ProtocolsOpenfabricDomainInterface{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (interface) */
// ProtocolsOpenfabricDomainInterfaceSelfIdentifier used by TagNodes to keep the id info
type ProtocolsOpenfabricDomainInterfaceSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (interface) */

	ProtocolsOpenfabricDomainInterface types.String `tfsdk:"interface" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (domain) */

	ProtocolsOpenfabricDomain types.String `tfsdk:"domain" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (openfabric) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (protocols) */
}

// ProtocolsOpenfabricDomainInterface describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ProtocolsOpenfabricDomainInterface struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *ProtocolsOpenfabricDomainInterfaceSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafProtocolsOpenfabricDomainInterfaceCsnpInterval    types.Number `tfsdk:"csnp_interval" vyos:"csnp-interval,omitempty"`
	LeafProtocolsOpenfabricDomainInterfaceHelloInterval   types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafProtocolsOpenfabricDomainInterfaceHelloMultIPlier types.Number `tfsdk:"hello_multiplier" vyos:"hello-multiplier,omitempty"`
	LeafProtocolsOpenfabricDomainInterfaceMetric          types.Number `tfsdk:"metric" vyos:"metric,omitempty"`
	LeafProtocolsOpenfabricDomainInterfacePassive         types.Bool   `tfsdk:"passive" vyos:"passive,omitempty"`
	LeafProtocolsOpenfabricDomainInterfacePsnpInterval    types.Number `tfsdk:"psnp_interval" vyos:"psnp-interval,omitempty"`

	// TagNodes

	// Nodes

	NodeProtocolsOpenfabricDomainInterfaceAddressFamily *ProtocolsOpenfabricDomainInterfaceAddressFamily `tfsdk:"address_family" vyos:"address-family,omitempty"`

	NodeProtocolsOpenfabricDomainInterfacePassword *ProtocolsOpenfabricDomainInterfacePassword `tfsdk:"password" vyos:"password,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsOpenfabricDomainInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsOpenfabricDomainInterface) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsOpenfabricDomainInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOpenfabricDomainInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"interface",
		o.SelfIdentifier.ProtocolsOpenfabricDomainInterface.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsOpenfabricDomainInterface) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (domain) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (openfabric) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (protocols) */
		"protocols", // Node

		"openfabric", // Node

		"domain",
		o.SelfIdentifier.ProtocolsOpenfabricDomain.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsOpenfabricDomainInterface) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (domain) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (domain) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (openfabric) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (protocols) */
		"protocols", // Node

		"openfabric", // Node

		"domain",
		o.SelfIdentifier.ProtocolsOpenfabricDomain.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOpenfabricDomainInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
					Description: `Interface params

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
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

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (domain) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (openfabric) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (protocols) */

				"domain": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `OpenFabric process name

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Domain name  |
`,
					Description: `OpenFabric process name

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Domain name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in domain, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  domain, value must match: ^[.:a-zA-Z0-9-_/]+$",
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

		"csnp_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (csnp-interval) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Complete Sequence Number Packets (CSNP) interval

    |  Format  |  Description               |
    |----------|----------------------------|
    |  1-600   |  CSNP interval in seconds  |
`,
			Description: `Complete Sequence Number Packets (CSNP) interval

    |  Format  |  Description               |
    |----------|----------------------------|
    |  1-600   |  CSNP interval in seconds  |
`,
		},

		"hello_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (hello-interval) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Hello interval

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-600   |  Hello interval in seconds  |
`,
			Description: `Hello interval

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  1-600   |  Hello interval in seconds  |
`,
		},

		"hello_multiplier":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (hello-multiplier) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Multiplier for Hello holding time

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  2-100   |  Multiplier for Hello holding time  |
`,
			Description: `Multiplier for Hello holding time

    |  Format  |  Description                        |
    |----------|-------------------------------------|
    |  2-100   |  Multiplier for Hello holding time  |
`,
		},

		"metric":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (metric) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interface metric value

    |  Format      |  Description             |
    |--------------|--------------------------|
    |  0-16777215  |  Interface metric value  |
`,
			Description: `Interface metric value

    |  Format      |  Description             |
    |--------------|--------------------------|
    |  0-16777215  |  Interface metric value  |
`,
		},

		"passive":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (passive) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not initiate adjacencies to the interface

`,
			Description: `Do not initiate adjacencies to the interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"psnp_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (psnp-interval) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Partial Sequence Number Packets (PSNP) interval

    |  Format  |  Description               |
    |----------|----------------------------|
    |  0-120   |  PSNP interval in seconds  |
`,
			Description: `Partial Sequence Number Packets (PSNP) interval

    |  Format  |  Description               |
    |----------|----------------------------|
    |  0-120   |  PSNP interval in seconds  |
`,
		},

		// TagNodes

		// Nodes

		"address_family": schema.SingleNestedAttribute{
			Attributes: ProtocolsOpenfabricDomainInterfaceAddressFamily{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Openfabric address family

`,
			Description: `Openfabric address family

`,
		},

		"password": schema.SingleNestedAttribute{
			Attributes: ProtocolsOpenfabricDomainInterfacePassword{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Authentication password for the interface

`,
			Description: `Authentication password for the interface

`,
		},
	}
}
