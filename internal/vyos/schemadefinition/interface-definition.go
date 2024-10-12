package schemadefinition

import (
	"fmt"
)

/*
Contains structs to use with vyos schemas
*/

// GetRootNode returns the top level node from the interface definition
func (i *InterfaceDefinition) GetRootNode() (*Node, error) {
	if len(i.Node) == 0 {
		return &Node{}, fmt.Errorf("no root node found under interface")
	} else if len(i.Node) != 1 {
		return &Node{}, fmt.Errorf("too many root nodes found under interface: %v", i.XMLName)
	}

	ret := i.Node[0]

	ret.InformLinage()

	return ret, nil
}

// BaseTagNodes are the resource basis. All TagNodes start with the assumption that they are base nodes, use overrides to change that.
func (i *InterfaceDefinition) BaseTagNodes() (tagNodes []*TagNode, ok bool) {
	rootNode, err := i.GetRootNode()
	if err != nil {
		fmt.Printf("BaseTagNodes Skipping rootnode:%v because: %s\n", i, err)
		return nil, false
	}

	var recursivelyGetBaseTagNodes func(rootNode NodeParent) []*TagNode
	recursivelyGetBaseTagNodes = func(rootNode NodeParent) []*TagNode {
		var ret []*TagNode

		children := rootNode.GetChildren()

		if children == nil {
			fmt.Printf("[%s] Skipping children:nil\n", rootNode.BaseName())
			return ret
		}

		// Recurse over tagnode children
		for _, n := range children.TagNodes() {
			if n.GetIsBaseNode() {
				ret = append(ret, n)
			}
			ret = append(ret, recursivelyGetBaseTagNodes(n)...)
		}

		// Recurse over node children
		for _, n := range children.Nodes() {
			ret = append(ret, recursivelyGetBaseTagNodes(n)...)
		}

		return ret
	}

	btn := recursivelyGetBaseTagNodes(rootNode)
	return btn, len(btn) > 0
}

// Nodes returns all Nodes that contain a LeafNode and is itself not inside a TagNode
func (i *InterfaceDefinition) BaseNodes() (nodes []*Node, ok bool) {
	var recursivelyGetBaseNodes func(NodeParent) []*Node
	recursivelyGetBaseNodes = func(parent NodeParent) []*Node {
		var ret []*Node
		children := parent.GetChildren()

		if children == nil {
			fmt.Printf("[%s] Skipping children:nil\n", parent.BaseName())
			return ret
		}
		if parent.GetIsBaseNode() {
			ret = append(ret, parent.(*Node))
		}

		// Recurse for children
		for _, n := range children.Nodes() {
			ret = append(ret, recursivelyGetBaseNodes(n)...)
		}

		return ret
	}

	rootNode, err := i.GetRootNode()
	if err != nil {
		fmt.Printf("BaseNodes Skipping rootnode:nil i:%v\n", i)
		return nil, false
	}

	nodes = recursivelyGetBaseNodes(rootNode)
	return nodes, len(nodes) > 0
}
