package interfacedefinition

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// BaseName returns the name of the node
func (t *TagNode) BaseName() string {
	return t.NodeNameAttr
}

// SetBaseName is only intended to be used on the base node to be able to set the entire "family tree name"
func (t *TagNode) SetBaseName(name string) {
	t.NodeNameAttr = name
}

// BaseNameC returns the capitalized BaseName
func (t *TagNode) BaseNameC() string {
	return cases.Title(language.Norwegian).String(t.BaseName())
}

// BaseNameS returns the terraform schema friendly version of BaseName
func (t *TagNode) BaseNameS() string {
	return regexp.MustCompile("[^0-9a-z_]").ReplaceAllString(t.BaseName(), "_")
}

// BaseNameG returns a go friendlier version of BaseName
func (t *TagNode) BaseNameG() string {
	ret := regexp.MustCompile("[^a-z_]").ReplaceAllString(t.BaseName(), "_")

	// Special cases
	switch ret {
	case "interface":
		ret = "iface"
	}

	return ret
}

// BaseNameCG returns a capitalized go friendlier version of BaseName
func (t *TagNode) BaseNameCG() string {
	return cases.Title(language.Norwegian).String(t.BaseNameG())
}

// GetChildren return Children object
func (t *TagNode) GetChildren() *Children {
	if t.Children != nil {
		return t.Children[0]
	}

	return &Children{}
}

// Description returns the formatted description of the node, including help text
func (t TagNode) Description() string {
	var desc string
	for _, p := range t.Properties {
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
