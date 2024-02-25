// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func interfaces_wireguard() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*interfacedefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "interfaces",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				TagNode: []*interfacedefinition.TagNode{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "tagNode",
					},
					NodeNameAttr: "wireguard",
					OwnerAttr:    "${vyos_conf_scripts_dir}/interfaces_wireguard.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"WireGuard Interface"},
						Constraint: []*interfacedefinition.Constraint{{
							XMLName: xml.Name{
								Local: "constraint",
							},
							Regex: []string{"wg[0-9]+"},
						}},
						ValueHelp: []*interfacedefinition.ValueHelp{{
							XMLName: xml.Name{
								Local: "valueHelp",
							},
							Format:      "wgN",
							Description: "WireGuard interface name",
						}},
						ConstraintErrorMessage: []string{"WireGuard interface must be named wgN"},
						Priority:               []string{"379"},
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
							NodeNameAttr: "mirror",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Mirror ingress/egress packets"},
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
									NodeNameAttr: "ingress",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Mirror ingress traffic to destination interface"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Destination interface name",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											Script: []string{"${vyos_completion_dir}/list_interfaces"},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "egress",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Mirror egress traffic to destination interface"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Destination interface name",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
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
							NodeNameAttr: "ip",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"IPv4 routing parameters"},
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
									NodeNameAttr: "adjust-mss",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Adjust TCP MSS value"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(clamp-mss-to-pmtu)"},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 536-65535",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "clamp-mss-to-pmtu",
											Description: "Automatically sets the MSS to the proper value",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:536-65535",
											Description: "TCP Maximum segment size in bytes",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"clamp-mss-to-pmtu"},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "arp-cache-timeout",
									DefaultValue: []string{"30"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"ARP cache entry timeout in seconds"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 1-86400",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:1-86400",
											Description: "ARP cache entry timout in seconds",
										}},
										ConstraintErrorMessage: []string{"ARP cache entry timeout must be between 1 and 86400 seconds"},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "disable-arp-filter",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Disable ARP filter on this interface"},
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
									NodeNameAttr: "disable-forwarding",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Disable IP forwarding on this interface"},
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
									NodeNameAttr: "enable-directed-broadcast",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Enable directed broadcast forwarding on this interface"},
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
									NodeNameAttr: "enable-arp-accept",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Enable ARP accept on this interface"},
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
									NodeNameAttr: "enable-arp-announce",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Enable ARP announce on this interface"},
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
									NodeNameAttr: "enable-arp-ignore",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Enable ARP ignore on this interface"},
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
									NodeNameAttr: "enable-proxy-arp",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Enable proxy-arp on this interface"},
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
									NodeNameAttr: "proxy-arp-pvlan",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Enable private VLAN proxy ARP on this interface"},
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
									NodeNameAttr: "source-validation",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Source validation by reversed path (RFC3704)"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(strict|loose|disable)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "strict",
											Description: "Enable Strict Reverse Path Forwarding as defined in RFC3704",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "loose",
											Description: "Enable Loose Reverse Path Forwarding as defined in RFC3704",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "disable",
											Description: "No source validation",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"strict loose disable"},
										}},
									}},
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "node",
							},
							NodeNameAttr: "ipv6",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"IPv6 routing parameters"},
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
									NodeNameAttr: "address",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IPv6 address configuration modes"},
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
											NodeNameAttr: "autoconf",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)"},
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
											NodeNameAttr: "eui64",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Prefix for IPv6 address with MAC-based EUI-64"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*interfacedefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ipv6-eui64-prefix",
													}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "<h:h:h:h:h:h:h:h/64>",
													Description: "IPv6 /64 network",
												}},
												ConstraintErrorMessage: []string{"EUI64 prefix length must be 64"},
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
											NodeNameAttr: "no-default-link-local",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Remove the default link-local address from the interface"},
												Valueless: []*interfacedefinition.Valueless{{
													XMLName: xml.Name{
														Local: "valueless",
													},
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
									NodeNameAttr: "adjust-mss",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Adjust TCP MSS value"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(clamp-mss-to-pmtu)"},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 536-65535",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "clamp-mss-to-pmtu",
											Description: "Automatically sets the MSS to the proper value",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:536-65535",
											Description: "TCP Maximum segment size in bytes",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"clamp-mss-to-pmtu"},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "disable-forwarding",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Disable IP forwarding on this interface"},
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
									NodeNameAttr: "accept-dad",
									DefaultValue: []string{"1"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Accept Duplicate Address Detection"},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "0",
											Description: "Disable DAD",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "1",
											Description: "Enable DAD",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "2",
											Description: "Enable DAD - disable IPv6 if MAC-based duplicate link-local address found",
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "dup-addr-detect-transmits",
									DefaultValue: []string{"1"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Number of NS messages to send while performing DAD"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--non-negative",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:0",
											Description: "Disable Duplicate Address Dectection (DAD)",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:1-n",
											Description: "Number of NS messages to send while performing DAD",
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "source-validation",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Source validation by reversed path (RFC3704)"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(strict|loose|disable)"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "strict",
											Description: "Enable Strict Reverse Path Forwarding as defined in RFC3704",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "loose",
											Description: "Enable Loose Reverse Path Forwarding as defined in RFC3704",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "disable",
											Description: "No source validation",
										}},
										CompletionHelp: []*interfacedefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"strict loose disable"},
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
							NodeNameAttr: "peer",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"peer alias"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[^ ]{1,100}"},
								}},
								ConstraintErrorMessage: []string{"peer alias too long (limit 100 characters)"},
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
									NodeNameAttr: "disable",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Disable instance"},
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
									NodeNameAttr: "description",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Description"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"[[:ascii:]]{0,256}"},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Description",
										}},
										ConstraintErrorMessage: []string{"Description too long (limit 256 characters)"},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "public-key",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"base64 encoded public key"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "base64",
											}},
										}},
										ConstraintErrorMessage: []string{"Key is not base64-encoded"},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "preshared-key",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"base64 encoded preshared key"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "base64",
											}},
										}},
										ConstraintErrorMessage: []string{"Key is not base64-encoded"},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "allowed-ips",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IP addresses allowed to traverse the peer"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ip-prefix",
											}},
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
									NodeNameAttr: "address",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IP address of tunnel endpoint"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ip-address",
											}, {
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ipv6-link-local",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv4",
											Description: "IPv4 address of remote tunnel endpoint",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv6",
											Description: "IPv6 address of remote tunnel endpoint",
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "port",
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
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "persistent-keepalive",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Interval to send keepalive messages"},
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
											Description: "Interval in seconds",
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
							NodeNameAttr: "address",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"IP address"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr: "ip-host",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ipv4net",
									Description: "IPv4 address and prefix length",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ipv6net",
									Description: "IPv6 address and prefix length",
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
							NodeNameAttr: "description",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Description"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[[:ascii:]]{0,256}"},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "txt",
									Description: "Description",
								}},
								ConstraintErrorMessage: []string{"Description too long (limit 256 characters)"},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "disable",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Administratively disable interface"},
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
							NodeNameAttr: "port",
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
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "mtu",
							DefaultValue: []string{"1420"},
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Maximum Transmission Unit (MTU)"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 68-16000",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:68-16000",
									Description: "Maximum Transmission Unit in byte",
								}},
								ConstraintErrorMessage: []string{"MTU must be between 68 and 16000"},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "fwmark",
							DefaultValue: []string{"0"},
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"A 32-bit fwmark value set on all outgoing packets"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 0-4294967295",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "number",
									Description: "value which marks the packet for QoS/shaper",
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "private-key",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Base64 encoded private key"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr: "base64",
									}},
								}},
								ConstraintErrorMessage: []string{"Key is not base64-encoded"},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "redirect",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Redirect incoming packet to destination"},
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
									Description: "Destination interface name",
								}},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									Script: []string{"${vyos_completion_dir}/list_interfaces"},
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "per-client-thread",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Process traffic from each client in a dedicated thread"},
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
							NodeNameAttr: "vrf",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"VRF instance name"},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "txt",
									Description: "VRF instance name",
								}},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									Path: []string{"vrf name"},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}
