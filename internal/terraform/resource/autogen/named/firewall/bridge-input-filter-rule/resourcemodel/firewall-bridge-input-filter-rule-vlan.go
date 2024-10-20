/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &FirewallBrIDgeInputFilterRuleVlan{}

// FirewallBrIDgeInputFilterRuleVlan describes the resource data model.
type FirewallBrIDgeInputFilterRuleVlan struct {
	// LeafNodes
	LeafFirewallBrIDgeInputFilterRuleVlanID           types.String `tfsdk:"id" vyos:"id,omitempty"`
	LeafFirewallBrIDgeInputFilterRuleVlanPriority     types.String `tfsdk:"priority" vyos:"priority,omitempty"`
	LeafFirewallBrIDgeInputFilterRuleVlanEthernetType types.String `tfsdk:"ethernet_type" vyos:"ethernet-type,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o FirewallBrIDgeInputFilterRuleVlan) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"id":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Vlan id

    |  Format       |  Description             |
    |---------------|--------------------------|
    |  0-4096       |  Vlan id                 |
    |  <start-end>  |  Vlan id range to match  |
`,
			Description: `Vlan id

    |  Format       |  Description             |
    |---------------|--------------------------|
    |  0-4096       |  Vlan id                 |
    |  <start-end>  |  Vlan id range to match  |
`,
		},

		"priority":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Vlan priority(pcp)

    |  Format       |  Description                   |
    |---------------|--------------------------------|
    |  0-7          |  Vlan priority                 |
    |  <start-end>  |  Vlan priority range to match  |
`,
			Description: `Vlan priority(pcp)

    |  Format       |  Description                   |
    |---------------|--------------------------------|
    |  0-7          |  Vlan priority                 |
    |  <start-end>  |  Vlan priority range to match  |
`,
		},

		"ethernet_type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Ethernet type

    |  Format   |  Description                  |
    |-----------|-------------------------------|
    |  802.1q   |  Customer VLAN tag type       |
    |  802.1ad  |  Service VLAN tag type        |
    |  arp      |  Adress Resolution Protocol   |
    |  _ipv4    |  Internet Protocol version 4  |
    |  _ipv6    |  Internet Protocol version 6  |
`,
			Description: `Ethernet type

    |  Format   |  Description                  |
    |-----------|-------------------------------|
    |  802.1q   |  Customer VLAN tag type       |
    |  802.1ad  |  Service VLAN tag type        |
    |  arp      |  Adress Resolution Protocol   |
    |  _ipv4    |  Internet Protocol version 4  |
    |  _ipv6    |  Internet Protocol version 6  |
`,
		},

		// Nodes

	}
}
