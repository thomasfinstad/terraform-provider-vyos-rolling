package helpers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/provider/data"
)

// VyosResource is used to support CRUD operations for terraform resources via helpers
type VyosResource interface {
	// returned model must be ptr
	GetModel() VyosTopResourceDataModel
	GetClient() client.Client
	GetProviderConfig() data.ProviderData
}

// VyosTopResourceDataModel defines common functions all models need in order to operate
type VyosTopResourceDataModel interface {
	VyosResourceDataModel

	// SetID configures the resource ID
	SetID(id []string)

	// GetTimeouts returns resource timeout config
	GetTimeouts() timeouts.Value

	// IsGlobalResource returns true if this is global
	// This is useful during CRUD delete
	IsGlobalResource() bool

	// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
	GetVyosPath() []string

	// GetVyosNamedParentPath returns the list of strings to use to get to the correct
	// vyos configuration for the nearest parent that is not a global resource.
	// If this is the top level named resource the list is zero elements long.
	// This is intended to use with the resource CRUD create function to check if the required parent exists.
	GetVyosNamedParentPath() []string

	// GetVyosParentPath returns the list of strings to use to get to the correct
	// vyos configuration for the nearest parent.
	// If this is the top level resource the list might end up returning the entire interface definition tree.
	// This is intended to use with the resource CRUD read function to check for empty resources.
	GetVyosParentPath() []string
}

// VyosResourceDataModel defines common functions all models need in order to operate
type VyosResourceDataModel interface {
	ResourceSchemaAttributes(context.Context) map[string]schema.Attribute
}
