// Package resourcemodel code generated by tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

// Validate compliance

var _ helpers.VyosResourceDataModel = &VrfNameProtocolsIsisFastRerouteLfaLocalTiebreaker{}

// VrfNameProtocolsIsisFastRerouteLfaLocalTiebreaker describes the resource data model.
type VrfNameProtocolsIsisFastRerouteLfaLocalTiebreaker struct {
	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)
	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
	NodeVrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstream         *VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstream         `tfsdk:"downstream" vyos:"downstream,omitempty"`
	NodeVrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerLowestBackupMetric *VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerLowestBackupMetric `tfsdk:"lowest_backup_metric" vyos:"lowest-backup-metric,omitempty"`
	NodeVrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtecting     *VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtecting     `tfsdk:"node_protecting" vyos:"node-protecting,omitempty"`
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o VrfNameProtocolsIsisFastRerouteLfaLocalTiebreaker) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		// LeafNodes

		// Nodes

		"downstream": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerDownstream{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Prefer backup path via downstream node

`,
			Description: `Prefer backup path via downstream node

`,
		},

		"lowest_backup_metric": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerLowestBackupMetric{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Prefer backup path with lowest total metric

`,
			Description: `Prefer backup path with lowest total metric

`,
		},

		"node_protecting": schema.SingleNestedAttribute{
			Attributes: VrfNameProtocolsIsisFastRerouteLfaLocalTiebreakerNodeProtecting{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Prefer node protecting backup path

`,
			Description: `Prefer node protecting backup path

`,
		},
	}
}
