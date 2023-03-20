package interfacedefinition

/*
Contains structs to use with vyos schemas
*/

//Generate schema
//go:generate go run github.com/xuri/xgen/cmd/xgen -p interfacedefinition -i ../../../../../.build/vyos/schema/interface-definition.xsd -o autogen-structs.go -l Go
