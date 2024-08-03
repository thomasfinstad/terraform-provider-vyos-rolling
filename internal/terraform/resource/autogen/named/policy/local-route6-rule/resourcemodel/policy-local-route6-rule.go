// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &PolicyLocalRoutesixRule{}

// PolicyLocalRoutesixRule describes the resource data model.
type PolicyLocalRoutesixRule struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafPolicyLocalRoutesixRuleFwmark           types.Number `tfsdk:"fwmark" vyos:"fwmark,omitempty"`
	LeafPolicyLocalRoutesixRuleProtocol         types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafPolicyLocalRoutesixRuleInboundInterface types.String `tfsdk:"inbound_interface" vyos:"inbound-interface,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodePolicyLocalRoutesixRuleSet         *PolicyLocalRoutesixRuleSet         `tfsdk:"set" vyos:"set,omitempty"`
	NodePolicyLocalRoutesixRuleSource      *PolicyLocalRoutesixRuleSource      `tfsdk:"source" vyos:"source,omitempty"`
	NodePolicyLocalRoutesixRuleDestination *PolicyLocalRoutesixRuleDestination `tfsdk:"destination" vyos:"destination,omitempty"`
}

// SetID configures the resource ID
func (o *PolicyLocalRoutesixRule) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *PolicyLocalRoutesixRule) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *PolicyLocalRoutesixRule) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *PolicyLocalRoutesixRule) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"rule",
		o.SelfIdentifier.Attributes()["rule"].(types.Number).ValueBigFloat().String(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *PolicyLocalRoutesixRule) GetVyosParentPath() []string {
	return []string{
		"policy",

		"local-route6",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *PolicyLocalRoutesixRule) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o PolicyLocalRoutesixRule) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.MapNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"rule": schema.NumberAttribute{
						Required: true,
						MarkdownDescription: `IPv6 policy local-route rule set number

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  1-32765  |  Local-route rule number (1-32765)  |
`,
						Description: `IPv6 policy local-route rule set number

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  1-32765  |  Local-route rule number (1-32765)  |
`,
						PlanModifiers: []planmodifier.Number{
							numberplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"fwmark": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Match fwmark value

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  1-2147483647  |  Address to match against  |
`,
			Description: `Match fwmark value

    |  Format        |  Description               |
    |----------------|----------------------------|
    |  1-2147483647  |  Address to match against  |
`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol to match (protocol name or number)

    |  Format      |  Description         |
    |--------------|----------------------|
    |  0-255       |  IP protocol number  |
    |  <protocol>  |  IP protocol name    |
`,
			Description: `Protocol to match (protocol name or number)

    |  Format      |  Description         |
    |--------------|----------------------|
    |  0-255       |  IP protocol number  |
    |  <protocol>  |  IP protocol name    |
`,
		},

		"inbound_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Inbound Interface

`,
			Description: `Inbound Interface

`,
		},

		// Nodes

		"set": schema.SingleNestedAttribute{
			Attributes: PolicyLocalRoutesixRuleSet{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Packet modifications

`,
			Description: `Packet modifications

`,
		},

		"source": schema.SingleNestedAttribute{
			Attributes: PolicyLocalRoutesixRuleSource{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Source parameters

`,
			Description: `Source parameters

`,
		},

		"destination": schema.SingleNestedAttribute{
			Attributes: PolicyLocalRoutesixRuleDestination{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Destination parameters

`,
			Description: `Destination parameters

`,
		},
	}
}
