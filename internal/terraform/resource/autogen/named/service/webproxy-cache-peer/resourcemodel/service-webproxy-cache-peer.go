// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

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

// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceWebproxyCachePeer{}

// ServiceWebproxyCachePeer describes the resource data model.
type ServiceWebproxyCachePeer struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceWebproxyCachePeerAddress  types.String `tfsdk:"address" vyos:"address,omitempty"`
	LeafServiceWebproxyCachePeerHTTPPort types.Number `tfsdk:"http_port" vyos:"http-port,omitempty"`
	LeafServiceWebproxyCachePeerIcpPort  types.Number `tfsdk:"icp_port" vyos:"icp-port,omitempty"`
	LeafServiceWebproxyCachePeerOptions  types.String `tfsdk:"options" vyos:"options,omitempty"`
	LeafServiceWebproxyCachePeerType     types.String `tfsdk:"type" vyos:"type,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

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
		o.SelfIdentifier.Attributes()["cache_peer"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceWebproxyCachePeer) GetVyosParentPath() []string {
	return []string{
		"service",

		"webproxy",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceWebproxyCachePeer) GetVyosNamedParentPath() []string {
	return []string{}
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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  cache_peer, value must match: ^[.:a-zA-Z0-9-_]+$",
							),
						),
					},
				},
			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"address": schema.StringAttribute{
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

		"http_port": schema.NumberAttribute{
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

		"icp_port": schema.NumberAttribute{
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

		"options": schema.StringAttribute{
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

		"type": schema.StringAttribute{
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

		// Nodes

	}
}
