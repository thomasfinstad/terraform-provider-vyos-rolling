// Code generated by /workspaces/terraform-provider-vyos/tools/build-resource-terraform-schemas/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/provider/vyos/schema/interfacedefinition"
)

func systemsflow() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{Local: "interfaceDefinition"},
		Node: []*interfacedefinition.Node{{
			XMLName: xml.Name{Local: "node"}, NodeNameAttr: "system",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{Local: "children"},
				Node: []*interfacedefinition.Node{{
					XMLName: xml.Name{Local: "node"}, NodeNameAttr: "sflow", OwnerAttr: "${vyos_conf_scripts_dir}/system_sflow.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName:  xml.Name{Local: "properties"},
						Help:     []string{"sFlow settings"},
						Priority: []string{"990"},
					}},
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{Local: "children"},
						TagNode: []*interfacedefinition.TagNode{{
							XMLName: xml.Name{Local: "tagNode"}, NodeNameAttr: "server",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{Local: "properties"},
								Help:    []string{"sFlow destination server"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName:   xml.Name{Local: "constraint"},
									Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv4-address"}, {XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv4", Description: "IPv4 server to export sFlow"}, {XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "IPv6 server to export sFlow"}},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{Local: "children"},
								LeafNode: []*interfacedefinition.LeafNode{{
									XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "port",
									DefaultValue: []string{"6343"},
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{Local: "properties"},
										Help:    []string{"Port number used by connection"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName:   xml.Name{Local: "constraint"},
											Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "numeric", ArgumentAttr: "--range 1-65535"}},
										}},
										ValueHelp:              []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "u32:1-65535", Description: "Numeric IP port"}},
										ConstraintErrorMessage: []string{"Port number must be in range 1 to 65535"},
									}},
								}},
							}},
						}},
						LeafNode: []*interfacedefinition.LeafNode{{
							XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "agent-address",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{Local: "properties"},
								Help:    []string{"sFlow agent IPv4 or IPv6 address"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName:   xml.Name{Local: "constraint"},
									Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv4-address"}, {XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-address"}, {XMLName: xml.Name{Local: "validator"}, NameAttr: "ipv6-link-local"}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv4", Description: "sFlow IPv4 agent address"}, {XMLName: xml.Name{Local: "valueHelp"}, Format: "ipv6", Description: "sFlow IPv6 agent address"}},
								CompletionHelp: []*interfacedefinition.CompletionHelp{{
									XMLName: xml.Name{Local: "completionHelp"},
									List:    []string{"auto"},
									Script:  []string{"${vyos_completion_dir}/list_local_ips.sh --both"},
								}},
							}},
						}, {
							XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "agent-interface",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{Local: "properties"},
								Help:    []string{"IP address associated with this interface"},
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
						}, {
							XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "drop-monitor-limit",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{Local: "properties"},
								Help:    []string{"Export headers of dropped by kernel packets"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName:   xml.Name{Local: "constraint"},
									Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "numeric", ArgumentAttr: "--range 1-65535"}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "u32:1-65535", Description: "Maximum rate limit of N drops per second send out in the sFlow datagrams"}},
							}},
						}, {
							XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "interface",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{Local: "properties"},
								Help:    []string{"Interface to use"},
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
								Multi: []*interfacedefinition.Multi{{XMLName: xml.Name{Local: "multi"}}},
							}},
						}, {
							XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "polling",
							DefaultValue: []string{"30"},
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{Local: "properties"},
								Help:    []string{"Schedule counter-polling in seconds"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName:   xml.Name{Local: "constraint"},
									Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "numeric", ArgumentAttr: "--range 1-600"}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "u32:1-600", Description: "Polling rate in seconds"}},
							}},
						}, {
							XMLName: xml.Name{Local: "leafNode"}, NodeNameAttr: "sampling-rate",
							DefaultValue: []string{"1000"},
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{Local: "properties"},
								Help:    []string{"sFlow sampling-rate"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName:   xml.Name{Local: "constraint"},
									Validator: []*interfacedefinition.Validator{{XMLName: xml.Name{Local: "validator"}, NameAttr: "numeric", ArgumentAttr: "--range 1-65535"}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{XMLName: xml.Name{Local: "valueHelp"}, Format: "u32:1-65535", Description: "Sampling rate (1 in N packets)"}},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}
