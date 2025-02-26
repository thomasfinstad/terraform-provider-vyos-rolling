/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl #package (<no value>) */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl #imports (.) */
import (
	"context"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (cache-peer) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceWebproxyCachePeer{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (cache-peer) */
// ServiceWebproxyCachePeerSelfIdentifier used by TagNodes to keep the id info
type ServiceWebproxyCachePeerSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (cache-peer) */

	ServiceWebproxyCachePeer types.String `tfsdk:"cache_peer" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (webproxy) */

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (service) */
}

// ServiceWebproxyCachePeer describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ServiceWebproxyCachePeer struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *ServiceWebproxyCachePeerSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafServiceWebproxyCachePeerAddress  types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafServiceWebproxyCachePeerHTTPPort types.Number `tfsdk:"http_port" vyos:"http-port,omitempty"`
	LeafServiceWebproxyCachePeerIcpPort  types.Number `tfsdk:"icp_port" vyos:"icp-port,omitempty"`
	LeafServiceWebproxyCachePeerOptions  types.String `tfsdk:"options" vyos:"options,omitempty"`
	LeafServiceWebproxyCachePeerType     types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *ServiceWebproxyCachePeer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceWebproxyCachePeer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceWebproxyCachePeer) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceWebproxyCachePeer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"cache-peer",
		o.SelfIdentifier.ServiceWebproxyCachePeer.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceWebproxyCachePeer) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (webproxy) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (service) */
		"service", // Node

		"webproxy", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceWebproxyCachePeer) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (webproxy) */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (service) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceWebproxyCachePeer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"cache_peer": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Specify other caches in a hierarchy

    |  Format    |  Description       |
    |------------|--------------------|
    |  hostname  |  Cache peers FQDN  |
`,
					Description: `Specify other caches in a hierarchy

    |  Format    |  Description       |
    |------------|--------------------|
    |  hostname  |  Cache peers FQDN  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in cache_peer, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  cache_peer, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (webproxy) */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (service) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"address":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (address) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Hostname or IP address of peer

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  ipv4      |  Squid cache-peer IPv4 address  |
    |  hostname  |  Squid cache-peer hostname      |
`,
			Description: `Hostname or IP address of peer

    |  Format    |  Description                    |
    |------------|---------------------------------|
    |  ipv4      |  Squid cache-peer IPv4 address  |
    |  hostname  |  Squid cache-peer hostname      |
`,
		},

		"http_port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (http-port) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Default Proxy Port

    |  Format      |  Description          |
    |--------------|-----------------------|
    |  1025-65535  |  Default port number  |
`,
			Description: `Default Proxy Port

    |  Format      |  Description          |
    |--------------|-----------------------|
    |  1025-65535  |  Default port number  |
`,

			// Default:          stringdefault.StaticString(`3128`),
			Computed: true,
		},

		"icp_port":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (icp-port) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Cache peer ICP port

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  0        |  Cache peer disabled  |
    |  1-65535  |  Cache peer ICP port  |
`,
			Description: `Cache peer ICP port

    |  Format   |  Description          |
    |-----------|-----------------------|
    |  0        |  Cache peer disabled  |
    |  1-65535  |  Cache peer ICP port  |
`,

			// Default:          stringdefault.StaticString(`0`),
			Computed: true,
		},

		"options":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (options) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Cache peer options

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  Cache peer options  |
`,
			Description: `Cache peer options

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  Cache peer options  |
`,

			// Default:          stringdefault.StaticString(`no-query default`),
			Computed: true,
		},

		"type":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (type) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Squid peer type (default parent)

    |  Format     |  Description                            |
    |-------------|-----------------------------------------|
    |  parent     |  Peer is a parent                       |
    |  sibling    |  Peer is a sibling                      |
    |  multicast  |  Peer is a member of a multicast group  |
`,
			Description: `Squid peer type (default parent)

    |  Format     |  Description                            |
    |-------------|-----------------------------------------|
    |  parent     |  Peer is a parent                       |
    |  sibling    |  Peer is a sibling                      |
    |  multicast  |  Peer is a member of a multicast group  |
`,

			// Default:          stringdefault.StaticString(`parent`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
