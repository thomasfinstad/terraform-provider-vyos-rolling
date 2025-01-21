/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (system) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceConfigSyncSectionSystem{}

// ServiceConfigSyncSectionSystem describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServiceConfigSyncSectionSystem struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceConfigSyncSectionSystemConntrack         types.Bool `tfsdk:"conntrack" vyos:"conntrack,omitempty"`
	LeafServiceConfigSyncSectionSystemFlowAccounting    types.Bool `tfsdk:"flow_accounting" vyos:"flow-accounting,omitempty"`
	LeafServiceConfigSyncSectionSystemOption            types.Bool `tfsdk:"option" vyos:"option,omitempty"`
	LeafServiceConfigSyncSectionSystemSflow             types.Bool `tfsdk:"sflow" vyos:"sflow,omitempty"`
	LeafServiceConfigSyncSectionSystemStaticHostMapping types.Bool `tfsdk:"static_host_mapping" vyos:"static-host-mapping,omitempty"`
	LeafServiceConfigSyncSectionSystemSysctl            types.Bool `tfsdk:"sysctl" vyos:"sysctl,omitempty"`
	LeafServiceConfigSyncSectionSystemTimeZone          types.Bool `tfsdk:"time_zone" vyos:"time-zone,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceConfigSyncSectionSystem) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceConfigSyncSectionSystem) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceConfigSyncSectionSystem) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceConfigSyncSectionSystem) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"system",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceConfigSyncSectionSystem) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (section) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (config-sync) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */
		"service", // Node

		"config-sync", // Node

		"section", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceConfigSyncSectionSystem) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (section) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (config-sync) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (service) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceConfigSyncSectionSystem) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"conntrack":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (conntrack) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Connection Tracking

`,
			Description: `Connection Tracking

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"flow_accounting":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (flow-accounting) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Flow accounting

`,
			Description: `Flow accounting

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"option":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (option) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `System Options

`,
			Description: `System Options

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"sflow":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (sflow) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `sFlow

`,
			Description: `sFlow

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"static_host_mapping":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (static-host-mapping) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Map host names to addresses

`,
			Description: `Map host names to addresses

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"sysctl":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (sysctl) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Configure kernel parameters at runtime

`,
			Description: `Configure kernel parameters at runtime

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"time_zone":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (time-zone) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Local time zone

`,
			Description: `Local time zone

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
