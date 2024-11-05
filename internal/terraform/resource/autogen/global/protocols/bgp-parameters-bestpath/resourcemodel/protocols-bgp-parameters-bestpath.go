/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsBgpParametersBestpath{}

// ProtocolsBgpParametersBestpath describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ProtocolsBgpParametersBestpath struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsBgpParametersBestpathBandwIDth       types.String `tfsdk:"bandwidth" vyos:"bandwidth,omitempty"`
	LeafProtocolsBgpParametersBestpathCompareRouterID types.Bool   `tfsdk:"compare_routerid" vyos:"compare-routerid,omitempty"`
	LeafProtocolsBgpParametersBestpathMed             types.List   `tfsdk:"med" vyos:"med,omitempty"`

	// TagNodes

	// Nodes

	ExistsNodeProtocolsBgpParametersBestpathAsPath bool `tfsdk:"-" vyos:"as-path,child"`

	ExistsNodeProtocolsBgpParametersBestpathPeerType bool `tfsdk:"-" vyos:"peer-type,child"`
}

// SetID configures the resource ID
func (o *ProtocolsBgpParametersBestpath) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsBgpParametersBestpath) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsBgpParametersBestpath) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsBgpParametersBestpath) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"bestpath",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsBgpParametersBestpath) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"protocols", // Node

		"bgp", // Node

		"parameters", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsBgpParametersBestpath) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsBgpParametersBestpath) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"bandwidth":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Link Bandwidth attribute

    |  Format                      |  Description                                                            |
    |------------------------------|-------------------------------------------------------------------------|
    |  default-weight-for-missing  |  Assign low default weight (1) to paths not having link bandwidth       |
    |  ignore                      |  Ignore link bandwidth (do regular ECMP, not weighted)                  |
    |  skip-missing                |  Ignore paths without link bandwidth for ECMP (if other paths have it)  |
`,
			Description: `Link Bandwidth attribute

    |  Format                      |  Description                                                            |
    |------------------------------|-------------------------------------------------------------------------|
    |  default-weight-for-missing  |  Assign low default weight (1) to paths not having link bandwidth       |
    |  ignore                      |  Ignore link bandwidth (do regular ECMP, not weighted)                  |
    |  skip-missing                |  Ignore paths without link bandwidth for ECMP (if other paths have it)  |
`,
		},

		"compare_routerid":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Compare the router-id for identical EBGP paths

`,
			Description: `Compare the router-id for identical EBGP paths

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"med":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `MED attribute comparison parameters

    |  Format            |  Description                                              |
    |--------------------|-----------------------------------------------------------|
    |  confed            |  Compare MEDs among confederation paths                   |
    |  missing-as-worst  |  Treat missing route as a MED as the least preferred one  |
`,
			Description: `MED attribute comparison parameters

    |  Format            |  Description                                              |
    |--------------------|-----------------------------------------------------------|
    |  confed            |  Compare MEDs among confederation paths                   |
    |  missing-as-worst  |  Treat missing route as a MED as the least preferred one  |
`,
		},

		// TagNodes

		// Nodes

	}
}
