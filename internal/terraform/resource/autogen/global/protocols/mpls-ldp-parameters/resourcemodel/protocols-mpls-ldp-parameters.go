// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

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

// Validate compliance
var _ helpers.VyosTopResourceDataModel = &ProtocolsMplsLdpParameters{}

// ProtocolsMplsLdpParameters describes the resource data model.
type ProtocolsMplsLdpParameters struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsMplsLdpParametersCiscoInteropTlv        types.Bool `tfsdk:"cisco_interop_tlv" vyos:"cisco-interop-tlv,omitempty"`
	LeafProtocolsMplsLdpParametersTransportPreferIPvfour types.Bool `tfsdk:"transport_prefer_ipv4" vyos:"transport-prefer-ipv4,omitempty"`
	LeafProtocolsMplsLdpParametersOrderedControl         types.Bool `tfsdk:"ordered_control" vyos:"ordered-control,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ProtocolsMplsLdpParameters) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsMplsLdpParameters) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsMplsLdpParameters) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsMplsLdpParameters) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"parameters",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsMplsLdpParameters) GetVyosParentPath() []string {
	return []string{
		"protocols",

		"mpls",

		"ldp",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsMplsLdpParameters) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsMplsLdpParameters) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"cisco_interop_tlv": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable Cisco non-compliant format capability TLV

`,
			Description: `Enable Cisco non-compliant format capability TLV

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"transport_prefer_ipv4": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Prefer IPv4 for TCP peer transport connection

`,
			Description: `Prefer IPv4 for TCP peer transport connection

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ordered_control": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable LDP ordered label distribution control mode

`,
			Description: `Enable LDP ordered label distribution control mode

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
