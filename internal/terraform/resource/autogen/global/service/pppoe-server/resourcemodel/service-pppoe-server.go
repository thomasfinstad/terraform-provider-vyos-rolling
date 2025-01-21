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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (pppoe-server) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServicePppoeServer{}

// ServicePppoeServer describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServicePppoeServer struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServicePppoeServerAccessConcentrator    types.String `tfsdk:"access_concentrator" vyos:"access-concentrator,omitempty"`
	LeafServicePppoeServerServiceName           types.List   `tfsdk:"service_name" vyos:"service-name,omitempty"`
	LeafServicePppoeServerAcceptAnyService      types.Bool   `tfsdk:"accept_any_service" vyos:"accept-any-service,omitempty"`
	LeafServicePppoeServerAcceptBlankService    types.Bool   `tfsdk:"accept_blank_service" vyos:"accept-blank-service,omitempty"`
	LeafServicePppoeServerSessionControl        types.String `tfsdk:"session_control" vyos:"session-control,omitempty"`
	LeafServicePppoeServerDefaultPool           types.String `tfsdk:"default_pool" vyos:"default-pool,omitempty"`
	LeafServicePppoeServerDefaultIPvsixPool     types.String `tfsdk:"default_ipv6_pool" vyos:"default-ipv6-pool,omitempty"`
	LeafServicePppoeServerGatewayAddress        types.String `tfsdk:"gateway_address" vyos:"gateway-address,omitempty"`
	LeafServicePppoeServerMaxConcurrentSessions types.Number `tfsdk:"max_concurrent_sessions" vyos:"max-concurrent-sessions,omitempty"`
	LeafServicePppoeServerMtu                   types.String `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafServicePppoeServerWinsServer            types.List   `tfsdk:"wins_server" vyos:"wins-server,omitempty"`
	LeafServicePppoeServerDescrIPtion           types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafServicePppoeServerNameServer            types.List   `tfsdk:"name_server" vyos:"name-server,omitempty"`

	// TagNodes

	ExistsTagServicePppoeServerInterface bool `tfsdk:"-" vyos:"interface,child"`

	ExistsTagServicePppoeServerPadoDelay bool `tfsdk:"-" vyos:"pado-delay,child"`

	ExistsTagServicePppoeServerClientIPPool bool `tfsdk:"-" vyos:"client-ip-pool,child"`

	ExistsTagServicePppoeServerClientIPvsixPool bool `tfsdk:"-" vyos:"client-ipv6-pool,child"`

	// Nodes

	ExistsNodeServicePppoeServerAuthentication bool `tfsdk:"-" vyos:"authentication,child"`

	ExistsNodeServicePppoeServerExtendedScrIPts bool `tfsdk:"-" vyos:"extended-scripts,child"`

	ExistsNodeServicePppoeServerLimits bool `tfsdk:"-" vyos:"limits,child"`

	ExistsNodeServicePppoeServerPppOptions bool `tfsdk:"-" vyos:"ppp-options,child"`

	ExistsNodeServicePppoeServerShaper bool `tfsdk:"-" vyos:"shaper,child"`

	ExistsNodeServicePppoeServerSnmp bool `tfsdk:"-" vyos:"snmp,child"`

	ExistsNodeServicePppoeServerLog bool `tfsdk:"-" vyos:"log,child"`
}

// SetID configures the resource ID
func (o *ServicePppoeServer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServicePppoeServer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServicePppoeServer) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServicePppoeServer) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"pppoe-server",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServicePppoeServer) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */
		"service", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServicePppoeServer) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (service) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServicePppoeServer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"access_concentrator":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (access-concentrator) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Access concentrator name

`,
			Description: `Access concentrator name

`,

			// Default:          stringdefault.StaticString(`vyos-ac`),
			Computed: true,
		},

		"service_name":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (service-name) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Service name

`,
			Description: `Service name

`,
		},

		"accept_any_service":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (accept-any-service) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Accept any service name in PPPoE Active Discovery Request (PADR)

`,
			Description: `Accept any service name in PPPoE Active Discovery Request (PADR)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"accept_blank_service":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (accept-blank-service) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Accept blank service name in PADR

`,
			Description: `Accept blank service name in PADR

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"session_control":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (session-control) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `control sessions count

    |  Format   |  Description                                        |
    |-----------|-----------------------------------------------------|
    |  disable  |  Disables session control                           |
    |  deny     |  Deny second session authorization                  |
    |  replace  |  Terminate first session when second is authorized  |
`,
			Description: `control sessions count

    |  Format   |  Description                                        |
    |-----------|-----------------------------------------------------|
    |  disable  |  Disables session control                           |
    |  deny     |  Deny second session authorization                  |
    |  replace  |  Terminate first session when second is authorized  |
`,

			// Default:          stringdefault.StaticString(`replace`),
			Computed: true,
		},

		"default_pool":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-pool) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default client IP pool name

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  Default IP pool  |
`,
			Description: `Default client IP pool name

    |  Format  |  Description      |
    |----------|-------------------|
    |  txt     |  Default IP pool  |
`,
		},

		"default_ipv6_pool":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (default-ipv6-pool) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Default client IPv6 pool name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  Default IPv6 pool  |
`,
			Description: `Default client IPv6 pool name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  Default IPv6 pool  |
`,
		},

		"gateway_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (gateway-address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Gateway IP address

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  ipv4    |  Default Gateway send to the client  |
`,
			Description: `Gateway IP address

    |  Format  |  Description                         |
    |----------|--------------------------------------|
    |  ipv4    |  Default Gateway send to the client  |
`,
		},

		"max_concurrent_sessions":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (max-concurrent-sessions) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum number of concurrent session start attempts

    |  Format   |  Description                                          |
    |-----------|-------------------------------------------------------|
    |  0-65535  |  Maximum number of concurrent session start attempts  |
`,
			Description: `Maximum number of concurrent session start attempts

    |  Format   |  Description                                          |
    |-----------|-------------------------------------------------------|
    |  0-65535  |  Maximum number of concurrent session start attempts  |
`,
		},

		"mtu":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mtu) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

`,
			Description: `Maximum Transmission Unit (MTU)

`,

			// Default:          stringdefault.StaticString(`1492`),
			Computed: true,
		},

		"wins_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (wins-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Windows Internet Name Service (WINS) servers propagated to client

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv4    |  Domain Name Server (DNS) IPv4 address  |
`,
			Description: `Windows Internet Name Service (WINS) servers propagated to client

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv4    |  Domain Name Server (DNS) IPv4 address  |
`,
		},

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

		"name_server":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (name-server) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv4    |  Domain Name Server (DNS) IPv4 address  |
    |  ipv6    |  Domain Name Server (DNS) IPv6 address  |
`,
			Description: `Domain Name Servers (DNS) addresses

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  ipv4    |  Domain Name Server (DNS) IPv4 address  |
    |  ipv6    |  Domain Name Server (DNS) IPv6 address  |
`,
		},

		// TagNodes

		// Nodes

	}
}
