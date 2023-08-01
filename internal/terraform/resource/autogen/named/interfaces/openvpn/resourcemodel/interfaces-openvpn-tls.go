// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// InterfacesOpenvpnTLS describes the resource data model.
type InterfacesOpenvpnTLS struct {
	// LeafNodes
	LeafInterfacesOpenvpnTLSAuthKey       types.String `tfsdk:"auth_key" vyos:"auth-key,omitempty"`
	LeafInterfacesOpenvpnTLSCertificate   types.String `tfsdk:"certificate" vyos:"certificate,omitempty"`
	LeafInterfacesOpenvpnTLSCaCertificate types.List   `tfsdk:"ca_certificate" vyos:"ca-certificate,omitempty"`
	LeafInterfacesOpenvpnTLSDhParams      types.String `tfsdk:"dh_params" vyos:"dh-params,omitempty"`
	LeafInterfacesOpenvpnTLSCryptKey      types.String `tfsdk:"crypt_key" vyos:"crypt-key,omitempty"`
	LeafInterfacesOpenvpnTLSTLSVersionMin types.String `tfsdk:"tls_version_min" vyos:"tls-version-min,omitempty"`
	LeafInterfacesOpenvpnTLSRole          types.String `tfsdk:"role" vyos:"role,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesOpenvpnTLS) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"auth_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `TLS shared secret key for tls-auth

`,
		},

		"certificate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate in PKI configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of certificate in PKI configuration  |

`,
		},

		"ca_certificate": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Certificate Authority chain in PKI configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of CA in PKI configuration  |

`,
		},

		"dh_params": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Diffie Hellman parameters (server only)

`,
		},

		"crypt_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Static key to use to authenticate control channel

`,
		},

		"tls_version_min": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify the minimum required TLS version

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  1.0  &emsp; |  TLS v1.0  |
    |  1.1  &emsp; |  TLS v1.1  |
    |  1.2  &emsp; |  TLS v1.2  |
    |  1.3  &emsp; |  TLS v1.3  |

`,
		},

		"role": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `TLS negotiation role

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  active  &emsp; |  Initiate TLS negotiation actively  |
    |  passive  &emsp; |  Wait for incoming TLS connection  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *InterfacesOpenvpnTLS) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *InterfacesOpenvpnTLS) UnmarshalJSON(_ []byte) error {
	return nil
}
