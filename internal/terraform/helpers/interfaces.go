package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
)

// VyosResource is used to support CRUD operations for terraform resources via helpers
type VyosResource interface {
	GetName() string
	GetModel() VyosTopResourceDataModel
	GetClient() *client.Client
}

// VyosTopResourceDataModel defines common functions all models need in order to operate
type VyosTopResourceDataModel interface {
	VyosResourceDataModel
	GetID() *types.String
	SetID(id types.String)
	GetVyosPath() (vyosPath []string)
}

// VyosResourceDataModel defines common functions all models need in order to operate
type VyosResourceDataModel interface {
	ResourceSchemaAttributes() map[string]schema.Attribute
}
