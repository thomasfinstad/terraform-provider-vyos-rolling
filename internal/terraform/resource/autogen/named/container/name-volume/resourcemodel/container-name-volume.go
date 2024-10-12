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

var _ helpers.VyosTopResourceDataModel = &ContainerNameVolume{}

// ContainerNameVolume describes the resource data model.
type ContainerNameVolume struct {
	ID types.String `tfsdk:"id" vyos:"-,tfsdk-id"`

	SelfIdentifier types.Object `tfsdk:"identifier" vyos:"-,self-id"`

	Timeouts timeouts.Value `tfsdk:"timeouts" vyos:"-,timeout"`

	// LeafNodes
	LeafContainerNameVolumeSource      types.String `tfsdk:"source" vyos:"source,omitempty"`
	LeafContainerNameVolumeDestination types.String `tfsdk:"destination" vyos:"destination,omitempty"`
	LeafContainerNameVolumeMode        types.String `tfsdk:"mode" vyos:"mode,omitempty"`
	LeafContainerNameVolumePropagation types.String `tfsdk:"propagation" vyos:"propagation,omitempty"`

	// TagNodes (bools that show if child resources have been configured if they are their own BaseNode)

	// Nodes
}

// SetID configures the resource ID
func (o *ContainerNameVolume) SetID(id []string) {
	o.ID = basetypes.NewStringValue(strings.Join(id, "__"))
}

// GetTimeouts returns resource timeout config
func (o *ContainerNameVolume) GetTimeouts() timeouts.Value {
	return o.Timeouts
}

// IsGlobalResource returns true if this is global
// This is useful during CRUD delete
func (o *ContainerNameVolume) IsGlobalResource() bool {
	return (false)
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ContainerNameVolume) GetVyosPath() []string {
	if o.ID.ValueString() != "" {
		return strings.Split(o.ID.ValueString(), "__")
	}

	return append(
		o.GetVyosParentPath(),
		"volume",
		o.SelfIdentifier.Attributes()["volume"].(types.String).ValueString(),
	)
}

// GetVyosParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent.
// If this is the top level resource the list might end up returning the entire interface definition tree.
// This is intended to use with the resource CRUD read function to check for empty resources.
func (o *ContainerNameVolume) GetVyosParentPath() []string {
	return []string{
		"container",

		"name",

		o.SelfIdentifier.Attributes()["name"].(types.String).ValueString(),
	}
}

// GetVyosNamedParentPath returns the list of strings to use to get to the correct
// vyos configuration for the nearest parent that is not a global resource.
// If this is the top level named resource the list is zero elements long.
// This is intended to use with the resource CRUD create function to check if the required parent exists.
func (o *ContainerNameVolume) GetVyosNamedParentPath() []string {
	return []string{
		"container",

		"name",

		o.SelfIdentifier.Attributes()["name"].(types.String).ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerNameVolume) ResourceSchemaAttributes(ctx context.Context) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource ID, full vyos path to the resource with each field separated by dunder (`__`).",
		},
		"identifier": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"volume": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Mount a volume into the container

`,
					Description: `Mount a volume into the container

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in volume, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  volume, value must match: ^[a-zA-Z0-9-_]*$",
							),
						),
					},
				},

				"name": schema.StringAttribute{
					Required: true,
					MarkdownDescription: `Container name

`,
					Description: `Container name

`,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					}, Validators: []validator.String{
						stringvalidator.All(
							helpers.StringNot(
								stringvalidator.RegexMatches(
									regexp.MustCompile(`^.*__.*$`),
									"double underscores in name, conflicts with the internal resource id",
								),
							),
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[a-zA-Z0-9-_]*$`),
								"illegal character in  name, value must match: ^[a-zA-Z0-9-_]*$",
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

		"source": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Source host directory

    |  Format  |  Description            |
    |----------|-------------------------|
    |  txt     |  Source host directory  |
`,
			Description: `Source host directory

    |  Format  |  Description            |
    |----------|-------------------------|
    |  txt     |  Source host directory  |
`,
		},

		"destination": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Destination container directory

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  txt     |  Destination container directory  |
`,
			Description: `Destination container directory

    |  Format  |  Description                      |
    |----------|-----------------------------------|
    |  txt     |  Destination container directory  |
`,
		},

		"mode": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Volume access mode ro/rw

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  ro      |  Volume mounted into the container as read-only   |
    |  rw      |  Volume mounted into the container as read-write  |
`,
			Description: `Volume access mode ro/rw

    |  Format  |  Description                                      |
    |----------|---------------------------------------------------|
    |  ro      |  Volume mounted into the container as read-only   |
    |  rw      |  Volume mounted into the container as read-write  |
`,

			// Default:          stringdefault.StaticString(`rw`),
			Computed: true,
		},

		"propagation": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Volume bind propagation

    |  Format    |  Description                                                                                                |
    |------------|-------------------------------------------------------------------------------------------------------------|
    |  shared    |  Sub-mounts of the original mount are exposed to replica mounts                                             |
    |  slave     |  Allow replica mount to see sub-mount from the original mount but not vice versa                            |
    |  private   |  Sub-mounts within a mount are not visible to replica mounts or the original mount                          |
    |  rshared   |  Allows sharing of mount points and their nested mount points between both the original and replica mounts  |
    |  rslave    |  Allows mount point and their nested mount points between original an replica mounts                        |
    |  rprivate  |  No mount points within original or replica mounts in any direction                                         |
`,
			Description: `Volume bind propagation

    |  Format    |  Description                                                                                                |
    |------------|-------------------------------------------------------------------------------------------------------------|
    |  shared    |  Sub-mounts of the original mount are exposed to replica mounts                                             |
    |  slave     |  Allow replica mount to see sub-mount from the original mount but not vice versa                            |
    |  private   |  Sub-mounts within a mount are not visible to replica mounts or the original mount                          |
    |  rshared   |  Allows sharing of mount points and their nested mount points between both the original and replica mounts  |
    |  rslave    |  Allows mount point and their nested mount points between original an replica mounts                        |
    |  rprivate  |  No mount points within original or replica mounts in any direction                                         |
`,

			// Default:          stringdefault.StaticString(`rprivate`),
			Computed: true,
		},

		// Nodes

	}
}
