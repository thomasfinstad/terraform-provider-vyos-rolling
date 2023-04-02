//go:build tools
// +build tools

package tools

import (
	// Ensure documentation tooling is not removed from go.mod.
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	_ "github.com/xuri/xgen/cmd/xgen"
	_ "mvdan.cc/gofumpt"
)
