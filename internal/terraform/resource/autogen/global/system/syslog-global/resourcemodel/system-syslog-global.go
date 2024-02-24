// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// SystemSyslogGlobal describes the resource data model.
type SystemSyslogGlobal struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafSystemSyslogGlobalPreserveFqdn types.Bool `tfsdk:"preserve_fqdn" vyos:"preserve-fqdn,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagSystemSyslogGlobalFacility bool `tfsdk:"-" vyos:"facility,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeSystemSyslogGlobalMarker bool `tfsdk:"-" vyos:"marker,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *SystemSyslogGlobal) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSyslogGlobal) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"system",

		"syslog",

		"global",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSyslogGlobal) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"preserve_fqdn": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `uses FQDN for logging

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
