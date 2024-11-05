/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceConfigSyncSectionProtocols{}

// ServiceConfigSyncSectionProtocols describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServiceConfigSyncSectionProtocols struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceConfigSyncSectionProtocolsBabel          types.Bool `tfsdk:"babel" vyos:"babel,omitempty"`
	LeafServiceConfigSyncSectionProtocolsBfd            types.Bool `tfsdk:"bfd" vyos:"bfd,omitempty"`
	LeafServiceConfigSyncSectionProtocolsBgp            types.Bool `tfsdk:"bgp" vyos:"bgp,omitempty"`
	LeafServiceConfigSyncSectionProtocolsFailover       types.Bool `tfsdk:"failover" vyos:"failover,omitempty"`
	LeafServiceConfigSyncSectionProtocolsIgmpProxy      types.Bool `tfsdk:"igmp_proxy" vyos:"igmp-proxy,omitempty"`
	LeafServiceConfigSyncSectionProtocolsIsis           types.Bool `tfsdk:"isis" vyos:"isis,omitempty"`
	LeafServiceConfigSyncSectionProtocolsMpls           types.Bool `tfsdk:"mpls" vyos:"mpls,omitempty"`
	LeafServiceConfigSyncSectionProtocolsNhrp           types.Bool `tfsdk:"nhrp" vyos:"nhrp,omitempty"`
	LeafServiceConfigSyncSectionProtocolsOspf           types.Bool `tfsdk:"ospf" vyos:"ospf,omitempty"`
	LeafServiceConfigSyncSectionProtocolsOspfvthree     types.Bool `tfsdk:"ospfv3" vyos:"ospfv3,omitempty"`
	LeafServiceConfigSyncSectionProtocolsPim            types.Bool `tfsdk:"pim" vyos:"pim,omitempty"`
	LeafServiceConfigSyncSectionProtocolsPimsix         types.Bool `tfsdk:"pim6" vyos:"pim6,omitempty"`
	LeafServiceConfigSyncSectionProtocolsRIP            types.Bool `tfsdk:"rip" vyos:"rip,omitempty"`
	LeafServiceConfigSyncSectionProtocolsRIPng          types.Bool `tfsdk:"ripng" vyos:"ripng,omitempty"`
	LeafServiceConfigSyncSectionProtocolsRpki           types.Bool `tfsdk:"rpki" vyos:"rpki,omitempty"`
	LeafServiceConfigSyncSectionProtocolsSegmentRouting types.Bool `tfsdk:"segment_routing" vyos:"segment-routing,omitempty"`
	LeafServiceConfigSyncSectionProtocolsStatic         types.Bool `tfsdk:"static" vyos:"static,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceConfigSyncSectionProtocols) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceConfigSyncSectionProtocols) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceConfigSyncSectionProtocols) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceConfigSyncSectionProtocols) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"protocols",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceConfigSyncSectionProtocols) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"service", // Node

		"config-sync", // Node

		"section", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceConfigSyncSectionProtocols) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceConfigSyncSectionProtocols) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"babel":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Babel Routing Protocol

`,
			Description: `Babel Routing Protocol

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"bfd":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Bidirectional Forwarding Detection (BFD)

`,
			Description: `Bidirectional Forwarding Detection (BFD)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"bgp":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Border Gateway Protocol (BGP)

`,
			Description: `Border Gateway Protocol (BGP)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"failover":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Failover route

`,
			Description: `Failover route

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"igmp_proxy":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Internet Group Management Protocol (IGMP) proxy

`,
			Description: `Internet Group Management Protocol (IGMP) proxy

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"isis":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Intermediate System to Intermediate System (IS-IS)

`,
			Description: `Intermediate System to Intermediate System (IS-IS)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mpls":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Multiprotocol Label Switching (MPLS)

`,
			Description: `Multiprotocol Label Switching (MPLS)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"nhrp":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Next Hop Resolution Protocol (NHRP) parameters

`,
			Description: `Next Hop Resolution Protocol (NHRP) parameters

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ospf":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Open Shortest Path First (OSPF)

`,
			Description: `Open Shortest Path First (OSPF)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ospfv3":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Open Shortest Path First (OSPF) for IPv6

`,
			Description: `Open Shortest Path First (OSPF) for IPv6

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"pim":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Protocol Independent Multicast (PIM) and IGMP

`,
			Description: `Protocol Independent Multicast (PIM) and IGMP

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"pim6":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Protocol Independent Multicast for IPv6 (PIMv6) and MLD

`,
			Description: `Protocol Independent Multicast for IPv6 (PIMv6) and MLD

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rip":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Routing Information Protocol (RIP) parameters

`,
			Description: `Routing Information Protocol (RIP) parameters

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"ripng":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Routing Information Protocol (RIPng) parameters

`,
			Description: `Routing Information Protocol (RIPng) parameters

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"rpki":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Resource Public Key Infrastructure (RPKI)

`,
			Description: `Resource Public Key Infrastructure (RPKI)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"segment_routing":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Segment Routing

`,
			Description: `Segment Routing

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"static":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Static Routing

`,
			Description: `Static Routing

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
