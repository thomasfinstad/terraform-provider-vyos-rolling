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
