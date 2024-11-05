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

var _ helpers.VyosTopResourceDataModel = &ProtocolsBgpAddressFamilyIPvfourMulticastDistance{}

// ProtocolsBgpAddressFamilyIPvfourMulticastDistance describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ProtocolsBgpAddressFamilyIPvfourMulticastDistance struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBgpAddressFamilyIPvfourMulticastDistanceExternal types.Number `tfsdk:"external" vyos:"external,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvfourMulticastDistanceInternal types.Number `tfsdk:"internal" vyos:"internal,omitempty"`
	LeafProtocolsBgpAddressFamilyIPvfourMulticastDistanceLocal    types.Number `tfsdk:"local" vyos:"local,omitempty"`

	// TagNodes

	ExistsTagProtocolsBgpAddressFamilyIPvfourMulticastDistancePrefix bool `tfsdk:"-" vyos:"prefix,child"`

	// Nodes
}

// SetID configures the resource ID
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastDistance) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastDistance) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastDistance) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastDistance) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"distance",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastDistance) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"protocols", // Node

		"bgp", // Node

		"address-family", // Node

		"ipv4-multicast", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsBgpAddressFamilyIPvfourMulticastDistance) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpAddressFamilyIPvfourMulticastDistance) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"external":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `eBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  eBGP routes administrative distance  |
`,
			Description: `eBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  eBGP routes administrative distance  |
`,
		},

		"internal":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `iBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  iBGP routes administrative distance  |
`,
			Description: `iBGP routes administrative distance

    |  Format  |  Description                          |
    |----------|---------------------------------------|
    |  1-255   |  iBGP routes administrative distance  |
`,
		},

		"local":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Locally originated BGP routes administrative distance

    |  Format  |  Description                                            |
    |----------|---------------------------------------------------------|
    |  1-255   |  Locally originated BGP routes administrative distance  |
`,
			Description: `Locally originated BGP routes administrative distance

    |  Format  |  Description                                            |
    |----------|---------------------------------------------------------|
    |  1-255   |  Locally originated BGP routes administrative distance  |
`,
		},

		// TagNodes

		// Nodes

	}
}
