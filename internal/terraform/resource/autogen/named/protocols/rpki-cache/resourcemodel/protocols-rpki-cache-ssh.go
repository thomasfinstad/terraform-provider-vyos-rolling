// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProtocolsRpkiCacheTCP describes the resource data model.
type ProtocolsRpkiCacheTCP struct {
	// LeafNodes
	LeafProtocolsRpkiCacheTCPKnownHostsFile types.String `tfsdk:"known_hosts_file" vyos:"known-hosts-file,omitempty"`
	LeafProtocolsRpkiCacheTCPPrivateKeyFile types.String `tfsdk:"private_key_file" vyos:"private-key-file,omitempty"`
	LeafProtocolsRpkiCacheTCPPublicKeyFile  types.String `tfsdk:"public_key_file" vyos:"public-key-file,omitempty"`
	LeafProtocolsRpkiCacheTCPUsername       types.String `tfsdk:"username" vyos:"username,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRpkiCacheTCP) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"known_hosts_file": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `RPKI SSH known hosts file

`,
		},

		"private_key_file": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `RPKI SSH private key file

`,
		},

		"public_key_file": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `RPKI SSH public key file path

`,
		},

		"username": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Username used for authentication

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Username  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ProtocolsRpkiCacheTCP) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ProtocolsRpkiCacheTCP) UnmarshalJSON(_ []byte) error {
	return nil
}
