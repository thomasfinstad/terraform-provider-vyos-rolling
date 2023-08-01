// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ServiceEventHandlerEvent describes the resource data model.
type ServiceEventHandlerEvent struct {
	SelfIdentifier types.String `tfsdk:"event_id" vyos:",self-id"`

	// LeafNodes

	// TagNodes (Bools that show if child resources have been configured)

	// Nodes
	NodeServiceEventHandlerEventFilter *ServiceEventHandlerEventFilter `tfsdk:"filter" vyos:"filter,omitempty"`
	NodeServiceEventHandlerEventScrIPt *ServiceEventHandlerEventScrIPt `tfsdk:"script" vyos:"script,omitempty"`
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ServiceEventHandlerEvent) GetVyosPath() []string {
	return []string{
		"service",

		"event-handler",

		"event",
		o.SelfIdentifier.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ServiceEventHandlerEvent) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, an amalgamation of the `event_id` and the parents `*_id` fields seperated by dunder `__` starting with top level ancestor.",
		},
		"event_id": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Event handler name

`,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},

		// LeafNodes

		// Nodes

		"filter": schema.SingleNestedAttribute{
			Attributes: ServiceEventHandlerEventFilter{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Logs filter settings

`,
		},

		"script": schema.SingleNestedAttribute{
			Attributes: ServiceEventHandlerEventScrIPt{}.ResourceSchemaAttributes(),
			Optional:   true,
			MarkdownDescription: `Event handler script file

`,
		},
	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ServiceEventHandlerEvent) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ServiceEventHandlerEvent) UnmarshalJSON(_ []byte) error {
	return nil
}
