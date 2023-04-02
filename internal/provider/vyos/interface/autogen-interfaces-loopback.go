// Code generated by /workspaces/terraform-provider-vyos/tools/build-resource-terraform-schemas/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/provider/vyos/schema/interfacedefinition"
)

func InterfacesLoopback() interfacedefinition.InterfaceDefinition {
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
						TagNode: []*interfacedefinition.TagNode{
							{
								XMLName: xml.Name{
									Local: `tagNode`},
								OwnerAttr: `${vyos_conf_scripts_dir}/interfaces-loopback.py`,
								Properties: []*interfacedefinition.Properties{
									{
										XMLName: xml.Name{
											Local: `properties`},
										Help: []string{
											`Loopback Interface`},
										Constraint: []*interfacedefinition.Constraint{
											{
												XMLName: xml.Name{
													Local: `constraint`},
												Regex: []string{
													`lo`},
												Validator: []*interfacedefinition.Validator(nil)}},
										ValueHelp: []*interfacedefinition.ValueHelp{
											{
												XMLName: xml.Name{
													Local: `valueHelp`},
												Format:      `lo`,
												Description: `Loopback interface`}},
										ConstraintErrorMessage: []string{
											`Loopback interface must be named lo`},
										Priority: []string{
											`300`},
										KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
								Children: []*interfacedefinition.Children{
									{
										XMLName: xml.Name{
											Local: `children`},
										Node: []*interfacedefinition.Node{
											{
												XMLName: xml.Name{
													Local: `node`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`IPv4 routing parameters`},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
												Children: []*interfacedefinition.Children{
													{
														XMLName: xml.Name{
															Local: `children`},
														LeafNode: []*interfacedefinition.LeafNode{
															{
																XMLName: xml.Name{
																	Local: `leafNode`},
																Properties: []*interfacedefinition.Properties{
																	{
																		XMLName: xml.Name{
																			Local: `properties`},
																		Help: []string{
																			`Source validation by reversed path (RFC3704)`},
																		Constraint: []*interfacedefinition.Constraint{
																			{
																				XMLName: xml.Name{
																					Local: `constraint`},
																				Regex: []string{
																					`(strict|loose|disable)`},
																				Validator: []*interfacedefinition.Validator(nil)}},
																		ValueHelp: []*interfacedefinition.ValueHelp{
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `strict`,
																				Description: `Enable Strict Reverse Path Forwarding as defined in RFC3704`},
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `loose`,
																				Description: `Enable Loose Reverse Path Forwarding as defined in RFC3704`},
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `disable`,
																				Description: `No source validation`}},
																		CompletionHelp: []*interfacedefinition.CompletionHelp{
																			{
																				XMLName: xml.Name{
																					Local: `completionHelp`},
																				List: []string{
																					`strict loose disable`},
																				Script: []string(nil)}},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}},
											{
												XMLName: xml.Name{
													Local: `node`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Mirror ingress/egress packets`},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
												Children: []*interfacedefinition.Children{
													{
														XMLName: xml.Name{
															Local: `children`},
														LeafNode: []*interfacedefinition.LeafNode{
															{
																XMLName: xml.Name{
																	Local: `leafNode`},
																Properties: []*interfacedefinition.Properties{
																	{
																		XMLName: xml.Name{
																			Local: `properties`},
																		Help: []string{
																			`Mirror ingress traffic to destination interface`},
																		ValueHelp: []*interfacedefinition.ValueHelp{
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `txt`,
																				Description: `Destination interface name`}},
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
																			`Mirror egress traffic to destination interface`},
																		ValueHelp: []*interfacedefinition.ValueHelp{
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `txt`,
																				Description: `Destination interface name`}},
																		CompletionHelp: []*interfacedefinition.CompletionHelp{
																			{
																				XMLName: xml.Name{
																					Local: `completionHelp`},
																				Script: []string{
																					`${vyos_completion_dir}/list_interfaces`}}},
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
															`IP address`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Validator: []*interfacedefinition.Validator{
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `ip-host`,
																		ArgumentAttr: ``}}}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `ipv4net`,
																Description: `IPv4 address and prefix length`},
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `ipv6net`,
																Description: `IPv6 address and prefix length`}},
														Multi: []*interfacedefinition.Multi{
															{
																XMLName: xml.Name{
																	Local: `multi`}}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Description`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Regex: []string{
																	`[[:ascii:]]{0,
256}`},
																Validator: []*interfacedefinition.Validator(nil)}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `txt`,
																Description: `Description`}},
														ConstraintErrorMessage: []string{
															`Description too long (limit 256 characters)`},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Redirect incoming packet to destination`},
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
																Description: `Destination interface name`}},
														CompletionHelp: []*interfacedefinition.CompletionHelp{
															{
																XMLName: xml.Name{
																	Local: `completionHelp`},
																Script: []string{
																	`${vyos_completion_dir}/list_interfaces`}}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
						LeafNode: []*interfacedefinition.LeafNode(nil)}}}}}
}
