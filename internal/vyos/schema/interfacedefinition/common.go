package interfacedefinition

import "log"

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
}

// NodeParent should match Node and TagNode
type NodeParent interface {
	NodeBase
	InformLinage()
	GetChildren() *Children
	SetBaseName(string)
}
