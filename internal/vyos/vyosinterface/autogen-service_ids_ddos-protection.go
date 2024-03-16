// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func service_ids_ddosprotection() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*interfacedefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "service",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*interfacedefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "ids",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Intrusion Detection System"},
					}},
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						Node: []*interfacedefinition.Node{{
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "ddos-protection",
							OwnerAttr:    "${vyos_conf_scripts_dir}/service_ids_ddos-protection.py",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help:     []string{"FastNetMon detection and protection parameters"},
								Priority: []string{"731"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								Node: []*interfacedefinition.Node{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "node",
									},
									NodeNameAttr: "sflow",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Sflow settings"},
									}},
									Children: []*interfacedefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										LeafNode: []*interfacedefinition.LeafNode{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "listen-address",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Local IPv4 addresses to listen on"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*interfacedefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv4-address",
													}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv4",
													Description: "IPv4 address to listen for incoming connections",
												}},
												CompletionHelp: []*interfacedefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													Script: []string{"${vyos_completion_dir}/list_local_ips.sh --ipv4"},
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "port",
											DefaultValue: []string{"6343"},
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Port number used by connection"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*interfacedefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 1-65535",
													}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
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
									NodeNameAttr: "threshold",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Attack limits thresholds"},
									}},
									Children: []*interfacedefinition.Children{{
										XMLName: xml.Name{
											Local: "children",
										},
										Node: []*interfacedefinition.Node{{
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "node",
											},
											NodeNameAttr: "general",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"General threshold"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{
													Local: "children",
												},
												LeafNode: []*interfacedefinition.LeafNode{{
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "fps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Flows per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Flows per second",
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "mbps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Megabits per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Megabits per second",
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "pps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Packets per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Packets per second",
														}},
													}},
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "node",
											},
											NodeNameAttr: "tcp",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"TCP threshold"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{
													Local: "children",
												},
												LeafNode: []*interfacedefinition.LeafNode{{
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "fps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Flows per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Flows per second",
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "mbps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Megabits per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Megabits per second",
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "pps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Packets per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Packets per second",
														}},
													}},
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "node",
											},
											NodeNameAttr: "udp",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"UDP threshold"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{
													Local: "children",
												},
												LeafNode: []*interfacedefinition.LeafNode{{
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "fps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Flows per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Flows per second",
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "mbps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Megabits per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Megabits per second",
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "pps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Packets per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Packets per second",
														}},
													}},
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "node",
											},
											NodeNameAttr: "icmp",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"ICMP threshold"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{
													Local: "children",
												},
												LeafNode: []*interfacedefinition.LeafNode{{
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "fps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Flows per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Flows per second",
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "mbps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Megabits per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Megabits per second",
														}},
													}},
												}, {
													IsBaseNode: false,
													XMLName: xml.Name{
														Local: "leafNode",
													},
													NodeNameAttr: "pps",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Packets per second"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{
																Local: "constraint",
															},
															Validator: []*interfacedefinition.Validator{{
																XMLName: xml.Name{
																	Local: "validator",
																},
																NameAttr:     "numeric",
																ArgumentAttr: "--range 0-4294967294",
															}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "u32:0-4294967294",
															Description: "Packets per second",
														}},
													}},
												}},
											}},
										}},
									}},
								}},
								LeafNode: []*interfacedefinition.LeafNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "alert-script",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Path to fastnetmon alert script"},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "ban-time",
									DefaultValue: []string{"1900"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"How long we should keep an IP in blocked state"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 1-4294967294",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:1-4294967294",
											Description: "Time in seconds",
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "direction",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Direction for processing traffic"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(in|out)"},
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"in out"},
										}},
										Multi: []*interfacedefinition.Multi{{
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
									NodeNameAttr: "excluded-network",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Specify IPv4 and IPv6 networks which are going to be excluded from protection"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ipv4-prefix",
											}, {
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ipv6-prefix",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv4net",
											Description: "IPv4 prefix(es) to exclude",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv6net",
											Description: "IPv6 prefix(es) to exclude",
										}},
										Multi: []*interfacedefinition.Multi{{
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
									NodeNameAttr: "listen-interface",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Listen interface for mirroring traffic"},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_interfaces"},
										}},
										Multi: []*interfacedefinition.Multi{{
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
									NodeNameAttr: "mode",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Traffic capture mode"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(mirror|sflow)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "mirror",
											Description: "Listen to mirrored traffic",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "sflow",
											Description: "Capture sFlow flows",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"mirror sflow"},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "network",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Specify IPv4 and IPv6 networks which belong to you"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ipv4-prefix",
											}, {
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ipv6-prefix",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv4net",
											Description: "Your IPv4 prefix(es)",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv6net",
											Description: "Your IPv6 prefix(es)",
										}},
										Multi: []*interfacedefinition.Multi{{
											XMLName: xml.Name{
												Local: "multi",
											},
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