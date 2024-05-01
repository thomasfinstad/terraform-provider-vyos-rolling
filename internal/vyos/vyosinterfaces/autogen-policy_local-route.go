// Code generated by /repo/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func policy_localroute() schemadefinition.InterfaceDefinition {
	return schemadefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*schemadefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "policy",
			Children: []*schemadefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*schemadefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "local-route",
					OwnerAttr:    "${vyos_conf_scripts_dir}/policy_local-route.py",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"IPv4 policy route of local traffic"},
						Priority: []string{"500"},
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
								Help: []string{"Policy local-route rule set number"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*schemadefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 1-32765",
									}},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-32765",
									Description: "Local-route rule number (1-32765)",
								}},
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
									NodeNameAttr: "set",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Packet modifications"},
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
											NodeNameAttr: "table",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Routing table to forward packet with"},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-200",
													Description: "Table number",
												}},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													List: []string{"main"},
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
										Help: []string{"Source parameters"},
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
												Help: []string{"IPv4 address or prefix"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv4-address",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ip-prefix",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv4",
													Description: "Address to match against",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv4net",
													Description: "Prefix to match against",
												}},
												Multi: []*schemadefinition.Multi{{
													XMLName: xml.Name{
														Local: "multi",
													},
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
												Help: []string{"Port number used by connection"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 1-65535",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-65535",
													Description: "Numeric IP port",
												}},
												ConstraintErrorMessage: []string{"Port number must be in range 1 to 65535"},
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
										Help: []string{"Destination parameters"},
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
												Help: []string{"IPv4 address or prefix"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv4-address",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ip-prefix",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv4",
													Description: "Address to match against",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv4net",
													Description: "Prefix to match against",
												}},
												Multi: []*schemadefinition.Multi{{
													XMLName: xml.Name{
														Local: "multi",
													},
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
												Help: []string{"Port number used by connection"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 1-65535",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-65535",
													Description: "Numeric IP port",
												}},
												ConstraintErrorMessage: []string{"Port number must be in range 1 to 65535"},
											}},
										}},
									}},
								}},
								LeafNode: []*schemadefinition.LeafNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "fwmark",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Match fwmark value"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 1-2147483647",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:1-2147483647",
											Description: "Address to match against",
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
										Help: []string{"Protocol to match (protocol name or number)"},
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
											Format:      "u32:0-255",
											Description: "IP protocol number",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "<protocol>",
											Description: "IP protocol name",
										}},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_protocols.sh"},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "inbound-interface",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Inbound Interface"},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_interfaces"},
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
					NodeNameAttr: "local-route6",
					OwnerAttr:    "${vyos_conf_scripts_dir}/policy_local-route.py",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"IPv6 policy route of local traffic"},
						Priority: []string{"500"},
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
								Help: []string{"IPv6 policy local-route rule set number"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*schemadefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 1-32765",
									}},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-32765",
									Description: "Local-route rule number (1-32765)",
								}},
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
									NodeNameAttr: "set",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Packet modifications"},
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
											NodeNameAttr: "table",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Routing table to forward packet with"},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-200",
													Description: "Table number",
												}},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													List: []string{"main"},
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
										Help: []string{"Source parameters"},
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
												Help: []string{"IPv6 address or prefix"},
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
													Description: "Address to match against",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6net",
													Description: "Prefix to match against",
												}},
												Multi: []*schemadefinition.Multi{{
													XMLName: xml.Name{
														Local: "multi",
													},
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
												Help: []string{"Port number used by connection"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 1-65535",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-65535",
													Description: "Numeric IP port",
												}},
												ConstraintErrorMessage: []string{"Port number must be in range 1 to 65535"},
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
										Help: []string{"Destination parameters"},
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
												Help: []string{"IPv6 address or prefix"},
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
													Description: "Address to match against",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6net",
													Description: "Prefix to match against",
												}},
												Multi: []*schemadefinition.Multi{{
													XMLName: xml.Name{
														Local: "multi",
													},
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
												Help: []string{"Port number used by connection"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 1-65535",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-65535",
													Description: "Numeric IP port",
												}},
												ConstraintErrorMessage: []string{"Port number must be in range 1 to 65535"},
											}},
										}},
									}},
								}},
								LeafNode: []*schemadefinition.LeafNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "fwmark",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Match fwmark value"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 1-2147483647",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:1-2147483647",
											Description: "Address to match against",
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
										Help: []string{"Protocol to match (protocol name or number)"},
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
											Format:      "u32:0-255",
											Description: "IP protocol number",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "<protocol>",
											Description: "IP protocol name",
										}},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_protocols.sh"},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "inbound-interface",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Inbound Interface"},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_interfaces"},
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
