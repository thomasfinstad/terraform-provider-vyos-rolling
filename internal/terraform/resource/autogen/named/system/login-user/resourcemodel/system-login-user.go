// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SystemLoginUser describes the resource data model.
type SystemLoginUser struct {
	SelfIdentifier types.String `tfsdk:"user_id" vyos:",self-id"`

	// LeafNodes
	LeafSystemLoginUserFullName      types.String `tfsdk:"full_name" vyos:"full-name,omitempty"`
	LeafSystemLoginUserHomeDirectory types.String `tfsdk:"home_directory" vyos:"home-directory,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeSystemLoginUserAuthentication *SystemLoginUserAuthentication `tfsdk:"authentication" vyos:"authentication,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemLoginUser) GetVyosPath() []string {
	return []string{
		"system",

		"login",

		"user",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginUser) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `user_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"user_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Local user account information

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"full_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Full name of the user (use quotes for names with spaces)

`,
		},

		"home_directory": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Home directory

`,
		},

		// Nodes

		"authentication": schema.SingleNestedAttribute{
			Attributes: SystemLoginUserAuthentication{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Authentication settings

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemLoginUser) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemLoginUser) UnmarshalJSON(_ []byte) error {
	return nil
}
