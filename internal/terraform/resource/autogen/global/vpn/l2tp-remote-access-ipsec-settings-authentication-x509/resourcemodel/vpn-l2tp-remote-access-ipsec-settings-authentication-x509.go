/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/global/resource-model.gotmpl */
// Validate compliance
var _ helpers.VyosTopResourceDataModel = &VpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine{}

// VpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine describes the resource data model.
type VpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafVpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronineCertificate   types.String `tfsdk:"certificate" vyos:"certificate,omitempty"`
	LeafVpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeroninePassphrase    types.String `tfsdk:"passphrase" vyos:"passphrase,omitempty"`
	LeafVpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronineCaCertificate types.List   `tfsdk:"ca_certificate" vyos:"ca-certificate,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *VpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"x509",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */
		"vpn",

		"l2tp",

		"remote-access",

		"ipsec-settings",

		"authentication",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *VpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnLtwotpRemoteAccessIPsecSettingsAuthenticationXfivezeronine) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"certificate":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate in PKI configuration

    |  Format  |  Description                               |
    |----------|--------------------------------------------|
    |  txt     |  Name of certificate in PKI configuration  |
`,
			Description: `Certificate in PKI configuration

    |  Format  |  Description                               |
    |----------|--------------------------------------------|
    |  txt     |  Name of certificate in PKI configuration  |
`,
		},

		"passphrase":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Private key passphrase

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  txt     |  Passphrase to decrypt the private key  |
`,
			Description: `Private key passphrase

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  txt     |  Passphrase to decrypt the private key  |
`,
		},

		"ca_certificate":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Certificate Authority chain in PKI configuration

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  txt     |  Name of CA in PKI configuration  |
`,
			Description: `Certificate Authority chain in PKI configuration

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  txt     |  Name of CA in PKI configuration  |
`,
		},
	}
}
