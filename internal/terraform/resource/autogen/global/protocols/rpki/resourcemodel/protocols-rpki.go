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
var _ helpers.VyosTopResourceDataModel = &ProtocolsRpki{}

// ProtocolsRpki describes the resource data model.
type ProtocolsRpki struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsRpkiExpireInterval types.Number `tfsdk:"expire_interval" vyos:"expire-interval,omitempty"`
	LeafProtocolsRpkiPollingPeriod  types.Number `tfsdk:"polling_period" vyos:"polling-period,omitempty"`
	LeafProtocolsRpkiRetryInterval  types.Number `tfsdk:"retry_interval" vyos:"retry-interval,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsRpkiCache bool `tfsdk:"-" vyos:"cache,child"`

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsRpki) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsRpki) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsRpki) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsRpki) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"rpki",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsRpki) GetVyosParentPath() []string {
	return []string{
		"protocols",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsRpki) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRpki) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"expire_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval to wait before expiring the cache

    |  Format      |  Description          |
    |--------------|-----------------------|
    |  600-172800  |  Interval in seconds  |
`,
			Description: `Interval to wait before expiring the cache

    |  Format      |  Description          |
    |--------------|-----------------------|
    |  600-172800  |  Interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`7200`),
			Computed: true,
		},

		"polling_period": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Cache polling interval

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  1-86400  |  Interval in seconds  |
`,
			Description: `Cache polling interval

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  1-86400  |  Interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`300`),
			Computed: true,
		},

		"retry_interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Retry interval to connect to the cache server

    |  Format  |  Description          |
    |----------|-----------------------|
    |  1-7200  |  Interval in seconds  |
`,
			Description: `Retry interval to connect to the cache server

    |  Format  |  Description          |
    |----------|-----------------------|
    |  1-7200  |  Interval in seconds  |
`,

			// Default:          stringdefault.StaticString(`600`),
			Computed: true,
		},
	}
}
