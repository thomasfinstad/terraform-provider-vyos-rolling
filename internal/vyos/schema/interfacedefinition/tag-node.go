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

// VyosPath returns the name of the node
func (t *TagNode) VyosPath() []string {
	return strings.Split(t.BaseName(), " ")
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

// MutateWithAncestors uses a hacky way to store information about ancestor nodes, letting us keep the autogeneration with less alterations
func (t *TagNode) MutateWithAncestors(nodes []NodeParent) {
	if len(t.Children) != 1 {
		panic(fmt.Sprintf("[%s] Can not add ancestors to node with anything other than 1 child container, node has %d child containers", t.BaseName(), len(t.Children)))
	}

	var c []*Children

	for _, n := range nodes {
		switch n := n.(type) {
		case *Node:
			c = append(c, &Children{Node: []*Node{n}})
		case *TagNode:
			c = append(c, &Children{TagNode: []*TagNode{n}})
		default:
			panic(fmt.Sprintf("[%s] Ancestor node type unknown: %T", t.BaseName(), n))
		}
	}

	t.Children = append(t.Children, c...)
}

// AncestorDescription returns the formatted description all ancestor nodes, including help text
func (t *TagNode) AncestorDescription() string {
	var desc string

	for _, c := range t.Children[1 : len(t.Children)-1] {
		for _, n := range c.Node {
			desc += n.Description()
		}
		for _, t := range c.TagNode {
			desc += t.Description()
		}
	}

	return desc
}
