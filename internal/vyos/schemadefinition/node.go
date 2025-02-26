package schemadefinition

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/terraform/helpers/tools"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Validate conformity
var _ NodeParent = &Node{}

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

// ImportStr returns the string used to create import documentation
func (o *Node) ImportStr() string {
	if o.Parent == nil {
		return o.BaseName()
	}
	return fmt.Sprintf("%s__%s", o.Parent.ImportStr(), o.BaseName())
}

// AbsName returns each name in the node hierarchy starting with the root as the first element and this node as the last element
func (o *Node) AbsName() []string {
	if o.Parent == nil {
		return []string{o.BaseName()}
	}

	return append(o.Parent.AbsName(), o.BaseName())
}

// BaseNameR returns name of this node,
// This is intended to be used as a way to generate resource names
func (o *Node) BaseNameR() string {
	return regexp.MustCompile(
		"[ -]",
	).ReplaceAllString(
		o.BaseName(),
		"_",
	)
}

// AbsNameR returns full name in the node hierarchy,
// starting with the root as the first element and this node as the last element
// concatenated with underscores to make it tf resource name friendly
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
	ret := regexp.MustCompile("[ -]").ReplaceAllString(o.BaseName(), "_")
	if regexp.MustCompile("^[0-9]").FindString(ret) != "" {
		ret = "_" + ret
	}
	return ret
}

// BaseNameSB returns a resource top level parameter safe version of BaseNameS()
func (o *Node) BaseNameSB() string {
	name := o.BaseNameS()

	if tools.ListContains(terraformReservedParameternames(), name) {
		return fmt.Sprintf("%s_param", name)
	}
	return name
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

		// Convert numbers for go compatibility
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

		// Special cases for go lint compatibility
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

func (o *Node) GetChildAbsPath(absChildPath []string) (child NodeBase, err error) {
	match := 0
	for i := range min(len(absChildPath), len(o.AbsName())) {
		if o.AbsName()[i] == absChildPath[i] {
			match = i + 1
		} else {
			break
		}
	}
	return o.getChild(absChildPath[match:], absChildPath)
}

func (o *Node) GetChild(childName string) (child NodeBase, err error) {
	return o.getChild([]string{childName}, append(o.AbsName(), childName))
}

func (o *Node) getChild(childName, absChildPath []string) (child NodeBase, err error) {
	if slices.Equal(o.AbsName(), absChildPath) {
		return o, nil
	}

	if len(childName) < 1 {
		return nil, fmt.Errorf("recursed too deep")
	}

	for _, t := range o.GetChildren().TagNode {
		if t.BaseName() == childName[0] {
			child, err = t.getChild(childName[1:], absChildPath)
			if err == nil {
				return child, nil
			}
		}
	}

	for _, n := range o.GetChildren().Node {
		if n.BaseName() == childName[0] {
			child, err = n.getChild(childName[1:], absChildPath)
			if err == nil {
				return child, nil
			}
		}
	}

	for _, l := range o.GetChildren().LeafNode {
		if slices.Equal(l.AbsName(), absChildPath) {
			return l, nil
		}
	}

	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("[%s] do not have child: %s", strings.Join(o.AbsName(), " "), strings.Join(childName, " "))
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
	regexReplaceWithString := map[string]string{}
	ret := o.MarkdownDescription()
	for regex, repl := range regexReplaceWithString {
		ret = regexp.MustCompile(regex).ReplaceAllString(ret, repl)
	}

	return ret
}

// MarkdownDescription returns the markdown formatted description of the node, including help text
func (o Node) MarkdownDescription() string {
	regexReplaceWithString := map[string]string{
		`u32:`: "",
		`\\n`:  "</br>",
	}

	var desc string
	for _, p := range o.Properties {
		if p.Help != nil {
			desc = fmt.Sprintf("%s\n\n", strings.Join(p.Help, "\n"))
		}

		// Adding &emsp; gives tables better readability on terraform registry documentation page
		if p.ValueHelp != nil {
			var valueHelp [][]string
			formatMaxLen := len("Format")
			descriptionMaxLen := len("Description")
			for _, vh := range p.ValueHelp {
				format := ""
				description := "N/A"
				if vh.Format != "N/A" {
					// Format Column

					format = vh.Format
					for regex, repl := range regexReplaceWithString {
						format = regexp.MustCompile(regex).ReplaceAllString(format, repl)
					}
					formatMaxLen = max(formatMaxLen, len(format))
				}

				if vh.Description != "" {
					// Description Column
					description = vh.Description
					for regex, repl := range regexReplaceWithString {
						description = regexp.MustCompile(regex).ReplaceAllString(description, repl)
					}
					descriptionMaxLen = max(descriptionMaxLen, len(description))
				}

				// Add modified help txt
				valueHelp = append(valueHelp, []string{format, description})
			}

			// Add paded text to description
			desc += fmt.Sprintf("    |  %-*s  |  %-*s  |\n", formatMaxLen, "Format", descriptionMaxLen, "Description") +
				fmt.Sprintf("    |--%s--|--%s--|\n", strings.Repeat("-", formatMaxLen), strings.Repeat("-", descriptionMaxLen))

			for _, row := range valueHelp {
				desc += fmt.Sprintf("    |  %-*s  |  %-*s  |\n", formatMaxLen, row[0], descriptionMaxLen, row[1])
			}
		}
	}

	return desc
}

// AncestorDescription returns the formatted description all ancestor nodes, using help text
func (o *Node) AncestorDescription() string {
	// Recursively fetch unformatted descriptions and add them with formatting
	var getOldies func(n NodeBase) string
	getOldies = func(n NodeBase) string {
		subdesc := ""
		if parent := n.GetParent(); parent != nil {
			subdesc += getOldies(parent)
		}

		if p := n.GetProperties(); p != nil {
			if h := p.Help; h != nil {
				subdesc += strings.Join(h, "\n")
			} else {
				subdesc += fmt.Sprintf("*%s*", n.BaseName())
			}
		} else {
			subdesc += fmt.Sprintf("*%s*", n.BaseName())
		}

		subdesc += "  \n⯯  \n"

		return subdesc
	}

	desc := ""

	// Ancestors descriptions
	if parent := o.GetParent(); parent != nil {
		desc += getOldies(parent)
	}

	// Make our own description bold
	if p := o.GetProperties(); p != nil {
		if h := p.Help; h != nil {
			desc += fmt.Sprintf("**%s**\n", strings.Join(h, "\n"))
		}
	}

	return desc
}

// NodeType returns a string of node type
func (o *Node) NodeType() string {
	return "Node"
}

// GetParent returns parent, can be nil
func (o *Node) GetParent() NodeParent {
	return o.Parent
}

// GetProperties returns nodes properties
func (o *Node) GetProperties() *Properties {
	if len(o.Properties) > 0 {
		return o.Properties[0]
	}
	return nil
}

// GetIsBaseNode returns if this parent is a basenode (root node for a resource)
func (o *Node) GetIsBaseNode() bool {
	return o.IsBaseNode
}

// SetIsBaseNode configures if this node is a basenode
func (o *Node) SetIsBaseNode(isBaseNode bool) {
	o.IsBaseNode = isBaseNode
}

// HasSubValue is used to inform if this or child NodeParents
// has LeafNodes where no intermediate NodeParent or this NodeParent is itself a BaseNode
func (o *Node) HasSubValue() bool {
	if o.GetIsBaseNode() {
		return false
	}

	if len(o.GetChildren().LeafNodes()) > 0 {
		return true
	}

	if np, ok := o.GetChildren().GetNodeParents(); ok {
		for _, c := range np {
			if c.HasSubValue() {
				return true
			}
		}
	}

	return false
}

// ValueType for Node will always be "NO_VALUE"
func (o *Node) ValueType() string {
	return "NO_VALUE"
}
