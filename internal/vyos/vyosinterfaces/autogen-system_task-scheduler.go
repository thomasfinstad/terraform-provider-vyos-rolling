// Code generated by tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func system_taskscheduler() schemadefinition.InterfaceDefinition {
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
					NodeNameAttr: "task-scheduler",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Task scheduler settings"},
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
							NodeNameAttr: "task",
							OwnerAttr:    "${vyos_conf_scripts_dir}/system_task-scheduler.py",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Scheduled task"},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "txt",
									Description: "Task name",
								}},
								Priority: []string{"999"},
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
									NodeNameAttr: "executable",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Executable path and arguments"},
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
											NodeNameAttr: "path",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Path to executable"},
											}},
										}, {
											IsBaseNode: false,
											XMLName: xml.Name{
												Local: "leafNode",
											},
											NodeNameAttr: "arguments",
											Properties: []*schemadefinition.Properties{{
												XMLName: xml.Name{
													Local: "properties",
												},
												Help: []string{"Arguments passed to the executable"},
											}},
										}},
									}},
								}},
								LeafNode: []*schemadefinition.LeafNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "crontab-spec",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"UNIX crontab time specification string"},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "interval",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Execution interval"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"[1-9]([0-9]&)([mhd]{0,1})"},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "<minutes>",
											Description: "Execution interval in minutes",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "<minutes>m",
											Description: "Execution interval in minutes",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "<hours>h",
											Description: "Execution interval in hours",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "<days>d",
											Description: "Execution interval in days",
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
