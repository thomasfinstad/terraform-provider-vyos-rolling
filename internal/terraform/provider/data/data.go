package data

import (
	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
)

/*
NewProviderData sets defaults
*/
func NewProviderData(c client.Client) ProviderData {
	return ProviderData{
		Client: c,
		Config: Config{
			CrudSkipCheckParentBeforeCreate: false,
			CrudSkipExistingResourceCheck:   false,
			CrudSkipCheckChildBeforeDelete:  false,
		},
	}
}

/*
ProviderData used to communicate prodiver configs and features to the resources
*/
type ProviderData struct {
	Client client.Client
	Config Config
}

/*
Config contains provider level configurations
*/
type Config struct {
	CrudSkipCheckParentBeforeCreate bool
	CrudSkipExistingResourceCheck   bool
	CrudSkipCheckChildBeforeDelete  bool
	CrudDefaultTimeouts             float64
}
