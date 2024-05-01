// Code generated by /repo/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func service_mdns_repeater() schemadefinition.InterfaceDefinition {
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
					NodeNameAttr: "mdns",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Multicast DNS (mDNS) parameters"},
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
							NodeNameAttr: "repeater",
							OwnerAttr:    "${vyos_conf_scripts_dir}/service_mdns_repeater.py",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help:     []string{"mDNS repeater configuration"},
								Priority: []string{"990"},
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
									NodeNameAttr: "ip-version",
									DefaultValue: []string{"both"},
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
									NodeNameAttr: "browse-domain",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"mDNS browsing domains in addition to the default one"},
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
											Description: "mDNS browsing domain",
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
									NodeNameAttr: "allow-service",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Allowed mDNS services to be repeated"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"[-_.a-zA-Z0-9]+"},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "txt",
											Description: "mDNS service",
										}},
										ConstraintErrorMessage: []string{"Service name must be alphanumeric and can contain hyphens and underscores"},
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
									NodeNameAttr: "vrrp-disable",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Disables mDNS repeater on VRRP interfaces not in MASTER state"},
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
			}},
		}},
	}
}
