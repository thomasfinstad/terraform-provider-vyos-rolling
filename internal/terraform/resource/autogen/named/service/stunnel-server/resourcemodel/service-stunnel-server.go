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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers"
)

/* tools/generate-terraform-resource-full/templates/resources/named/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &ServiceStunnelServer{}

// ServiceStunnelServer describes the resource data model.
type ServiceStunnelServer struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl */
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafServiceStunnelServerProtocol types.String `tfsdk:"protocol" vyos:"protocol,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	ExistsTagServiceStunnelServerPsk bool `tfsdk:"-" vyos:"psk,child"`

	// Nodes
	NodeServiceStunnelServerConnect *ServiceStunnelServerConnect `tfsdk:"connect" vyos:"connect,omitempty"`
	NodeServiceStunnelServerListen  *ServiceStunnelServerListen  `tfsdk:"listen" vyos:"listen,omitempty"`
	NodeServiceStunnelServerSsl     *ServiceStunnelServerSsl     `tfsdk:"ssl" vyos:"ssl,omitempty"`
}

// SetID configures the resource ID
func (o *ServiceStunnelServer) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ServiceStunnelServer) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ServiceStunnelServer) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceStunnelServer) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"server",
		o.SelfIdentifier.Attributes()["server"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ServiceStunnelServer) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-vyos-path-hack.gotmpl */
		"service",

		"stunnel",
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ServiceStunnelServer) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack-for-non-global.gotmpl */

	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceStunnelServer) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"server": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Stunnel server config

`,
					Description: `Stunnel server config

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in server, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  server, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"protocol":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Application protocol to negotiate TLS

    |  Format  |  Description                                                                      |
    |----------|-----------------------------------------------------------------------------------|
    |  cifs    |  Proprietary (undocummented) extension of CIFS protocol                           |
    |  imap    |  Based on RFC 2595 - Using TLS with IMAP, POP3 and ACAP                           |
    |  pgsql   |  Based on PostgreSQL frontend/backend protocol                                    |
    |  pop3    |  Based on RFC 2449 - POP3 Extension Mechanism                                     |
    |  proxy   |  Passing of the original client IP address with HAProxy PROXY protocol version 1  |
    |  smtp    |  Based on RFC 2487 - SMTP Service Extension for Secure SMTP over TLS              |
    |  socks   |  SOCKS versions 4, 4a, and 5 are supported                                        |
`,
			Description: `Application protocol to negotiate TLS

    |  Format  |  Description                                                                      |
    |----------|-----------------------------------------------------------------------------------|
    |  cifs    |  Proprietary (undocummented) extension of CIFS protocol                           |
    |  imap    |  Based on RFC 2595 - Using TLS with IMAP, POP3 and ACAP                           |
    |  pgsql   |  Based on PostgreSQL frontend/backend protocol                                    |
    |  pop3    |  Based on RFC 2449 - POP3 Extension Mechanism                                     |
    |  proxy   |  Passing of the original client IP address with HAProxy PROXY protocol version 1  |
    |  smtp    |  Based on RFC 2487 - SMTP Service Extension for Secure SMTP over TLS              |
    |  socks   |  SOCKS versions 4, 4a, and 5 are supported                                        |
`,
		},

		// Nodes

		"connect": schema.SingleNestedAttribute{
			Attributes: ServiceStunnelServerConnect{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Connect to a remote address

`,
			Description: `Connect to a remote address

`,
		},

		"listen": schema.SingleNestedAttribute{
			Attributes: ServiceStunnelServerListen{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `Accept connections on specified address

`,
			Description: `Accept connections on specified address

`,
		},

		"ssl": schema.SingleNestedAttribute{
			Attributes: ServiceStunnelServerSsl{}.ResourceSchemaAttributes(ctx),
			Optional:   true,
			MarkdownDescription: `SSL Certificate, SSL Key and CA

`,
			Description: `SSL Certificate, SSL Key and CA

`,
		},
	}
}
