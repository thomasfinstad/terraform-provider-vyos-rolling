module github.com/thomasfinstad/terraform-provider-vyos-rolling/tools/generate-terraform-resource-full

go 1.23.1

replace github.com/thomasfinstad/terraform-provider-vyos-rolling => ../../

require github.com/thomasfinstad/terraform-provider-vyos-rolling v1.6.202408210

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
