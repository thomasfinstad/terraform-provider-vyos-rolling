// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SystemSyslogFileFacility describes the resource data model.
type SystemSyslogFileFacility struct {
	SelfIdentifier types.String `tfsdk:"facility_id" vyos:",self-id"`

	ParentIDSystemSyslogFile types.String `tfsdk:"file" vyos:"file,parent-id"`

	// LeafNodes
	LeafSystemSyslogFileFacilityLevel types.String `tfsdk:"level" vyos:"level,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *SystemSyslogFileFacility) GetVyosPath() []string {
	return []string{
		"system",

		"syslog",

		"file",
		o.ParentIDSystemSyslogFile.ValueString(),

		"facility",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o SystemSyslogFileFacility) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `facility_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"facility_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Facility for logging

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  all  &emsp; |  All facilities excluding "mark"  |
    |  auth  &emsp; |  Authentication and authorization  |
    |  authpriv  &emsp; |  Non-system authorization  |
    |  cron  &emsp; |  Cron daemon  |
    |  daemon  &emsp; |  System daemons  |
    |  kern  &emsp; |  Kernel  |
    |  lpr  &emsp; |  Line printer spooler  |
    |  mail  &emsp; |  Mail subsystem  |
    |  mark  &emsp; |  Timestamp  |
    |  news  &emsp; |  USENET subsystem  |
    |  protocols  &emsp; |  depricated will be set to local7  |
    |  security  &emsp; |  depricated will be set to auth  |
    |  syslog  &emsp; |  Authentication and authorization  |
    |  user  &emsp; |  Application processes  |
    |  uucp  &emsp; |  UUCP subsystem  |
    |  local0  &emsp; |  Local facility 0  |
    |  local1  &emsp; |  Local facility 1  |
    |  local2  &emsp; |  Local facility 2  |
    |  local3  &emsp; |  Local facility 3  |
    |  local4  &emsp; |  Local facility 4  |
    |  local5  &emsp; |  Local facility 5  |
    |  local6  &emsp; |  Local facility 6  |
    |  local7  &emsp; |  Local facility 7  |

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		"file_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Logging to a file

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		"level": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Logging level

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  emerg  &emsp; |  Emergency messages  |
    |  alert  &emsp; |  Urgent messages  |
    |  crit  &emsp; |  Critical messages  |
    |  err  &emsp; |  Error messages  |
    |  warning  &emsp; |  Warning messages  |
    |  notice  &emsp; |  Messages for further investigation  |
    |  info  &emsp; |  Informational messages  |
    |  debug  &emsp; |  Debug messages  |
    |  all  &emsp; |  Log everything  |

`,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *SystemSyslogFileFacility) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *SystemSyslogFileFacility) UnmarshalJSON(_ []byte) error {
	return nil
}
