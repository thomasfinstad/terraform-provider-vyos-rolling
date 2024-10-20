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

/* tools/generate-terraform-resource-full/templates/resources/global/resource-model.gotmpl */
// Validate compliance
var _ helpers.VyosTopResourceDataModel = &ServiceMDNSRepeater{}

// ServiceMDNSRepeater describes the resource data model.
type ServiceMDNSRepeater struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceMDNSRepeaterDisable      types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafServiceMDNSRepeaterInterface    types.List   `tfsdk:"interface" vyos:"interface,omitempty"`
	LeafServiceMDNSRepeaterIPVersion    types.String `tfsdk:"ip_version" vyos:"ip-version,omitempty"`
	LeafServiceMDNSRepeaterBrowseDomain types.List   `tfsdk:"browse_domain" vyos:"browse-domain,omitempty"`
	LeafServiceMDNSRepeaterAllowService types.List   `tfsdk:"allow_service" vyos:"allow-service,omitempty"`
	LeafServiceMDNSRepeaterVrrpDisable  types.Bool   `tfsdk:"vrrp_disable" vyos:"vrrp-disable,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
}

// SetID configures the resource ID
func (o *ServiceMDNSRepeater) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceMDNSRepeater) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceMDNSRepeater) IsGlobalResource() bool {
	return (true)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceMDNSRepeater) GetVyosPath() []string {
	return append(
		o.GetVyosParentPath(),
		"repeater",
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceMDNSRepeater) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/global/resource-model-parent-vyos-path-hack.gotmpl */
		"service",

		"mdns",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
// ! Since this is a global resource it MUST NOT have a named resource as a parent and should therefore always return an empty string
func (o *ServiceMDNSRepeater) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceMDNSRepeater) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"disable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Description: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"interface":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Interface to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
			Description: `Interface to use

    |  Format  |  Description     |
    |----------|------------------|
    |  txt     |  Interface name  |
`,
		},

		"ip_version":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `IP address version to use

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  _ipv4   |  Use only IPv4 address           |
    |  _ipv6   |  Use only IPv6 address           |
    |  both    |  Use both IPv4 and IPv6 address  |
`,
			Description: `IP address version to use

    |  Format  |  Description                     |
    |----------|----------------------------------|
    |  _ipv4   |  Use only IPv4 address           |
    |  _ipv6   |  Use only IPv6 address           |
    |  both    |  Use both IPv4 and IPv6 address  |
`,

			// Default:          stringdefault.StaticString(`both`),
			Computed: true,
		},

		"browse_domain":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `mDNS browsing domains in addition to the default one

    |  Format  |  Description           |
    |----------|------------------------|
    |  txt     |  mDNS browsing domain  |
`,
			Description: `mDNS browsing domains in addition to the default one

    |  Format  |  Description           |
    |----------|------------------------|
    |  txt     |  mDNS browsing domain  |
`,
		},

		"allow_service":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype-multi.gotmpl */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Allowed mDNS services to be repeated

    |  Format  |  Description   |
    |----------|----------------|
    |  txt     |  mDNS service  |
`,
			Description: `Allowed mDNS services to be repeated

    |  Format  |  Description   |
    |----------|----------------|
    |  txt     |  mDNS service  |
`,
		},

		"vrrp_disable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disables mDNS repeater on VRRP interfaces not in MASTER state

`,
			Description: `Disables mDNS repeater on VRRP interfaces not in MASTER state

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},
	}
}
