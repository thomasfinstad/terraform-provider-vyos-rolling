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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (rule) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &LoadBalancingHaproxyServiceRule{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (rule) */
// LoadBalancingHaproxyServiceRuleSelfIdentifier used by TagNodes to keep the id info
type LoadBalancingHaproxyServiceRuleSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (rule) */

	LoadBalancingHaproxyServiceRule types.Number `tfsdk:"rule" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (service) */

	LoadBalancingHaproxyService types.String `tfsdk:"service" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (haproxy) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (load-balancing) */
}

// LoadBalancingHaproxyServiceRule describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type LoadBalancingHaproxyServiceRule struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *LoadBalancingHaproxyServiceRuleSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafLoadBalancingHaproxyServiceRuleDomainName types.List   `tfsdk:"domain_name" vyos:"domain-name,omitempty"`
	LeafLoadBalancingHaproxyServiceRuleSsl        types.String `tfsdk:"ssl" vyos:"ssl,omitempty"`

	// TagNodes

	// Nodes

	NodeLoadBalancingHaproxyServiceRuleSet *LoadBalancingHaproxyServiceRuleSet `tfsdk:"set" vyos:"set,omitempty"`

	NodeLoadBalancingHaproxyServiceRuleURLPath *LoadBalancingHaproxyServiceRuleURLPath `tfsdk:"url_path" vyos:"url-path,omitempty"`
}

// SetID configures the resource ID
func (o *LoadBalancingHaproxyServiceRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *LoadBalancingHaproxyServiceRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *LoadBalancingHaproxyServiceRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingHaproxyServiceRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"rule",
		o.SelfIdentifier.LoadBalancingHaproxyServiceRule.ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *LoadBalancingHaproxyServiceRule) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (haproxy) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (load-balancing) */
		"load-balancing", // Node

		"haproxy", // Node

		"service",
		o.SelfIdentifier.LoadBalancingHaproxyService.ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *LoadBalancingHaproxyServiceRule) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (service) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (haproxy) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (load-balancing) */
		"load-balancing", // Node

		"haproxy", // Node

		"service",
		o.SelfIdentifier.LoadBalancingHaproxyService.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingHaproxyServiceRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"rule": schema.NumberAttribute{
					Required: true,
					MarkdownDescription: `Proxy rule number

    |  Format   |  Description                 |
    |-----------|------------------------------|
    |  1-10000  |  Number for this proxy rule  |
`,
					Description: `Proxy rule number

    |  Format   |  Description                 |
    |-----------|------------------------------|
    |  1-10000  |  Number for this proxy rule  |
`,
					PlanModifiers: []planmodifier.Number{
						numberplanmodifier.RequiresReplace(),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (service) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (haproxy) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (load-balancing) */

				"service": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Frontend service name

`,
					Description: `Frontend service name

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in service, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  service, value must match: ^[.:a-zA-Z0-9-_/]+$",
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

		"domain_name":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (domain-name) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain name to match

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  Domain address to match  |
`,
			Description: `Domain name to match

    |  Format  |  Description              |
    |----------|---------------------------|
    |  txt     |  Domain address to match  |
`,
		},

		"ssl":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (ssl) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `SSL match options

    |  Format          |  Description                                                  |
    |------------------|---------------------------------------------------------------|
    |  req-ssl-sni     |  SSL Server Name Indication (SNI) request match               |
    |  ssl-fc-sni      |  SSL frontend connection Server Name Indication match         |
    |  ssl-fc-sni-end  |  SSL frontend match end of connection Server Name Indication  |
`,
			Description: `SSL match options

    |  Format          |  Description                                                  |
    |------------------|---------------------------------------------------------------|
    |  req-ssl-sni     |  SSL Server Name Indication (SNI) request match               |
    |  ssl-fc-sni      |  SSL frontend connection Server Name Indication match         |
    |  ssl-fc-sni-end  |  SSL frontend match end of connection Server Name Indication  |
`,
		},

		// TagNodes

		// Nodes

		"set": schema.SingleNestedAttribute{
			Attributes: LoadBalancingHaproxyServiceRuleSet{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Proxy modifications

`,
			Description: `Proxy modifications

`,
		},

		"url_path": schema.SingleNestedAttribute{
			Attributes: LoadBalancingHaproxyServiceRuleURLPath{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `URL path match

`,
			Description: `URL path match

`,
		},
	}
}
