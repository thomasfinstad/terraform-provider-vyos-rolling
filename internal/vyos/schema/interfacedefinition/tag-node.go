package interfacedefinition

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// InformLinage will recursively set the Parent field in all children
func (t *TagNode) InformLinage() {
	for _, cn := range t.GetChildren().Node {
		cn.Parent = t
		cn.InformLinage()
	}

	for _, ct := range t.GetChildren().TagNode {
		ct.Parent = t
		ct.InformLinage()
	}

	for _, cl := range t.GetChildren().LeafNode {
		cl.Parent = t
	}
}

// AbsName returns each name in the node hirarchy starting with the root as the first element and this node as the last element
func (t *TagNode) AbsName() []string {
	if t.Parent == nil {
		return []string{t.BaseName()}
	}

	return append(t.Parent.AbsName(), t.BaseName())
}

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
	return regexp.MustCompile("[ -]").ReplaceAllString(t.BaseName(), "_")
}

// BaseNameG returns a go friendlier version of BaseName
func (t *TagNode) BaseNameG() string {
	bn := t.BaseName()
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
