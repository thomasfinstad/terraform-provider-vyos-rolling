/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesMacsecIPvsix{}

// InterfacesMacsecIPvsix describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesMacsecIPvsix struct {
	// LeafNodes
	LeafInterfacesMacsecIPvsixAdjustMss              types.String `tfsdk:"adjust_mss" vyos:"adjust-mss,omitempty"`
	LeafInterfacesMacsecIPvsixBaseReachableTime      types.Number `tfsdk:"base_reachable_time" vyos:"base-reachable-time,omitempty"`
	LeafInterfacesMacsecIPvsixDisableForwarding      types.Bool   `tfsdk:"disable_forwarding" vyos:"disable-forwarding,omitempty"`
	LeafInterfacesMacsecIPvsixAcceptDad              types.String `tfsdk:"accept_dad" vyos:"accept-dad,omitempty"`
	LeafInterfacesMacsecIPvsixDupAddrDetectTransmits types.Number `tfsdk:"dup_addr_detect_transmits" vyos:"dup-addr-detect-transmits,omitempty"`
	LeafInterfacesMacsecIPvsixSourceValIDation       types.String `tfsdk:"source_validation" vyos:"source-validation,omitempty"`

	// TagNodes

	// Nodes

	NodeInterfacesMacsecIPvsixAddress *InterfacesMacsecIPvsixAddress `tfsdk:"address" vyos:"address,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesMacsecIPvsix) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"adjust_mss":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Adjust TCP MSS value

    |  Format             |  Description                                     |
    |---------------------|--------------------------------------------------|
    |  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
    |  536-65535          |  TCP Maximum segment size in bytes               |
`,
			Description: `Adjust TCP MSS value

    |  Format             |  Description                                     |
    |---------------------|--------------------------------------------------|
    |  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
    |  536-65535          |  TCP Maximum segment size in bytes               |
`,
		},

		"base_reachable_time":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Base reachable time in seconds

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  1-86400  |  Base reachable time in seconds  |
`,
			Description: `Base reachable time in seconds

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  1-86400  |  Base reachable time in seconds  |
`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"disable_forwarding":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
			Description: `Disable IP forwarding on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"accept_dad":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Accept Duplicate Address Detection

    |  Format  |  Description                                                                |
    |----------|-----------------------------------------------------------------------------|
    |  0       |  Disable DAD                                                                |
    |  1       |  Enable DAD                                                                 |
    |  2       |  Enable DAD - disable IPv6 if MAC-based duplicate link-local address found  |
`,
			Description: `Accept Duplicate Address Detection

    |  Format  |  Description                                                                |
    |----------|-----------------------------------------------------------------------------|
    |  0       |  Disable DAD                                                                |
    |  1       |  Enable DAD                                                                 |
    |  2       |  Enable DAD - disable IPv6 if MAC-based duplicate link-local address found  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"dup_addr_detect_transmits":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Number of NS messages to send while performing DAD

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  0       |  Disable Duplicate Address Dectection (DAD)          |
    |  1-n     |  Number of NS messages to send while performing DAD  |
`,
			Description: `Number of NS messages to send while performing DAD

    |  Format  |  Description                                         |
    |----------|------------------------------------------------------|
    |  0       |  Disable Duplicate Address Dectection (DAD)          |
    |  1-n     |  Number of NS messages to send while performing DAD  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"source_validation":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source validation by reversed path (RFC3704)

    |  Format   |  Description                                                  |
    |-----------|---------------------------------------------------------------|
    |  strict   |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose    |  Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    |  disable  |  No source validation                                         |
`,
			Description: `Source validation by reversed path (RFC3704)

    |  Format   |  Description                                                  |
    |-----------|---------------------------------------------------------------|
    |  strict   |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose    |  Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    |  disable  |  No source validation                                         |
`,
		},

		// TagNodes

		// Nodes

		"address": schema.SingleNestedAttribute{
			Attributes: InterfacesMacsecIPvsixAddress{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 address configuration modes

`,
			Description: `IPv6 address configuration modes

`,
		},
	}
}
