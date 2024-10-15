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

var _ helpers.VyosTopResourceDataModel = &SystemSyslogUserFacility{}

// SystemSyslogUserFacility describes the resource data model.
type SystemSyslogUserFacility struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafSystemSyslogUserFacilityLevel types.String `tfsdk:"level" vyos:"level,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *SystemSyslogUserFacility) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *SystemSyslogUserFacility) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *SystemSyslogUserFacility) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSyslogUserFacility) GetVyosPath() []string {
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
func (o *SystemSyslogUserFacility) GetVyosParentPath() []string {
	return []string{
		"system",

		"syslog",

		"user",

		o.SelfIdentifier.Attributes()["user"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *SystemSyslogUserFacility) GetVyosNamedParentPath() []string {
	return []string{
		"system",

		"syslog",

		"user",

		o.SelfIdentifier.Attributes()["user"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSyslogUserFacility) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
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
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  facility, value must match: ^[.:a-zA-Z0-9-_]+$",
							),
						),
					},
				},

				"user": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Logging to specific terminal of given user

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  Local user account  |
`,
					Description: `Logging to specific terminal of given user

    |  Format  |  Description         |
    |----------|----------------------|
    |  txt     |  Local user account  |
`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in user, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[.:a-zA-Z0-9-_]+$`),
								"illegal character in  user, value must match: ^[.:a-zA-Z0-9-_]+$",
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

		"level": schema.StringAttribute{
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

		// Nodes

	}
}
