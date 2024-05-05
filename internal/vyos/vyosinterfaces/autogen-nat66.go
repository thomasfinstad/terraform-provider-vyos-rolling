// Code generated by /repo/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func nat66() schemadefinition.InterfaceDefinition {
	return schemadefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*schemadefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "nat66",
			OwnerAttr:    "${vyos_conf_scripts_dir}/nat66.py",
			Properties: []*schemadefinition.Properties{{
				XMLName: xml.Name{
					Local: "properties",
				},
				Help:     []string{"Network Prefix Translation (NAT66/NPTv6) parameters"},
				Priority: []string{"500"},
			}},
			Children: []*schemadefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*schemadefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "source",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Prefix mapping of IPv6 source address translation"},
					}},
					Children: []*schemadefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						TagNode: []*schemadefinition.TagNode{{
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "rule",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Source NAT66 rule number"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*schemadefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 1-999999",
									}},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-999999",
									Description: "Number for this rule",
								}},
								ConstraintErrorMessage: []string{"NAT66 rule number must be between 1 and 999999"},
							}},
							Children: []*schemadefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								Node: []*schemadefinition.Node{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "outbound-interface",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Match outbound-interface"},
									}},
									Children: []*schemadefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*schemadefinition.LeafNode{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "name",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Match interface"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"(\\!?)(bond|br|dum|en|ersp|eth|gnv|ifb|lan|l2tp|l2tpeth|macsec|peth|ppp|pppoe|pptp|sstp|tun|veth|vti|vtun|vxlan|wg|wlan|wwan)([0-9]?)(\\&?)(.+)?|(\\!?)lo"},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "vrf-name",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Interface name",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt&",
													Description: "Interface name with wildcard",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "!txt",
													Description: "Inverted interface name to match",
												}},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													Path:   []string{"vrf name"},
													Script: []string{"${vyos_completion_dir}/list_interfaces"},
												}},
											}},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "destination",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IPv6 destination prefix options"},
									}},
									Children: []*schemadefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*schemadefinition.LeafNode{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "prefix",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"IPv6 prefix to be translated"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-prefix",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-prefix-exclude",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6net",
													Description: "IPv6 prefix",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "!ipv6net",
													Description: "Match everything except the specified IPv6 prefix",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "port",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Port number"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "port-multi",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Named port (any name in /etc/services, e.g., http)",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-65535",
													Description: "Numeric IP port",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "start-end",
													Description: "Numbered port range (e.g. 1001-1005)",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Description: "\\n\\nMultiple destination ports can be specified as a comma-separated list.\\nThe whole list can also be negated using '!'.\\nFor example: '!22,telnet,http,123,1001-1005'",
												}},
											}},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "source",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IPv6 source prefix options"},
									}},
									Children: []*schemadefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*schemadefinition.LeafNode{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "prefix",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"IPv6 prefix to be translated"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-prefix",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-prefix-exclude",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6net",
													Description: "IPv6 prefix",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "!ipv6net",
													Description: "Match everything except the specified IPv6 prefix",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "port",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Port number"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "port-multi",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Named port (any name in /etc/services, e.g., http)",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-65535",
													Description: "Numeric IP port",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "start-end",
													Description: "Numbered port range (e.g. 1001-1005)",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Description: "\\n\\nMultiple destination ports can be specified as a comma-separated list.\\nThe whole list can also be negated using '!'.\\nFor example: '!22,telnet,http,123,1001-1005'",
												}},
											}},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "translation",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Translated IPv6 address options"},
									}},
									Children: []*schemadefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*schemadefinition.LeafNode{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "address",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"IPv6 address to translate to"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"(masquerade)"},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-address",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-prefix",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6",
													Description: "IPv6 address",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6net",
													Description: "IPv6 prefix",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "masquerade",
													Description: "NAT to the primary address of outbound-interface",
												}},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													List: []string{"masquerade"},
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "port",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Port number"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "port-range",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-65535",
													Description: "Numeric IP port",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "range",
													Description: "Numbered port range (e.g., 1001-1005)",
												}},
											}},
										}},
									}},
								}},
								LeafNode: []*schemadefinition.LeafNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "description",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Description"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{".{0,255}"},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Description",
										}},
										ConstraintErrorMessage: []string{"Description too long (limit 255 characters)"},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "disable",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Disable instance"},
										Valueless: []*schemadefinition.Valueless{{
											XMLName: xml.Name{
												Local: "valueless",
											},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "exclude",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Exclude packets matching this rule from NAT"},
										Valueless: []*schemadefinition.Valueless{{
											XMLName: xml.Name{
												Local: "valueless",
											},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "log",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Log packets hitting this rule"},
										Valueless: []*schemadefinition.Valueless{{
											XMLName: xml.Name{
												Local: "valueless",
											},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "protocol",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Protocol to match (protocol name, number, or \"all\")"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ip-protocol",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "all",
											Description: "All IP protocols",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "tcp_udp",
											Description: "Both TCP and UDP",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:0-255",
											Description: "IP protocol number",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "<protocol>",
											Description: "IP protocol name",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "!<protocol>",
											Description: "IP protocol name",
										}},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List:   []string{"all tcp_udp"},
											Script: []string{"${vyos_completion_dir}/list_protocols.sh"},
										}},
									}},
								}},
							}},
						}},
					}},
				}, {
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "destination",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Prefix mapping for IPv6 destination address translation"},
					}},
					Children: []*schemadefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						TagNode: []*schemadefinition.TagNode{{
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "rule",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Destination NAT66 rule number"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*schemadefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 1-999999",
									}},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-999999",
									Description: "Number for this rule",
								}},
								ConstraintErrorMessage: []string{"NAT66 rule number must be between 1 and 999999"},
							}},
							Children: []*schemadefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								Node: []*schemadefinition.Node{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "inbound-interface",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Match inbound-interface"},
									}},
									Children: []*schemadefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*schemadefinition.LeafNode{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "name",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Match interface"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"(\\!?)(bond|br|dum|en|ersp|eth|gnv|ifb|lan|l2tp|l2tpeth|macsec|peth|ppp|pppoe|pptp|sstp|tun|veth|vti|vtun|vxlan|wg|wlan|wwan)([0-9]?)(\\&?)(.+)?|(\\!?)lo"},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "vrf-name",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Interface name",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt&",
													Description: "Interface name with wildcard",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "!txt",
													Description: "Inverted interface name to match",
												}},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													Path:   []string{"vrf name"},
													Script: []string{"${vyos_completion_dir}/list_interfaces"},
												}},
											}},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "destination",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IPv6 destination prefix options"},
									}},
									Children: []*schemadefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*schemadefinition.LeafNode{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "address",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"IPv6 address or prefix to be translated"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-address",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-prefix",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-address-exclude",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-prefix-exclude",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6",
													Description: "IPv6 address",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6net",
													Description: "IPv6 prefix",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "!ipv6",
													Description: "Match everything except the specified IPv6 address",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "!ipv6net",
													Description: "Match everything except the specified IPv6 prefix",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "port",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Port number"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "port-multi",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Named port (any name in /etc/services, e.g., http)",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-65535",
													Description: "Numeric IP port",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "start-end",
													Description: "Numbered port range (e.g. 1001-1005)",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Description: "\\n\\nMultiple destination ports can be specified as a comma-separated list.\\nThe whole list can also be negated using '!'.\\nFor example: '!22,telnet,http,123,1001-1005'",
												}},
											}},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "source",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IPv6 source prefix options"},
									}},
									Children: []*schemadefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*schemadefinition.LeafNode{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "address",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"IPv6 address or prefix to be translated"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-address",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-prefix",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-address-exclude",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-prefix-exclude",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6",
													Description: "IPv6 address",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6net",
													Description: "IPv6 prefix",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "!ipv6",
													Description: "Match everything except the specified IPv6 address",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "!ipv6net",
													Description: "Match everything except the specified IPv6 prefix",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "port",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Port number"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "port-multi",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Named port (any name in /etc/services, e.g., http)",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-65535",
													Description: "Numeric IP port",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "start-end",
													Description: "Numbered port range (e.g. 1001-1005)",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Description: "\\n\\nMultiple destination ports can be specified as a comma-separated list.\\nThe whole list can also be negated using '!'.\\nFor example: '!22,telnet,http,123,1001-1005'",
												}},
											}},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "translation",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Translated IPv6 address options"},
									}},
									Children: []*schemadefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*schemadefinition.LeafNode{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "address",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"IPv6 address or prefix to translate to"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-address",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-prefix",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6",
													Description: "IPv6 address",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6net",
													Description: "IPv6 prefix",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "port",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Port number"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "port-range",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-65535",
													Description: "Numeric IP port",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "range",
													Description: "Numbered port range (e.g., 1001-1005)",
												}},
											}},
										}},
									}},
								}},
								LeafNode: []*schemadefinition.LeafNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "description",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Description"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{".{0,255}"},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Description",
										}},
										ConstraintErrorMessage: []string{"Description too long (limit 255 characters)"},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "disable",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Disable instance"},
										Valueless: []*schemadefinition.Valueless{{
											XMLName: xml.Name{
												Local: "valueless",
											},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "exclude",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Exclude packets matching this rule from NAT"},
										Valueless: []*schemadefinition.Valueless{{
											XMLName: xml.Name{
												Local: "valueless",
											},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "log",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"NAT66 rule logging"},
										Valueless: []*schemadefinition.Valueless{{
											XMLName: xml.Name{
												Local: "valueless",
											},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "protocol",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Protocol to match (protocol name, number, or \"all\")"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ip-protocol",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "all",
											Description: "All IP protocols",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "tcp_udp",
											Description: "Both TCP and UDP",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:0-255",
											Description: "IP protocol number",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "<protocol>",
											Description: "IP protocol name",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "!<protocol>",
											Description: "IP protocol name",
										}},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List:   []string{"all tcp_udp"},
											Script: []string{"${vyos_completion_dir}/list_protocols.sh"},
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
