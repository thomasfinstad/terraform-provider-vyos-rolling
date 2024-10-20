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

/* tools/generate-terraform-resource-full/templates/resources/global/resource-model.gotmpl */
// Validate compliance
var _ helpers.VyosTopResourceDataModel = &SystemFrrSnmp{}

// SystemFrrSnmp describes the resource data model.
type SystemFrrSnmp struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemFrrSnmpBgpd     types.Bool `tfsdk:"bgpd" vyos:"bgpd,omitempty"`
	LeafSystemFrrSnmpIsisd    types.Bool `tfsdk:"isisd" vyos:"isisd,omitempty"`
	LeafSystemFrrSnmpLdpd     types.Bool `tfsdk:"ldpd" vyos:"ldpd,omitempty"`
	LeafSystemFrrSnmpOspfsixd types.Bool `tfsdk:"ospf6d" vyos:"ospf6d,omitempty"`
	LeafSystemFrrSnmpOspfd    types.Bool `tfsdk:"ospfd" vyos:"ospfd,omitempty"`
	LeafSystemFrrSnmpRIPd     types.Bool `tfsdk:"ripd" vyos:"ripd,omitempty"`
	LeafSystemFrrSnmpZebra    types.Bool `tfsdk:"zebra" vyos:"zebra,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemFrrSnmp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemFrrSnmp) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemFrrSnmp) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemFrrSnmp) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"snmp",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemFrrSnmp) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */
		"system",

		"frr",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *SystemFrrSnmp) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemFrrSnmp) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"bgpd":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `BGP

`,
			Description: `BGP

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"isisd":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `IS-IS

`,
			Description: `IS-IS

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ldpd":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `LDP

`,
			Description: `LDP

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ospf6d":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `OSPFv3

`,
			Description: `OSPFv3

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ospfd":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `OSPFv2

`,
			Description: `OSPFv2

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ripd":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `RIP

`,
			Description: `RIP

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"zebra":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Zebra (IP routing manager)

`,
			Description: `Zebra (IP routing manager)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
