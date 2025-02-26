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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (flow-accounting) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &SystemFlowAccounting{}

// SystemFlowAccounting describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type SystemFlowAccounting struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemFlowAccountingBufferSize     types.Number `tfsdk:"buffer_size" vyos:"buffer-size,omitempty"`
	LeafSystemFlowAccountingPacketLength   types.Number `tfsdk:"packet_length" vyos:"packet-length,omitempty"`
	LeafSystemFlowAccountingEnableEgress   types.Bool   `tfsdk:"enable_egress" vyos:"enable-egress,omitempty"`
	LeafSystemFlowAccountingDisableImt     types.Bool   `tfsdk:"disable_imt" vyos:"disable-imt,omitempty"`
	LeafSystemFlowAccountingSyslogFacility types.String `tfsdk:"syslog_facility" vyos:"syslog-facility,omitempty"`
	LeafSystemFlowAccountingInterface      types.List   `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafSystemFlowAccountingVrf            types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes

	// Nodes

	ExistsNodeSystemFlowAccountingNetflow bool `tfsdk:"-" vyos:"netflow,child"`
}

// SetID configures the resource ID
func (o *SystemFlowAccounting) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemFlowAccounting) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemFlowAccounting) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemFlowAccounting) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"flow-accounting",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemFlowAccounting) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (system) */
		"system", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *SystemFlowAccounting) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (system) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemFlowAccounting) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"buffer_size":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (buffer-size) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Buffer size

    |  Format  |  Description         |
    |----------|----------------------|
    |  u32     |  Buffer size in MiB  |
`,
			Description: `Buffer size

    |  Format  |  Description         |
    |----------|----------------------|
    |  u32     |  Buffer size in MiB  |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"packet_length":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (packet-length) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Specifies the maximum number of bytes to capture for each packet

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  128-750  |  Packet length in bytes  |
`,
			Description: `Specifies the maximum number of bytes to capture for each packet

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  128-750  |  Packet length in bytes  |
`,

			// Default:          stringdefault.StaticString(`128`),
			Computed: true,
		},

		"enable_egress":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (enable-egress) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable egress flow accounting

`,
			Description: `Enable egress flow accounting

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_imt":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable-imt) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable in memory table plugin

`,
			Description: `Disable in memory table plugin

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"syslog_facility":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (syslog-facility) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Syslog facility for flow-accounting

    |  Format     |  Description                       |
    |-------------|------------------------------------|
    |  auth       |  Authentication and authorization  |
    |  authpriv   |  Non-system authorization          |
    |  cron       |  Cron daemon                       |
    |  daemon     |  System daemons                    |
    |  kern       |  Kernel                            |
    |  lpr        |  Line printer spooler              |
    |  mail       |  Mail subsystem                    |
    |  mark       |  Timestamp                         |
    |  news       |  USENET subsystem                  |
    |  protocols  |  Routing protocols (local7)        |
    |  security   |  Authentication and authorization  |
    |  syslog     |  Authentication and authorization  |
    |  user       |  Application processes             |
    |  uucp       |  UUCP subsystem                    |
    |  local0     |  Local facility 0                  |
    |  local1     |  Local facility 1                  |
    |  local2     |  Local facility 2                  |
    |  local3     |  Local facility 3                  |
    |  local4     |  Local facility 4                  |
    |  local5     |  Local facility 5                  |
    |  local6     |  Local facility 6                  |
    |  local7     |  Local facility 7                  |
    |  all        |  Authentication and authorization  |
`,
			Description: `Syslog facility for flow-accounting

    |  Format     |  Description                       |
    |-------------|------------------------------------|
    |  auth       |  Authentication and authorization  |
    |  authpriv   |  Non-system authorization          |
    |  cron       |  Cron daemon                       |
    |  daemon     |  System daemons                    |
    |  kern       |  Kernel                            |
    |  lpr        |  Line printer spooler              |
    |  mail       |  Mail subsystem                    |
    |  mark       |  Timestamp                         |
    |  news       |  USENET subsystem                  |
    |  protocols  |  Routing protocols (local7)        |
    |  security   |  Authentication and authorization  |
    |  syslog     |  Authentication and authorization  |
    |  user       |  Application processes             |
    |  uucp       |  UUCP subsystem                    |
    |  local0     |  Local facility 0                  |
    |  local1     |  Local facility 1                  |
    |  local2     |  Local facility 2                  |
    |  local3     |  Local facility 3                  |
    |  local4     |  Local facility 4                  |
    |  local5     |  Local facility 5                  |
    |  local6     |  Local facility 6                  |
    |  local7     |  Local facility 7                  |
    |  all        |  Authentication and authorization  |
`,
		},

		"interface":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (interface) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Interface

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		"vrf":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (vrf) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			Description: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
		},

		// TagNodes

		// Nodes

	}
}
