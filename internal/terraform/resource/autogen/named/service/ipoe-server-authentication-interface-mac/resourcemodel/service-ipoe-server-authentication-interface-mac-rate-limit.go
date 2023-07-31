// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceIPoeServerAuthenticationInterfaceMacRateLimit describes the resource data model.
type ServiceIPoeServerAuthenticationInterfaceMacRateLimit struct {
	// LeafNodes
	LeafServiceIPoeServerAuthenticationInterfaceMacRateLimitUpload   types.String `tfsdk:"upload" vyos:"upload,omitempty"`
	LeafServiceIPoeServerAuthenticationInterfaceMacRateLimitDownload types.String `tfsdk:"download" vyos:"download,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerAuthenticationInterfaceMacRateLimit) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"upload": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Upload bandwidth limit in kbits/sec

`,
		},

		"download": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Download bandwidth limit in kbits/sec

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceIPoeServerAuthenticationInterfaceMacRateLimit) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceIPoeServerAuthenticationInterfaceMacRateLimit) UnmarshalJSON(_ []byte) error {
	return nil
}
