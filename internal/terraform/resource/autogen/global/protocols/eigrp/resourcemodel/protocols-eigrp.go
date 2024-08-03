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
var _ helpers.VyosTopResourceDataModel = &ProtocolsEigrp{}

// ProtocolsEigrp describes the resource data model.
type ProtocolsEigrp struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafProtocolsEigrpSystemAs         types.Number `tfsdk:"system_as" vyos:"system-as,omitempty"`
	LeafProtocolsEigrpMaximumPaths     types.Number `tfsdk:"maximum_paths" vyos:"maximum-paths,omitempty"`
	LeafProtocolsEigrpNetwork          types.List   `tfsdk:"network" vyos:"network,omitempty"`
	LeafProtocolsEigrpPassiveInterface types.List   `tfsdk:"passive_interface" vyos:"passive-interface,omitempty"`
	LeafProtocolsEigrpRedistribute     types.List   `tfsdk:"redistribute" vyos:"redistribute,omitempty"`
	LeafProtocolsEigrpRouterID         types.String `tfsdk:"router_id" vyos:"router-id,omitempty"`
	LeafProtocolsEigrpVariance         types.Number `tfsdk:"variance" vyos:"variance,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeProtocolsEigrpMetric bool `tfsdk:"-" vyos:"metric,child"`
}

// SetID configures the resource ID
func (o *ProtocolsEigrp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ProtocolsEigrp) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ProtocolsEigrp) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ProtocolsEigrp) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"eigrp",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ProtocolsEigrp) GetVyosParentPath() []string {
	return []string{
		"protocols",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ProtocolsEigrp) GetVyosNamedParentPath() []string {
	return []string{}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ProtocolsEigrp) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"system_as": schema.NumberAttribute{
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

		"maximum_paths": schema.NumberAttribute{
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

		"network": schema.ListAttribute{
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

		"passive_interface": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Suppress routing updates on an interface

`,
			Description: `Suppress routing updates on an interface

`,
		},

		"redistribute": schema.ListAttribute{
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

		"router_id": schema.StringAttribute{
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

		"variance": schema.NumberAttribute{
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
	}
}
