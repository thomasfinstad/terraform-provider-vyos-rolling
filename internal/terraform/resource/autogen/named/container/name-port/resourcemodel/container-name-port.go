// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
)

// ContainerNamePort describes the resource data model.
type ContainerNamePort struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.String `tfsdk:"port_id" vyos:"-,self-id"`

	ParentIDContainerName types.String `tfsdk:"name_id" vyos:"name,parent-id"`

	// LeafNodes
	LeafContainerNamePortListenAddress types.List   `tfsdk:"listen_address" vyos:"listen-address,omitempty"`
	LeafContainerNamePortSource        types.String `tfsdk:"source" vyos:"source,omitempty"`
	LeafContainerNamePortDestination   types.String `tfsdk:"destination" vyos:"destination,omitempty"`
	LeafContainerNamePortProtocol      types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// SetID configures the resource ID
func (o *ContainerNamePort) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ContainerNamePort) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"container",

		"name",
		o.ParentIDContainerName.ValueString(),

		"port",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerNamePort) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},
		"port_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Publish port to the container

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			}, Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in port_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  port_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		"name_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Container name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Validators: []validator.String{
				stringvalidator.All(
					helpers.StringNot(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^.*__.*$`),
							"double underscores in name_id, conflicts with the internal resource id",
						),
					),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
						"illigal character in  name_id, value must match: ^[a-zA-Z0-9-_]*$",
					),
				),
			},
		},

		// LeafNodes

		"listen_address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Local IP addresses to listen on

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address to listen for incoming connections  |
    |  ipv6  &emsp; |  IPv6 address to listen for incoming connections  |

`,
		},

		"source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source host port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Source host port  |
    |  start-end  &emsp; |  Source host port range (e.g. 10025-10030)  |

`,
		},

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Destination container port

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Destination container port  |
    |  start-end  &emsp; |  Destination container port range (e.g. 10025-10030)  |

`,
		},

		"protocol": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Transport protocol used for port mapping

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  tcp  &emsp; |  Use Transmission Control Protocol for given port  |
    |  udp  &emsp; |  Use User Datagram Protocol for given port  |

`,

			// Default:          stringdefault.StaticString(`tcp`),
			Computed: true,
		},

		// Nodes

	}
}
