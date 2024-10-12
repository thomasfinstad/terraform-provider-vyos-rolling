// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesSstpcSsl{}

// InterfacesSstpcSsl describes the resource data model.
type InterfacesSstpcSsl struct {
	// LeafNodes
	LeafInterfacesSstpcSslCaCertificate types.String `tfsdk:"ca_certificate" vyos:"ca-certificate,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesSstpcSsl) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"ca_certificate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate Authority in PKI configuration

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  txt     |  Name of CA in PKI configuration  |
`,
			Description: `Certificate Authority in PKI configuration

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  txt     |  Name of CA in PKI configuration  |
`,
		},

		// Nodes

	}
}
