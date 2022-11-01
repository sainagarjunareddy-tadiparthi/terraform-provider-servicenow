# Terraform ServiceNow Provider

A custom provider for Terraform to manage objects in a ServiceNow instance for developping system applications outside of an instance. This is especially useful when you want to create an Application with proper source control and continuous development integration.

<img src="https://www.terraform.io/assets/images/og-image-8b3e4f7d.png" height="200" alt="Terraform Logo"/><img src="https://community.servicenow.com/c4fe846adbb95f0037015e77dc961918.iix" height="200" alt="ServiceNow Logo"/>

[![Release Badge](https://img.shields.io/github/release/sainagarjunareddy-tadiparthi/terraform-provider-servicenow.svg)](https://github.com/sainagarjunareddy-tadiparthi/terraform-provider-servicenow/releases/latest)
[![License Badge](https://img.shields.io/github/license/sainagarjunareddy-tadiparthi/terraform-provider-servicenow.svg)](LICENSE)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html)
- [Go](https://golang.org/doc/install) (to build the provider plugin)

## Installation

### Windows

1. Clone repository to: `%GOPATH%/src/github.com/terraform-providers/terraform-provider-servicenow`
1. Build the executable using `go build -o terraform-provider-servicenow.exe`
1. Copy the file to `%APPDATA%\terraform.d\plugins`

### Linux

1. Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-servicenow`
1. Build the executable using `go build -o terraform-provider-servicenow`
1. Copy the file to `~.terraform.d/plugins`

### Other

You can also download the [latest release](https://github.com/sainagarjunareddy-tadiparthi/terraform-provider-servicenow/releases) binaries and place them in your working directory, since Terraform will look for providers in the working directory also.

## Supported Resources

Check out the [Wiki](https://github.com/sainagarjunareddy-tadiparthi/terraform-provider-servicenow/wiki) !

Build Instructions:

1. go mod init github.com/sainagarjunareddy-tadiparthi/terraform-provider-servicenow
2. go mod tidy
3. go build -o terraform-provider-servicenow
4. Copy the binary file to ~/.terraform.d/plugins/terraform.local/local/servicenow/1.0.0/_________
5. terraform init
6. terraform plan
7. terraform apply

## Setup your local for testing

1. Init your project and get depdendencies.
2. Build the project to get the binary file.
3. Copy the generated binary file to ~/.terraform.d/plugins/terraform.local/local/<provider_name>/<provider_version>/darwin_amd64/<your_binary_file>
4. If above folders are not available. Create them manually.
5. Initialize the terrform, plan and apply to make it work.