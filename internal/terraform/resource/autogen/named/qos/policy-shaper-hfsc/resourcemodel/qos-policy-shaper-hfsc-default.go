// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &QosPolicyShaperHfscDefault{}

// QosPolicyShaperHfscDefault describes the resource data model.
type QosPolicyShaperHfscDefault struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeQosPolicyShaperHfscDefaultLinkshare  *QosPolicyShaperHfscDefaultLinkshare  `tfsdk:"linkshare" vyos:"linkshare,omitempty"`
	NodeQosPolicyShaperHfscDefaultRealtime   *QosPolicyShaperHfscDefaultRealtime   `tfsdk:"realtime" vyos:"realtime,omitempty"`
	NodeQosPolicyShaperHfscDefaultUpperlimit *QosPolicyShaperHfscDefaultUpperlimit `tfsdk:"upperlimit" vyos:"upperlimit,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o QosPolicyShaperHfscDefault) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"linkshare": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscDefaultLinkshare{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Linkshare class settings

`,
			Description: `Linkshare class settings

`,
		},

		"realtime": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscDefaultRealtime{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Realtime class settings

`,
			Description: `Realtime class settings

`,
		},

		"upperlimit": schema.SingleNestedAttribute{
			Attributes: QosPolicyShaperHfscDefaultUpperlimit{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Upperlimit class settings

`,
			Description: `Upperlimit class settings

`,
		},
	}
}
