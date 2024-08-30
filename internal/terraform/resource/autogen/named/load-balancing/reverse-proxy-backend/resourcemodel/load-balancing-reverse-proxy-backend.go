// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

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

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &LoadBalancingReverseProxyBackend{}

// LoadBalancingReverseProxyBackend describes the resource data model.
type LoadBalancingReverseProxyBackend struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafLoadBalancingReverseProxyBackendBalance     types.String `tfsdk:"balance" vyos:"balance,omitempty"`
	LeafLoadBalancingReverseProxyBackendDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafLoadBalancingReverseProxyBackendMode        types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafLoadBalancingReverseProxyBackendHealthCheck types.String `tfsdk:"health_check" vyos:"health-check,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagLoadBalancingReverseProxyBackendHTTPResponseHeaders bool `tfsdk:"-" vyos:"http-response-headers,child"`

	ExistsTagLoadBalancingReverseProxyBackendRule bool `tfsdk:"-" vyos:"rule,child"`

	ExistsTagLoadBalancingReverseProxyBackendServer bool `tfsdk:"-" vyos:"server,child"`

	// Nodes
	NodeLoadBalancingReverseProxyBackendLogging   *LoadBalancingReverseProxyBackendLogging   `tfsdk:"logging" vyos:"logging,omitempty"`
	NodeLoadBalancingReverseProxyBackendHTTPCheck *LoadBalancingReverseProxyBackendHTTPCheck `tfsdk:"http_check" vyos:"http-check,omitempty"`
	NodeLoadBalancingReverseProxyBackendSsl       *LoadBalancingReverseProxyBackendSsl       `tfsdk:"ssl" vyos:"ssl,omitempty"`
	NodeLoadBalancingReverseProxyBackendTimeout   *LoadBalancingReverseProxyBackendTimeout   `tfsdk:"timeout" vyos:"timeout,omitempty"`
}

// SetID configures the resource ID
func (o *LoadBalancingReverseProxyBackend) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *LoadBalancingReverseProxyBackend) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *LoadBalancingReverseProxyBackend) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingReverseProxyBackend) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"backend",
		o.SelfIdentifier.Attributes()["backend"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *LoadBalancingReverseProxyBackend) GetVyosParentPath() []string {
	return []string{
		"load-balancing",

		"reverse-proxy",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *LoadBalancingReverseProxyBackend) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingReverseProxyBackend) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"backend": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Backend server name

`,
					Description: `Backend server name

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in backend, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  backend, value must match: ^[a-zA-Z0-9-_]*$",
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

		"balance": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Load-balancing algorithm

    |  Format            |  Description                         |
    |--------------------|--------------------------------------|
    |  source-address    |  Based on hash of source IP address  |
    |  round-robin       |  Round robin                         |
    |  least-connection  |  Least connection                    |
`,
			Description: `Load-balancing algorithm

    |  Format            |  Description                         |
    |--------------------|--------------------------------------|
    |  source-address    |  Based on hash of source IP address  |
    |  round-robin       |  Round robin                         |
    |  least-connection  |  Least connection                    |
`,

			// Default:          stringdefault.StaticString(`round-robin`),
			Computed: true,
		},

		"description": schema.StringAttribute{
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

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Proxy mode

    |  Format  |  Description      |
    |----------|-------------------|
    |  http    |  HTTP proxy mode  |
    |  tcp     |  TCP proxy mode   |
`,
			Description: `Proxy mode

    |  Format  |  Description      |
    |----------|-------------------|
    |  http    |  HTTP proxy mode  |
    |  tcp     |  TCP proxy mode   |
`,

			// Default:          stringdefault.StaticString(`http`),
			Computed: true,
		},

		"health_check": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Non HTTP health check options

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  ldap    |  LDAP protocol check        |
    |  mysql   |  MySQL protocol check       |
    |  pgsql   |  PostgreSQL protocol check  |
    |  redis   |  Redis protocol check       |
    |  smtp    |  SMTP protocol check        |
`,
			Description: `Non HTTP health check options

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  ldap    |  LDAP protocol check        |
    |  mysql   |  MySQL protocol check       |
    |  pgsql   |  PostgreSQL protocol check  |
    |  redis   |  Redis protocol check       |
    |  smtp    |  SMTP protocol check        |
`,
		},

		// Nodes

		"logging": schema.SingleNestedAttribute{
			Attributes: LoadBalancingReverseProxyBackendLogging{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Logging parameters

`,
			Description: `Logging parameters

`,
		},

		"http_check": schema.SingleNestedAttribute{
			Attributes: LoadBalancingReverseProxyBackendHTTPCheck{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `HTTP check configuration

`,
			Description: `HTTP check configuration

`,
		},

		"ssl": schema.SingleNestedAttribute{
			Attributes: LoadBalancingReverseProxyBackendSsl{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `SSL Certificate, SSL Key and CA

`,
			Description: `SSL Certificate, SSL Key and CA

`,
		},

		"timeout": schema.SingleNestedAttribute{
			Attributes: LoadBalancingReverseProxyBackendTimeout{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Timeout options

`,
			Description: `Timeout options

`,
		},
	}
}
