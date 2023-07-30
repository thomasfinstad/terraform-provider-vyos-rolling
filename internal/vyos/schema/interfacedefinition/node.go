package interfacedefinition

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// InformLinage will recursively set the Parent field in all children
func (o *Node) InformLinage() {
	for _, cn := range o.GetChildren().Node {
		cn.Parent = o
		cn.InformLinage()
	}

	for _, ct := range o.GetChildren().TagNode {
		ct.Parent = o
		ct.InformLinage()
	}

	for _, cl := range o.GetChildren().LeafNode {
		cl.Parent = o
	}
}

// AbsName returns each name in the node hirarchy starting with the root as the first element and this node as the last element
func (o *Node) AbsName() []string {
	if o.Parent == nil {
		return []string{o.BaseName()}
	}

	return append(o.Parent.AbsName(), o.BaseName())
}

// AbsNameR returns full name in the node hirarchy,
// starting with the root as the first element and this node as the last element
// concatinated with underscores to make it tf resource name friendly
// This is intended to be used as a way to generate resource names
func (o *Node) AbsNameR() string {
	return regexp.MustCompile(
		"[ -]",
	).ReplaceAllString(
		strings.Join(
			o.AbsName(),
			" ",
		),
		"_",
	)
}

// BaseName returns the name of the node
func (o *Node) BaseName() string {
	if o.NodeNameAttr == "" {
		return "UNDEF"
	}
	return o.NodeNameAttr
}

// BaseNameC returns the capitalized BaseName
func (o *Node) BaseNameC() string {
	return cases.Title(language.Norwegian).String(o.BaseName())
}

// BaseNameS returns the terraform schema friendly version of BaseName
func (o *Node) BaseNameS() string {
	return regexp.MustCompile("[ -]").ReplaceAllString(o.BaseName(), "_")
}

// BaseNameG returns a go friendlier version of BaseName
func (o *Node) BaseNameG() string {
	capName := o.BaseNameCG()
	r, i := utf8.DecodeRuneInString(capName)
	return string(unicode.ToLower(r)) + capName[i:]
}

// BaseNameCG returns a capitalized go friendlier version of BaseName
func (o *Node) BaseNameCG() string {
	var ret []string

	for _, nameComponent := range o.AbsName() {

		// Convert numbers for go compability
		nameComponent = strings.ReplaceAll(nameComponent, "0", "Zero")
		nameComponent = strings.ReplaceAll(nameComponent, "1", "One")
		nameComponent = strings.ReplaceAll(nameComponent, "2", "Two")
		nameComponent = strings.ReplaceAll(nameComponent, "3", "Three")
		nameComponent = strings.ReplaceAll(nameComponent, "4", "Four")
		nameComponent = strings.ReplaceAll(nameComponent, "5", "Five")
		nameComponent = strings.ReplaceAll(nameComponent, "6", "Six")
		nameComponent = strings.ReplaceAll(nameComponent, "7", "Seven")
		nameComponent = strings.ReplaceAll(nameComponent, "8", "Eight")
		nameComponent = strings.ReplaceAll(nameComponent, "9", "Nine")

		// Pretend non-word characters start a new word
		nameComponent = regexp.MustCompile(`[^A-Za-z]+`).ReplaceAllString(nameComponent, ` `)
		// Capitalize words for improved readability
		nameComponent = cases.Title(language.Norwegian).String(nameComponent)

		// Special cases for go lint compability
		nameComponent = regexp.MustCompile(`(?i)sla`).ReplaceAllString(nameComponent, `SLA`)
		nameComponent = regexp.MustCompile(`(?i)tcp`).ReplaceAllString(nameComponent, `TCP`)
		nameComponent = regexp.MustCompile(`(?i)udp`).ReplaceAllString(nameComponent, `UDP`)
		nameComponent = regexp.MustCompile(`(?i)ip`).ReplaceAllString(nameComponent, `IP`)
		nameComponent = regexp.MustCompile(`(?i)id`).ReplaceAllString(nameComponent, `ID`)
		nameComponent = regexp.MustCompile(`(?i)ttl`).ReplaceAllString(nameComponent, `TTL`)
		nameComponent = regexp.MustCompile(`(?i)dns`).ReplaceAllString(nameComponent, `DNS`)
		nameComponent = regexp.MustCompile(`(?i)url`).ReplaceAllString(nameComponent, `URL`)
		nameComponent = regexp.MustCompile(`(?i)ssh`).ReplaceAllString(nameComponent, `TCP`)
		nameComponent = regexp.MustCompile(`(?i)https`).ReplaceAllString(nameComponent, `HTTPS`)
		nameComponent = regexp.MustCompile(`(?i)http`).ReplaceAllString(nameComponent, `HTTP`)
		nameComponent = regexp.MustCompile(`(?i)tls`).ReplaceAllString(nameComponent, `TLS`)
		nameComponent = regexp.MustCompile(`(?i)smtp`).ReplaceAllString(nameComponent, `SMTP`)
		nameComponent = regexp.MustCompile(`(?i)api`).ReplaceAllString(nameComponent, `API`)
		nameComponent = regexp.MustCompile(`(?i)uri`).ReplaceAllString(nameComponent, `URI`)

		ret = append(ret, strings.Split(nameComponent, " ")...)
	}

	return strings.Join(ret, "")
}

// VyosPath returns the name of the node
func (o *Node) VyosPath() []string {
	return strings.Split(o.BaseName(), " ")
}

// GetChildren return Children object
func (o *Node) GetChildren() *Children {
	if o.Children != nil {
		return o.Children[0]
	}
	return &Children{}
}

// Description returns the formatted description of the node, including help text
func (o Node) Description() string {
	var desc string
	for _, p := range o.Properties {
		if p.Help != nil {
			desc = fmt.Sprintf("%s\n\n", strings.Join(p.Help, "\n"))
		}

		if p.ValueHelp != nil {
			desc += "    |  Format  |  Description  |\n"
			desc += "    |----------|---------------|\n"
			for _, vh := range p.ValueHelp {
				if vh.Format != "" {
					desc += fmt.Sprintf("    |  %s  |", regexp.MustCompile(`\r?\n`).ReplaceAllString(vh.Format, " "))
				} else {
					desc += "    |   |"
				}
				if vh.Format != "" {
					desc += fmt.Sprintf("  %s  |\n", regexp.MustCompile(`\r?\n`).ReplaceAllString(vh.Description, " "))
				} else {
					desc += "   |\n"
				}
			}

			desc += "\n"
		}
	}

	return desc
}

// AncestorDescription returns the formatted description all ancestor nodes, including help text
func (o *Node) AncestorDescription() string {
	var desc string

	if o.Parent != nil {
		desc = o.Parent.AncestorDescription()
	}

	desc += o.Description()

	return desc
}

// NodeType returns a string of node type
func (o *Node) NodeType() string {
	return "Node"
}
