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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (dhcp-relay) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceDhcpRelay{}

// ServiceDhcpRelay describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServiceDhcpRelay struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceDhcpRelayDisable           types.Bool `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafServiceDhcpRelayInterface         types.List `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafServiceDhcpRelayListenInterface   types.List `tfsdk:"listen_interface" vyos:"listen-interface,omitempty"`
	LeafServiceDhcpRelayUpstreamInterface types.List `tfsdk:"upstream_interface" vyos:"upstream-interface,omitempty"`
	LeafServiceDhcpRelayServer            types.List `tfsdk:"server" vyos:"server,omitempty"`

	// TagNodes

	// Nodes

	ExistsNodeServiceDhcpRelayRelayOptions bool `tfsdk:"-" vyos:"relay-options,child"`
}

// SetID configures the resource ID
func (o *ServiceDhcpRelay) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceDhcpRelay) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceDhcpRelay) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceDhcpRelay) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"dhcp-relay",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceDhcpRelay) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */
		"service", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceDhcpRelay) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (service) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceDhcpRelay) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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

		"interface":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (interface) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Interface

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		"listen_interface":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (listen-interface) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface for DHCP Relay Agent to listen for requests

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Interface for DHCP Relay Agent to listen for requests

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		"upstream_interface":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (upstream-interface) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface for DHCP Relay Agent forward requests out

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Interface for DHCP Relay Agent forward requests out

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		"server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `DHCP server address

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  DHCP server IPv4 address  |
`,
			Description: `DHCP server address

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  DHCP server IPv4 address  |
`,
		},

		// TagNodes

		// Nodes

	}
}
