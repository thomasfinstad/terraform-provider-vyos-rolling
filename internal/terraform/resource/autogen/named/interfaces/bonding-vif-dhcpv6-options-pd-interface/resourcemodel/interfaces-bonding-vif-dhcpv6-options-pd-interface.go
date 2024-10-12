// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
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

var _ helpers.VyosTopResourceDataModel = &InterfacesBondingVifDhcpvsixOptionsPdInterface{}

// InterfacesBondingVifDhcpvsixOptionsPdInterface describes the resource data model.
type InterfacesBondingVifDhcpvsixOptionsPdInterface struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafInterfacesBondingVifDhcpvsixOptionsPdInterfaceAddress types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesBondingVifDhcpvsixOptionsPdInterfaceSLAID   types.Number `tfsdk:"sla_id" vyos:"sla-id,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *InterfacesBondingVifDhcpvsixOptionsPdInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesBondingVifDhcpvsixOptionsPdInterface) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesBondingVifDhcpvsixOptionsPdInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesBondingVifDhcpvsixOptionsPdInterface) GetVyosPath() []string {
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
func (o *InterfacesBondingVifDhcpvsixOptionsPdInterface) GetVyosParentPath() []string {
	return []string{
		"interfaces",

		"bonding",

		o.SelfIdentifier.Attributes()["bonding"].(types.String).ValueString(),

		"vif",

		o.SelfIdentifier.Attributes()["vif"].(types.Number).ValueBigFloat().String(),

		"dhcpv6-options",

		"pd",

		o.SelfIdentifier.Attributes()["pd"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesBondingVifDhcpvsixOptionsPdInterface) GetVyosNamedParentPath() []string {
	return []string{
		"interfaces",

		"bonding",

		o.SelfIdentifier.Attributes()["bonding"].(types.String).ValueString(),

		"vif",

		o.SelfIdentifier.Attributes()["vif"].(types.Number).ValueBigFloat().String(),

		"dhcpv6-options",

		"pd",

		o.SelfIdentifier.Attributes()["pd"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesBondingVifDhcpvsixOptionsPdInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
					MarkdownDescription: `Delegate IPv6 prefix from provider to this interface

`,
					Description: `Delegate IPv6 prefix from provider to this interface

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
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  interface, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},

				"bonding": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Bonding Interface/Link Aggregation

    |  Format  |  Description             |
    |----------|--------------------------|
    |  bondN   |  Bonding interface name  |
`,
					Description: `Bonding Interface/Link Aggregation

    |  Format  |  Description             |
    |----------|--------------------------|
    |  bondN   |  Bonding interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in bonding, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  bonding, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},

				"vif": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `Virtual Local Area Network (VLAN) ID

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  0-4094  |  Virtual Local Area Network (VLAN) ID  |
`,
					Description: `Virtual Local Area Network (VLAN) ID

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  0-4094  |  Virtual Local Area Network (VLAN) ID  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

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
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local interface address assigned to interface (default: EUI-64)

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  >0      |  Used to form IPv6 interface address  |
`,
			Description: `Local interface address assigned to interface (default: EUI-64)

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  >0      |  Used to form IPv6 interface address  |
`,
		},

		"sla_id": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interface site-Level aggregator (SLA)

    |  Format   |  Description                                          |
    |-----------|-------------------------------------------------------|
    |  0-65535  |  Decimal integer which fits in the length of SLA IDs  |
`,
			Description: `Interface site-Level aggregator (SLA)

    |  Format   |  Description                                          |
    |-----------|-------------------------------------------------------|
    |  0-65535  |  Decimal integer which fits in the length of SLA IDs  |
`,
		},

		// Nodes

	}
}
