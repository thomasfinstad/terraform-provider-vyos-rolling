/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServicePppoeServerAuthenticationRadius{}

// ServicePppoeServerAuthenticationRadius describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServicePppoeServerAuthenticationRadius struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServicePppoeServerAuthenticationRadiusSourceAddress             types.String `tfsdk:"source_address" vyos:"source-address,omitempty"`
	LeafServicePppoeServerAuthenticationRadiusAccountingInterimInterval types.Number `tfsdk:"accounting_interim_interval" vyos:"accounting-interim-interval,omitempty"`
	LeafServicePppoeServerAuthenticationRadiusAcctInterimJitter         types.Number `tfsdk:"acct_interim_jitter" vyos:"acct-interim-jitter,omitempty"`
	LeafServicePppoeServerAuthenticationRadiusTimeout                   types.Number `tfsdk:"timeout" vyos:"timeout,omitempty"`
	LeafServicePppoeServerAuthenticationRadiusAcctTimeout               types.Number `tfsdk:"acct_timeout" vyos:"acct-timeout,omitempty"`
	LeafServicePppoeServerAuthenticationRadiusMaxTry                    types.Number `tfsdk:"max_try" vyos:"max-try,omitempty"`
	LeafServicePppoeServerAuthenticationRadiusNasIDentifier             types.String `tfsdk:"nas_identifier" vyos:"nas-identifier,omitempty"`
	LeafServicePppoeServerAuthenticationRadiusNasIPAddress              types.String `tfsdk:"nas_ip_address" vyos:"nas-ip-address,omitempty"`
	LeafServicePppoeServerAuthenticationRadiusPreallocateVif            types.Bool   `tfsdk:"preallocate_vif" vyos:"preallocate-vif,omitempty"`
	LeafServicePppoeServerAuthenticationRadiusCalledSIDFormat           types.String `tfsdk:"called_sid_format" vyos:"called-sid-format,omitempty"`

	// TagNodes

	ExistsTagServicePppoeServerAuthenticationRadiusServer bool `tfsdk:"-" vyos:"server,child"`

	// Nodes

	ExistsNodeServicePppoeServerAuthenticationRadiusDynamicAuthor bool `tfsdk:"-" vyos:"dynamic-author,child"`

	ExistsNodeServicePppoeServerAuthenticationRadiusRateLimit bool `tfsdk:"-" vyos:"rate-limit,child"`
}

// SetID configures the resource ID
func (o *ServicePppoeServerAuthenticationRadius) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServicePppoeServerAuthenticationRadius) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServicePppoeServerAuthenticationRadius) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServicePppoeServerAuthenticationRadius) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"radius",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServicePppoeServerAuthenticationRadius) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"service", // Node

		"pppoe-server", // Node

		"authentication", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServicePppoeServerAuthenticationRadius) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServicePppoeServerAuthenticationRadius) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"source_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IPv4 address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
`,
			Description: `IPv4 address used to initiate connection

    |  Format  |  Description          |
    |----------|-----------------------|
    |  ipv4    |  IPv4 source address  |
`,
		},

		"accounting_interim_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval in seconds to send accounting information

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  1-3600  |  Interval in seconds to send accounting information  |
`,
			Description: `Interval in seconds to send accounting information

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  1-3600  |  Interval in seconds to send accounting information  |
`,
		},

		"acct_interim_jitter":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum jitter value in seconds to be applied to accounting information interval

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  1-60    |  Maximum jitter value in seconds  |
`,
			Description: `Maximum jitter value in seconds to be applied to accounting information interval

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  1-60    |  Maximum jitter value in seconds  |
`,
		},

		"timeout":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout in seconds to wait response from RADIUS server

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-60    |  Timeout in seconds  |
`,
			Description: `Timeout in seconds to wait response from RADIUS server

    |  Format  |  Description         |
    |----------|----------------------|
    |  1-60    |  Timeout in seconds  |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"acct_timeout":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Timeout for Interim-Update packets, terminate session afterwards

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  0-60    |  Timeout in seconds, 0 to keep active  |
`,
			Description: `Timeout for Interim-Update packets, terminate session afterwards

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  0-60    |  Timeout in seconds, 0 to keep active  |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"max_try":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of tries to send Access-Request/Accounting-Request queries

    |  Format  |  Description    |
    |----------|-----------------|
    |  1-20    |  Maximum tries  |
`,
			Description: `Number of tries to send Access-Request/Accounting-Request queries

    |  Format  |  Description    |
    |----------|-----------------|
    |  1-20    |  Maximum tries  |
`,

			// Default:          stringdefault.StaticString(`3`),
			Computed: true,
		},

		"nas_identifier":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAS-Identifier attribute sent to RADIUS

`,
			Description: `NAS-Identifier attribute sent to RADIUS

`,
		},

		"nas_ip_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `NAS-IP-Address attribute sent to RADIUS

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  NAS-IP-Address attribute  |
`,
			Description: `NAS-IP-Address attribute sent to RADIUS

    |  Format  |  Description               |
    |----------|----------------------------|
    |  ipv4    |  NAS-IP-Address attribute  |
`,
		},

		"preallocate_vif":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable attribute NAS-Port-Id in Access-Request

`,
			Description: `Enable attribute NAS-Port-Id in Access-Request

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"called_sid_format":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Format of Called-Station-Id attribute

    |  Format      |  Description                                                                                            |
    |--------------|---------------------------------------------------------------------------------------------------------|
    |  ifname      |  NAS-Port-Id - should contain root interface name (NAS-Port-Id=eth1)                                    |
    |  ifname:mac  |  NAS-Port-Id - should contain root interface name and mac address (NAS-Port-Id=eth1:00:00:00:00:00:00)  |
`,
			Description: `Format of Called-Station-Id attribute

    |  Format      |  Description                                                                                            |
    |--------------|---------------------------------------------------------------------------------------------------------|
    |  ifname      |  NAS-Port-Id - should contain root interface name (NAS-Port-Id=eth1)                                    |
    |  ifname:mac  |  NAS-Port-Id - should contain root interface name and mac address (NAS-Port-Id=eth1:00:00:00:00:00:00)  |
`,
		},

		// TagNodes

		// Nodes

	}
}
