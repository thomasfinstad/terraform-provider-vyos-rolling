// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &PkiCertificate{}

// PkiCertificate describes the resource data model.
type PkiCertificate struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"certificate_id" vyos:"-,self-id"`

	// LeafNodes
	LeafPkiCertificateCertificate types.String `tfsdk:"certificate" vyos:"certificate,omitempty"`
	LeafPkiCertificateDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPkiCertificateRevoke      types.Bool   `tfsdk:"revoke" vyos:"revoke,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodePkiCertificateAcme    *PkiCertificateAcme    `tfsdk:"acme" vyos:"acme,omitempty"`
	NodePkiCertificatePrivate *PkiCertificatePrivate `tfsdk:"private" vyos:"private,omitempty"`
}

// SetID configures the resource ID
func (o *PkiCertificate) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PkiCertificate) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"pki",

		"certificate",
		o.SelfIdentifier.ValueString(),
	}
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PkiCertificate) GetVyosParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiCertificate) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"certificate_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Certificate

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in certificate_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  certificate_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
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

		"revoke": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Include certificate in parent CRL

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"acme": schema.SingleNestedAttribute{
			Attributes: PkiCertificateAcme{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Automatic Certificate Management Environment (ACME) request

`,
		},

		"private": schema.SingleNestedAttribute{
			Attributes: PkiCertificatePrivate{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Certificate private key

`,
		},
	}
}
