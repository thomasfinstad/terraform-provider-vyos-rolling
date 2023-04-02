// Code generated by /workspaces/terraform-provider-vyos/tools/build-resource-terraform-schemas/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/provider/vyos/schema/interfacedefinition"
)

func SystemSflow() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: `interfaceDefinition`},
		Node: []*interfacedefinition.Node{
			{
				XMLName: xml.Name{
					Local: `node`},
				Children: []*interfacedefinition.Children{
					{
						XMLName: xml.Name{
							Local: `children`},
						Node: []*interfacedefinition.Node{
							{
								XMLName: xml.Name{
									Local: `node`},
								OwnerAttr: `${vyos_conf_scripts_dir}/system_sflow.py`,
								Properties: []*interfacedefinition.Properties{
									{
										XMLName: xml.Name{
											Local: `properties`},
										Help: []string{
											`sFlow settings`},
										Priority: []string{
											`990`},
										KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
								Children: []*interfacedefinition.Children{
									{
										XMLName: xml.Name{
											Local: `children`},
										TagNode: []*interfacedefinition.TagNode{
											{
												XMLName: xml.Name{
													Local: `tagNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`sFlow destination server`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Validator: []*interfacedefinition.Validator{
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `ipv4-address`,
																		ArgumentAttr: ``},
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `ipv6-address`,
																		ArgumentAttr: ``}}}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `ipv4`,
																Description: `IPv4 server to export sFlow`},
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `ipv6`,
																Description: `IPv6 server to export sFlow`}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
												Children: []*interfacedefinition.Children{
													{
														XMLName: xml.Name{
															Local: `children`},
														LeafNode: []*interfacedefinition.LeafNode{
															{
																XMLName: xml.Name{
																	Local: `leafNode`},
																DefaultValue: []string{
																	`6343`},
																Properties: []*interfacedefinition.Properties{
																	{
																		XMLName: xml.Name{
																			Local: `properties`},
																		Help: []string{
																			`Port number used by connection`},
																		Constraint: []*interfacedefinition.Constraint{
																			{
																				XMLName: xml.Name{
																					Local: `constraint`},
																				Validator: []*interfacedefinition.Validator{
																					{
																						XMLName: xml.Name{
																							Local: `validator`},
																						NameAttr:     `numeric`,
																						ArgumentAttr: `--range 1-65535`}}}},
																		ValueHelp: []*interfacedefinition.ValueHelp{
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `u32:1-65535`,
																				Description: `Numeric IP port`}},
																		ConstraintErrorMessage: []string{
																			`Port number must be in range 1 to 65535`},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
										LeafNode: []*interfacedefinition.LeafNode{
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`sFlow agent IPv4 or IPv6 address`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Validator: []*interfacedefinition.Validator{
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `ipv4-address`,
																		ArgumentAttr: ``},
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `ipv6-address`,
																		ArgumentAttr: ``},
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `ipv6-link-local`,
																		ArgumentAttr: ``}}}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `ipv4`,
																Description: `sFlow IPv4 agent address`},
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `ipv6`,
																Description: `sFlow IPv6 agent address`}},
														CompletionHelp: []*interfacedefinition.CompletionHelp{
															{
																XMLName: xml.Name{
																	Local: `completionHelp`},
																List: []string{
																	`auto`},
																Script: []string{
																	`${vyos_completion_dir}/list_local_ips.sh --both`}}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`IP address associated with this interface`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Regex: []string{
																	`(bond|br|dum|en|ersp|eth|gnv|ifb|lan|l2tp|l2tpeth|macsec|peth|ppp|pppoe|pptp|sstp|tun|veth|vti|vtun|vxlan|wg|wlan|wwan)[0-9]+(.\\d+)?|lo`},
																Validator: []*interfacedefinition.Validator{
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `file-path --lookup-path /sys/class/net --directory`,
																		ArgumentAttr: ``}}}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `txt`,
																Description: `Interface name`}},
														CompletionHelp: []*interfacedefinition.CompletionHelp{
															{
																XMLName: xml.Name{
																	Local: `completionHelp`},
																Script: []string{
																	`${vyos_completion_dir}/list_interfaces`}}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Interface to use`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Regex: []string{
																	`(bond|br|dum|en|ersp|eth|gnv|ifb|lan|l2tp|l2tpeth|macsec|peth|ppp|pppoe|pptp|sstp|tun|veth|vti|vtun|vxlan|wg|wlan|wwan)[0-9]+(.\\d+)?|lo`},
																Validator: []*interfacedefinition.Validator{
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `file-path --lookup-path /sys/class/net --directory`,
																		ArgumentAttr: ``}}}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `txt`,
																Description: `Interface name`}},
														CompletionHelp: []*interfacedefinition.CompletionHelp{
															{
																XMLName: xml.Name{
																	Local: `completionHelp`},
																Script: []string{
																	`${vyos_completion_dir}/list_interfaces`}}},
														Multi: []*interfacedefinition.Multi{
															{
																XMLName: xml.Name{
																	Local: `multi`}}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												DefaultValue: []string{
													`30`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Schedule counter-polling in seconds`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Validator: []*interfacedefinition.Validator{
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `numeric`,
																		ArgumentAttr: `--range 1-600`}}}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `u32:1-600`,
																Description: `Polling rate in seconds`}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												DefaultValue: []string{
													`1000`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`sFlow sampling-rate`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Validator: []*interfacedefinition.Validator{
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `numeric`,
																		ArgumentAttr: `--range 1-65535`}}}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `u32:1-65535`,
																Description: `Sampling rate (1 in N packets)`}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
						LeafNode: []*interfacedefinition.LeafNode(nil)}}}}}
}
