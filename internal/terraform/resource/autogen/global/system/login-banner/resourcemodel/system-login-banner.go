// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemLoginBanner describes the resource data model.
type SystemLoginBanner struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafSystemLoginBannerPostLogin types.String `tfsdk:"post_login" vyos:"post-login,omitempty"`
	LeafSystemLoginBannerPreLogin  types.String `tfsdk:"pre_login" vyos:"pre-login,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *SystemLoginBanner) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemLoginBanner) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"system",

		"login",

		"banner",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemLoginBanner) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"post_login": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `A system banner after the user logs in

`,
		},

		"pre_login": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `A system banner before the user logs in

`,
		},
	}
}
