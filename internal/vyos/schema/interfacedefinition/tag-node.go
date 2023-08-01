package interfacedefinition

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// InformLinage will recursively set the Parent field in all children
func (o *TagNode) InformLinage() {
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
func (o *TagNode) AbsName() []string {
	if o.Parent == nil {
		return []string{o.BaseName()}
	}

	return append(o.Parent.AbsName(), o.BaseName())
}

// AbsNameR returns full name in the node hirarchy,
// starting with the root as the first element and this node as the last element
// concatinated with underscores to make it tf resource name friendly
// This is intended to be used as a way to generate resource names
func (o *TagNode) AbsNameR() string {
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
func (o *TagNode) BaseName() string {
	return o.NodeNameAttr
}

// BaseNameC returns the capitalized BaseName
func (o *TagNode) BaseNameC() string {
	return cases.Title(language.Norwegian).String(o.BaseName())
}

// BaseNameS returns the terraform schema friendly version of BaseName
func (o *TagNode) BaseNameS() string {
	return regexp.MustCompile("[ -]").ReplaceAllString(o.BaseName(), "_")
}

// BaseNameSB returns a resource top level parameter safe version of BaseNameS()
func (o *TagNode) BaseNameSB() string {
	name := o.BaseNameS()

	if helpers.ListContains(terraformReservedParameternames(), name) {
		panic(name)
		return fmt.Sprintf("%s_param", name)
	}
	return name
}

// BaseNameG returns a go friendlier version of BaseName
func (o *TagNode) BaseNameG() string {
	capName := o.BaseNameCG()
	r, i := utf8.DecodeRuneInString(capName)
	return string(unicode.ToLower(r)) + capName[i:]
}

// BaseNameCG returns a capitalized go friendlier version of BaseName
func (o *TagNode) BaseNameCG() string {
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
func (o *TagNode) VyosPath() []string {
	return strings.Split(o.BaseName(), " ")
}

// GetChildren return Children object
func (o *TagNode) GetChildren() *Children {
	if o.Children != nil {
		return o.Children[0]
	}

	return &Children{}
}

// Description returns the formatted description of the node, including help text
func (o TagNode) Description() string {
	strReplaceFormat := map[string]string{
		`\r?\n`: "<br>",
		`u32:`:  "number: ",
	}

	var desc string
	for _, p := range o.Properties {
		if p.Help != nil {
			desc = fmt.Sprintf("%s\n\n", strings.Join(p.Help, "\n"))
		}

		// Adding &emsp; gives tables better readability on terraform registry documentation page
		if p.ValueHelp != nil {
			desc += "    |  Format &emsp; | Description  |\n"
			desc += "    |----------|---------------|\n"
			for _, vh := range p.ValueHelp {
				// Format Column
				if vh.Format != "" {
					formatStr := vh.Format
					for regex, repl := range strReplaceFormat {
						formatStr = regexp.MustCompile(regex).ReplaceAllString(formatStr, repl)
					}
					desc += fmt.Sprintf("    |  %s  &emsp; |", formatStr)
				} else {
					desc += "    |   &emsp; |"
				}

				// Description Column
				if vh.Format != "" {
					formatStr := vh.Description
					for regex, repl := range strReplaceFormat {
						formatStr = regexp.MustCompile(regex).ReplaceAllString(formatStr, repl)
					}
					desc += fmt.Sprintf("  %s  |\n", formatStr)
				} else {
					desc += "   |\n"
				}
			}

			desc += "\n"
		}
	}

	return desc
}

// MutateWithAncestors uses a hacky way to store information about ancestor nodes, letting us keep the autogeneration with less alterations
func (o *TagNode) MutateWithAncestors(nodes []NodeParent) {
	if len(o.Children) != 1 {
		panic(fmt.Sprintf("[%s] Can not add ancestors to node with anything other than 1 child container, node has %d child containers", o.BaseName(), len(o.Children)))
	}

	var c []*Children

	for _, n := range nodes {
		switch n := n.(type) {
		case *Node:
			c = append(c, &Children{Node: []*Node{n}})
		case *TagNode:
			c = append(c, &Children{TagNode: []*TagNode{n}})
		default:
			panic(fmt.Sprintf("[%s] Ancestor node type unknown: %T", o.BaseName(), n))
		}
	}

	o.Children = append(o.Children, c...)
}

// AncestorDescription returns the formatted description all ancestor nodes, using help text
func (o *TagNode) AncestorDescription() string {
	// Recursively fetch unformatted descriptions and add them with formatting
	var getOldies func(n NodeBase) string
	getOldies = func(n NodeBase) string {
		subdesc := ""
		if parent := n.GetParent(); parent != nil {
			subdesc += getOldies(parent)
		}

		if p := n.GetProperties(); p != nil {
			if h := p.Help; h != nil {
				subdesc += fmt.Sprintf("%s\n\n", strings.Join(h, "\n<br>"))
			} else {
				subdesc += fmt.Sprintf("<i>%s</i>\n\n", n.BaseName())
			}
		} else {
			subdesc += fmt.Sprintf("<i>%s</i>\n\n", n.BaseName())
		}

		subdesc += "<br>\n"
		subdesc += "&darr;\n"
		subdesc += "<br>\n"

		return subdesc
	}

	// Center text
	desc := "<div style=\"text-align: center\">\n"

	// Ancestors descriptions
	if parent := o.GetParent(); parent != nil {
		desc += getOldies(parent)
	}

	// Make our own description bold
	desc += "<b>\n"
	if h := o.GetProperties().Help; h != nil {
		desc += fmt.Sprintf("%s\n", strings.Join(h, "\n"))
	}
	desc += "</b>\n"
	desc += "</div>\n"

	return desc
}

// ValueType the type of value that can be expected for this node, one of:
// string (default)
// number (if value help only lists u32)
func (o *TagNode) ValueType() string {
	isNum := false
	for _, vh := range o.Properties[0].ValueHelp {
		if strings.HasPrefix(vh.Format, "u32") {
			isNum = true
		} else {
			isNum = false
			break
		}
	}
	if isNum {
		return "number"
	}

	return "string"
}

// NodeType returns a string of node type
func (o *TagNode) NodeType() string {
	return "TagNode"
}

// GetParent returns parent, can be nil
func (o *TagNode) GetParent() NodeParent {
	return o.Parent
}

// GetProperties returns nodes properties
func (o *TagNode) GetProperties() *Properties {
	if len(o.Properties) > 0 {
		return o.Properties[0]
	}
	return nil
}
