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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceIPoeServerAuthenticationInterfaceMac{}

// ServiceIPoeServerAuthenticationInterfaceMac describes the resource data model.
type ServiceIPoeServerAuthenticationInterfaceMac struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl */
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceIPoeServerAuthenticationInterfaceMacVlan types.Number `tfsdk:"vlan" vyos:"vlan,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeServiceIPoeServerAuthenticationInterfaceMacRateLimit *ServiceIPoeServerAuthenticationInterfaceMacRateLimit `tfsdk:"rate_limit" vyos:"rate-limit,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceIPoeServerAuthenticationInterfaceMac) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceIPoeServerAuthenticationInterfaceMac) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceIPoeServerAuthenticationInterfaceMac) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerAuthenticationInterfaceMac) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"mac",
		o.SelfIdentifier.Attributes()["mac"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceIPoeServerAuthenticationInterfaceMac) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */
		"service",

		"ipoe-server",

		"authentication",

		"interface",

		o.SelfIdentifier.Attributes()["interface"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceIPoeServerAuthenticationInterfaceMac) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */
		"service",

		"ipoe-server",

		"authentication",

		"interface",

		o.SelfIdentifier.Attributes()["interface"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerAuthenticationInterfaceMac) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"mac": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
					Description: `Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in mac, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  mac, value must match: ^[.:a-zA-Z0-9-_]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				"interface": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Network interface for client MAC addresses

`,
					Description: `Network interface for client MAC addresses

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in interface, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  interface, value must match: ^[.:a-zA-Z0-9-_]+$",
							),
						),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"vlan":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `VLAN monitor for automatic creation of VLAN interfaces

    |  Format  |  Description     |
    |----------|------------------|
    |  1-4094  |  Client VLAN id  |
`,
			Description: `VLAN monitor for automatic creation of VLAN interfaces

    |  Format  |  Description     |
    |----------|------------------|
    |  1-4094  |  Client VLAN id  |
`,
		},

		// Nodes

		"rate_limit": schema.SingleNestedAttribute{
			Attributes: ServiceIPoeServerAuthenticationInterfaceMacRateLimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Upload/Download speed limits

`,
			Description: `Upload/Download speed limits

`,
		},
	}
}
