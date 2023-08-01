// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SystemSyslogFile describes the resource data model.
type SystemSyslogFile struct {
	SelfIdentifier types.String `tfsdk:"file_id" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagSystemSyslogFileFacility bool `tfsdk:"facility" vyos:"facility,child"`

	// Nodes
	NodeSystemSyslogFileArchive *SystemSyslogFileArchive `tfsdk:"archive" vyos:"archive,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSyslogFile) GetVyosPath() []string {
	return []string{
		"system",

		"syslog",

		"file",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSyslogFile) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `file_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"file_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Logging to a file

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		// Nodes

		"archive": schema.SingleNestedAttribute{
			Attributes: SystemSyslogFileArchive{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Log file size and rotation characteristics

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemSyslogFile) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemSyslogFile) UnmarshalJSON(_ []byte) error {
	return nil
}
