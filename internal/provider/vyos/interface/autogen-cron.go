// Code generated by /workspaces/terraform-provider-vyos/tools/build-resource-terraform-schemas/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/provider/vyos/schema/interfacedefinition"
)

func Cron() interfacedefinition.InterfaceDefinition {
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
								Properties: []*interfacedefinition.Properties{
									{
										XMLName: xml.Name{
											Local: `properties`},
										Help: []string{
											`Task scheduler settings`},
										KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
								Children: []*interfacedefinition.Children{
									{
										XMLName: xml.Name{
											Local: `children`},
										TagNode: []*interfacedefinition.TagNode{
											{
												XMLName: xml.Name{
													Local: `tagNode`},
												OwnerAttr: `${vyos_conf_scripts_dir}/task_scheduler.py`,
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Scheduled task`},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `txt`,
																Description: `Task name`}},
														Priority: []string{
															`999`},
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
																			`Executable path and arguments`},
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
																							`Path to executable`},
																						KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
																			{
																				XMLName: xml.Name{
																					Local: `leafNode`},
																				Properties: []*interfacedefinition.Properties{
																					{
																						XMLName: xml.Name{
																							Local: `properties`},
																						Help: []string{
																							`Arguments passed to the executable`},
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
																			`UNIX crontab time specification string`},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
															{
																XMLName: xml.Name{
																	Local: `leafNode`},
																Properties: []*interfacedefinition.Properties{
																	{
																		XMLName: xml.Name{
																			Local: `properties`},
																		Help: []string{
																			`Execution interval`},
																		Constraint: []*interfacedefinition.Constraint{
																			{
																				XMLName: xml.Name{
																					Local: `constraint`},
																				Regex: []string{
																					`[1-9]([0-9]&)([mhd]{0,
1})`},
																				Validator: []*interfacedefinition.Validator(nil)}},
																		ValueHelp: []*interfacedefinition.ValueHelp{
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `<minutes>`,
																				Description: `Execution interval in minutes`},
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `<minutes>m`,
																				Description: `Execution interval in minutes`},
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `<hours>h`,
																				Description: `Execution interval in hours`},
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `<days>d`,
																				Description: `Execution interval in days`}},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
										LeafNode: []*interfacedefinition.LeafNode(nil)}}}},
						LeafNode: []*interfacedefinition.LeafNode(nil)}}}}}
}
