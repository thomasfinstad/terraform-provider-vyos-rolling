/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsRIPng{}

// ProtocolsRIPng describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ProtocolsRIPng struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsRIPngAggregateAddress types.List   `tfsdk:"aggregate_address" vyos:"aggregate-address,omitempty"`
	LeafProtocolsRIPngDefaultMetric    types.Number `tfsdk:"default_metric" vyos:"default-metric,omitempty"`
	LeafProtocolsRIPngNetwork          types.List   `tfsdk:"network" vyos:"network,omitempty"`
	LeafProtocolsRIPngPassiveInterface types.List   `tfsdk:"passive_interface" vyos:"passive-interface,omitempty"`
	LeafProtocolsRIPngRoute            types.List   `tfsdk:"route" vyos:"route,omitempty"`
	LeafProtocolsRIPngRouteMap         types.String `tfsdk:"route_map" vyos:"route-map,omitempty"`

	// TagNodes

	ExistsTagProtocolsRIPngInterface bool `tfsdk:"-" vyos:"interface,child"`

	// Nodes

	ExistsNodeProtocolsRIPngDefaultInformation bool `tfsdk:"-" vyos:"default-information,child"`

	NodeProtocolsRIPngDistributeList *ProtocolsRIPngDistributeList `tfsdk:"distribute_list" vyos:"distribute-list,omitempty"`

	NodeProtocolsRIPngRedistribute *ProtocolsRIPngRedistribute `tfsdk:"redistribute" vyos:"redistribute,omitempty"`

	ExistsNodeProtocolsRIPngTimers bool `tfsdk:"-" vyos:"timers,child"`
}

// SetID configures the resource ID
func (o *ProtocolsRIPng) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsRIPng) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsRIPng) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsRIPng) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"ripng",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsRIPng) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"protocols", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsRIPng) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsRIPng) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"aggregate_address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Aggregate RIPng route announcement

    |  Format   |  Description                         |
    |-----------|--------------------------------------|
    |  ipv6net  |  Aggregate RIPng route announcement  |
`,
			Description: `Aggregate RIPng route announcement

    |  Format   |  Description                         |
    |-----------|--------------------------------------|
    |  ipv6net  |  Aggregate RIPng route announcement  |
`,
		},

		"default_metric":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Metric of redistributed routes

    |  Format  |  Description     |
    |----------|------------------|
    |  1-16    |  Default metric  |
`,
			Description: `Metric of redistributed routes

    |  Format  |  Description     |
    |----------|------------------|
    |  1-16    |  Default metric  |
`,
		},

		"network":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `RIPng network

    |  Format   |  Description    |
    |-----------|-----------------|
    |  ipv6net  |  RIPng network  |
`,
			Description: `RIPng network

    |  Format   |  Description    |
    |-----------|-----------------|
    |  ipv6net  |  RIPng network  |
`,
		},

		"passive_interface":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Passive interface

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  txt     |  Suppress routing updates on interface  |
`,
			Description: `Passive interface

    |  Format  |  Description                            |
    |----------|-----------------------------------------|
    |  txt     |  Suppress routing updates on interface  |
`,
		},

		"route":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `RIPng static route

    |  Format   |  Description         |
    |-----------|----------------------|
    |  ipv6net  |  RIPng static route  |
`,
			Description: `RIPng static route

    |  Format   |  Description         |
    |-----------|----------------------|
    |  ipv6net  |  RIPng static route  |
`,
		},

		"route_map":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specify route-map name to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
			Description: `Specify route-map name to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Route map name  |
`,
		},

		// TagNodes

		// Nodes

		"distribute_list": schema.SingleNestedAttribute{
			Attributes: ProtocolsRIPngDistributeList{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Filter networks in routing updates

`,
			Description: `Filter networks in routing updates

`,
		},

		"redistribute": schema.SingleNestedAttribute{
			Attributes: ProtocolsRIPngRedistribute{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Redistribute information from another routing protocol

`,
			Description: `Redistribute information from another routing protocol

`,
		},
	}
}
