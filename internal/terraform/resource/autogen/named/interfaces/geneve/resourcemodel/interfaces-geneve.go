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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl #resource-model (geneve) */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &InterfacesGeneve{}

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-tag-node-identifier (geneve) */
// InterfacesGeneveSelfIdentifier used by TagNodes to keep the id info
type InterfacesGeneveSelfIdentifier struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (geneve) */

	InterfacesGeneve types.String `tfsdk:"geneve" vyos:"-,self-id"`

	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl #resource-model-parent-id-hack  (interfaces) */
}

// InterfacesGeneve describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type InterfacesGeneve struct {
	ID       types.String   `tfsdk:"id" vyos:"-,tfsdk-id"`
	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	SelfIdentifier *InterfacesGeneveSelfIdentifier `tfsdk:"identifier" vyos:"-,self-id"`

	// LeafNodes
	LeafInterfacesGeneveAddress     types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafInterfacesGeneveDescrIPtion types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafInterfacesGeneveDisable     types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafInterfacesGeneveMac         types.String `tfsdk:"mac" vyos:"mac,omitempty"`
	LeafInterfacesGeneveMtu         types.Number `tfsdk:"mtu" vyos:"mtu,omitempty"`
	LeafInterfacesGeneveRedirect    types.String `tfsdk:"redirect" vyos:"redirect,omitempty"`
	LeafInterfacesGeneveRemote      types.String `tfsdk:"remote" vyos:"remote,omitempty"`
	LeafInterfacesGeneveVrf         types.String `tfsdk:"vrf" vyos:"vrf,omitempty"`
	LeafInterfacesGeneveVni         types.Number `tfsdk:"vni" vyos:"vni,omitempty"`

	// TagNodes

	// Nodes

	NodeInterfacesGeneveIP *InterfacesGeneveIP `tfsdk:"ip" vyos:"ip,omitempty"`

	NodeInterfacesGeneveIPvsix *InterfacesGeneveIPvsix `tfsdk:"ipv6" vyos:"ipv6,omitempty"`

	NodeInterfacesGeneveParameters *InterfacesGeneveParameters `tfsdk:"parameters" vyos:"parameters,omitempty"`

	NodeInterfacesGeneveMirror *InterfacesGeneveMirror `tfsdk:"mirror" vyos:"mirror,omitempty"`
}

// SetID configures the resource ID
func (o *InterfacesGeneve) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *InterfacesGeneve) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *InterfacesGeneve) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *InterfacesGeneve) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"geneve",
		o.SelfIdentifier.InterfacesGeneve.ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *InterfacesGeneve) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack (interfaces) */
		"interfaces", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *InterfacesGeneve) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global (interfaces) */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o InterfacesGeneve) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"geneve": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Generic Network Virtualization Encapsulation (GENEVE) Interface

    |  Format  |  Description            |
    |----------|-------------------------|
    |  gnvN    |  GENEVE interface name  |
`,
					Description: `Generic Network Virtualization Encapsulation (GENEVE) Interface

    |  Format  |  Description            |
    |----------|-------------------------|
    |  gnvN    |  GENEVE interface name  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in geneve, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  geneve, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl #resource-model-parent-schema-hack (interfaces) */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"address":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi (address) */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `IP address

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
    |  ipv6net  |  IPv6 address and prefix length  |
`,
			Description: `IP address

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
    |  ipv6net  |  IPv6 address and prefix length  |
`,
		},

		"description":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (description) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
			Description: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt     |  Description  |
`,
		},

		"disable":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (disable) */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Administratively disable interface

`,
			Description: `Administratively disable interface

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"mac":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mac) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
			Description: `Media Access Control (MAC) address

    |  Format   |  Description             |
    |-----------|--------------------------|
    |  macaddr  |  Hardware (MAC) address  |
`,
		},

		"mtu":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (mtu) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Maximum Transmission Unit (MTU)

    |  Format      |  Description                        |
    |--------------|-------------------------------------|
    |  1200-16000  |  Maximum Transmission Unit in byte  |
`,
			Description: `Maximum Transmission Unit (MTU)

    |  Format      |  Description                        |
    |--------------|-------------------------------------|
    |  1200-16000  |  Maximum Transmission Unit in byte  |
`,

			// Default:          stringdefault.StaticString(`1500`),
			Computed: true,
		},

		"redirect":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (redirect) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Redirect incoming packet to destination

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
			Description: `Redirect incoming packet to destination

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  txt     |  Destination interface name  |
`,
		},

		"remote":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (remote) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Tunnel remote address

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  Tunnel remote IPv4 address  |
    |  ipv6    |  Tunnel remote IPv6 address  |
`,
			Description: `Tunnel remote address

    |  Format  |  Description                 |
    |----------|------------------------------|
    |  ipv4    |  Tunnel remote IPv4 address  |
    |  ipv6    |  Tunnel remote IPv6 address  |
`,
		},

		"vrf":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (vrf) */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
			Description: `VRF instance name

    |  Format  |  Description        |
    |----------|---------------------|
    |  txt     |  VRF instance name  |
`,
		},

		"vni":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype (vni) */
		schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Virtual Network Identifier

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  0-16777214  |  VXLAN virtual network identifier  |
`,
			Description: `Virtual Network Identifier

    |  Format      |  Description                       |
    |--------------|------------------------------------|
    |  0-16777214  |  VXLAN virtual network identifier  |
`,
		},

		// TagNodes

		// Nodes

		"ip": schema.SingleNestedAttribute{
			Attributes: InterfacesGeneveIP{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv4 routing parameters

`,
			Description: `IPv4 routing parameters

`,
		},

		"ipv6": schema.SingleNestedAttribute{
			Attributes: InterfacesGeneveIPvsix{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `IPv6 routing parameters

`,
			Description: `IPv6 routing parameters

`,
		},

		"parameters": schema.SingleNestedAttribute{
			Attributes: InterfacesGeneveParameters{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `GENEVE tunnel parameters

`,
			Description: `GENEVE tunnel parameters

`,
		},

		"mirror": schema.SingleNestedAttribute{
			Attributes: InterfacesGeneveMirror{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Mirror ingress/egress packets

`,
			Description: `Mirror ingress/egress packets

`,
		},
	}
}
