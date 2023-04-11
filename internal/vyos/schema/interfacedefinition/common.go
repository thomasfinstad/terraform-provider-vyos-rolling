package interfacedefinition

import "log"

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// NodeBase should match Node, TagNode and LeafNode
type NodeBase interface {
	BaseName() string
	Description() string
}

// NodeParent should match Node and TagNode
type NodeParent interface {
	NodeBase
	GetChildren() *Children
	SetBaseName(string)
}
