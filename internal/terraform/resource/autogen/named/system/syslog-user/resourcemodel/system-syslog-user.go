// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SystemSyslogUser describes the resource data model.
type SystemSyslogUser struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"user_id" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagSystemSyslogUserFacility bool `tfsdk:"facility" vyos:"facility,child"`

	// Nodes
}

// GetID returns the resource ID
func (o SystemSyslogUser) GetID() *types.String {
	return &o.ID
}

// SetID configures the resource ID
func (o SystemSyslogUser) SetID(id types.String) {
	o.ID = id
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
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"user_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Logging to specific terminal of given user

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  username  &emsp; |  user login name  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		// Nodes

	}
}
