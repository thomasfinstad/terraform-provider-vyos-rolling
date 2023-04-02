// Code generated by /workspaces/terraform-provider-vyos/tools/build-resource-terraform-schemas/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/provider/vyos/schema/interfacedefinition"
)

func dhcpv6server() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{Local: "interfaceDefinition"},
		Node: []*interfacedefinition.Node{{
			XMLName: xml.Name{Local: "node"}, NodeNameAttr: "service",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{Local: "children"},
				Node: []*interfacedefinition.Node{{
					XMLName: xml.Name{Local: "node"}, NodeNameAttr: "dhcpv6-server", OwnerAttr: "${vyos_conf_scripts_dir}/dhcpv6_server.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName:  xml.Name{Local: "properties"},
						Help:     []string{"DHCP for IPv6 (DHCPv6) server"},
						Priority: []string{"900"},
					}},
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{Local: "children"},
						Node: []*interfacedefinition.Node{{
							XMLName: xml.Name{Local: "node"}, NodeNameAttr: "global-parameters",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{Local: "properties"},
								Help:    []string{"Additional global parameters for DHCPv6 server"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{Local: "children"},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "name-server",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{Local: "properties"},
										Help:    []string{"Domain Name Servers (DNS) addresses"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName:   xml.Name{Local: "constraint"},
											Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "Domain Name Server (DNS) IPv6 address"}},
										Multi:     []*interfacedefinition.Multi{{XMLName: xml.Name{Local: "multi"}}},
									}},
								}},
							}},
						}},
						TagNode: []*interfacedefinition.TagNode{{
							XMLName: xml.Name{Local: "tagNode"}, NodeNameAttr: "shared-network-name",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{Local: "properties"},
								Help:    []string{"DHCPv6 shared network name"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{Local: "constraint"},
									Regex:   []string{"[-_a-zA-Z0-9.]+"},
								}},
								ConstraintErrorMessage: []string{"Invalid DHCPv6 shared network name. May only contain letters, numbers and .-_"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{Local: "children"},
								Node: []*interfacedefinition.Node{{
									XMLName: xml.Name{Local: "node"}, NodeNameAttr: "common-options",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{Local: "properties"},
										Help:    []string{"Common options to distribute to all clients, including stateless clients"},
									}},
									Children: []*interfacedefinition.Children{{
										XMLName: xml.Name{Local: "children"},
										LeafNode: []*interfacedefinition.LeafNode{{
											XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "info-refresh-time",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"Time (in seconds) that stateless clients should wait between refreshing the information they were given"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName:   xml.Name{Local: "constraint"},
													Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "numeric", ArgumentAttr: "--range 1-4294967295"}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "u32:1-4294967295", Description: "DHCPv6 information refresh time"}},
											}},
										}, {
											XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "domain-search",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"Client Domain Name search list"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName:   xml.Name{Local: "constraint"},
													Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "fqdn"}},
												}},
												ConstraintErrorMessage: []string{"Invalid domain name (RFC 1123 section 2).\\nMay only contain letters, numbers, period, and underscore."},
												Multi:                  []*interfacedefinition.Multi{{XMLName: xml.Name{Local: "multi"}}},
											}},
										}, {
											XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "name-server",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"Domain Name Servers (DNS) addresses"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName:   xml.Name{Local: "constraint"},
													Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "Domain Name Server (DNS) IPv6 address"}},
												Multi:     []*interfacedefinition.Multi{{XMLName: xml.Name{Local: "multi"}}},
											}},
										}},
									}},
								}},
								TagNode: []*interfacedefinition.TagNode{{
									XMLName: xml.Name{Local: "tagNode"}, NodeNameAttr: "subnet",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{Local: "properties"},
										Help:    []string{"IPv6 DHCP subnet for this shared network"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName:   xml.Name{Local: "constraint"},
											Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-prefix"}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6net", Description: "IPv6 address and prefix length"}},
									}},
									Children: []*interfacedefinition.Children{{
										XMLName: xml.Name{Local: "children"},
										Node: []*interfacedefinition.Node{{
											XMLName: xml.Name{Local: "node"}, NodeNameAttr: "address-range",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"Parameters setting ranges for assigning IPv6 addresses"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{Local: "children"},
												TagNode: []*interfacedefinition.TagNode{{
													XMLName: xml.Name{Local: "tagNode"}, NodeNameAttr: "prefix",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"IPv6 prefix defining range of addresses to assign"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName:   xml.Name{Local: "constraint"},
															Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-prefix"}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6net", Description: "IPv6 address and prefix length"}},
													}},
													Children: []*interfacedefinition.Children{{
														XMLName: xml.Name{Local: "children"},
														LeafNode: []*interfacedefinition.LeafNode{{
															XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "temporary",
															Properties: []*interfacedefinition.Properties{{
																XMLName:   xml.Name{Local: "properties"},
																Help:      []string{"Address range will be used for temporary addresses"},
																Valueless: []*interfacedefinition.Valueless{{XMLName: xml.Name{Local: "valueless"}}},
															}},
														}},
													}},
												}, {
													XMLName: xml.Name{Local: "tagNode"}, NodeNameAttr: "start",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"First in range of consecutive IPv6 addresses to assign"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName:   xml.Name{Local: "constraint"},
															Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "IPv6 address"}},
													}},
													Children: []*interfacedefinition.Children{{
														XMLName: xml.Name{Local: "children"},
														LeafNode: []*interfacedefinition.LeafNode{{
															XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "stop",
															Properties: []*interfacedefinition.Properties{{
																XMLName: xml.Name{Local: "properties"},
																Help:    []string{"Last in range of consecutive IPv6 addresses"},
																Constraint: []*interfacedefinition.Constraint{{
																	XMLName:   xml.Name{Local: "constraint"},
																	Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
																}},
																ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "IPv6 address"}},
															}},
														}},
													}},
												}},
											}},
										}, {
											XMLName: xml.Name{Local: "node"}, NodeNameAttr: "lease-time",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"Parameters relating to the lease time"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{Local: "children"},
												LeafNode: []*interfacedefinition.LeafNode{{
													XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "default",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"Default time (in seconds) that will be assigned to a lease"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName:   xml.Name{Local: "constraint"},
															Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "numeric", ArgumentAttr: "--range 1-4294967295"}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "u32:1-4294967295", Description: "DHCPv6 valid lifetime"}},
													}},
												}, {
													XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "maximum",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"Maximum time (in seconds) that will be assigned to a lease"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName:   xml.Name{Local: "constraint"},
															Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "numeric", ArgumentAttr: "--range 1-4294967295"}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "u32:1-4294967295", Description: "Maximum lease time in seconds"}},
													}},
												}, {
													XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "minimum",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"Minimum time (in seconds) that will be assigned to a lease"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName:   xml.Name{Local: "constraint"},
															Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "numeric", ArgumentAttr: "--range 1-4294967295"}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "u32:1-4294967295", Description: "Minimum lease time in seconds"}},
													}},
												}},
											}},
										}, {
											XMLName: xml.Name{Local: "node"}, NodeNameAttr: "prefix-delegation",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"Parameters relating to IPv6 prefix delegation"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{Local: "children"},
												TagNode: []*interfacedefinition.TagNode{{
													XMLName: xml.Name{Local: "tagNode"}, NodeNameAttr: "start",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"First in range of IPv6 addresses to be used in prefix delegation"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName:   xml.Name{Local: "constraint"},
															Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "IPv6 address used in prefix delegation"}},
													}},
													Children: []*interfacedefinition.Children{{
														XMLName: xml.Name{Local: "children"},
														LeafNode: []*interfacedefinition.LeafNode{{
															XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "prefix-length",
															Properties: []*interfacedefinition.Properties{{
																XMLName: xml.Name{Local: "properties"},
																Help:    []string{"Length in bits of prefixes to be delegated"},
																Constraint: []*interfacedefinition.Constraint{{
																	XMLName:   xml.Name{Local: "constraint"},
																	Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "numeric", ArgumentAttr: "--range 32-64"}},
																}},
																ValueHelp:              []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "u32:32-64", Description: "Delagated prefix length (32-64)"}},
																ConstraintErrorMessage: []string{"Delegated prefix length must be between 32 and 64"},
															}},
														}, {
															XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "stop",
															Properties: []*interfacedefinition.Properties{{
																XMLName: xml.Name{Local: "properties"},
																Help:    []string{"Last in range of IPv6 addresses to be used in prefix delegation"},
																Constraint: []*interfacedefinition.Constraint{{
																	XMLName:   xml.Name{Local: "constraint"},
																	Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
																}},
																ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "IPv6 address used in prefix delegation"}},
															}},
														}},
													}},
												}},
											}},
										}, {
											XMLName: xml.Name{Local: "node"}, NodeNameAttr: "vendor-option",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"Vendor Specific Options"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{Local: "children"},
												Node: []*interfacedefinition.Node{{
													XMLName: xml.Name{Local: "node"}, NodeNameAttr: "cisco",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"Cisco specific parameters"},
													}},
													Children: []*interfacedefinition.Children{{
														XMLName: xml.Name{Local: "children"},
														LeafNode: []*interfacedefinition.LeafNode{{
															XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "tftp-server",
															Properties: []*interfacedefinition.Properties{{
																XMLName: xml.Name{Local: "properties"},
																Help:    []string{"TFTP server name"},
																Constraint: []*interfacedefinition.Constraint{{
																	XMLName:   xml.Name{Local: "constraint"},
																	Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
																}},
																ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "TFTP server IPv6 address"}},
																Multi:     []*interfacedefinition.Multi{{XMLName: xml.Name{Local: "multi"}}},
															}},
														}},
													}},
												}},
											}},
										}},
										TagNode: []*interfacedefinition.TagNode{{
											XMLName: xml.Name{Local: "tagNode"}, NodeNameAttr: "static-mapping",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"Name of static mapping"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{Local: "constraint"},
													Regex:   []string{"[-_a-zA-Z0-9.]+"},
												}},
												ConstraintErrorMessage: []string{"Invalid static mapping name. May only contain letters, numbers and .-_"},
											}},
											Children: []*interfacedefinition.Children{{
												XMLName: xml.Name{Local: "children"},
												LeafNode: []*interfacedefinition.LeafNode{{
													XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "disable",
													Properties: []*interfacedefinition.Properties{{
														XMLName:   xml.Name{Local: "properties"},
														Help:      []string{"Disable instance"},
														Valueless: []*interfacedefinition.Valueless{{XMLName: xml.Name{Local: "valueless"}}},
													}},
												}, {
													XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "identifier",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"Client identifier (DUID) for this static mapping"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName: xml.Name{Local: "constraint"},
															Regex:   []string{"([0-9A-Fa-f]{1,2}[:])&([0-9A-Fa-f]{1,2})"},
														}},
														ValueHelp:              []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "h[[:h]...]", Description: "DUID: colon-separated hex list (as used by isc-dhcp option dhcpv6.client-id)"}},
														ConstraintErrorMessage: []string{"Invalid DUID, must be in the format h[[:h]...]"},
													}},
												}, {
													XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "ipv6-address",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"Client IPv6 address for this static mapping"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName:   xml.Name{Local: "constraint"},
															Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "IPv6 address for this static mapping"}},
													}},
												}, {
													XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "ipv6-prefix",
													Properties: []*interfacedefinition.Properties{{
														XMLName: xml.Name{Local: "properties"},
														Help:    []string{"Client IPv6 prefix for this static mapping"},
														Constraint: []*interfacedefinition.Constraint{{
															XMLName:   xml.Name{Local: "constraint"},
															Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-prefix"}},
														}},
														ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6net", Description: "IPv6 prefix for this static mapping"}},
													}},
												}},
											}},
										}},
										LeafNode: []*interfacedefinition.LeafNode{{
											XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "domain-search",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"Client Domain Name search list"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName:   xml.Name{Local: "constraint"},
													Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "fqdn"}},
												}},
												ConstraintErrorMessage: []string{"Invalid domain name (RFC 1123 section 2).\\nMay only contain letters, numbers, period, and underscore."},
												Multi:                  []*interfacedefinition.Multi{{XMLName: xml.Name{Local: "multi"}}},
											}},
										}, {
											XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "name-server",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"Domain Name Servers (DNS) addresses"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName:   xml.Name{Local: "constraint"},
													Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "Domain Name Server (DNS) IPv6 address"}},
												Multi:     []*interfacedefinition.Multi{{XMLName: xml.Name{Local: "multi"}}},
											}},
										}, {
											XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "nis-domain",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"NIS domain name for client to use"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{Local: "constraint"},
													Regex:   []string{"[-_a-zA-Z0-9.]+"},
												}},
												ConstraintErrorMessage: []string{"Invalid NIS domain name"},
											}},
										}, {
											XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "nis-server",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"IPv6 address of a NIS Server"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName:   xml.Name{Local: "constraint"},
													Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "IPv6 address of NIS server"}},
												Multi:     []*interfacedefinition.Multi{{XMLName: xml.Name{Local: "multi"}}},
											}},
										}, {
											XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "nisplus-domain",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"NIS+ domain name for client to use"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName: xml.Name{Local: "constraint"},
													Regex:   []string{"[-_a-zA-Z0-9.]+"},
												}},
												ConstraintErrorMessage: []string{"Invalid NIS+ domain name. May only contain letters, numbers and .-_"},
											}},
										}, {
											XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "nisplus-server",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"IPv6 address of a NIS+ Server"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName:   xml.Name{Local: "constraint"},
													Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "IPv6 address of NIS+ server"}},
												Multi:     []*interfacedefinition.Multi{{XMLName: xml.Name{Local: "multi"}}},
											}},
										}, {
											XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "sip-server",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"IPv6 address of SIP server"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName:   xml.Name{Local: "constraint"},
													Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}, {XMLName: xml.Name{Local: "validator"}, NameAttr: "fqdn"}},
												}},
												ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "IPv6 address of SIP server"}, {XMLName: xml.Name{Local: "valueHelp"}, Format: "hostname", Description: "FQDN of SIP server"}},
												Multi:     []*interfacedefinition.Multi{{XMLName: xml.Name{Local: "multi"}}},
											}},
										}, {
											XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "sntp-server",
											Properties: []*interfacedefinition.Properties{{
												XMLName: xml.Name{Local: "properties"},
												Help:    []string{"IPv6 address of an SNTP server for client to use"},
												Constraint: []*interfacedefinition.Constraint{{
													XMLName:   xml.Name{Local: "constraint"},
													Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
												}},
												Multi: []*interfacedefinition.Multi{{XMLName: xml.Name{Local: "multi"}}},
											}},
										}},
									}},
								}},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "disable",
									Properties: []*interfacedefinition.Properties{{
										XMLName:   xml.Name{Local: "properties"},
										Help:      []string{"Disable instance"},
										Valueless: []*interfacedefinition.Valueless{{XMLName: xml.Name{Local: "valueless"}}},
									}},
								}, {
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
								}},
							}},
						}},
						LeafNode: []*interfacedefinition.LeafNode{{
							XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "disable",
							Properties: []*interfacedefinition.Properties{{
								XMLName:   xml.Name{Local: "properties"},
								Help:      []string{"Disable instance"},
								Valueless: []*interfacedefinition.Valueless{{XMLName: xml.Name{Local: "valueless"}}},
							}},
						}, {
							XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "preference",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{Local: "properties"},
								Help:    []string{"Preference of this DHCPv6 server compared with others"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName:   xml.Name{Local: "constraint"},
									Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "numeric", ArgumentAttr: "--range 0-255"}},
								}},
								ValueHelp:              []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "u32:0-255", Description: "DHCPv6 server preference (0-255)"}},
								ConstraintErrorMessage: []string{"Preference must be between 0 and 255"},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}
