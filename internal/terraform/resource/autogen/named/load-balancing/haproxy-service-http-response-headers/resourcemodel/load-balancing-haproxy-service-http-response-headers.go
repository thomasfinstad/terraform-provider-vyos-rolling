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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (http-response-headers) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &LoadBalancingHaproxyServiceHTTPResponseHeaders{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (http-response-headers) */
// LoadBalancingHaproxyServiceHTTPResponseHeadersSelfIdentifier used by TagNodes to keep the id info
type LoadBalancingHaproxyServiceHTTPResponseHeadersSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (http-response-headers) */

	LoadBalancingHaproxyServiceHTTPResponseHeaders types.String `tfsdk:"http_response_headers" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (service) */

	LoadBalancingHaproxyService types.String `tfsdk:"service" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (haproxy) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (load-balancing) */
}

// LoadBalancingHaproxyServiceHTTPResponseHeaders describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type LoadBalancingHaproxyServiceHTTPResponseHeaders struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *LoadBalancingHaproxyServiceHTTPResponseHeadersSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafLoadBalancingHaproxyServiceHTTPResponseHeadersValue types.String `tfsdk:"value" vyos:"value,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *LoadBalancingHaproxyServiceHTTPResponseHeaders) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *LoadBalancingHaproxyServiceHTTPResponseHeaders) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *LoadBalancingHaproxyServiceHTTPResponseHeaders) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingHaproxyServiceHTTPResponseHeaders) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"http-response-headers",
		o.SelfIdentifier.LoadBalancingHaproxyServiceHTTPResponseHeaders.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *LoadBalancingHaproxyServiceHTTPResponseHeaders) GetVyosParentPath() []string {
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
func (o *LoadBalancingHaproxyServiceHTTPResponseHeaders) GetVyosNamedParentPath() []string {
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
func (o LoadBalancingHaproxyServiceHTTPResponseHeaders) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"http_response_headers": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Headers to include in HTTP response

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  HTTP header name  |
`,
					Description: `Headers to include in HTTP response

    |  Format  |  Description       |
    |----------|--------------------|
    |  txt     |  HTTP header name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in http_response_headers, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  http_response_headers, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
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

		"value":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (value) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `HTTP header value

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  HTTP header value  |
`,
			Description: `HTTP header value

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  HTTP header value  |
`,
		},

		// TagNodes

		// Nodes

	}
}
