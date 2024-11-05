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

var _ helpers.VyosTopResourceDataModel = &ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMapping{}

// ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMapping describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMapping struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl */
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDhcpvsixServerSharedNetworkNameSubnetStaticMappingDisable       types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameSubnetStaticMappingMac           types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameSubnetStaticMappingDuID          types.String `tfsdk:"duid" vyos:"duid,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameSubnetStaticMappingIPvsixAddress types.String `tfsdk:"ipv6_address" vyos:"ipv6-address,omitempty"`
	LeafServiceDhcpvsixServerSharedNetworkNameSubnetStaticMappingIPvsixPrefix  types.String `tfsdk:"ipv6_prefix" vyos:"ipv6-prefix,omitempty"`

	// TagNodes

	// Nodes

	NodeServiceDhcpvsixServerSharedNetworkNameSubnetStaticMappingOption *ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMappingOption `tfsdk:"option" vyos:"option,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMapping) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMapping) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMapping) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMapping) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"static-mapping",
		o.SelfIdentifier.Attributes()["static_mapping"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMapping) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"service", // Node

		"dhcpv6-server", // Node

		"shared-network-name",
		o.SelfIdentifier.Attributes()["shared_network_name"].(types.String).ValueString(),

		"subnet",
		o.SelfIdentifier.Attributes()["subnet"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMapping) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"service", // Node

		"dhcpv6-server", // Node

		"shared-network-name",
		o.SelfIdentifier.Attributes()["shared_network_name"].(types.String).ValueString(),

		"subnet",
		o.SelfIdentifier.Attributes()["subnet"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMapping) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"static_mapping": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Hostname for static mapping reservation

`,
					Description: `Hostname for static mapping reservation

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in static_mapping, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  static_mapping, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				"shared_network_name": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `DHCPv6 shared network name

`,
					Description: `DHCPv6 shared network name

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in shared_network_name, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  shared_network_name, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				"subnet": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `IPv6 DHCP subnet for this shared network

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
					Description: `IPv6 DHCP subnet for this shared network

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in subnet, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  subnet, value must match: ^[.:a-zA-Z0-9-_/]+$",
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

		"disable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mac":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
			Description: `Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
		},

		"duid":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `DHCP unique identifier (DUID) to be sent by client

    |  Format  |  Description             |
    |----------|--------------------------|
    |  duid    |  DHCP unique identifier  |
`,
			Description: `DHCP unique identifier (DUID) to be sent by client

    |  Format  |  Description             |
    |----------|--------------------------|
    |  duid    |  DHCP unique identifier  |
`,
		},

		"ipv6_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client IPv6 address for this static mapping

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  ipv6    |  IPv6 address for this static mapping  |
`,
			Description: `Client IPv6 address for this static mapping

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  ipv6    |  IPv6 address for this static mapping  |
`,
		},

		"ipv6_prefix":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client IPv6 prefix for this static mapping

    |  Format   |  Description                          |
    |-----------|---------------------------------------|
    |  ipv6net  |  IPv6 prefix for this static mapping  |
`,
			Description: `Client IPv6 prefix for this static mapping

    |  Format   |  Description                          |
    |-----------|---------------------------------------|
    |  ipv6net  |  IPv6 prefix for this static mapping  |
`,
		},

		// TagNodes

		// Nodes

		"option": schema.SingleNestedAttribute{
			Attributes: ServiceDhcpvsixServerSharedNetworkNameSubnetStaticMappingOption{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCPv6 option

`,
			Description: `DHCPv6 option

`,
		},
	}
}
