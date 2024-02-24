package interfacedefinition

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
	}

	ret := i.Node[0]

	ret.InformLinage()

	return ret, nil
}

// BaseTagNodes returns highest level TagNodes
func (i *InterfaceDefinition) BaseTagNodes() (tagNodes []*TagNode, ok bool) {
	var recurse func(NodeParent) []*TagNode
	recurse = func(parent NodeParent) []*TagNode {
		var ret []*TagNode
		children := parent.GetChildren()

		if children == nil {
			fmt.Printf("[%s] Skipping children:nil\n", parent.BaseName())
			return ret
		}

		for _, n := range children.Node {
			ret = append(ret, recurse(n)...)
		}

		ret = append(ret, children.TagNode...)

		return ret
	}

	rootNode, err := i.GetRootNode()
	if err != nil {
		fmt.Printf("BaseTagNodes Skipping rootnode:nil i:%v\n", i)
		return nil, false
	}

	tagNodes = recurse(rootNode)
	ok = len(tagNodes) > 0

	// Let tagnode know it is the basenode
	for _, tagNode := range tagNodes {
		tagNode.IsBaseNode = true
	}

	return tagNodes, ok
}

// TagNodes returns all TagNodes (recursively)
func (i *InterfaceDefinition) TagNodes() (tagNodes []*TagNode, ok bool) {
	var recurse func(NodeParent) []*TagNode
	recurse = func(parent NodeParent) []*TagNode {
		var ret []*TagNode
		children := parent.GetChildren()

		if children == nil {
			fmt.Printf("[%s] Skipping children:nil\n", parent.BaseName())
			return ret
		}

		// Add own tag node children
		ret = append(ret, children.TagNode...)

		// Recurse for children
		for _, n := range children.Nodes() {
			ret = append(ret, recurse(n)...)
		}

		for _, t := range children.TagNodes() {
			ret = append(ret, recurse(t)...)
		}

		return ret
	}

	rootNode, err := i.GetRootNode()
	if err != nil {
		fmt.Printf("BaseTagNodes Skipping rootnode:nil i:%v\n", i)
		return nil, false
	}

	tagNodes = recurse(rootNode)
	ok = len(tagNodes) > 0

	// Let tagnode know it is the basenode
	for _, tagNode := range tagNodes {
		tagNode.IsBaseNode = true
	}

	return tagNodes, ok
}

// Nodes returns all Nodes that contain a LeafNode that is not inside a TagNode
func (i *InterfaceDefinition) Nodes() (nodes []*Node, ok bool) {
	var recurse func(NodeParent) []*Node
	recurse = func(parent NodeParent) []*Node {
		var ret []*Node
		children := parent.GetChildren()

		if children == nil {
			fmt.Printf("[%s] Skipping children:nil\n", parent.BaseName())
			return ret
		}

		// Add self is containing LeafNodes
		if len(children.LeafNode) > 0 {
			ret = append(ret, parent.(*Node))
		}

		// Recurse for children
		for _, n := range children.Nodes() {
			ret = append(ret, recurse(n)...)
		}

		return ret
	}

	rootNode, err := i.GetRootNode()
	if err != nil {
		fmt.Printf("BaseNodes Skipping rootnode:nil i:%v\n", i)
		return nil, false
	}

	nodes = recurse(rootNode)
	ok = len(nodes) > 0

	// Let node know it is the basenode (unlikely to be useful)
	for _, node := range nodes {
		node.IsBaseNode = true
	}

	return nodes, ok
}
