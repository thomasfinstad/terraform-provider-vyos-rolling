// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceSaltMinion describes the resource data model.
type ServiceSaltMinion struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafServiceSaltMinionHash            types.String `tfsdk:"hash" vyos:"hash,omitempty"`
	LeafServiceSaltMinionMaster          types.List   `tfsdk:"master" vyos:"master,omitempty"`
	LeafServiceSaltMinionID              types.String `tfsdk:"id_param" vyos:"id,omitempty"`
	LeafServiceSaltMinionInterval        types.Number `tfsdk:"interval" vyos:"interval,omitempty"`
	LeafServiceSaltMinionMasterKey       types.String `tfsdk:"master_key" vyos:"master-key,omitempty"`
	LeafServiceSaltMinionSourceInterface types.String `tfsdk:"source_interface" vyos:"source-interface,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ServiceSaltMinion) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSaltMinion) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"salt-minion",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSaltMinion) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"hash": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hash used when discovering file on master server (default: sha256)

`,

			// Default:          stringdefault.StaticString(`sha256`),
			Computed: true,
		},

		"master": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Hostname or IP address of the Salt master server

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Salt server IPv4 address  |
    |  ipv6  &emsp; |  Salt server IPv6 address  |
    |  hostname  &emsp; |  Salt server FQDN address  |

`,
		},

		"id_param": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Explicitly declare ID for this minion to use (default: hostname)

`,
		},

		"interval": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Interval in minutes between updates (default: 60)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-1440  &emsp; |  Update interval in minutes  |

`,

			// Default:          stringdefault.StaticString(`60`),
			Computed: true,
		},

		"master_key": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `URL with signature of master for auth reply verification

`,
		},

		"source_interface": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Interface used to establish connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  interface  &emsp; |  Interface name  |

`,
		},
	}
}
