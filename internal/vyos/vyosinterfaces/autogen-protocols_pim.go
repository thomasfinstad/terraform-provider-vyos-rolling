// Code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func protocols_pim() schemadefinition.InterfaceDefinition {
	return schemadefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*schemadefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "protocols",
			Children: []*schemadefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*schemadefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "pim",
					OwnerAttr:    "${vyos_conf_scripts_dir}/protocols_pim.py",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"Protocol Independent Multicast (PIM) and IGMP"},
						Priority: []string{"400"},
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
							NodeNameAttr: "ecmp",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Enable PIM ECMP"},
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
									NodeNameAttr: "rebalance",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Enable PIM ECMP Rebalance"},
										Valueless: []*schemadefinition.Valueless{{
											XMLName: xml.Name{
												Local: "valueless",
											},
										}},
									}},
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "igmp",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Internet Group Management Protocol (IGMP) options"},
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
									NodeNameAttr: "watermark-warning",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Configure group limit for watermark warning"},
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
											Description: "Group count to generate watermark warning",
										}},
									}},
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "register-accept-list",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Only accept registers from a specific source prefix list"},
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
									NodeNameAttr: "prefix-list",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Prefix-list to use"},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Prefix-list to apply (IPv4)",
										}},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Path: []string{"policy prefix-list"},
										}},
									}},
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "rp",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Rendezvous Point"},
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
									NodeNameAttr: "address",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Rendezvous Point address"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ipv4-address",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv4",
											Description: "Rendezvous Point address",
										}},
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
											NodeNameAttr: "group",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Group Address range"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
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
													Format:      "ipv4net",
													Description: "Group Address range RFC 3171",
												}},
												Multi: []*schemadefinition.Multi{{
													XMLName: xml.Name{
														Local: "multi",
													},
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
									NodeNameAttr: "keep-alive-timer",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Keep alive Timer"},
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
											Description: "Keep alive Timer in seconds",
										}},
									}},
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "spt-switchover",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Shortest-path tree (SPT) switchover"},
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
									NodeNameAttr: "infinity-and-beyond",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Never switch to SPT Tree"},
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
											NodeNameAttr: "prefix-list",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Prefix-list to use"},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Prefix-list to apply (IPv4)",
												}},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													Path: []string{"policy prefix-list"},
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
							NodeNameAttr: "ssm",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Source-Specific Multicast"},
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
									NodeNameAttr: "prefix-list",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Prefix-list to use"},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Prefix-list to apply (IPv4)",
										}},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Path: []string{"policy prefix-list"},
										}},
									}},
								}},
							}},
						}},
						TagNode: []*schemadefinition.TagNode{{
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "interface",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"PIM interface"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(bond|br|dum|en|ersp|eth|gnv|ifb|ipoe|lan|l2tp|l2tpeth|macsec|peth|ppp|pppoe|pptp|sstp|sstpc|tun|veth|vti|vtun|vxlan|wg|wlan|wwan)[0-9]+(.\\d+)?|lo"},
									Validator: []*schemadefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr: "file-path --lookup-path /sys/class/net --directory",
									}},
								}},
								CompletionHelp: []*schemadefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									Script: []string{"${vyos_completion_dir}/list_interfaces"},
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
									NodeNameAttr: "bfd",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Enable Bidirectional Forwarding Detection (BFD)"},
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
											NodeNameAttr: "profile",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Use settings from BFD profile"},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "BFD profile name",
												}},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													Path: []string{"protocols bfd profile"},
												}},
											}},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "igmp",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Internet Group Management Protocol (IGMP) options"},
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
											NodeNameAttr: "join",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"IGMP join multicast group"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv4-address",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv4",
													Description: "Multicast group address",
												}},
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
													NodeNameAttr: "source-address",
													Properties: []*schemadefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"IPv4 source address used to initiate connection"},
														Constraint: []*schemadefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*schemadefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr: "ipv4-address",
															}},
														}},
														ValueHelp: []*schemadefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "ipv4",
															Description: "IPv4 source address",
														}},
														CompletionHelp: []*schemadefinition.CompletionHelp{{
															XMLName: xml.Name{
																Local: "completionHelp",
															},
															Script: []string{"${vyos_completion_dir}/list_local_ips.sh --ipv4"},
														}},
														Multi: []*schemadefinition.Multi{{
															XMLName: xml.Name{
																Local: "multi",
															},
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
											NodeNameAttr: "query-interval",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"IGMP host query interval"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 1-1800",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:1-1800",
													Description: "Query interval in seconds",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "query-max-response-time",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"IGMP max query response time"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 10-250",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:10-250",
													Description: "Query response value in deci-seconds",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "version",
											DefaultValue: []string{"3"},
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Interface IGMP version"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 2-3",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "2",
													Description: "IGMP version 2",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "3",
													Description: "IGMP version 3",
												}},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													List: []string{"2 3"},
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
									NodeNameAttr: "no-bsm",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Do not process bootstrap messages"},
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
									NodeNameAttr: "no-unicast-bsm",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Do not process unicast bootstrap messages"},
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
									NodeNameAttr: "dr-priority",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Designated router election priority"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 1-4294967295",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:1-4294967295",
											Description: "DR Priority",
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "hello",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Hello Interval"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 1-180",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:1-180",
											Description: "Hello Interval in seconds",
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "passive",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Disable sending and receiving PIM control packets on the interface"},
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
									NodeNameAttr: "source-address",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IPv4 source address used to initiate connection"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ipv4-address",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv4",
											Description: "IPv4 source address",
										}},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_local_ips.sh --ipv4"},
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
							NodeNameAttr: "join-prune-interval",
							DefaultValue: []string{"60"},
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Join prune send interval"},
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
									Description: "Interval in seconds",
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "keep-alive-timer",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Keep alive Timer"},
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
									Description: "Keep alive Timer in seconds",
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "packets",
							DefaultValue: []string{"3"},
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Packets to process at once"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*schemadefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 1-255",
									}},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-255",
									Description: "Number of packets",
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "register-suppress-time",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Register suppress timer"},
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
									Description: "Timer in seconds",
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "no-v6-secondary",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Disable IPv6 secondary address in hello packets"},
								Valueless: []*schemadefinition.Valueless{{
									XMLName: xml.Name{
										Local: "valueless",
									},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}
