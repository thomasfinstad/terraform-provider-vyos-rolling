// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func protocols_segmentrouting() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*interfacedefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "protocols",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*interfacedefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "segment-routing",
					OwnerAttr:    "${vyos_conf_scripts_dir}/protocols_segment-routing.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"Segment Routing"},
						Priority: []string{"900"},
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
							NodeNameAttr: "srv6",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Segment-Routing SRv6 configuration"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								TagNode: []*interfacedefinition.TagNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "tagNode",
									},
									NodeNameAttr: "locator",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Segment Routing SRv6 locator"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"[-_a-zA-Z0-9]+"},
										}},
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
											NodeNameAttr: "behavior-usid",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Set SRv6 behavior uSID"},
												Valueless: []*interfacedefinition.Valueless{{
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
											NodeNameAttr: "prefix",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"SRv6 locator prefix"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*interfacedefinition.Validator{{
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
													Format:      "ipv6net",
													Description: "SRv6 locator prefix",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "block-len",
											DefaultValue: []string{"40"},
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Configure SRv6 locator block length in bits"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*interfacedefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 16-64",
													}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:16-64",
													Description: "Specify SRv6 locator block length in bits",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "func-bits",
											DefaultValue: []string{"16"},
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Configure SRv6 locator function length in bits"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*interfacedefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 0-64",
													}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:0-64",
													Description: "Specify SRv6 locator function length in bits",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "node-len",
											DefaultValue: []string{"24"},
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Configure SRv6 locator node length in bits"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*interfacedefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 16-64",
													}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:16-64",
													Description: "Configure SRv6 locator node length in bits",
												}},
											}},
										}},
									}},
								}},
							}},
						}},
						TagNode: []*interfacedefinition.TagNode{{
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "interface",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Interface specific Segment Routing options"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(bond|br|dum|en|ersp|eth|gnv|ifb|ipoe|lan|l2tp|l2tpeth|macsec|peth|ppp|pppoe|pptp|sstp|sstpc|tun|veth|vti|vtun|vxlan|wg|wlan|wwan)[0-9]+(.\\d+)?|lo"},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr: "file-path --lookup-path /sys/class/net --directory",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "txt",
									Description: "Interface name",
								}},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									Script: []string{"${vyos_completion_dir}/list_interfaces"},
								}},
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
									NodeNameAttr: "srv6",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Accept SR-enabled IPv6 packets on this interface"},
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
											NodeNameAttr: "hmac",
											DefaultValue: []string{"accept"},
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Define HMAC policy for ingress SR-enabled packets on this interface"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"(accept|drop|ignore)"},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "accept",
													Description: "Accept packets without HMAC, validate packets with HMAC",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "drop",
													Description: "Drop packets without HMAC, validate packets with HMAC",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ignore",
													Description: "Ignore HMAC field.",
												}},
												CompletionHelp: []*interfacedefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													List: []string{"accept drop ignore"},
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
