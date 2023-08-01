// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceWebproxyURLFilteringSquIDguardSourceGroup describes the resource data model.
type ServiceWebproxyURLFilteringSquIDguardSourceGroup struct {
	SelfIdentifier types.String `tfsdk:"source_group_id" vyos:",self-id"`

	// LeafNodes
	LeafServiceWebproxyURLFilteringSquIDguardSourceGroupAddress        types.List   `tfsdk:"address" vyos:"address,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardSourceGroupDescrIPtion    types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardSourceGroupDomain         types.List   `tfsdk:"domain" vyos:"domain,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardSourceGroupLdapIPSearch   types.List   `tfsdk:"ldap_ip_search" vyos:"ldap-ip-search,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardSourceGroupLdapUserSearch types.List   `tfsdk:"ldap_user_search" vyos:"ldap-user-search,omitempty"`
	LeafServiceWebproxyURLFilteringSquIDguardSourceGroupUser           types.String `tfsdk:"user" vyos:"user,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceWebproxyURLFilteringSquIDguardSourceGroup) GetVyosPath() []string {
	return []string{
		"service",

		"webproxy",

		"url-filtering",

		"squidguard",

		"source-group",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceWebproxyURLFilteringSquIDguardSourceGroup) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"source_group_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Source group name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  name  &emsp; |  Name of source group  |

`,
		},

		// LeafNodes

		"address": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Address for source-group

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address to match  |
    |  ipv4net  &emsp; |  IPv4 prefix to match  |
    |  ipv4range  &emsp; |  IPv4 address range to match  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description for source-group

`,
		},

		"domain": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Domain for source-group

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  domain  &emsp; |  Domain name for the source-group  |

`,
		},

		"ldap_ip_search": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `LDAP search expression for an IP address list

`,
		},

		"ldap_user_search": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `LDAP search expression for a user group

`,
		},

		"user": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `List of user names

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceWebproxyURLFilteringSquIDguardSourceGroup) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceWebproxyURLFilteringSquIDguardSourceGroup) UnmarshalJSON(_ []byte) error {
	return nil
}