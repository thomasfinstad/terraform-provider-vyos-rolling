/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (tls) */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesOpenvpnTLS{}

// InterfacesOpenvpnTLS describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesOpenvpnTLS struct {
	// LeafNodes
	LeafInterfacesOpenvpnTLSAuthKey         types.String `tfsdk:"auth_key" vyos:"auth-key,omitempty"`
	LeafInterfacesOpenvpnTLSCertificate     types.String `tfsdk:"certificate" vyos:"certificate,omitempty"`
	LeafInterfacesOpenvpnTLSCaCertificate   types.List   `tfsdk:"ca_certificate" vyos:"ca-certificate,omitempty"`
	LeafInterfacesOpenvpnTLSDhParams        types.String `tfsdk:"dh_params" vyos:"dh-params,omitempty"`
	LeafInterfacesOpenvpnTLSCryptKey        types.String `tfsdk:"crypt_key" vyos:"crypt-key,omitempty"`
	LeafInterfacesOpenvpnTLSPeerFingerprint types.List   `tfsdk:"peer_fingerprint" vyos:"peer-fingerprint,omitempty"`
	LeafInterfacesOpenvpnTLSTLSVersionMin   types.String `tfsdk:"tls_version_min" vyos:"tls-version-min,omitempty"`
	LeafInterfacesOpenvpnTLSRole            types.String `tfsdk:"role" vyos:"role,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnTLS) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"auth_key":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (auth-key) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `TLS shared secret key for tls-auth

`,
			Description: `TLS shared secret key for tls-auth

`,
		},

		"certificate":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (certificate) */
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

		"ca_certificate":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (ca-certificate) */
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

		"dh_params":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (dh-params) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Diffie Hellman parameters (server only)

`,
			Description: `Diffie Hellman parameters (server only)

`,
		},

		"crypt_key":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (crypt-key) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Static key to use to authenticate control channel

`,
			Description: `Static key to use to authenticate control channel

`,
		},

		"peer_fingerprint":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (peer-fingerprint) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Peer certificate SHA256 fingerprint

`,
			Description: `Peer certificate SHA256 fingerprint

`,
		},

		"tls_version_min":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (tls-version-min) */
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
		},

		"role":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (role) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `TLS negotiation role

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  active   |  Initiate TLS negotiation actively  |
    |  passive  |  Wait for incoming TLS connection   |
`,
			Description: `TLS negotiation role

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  active   |  Initiate TLS negotiation actively  |
    |  passive  |  Wait for incoming TLS connection   |
`,
		},

		// TagNodes

		// Nodes

	}
}
