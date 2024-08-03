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
var _ helpers.VyosTopResourceDataModel = &SystemLcd{}

// SystemLcd describes the resource data model.
type SystemLcd struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemLcdModel  types.String `tfsdk:"model" vyos:"model,omitempty"`
	LeafSystemLcdDevice types.String `tfsdk:"device" vyos:"device,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemLcd) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemLcd) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemLcd) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemLcd) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"lcd",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *SystemLcd) GetVyosParentPath() []string {
	return []string{
		"system",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *SystemLcd) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLcd) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"model": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Model of the display attached to this system

    |  Format   |  Description                                            |
    |-----------|---------------------------------------------------------|
    |  cfa-533  |  Crystalfontz CFA-533                                   |
    |  cfa-631  |  Crystalfontz CFA-631                                   |
    |  cfa-633  |  Crystalfontz CFA-633                                   |
    |  cfa-635  |  Crystalfontz CFA-635                                   |
    |  hd44780  |  Hitachi HD44780, Caswell Appliances                    |
    |  sdec     |  Lanner, Watchguard, Nexcom NSA, Sophos UTM appliances  |
`,
			Description: `Model of the display attached to this system

    |  Format   |  Description                                            |
    |-----------|---------------------------------------------------------|
    |  cfa-533  |  Crystalfontz CFA-533                                   |
    |  cfa-631  |  Crystalfontz CFA-631                                   |
    |  cfa-633  |  Crystalfontz CFA-633                                   |
    |  cfa-635  |  Crystalfontz CFA-635                                   |
    |  hd44780  |  Hitachi HD44780, Caswell Appliances                    |
    |  sdec     |  Lanner, Watchguard, Nexcom NSA, Sophos UTM appliances  |
`,
		},

		"device": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Physical device used by LCD display

    |  Format    |  Description                           |
    |------------|----------------------------------------|
    |  ttySXX    |  TTY device name, regular serial port  |
    |  usbNbXpY  |  TTY device name, USB based            |
`,
			Description: `Physical device used by LCD display

    |  Format    |  Description                           |
    |------------|----------------------------------------|
    |  ttySXX    |  TTY device name, regular serial port  |
    |  usbNbXpY  |  TTY device name, USB based            |
`,
		},
	}
}
