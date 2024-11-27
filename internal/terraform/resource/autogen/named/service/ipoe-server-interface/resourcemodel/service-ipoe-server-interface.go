/* tools/generate-terraform-resource-full/templates/resources/common/package.gotmpl */
// Package resourcemodel code generated by tools/generate-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

/* tools/generate-terraform-resource-full/templates/resources/common/imports.gotmpl */
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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceIPoeServerInterface{}

// ServiceIPoeServerInterface describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type ServiceIPoeServerInterface struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl */
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceIPoeServerInterfaceMode         types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafServiceIPoeServerInterfaceNetwork      types.String `tfsdk:"network" vyos:"network,omitempty"`
	LeafServiceIPoeServerInterfaceClientSubnet types.String `tfsdk:"client_subnet" vyos:"client-subnet,omitempty"`
	LeafServiceIPoeServerInterfaceLuaUsername  types.String `tfsdk:"lua_username" vyos:"lua-username,omitempty"`
	LeafServiceIPoeServerInterfaceVlan         types.List   `tfsdk:"vlan" vyos:"vlan,omitempty"`
	LeafServiceIPoeServerInterfaceVlanMon      types.Bool   `tfsdk:"vlan_mon" vyos:"vlan-mon,omitempty"`

	// TagNodes

	// Nodes

	NodeServiceIPoeServerInterfaceExternalDhcp *ServiceIPoeServerInterfaceExternalDhcp `tfsdk:"external_dhcp" vyos:"external-dhcp,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceIPoeServerInterface) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceIPoeServerInterface) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceIPoeServerInterface) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerInterface) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"interface",
		o.SelfIdentifier.Attributes()["interface"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceIPoeServerInterface) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"service", // Node

		"ipoe-server", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceIPoeServerInterface) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerInterface) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"interface": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Interface to listen dhcp or unclassified packets

`,
					Description: `Interface to listen dhcp or unclassified packets

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in interface, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  interface, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

			},
		},

		"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
			Create: true,
		}),

		// LeafNodes

		"mode":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client connectivity mode

    |  Format  |  Description                                 |
    |----------|----------------------------------------------|
    |  l2      |  Client located on same interface as server  |
    |  l3      |  Client located behind a router              |
`,
			Description: `Client connectivity mode

    |  Format  |  Description                                 |
    |----------|----------------------------------------------|
    |  l2      |  Client located on same interface as server  |
    |  l3      |  Client located behind a router              |
`,

			// Default:          stringdefault.StaticString(`l2`),
			Computed: true,
		},

		"network":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Enables clients to share the same network or each client has its own vlan

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  shared  |  Multiple clients share the same network  |
    |  vlan    |  One VLAN per client                      |
`,
			Description: `Enables clients to share the same network or each client has its own vlan

    |  Format  |  Description                              |
    |----------|-------------------------------------------|
    |  shared  |  Multiple clients share the same network  |
    |  vlan    |  One VLAN per client                      |
`,

			// Default:          stringdefault.StaticString(`shared`),
			Computed: true,
		},

		"client_subnet":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Client address pool

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
`,
			Description: `Client address pool

    |  Format   |  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  |  IPv4 address and prefix length  |
`,
		},

		"lua_username":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Username function

    |  Format  |  Description                                                       |
    |----------|--------------------------------------------------------------------|
    |  txt     |  Name of the function in the Lua file to construct usernames with  |
`,
			Description: `Username function

    |  Format  |  Description                                                       |
    |----------|--------------------------------------------------------------------|
    |  txt     |  Name of the function in the Lua file to construct usernames with  |
`,
		},

		"vlan":
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype-multi */
		schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `VLAN monitor for automatic creation of VLAN interfaces

    |  Format     |  Description                                      |
    |-------------|---------------------------------------------------|
    |  1-4094     |  VLAN for automatic creation                      |
    |  start-end  |  VLAN range for automatic creation (e.g. 1-4094)  |
`,
			Description: `VLAN monitor for automatic creation of VLAN interfaces

    |  Format     |  Description                                      |
    |-------------|---------------------------------------------------|
    |  1-4094     |  VLAN for automatic creation                      |
    |  start-end  |  VLAN range for automatic creation (e.g. 1-4094)  |
`,
		},

		"vlan_mon":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Automatically create VLAN interfaces

`,
			Description: `Automatically create VLAN interfaces

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		// TagNodes

		// Nodes

		"external_dhcp": schema.SingleNestedAttribute{
			Attributes: ServiceIPoeServerInterfaceExternalDhcp{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `DHCP requests will be forwarded

`,
			Description: `DHCP requests will be forwarded

`,
		},
	}
}
