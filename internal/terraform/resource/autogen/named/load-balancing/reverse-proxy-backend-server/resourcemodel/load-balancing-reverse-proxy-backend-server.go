// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
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

var _ helpers.VyosTopResourceDataModel = &LoadBalancingReverseProxyBackendServer{}

// LoadBalancingReverseProxyBackendServer describes the resource data model.
type LoadBalancingReverseProxyBackendServer struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafLoadBalancingReverseProxyBackendServerAddress       types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafLoadBalancingReverseProxyBackendServerBackup        types.Bool   `tfsdk:"backup" vyos:"backup,omitempty"`
	LeafLoadBalancingReverseProxyBackendServerCheck         types.Bool   `tfsdk:"check" vyos:"check,omitempty"`
	LeafLoadBalancingReverseProxyBackendServerPort          types.Number `tfsdk:"port" vyos:"port,omitempty"`
	LeafLoadBalancingReverseProxyBackendServerSendProxy     types.Bool   `tfsdk:"send_proxy" vyos:"send-proxy,omitempty"`
	LeafLoadBalancingReverseProxyBackendServerSendProxyVtwo types.Bool   `tfsdk:"send_proxy_v2" vyos:"send-proxy-v2,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *LoadBalancingReverseProxyBackendServer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *LoadBalancingReverseProxyBackendServer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *LoadBalancingReverseProxyBackendServer) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingReverseProxyBackendServer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"server",
		o.SelfIdentifier.Attributes()["server"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *LoadBalancingReverseProxyBackendServer) GetVyosParentPath() []string {
	return []string{
		"load-balancing",

		"reverse-proxy",

		"backend",

		o.SelfIdentifier.Attributes()["backend"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *LoadBalancingReverseProxyBackendServer) GetVyosNamedParentPath() []string {
	return []string{
		"load-balancing",

		"reverse-proxy",

		"backend",

		o.SelfIdentifier.Attributes()["backend"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingReverseProxyBackendServer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"server": schema.StringAttribute{
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
										"double underscores in server, conflicts with the internal resource id",
									),
								),
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
									"illegal character in  server, value must match: ^[a-zA-Z0-9-_]*$",
								),
							),
						},
					},

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
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"address": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Backend server address

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  ipv4    |  IPv4 unicast peer address  |
    |  ipv6    |  IPv6 unicast peer address  |
`,
			Description: `Backend server address

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  ipv4    |  IPv4 unicast peer address  |
    |  ipv6    |  IPv6 unicast peer address  |
`,
		},

		"backup": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Use backup server if other servers are not available

`,
			Description: `Use backup server if other servers are not available

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"check": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Active health check backend server

`,
			Description: `Active health check backend server

`,
			Default:  booldefault.StaticBool(false),
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

		"send_proxy": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send a Proxy Protocol version 1 header (text format)

`,
			Description: `Send a Proxy Protocol version 1 header (text format)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"send_proxy_v2": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Send a Proxy Protocol version 2 header (binary format)

`,
			Description: `Send a Proxy Protocol version 2 header (binary format)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

	}
}
