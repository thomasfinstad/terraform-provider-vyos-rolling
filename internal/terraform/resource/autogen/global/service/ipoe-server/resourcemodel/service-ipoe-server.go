// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceIPoeServer describes the resource data model.
type ServiceIPoeServer struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafServiceIPoeServerNameServer types.List `tfsdk:"name_server" vyos:"name-server,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceIPoeServerInterface bool `tfsdk:"-" vyos:"interface,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceIPoeServerClientIPPool     bool `tfsdk:"-" vyos:"client-ip-pool,ignore,omitempty"`
	ExistsNodeServiceIPoeServerClientIPvsixPool bool `tfsdk:"-" vyos:"client-ipv6-pool,ignore,omitempty"`
	ExistsNodeServiceIPoeServerAuthentication   bool `tfsdk:"-" vyos:"authentication,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceIPoeServer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"ipoe-server",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServer) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"name_server": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain Name Servers (DNS) addresses

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Domain Name Server (DNS) IPv4 address  |
    |  ipv6  &emsp; |  Domain Name Server (DNS) IPv6 address  |

`,
		},
	}
}