//go:build tools
// +build tools

package tools

// Ensure documentation tooling is not removed from go.mod.
import (
	// XSD schema tool, not a main package so to ensure versioning it can not be installed with `go install ....`
	_ "github.com/xuri/xgen/cmd/xgen"
)
