// Code generated by /workspaces/terraform-provider-vyos/tools/build-resource-terraform-schemas/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/provider/vyos/schema/interfacedefinition"
)

func protocolsstaticarp() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{Local: "interfaceDefinition"},
		Node: []*interfacedefinition.Node{{
			XMLName: xml.Name{Local: "node"}, NodeNameAttr: "protocols",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{Local: "children"},
				Node: []*interfacedefinition.Node{{
					XMLName: xml.Name{Local: "node"}, NodeNameAttr: "static",
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{Local: "children"},
						Node: []*interfacedefinition.Node{{
							XMLName: xml.Name{Local: "node"}, NodeNameAttr: "arp", OwnerAttr: "${vyos_conf_scripts_dir}/arp.py",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{Local: "properties"},
								Help:    []string{"Static ARP translation"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{Local: "children"},
								TagNode: []*interfacedefinition.TagNode{{
									XMLName: xml.Name{Local: "tagNode"}, NodeNameAttr: "interface",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{Local: "properties"},
										Help:    []string{"Interface configuration"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName:   xml.Name{Local: "constraint"},
											Regex:     []string{"(bond|br|dum|en|ersp|eth|gnv|ifb|lan|l2tp|l2tpeth|macsec|peth|ppp|pppoe|pptp|sstp|tun|veth|vti|vtun|vxlan|wg|wlan|wwan)[0-9]+(.\\d+)?|lo"},
											Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "file-path --lookup-path /sys/class/net --directory"}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "txt", Description: "Interface name"}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{Local: "completionHelp"},
											Script:  []string{"${vyos_completion_dir}/list_interfaces"},
										}},
									}},
									Children: []*interfacedefinition.Children{{
										XMLName: xml.Name{Local: "children"},
										TagNode: []*interfacedefinition.TagNode{{
											XMLName: xml.Name{Local: "tagNode"}, NodeNameAttr: "address",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"IP address for static ARP entry"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName:   xml.Name{Local: "constraint"},
													Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv4-address"}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv4", Description: "IPv4 destination address"}},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{Local: "children"},
												LeafNode: []*interfacedefinition.LeafNode{{
													XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "description",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"Description"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{Local: "constraint"},
															Regex:   []string{"[[:ascii:]]{0,256}"},
														}},
														ValueHelp:              []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "txt", Description: "Description"}},
														ConstraintErrorMessage: []string{"Description too long (limit 256 characters)"},
													}},
												}, {
													XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "mac",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"Media Access Control (MAC) address"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName:   xml.Name{Local: "constraint"},
															Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "mac-address"}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "macaddr", Description: "Hardware (MAC) address"}},
													}},
												}},
											}},
										}},
									}},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}
