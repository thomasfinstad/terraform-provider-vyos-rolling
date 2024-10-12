package schemadefinition_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
	vyosinterface "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/vyosinterfaces"
)

func TestTagNodeOverride(t *testing.T) {

	// var baseNodes []schemadefinition.NodeParent
	// for _, i := range vyosinterface.GetInterfaces() {
	// 	if btn, ok := i.BaseTagNodes(); ok {
	// 		for _, n := range btn {
	// 			baseNodes = append(baseNodes, n)
	// 		}
	// 	} else {
	// 		t.Logf("No base TagNodes returned for %s", i.Node[0].NodeNameAttr)
	// 	}
	// 	if bn, ok := i.BaseNodes(); ok {
	// 		for _, n := range bn {
	// 			baseNodes = append(baseNodes, n)
	// 		}
	// 	} else {
	// 		t.Logf("No base Nodes returned for %s", i.Node[0].NodeNameAttr)
	// 	}
	// }

	var rootNodes []*schemadefinition.Node
	for _, iface := range vyosinterface.GetInterfaces() {
		if rn, err := iface.GetRootNode(); err == nil {
			rootNodes = append(rootNodes, rn)
		}
	}

	t.Logf("Total root nodes: %d", len(rootNodes))

	// BaseNodeLoop:
	for _, rootNode := range rootNodes {
		t.Logf("checking rootNode: %v", rootNode.AbsName())
	OverrideLoop:
		for _, override := range schemadefinition.MergeBaseNodeOverrides {
			t.Logf("checking override: %v", override.Full())
			child, err := rootNode.GetChildAbsPath(override.Full())
			if err != nil {
				t.Logf("%v", err)
				continue OverrideLoop
			}

			t.Logf("matching child: %v", child.AbsName())
			c := child
		ParentWalkLoop:
			for range len(c.AbsName()) {
				t.Logf("checking: %v", c.AbsName())
				switch c := c.(type) {
				case schemadefinition.NodeParent:
					if c.GetIsBaseNode() {
						t.Errorf("should not be baseNode: %v", c.AbsName())
					} else {
						t.Logf("node correctly marked as non-baseNode: %v", c.AbsName())
					}
				case *schemadefinition.LeafNode:
					t.Logf("node is a leafnode: %v", c.AbsName())
					continue ParentWalkLoop
				default:
					panic(fmt.Sprintf("unreachable. %#v %#v", override, c))
				}

				c = c.GetParent()
				if slices.Equal(c.AbsName(), override.From()) {
					t.Logf("Node %v is at start of override scope %v and is allowed to be a baseNode", c.AbsName(), override)
					continue OverrideLoop
				}
			}
		}
	}
	// t.Errorf("Artificial failure to view log lines")
}
