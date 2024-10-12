// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VpnIPsecSiteToSitePeerAuthenticationRsa{}

// VpnIPsecSiteToSitePeerAuthenticationRsa describes the resource data model.
type VpnIPsecSiteToSitePeerAuthenticationRsa struct {
	// LeafNodes
	LeafVpnIPsecSiteToSitePeerAuthenticationRsaLocalKey   types.String `tfsdk:"local_key" vyos:"local-key,omitempty"`
	LeafVpnIPsecSiteToSitePeerAuthenticationRsaPassphrase types.String `tfsdk:"passphrase" vyos:"passphrase,omitempty"`
	LeafVpnIPsecSiteToSitePeerAuthenticationRsaRemoteKey  types.String `tfsdk:"remote_key" vyos:"remote-key,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecSiteToSitePeerAuthenticationRsa) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"local_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Name of PKI key-pair with local private key

`,
			Description: `Name of PKI key-pair with local private key

`,
		},

		"passphrase": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Local private key passphrase

`,
			Description: `Local private key passphrase

`,
		},

		"remote_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Name of PKI key-pair with remote public key

`,
			Description: `Name of PKI key-pair with remote public key

`,
		},

		// Nodes

	}
}
