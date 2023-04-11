package interfacedefinition

import (
	"fmt"
)

/*
Contains structs to use with vyos schemas
*/

// GetRootNode returns the top level node from the interface definition
func (i *InterfaceDefinition) GetRootNode() (*Node, error) {
	if len(i.Node) > 0 {
		return i.Node[0], nil
	}

	return &Node{}, fmt.Errorf("no root node found under interface")
}

// GetAncestory returns a list of (tag)nodes based on the pointer address of n
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
func (i *InterfaceDefinition) BaseTagNodes() []*TagNode {
	var n NodeParent
	var t []*TagNode

	n, err := i.GetRootNode()
	if err != nil {
		fmt.Printf("BaseTagNodes Skipping rootnode:nil i:%v\n", i)
		return t
	}

	if n.BaseName() == "" {
		fmt.Printf("BaseTagNodes Skipping root node has no name. i:%v\n", i)
		return t
	}

	// TODO Investigate need to convert to recursive inline func and return full list of tag nodes, to ensure not only first tagnodes are returned
	fmt.Printf("BaseTagNodes ")
	for {
		fmt.Printf("[%s]", n.BaseName())

		children := n.GetChildren()
		if children == nil {
			fmt.Printf("[%s] Skipping c:nil\n", i.Node[0].NodeNameAttr)
			break
		}
		if len(children.TagNode) >= 1 {
			t = children.TagNode
			fmt.Printf(":%d\n", len(t))
			break
		} else if len(children.Node) == 1 {
			n = children.Node[0]
		} else {
			fmt.Printf("[%s] Skipping n:%d t:%d l:%d\n", i.Node[0].NodeNameAttr, len(n.GetChildren().Node), len(n.GetChildren().TagNode), len(n.GetChildren().LeafNode))
			break
		}
	}

	return t
}
