/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosResourceDataModel = &InterfacesEthernetVifSVifCIP{}

// InterfacesEthernetVifSVifCIP describes the resource data model.
// This is not a basenode!
// Top level basenode type: `N/A`
type InterfacesEthernetVifSVifCIP struct {
	// LeafNodes
	LeafInterfacesEthernetVifSVifCIPAdjustMss               types.String `tfsdk:"adjust_mss" vyos:"adjust-mss,omitempty"`
	LeafInterfacesEthernetVifSVifCIPArpCacheTimeout         types.Number `tfsdk:"arp_cache_timeout" vyos:"arp-cache-timeout,omitempty"`
	LeafInterfacesEthernetVifSVifCIPDisableArpFilter        types.Bool   `tfsdk:"disable_arp_filter" vyos:"disable-arp-filter,omitempty"`
	LeafInterfacesEthernetVifSVifCIPDisableForwarding       types.Bool   `tfsdk:"disable_forwarding" vyos:"disable-forwarding,omitempty"`
	LeafInterfacesEthernetVifSVifCIPEnableDirectedBroadcast types.Bool   `tfsdk:"enable_directed_broadcast" vyos:"enable-directed-broadcast,omitempty"`
	LeafInterfacesEthernetVifSVifCIPEnableArpAccept         types.Bool   `tfsdk:"enable_arp_accept" vyos:"enable-arp-accept,omitempty"`
	LeafInterfacesEthernetVifSVifCIPEnableArpAnnounce       types.Bool   `tfsdk:"enable_arp_announce" vyos:"enable-arp-announce,omitempty"`
	LeafInterfacesEthernetVifSVifCIPEnableArpIgnore         types.Bool   `tfsdk:"enable_arp_ignore" vyos:"enable-arp-ignore,omitempty"`
	LeafInterfacesEthernetVifSVifCIPEnableProxyArp          types.Bool   `tfsdk:"enable_proxy_arp" vyos:"enable-proxy-arp,omitempty"`
	LeafInterfacesEthernetVifSVifCIPProxyArpPvlan           types.Bool   `tfsdk:"proxy_arp_pvlan" vyos:"proxy-arp-pvlan,omitempty"`
	LeafInterfacesEthernetVifSVifCIPSourceValIDation        types.String `tfsdk:"source_validation" vyos:"source-validation,omitempty"`

	// TagNodes

	// Nodes
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesEthernetVifSVifCIP) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		"adjust_mss":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Adjust TCP MSS value

    |  Format             |  Description                                     |
    |---------------------|--------------------------------------------------|
    |  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
    |  536-65535          |  TCP Maximum segment size in bytes               |
`,
			Description: `Adjust TCP MSS value

    |  Format             |  Description                                     |
    |---------------------|--------------------------------------------------|
    |  clamp-mss-to-pmtu  |  Automatically sets the MSS to the proper value  |
    |  536-65535          |  TCP Maximum segment size in bytes               |
`,
		},

		"arp_cache_timeout":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `ARP cache entry timeout in seconds

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  1-86400  |  ARP cache entry timout in seconds  |
`,
			Description: `ARP cache entry timeout in seconds

    |  Format   |  Description                        |
    |-----------|-------------------------------------|
    |  1-86400  |  ARP cache entry timout in seconds  |
`,

			// Default:          stringdefault.StaticString(`30`),
			Computed: true,
		},

		"disable_arp_filter":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable ARP filter on this interface

`,
			Description: `Disable ARP filter on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"disable_forwarding":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable IP forwarding on this interface

`,
			Description: `Disable IP forwarding on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_directed_broadcast":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable directed broadcast forwarding on this interface

`,
			Description: `Enable directed broadcast forwarding on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_arp_accept":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable ARP accept on this interface

`,
			Description: `Enable ARP accept on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_arp_announce":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable ARP announce on this interface

`,
			Description: `Enable ARP announce on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_arp_ignore":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable ARP ignore on this interface

`,
			Description: `Enable ARP ignore on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"enable_proxy_arp":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable proxy-arp on this interface

`,
			Description: `Enable proxy-arp on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"proxy_arp_pvlan":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Enable private VLAN proxy ARP on this interface

`,
			Description: `Enable private VLAN proxy ARP on this interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"source_validation":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source validation by reversed path (RFC3704)

    |  Format   |  Description                                                  |
    |-----------|---------------------------------------------------------------|
    |  strict   |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose    |  Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    |  disable  |  No source validation                                         |
`,
			Description: `Source validation by reversed path (RFC3704)

    |  Format   |  Description                                                  |
    |-----------|---------------------------------------------------------------|
    |  strict   |  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose    |  Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    |  disable  |  No source validation                                         |
`,
		},

		// TagNodes

		// Nodes

	}
}
