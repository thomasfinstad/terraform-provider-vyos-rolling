// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ServiceHTTPSCertificates describes the resource data model.
type ServiceHTTPSCertificates struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	// LeafNodes
	LeafServiceHTTPSCertificatesCaCertificate types.String `tfsdk:"ca_certificate" vyos:"ca-certificate,omitempty"`
	LeafServiceHTTPSCertificatesCertificate   types.String `tfsdk:"certificate" vyos:"certificate,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes (Bools that show if child resources have been configured)
	ExistsNodeServiceHTTPSCertificatesCertbot bool `tfsdk:"-" vyos:"certbot,child"`
}

// SetID configures the resource ID
func (o *ServiceHTTPSCertificates) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceHTTPSCertificates) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return []string{
		"service",

		"https",

		"certificates",
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceHTTPSCertificates) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).",
		},

		// LeafNodes

		"ca_certificate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate Authority in PKI configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of CA in PKI configuration  |

`,
		},

		"certificate": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Certificate in PKI configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of certificate in PKI configuration  |

`,
		},
	}
}
