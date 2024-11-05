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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsEigrp{}

// VrfNameProtocolsEigrp describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type VrfNameProtocolsEigrp struct {
	// LeafNodes
	LeafVrfNameProtocolsEigrpSystemAs         types.Number `tfsdk:"system_as" vyos:"system-as,omitempty"`
	LeafVrfNameProtocolsEigrpMaximumPaths     types.Number `tfsdk:"maximum_paths" vyos:"maximum-paths,omitempty"`
	LeafVrfNameProtocolsEigrpNetwork          types.List   `tfsdk:"network" vyos:"network,omitempty"`
	LeafVrfNameProtocolsEigrpPassiveInterface types.List   `tfsdk:"passive_interface" vyos:"passive-interface,omitempty"`
	LeafVrfNameProtocolsEigrpRedistribute     types.List   `tfsdk:"redistribute" vyos:"redistribute,omitempty"`
	LeafVrfNameProtocolsEigrpRouterID         types.String `tfsdk:"router_id" vyos:"router-id,omitempty"`
	LeafVrfNameProtocolsEigrpVariance         types.Number `tfsdk:"variance" vyos:"variance,omitempty"`

	// TagNodes

	// Nodes

	NodeVrfNameProtocolsEigrpMetric *VrfNameProtocolsEigrpMetric `tfsdk:"metric" vyos:"metric,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsEigrp) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"system_as":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Autonomous System Number (ASN)

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  1-65535  |  Autonomous System Number  |
`,
			Description: `Autonomous System Number (ASN)

    |  Format   |  Description               |
    |-----------|----------------------------|
    |  1-65535  |  Autonomous System Number  |
`,
		},

		"maximum_paths":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Forward packets over multiple paths

    |  Format  |  Description      |
    |----------|-------------------|
    |  1-32    |  Number of paths  |
`,
			Description: `Forward packets over multiple paths

    |  Format  |  Description      |
    |----------|-------------------|
    |  1-32    |  Number of paths  |
`,
		},

		"network":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Enable routing on an IP network

    |  Format   |  Description           |
    |-----------|------------------------|
    |  ipv4net  |  EIGRP network prefix  |
`,
			Description: `Enable routing on an IP network

    |  Format   |  Description           |
    |-----------|------------------------|
    |  ipv4net  |  EIGRP network prefix  |
`,
		},

		"passive_interface":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Suppress routing updates on an interface

`,
			Description: `Suppress routing updates on an interface

`,
		},

		"redistribute":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Redistribute information from another routing protocol

    |  Format     |  Description                          |
    |-------------|---------------------------------------|
    |  bgp        |  Border Gateway Protocol (BGP)        |
    |  connected  |  Connected routes                     |
    |  nhrp       |  Next Hop Resolution Protocol (NHRP)  |
    |  ospf       |  Open Shortest Path First (OSPFv2)    |
    |  rip        |  Routing Information Protocol (RIP)   |
    |  babel      |  Babel routing protocol (Babel)       |
    |  static     |  Statically configured routes         |
    |  vnc        |  Virtual Network Control (VNC)        |
`,
			Description: `Redistribute information from another routing protocol

    |  Format     |  Description                          |
    |-------------|---------------------------------------|
    |  bgp        |  Border Gateway Protocol (BGP)        |
    |  connected  |  Connected routes                     |
    |  nhrp       |  Next Hop Resolution Protocol (NHRP)  |
    |  ospf       |  Open Shortest Path First (OSPFv2)    |
    |  rip        |  Routing Information Protocol (RIP)   |
    |  babel      |  Babel routing protocol (Babel)       |
    |  static     |  Statically configured routes         |
    |  vnc        |  Virtual Network Control (VNC)        |
`,
		},

		"router_id":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override default router identifier

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv4    |  Router-ID in IP address format  |
`,
			Description: `Override default router identifier

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  ipv4    |  Router-ID in IP address format  |
`,
		},

		"variance":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Control load balancing variance

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  1-128   |  Metric variance multiplier  |
`,
			Description: `Control load balancing variance

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  1-128   |  Metric variance multiplier  |
`,
		},

		// TagNodes

		// Nodes

		"metric": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsEigrpMetric{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Modify metrics and parameters for advertisement

`,
			Description: `Modify metrics and parameters for advertisement

`,
		},
	}
}
