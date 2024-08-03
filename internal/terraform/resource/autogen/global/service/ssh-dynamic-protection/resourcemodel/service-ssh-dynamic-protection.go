// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &ServiceTCPDynamicProtection{}

// ServiceTCPDynamicProtection describes the resource data model.
type ServiceTCPDynamicProtection struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceTCPDynamicProtectionBlockTime  types.Number `tfsdk:"block_time" vyos:"block-time,omitempty"`
	LeafServiceTCPDynamicProtectionDetectTime types.Number `tfsdk:"detect_time" vyos:"detect-time,omitempty"`
	LeafServiceTCPDynamicProtectionThreshold  types.Number `tfsdk:"threshold" vyos:"threshold,omitempty"`
	LeafServiceTCPDynamicProtectionAllowFrom  types.List   `tfsdk:"allow_from" vyos:"allow-from,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ServiceTCPDynamicProtection) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceTCPDynamicProtection) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceTCPDynamicProtection) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceTCPDynamicProtection) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"dynamic-protection",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceTCPDynamicProtection) GetVyosParentPath() []string {
	return []string{
		"service",

		"ssh",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ServiceTCPDynamicProtection) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceTCPDynamicProtection) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"block_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Block source IP in seconds. Subsequent blocks increase by a factor of 1.5

    |  Format   |  Description                            |
    |-----------|-----------------------------------------|
    |  1-65535  |  Time interval in seconds for blocking  |
`,
			Description: `Block source IP in seconds. Subsequent blocks increase by a factor of 1.5

    |  Format   |  Description                            |
    |-----------|-----------------------------------------|
    |  1-65535  |  Time interval in seconds for blocking  |
`,

			// Default:          stringdefault.StaticString(`120`),
			Computed: true,
		},

		"detect_time": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Remember source IP in seconds before reset their score

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  1-65535  |  Time interval in seconds  |
`,
			Description: `Remember source IP in seconds before reset their score

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  1-65535  |  Time interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`1800`),
			Computed: true,
		},

		"threshold": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Block source IP when their cumulative attack score exceeds threshold

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Threshold score  |
`,
			Description: `Block source IP when their cumulative attack score exceeds threshold

    |  Format   |  Description      |
    |-----------|-------------------|
    |  1-65535  |  Threshold score  |
`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"allow_from": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Always allow inbound connections from these systems

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4     |  Address to match against        |
    |  ipv4net  |  IPv4 address and prefix length  |
    |  ipv6     |  IPv6 address to match against   |
    |  ipv6net  |  IPv6 address and prefix length  |
`,
			Description: `Always allow inbound connections from these systems

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4     |  Address to match against        |
    |  ipv4net  |  IPv4 address and prefix length  |
    |  ipv6     |  IPv6 address to match against   |
    |  ipv6net  |  IPv6 address and prefix length  |
`,
		},
	}
}
