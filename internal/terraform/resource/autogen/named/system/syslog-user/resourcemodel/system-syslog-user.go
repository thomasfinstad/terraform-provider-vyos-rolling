// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SystemSyslogUser describes the resource data model.
type SystemSyslogUser struct {
	SelfIdentifier types.String `tfsdk:"user_id" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagSystemSyslogUserFacility bool `tfsdk:"facility" vyos:"facility,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSyslogUser) GetVyosPath() []string {
	return []string{
		"system",

		"syslog",

		"user",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSyslogUser) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"user_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Logging to specific terminal of given user

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  username  &emsp; |  user login name  |

`,
		},

		// LeafNodes

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemSyslogUser) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemSyslogUser) UnmarshalJSON(_ []byte) error {
	return nil
}
