// Code generated by /repo/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func interfaces_input() schemadefinition.InterfaceDefinition {
	return schemadefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*schemadefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "interfaces",
			Children: []*schemadefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				TagNode: []*schemadefinition.TagNode{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "tagNode",
					},
					NodeNameAttr: "input",
					OwnerAttr:    "${vyos_conf_scripts_dir}/interfaces_input.py",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Input Functional Block (IFB) interface name"},
						Constraint: []*schemadefinition.Constraint{{
							XMLName: xml.Name{
								Local: "constraint",
							},
							Regex: []string{"ifb[0-9]+"},
						}},
						ValueHelp: []*schemadefinition.ValueHelp{{
							XMLName: xml.Name{
								Local: "valueHelp",
							},
							Format:      "ifbN",
							Description: "Input interface name",
						}},
						ConstraintErrorMessage: []string{"Input interface must be named ifbN"},
						Priority:               []string{"310"},
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
							NodeNameAttr: "disable",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Administratively disable interface"},
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
							NodeNameAttr: "redirect",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Redirect incoming packet to destination"},
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
									Description: "Destination interface name",
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
			}},
		}},
	}
}
