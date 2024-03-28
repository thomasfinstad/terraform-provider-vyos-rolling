// Code generated by /workspaces/terraform-provider-vyos-rolling/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func service_dns_dynamic() schemadefinition.InterfaceDefinition {
	return schemadefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*schemadefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "service",
			Children: []*schemadefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				Node: []*schemadefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "dns",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Domain Name System (DNS) related services"},
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
							NodeNameAttr: "dynamic",
							OwnerAttr:    "${vyos_conf_scripts_dir}/service_dns_dynamic.py",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help:     []string{"Dynamic DNS"},
								Priority: []string{"990"},
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
									NodeNameAttr: "name",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Dynamic DNS configuration"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"[-_a-zA-Z0-9]+"},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "Dynamic DNS service name",
										}},
										ConstraintErrorMessage: []string{"Dynamic DNS service name must be alphanumeric and can contain hyphens and underscores"},
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
											NodeNameAttr: "address",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Obtain IP address to send Dynamic DNS update for"},
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
													NodeNameAttr: "web",
													Properties: []*schemadefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"HTTP(S) web request to use"},
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
															NodeNameAttr: "url",
															Properties: []*schemadefinition.Properties{{
																XMLName: xml.Name{
																	Local: "properties",
																},
																Help: []string{"Remote URL"},
																Constraint: []*schemadefinition.Constraint{{
																	XMLName: xml.Name{
																		Local: "constraint",
																	},
																	Validator: []*schemadefinition.Validator{{
																		XMLName: xml.Name{
																			Local: "validator",
																		},
																		NameAttr:     "url",
																		ArgumentAttr: "--scheme http --scheme https",
																	}},
																}},
																ValueHelp: []*schemadefinition.ValueHelp{{
																	XMLName: xml.Name{
																		Local: "valueHelp",
																	},
																	Format:      "url",
																	Description: "Remote HTTP(S) URL",
																}},
																ConstraintErrorMessage: []string{"Invalid HTTP(S) URL format"},
															}},
														}, {
															IsBaseNode: false,
															XMLName: xml.Name{
																Local: "leafNode",
															},
															NodeNameAttr: "skip",
															Properties: []*schemadefinition.Properties{{
																XMLName: xml.Name{
																	Local: "properties",
																},
																Help: []string{"Pattern to skip from the HTTP(S) respose"},
																ValueHelp: []*schemadefinition.ValueHelp{{
																	XMLName: xml.Name{
																		Local: "valueHelp",
																	},
																	Format:      "txt",
																	Description: "Pattern to skip from the HTTP(S) respose to extract the external IP address",
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
													NodeNameAttr: "interface",
													Properties: []*schemadefinition.Properties{{
														XMLName: xml.Name{
															Local: "properties",
														},
														Help: []string{"Interface to use"},
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
														ValueHelp: []*schemadefinition.ValueHelp{{
															XMLName: xml.Name{
																Local: "valueHelp",
															},
															Format:      "txt",
															Description: "Interface name",
														}},
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
											NodeNameAttr: "protocol",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"ddclient protocol used for Dynamic DNS service"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ddclient-protocol",
													}},
												}},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													Script: []string{"${vyos_completion_dir}/list_ddclient_protocols.sh"},
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "ip-version",
											DefaultValue: []string{"ipv4"},
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"IP address version to use"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"(ipv[46]|both)"},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "_ipv4",
													Description: "Use only IPv4 address",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "_ipv6",
													Description: "Use only IPv6 address",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "both",
													Description: "Use both IPv4 and IPv6 address",
												}},
												ConstraintErrorMessage: []string{"IP Version must be literal 'ipv4', 'ipv6' or 'both'"},
												CompletionHelp: []*schemadefinition.CompletionHelp{{
													XMLName: xml.Name{
														Local: "completionHelp",
													},
													List: []string{"ipv4 ipv6 both"},
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "host-name",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Hostname to register with Dynamic DNS service"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"[A-Za-z0-9][-.A-Za-z0-9]&[A-Za-z0-9]", "(\\@|\\&)[-.A-Za-z0-9]&"},
												}},
												ConstraintErrorMessage: []string{"Host-name must be alphanumeric, can contain hyphens and can be prefixed with '@' or '&'"},
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
											NodeNameAttr: "server",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Remote Dynamic DNS server to send updates to"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "ip-address",
													}, {
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "fqdn",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv4",
													Description: "IPv4 address of the remote server",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "ipv6",
													Description: "IPv6 address of the remote server",
												}, {
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "hostname",
													Description: "Fully qualified domain name of the remote server",
												}},
												ConstraintErrorMessage: []string{"Remote server must be IP address or fully qualified domain name"},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "zone",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"DNS zone to be updated"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr: "fqdn",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Name of DNS zone",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "username",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Username used for authentication"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"[[:ascii:]]{1,128}"},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Username",
												}},
												ConstraintErrorMessage: []string{"Username is limited to ASCII characters only, with a total length of 128"},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "password",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Password used for authentication"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Regex: []string{"[[:ascii:]]{1,128}"},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "txt",
													Description: "Password",
												}},
												ConstraintErrorMessage: []string{"Password is limited to ASCII characters only, with a total length of 128"},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "key",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"File containing TSIG authentication key for RFC2136 nsupdate on remote DNS server"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "file-path",
														ArgumentAttr: "--strict --parent-dir /config/auth",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "filename",
													Description: "File in /config/auth directory",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "ttl",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Time-to-live (TTL)"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 0-2147483647",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:0-2147483647",
													Description: "TTL in seconds",
												}},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "wait-time",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Time in seconds to wait between update attempts"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 60-86400",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:60-86400",
													Description: "Time in seconds",
												}},
												ConstraintErrorMessage: []string{"Wait time must be between 60 and 86400 seconds"},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "expiry-time",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Time in seconds for the hostname to be marked expired in cache"},
												Constraint: []*schemadefinition.Constraint{{
													XMLName: xml.Name{
														Local: "constraint",
													},
													Validator: []*schemadefinition.Validator{{
														XMLName: xml.Name{
															Local: "validator",
														},
														NameAttr:     "numeric",
														ArgumentAttr: "--range 300-2160000",
													}},
												}},
												ValueHelp: []*schemadefinition.ValueHelp{{
													XMLName: xml.Name{
														Local: "valueHelp",
													},
													Format:      "u32:300-2160000",
													Description: "Time in seconds",
												}},
												ConstraintErrorMessage: []string{"Expiry time must be between 300 and 2160000 seconds"},
											}},
										}},
									}},
								}},
								LeafNode: []*schemadefinition.LeafNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "interval",
									DefaultValue: []string{"300"},
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Interval in seconds to wait between Dynamic DNS updates"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr:     "numeric",
												ArgumentAttr: "--range 60-3600",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "u32:60-3600",
											Description: "Time in seconds",
										}},
										ConstraintErrorMessage: []string{"Interval must be between 60 and 3600 seconds"},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "vrf",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"VRF instance name"},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "VRF instance name",
										}},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
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
			}},
		}},
	}
}
