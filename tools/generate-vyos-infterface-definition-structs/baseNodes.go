package main

import (
	"fmt"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func markBaseTagNodes(rootNode schemadefinition.NodeParent) {
	children := rootNode.GetChildren()

	if children == nil {
		fmt.Printf("[%s] Skipping children:nil\n", rootNode.BaseName())
		return
	}

	// Recurse tagnode children
	for _, n := range children.TagNodes() {
		n.SetIsBaseNode(true)
		markBaseTagNodes(n)
	}

	// Recurse node children
	for _, n := range children.Nodes() {
		markBaseTagNodes(n)
	}

	// Apply overrides to all sub-nodes that are themselves parents
	if np, ok := children.GetNodeParents(); ok {
		schemadefinition.ApplyOverrides(np)
	}
}

func markBaseNodes(rootNode schemadefinition.NodeParent) {
	children := rootNode.GetChildren()

	if children == nil {
		fmt.Printf("[%s] Skipping children:nil\n", rootNode.BaseName())
		return
	}

	// Add self is containing LeafNodes
	if len(children.LeafNodes()) > 0 {
		rootNode.SetIsBaseNode(true)
	} else {
		rootNode.SetIsBaseNode(false)
	}

	// Recurse for children
	for _, n := range children.Nodes() {
		markBaseNodes(n)
	}

	// Apply overrides to all sub-nodes that are themselves parents
	if np, ok := children.GetNodeParents(); ok {
		schemadefinition.ApplyOverrides(np)
	}
}
