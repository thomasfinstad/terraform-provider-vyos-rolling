package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
)

// CustomResourceDataModel defines common functions all models need in order to operate
type CustomResourceDataModel interface {
	GetValues() (values map[string]attr.Value)
	GetVyosPath() (vyosPath []string)
}
