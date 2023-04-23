package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
)

// CustomResourceDataModel defines common functions all models need in order to operate
type CustomResourceDataModel interface {
	GetValues() (vyosPath []string, values map[string]attr.Value)
}
