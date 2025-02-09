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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (static-mapping) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceDhcpServerSharedNetworkNameSubnetStaticMapping{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (static-mapping) */
// ServiceDhcpServerSharedNetworkNameSubnetStaticMappingSelfIdentifier used by TagNodes to keep the id info
type ServiceDhcpServerSharedNetworkNameSubnetStaticMappingSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (static-mapping) */

	ServiceDhcpServerSharedNetworkNameSubnetStaticMapping types.String `tfsdk:"static_mapping" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (subnet) */

	ServiceDhcpServerSharedNetworkNameSubnet types.String `tfsdk:"subnet" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (shared-network-name) */

	ServiceDhcpServerSharedNetworkName types.String `tfsdk:"shared_network_name" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (dhcp-server) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (service) */
}

// ServiceDhcpServerSharedNetworkNameSubnetStaticMapping describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ServiceDhcpServerSharedNetworkNameSubnetStaticMapping struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *ServiceDhcpServerSharedNetworkNameSubnetStaticMappingSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingDisable     types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingIPAddress   types.String `tfsdk:"ip_address" vyos:"ip-address,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingMac         types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafServiceDhcpServerSharedNetworkNameSubnetStaticMappingDuID        types.String `tfsdk:"duid" vyos:"duid,omitempty"`

	// TagNodes

	// Nodes

	NodeServiceDhcpServerSharedNetworkNameSubnetStaticMappingOption *ServiceDhcpServerSharedNetworkNameSubnetStaticMappingOption `tfsdk:"option" vyos:"option,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"static-mapping",
		o.SelfIdentifier.ServiceDhcpServerSharedNetworkNameSubnetStaticMapping.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (subnet) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (shared-network-name) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (dhcp-server) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */
		"service", // Node

		"dhcp-server", // Node

		"shared-network-name",
		o.SelfIdentifier.ServiceDhcpServerSharedNetworkName.ValueString(),

		"subnet",
		o.SelfIdentifier.ServiceDhcpServerSharedNetworkNameSubnet.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (subnet) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (shared-network-name) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (shared-network-name) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (dhcp-server) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */
		"service", // Node

		"dhcp-server", // Node

		"shared-network-name",
		o.SelfIdentifier.ServiceDhcpServerSharedNetworkName.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServerSharedNetworkNameSubnetStaticMapping) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (subnet) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (shared-network-name) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (dhcp-server) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (service) */

				"shared_network_name": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Name of DHCP shared network

`,
					Description: `Name of DHCP shared network

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
					MarkdownDescription: `DHCP subnet for shared network

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
`,
					Description: `DHCP subnet for shared network

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
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

		"description":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (description) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"disable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ip_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ip-address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Fixed IP address of static mapping

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  ipv4    |  IPv4 address used in static mapping  |
`,
			Description: `Fixed IP address of static mapping

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  ipv4    |  IPv4 address used in static mapping  |
`,
		},

		"mac":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mac) */
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

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (duid) */
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

		// TagNodes

		// Nodes

		"option": schema.SingleNestedAttribute{
			Attributes: ServiceDhcpServerSharedNetworkNameSubnetStaticMappingOption{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCP option

`,
			Description: `DHCP option

`,
		},
	}
}
