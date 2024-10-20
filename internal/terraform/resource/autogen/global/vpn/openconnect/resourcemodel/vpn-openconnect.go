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

/* tools/generate-terraform-resource-full/templates/resources/global/resource-model.gotmpl */
// Validate compliance
var _ helpers.VyosTopResourceDataModel = &VpnOpenconnect{}

// VpnOpenconnect describes the resource data model.
type VpnOpenconnect struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnOpenconnectListenAddress       types.String `tfsdk:"listen_address" vyos:"listen-address,omitempty"`
	LeafVpnOpenconnectHTTPSecURItyHeaders types.Bool   `tfsdk:"http_security_headers" vyos:"http-security-headers,omitempty"`
	LeafVpnOpenconnectTLSVersionMin       types.String `tfsdk:"tls_version_min" vyos:"tls-version-min,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeVpnOpenconnectAccounting      bool `tfsdk:"-" vyos:"accounting,child"`
	ExistsNodeVpnOpenconnectAuthentication  bool `tfsdk:"-" vyos:"authentication,child"`
	ExistsNodeVpnOpenconnectListenPorts     bool `tfsdk:"-" vyos:"listen-ports,child"`
	ExistsNodeVpnOpenconnectSsl             bool `tfsdk:"-" vyos:"ssl,child"`
	ExistsNodeVpnOpenconnectNetworkSettings bool `tfsdk:"-" vyos:"network-settings,child"`
}

// SetID configures the resource ID
func (o *VpnOpenconnect) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnOpenconnect) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnOpenconnect) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnOpenconnect) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"openconnect",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnOpenconnect) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */
		"vpn",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *VpnOpenconnect) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnOpenconnect) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"listen_address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
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

			// Default:          stringdefault.StaticString(`0.0.0.0`),
			Computed: true,
		},

		"http_security_headers":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable HTTP security headers

`,
			Description: `Enable HTTP security headers

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"tls_version_min":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify the minimum required TLS version

    |  Format  |  Description  |
    |----------|---------------|
    |  1.0     |  TLS v1.0     |
    |  1.1     |  TLS v1.1     |
    |  1.2     |  TLS v1.2     |
    |  1.3     |  TLS v1.3     |
`,
			Description: `Specify the minimum required TLS version

    |  Format  |  Description  |
    |----------|---------------|
    |  1.0     |  TLS v1.0     |
    |  1.1     |  TLS v1.1     |
    |  1.2     |  TLS v1.2     |
    |  1.3     |  TLS v1.3     |
`,

			// Default:          stringdefault.StaticString(`1.2`),
			Computed: true,
		},
	}
}
