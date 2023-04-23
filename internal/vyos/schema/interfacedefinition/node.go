package interfacedefinition

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// InformLinage will recursively set the Parent field in all children
func (n *Node) InformLinage() {
	for _, cn := range n.GetChildren().Node {
		cn.Parent = n
		cn.InformLinage()
	}

	for _, ct := range n.GetChildren().TagNode {
		ct.Parent = n
		ct.InformLinage()
	}

	for _, cl := range n.GetChildren().LeafNode {
		cl.Parent = n
	}
}

// AbsName returns each name in the node hirarchy starting with the root as the first element and this node as the last element
func (n *Node) AbsName() []string {
	if n.Parent == nil {
		return []string{n.BaseName()}
	}

	return append(n.Parent.AbsName(), n.BaseName())
}

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
	return regexp.MustCompile("[ -]").ReplaceAllString(n.BaseName(), "_")
}

// BaseNameG returns a go friendlier version of BaseName
func (n *Node) BaseNameG() string {
	bn := n.BaseName()
	bn = strings.ReplaceAll(bn, "0", "zero")
	bn = strings.ReplaceAll(bn, "1", "one")
	bn = strings.ReplaceAll(bn, "2", "two")
	bn = strings.ReplaceAll(bn, "3", "three")
	bn = strings.ReplaceAll(bn, "4", "four")
	bn = strings.ReplaceAll(bn, "5", "five")
	bn = strings.ReplaceAll(bn, "6", "six")
	bn = strings.ReplaceAll(bn, "7", "seven")
	bn = strings.ReplaceAll(bn, "8", "eight")
	bn = strings.ReplaceAll(bn, "9", "nine")
	ret := regexp.MustCompile("[^a-z_]").ReplaceAllString(bn, "_")

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

// VyosPath returns the name of the node
func (n *Node) VyosPath() []string {
	return strings.Split(n.BaseName(), " ")
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
			desc += "|  Format  |  Description  |\n"
			desc += "|----------|---------------|\n"
			for _, vh := range p.ValueHelp {
				if vh.Format != "" {
					desc += fmt.Sprintf("|  %s  |", vh.Format)
				} else {
					desc += "|   |"
				}
				if vh.Format != "" {
					desc += fmt.Sprintf("  %s  |\n", vh.Description)
				} else {
					desc += "   |\n"
				}
			}
		}
	}

	return desc
}
