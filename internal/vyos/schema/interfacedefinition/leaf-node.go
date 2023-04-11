package interfacedefinition

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// BaseName returns the name of the node
func (l *LeafNode) BaseName() string {
	return l.NodeNameAttr
}

// BaseNameC returns the capitalized BaseName
func (l *LeafNode) BaseNameC() string {
	return cases.Title(language.Norwegian).String(l.BaseName())
}

// BaseNameS returns the terraform schema friendly version of BaseName
func (l *LeafNode) BaseNameS() string {
	return regexp.MustCompile("[^0-9a-z_]").ReplaceAllString(l.BaseName(), "_")
}

// BaseNameG returns a go friendlier version of BaseName
func (l *LeafNode) BaseNameG() string {
	ret := regexp.MustCompile("[^a-z_]").ReplaceAllString(l.BaseName(), "_")

	// Special cases
	switch ret {
	case "interface":
		ret = "iface"
	}

	return ret
}

// BaseNameCG returns a capitalized go friendlier version of BaseName
func (l *LeafNode) BaseNameCG() string {
	return cases.Title(language.Norwegian).String(l.BaseNameG())
}

// Description returns the formatted description of the node, including help text
func (l LeafNode) Description() string {
	var desc string
	for _, p := range l.Properties {
		if p.Help != nil {
			desc = fmt.Sprintf("%s\n\n", strings.Join(p.Help, "\n"))
		}

		if p.ValueHelp != nil {
			for _, vh := range p.ValueHelp {
				if vh.Format != "" {
					desc += fmt.Sprintf("Format: %s\n", vh.Format)
				}
				if vh.Format != "" {
					desc += fmt.Sprintf("%s\n", vh.Description)
				}
			}
		}
	}

	return desc
}

// Default returns the default value or empty string if no default is provided
func (l LeafNode) Default() string {
	var def string

	def += fmt.Sprintf("`%s`", strings.Join(l.DefaultValue, "`,`"))

	if def != "``" {
		return def
	}

	return ""
}
