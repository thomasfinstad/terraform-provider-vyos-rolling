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

/* tools/generate-terraform-resource-full/templates/resources/common/resource-model.gotmpl */
// Validate compliance

var _ helpers.VyosTopResourceDataModel = &LoadBalancingHaproxyServiceLoggingFacility{}

// LoadBalancingHaproxyServiceLoggingFacility describes the resource data model.
// This is a basenode!
// Top level basenode type: `TagNode`
type LoadBalancingHaproxyServiceLoggingFacility struct {
	/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-tag-node-identifier.gotmpl */
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafLoadBalancingHaproxyServiceLoggingFacilityLevel types.String `tfsdk:"level" vyos:"level,omitempty"`

	// TagNodes

	// Nodes
}

// SetID configures the resource ID
func (o *LoadBalancingHaproxyServiceLoggingFacility) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *LoadBalancingHaproxyServiceLoggingFacility) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *LoadBalancingHaproxyServiceLoggingFacility) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *LoadBalancingHaproxyServiceLoggingFacility) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"facility",
		o.SelfIdentifier.Attributes()["facility"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *LoadBalancingHaproxyServiceLoggingFacility) GetVyosParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"load-balancing", // Node

		"haproxy", // Node

		"service",
		o.SelfIdentifier.Attributes()["service"].(types.String).ValueString(),

		"logging", // Node

	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *LoadBalancingHaproxyServiceLoggingFacility) GetVyosNamedParentPath() []string {
	return []string{
		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack-for-non-global */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-parent-vyos-path-hack.gotmpl #resource-model-parent-vyos-path-hack */
		"load-balancing", // Node

		"haproxy", // Node

		"service",
		o.SelfIdentifier.Attributes()["service"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o LoadBalancingHaproxyServiceLoggingFacility) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"facility": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Facility for logging

    |  Format    |  Description                       |
    |------------|------------------------------------|
    |  all       |  All facilities excluding "mark"   |
    |  auth      |  Authentication and authorization  |
    |  authpriv  |  Non-system authorization          |
    |  cron      |  Cron daemon                       |
    |  daemon    |  System daemons                    |
    |  kern      |  Kernel                            |
    |  lpr       |  Line printer spooler              |
    |  mail      |  Mail subsystem                    |
    |  mark      |  Timestamp                         |
    |  news      |  USENET subsystem                  |
    |  syslog    |  Authentication and authorization  |
    |  user      |  Application processes             |
    |  uucp      |  UUCP subsystem                    |
    |  local0    |  Local facility 0                  |
    |  local1    |  Local facility 1                  |
    |  local2    |  Local facility 2                  |
    |  local3    |  Local facility 3                  |
    |  local4    |  Local facility 4                  |
    |  local5    |  Local facility 5                  |
    |  local6    |  Local facility 6                  |
    |  local7    |  Local facility 7                  |
`,
					Description: `Facility for logging

    |  Format    |  Description                       |
    |------------|------------------------------------|
    |  all       |  All facilities excluding "mark"   |
    |  auth      |  Authentication and authorization  |
    |  authpriv  |  Non-system authorization          |
    |  cron      |  Cron daemon                       |
    |  daemon    |  System daemons                    |
    |  kern      |  Kernel                            |
    |  lpr       |  Line printer spooler              |
    |  mail      |  Mail subsystem                    |
    |  mark      |  Timestamp                         |
    |  news      |  USENET subsystem                  |
    |  syslog    |  Authentication and authorization  |
    |  user      |  Application processes             |
    |  uucp      |  UUCP subsystem                    |
    |  local0    |  Local facility 0                  |
    |  local1    |  Local facility 1                  |
    |  local2    |  Local facility 2                  |
    |  local3    |  Local facility 3                  |
    |  local4    |  Local facility 4                  |
    |  local5    |  Local facility 5                  |
    |  local6    |  Local facility 6                  |
    |  local7    |  Local facility 7                  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in facility, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  facility, value must match: ^[.:a-zA-Z0-9-_/]+$",
							),
						),
					},
				},

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				/* tools/generate-terraform-resource-full/templates/resources/named/resource-model-parent-schema-hack.gotmpl */

				"service": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Frontend service name

`,
					Description: `Frontend service name

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in service, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_/]+$`),
								"illegal character in  service, value must match: ^[.:a-zA-Z0-9-_/]+$",
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

		"level":

		/* tools/generate-terraform-resource-full/templates/resources/common/resource-model-schema-attrtype.gotmpl #resource-model-schema-attrtype */
		schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Logging level

    |  Format   |  Description                         |
    |-----------|--------------------------------------|
    |  emerg    |  Emergency messages                  |
    |  alert    |  Urgent messages                     |
    |  crit     |  Critical messages                   |
    |  err      |  Error messages                      |
    |  warning  |  Warning messages                    |
    |  notice   |  Messages for further investigation  |
    |  info     |  Informational messages              |
    |  debug    |  Debug messages                      |
    |  all      |  Log everything                      |
`,
			Description: `Logging level

    |  Format   |  Description                         |
    |-----------|--------------------------------------|
    |  emerg    |  Emergency messages                  |
    |  alert    |  Urgent messages                     |
    |  crit     |  Critical messages                   |
    |  err      |  Error messages                      |
    |  warning  |  Warning messages                    |
    |  notice   |  Messages for further investigation  |
    |  info     |  Informational messages              |
    |  debug    |  Debug messages                      |
    |  all      |  Log everything                      |
`,

			// Default:          stringdefault.StaticString(`err`),
			Computed: true,
		},

		// TagNodes

		// Nodes

	}
}
