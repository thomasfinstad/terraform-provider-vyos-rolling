// Code generated by tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func system_ipv6() schemadefinition.InterfaceDefinition {
	return schemadefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*schemadefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "system",
			Children: []*schemadefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*schemadefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "ipv6",
					OwnerAttr:    "${vyos_conf_scripts_dir}/system_ipv6.py",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"IPv6 Settings"},
						Priority: []string{"290"},
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
							NodeNameAttr: "multipath",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"IPv6 multipath settings"},
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
									NodeNameAttr: "layer4-hashing",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Use layer 4 information for ECMP hashing"},
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
							NodeNameAttr: "neighbor",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Parameters for neighbor discovery cache"},
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
									NodeNameAttr: "table-size",
									DefaultValue: []string{"8192"},
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Maximum number of entries to keep in the cache"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(1024|2048|4096|8192|16384|32768)"},
										}},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"1024 2048 4096 8192 16384 32768"},
										}},
									}},
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "nht",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Filter Next Hop tracking route resolution"},
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
									NodeNameAttr: "no-resolve-via-default",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Do not resolve via default route"},
										Valueless: []*schemadefinition.Valueless{{
											XMLName: xml.Name{
												Local: "valueless",
											},
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
							NodeNameAttr: "protocol",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Filter routing info exchanged between routing protocol and zebra"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(any|babel|bgp|connected|isis|kernel|ospfv3|ripng|static|table)"},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "any",
									Description: "Any of the above protocols",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "babel",
									Description: "Babel routing protocol",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "bgp",
									Description: "Border Gateway Protocol",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "connected",
									Description: "Connected routes (directly attached subnet or host)",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "isis",
									Description: "Intermediate System to Intermediate System",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "kernel",
									Description: "Kernel routes (not installed via the zebra RIB)",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ospfv3",
									Description: "Open Shortest Path First (OSPFv3)",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ripng",
									Description: "Routing Information Protocol next-generation",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "static",
									Description: "Statically configured routes",
								}},
								CompletionHelp: []*schemadefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									List: []string{"any babel bgp connected isis kernel ospfv3 ripng static table"},
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
									NodeNameAttr: "route-map",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Specify route-map name to use"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"[-_a-zA-Z0-9][\\w\\-\\.\\+]&"},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Route map name",
										}},
										ConstraintErrorMessage: []string{"Name of route-map can only contain alpha-numeric letters, hyphen and underscores"},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Path: []string{"policy route-map"},
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
							NodeNameAttr: "disable-forwarding",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Disable IPv6 forwarding on all interfaces"},
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
							NodeNameAttr: "strict-dad",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Disable IPv6 operation on interface when DAD fails on LL addr"},
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
