/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (dhcp-options) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesVirtualEthernetVifDhcpOptions{}

// InterfacesVirtualEthernetVifDhcpOptions describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesVirtualEthernetVifDhcpOptions struct {
	// LeafNodes
	LeafInterfacesVirtualEthernetVifDhcpOptionsClientID             types.String `tfsdk:"client_id" vyos:"client-id,omitempty"`
	LeafInterfacesVirtualEthernetVifDhcpOptionsHostName             types.String `tfsdk:"host_name" vyos:"host-name,omitempty"`
	LeafInterfacesVirtualEthernetVifDhcpOptionsMtu                  types.Bool   `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesVirtualEthernetVifDhcpOptionsVendorClassID        types.String `tfsdk:"vendor_class_id" vyos:"vendor-class-id,omitempty"`
	LeafInterfacesVirtualEthernetVifDhcpOptionsUserClass            types.String `tfsdk:"user_class" vyos:"user-class,omitempty"`
	LeafInterfacesVirtualEthernetVifDhcpOptionsNoDefaultRoute       types.Bool   `tfsdk:"no_default_route" vyos:"no-default-route,omitempty"`
	LeafInterfacesVirtualEthernetVifDhcpOptionsDefaultRouteDistance types.Number `tfsdk:"default_route_distance" vyos:"default-route-distance,omitempty"`
	LeafInterfacesVirtualEthernetVifDhcpOptionsReject               types.List   `tfsdk:"reject" vyos:"reject,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesVirtualEthernetVifDhcpOptions) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"client_id":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (client-id) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Identifier used by client to identify itself to the DHCP server

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
			Description: `Identifier used by client to identify itself to the DHCP server

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
		},

		"host_name":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (host-name) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override system host-name sent to DHCP server

`,
			Description: `Override system host-name sent to DHCP server

`,
		},

		"mtu":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mtu) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use MTU value from DHCP server - ignore interface setting

`,
			Description: `Use MTU value from DHCP server - ignore interface setting

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"vendor_class_id":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (vendor-class-id) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Identify the vendor client type to the DHCP server

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
			Description: `Identify the vendor client type to the DHCP server

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
		},

		"user_class":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (user-class) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Identify to the DHCP server, user configurable option

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
			Description: `Identify to the DHCP server, user configurable option

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  DHCP option string  |
`,
		},

		"no_default_route":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (no-default-route) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Do not install default route to system

`,
			Description: `Do not install default route to system

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"default_route_distance":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-route-distance) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Distance for installed default route

    |  Format  |  Description                                              |
    |----------|-----------------------------------------------------------|
    |  1-255   |  Distance for the default route received from the server  |
`,
			Description: `Distance for installed default route

    |  Format  |  Description                                              |
    |----------|-----------------------------------------------------------|
    |  1-255   |  Distance for the default route received from the server  |
`,

			// Default:          stringdefault.StaticString(`210`),
			Computed: true,
		},

		"reject":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (reject) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP addresses or subnets from which to reject DHCP leases

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  ipv4     |  IPv4 address to match  |
    |  ipv4net  |  IPv4 prefix to match   |
`,
			Description: `IP addresses or subnets from which to reject DHCP leases

    |  Format   |  Description            |
    |-----------|-------------------------|
    |  ipv4     |  IPv4 address to match  |
    |  ipv4net  |  IPv4 prefix to match   |
`,
		},

		// TagNodes

		// Nodes

	}
}
