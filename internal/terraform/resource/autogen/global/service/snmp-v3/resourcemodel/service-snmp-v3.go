// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceSnmpVthree describes the resource data model.
type ServiceSnmpVthree struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafServiceSnmpVthreeEngineID types.String `tfsdk:"engineid" vyos:"engineid,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceSnmpVthreeGroup      bool `tfsdk:"-" vyos:"group,ignore,child"`
	ExistsTagServiceSnmpVthreeTrapTarget bool `tfsdk:"-" vyos:"trap-target,ignore,child"`
	ExistsTagServiceSnmpVthreeUser       bool `tfsdk:"-" vyos:"user,ignore,child"`
	ExistsTagServiceSnmpVthreeView       bool `tfsdk:"-" vyos:"view,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ServiceSnmpVthree) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSnmpVthree) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"snmp",

		"v3",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmpVthree) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"engineid": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Specifies the EngineID that uniquely identify an agent (e.g. 000000000000000000000002)

`,
		},
	}
}
