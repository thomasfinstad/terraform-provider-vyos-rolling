package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider/data"
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

	// GetVyosPath returns the list of strings to use to get to the correct vyos configuration
	GetVyosPath() []string

	// GetVyosParentPath returns the list of strings to use to get to the correct
	// vyos configuration for the nearest parent that is not a global resource.
	// If this is the top level named resource the list is zero elements long.
	// This is intended to use with the resource CRUD create function to check if the required parent exists.
	GetVyosParentPath() []string
}

// VyosResourceDataModel defines common functions all models need in order to operate
type VyosResourceDataModel interface {
	ResourceSchemaAttributes() map[string]schema.Attribute
}
