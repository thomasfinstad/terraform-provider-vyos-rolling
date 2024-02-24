// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceHTTPS describes the resource data model.
type ServiceHTTPS struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafServiceHTTPSVrf types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceHTTPSVirtualHost bool `tfsdk:"-" vyos:"virtual-host,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceHTTPSAPI          bool `tfsdk:"-" vyos:"api,ignore,omitempty"`
	ExistsNodeServiceHTTPSAPIRestrict  bool `tfsdk:"-" vyos:"api-restrict,ignore,omitempty"`
	ExistsNodeServiceHTTPSCertificates bool `tfsdk:"-" vyos:"certificates,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceHTTPS) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceHTTPS) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"https",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceHTTPS) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"vrf": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |

`,
		},
	}
}
