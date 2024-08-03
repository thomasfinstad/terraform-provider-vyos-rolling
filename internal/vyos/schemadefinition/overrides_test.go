package schemadefinition_test

import (
	"slices"
	"testing"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
	vyosinterface "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/vyosinterfaces"
)

func TestTagNodeOverride(t *testing.T) {
	var baseNodes []*schemadefinition.TagNode
	for _, i := range vyosinterface.GetInterfaces() {
		if n, ok := i.BaseTagNodes(); ok {
			baseNodes = append(baseNodes, n...)
		} else {
			t.Logf("No base tag nodes returned for %s", i.Node[0].NodeNameAttr)
		}
	}

	var overrideCheck func(o schemadefinition.BaseNodeOverride, n schemadefinition.NodeParent) (mergedSuccessfully bool)
	overrideCheck = func(o schemadefinition.BaseNodeOverride, n schemadefinition.NodeParent) (mergedSuccessfully bool) {
		if slices.Equal(n.AbsName(), o.Full()) && !n.GetIsBaseNode() {
			return true
		}

		for _, c := range n.GetChildren().TagNode {
			if overrideCheck(o, c) {
				return true
			}
		}

		for _, c := range n.GetChildren().Node {
			if overrideCheck(o, c) {
				return true
			}
		}

		return false
	}

	for _, o := range schemadefinition.BaseNodeOverrides {
		merged := false
		for _, n := range baseNodes {
			if overrideCheck(o, n) {
				merged = true
				break
			}
		}
		if !merged {
			t.Errorf("Node not merged correctly: [%v]", o)
		}
	}
}
