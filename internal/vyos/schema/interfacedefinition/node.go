package interfacedefinition

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// BaseName returns the name of the node
func (n *Node) BaseName() string {
	if n.NodeNameAttr == "" {
		return "UNDEF"
	}
	return n.NodeNameAttr
}

// SetBaseName is only intended to be used on the base node to be able to set the entire "family tree name"
func (n *Node) SetBaseName(name string) {
	n.NodeNameAttr = name
}

// BaseNameC returns the capitalized BaseName
func (n *Node) BaseNameC() string {
	return cases.Title(language.Norwegian).String(n.BaseName())
}

// BaseNameS returns the terraform schema friendly version of BaseName
func (n *Node) BaseNameS() string {
	return regexp.MustCompile("[^0-9a-z_]").ReplaceAllString(n.BaseName(), "_")
}

// BaseNameG returns a go friendlier version of BaseName
func (n *Node) BaseNameG() string {
	ret := regexp.MustCompile("[^a-z_]").ReplaceAllString(n.BaseName(), "_")

	// Special cases
	switch ret {
	case "interface":
		ret = "iface"
	}

	return ret
}

// BaseNameCG returns a capitalized go friendlier version of BaseName
func (n *Node) BaseNameCG() string {
	return cases.Title(language.Norwegian).String(n.BaseNameG())
}

// GetChildren return Children object
func (n *Node) GetChildren() *Children {

	if n.Children != nil {

		return n.Children[0]
	}
	return &Children{}
}

// Description returns the formatted description of the node, including help text
func (n Node) Description() string {
	var desc string
	for _, p := range n.Properties {
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
