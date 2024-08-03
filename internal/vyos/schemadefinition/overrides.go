package schemadefinition

// Allows tag nodes to be configured to not become its own
// resource by keeping it as a child config in the parent.
// This is needed as some nodes requires child tag nodes
// to be configured
//
// This is hopefully a temporary measure until
// this feature can be implemented in the schemas
// https://vyos.dev/T5083

var BaseNodeOverrides = []BaseNodeOverride{
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
}

type BaseNodeOverride struct {
	from []string
	to   []string
}

func (n BaseNodeOverride) Full() []string {
	return append(n.from, n.to...)
}
