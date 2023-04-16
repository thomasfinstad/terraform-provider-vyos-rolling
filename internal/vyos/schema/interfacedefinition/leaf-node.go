package interfacedefinition

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AbsName returns each name in the node hirarchy starting with the root as the first element and this node as the last element
func (l *LeafNode) AbsName() []string {
	if l.Parent == nil {
		return []string{l.BaseName()}
	}

	return append(l.Parent.AbsName(), l.BaseName())
}

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
	bn := l.BaseName()
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

// Default returns the default value or empty string if no default is provided
func (l LeafNode) Default() string {
	var def string

	def += fmt.Sprintf("`%s`", strings.Join(l.DefaultValue, "`,`"))

	if def != "``" {
		return def
	}

	return ""
}
