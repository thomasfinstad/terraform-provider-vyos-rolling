/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ProtocolsOspfAreaVirtualLink{}

// ProtocolsOspfAreaVirtualLink describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ProtocolsOspfAreaVirtualLink struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl */
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsOspfAreaVirtualLinkDeadInterval       types.Number `tfsdk:"dead_interval" vyos:"dead-interval,omitempty"`
	LeafProtocolsOspfAreaVirtualLinkHelloInterval      types.Number `tfsdk:"hello_interval" vyos:"hello-interval,omitempty"`
	LeafProtocolsOspfAreaVirtualLinkRetransmitInterval types.Number `tfsdk:"retransmit_interval" vyos:"retransmit-interval,omitempty"`
	LeafProtocolsOspfAreaVirtualLinkTransmitDelay      types.Number `tfsdk:"transmit_delay" vyos:"transmit-delay,omitempty"`
	LeafProtocolsOspfAreaVirtualLinkRetransmitWindow   types.Number `tfsdk:"retransmit_window" vyos:"retransmit-window,omitempty"`

	// TagNodes

	// Nodes

	NodeProtocolsOspfAreaVirtualLinkAuthentication *ProtocolsOspfAreaVirtualLinkAuthentication `tfsdk:"authentication" vyos:"authentication,omitempty"`
}

// SetID configures the resource ID
func (o *ProtocolsOspfAreaVirtualLink) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsOspfAreaVirtualLink) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsOspfAreaVirtualLink) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsOspfAreaVirtualLink) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"virtual-link",
		o.SelfIdentifier.Attributes()["virtual_link"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsOspfAreaVirtualLink) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"protocols", // Node

		"ospf", // Node

		"area",
		o.SelfIdentifier.Attributes()["area"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ProtocolsOspfAreaVirtualLink) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"protocols", // Node

		"ospf", // Node

		"area",
		o.SelfIdentifier.Attributes()["area"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsOspfAreaVirtualLink) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"virtual_link": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Virtual link

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  ipv4    |  OSPF area in dotted decimal notation  |
`,
					Description: `Virtual link

    |  Format  |  Description                           |
    |----------|----------------------------------------|
    |  ipv4    |  OSPF area in dotted decimal notation  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in virtual_link, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  virtual_link, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				"area": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `OSPF area settings

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  u32     |  OSPF area number in decimal notation         |
    |  ipv4    |  OSPF area number in dotted decimal notation  |
`,
					Description: `OSPF area settings

    |  Format  |  Description                                  |
    |----------|-----------------------------------------------|
    |  u32     |  OSPF area number in decimal notation         |
    |  ipv4    |  OSPF area number in dotted decimal notation  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in area, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  area, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"dead_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval after which a neighbor is declared dead

    |  Format   |  Description                       |
    |-----------|------------------------------------|
    |  1-65535  |  Neighbor dead interval (seconds)  |
`,
			Description: `Interval after which a neighbor is declared dead

    |  Format   |  Description                       |
    |-----------|------------------------------------|
    |  1-65535  |  Neighbor dead interval (seconds)  |
`,

			// Default:          stringdefault.StaticString(`40`),
			Computed: true,
		},

		"hello_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval between hello packets

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  1-65535  |  Hello interval (seconds)  |
`,
			Description: `Interval between hello packets

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  1-65535  |  Hello interval (seconds)  |
`,

			// Default:          stringdefault.StaticString(`10`),
			Computed: true,
		},

		"retransmit_interval":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval between retransmitting lost link state advertisements

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Retransmit interval (seconds)  |
`,
			Description: `Interval between retransmitting lost link state advertisements

    |  Format   |  Description                    |
    |-----------|---------------------------------|
    |  1-65535  |  Retransmit interval (seconds)  |
`,

			// Default:          stringdefault.StaticString(`5`),
			Computed: true,
		},

		"transmit_delay":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Link state transmit delay

    |  Format   |  Description                          |
    |-----------|---------------------------------------|
    |  1-65535  |  Link state transmit delay (seconds)  |
`,
			Description: `Link state transmit delay

    |  Format   |  Description                          |
    |-----------|---------------------------------------|
    |  1-65535  |  Link state transmit delay (seconds)  |
`,

			// Default:          stringdefault.StaticString(`1`),
			Computed: true,
		},

		"retransmit_window":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Window for LSA retransmit

    |  Format   |  Description                                             |
    |-----------|----------------------------------------------------------|
    |  20-1000  |  Retransmit LSAs expiring in this window (milliseconds)  |
`,
			Description: `Window for LSA retransmit

    |  Format   |  Description                                             |
    |-----------|----------------------------------------------------------|
    |  20-1000  |  Retransmit LSAs expiring in this window (milliseconds)  |
`,

			// Default:          stringdefault.StaticString(`50`),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: ProtocolsOspfAreaVirtualLinkAuthentication{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Authentication

`,
			Description: `Authentication

`,
		},
	}
}
