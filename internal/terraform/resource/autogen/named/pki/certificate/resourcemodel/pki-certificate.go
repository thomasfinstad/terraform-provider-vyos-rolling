/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &PkiCertificate{}

// PkiCertificate describes the resource data model.
type PkiCertificate struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl */
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafPkiCertificateCertificate types.String `tfsdk:"certificate" vyos:"certificate,omitempty"`
	LeafPkiCertificateDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafPkiCertificateRevoke      types.Bool   `tfsdk:"revoke" vyos:"revoke,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodePkiCertificateAcme    *PkiCertificateAcme    `tfsdk:"acme" vyos:"acme,omitempty"`
	NodePkiCertificatePrivate *PkiCertificatePrivate `tfsdk:"private" vyos:"private,omitempty"`
}

// SetID configures the resource ID
func (o *PkiCertificate) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *PkiCertificate) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PkiCertificate) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PkiCertificate) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"certificate",
		o.SelfIdentifier.Attributes()["certificate"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *PkiCertificate) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */
		"pki",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PkiCertificate) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PkiCertificate) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"certificate": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Certificate

`,
					Description: `Certificate

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in certificate, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  certificate, value must match: ^[.:a-zA-Z0-9-_]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"certificate":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate in PEM format

`,
			Description: `Certificate in PEM format

`,
		},

		"description":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"revoke":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Include certificate in parent CRL

`,
			Description: `Include certificate in parent CRL

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"acme": schema.SingleNestedAttribute{
			Attributes: PkiCertificateAcme{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Automatic Certificate Management Environment (ACME) request

`,
			Description: `Automatic Certificate Management Environment (ACME) request

`,
		},

		"private": schema.SingleNestedAttribute{
			Attributes: PkiCertificatePrivate{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Certificate private key

`,
			Description: `Certificate private key

`,
		},
	}
}
