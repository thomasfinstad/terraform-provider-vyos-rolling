// Package resourcemodel code generated by /workspaces/terraform-provider-vyos/tools/build-terraform-resource-full/main.go. DO NOT EDIT.
package resourcemodel

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ContainerName describes the resource data model.
type ContainerName struct {
	ID types.String `tfsdk:"identifier" vyos:",self-id"`

	// LeafNodes
	LeafContainerNameAllowHostNetworks types.Bool   `tfsdk:"allow_host_networks" vyos:"allow-host-networks,omitempty"`
	LeafContainerNameCapAdd            types.List   `tfsdk:"cap_add" vyos:"cap-add,omitempty"`
	LeafContainerNameDescrIPtion       types.String `tfsdk:"description" vyos:"description,omitempty"`
	LeafContainerNameDisable           types.Bool   `tfsdk:"disable" vyos:"disable,omitempty"`
	LeafContainerNameEntrypoint        types.String `tfsdk:"entrypoint" vyos:"entrypoint,omitempty"`
	LeafContainerNameHostName          types.String `tfsdk:"host_name" vyos:"host-name,omitempty"`
	LeafContainerNameImage             types.String `tfsdk:"image" vyos:"image,omitempty"`
	LeafContainerNameCommand           types.String `tfsdk:"command" vyos:"command,omitempty"`
	LeafContainerNameArguments         types.String `tfsdk:"arguments" vyos:"arguments,omitempty"`
	LeafContainerNameMemory            types.Number `tfsdk:"memory" vyos:"memory,omitempty"`
	LeafContainerNameSharedMemory      types.Number `tfsdk:"shared_memory" vyos:"shared-memory,omitempty"`
	LeafContainerNameRestart           types.String `tfsdk:"restart" vyos:"restart,omitempty"`

	// TagNodes (Bools that show if child resources have been configured)
	ExistsTagContainerNameDevice      bool `tfsdk:"device" vyos:"device,child"`
	ExistsTagContainerNameEnvironment bool `tfsdk:"environment" vyos:"environment,child"`
	ExistsTagContainerNameNetwork     bool `tfsdk:"network" vyos:"network,child"`
	ExistsTagContainerNamePort        bool `tfsdk:"port" vyos:"port,child"`
	ExistsTagContainerNameVolume      bool `tfsdk:"volume" vyos:"volume,child"`

	// Nodes
}

// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
func (o *ContainerName) GetVyosPath() []string {
	return []string{
		"container",

		"name",
		o.ID.ValueString(),
	}
}

// ResourceSchemaAttributes generates the schema attributes for the resource at this level
func (o ContainerName) ResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"identifier": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Container name

`,
		},

		// LeafNodes

		"allow_host_networks": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Allow host networks in container

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"cap_add": schema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			MarkdownDescription: `Container capabilities/permissions

    |  Format  |  Description  |
    |----------|---------------|
    |  net-admin  |  Network operations (interface, firewall, routing tables)  |
    |  net-bind-service  |  Bind a socket to privileged ports (port numbers less than 1024)  |
    |  net-raw  |  Permission to create raw network sockets  |
    |  setpcap  |  Capability sets (from bounded or inherited set)  |
    |  sys-admin  |  Administation operations (quotactl, mount, sethostname, setdomainame)  |
    |  sys-time  |  Permission to set system clock  |

`,
		},

		"description": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Description

    |  Format  |  Description  |
    |----------|---------------|
    |  txt  |  Description  |

`,
		},

		"disable": schema.BoolAttribute{
			Optional: true,
			MarkdownDescription: `Disable instance

`,
			Default:  booldefault.StaticBool(false),
			Computed: true,
		},

		"entrypoint": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override the default ENTRYPOINT from the image

`,
		},

		"host_name": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Container host name

`,
		},

		"image": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Image name in the hub-registry

`,
		},

		"command": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Override the default CMD from the image

`,
		},

		"arguments": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `The command's arguments for this container

`,
		},

		"memory": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Memory (RAM) available to this container

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0  |  Unlimited  |
    |  u32:1-16384  |  Container memory in megabytes (MB)  |

`,

			// Default:          stringdefault.StaticString(`512`),
			Computed: true,
		},

		"shared_memory": schema.NumberAttribute{
			Optional: true,
			MarkdownDescription: `Shared memory available to this container

    |  Format  |  Description  |
    |----------|---------------|
    |  u32:0  |  Unlimited  |
    |  u32:1-8192  |  Container memory in megabytes (MB)  |

`,

			// Default:          stringdefault.StaticString(`64`),
			Computed: true,
		},

		"restart": schema.StringAttribute{
			Optional: true,
			MarkdownDescription: `Restart options for container

    |  Format  |  Description  |
    |----------|---------------|
    |  no  |  Do not restart containers on exit  |
    |  on-failure  |  Restart containers when they exit with a non-zero exit code, retrying indefinitely  |
    |  always  |  Restart containers when they exit, regardless of status, retrying indefinitely  |

`,

			// Default:          stringdefault.StaticString(`on-failure`),
			Computed: true,
		},

		// Nodes

	}
}

// MarshalJSON returns json encoded string as bytes or error if marshalling did not go well
func (o *ContainerName) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON unmarshals json byte array into this object
func (o *ContainerName) UnmarshalJSON(_ []byte) error {
	return nil
}
