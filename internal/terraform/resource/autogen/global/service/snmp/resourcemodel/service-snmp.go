// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceSnmp describes the resource data model.
type ServiceSnmp struct {
	ID types.String `tfsdk:"id" vyos:"_,tfsdk-id"`

	// LeafNodes
	LeafServiceSnmpContact     types.String `tfsdk:"contact" vyos:"contact,omitempty"`
	LeafServiceSnmpDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafServiceSnmpLocation    types.String `tfsdk:"location" vyos:"location,omitempty"`
	LeafServiceSnmpOIDEnable   types.List   `tfsdk:"oid_enable" vyos:"oid-enable,omitempty"`
	LeafServiceSnmpProtocol    types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`
	LeafServiceSnmpSmuxPeer    types.List   `tfsdk:"smux_peer" vyos:"smux-peer,omitempty"`
	LeafServiceSnmpTrapSource  types.String `tfsdk:"trap_source" vyos:"trap-source,omitempty"`
	LeafServiceSnmpVrf         types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagServiceSnmpCommunity     bool `tfsdk:"-" vyos:"community,ignore,child"`
	ExistsTagServiceSnmpListenAddress bool `tfsdk:"-" vyos:"listen-address,ignore,child"`
	ExistsTagServiceSnmpTrapTarget    bool `tfsdk:"-" vyos:"trap-target,ignore,child"`

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceSnmpVthree           bool `tfsdk:"-" vyos:"v3,ignore,omitempty"`
	ExistsNodeServiceSnmpScrIPtExtensions bool `tfsdk:"-" vyos:"script-extensions,ignore,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceSnmp) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceSnmp) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"snmp",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceSnmp) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"contact": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Contact information

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description information

`,
		},

		"location": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Location information

`,
		},

		"oid_enable": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Enable specific OIDs that by default are disable

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ip-forward  &emsp; |  Enable ipForward: .1.3.6.1.2.1.4.24  |
    |  ip-route-table  &emsp; |  Enable ipRouteTable: .1.3.6.1.2.1.4.21  |
    |  ip-net-to-media-table  &emsp; |  Enable ipNetToMediaTable: .1.3.6.1.2.1.4.22  |
    |  ip-net-to-physical-phys-address  &emsp; |  Enable ipNetToPhysicalPhysAddress: .1.3.6.1.2.1.4.35  |

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Protocol to be used (TCP/UDP)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  udp  &emsp; |  Listen protocol UDP  |
    |  tcp  &emsp; |  Listen protocol TCP  |

`,

			// Default:          stringdefault.StaticString(`udp`),
			Computed: true,
		},

		"smux_peer": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Register a subtree for SMUX-based processing

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  SNMP Object Identifier  |

`,
		},

		"trap_source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `SNMP trap source address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address  |
    |  ipv6  &emsp; |  IPv6 address  |

`,
		},

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
