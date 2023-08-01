// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceIPoeServerClientIPvsixPoolDelegate describes the resource data model.
type ServiceIPoeServerClientIPvsixPoolDelegate struct {
	SelfIdentifier types.String `tfsdk:"delegate_id" vyos:",self-id"`

	// LeafNodes
	LeafServiceIPoeServerClientIPvsixPoolDelegateDelegationPrefix types.Number `tfsdk:"delegation_prefix" vyos:"delegation-prefix,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceIPoeServerClientIPvsixPoolDelegate) GetVyosPath() []string {
	return []string{
		"service",

		"ipoe-server",

		"client-ipv6-pool",

		"delegate",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceIPoeServerClientIPvsixPoolDelegate) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"delegate_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Subnet used to delegate prefix through DHCPv6-PD (RFC3633)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |

`,
		},

		// LeafNodes

		"delegation_prefix": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Prefix length delegated to client

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 32-64  &emsp; |  Delegated prefix length  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceIPoeServerClientIPvsixPoolDelegate) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceIPoeServerClientIPvsixPoolDelegate) UnmarshalJSON(_ []byte) error {
	return nil
}