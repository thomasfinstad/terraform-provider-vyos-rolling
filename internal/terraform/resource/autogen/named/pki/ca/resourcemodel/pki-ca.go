// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// PkiCa describes the resource data model.
type PkiCa struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"ca_id" vyos:"-,self-id"`

	// LeafNodes
	LeafPkiCaCertificate types.String `tfsdk:"certificate" vyos:"certificate,omitempty"`
	LeafPkiCaDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPkiCaCrl         types.List   `tfsdk:"crl" vyos:"crl,omitempty"`
	LeafPkiCaRevoke      types.Bool   `tfsdk:"revoke" vyos:"revoke,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePkiCaPrivate *PkiCaPrivate `tfsdk:"private" vyos:"private,omitempty"`
}

// SetID configures the resource ID
func (o *PkiCa) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PkiCa) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"pki",

		"ca",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiCa) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"ca_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Certificate Authority

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
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

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Description  |

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
			MarkdownDescription: `Include certificate in parent CRL

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
