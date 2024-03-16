package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
)

// VyosResource is used to support CRUD operations for terraform resources via helpers
type VyosResource interface {
	// returned model must be ptr
	GetModel() VyosTopResourceDataModel
	GetClient() client.Client
}

// VyosTopResourceDataModel defines common functions all models need in order to operate
type VyosTopResourceDataModel interface {
	VyosResourceDataModel
	SetID(id []string)
	GetVyosPath() []string
}

// VyosResourceDataModel defines common functions all models need in order to operate
type VyosResourceDataModel interface {
	ResourceSchemaAttributes() map[string]schema.Attribute
}
