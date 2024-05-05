package main

import (
	"fmt"
	"log/slog"
	"math"
	"reflect"
	"sort"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

type duplicateIndex struct {
	// Count the one with the lower index as original, this is kind of arbetrary
	original int

	// Count the one with the higher index as duplicate, this is kind of arbetrary
	duplicate int
}

type duplicateInfo struct {
	leaf []duplicateIndex
	tag  []duplicateIndex
	node []duplicateIndex
}

func (o *duplicateInfo) addDup(nodeType string, original int, duplicate int) {
	switch nodeType {
	case "leaf":
		o.leaf = append(o.leaf, duplicateIndex{original: original, duplicate: duplicate})
	case "tag":
		o.tag = append(o.tag, duplicateIndex{original: original, duplicate: duplicate})
	case "node":
		o.node = append(o.node, duplicateIndex{original: original, duplicate: duplicate})
	default:
		panic(fmt.Sprintf("nodeType string must be one of [leaf, tag, node], got: '%s'", nodeType))
	}
}

func (o *duplicateInfo) isDup(nodeType string, index int) bool {
	switch nodeType {
	case "leaf":
		for _, idx := range o.leaf {
			if idx.duplicate == index {
				return true
			}
		}
	case "tag":
		for _, idx := range o.tag {
			if idx.duplicate == index {
				return true
			}
		}
	case "node":
		for _, idx := range o.node {
			if idx.duplicate == index {
				return true
			}
		}
	default:
		panic(fmt.Sprintf("nodeType string must be one of [leaf, tag, node], got: '%s'", nodeType))
	}

	return false
}

func populateDuplicateChildIndexes(node schemadefinition.NodeParent) (dups duplicateInfo) {
	children := node.GetChildren()

	for idx1, c1 := range children.LeafNodes() {
		for idx2, c2 := range children.LeafNodes() {
			if idx1 != idx2 {
				if c1.AbsNameR() == c2.AbsNameR() {
					originalIdx := int(math.Min(float64(idx1), float64(idx2)))
					duplicateIdx := int(math.Max(float64(idx1), float64(idx2)))
					slog.Info(fmt.Sprintf("Tag[%s] Has duplicate, indexes: %d & %d. ", c1.AbsName(), originalIdx, duplicateIdx))

					if dups.isDup("leaf", duplicateIdx) {
						slog.Debug(fmt.Sprintf("Index %d already marked as duplicate\n", duplicateIdx))
					} else {
						slog.Debug(fmt.Sprintf("Adding index %d to duplicate list\n", duplicateIdx))
						dups.addDup("leaf", originalIdx, duplicateIdx)
					}
				}
			}
		}
	}

	for idx1, c1 := range children.TagNodes() {
		for idx2, c2 := range children.TagNodes() {
			if idx1 != idx2 {
				if c1.AbsNameR() == c2.AbsNameR() {
					originalIdx := int(math.Min(float64(idx1), float64(idx2)))
					duplicateIdx := int(math.Max(float64(idx1), float64(idx2)))
					slog.Info(fmt.Sprintf("Tag[%s] Has duplicate, indexes: %d & %d. ", c1.AbsName(), originalIdx, duplicateIdx))

					if dups.isDup("tag", duplicateIdx) {
						slog.Debug(fmt.Sprintf("Index %d already marked as duplicate", duplicateIdx))
					} else {
						slog.Debug(fmt.Sprintf("Adding index %d to duplicate list", duplicateIdx))
						dups.addDup("tag", originalIdx, duplicateIdx)
					}
				}
			}
		}
	}

	for idx1, c1 := range children.Nodes() {
		for idx2, c2 := range children.Nodes() {
			if idx1 != idx2 {
				if c1.AbsNameR() == c2.AbsNameR() {
					originalIdx := int(math.Min(float64(idx1), float64(idx2)))
					duplicateIdx := int(math.Max(float64(idx1), float64(idx2)))
					slog.Info(fmt.Sprintf("Tag[%s] Has duplicate, indexes: %d & %d. ", c1.AbsName(), originalIdx, duplicateIdx))

					if dups.isDup("node", duplicateIdx) {
						slog.Debug(fmt.Sprintf("Index %d already marked as duplicate", duplicateIdx))
					} else {
						slog.Debug(fmt.Sprintf("Adding index %d to duplicate list", duplicateIdx))
						dups.addDup("node", originalIdx, duplicateIdx)
					}
				}
			}
		}
	}

	return dups
}

func mergeNodeParents(rootNode schemadefinition.NodeParent) {
	// Find and move duplicate data from TagNode children
	if reflect.ValueOf(rootNode).Kind() != reflect.Ptr {
		panic("node parameter must be a pointer for merge to have any effect")
	}

	siblings := rootNode.GetChildren()
	duplicateSibling := populateDuplicateChildIndexes(rootNode)

	// Merge children by moving children from dups into the original

	// TODO improve error logging during failed node merging
	//  milestone: 2
	//  save or show output in a more visible way to indicate what was
	//  lost due to being unable to merge

	// Duplicate Leaf
	for _, dupIndex := range duplicateSibling.leaf {
		slog.Warn(fmt.Sprintf("Can not merge leaf [%s] from: index: %d to index: %d", siblings.LeafNode[dupIndex.original].AbsName(), dupIndex.duplicate, dupIndex.original))
	}

	// Duplicate Tag
	for _, dupIndex := range duplicateSibling.tag {
		slog.Info(fmt.Sprintf("Merging tag [%s] from: index: %d to index: %d", siblings.TagNode[dupIndex.original].AbsName(), dupIndex.duplicate, dupIndex.original))

		childrenOfDuplicate := siblings.TagNode[dupIndex.duplicate].GetChildren()
		childrenOfOriginal := siblings.TagNode[dupIndex.original].GetChildren()

		// Add all children of duplicate to children of original
		childrenOfOriginal.LeafNode = append(childrenOfOriginal.LeafNode, childrenOfDuplicate.LeafNode...)
		childrenOfOriginal.TagNode = append(childrenOfOriginal.TagNode, childrenOfDuplicate.TagNode...)
		childrenOfOriginal.Node = append(childrenOfOriginal.Node, childrenOfDuplicate.Node...)
	}

	// Duplicate Node
	for _, dupIndex := range duplicateSibling.node {
		slog.Info(fmt.Sprintf("Merging node [%s] from: index: %d to index: %d", siblings.Node[dupIndex.original].AbsName(), dupIndex.duplicate, dupIndex.original))

		childrenOfDuplicate := siblings.Node[dupIndex.duplicate].GetChildren()
		childrenOfOriginal := siblings.Node[dupIndex.original].GetChildren()

		// Add all children of duplicate to children of original
		childrenOfOriginal.LeafNode = append(childrenOfOriginal.LeafNode, childrenOfDuplicate.LeafNode...)
		childrenOfOriginal.TagNode = append(childrenOfOriginal.TagNode, childrenOfDuplicate.TagNode...)
		childrenOfOriginal.Node = append(childrenOfOriginal.Node, childrenOfDuplicate.Node...)
	}

	// Delete duplicates from current nodes children lists

	// Simple sortable list of indexes to remove
	var leafToRemove []int
	var tagToRemove []int
	var nodeToRemove []int
	for _, dup := range duplicateSibling.leaf {
		leafToRemove = append(leafToRemove, dup.duplicate)
	}
	for _, dup := range duplicateSibling.tag {
		tagToRemove = append(tagToRemove, dup.duplicate)
	}
	for _, dup := range duplicateSibling.node {
		nodeToRemove = append(nodeToRemove, dup.duplicate)
	}

	// Remove from highest duplicate index to lowest to avoid reshuffling list indexes during the "purge"

	// Leaf
	sort.Sort(sort.Reverse(sort.IntSlice(leafToRemove)))
	for _, idx := range leafToRemove {
		slog.Info(fmt.Sprintf("Removing leaf: [%s] idx: %d", siblings.LeafNode[idx].AbsName(), idx))
		siblings.LeafNode = append(
			siblings.LeafNode[:idx],
			siblings.LeafNode[idx+1:]...,
		)
	}

	// Tag
	sort.Sort(sort.Reverse(sort.IntSlice(tagToRemove)))
	for _, idx := range tagToRemove {
		slog.Info(fmt.Sprintf("Removing tag: [%s] idx: %d", siblings.TagNode[idx].AbsName(), idx))
		siblings.TagNode = append(
			siblings.TagNode[:idx],
			siblings.TagNode[idx+1:]...,
		)
	}

	// Node
	sort.Sort(sort.Reverse(sort.IntSlice(nodeToRemove)))
	for _, idx := range nodeToRemove {
		slog.Info(fmt.Sprintf("Removing node: [%s] idx: %d", siblings.Node[idx].AbsName(), idx))
		siblings.Node = append(
			siblings.Node[:idx],
			siblings.Node[idx+1:]...,
		)
	}

	// Recurse
	for _, c := range siblings.TagNodes() {
		mergeNodeParents(c)
	}
	for _, c := range siblings.Nodes() {
		mergeNodeParents(c)
	}
}
