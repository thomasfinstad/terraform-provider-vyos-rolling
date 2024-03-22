package schemadefinition

// Nodes will return children of type Node
func (c Children) Nodes() []*Node {
	return c.Node
}

// TagNodes will return children of type TagNode
func (c Children) TagNodes() []*TagNode {
	return c.TagNode
}

// LeafNodes will return children of type LeafNode
func (c Children) LeafNodes() []*LeafNode {
	if c.LeafNode != nil {
		return c.LeafNode
	}

	return []*LeafNode{}
}

// Count will return the total count of children
func (c Children) Count() int {
	return len(c.LeafNode) + len(c.Node) + len(c.TagNode)
}
