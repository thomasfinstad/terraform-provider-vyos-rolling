package helpers

import "github.com/hashicorp/terraform-plugin-framework/resource/schema"

// VyosTopResourceDataModel defines common functions all models need in order to operate
type VyosTopResourceDataModel interface {
	VyosResourceDataModel
	GetVyosPath() (vyosPath []string)
}

// VyosResourceDataModel defines common functions all models need in order to operate
type VyosResourceDataModel interface {
	ResourceSchemaAttributes() map[string]schema.Attribute
}
