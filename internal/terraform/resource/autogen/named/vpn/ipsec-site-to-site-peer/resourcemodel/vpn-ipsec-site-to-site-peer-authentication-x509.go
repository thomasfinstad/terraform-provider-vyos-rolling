// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// VpnIPsecSiteToSitePeerAuthenticationXfivezeronine describes the resource data model.
type VpnIPsecSiteToSitePeerAuthenticationXfivezeronine struct {
	// LeafNodes
	LeafVpnIPsecSiteToSitePeerAuthenticationXfivezeronineCertificate   types.String `tfsdk:"certificate" vyos:"certificate,omitempty"`
	LeafVpnIPsecSiteToSitePeerAuthenticationXfivezeroninePassphrase    types.String `tfsdk:"passphrase" vyos:"passphrase,omitempty"`
	LeafVpnIPsecSiteToSitePeerAuthenticationXfivezeronineCaCertificate types.String `tfsdk:"ca_certificate" vyos:"ca-certificate,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnIPsecSiteToSitePeerAuthenticationXfivezeronine) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"certificate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate in PKI configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of certificate in PKI configuration  |

`,
		},

		"passphrase": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Private key passphrase

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Passphrase to decrypt the private key  |

`,
		},

		"ca_certificate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate Authority in PKI configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of CA in PKI configuration  |

`,
		},

		// Nodes

	}
}
