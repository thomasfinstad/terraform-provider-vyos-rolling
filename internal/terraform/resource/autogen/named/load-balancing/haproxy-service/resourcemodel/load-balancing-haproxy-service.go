// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

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

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &LoadBalancingHaproxyService{}

// LoadBalancingHaproxyService describes the resource data model.
type LoadBalancingHaproxyService struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafLoadBalancingHaproxyServiceBackend             types.List   `tfsdk:"backend" vyos:"backend,omitempty"`
	LeafLoadBalancingHaproxyServiceDescrIPtion         types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafLoadBalancingHaproxyServiceListenAddress       types.List   `tfsdk:"listen_address" vyos:"listen-address,omitempty"`
	LeafLoadBalancingHaproxyServiceMode                types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafLoadBalancingHaproxyServicePort                types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafLoadBalancingHaproxyServiceRedirectHTTPToHTTPS types.Bool   `tfsdk:"redirect_http_to_https" vyos:"redirect-http-to-https,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagLoadBalancingHaproxyServiceRule bool `tfsdk:"-" vyos:"rule,child"`

	ExistsTagLoadBalancingHaproxyServiceHTTPResponseHeaders bool `tfsdk:"-" vyos:"http-response-headers,child"`

	// Nodes
	NodeLoadBalancingHaproxyServiceLogging    *LoadBalancingHaproxyServiceLogging    `tfsdk:"logging" vyos:"logging,omitempty"`
	NodeLoadBalancingHaproxyServiceTCPRequest *LoadBalancingHaproxyServiceTCPRequest `tfsdk:"tcp_request" vyos:"tcp-request,omitempty"`
	NodeLoadBalancingHaproxyServiceSsl        *LoadBalancingHaproxyServiceSsl        `tfsdk:"ssl" vyos:"ssl,omitempty"`
}

// SetID configures the resource ID
func (o *LoadBalancingHaproxyService) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *LoadBalancingHaproxyService) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *LoadBalancingHaproxyService) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingHaproxyService) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"service",
		o.SelfIdentifier.Attributes()["service"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *LoadBalancingHaproxyService) GetVyosParentPath() []string {
	return []string{
		"load-balancing",

		"haproxy",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *LoadBalancingHaproxyService) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingHaproxyService) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  service, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"backend": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Backend member

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  txt     |  Name of haproxy backend system  |
`,
			Description: `Backend member

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  txt     |  Name of haproxy backend system  |
`,
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

		"listen_address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local IP addresses to listen on

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  ipv4    |  IPv4 address to listen for incoming connections  |
    |  ipv6    |  IPv6 address to listen for incoming connections  |
`,
			Description: `Local IP addresses to listen on

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  ipv4    |  IPv4 address to listen for incoming connections  |
    |  ipv6    |  IPv6 address to listen for incoming connections  |
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

		"port": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
			Description: `Port number used by connection

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Numeric IP port  |
`,
		},

		"redirect_http_to_https": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Redirect HTTP to HTTPS

`,
			Description: `Redirect HTTP to HTTPS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"logging": schema.SingleNestedAttribute{
			Attributes: LoadBalancingHaproxyServiceLogging{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Logging parameters

`,
			Description: `Logging parameters

`,
		},

		"tcp_request": schema.SingleNestedAttribute{
			Attributes: LoadBalancingHaproxyServiceTCPRequest{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `TCP request directive

`,
			Description: `TCP request directive

`,
		},

		"ssl": schema.SingleNestedAttribute{
			Attributes: LoadBalancingHaproxyServiceSsl{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `SSL Certificate, SSL Key and CA

`,
			Description: `SSL Certificate, SSL Key and CA

`,
		},
	}
}
