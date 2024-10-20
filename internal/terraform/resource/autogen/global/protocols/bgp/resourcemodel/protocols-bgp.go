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

/* tools/generate-terraform-resource-full/templates/resources/global/resource-model.gotmpl */
// Validate compliance
var _ helpers.VyosTopResourceDataModel = &ProtocolsBgp{}

// ProtocolsBgp describes the resource data model.
type ProtocolsBgp struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBgpSystemAs types.Number `tfsdk:"system_as" vyos:"system-as,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagProtocolsBgpInterface bool `tfsdk:"-" vyos:"interface,child"`
	ExistsTagProtocolsBgpNeighbor  bool `tfsdk:"-" vyos:"neighbor,child"`
	ExistsTagProtocolsBgpPeerGroup bool `tfsdk:"-" vyos:"peer-group,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsBgpAddressFamily bool `tfsdk:"-" vyos:"address-family,child"`
	ExistsNodeProtocolsBgpBmp           bool `tfsdk:"-" vyos:"bmp,child"`
	ExistsNodeProtocolsBgpListen        bool `tfsdk:"-" vyos:"listen,child"`
	ExistsNodeProtocolsBgpParameters    bool `tfsdk:"-" vyos:"parameters,child"`
	ExistsNodeProtocolsBgpSrvsix        bool `tfsdk:"-" vyos:"srv6,child"`
	ExistsNodeProtocolsBgpSID           bool `tfsdk:"-" vyos:"sid,child"`
	ExistsNodeProtocolsBgpTimers        bool `tfsdk:"-" vyos:"timers,child"`
}

// SetID configures the resource ID
func (o *ProtocolsBgp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBgp) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBgp) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgp) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"bgp",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBgp) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */
		"protocols",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsBgp) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgp) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"system_as":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Autonomous System Number (ASN)

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  1-4294967294  |  Autonomous System Number  |
`,
			Description: `Autonomous System Number (ASN)

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  1-4294967294  |  Autonomous System Number  |
`,
		},
	}
}
