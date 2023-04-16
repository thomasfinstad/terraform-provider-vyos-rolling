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

// GetAncestory returns a list of (tag)nodes based on the pointer address of n
// TODO replace with parent based traversal
func (i *InterfaceDefinition) GetAncestory(target NodeParent) []NodeParent {
	rootNode, err := i.GetRootNode()
	die(err)
	targetAddr := fmt.Sprint(target)

	if fmt.Sprint(rootNode) == targetAddr {
		return []NodeParent{rootNode}
	}

	var recurse func(n NodeParent) []NodeParent
	recurse = func(n NodeParent) []NodeParent {
		for _, child := range n.GetChildren().Node {
			// If self pointer address matches target pointer address
			if fmt.Sprint(child) == targetAddr {
				// Return self
				return []NodeParent{child}
			}

			// If a match was found downstream
			if ret := recurse(child); len(ret) > 0 {
				// Insert self and return
				return append([]NodeParent{child}, ret...)
			}
		}

		for _, child := range n.GetChildren().TagNode {
			// If self pointer address matches target pointer address
			if fmt.Sprint(child) == targetAddr {
				// Return self
				return []NodeParent{child}
			}

			// If a match was found downstream
			if ret := recurse(child); len(ret) > 0 {
				// Insert self and return
				return append([]NodeParent{child}, ret...)
			}
		}

		return nil
	}

	ret := recurse(rootNode)
	if ret != nil {
		ret = append([]NodeParent{rootNode}, ret...)
	}

	return ret
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
		return tagNodes, false
	}

	tagNodes = recurse(rootNode)
	ok = len(tagNodes) > 0

	return tagNodes, ok
}
