terraform {
  required_providers {
    servicenow = {
      source  = "terraform.local/local/servicenow"
      version = "1.0.0"
    }
  }
}

provider "servicenow" {
  instance_url = var.ServiceNowUrl
  username     = var.ServiceNowUser
  password     = var.ServiceNowPwd
}

variable "ServiceNowUrl" {}

variable "ServiceNowUser" {}

variable "ServiceNowPwd" {}

resource "servicenow_ui_page" "test-page" {
  name              = "My test page"
  description       = "This is a description!"
  html              = "<p>Hello World!</p>"
  client_script     = "alert('Hello World!')"
  processing_script = "gs.info('Hello World!')"
  category          = "general"
  direct            = false
}