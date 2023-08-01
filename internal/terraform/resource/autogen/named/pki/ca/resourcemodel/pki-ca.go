// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PkiCa describes the resource data model.
type PkiCa struct {
	SelfIdentifier types.String `tfsdk:"ca_id" vyos:",self-id"`

	// LeafNodes
	LeafPkiCaCertificate types.String `tfsdk:"certificate" vyos:"certificate,omitempty"`
	LeafPkiCaDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPkiCaCrl         types.List   `tfsdk:"crl" vyos:"crl,omitempty"`
	LeafPkiCaRevoke      types.Bool   `tfsdk:"revoke" vyos:"revoke,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePkiCaPrivate *PkiCaPrivate `tfsdk:"private" vyos:"private,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PkiCa) GetVyosPath() []string {
	return []string{
		"pki",

		"ca",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiCa) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"ca_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Certificate Authority

`,
		},

		// LeafNodes

		"certificate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `CA certificate in PEM format

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

`,
		},

		"crl": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Certificate revocation list in PEM format

`,
		},

		"revoke": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `If parent CA is present, this CA certificate will be included in generated CRLs

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"private": schema.SingleNestedAttribute{
			Attributes: PkiCaPrivate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `CA private key in PEM format

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *PkiCa) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *PkiCa) UnmarshalJSON(_ []byte) error {
	return nil
}
