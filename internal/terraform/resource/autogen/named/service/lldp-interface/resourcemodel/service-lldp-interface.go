// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceLldpInterface describes the resource data model.
type ServiceLldpInterface struct {
	SelfIdentifier types.String `tfsdk:"interface_id" vyos:",self-id"`

	// LeafNodes
	LeafServiceLldpInterfaceDisable types.Bool `tfsdk:"disable" vyos:"disable,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeServiceLldpInterfaceLocation *ServiceLldpInterfaceLocation `tfsdk:"location" vyos:"location,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceLldpInterface) GetVyosPath() []string {
	return []string{
		"service",

		"lldp",

		"interface",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceLldpInterface) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Location data for interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  all  &emsp; |  Location data all interfaces  |
    |  txt  &emsp; |  Location data for a specific interface  |

`,
		},

		// LeafNodes

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// Nodes

		"location": schema.SingleNestedAttribute{
			Attributes: ServiceLldpInterfaceLocation{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `LLDP-MED location data

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceLldpInterface) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceLldpInterface) UnmarshalJSON(_ []byte) error {
	return nil
}
