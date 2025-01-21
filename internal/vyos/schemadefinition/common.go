package schemadefinition

func terraformReservedParameternames() []string {
	return []string{
		"id",
		"connection",
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
	NodeType() string
	ValueType() string
}

// NodeParent should match Node and TagNode
type NodeParent interface {
	NodeBase
	InformLinage()
	GetChildren() *Children
	GetChild(childName string) (child NodeBase, err error)
	GetChildAbsPath([]string) (NodeBase, error)
	GetIsBaseNode() bool
	SetIsBaseNode(bool)
	ImportStr() string
	HasSubValue() bool
}
