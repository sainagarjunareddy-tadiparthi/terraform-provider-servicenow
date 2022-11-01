# Terraform ServiceNow Provider

A custom provider for Terraform to manage objects in a ServiceNow instance for developping system applications outside of an instance. This is especially useful when you want to create an Application with proper source control and continuous development integration.

<img src="https://camo.githubusercontent.com/1a4ed08978379480a9b1ca95d7f4cc8eb80b45ad47c056a7cfb5c597e9315ae5/68747470733a2f2f7777772e6461746f636d732d6173736574732e636f6d2f323838352f313632393934313234322d6c6f676f2d7465727261666f726d2d6d61696e2e737667" height="200" width="200" alt="Terraform Logo"/><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/Plus_symbol.svg/640px-Plus_symbol.svg.png" height="200" width="200" alt="plus logo"/><img src="https://www.servicenow.com/content/dam/servicenow-assets/images/meganav/servicenow-header-logo.svg" height="200" width="200" alt="ServiceNow Logo"/>

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
