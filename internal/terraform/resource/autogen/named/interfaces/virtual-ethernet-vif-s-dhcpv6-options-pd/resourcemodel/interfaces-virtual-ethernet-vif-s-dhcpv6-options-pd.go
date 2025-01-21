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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (pd) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &InterfacesVirtualEthernetVifSDhcpvsixOptionsPd{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (pd) */
// InterfacesVirtualEthernetVifSDhcpvsixOptionsPdSelfIdentifier used by TagNodes to keep the id info
type InterfacesVirtualEthernetVifSDhcpvsixOptionsPdSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (pd) */

	InterfacesVirtualEthernetVifSDhcpvsixOptionsPd types.String `tfsdk:"pd" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (dhcpv6-options) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (vif-s) */

	InterfacesVirtualEthernetVifS types.Number `tfsdk:"vif_s" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (virtual-ethernet) */

	InterfacesVirtualEthernet types.String `tfsdk:"virtual_ethernet" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (interfaces) */
}

// InterfacesVirtualEthernetVifSDhcpvsixOptionsPd describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type InterfacesVirtualEthernetVifSDhcpvsixOptionsPd struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *InterfacesVirtualEthernetVifSDhcpvsixOptionsPdSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafInterfacesVirtualEthernetVifSDhcpvsixOptionsPdLength types.Number `tfsdk:"length" vyos:"length,omitempty"`

	// TagNodes

	ExistsTagInterfacesVirtualEthernetVifSDhcpvsixOptionsPdInterface bool `tfsdk:"-" vyos:"interface,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *InterfacesVirtualEthernetVifSDhcpvsixOptionsPd) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesVirtualEthernetVifSDhcpvsixOptionsPd) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesVirtualEthernetVifSDhcpvsixOptionsPd) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesVirtualEthernetVifSDhcpvsixOptionsPd) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"pd",
		o.SelfIdentifier.InterfacesVirtualEthernetVifSDhcpvsixOptionsPd.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesVirtualEthernetVifSDhcpvsixOptionsPd) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (dhcpv6-options) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vif-s) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (virtual-ethernet) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (interfaces) */
		"interfaces", // Node

		"virtual-ethernet",
		o.SelfIdentifier.InterfacesVirtualEthernet.ValueString(),

		"vif-s",
		o.SelfIdentifier.InterfacesVirtualEthernetVifS.ValueBigFloat().String(),

		"dhcpv6-options", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesVirtualEthernetVifSDhcpvsixOptionsPd) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (dhcpv6-options) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (vif-s) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vif-s) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (virtual-ethernet) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (interfaces) */
		"interfaces", // Node

		"virtual-ethernet",
		o.SelfIdentifier.InterfacesVirtualEthernet.ValueString(),

		"vif-s",
		o.SelfIdentifier.InterfacesVirtualEthernetVifS.ValueBigFloat().String(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesVirtualEthernetVifSDhcpvsixOptionsPd) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"pd": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `DHCPv6 prefix delegation interface statement

    |  Format           |  Description                        |
    |-------------------|-------------------------------------|
    |  instance number  |  Prefix delegation instance (>= 0)  |
`,
					Description: `DHCPv6 prefix delegation interface statement

    |  Format           |  Description                        |
    |-------------------|-------------------------------------|
    |  instance number  |  Prefix delegation instance (>= 0)  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in pd, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  pd, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (dhcpv6-options) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (vif-s) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (virtual-ethernet) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (interfaces) */

				"virtual_ethernet": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Virtual Ethernet (veth) Interface

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  vethN   |  Virtual Ethernet interface name  |
`,
					Description: `Virtual Ethernet (veth) Interface

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  vethN   |  Virtual Ethernet interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in virtual_ethernet, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  virtual_ethernet, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				"vif_s": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `QinQ TAG-S Virtual Local Area Network (VLAN) ID

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  0-4094  |  QinQ Virtual Local Area Network (VLAN) ID  |
`,
					Description: `QinQ TAG-S Virtual Local Area Network (VLAN) ID

    |  Format  |  Description                                |
    |----------|---------------------------------------------|
    |  0-4094  |  QinQ Virtual Local Area Network (VLAN) ID  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"length":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (length) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Request IPv6 prefix length from peer

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  32-64   |  Length of delegated prefix  |
`,
			Description: `Request IPv6 prefix length from peer

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  32-64   |  Length of delegated prefix  |
`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
