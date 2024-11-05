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

var _ helpers.VyosTopResourceDataModel = &ServiceConfigSyncSection{}

// ServiceConfigSyncSection describes the resource data model.
// This is a basenode!
// Top level basenode type: `Node`
type ServiceConfigSyncSection struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceConfigSyncSectionFirewall  types.Bool `tfsdk:"firewall" vyos:"firewall,omitempty"`
	LeafServiceConfigSyncSectionNat       types.Bool `tfsdk:"nat" vyos:"nat,omitempty"`
	LeafServiceConfigSyncSectionNatsixsix types.Bool `tfsdk:"nat66" vyos:"nat66,omitempty"`
	LeafServiceConfigSyncSectionPki       types.Bool `tfsdk:"pki" vyos:"pki,omitempty"`
	LeafServiceConfigSyncSectionPolicy    types.Bool `tfsdk:"policy" vyos:"policy,omitempty"`
	LeafServiceConfigSyncSectionVpn       types.Bool `tfsdk:"vpn" vyos:"vpn,omitempty"`
	LeafServiceConfigSyncSectionVrf       types.Bool `tfsdk:"vrf" vyos:"vrf,omitempty"`

	// TagNodes

	// Nodes

	ExistsNodeServiceConfigSyncSectionInterfaces bool `tfsdk:"-" vyos:"interfaces,child"`

	ExistsNodeServiceConfigSyncSectionProtocols bool `tfsdk:"-" vyos:"protocols,child"`

	ExistsNodeServiceConfigSyncSectionQos bool `tfsdk:"-" vyos:"qos,child"`

	ExistsNodeServiceConfigSyncSectionService bool `tfsdk:"-" vyos:"service,child"`

	ExistsNodeServiceConfigSyncSectionSystem bool `tfsdk:"-" vyos:"system,child"`
}

// SetID configures the resource ID
func (o *ServiceConfigSyncSection) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceConfigSyncSection) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceConfigSyncSection) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceConfigSyncSection) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"section",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceConfigSyncSection) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"service", // Node

		"config-sync", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceConfigSyncSection) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceConfigSyncSection) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"firewall":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Firewall

`,
			Description: `Firewall

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"nat":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `NAT

`,
			Description: `NAT

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"nat66":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `NAT66

`,
			Description: `NAT66

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"pki":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Public key infrastructure (PKI)

`,
			Description: `Public key infrastructure (PKI)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"policy":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Routing policy

`,
			Description: `Routing policy

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"vpn":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Private Network (VPN)

`,
			Description: `Virtual Private Network (VPN)

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"vrf":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Routing and Forwarding

`,
			Description: `Virtual Routing and Forwarding

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
