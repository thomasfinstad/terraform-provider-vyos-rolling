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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (virtual-server) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &HighAvailabilityVirtualServer{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (virtual-server) */
// HighAvailabilityVirtualServerSelfIdentifier used by TagNodes to keep the id info
type HighAvailabilityVirtualServerSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (virtual-server) */

	HighAvailabilityVirtualServer types.String `tfsdk:"virtual_server" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (high-availability) */
}

// HighAvailabilityVirtualServer describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type HighAvailabilityVirtualServer struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *HighAvailabilityVirtualServerSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafHighAvailabilityVirtualServerAddress            types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafHighAvailabilityVirtualServerAlgorithm          types.String `tfsdk:"algorithm" vyos:"algorithm,omitempty"`
	LeafHighAvailabilityVirtualServerDelayLoop          types.Number `tfsdk:"delay_loop" vyos:"delay-loop,omitempty"`
	LeafHighAvailabilityVirtualServerForwardMethod      types.String `tfsdk:"forward_method" vyos:"forward-method,omitempty"`
	LeafHighAvailabilityVirtualServerFwmark             types.Number `tfsdk:"fwmark" vyos:"fwmark,omitempty"`
	LeafHighAvailabilityVirtualServerPort               types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafHighAvailabilityVirtualServerPersistenceTimeout types.Number `tfsdk:"persistence_timeout" vyos:"persistence-timeout,omitempty"`
	LeafHighAvailabilityVirtualServerProtocol           types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`

	// TagNodes

	ExistsTagHighAvailabilityVirtualServerRealServer bool `tfsdk:"-" vyos:"real-server,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *HighAvailabilityVirtualServer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *HighAvailabilityVirtualServer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *HighAvailabilityVirtualServer) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *HighAvailabilityVirtualServer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"virtual-server",
		o.SelfIdentifier.HighAvailabilityVirtualServer.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *HighAvailabilityVirtualServer) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (high-availability) */
		"high-availability", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *HighAvailabilityVirtualServer) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (high-availability) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVirtualServer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"virtual_server": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Load-balancing virtual server alias

`,
					Description: `Load-balancing virtual server alias

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in virtual_server, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  virtual_server, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (high-availability) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
    |  ipv6    |  IPv6 address  |
`,
			Description: `IP address

    |  Format  |  Description   |
    |----------|----------------|
    |  ipv4    |  IPv4 address  |
    |  ipv6    |  IPv6 address  |
`,
		},

		"algorithm":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (algorithm) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Schedule algorithm (default - least-connection)

    |  Format                           |  Description                      |
    |-----------------------------------|-----------------------------------|
    |  round-robin                      |  Round robin                      |
    |  weighted-round-robin             |  Weighted round robin             |
    |  least-connection                 |  Least connection                 |
    |  weighted-least-connection        |  Weighted least connection        |
    |  source-hashing                   |  Source hashing                   |
    |  destination-hashing              |  Destination hashing              |
    |  locality-based-least-connection  |  Locality-Based least connection  |
`,
			Description: `Schedule algorithm (default - least-connection)

    |  Format                           |  Description                      |
    |-----------------------------------|-----------------------------------|
    |  round-robin                      |  Round robin                      |
    |  weighted-round-robin             |  Weighted round robin             |
    |  least-connection                 |  Least connection                 |
    |  weighted-least-connection        |  Weighted least connection        |
    |  source-hashing                   |  Source hashing                   |
    |  destination-hashing              |  Destination hashing              |
    |  locality-based-least-connection  |  Locality-Based least connection  |
`,

			// Default:          stringdefault.StaticString(`least-connection`),
			Computed: true,
		},

		"delay_loop":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (delay-loop) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval between health-checks (in seconds)

    |  Format  |  Description          |
    |----------|-----------------------|
    |  1-600   |  Interval in seconds  |
`,
			Description: `Interval between health-checks (in seconds)

    |  Format  |  Description          |
    |----------|-----------------------|
    |  1-600   |  Interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"forward_method":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (forward-method) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Forwarding method

    |  Format  |  Description     |
    |----------|------------------|
    |  direct  |  Direct routing  |
    |  nat     |  NAT             |
    |  tunnel  |  Tunneling       |
`,
			Description: `Forwarding method

    |  Format  |  Description     |
    |----------|------------------|
    |  direct  |  Direct routing  |
    |  nat     |  NAT             |
    |  tunnel  |  Tunneling       |
`,

			// Default:          stringdefault.StaticString(`nat`),
			Computed: true,
		},

		"fwmark":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (fwmark) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match fwmark value

    |  Format        |  Description                |
    |----------------|-----------------------------|
    |  1-2147483647  |  Match firewall mark value  |
`,
			Description: `Match fwmark value

    |  Format        |  Description                |
    |----------------|-----------------------------|
    |  1-2147483647  |  Match firewall mark value  |
`,
		},

		"port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (port) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  0-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  0-65535  |  Numeric IP port  |
`,
		},

		"persistence_timeout":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (persistence-timeout) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout for persistent connections

    |  Format   |  Description                         |
    |-----------|--------------------------------------|
    |  1-86400  |  Timeout for persistent connections  |
`,
			Description: `Timeout for persistent connections

    |  Format   |  Description                         |
    |-----------|--------------------------------------|
    |  1-86400  |  Timeout for persistent connections  |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"protocol":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (protocol) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol for port checks

    |  Format  |  Description  |
    |----------|---------------|
    |  tcp     |  TCP          |
    |  udp     |  UDP          |
`,
			Description: `Protocol for port checks

    |  Format  |  Description  |
    |----------|---------------|
    |  tcp     |  TCP          |
    |  udp     |  UDP          |
`,

			// Default:          stringdefault.StaticString(`tcp`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
