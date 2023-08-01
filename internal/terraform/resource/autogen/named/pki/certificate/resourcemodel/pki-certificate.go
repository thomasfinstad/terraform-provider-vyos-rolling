// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PkiCertificate describes the resource data model.
type PkiCertificate struct {
	SelfIdentifier types.String `tfsdk:"certificate_id" vyos:",self-id"`

	// LeafNodes
	LeafPkiCertificateCertificate types.String `tfsdk:"certificate" vyos:"certificate,omitempty"`
	LeafPkiCertificateDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPkiCertificateRevoke      types.Bool   `tfsdk:"revoke" vyos:"revoke,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePkiCertificatePrivate *PkiCertificatePrivate `tfsdk:"private" vyos:"private,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PkiCertificate) GetVyosPath() []string {
	return []string{
		"pki",

		"certificate",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiCertificate) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"certificate_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Certificate

`,
		},

		// LeafNodes

		"certificate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate in PEM format

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

`,
		},

		"revoke": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `If CA is present, this certificate will be included in generated CRLs

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"private": schema.SingleNestedAttribute{
			Attributes: PkiCertificatePrivate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Certificate private key

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PkiCertificate) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PkiCertificate) UnmarshalJSON(_ []byte) error {
	return nil
}
