// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VpnIPsecRemoteAccessConnectionAuthentication{}

// VpnIPsecRemoteAccessConnectionAuthentication describes the resource data model.
type VpnIPsecRemoteAccessConnectionAuthentication struct {
	// LeafNodes
	LeafVpnIPsecRemoteAccessConnectionAuthenticationLocalID         types.String `tfsdk:"local_id" vyos:"local-id,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionAuthenticationEapID           types.String `tfsdk:"eap_id" vyos:"eap-id,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionAuthenticationClientMode      types.String `tfsdk:"client_mode" vyos:"client-mode,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionAuthenticationServerMode      types.String `tfsdk:"server_mode" vyos:"server-mode,omitempty"`
	LeafVpnIPsecRemoteAccessConnectionAuthenticationPreSharedSecret types.String `tfsdk:"pre_shared_secret" vyos:"pre-shared-secret,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine *VpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine `tfsdk:"x509" vyos:"x509,omitempty"`
	NodeVpnIPsecRemoteAccessConnectionAuthenticationLocalUsers    *VpnIPsecRemoteAccessConnectionAuthenticationLocalUsers    `tfsdk:"local_users" vyos:"local-users,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecRemoteAccessConnectionAuthentication) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"local_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local ID for peer authentication

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  txt     |  Local ID used for peer authentication  |
`,
			Description: `Local ID for peer authentication

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  txt     |  Local ID used for peer authentication  |
`,
		},

		"eap_id": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Remote EAP ID for client authentication

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  txt     |  Remote EAP ID for client authentication  |
    |  any     |  Allow any EAP ID                         |
`,
			Description: `Remote EAP ID for client authentication

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  txt     |  Remote EAP ID for client authentication  |
    |  any     |  Allow any EAP ID                         |
`,

			// Default:          stringdefault.StaticString(`any`),
			Computed: true,
		},

		"client_mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client authentication mode

    |  Format        |  Description                                 |
    |----------------|----------------------------------------------|
    |  x509          |  Use IPsec x.509 certificate authentication  |
    |  eap-tls       |  Use EAP-TLS authentication                  |
    |  eap-mschapv2  |  Use EAP-MSCHAPv2 authentication             |
    |  eap-radius    |  Use EAP-RADIUS authentication               |
`,
			Description: `Client authentication mode

    |  Format        |  Description                                 |
    |----------------|----------------------------------------------|
    |  x509          |  Use IPsec x.509 certificate authentication  |
    |  eap-tls       |  Use EAP-TLS authentication                  |
    |  eap-mschapv2  |  Use EAP-MSCHAPv2 authentication             |
    |  eap-radius    |  Use EAP-RADIUS authentication               |
`,

			// Default:          stringdefault.StaticString(`eap-mschapv2`),
			Computed: true,
		},

		"server_mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Server authentication mode

    |  Format             |  Description                  |
    |---------------------|-------------------------------|
    |  pre-shared-secret  |  Use a pre-shared secret key  |
    |  x509               |  Use x.509 certificate        |
`,
			Description: `Server authentication mode

    |  Format             |  Description                  |
    |---------------------|-------------------------------|
    |  pre-shared-secret  |  Use a pre-shared secret key  |
    |  x509               |  Use x.509 certificate        |
`,

			// Default:          stringdefault.StaticString(`x509`),
			Computed: true,
		},

		"pre_shared_secret": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Pre-shared secret key

    |  Format  |  Description            |
    |----------|-------------------------|
    |  txt     |  Pre-shared secret key  |
`,
			Description: `Pre-shared secret key

    |  Format  |  Description            |
    |----------|-------------------------|
    |  txt     |  Pre-shared secret key  |
`,
		},

		// Nodes

		"x509": schema.SingleNestedAttribute{
			Attributes: VpnIPsecRemoteAccessConnectionAuthenticationXfivezeronine{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `X.509 certificate

`,
			Description: `X.509 certificate

`,
		},

		"local_users": schema.SingleNestedAttribute{
			Attributes: VpnIPsecRemoteAccessConnectionAuthenticationLocalUsers{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Local user authentication

`,
			Description: `Local user authentication

`,
		},
	}
}