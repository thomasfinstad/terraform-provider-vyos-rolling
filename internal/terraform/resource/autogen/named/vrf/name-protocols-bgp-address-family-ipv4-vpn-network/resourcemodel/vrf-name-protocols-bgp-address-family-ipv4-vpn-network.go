// Package resourcemodel code generated by /repo/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork{}

// VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork describes the resource data model.
type VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"network_id" vyos:"-,self-id"`

	ParentIDVrfName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourVpnNetworkRd    types.String `tfsdk:"rd" vyos:"rd,omitempty"`
	LeafVrfNameProtocolsBgpAddressFamilyIPvfourVpnNetworkLabel types.Number `tfsdk:"label" vyos:"label,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"network",
		o.SelfIdentifier.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork) GetVyosParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),

		"protocols",

		"bgp",

		"address-family",

		"ipv4-vpn",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork) GetVyosNamedParentPath() []string {
	return []string{
		"vrf",

		"name",
		o.ParentIDVrfName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsBgpAddressFamilyIPvfourVpnNetwork) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"network_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Import BGP network/prefix into unicast VPN IPv4 RIB

    |  Format   |  Description                          |
    |-----------|---------------------------------------|
    |  ipv4net  |  Unicast VPN IPv4 BGP network/prefix  |
`,
			Description: `Import BGP network/prefix into unicast VPN IPv4 RIB

    |  Format   |  Description                          |
    |-----------|---------------------------------------|
    |  ipv4net  |  Unicast VPN IPv4 BGP network/prefix  |
`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in network_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  network_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"name_id": schema.StringAttribute{
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
							"double underscores in name_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illegal character in  name_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"rd": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Route Distinguisher

    |  Format                   |  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
			Description: `Route Distinguisher

    |  Format                   |  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
`,
		},

		"label": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `MPLS label value assigned to route

    |  Format     |  Description       |
    |-------------|--------------------|
    |  0-1048575  |  MPLS label value  |
`,
			Description: `MPLS label value assigned to route

    |  Format     |  Description       |
    |-------------|--------------------|
    |  0-1048575  |  MPLS label value  |
`,
		},

		// Nodes

	}
}
