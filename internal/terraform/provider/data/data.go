package data

import (
	"context"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/client"
)

/*
ProviderData used to communicate prodiver configs and features to the resources
*/
type ProviderData struct {
	Client        client.Client
	Config        Config
	ctxMutilators []CtxMutilator
}

// CtxMutilatorAdd adds mutilator functions to the ProviderData,
// this is useful for creating provider wide ctx mutilators
// as they must be added everywhere
func (d *ProviderData) CtxMutilatorAdd(cm ...CtxMutilator) {
	d.ctxMutilators = append(d.ctxMutilators, cm...)
}

// CtxMutilatorRun runs all provider configured context mutilators and returns the
// modifiex context
func (d ProviderData) CtxMutilatorRun(ctx context.Context) context.Context {
	for _, fn := range d.ctxMutilators {
		ctx = fn(ctx)
	}

	return ctx
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
