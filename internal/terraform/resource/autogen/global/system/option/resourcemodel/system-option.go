// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemOption describes the resource data model.
type SystemOption struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafSystemOptionCtrlAltDelete           types.String `tfsdk:"ctrl_alt_delete" vyos:"ctrl-alt-delete,omitempty"`
	LeafSystemOptionKeyboardLayout          types.String `tfsdk:"keyboard_layout" vyos:"keyboard-layout,omitempty"`
	LeafSystemOptionPerformance             types.String `tfsdk:"performance" vyos:"performance,omitempty"`
	LeafSystemOptionRebootOnPanic           types.Bool   `tfsdk:"reboot_on_panic" vyos:"reboot-on-panic,omitempty"`
	LeafSystemOptionStartupBeep             types.Bool   `tfsdk:"startup_beep" vyos:"startup-beep,omitempty"`
	LeafSystemOptionRootPartitionAutoResize types.Bool   `tfsdk:"root_partition_auto_resize" vyos:"root-partition-auto-resize,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeSystemOptionHTTPClient bool `tfsdk:"-" vyos:"http-client,child"`
	ExistsNodeSystemOptionTCPClient  bool `tfsdk:"-" vyos:"ssh-client,child"`
}

// SetID configures the resource ID
func (o *SystemOption) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemOption) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"system",

		"option",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemOption) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"ctrl_alt_delete": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `System action on Ctrl-Alt-Delete keystroke

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ignore  &emsp; |  Ignore key sequence  |
    |  reboot  &emsp; |  Reboot system  |
    |  poweroff  &emsp; |  Poweroff system  |

`,
		},

		"keyboard_layout": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `System keyboard layout, type ISO2

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  us  &emsp; |  United States  |
    |  uk  &emsp; |  United Kingdom  |
    |  fr  &emsp; |  France  |
    |  de  &emsp; |  Germany  |
    |  es  &emsp; |  Spain  |
    |  fi  &emsp; |  Finland  |
    |  jp106  &emsp; |  Japan  |
    |  no  &emsp; |  Norway  |
    |  dk  &emsp; |  Denmark  |
    |  dvorak  &emsp; |  Dvorak  |

`,

			// Default:          stringdefault.StaticString(`us`),
			Computed: true,
		},

		"performance": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Tune system performance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  throughput  &emsp; |  Tune for maximum network throughput  |
    |  latency  &emsp; |  Tune for low network latency  |

`,
		},

		"reboot_on_panic": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Reboot system on kernel panic

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"startup_beep": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `plays sound via system speaker when you can login

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"root_partition_auto_resize": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable root partition auto-extention on system boot

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
