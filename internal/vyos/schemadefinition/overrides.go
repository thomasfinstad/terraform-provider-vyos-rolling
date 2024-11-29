package schemadefinition

import "fmt"

// Allows tag nodes to be configured to not become its own
// resource by keeping it as a child config in the parent.
// This is needed as some nodes requires child tag nodes
// to be configured
//
// This is hopefully a temporary measure until
// this feature can be implemented in the schemas
// https://vyos.dev/T5083

var MergeBaseNodeOverrides = []MergeBaseNodeOverride{
	// container
	{
		from: []string{"container", "name"},
		to:   []string{"network"},
	},

	// interfaces, openvpn
	{
		from: []string{"interfaces", "openvpn"},
		to:   []string{"local-address"},
	},
	{
		from: []string{"interfaces", "openvpn"},
		to:   []string{"server", "client"},
	},

	// service, ntp
	{
		from: []string{"service", "ntp"},
		to:   []string{"server"},
	},
	{
		from: []string{"service", "ntp"},
		to:   []string{"allow-client", "address"},
	},

	// service, dhcp-server
	{
		from: []string{"service", "dhcp-server", "shared-network-name"},
		to:   []string{"subnet", "range"},
	},

	// service, conntrack-sync
	{
		from: []string{"service", "conntrack-sync"},
		to:   []string{"failover-mechanism", "vrrp", "sync-group"},
	},
	{
		from: []string{"service", "conntrack-sync"},
		to:   []string{"interface"},
	},
}

type MergeBaseNodeOverride struct {
	from []string
	to   []string
}

func (n MergeBaseNodeOverride) Full() []string {
	return append(n.from, n.to...)
}

func (n MergeBaseNodeOverride) From() []string {
	return n.from
}

// ApplyOverrides mutates the IsBaseNode setting based on BaseNodeOverrides
func ApplyOverrides(nodes []NodeParent) {
	for _, node := range nodes {
		children := node.GetChildren()

		if children == nil {
			fmt.Printf("[%s] Skipping children:nil\n", node.BaseName())
			return
		}

		// Testing for (merge)override
		if np, ok := children.GetNodeParents(); ok {
		NodesLoop:
			for _, n := range np {
			MergeBaseNodeOverridesLoop:
				for _, override := range MergeBaseNodeOverrides {
					// fmt.Print("Checking override for Node: ", n.AbsName(), "\tOverride: ", override.Full(), "\t")

					// skip any non-matches
					for i := range min(len(n.AbsName()), len(override.Full())) {
						if n.AbsName()[i] != override.Full()[i] {
							// fmt.Print("wrong subnode.", "\n")
							continue MergeBaseNodeOverridesLoop
						}
					}

					// recurse and skip if node name is shorter than "from" override scope
					if len(n.AbsName()) <= len(override.from) {
						// fmt.Print("before override scope start.", "\n")
						if np, ok := n.GetChildren().GetNodeParents(); ok {
							ApplyOverrides(np)
						}
						continue MergeBaseNodeOverridesLoop
					}

					// skip if node name is longer than "to" override scope
					if len(n.AbsName()) > len(override.Full()) {
						// fmt.Print("byond override scope end.", "\n")
						continue MergeBaseNodeOverridesLoop
					}

					// fmt.Print("override test result:\t")

					// check each if each name segment is in the override
					for i := 0; i < len(n.AbsName()); i++ {

						// skip if not matching
						if n.AbsName()[i] != override.Full()[i] {
							// fmt.Print("mismatch at index: ", i, " subnode: ", n.AbsName()[i], "\n")
							continue MergeBaseNodeOverridesLoop
						}

						// if we are on the last segment of the node name and reach this point then it should be merged into parent
						if i+1 == len(n.AbsName()) {
							fmt.Print("\tOverride (merge base node) match found, Node: ", n.AbsName(), "\tOverride: (from ", override.from, ") (to ", override.Full(), ")\n")
							n.SetIsBaseNode(false)
							continue NodesLoop
						}
					}
				}
			}
		}
	}
}
