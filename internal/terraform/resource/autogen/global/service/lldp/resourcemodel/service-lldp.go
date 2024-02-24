// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceLldp describes the resource data model.
type ServiceLldp struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafServiceLldpManagementAddress types.List `tfsdk:"management_address" vyos:"management-address,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceLldpInterface bool `tfsdk:"-" vyos:"interface,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceLldpLegacyProtocols bool `tfsdk:"-" vyos:"legacy-protocols,ignore,omitempty"`
	ExistsNodeServiceLldpSnmp            bool `tfsdk:"-" vyos:"snmp,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceLldp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceLldp) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"lldp",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceLldp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"management_address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Management IP Address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 Management Address  |
    |  ipv6  &emsp; |  IPv6 Management Address  |

`,
		},
	}
}
