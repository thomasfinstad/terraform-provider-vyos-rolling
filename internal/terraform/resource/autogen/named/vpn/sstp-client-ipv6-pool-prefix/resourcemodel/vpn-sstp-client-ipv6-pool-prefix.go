/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (prefix) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &VpnSstpClientIPvsixPoolPrefix{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (prefix) */
// VpnSstpClientIPvsixPoolPrefixSelfIdentifier used by TagNodes to keep the id info
type VpnSstpClientIPvsixPoolPrefixSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (prefix) */

	VpnSstpClientIPvsixPoolPrefix types.String `tfsdk:"prefix" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (client-ipv6-pool) */

	VpnSstpClientIPvsixPool types.String `tfsdk:"client_ipv6_pool" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (sstp) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (vpn) */
}

// VpnSstpClientIPvsixPoolPrefix describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type VpnSstpClientIPvsixPoolPrefix struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *VpnSstpClientIPvsixPoolPrefixSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafVpnSstpClientIPvsixPoolPrefixMask types.Number `tfsdk:"mask" vyos:"mask,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *VpnSstpClientIPvsixPoolPrefix) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *VpnSstpClientIPvsixPoolPrefix) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *VpnSstpClientIPvsixPoolPrefix) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *VpnSstpClientIPvsixPoolPrefix) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"prefix",
		o.SelfIdentifier.VpnSstpClientIPvsixPoolPrefix.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *VpnSstpClientIPvsixPoolPrefix) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (client-ipv6-pool) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (sstp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vpn) */
		"vpn", // Node

		"sstp", // Node

		"client-ipv6-pool",
		o.SelfIdentifier.VpnSstpClientIPvsixPool.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *VpnSstpClientIPvsixPoolPrefix) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (client-ipv6-pool) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (client-ipv6-pool) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (sstp) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (vpn) */
		"vpn", // Node

		"sstp", // Node

		"client-ipv6-pool",
		o.SelfIdentifier.VpnSstpClientIPvsixPool.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VpnSstpClientIPvsixPoolPrefix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"prefix": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Pool of addresses used to assign to clients

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
					Description: `Pool of addresses used to assign to clients

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  |  IPv6 address and prefix length  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in prefix, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  prefix, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (client-ipv6-pool) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (sstp) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (vpn) */

				"client_ipv6_pool": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Pool of client IPv6 addresses

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  Name of IPv6 pool  |
`,
					Description: `Pool of client IPv6 addresses

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  Name of IPv6 pool  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in client_ipv6_pool, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  client_ipv6_pool, value must match: ^[.:a-zA-Z0-9-_/]+$",
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

		"mask":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mask) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length used for individual client

    |  Format  |  Description           |
    |----------|------------------------|
    |  48-128  |  Client prefix length  |
`,
			Description: `Prefix length used for individual client

    |  Format  |  Description           |
    |----------|------------------------|
    |  48-128  |  Client prefix length  |
`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
