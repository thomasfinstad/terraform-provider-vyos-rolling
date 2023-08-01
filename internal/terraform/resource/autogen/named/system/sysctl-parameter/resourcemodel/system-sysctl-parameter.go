// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SystemSysctlParameter describes the resource data model.
type SystemSysctlParameter struct {
	SelfIdentifier types.String `tfsdk:"parameter_id" vyos:",self-id"`

	// LeafNodes
	LeafSystemSysctlParameterValue types.String `tfsdk:"value" vyos:"value,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSysctlParameter) GetVyosPath() []string {
	return []string{
		"system",

		"sysctl",

		"parameter",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSysctlParameter) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `parameter_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"parameter_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Sysctl key name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Sysctl key name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"value": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Sysctl configuration value

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemSysctlParameter) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemSysctlParameter) UnmarshalJSON(_ []byte) error {
	return nil
}
