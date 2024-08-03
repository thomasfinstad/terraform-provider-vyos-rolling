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

var _ helpers.VyosTopResourceDataModel = &InterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd{}

// InterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd describes the resource data model.
type InterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafInterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPdLength types.Number `tfsdk:"length" vyos:"length,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagInterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPdInterface bool `tfsdk:"-" vyos:"interface,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *InterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"pd",
		o.SelfIdentifier.Attributes()["pd"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd) GetVyosParentPath() []string {
	return []string{
		"interfaces",

		"virtual-ethernet",

		o.SelfIdentifier.Attributes()["virtual_ethernet"].(types.String).ValueString(),

		"vif-s",

		o.SelfIdentifier.Attributes()["vif_s"].(types.Number).ValueBigFloat().String(),

		"vif-c",

		o.SelfIdentifier.Attributes()["vif_c"].(types.String).ValueString(),

		"dhcpv6-options",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd) GetVyosNamedParentPath() []string {
	return []string{
		"interfaces",

		"virtual-ethernet",

		o.SelfIdentifier.Attributes()["virtual_ethernet"].(types.String).ValueString(),

		"vif-s",

		o.SelfIdentifier.Attributes()["vif_s"].(types.Number).ValueBigFloat().String(),

		"vif-c",

		o.SelfIdentifier.Attributes()["vif_c"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesVirtualEthernetVifSVifCDhcpvsixOptionsPd) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
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
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  pd, value must match: ^[a-zA-Z0-9-_]*$",
								),
							),
						},
					},

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
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  virtual_ethernet, value must match: ^[a-zA-Z0-9-_]*$",
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

					"vif_c": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `QinQ TAG-C Virtual Local Area Network (VLAN) ID

`,
						Description: `QinQ TAG-C Virtual Local Area Network (VLAN) ID

`,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						}, Validators: []validator.String{
							stringvalidator.All(
								helpers.StringNot(
									stringvalidator.RegexMatches(
										regexp.MustCompile(`^.*__.*$`),
										"double underscores in vif_c, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  vif_c, value must match: ^[a-zA-Z0-9-_]*$",
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

		"length": schema.NumberAttribute{
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

		// Nodes

	}
}
