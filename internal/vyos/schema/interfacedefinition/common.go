package interfacedefinition

import "log"

func terraformReservedParameternames() []string {
	return []string{
		"id",
		"connection",
	}
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// NodeBase should match Node, TagNode and LeafNode
type NodeBase interface {
	AbsName() []string
	BaseName() string
	Description() string
	AncestorDescription() string
	GetParent() NodeParent
	GetProperties() *Properties
}

// NodeParent should match Node and TagNode
type NodeParent interface {
	NodeBase
	InformLinage()
	GetChildren() *Children
}
