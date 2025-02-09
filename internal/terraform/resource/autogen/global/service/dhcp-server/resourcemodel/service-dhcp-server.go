/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (dhcp-server) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceDhcpServer{}

// ServiceDhcpServer describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServiceDhcpServer struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDhcpServerDisable          types.Bool `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafServiceDhcpServerDynamicDNSUpdate types.Bool `tfsdk:"dynamic_dns_update" vyos:"dynamic-dns-update,omitempty"`
	LeafServiceDhcpServerHostfileUpdate   types.Bool `tfsdk:"hostfile_update" vyos:"hostfile-update,omitempty"`
	LeafServiceDhcpServerListenAddress    types.List `tfsdk:"listen_address" vyos:"listen-address,omitempty"`
	LeafServiceDhcpServerListenInterface  types.List `tfsdk:"listen_interface" vyos:"listen-interface,omitempty"`

	// TagNodes

	ExistsTagServiceDhcpServerSharedNetworkName bool `tfsdk:"-" vyos:"shared-network-name,child"`

	// Nodes

	ExistsNodeServiceDhcpServerHighAvailability bool `tfsdk:"-" vyos:"high-availability,child"`
}

// SetID configures the resource ID
func (o *ServiceDhcpServer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDhcpServer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDhcpServer) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDhcpServer) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"dhcp-server",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceDhcpServer) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */
		"service", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceDhcpServer) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (service) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpServer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

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

		"dynamic_dns_update":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (dynamic-dns-update) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Dynamically update Domain Name System (RFC4702)

`,
			Description: `Dynamically update Domain Name System (RFC4702)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"hostfile_update":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (hostfile-update) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Updating /etc/hosts file (per client lease)

`,
			Description: `Updating /etc/hosts file (per client lease)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"listen_address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (listen-address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local IPv4 addresses to listen on

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  ipv4    |  IPv4 address to listen for incoming connections  |
`,
			Description: `Local IPv4 addresses to listen on

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  ipv4    |  IPv4 address to listen for incoming connections  |
`,
		},

		"listen_interface":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (listen-interface) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface to listen on

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Interface to listen on

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		// TagNodes

		// Nodes

	}
}
