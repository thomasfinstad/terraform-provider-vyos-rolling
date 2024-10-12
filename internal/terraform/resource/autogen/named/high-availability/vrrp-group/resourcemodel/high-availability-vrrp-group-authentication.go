// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &HighAvailabilityVrrpGroupAuthentication{}

// HighAvailabilityVrrpGroupAuthentication describes the resource data model.
type HighAvailabilityVrrpGroupAuthentication struct {
	// LeafNodes
	LeafHighAvailabilityVrrpGroupAuthenticationPassword types.String `tfsdk:"password" vyos:"password,omitempty"`
	LeafHighAvailabilityVrrpGroupAuthenticationType     types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o HighAvailabilityVrrpGroupAuthentication) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"password": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRRP password

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  txt     |  Password string (up to 8 characters)  |
`,
			Description: `VRRP password

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  txt     |  Password string (up to 8 characters)  |
`,
		},

		"type": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Authentication type

    |  Format              |  Description                   |
    |----------------------|--------------------------------|
    |  plaintext-password  |  Simple password string        |
    |  ah                  |  AH - IPSEC (not recommended)  |
`,
			Description: `Authentication type

    |  Format              |  Description                   |
    |----------------------|--------------------------------|
    |  plaintext-password  |  Simple password string        |
    |  ah                  |  AH - IPSEC (not recommended)  |
`,
		},

		// Nodes

	}
}
