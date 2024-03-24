package data

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
)

// CtxMutilator changes the context object and returns the new value
type CtxMutilator func(context.Context) context.Context

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

// CtxMutilators is separated out for testing convenience
// only, do not use outside this package
func CtxMutilators(apiEndpoint, apiKey string) []CtxMutilator {
	return []CtxMutilator{
		func(ctx context.Context) context.Context {
			return tflog.MaskLogRegexes(ctx, regexp.MustCompile(apiKey))
		},
		func(ctx context.Context) context.Context {
			return tflog.MaskLogRegexes(ctx, regexp.MustCompile(apiEndpoint))
		},
	}
}
