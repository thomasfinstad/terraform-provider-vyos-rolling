// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceRouterAdvertInterfacePrefix describes the resource data model.
type ServiceRouterAdvertInterfacePrefix struct {
	SelfIdentifier types.String `tfsdk:"prefix_id" vyos:",self-id"`

	ParentIDServiceRouterAdvertInterface types.String `tfsdk:"interface" vyos:"interface,parent-id"`

	// LeafNodes
	LeafServiceRouterAdvertInterfacePrefixNoAutonomousFlag  types.Bool   `tfsdk:"no_autonomous_flag" vyos:"no-autonomous-flag,omitempty"`
	LeafServiceRouterAdvertInterfacePrefixNoOnLinkFlag      types.Bool   `tfsdk:"no_on_link_flag" vyos:"no-on-link-flag,omitempty"`
	LeafServiceRouterAdvertInterfacePrefixDeprecatePrefix   types.Bool   `tfsdk:"deprecate_prefix" vyos:"deprecate-prefix,omitempty"`
	LeafServiceRouterAdvertInterfacePrefixDecrementLifetime types.Bool   `tfsdk:"decrement_lifetime" vyos:"decrement-lifetime,omitempty"`
	LeafServiceRouterAdvertInterfacePrefixPreferredLifetime types.String `tfsdk:"preferred_lifetime" vyos:"preferred-lifetime,omitempty"`
	LeafServiceRouterAdvertInterfacePrefixValIDLifetime     types.String `tfsdk:"valid_lifetime" vyos:"valid-lifetime,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceRouterAdvertInterfacePrefix) GetVyosPath() []string {
	return []string{
		"service",

		"router-advert",

		"interface",
		o.ParentIDServiceRouterAdvertInterface.ValueString(),

		"prefix",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceRouterAdvertInterfacePrefix) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"prefix_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `IPv6 prefix to be advertised in Router Advertisements (RAs)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  IPv6 prefix to be advertized  |

`,
		},

		"interface_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Interface to send RA on

`,
		},

		// LeafNodes

		"no_autonomous_flag": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Prefix can not be used for stateless address auto-configuration

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"no_on_link_flag": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Prefix can not be used for on-link determination

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"deprecate_prefix": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Upon shutdown, this option will deprecate the prefix by announcing it in the shutdown RA

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"decrement_lifetime": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Lifetime is decremented by the number of seconds since the last RA - use in conjunction with a DHCPv6-PD prefix

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"preferred_lifetime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time in seconds that the prefix will remain preferred

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  u32  &emsp; |  Time in seconds that the prefix will remain preferred  |
    |  infinity  &emsp; |  Prefix will remain preferred forever  |

`,

			// Default:          stringdefault.StaticString(`14400`),
			Computed: true,
		},

		"valid_lifetime": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Time in seconds that the prefix will remain valid

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967295  &emsp; |  Time in seconds that the prefix will remain valid  |
    |  infinity  &emsp; |  Prefix will remain preferred forever  |

`,

			// Default:          stringdefault.StaticString(`2592000`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceRouterAdvertInterfacePrefix) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceRouterAdvertInterfacePrefix) UnmarshalJSON(_ []byte) error {
	return nil
}