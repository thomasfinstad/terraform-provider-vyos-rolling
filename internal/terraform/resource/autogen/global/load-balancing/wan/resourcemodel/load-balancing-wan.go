// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &LoadBalancingWan{}

// LoadBalancingWan describes the resource data model.
type LoadBalancingWan struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafLoadBalancingWanDisableSourceNat   types.Bool   `tfsdk:"disable_source_nat" vyos:"disable-source-nat,omitempty"`
	LeafLoadBalancingWanEnableLocalTraffic types.Bool   `tfsdk:"enable_local_traffic" vyos:"enable-local-traffic,omitempty"`
	LeafLoadBalancingWanFlushConnections   types.Bool   `tfsdk:"flush_connections" vyos:"flush-connections,omitempty"`
	LeafLoadBalancingWanHook               types.String `tfsdk:"hook" vyos:"hook,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagLoadBalancingWanInterfaceHealth bool `tfsdk:"-" vyos:"interface-health,child"`
	ExistsTagLoadBalancingWanRule            bool `tfsdk:"-" vyos:"rule,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeLoadBalancingWanStickyConnections bool `tfsdk:"-" vyos:"sticky-connections,child"`
}

// SetID configures the resource ID
func (o *LoadBalancingWan) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *LoadBalancingWan) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *LoadBalancingWan) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingWan) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"wan",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *LoadBalancingWan) GetVyosParentPath() []string {
	return []string{
		"load-balancing",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *LoadBalancingWan) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingWan) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"disable_source_nat": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable source NAT rules from being configured for WAN load balancing

`,
			Description: `Disable source NAT rules from being configured for WAN load balancing

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_local_traffic": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable WAN load balancing for locally sourced traffic

`,
			Description: `Enable WAN load balancing for locally sourced traffic

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"flush_connections": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Flush connection tracking tables on connection state change

`,
			Description: `Flush connection tracking tables on connection state change

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"hook": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Script to be executed on interface status change

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  txt     |  Script in /config/scripts  |
`,
			Description: `Script to be executed on interface status change

    |  Format  |  Description                |
    |----------|-----------------------------|
    |  txt     |  Script in /config/scripts  |
`,
		},
	}
}
