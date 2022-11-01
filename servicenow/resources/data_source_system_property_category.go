package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sainagarjunareddy-tadiparthi/terraform-provider-servicenow/servicenow/client"
)

// DataSourceSystemPropertyCategory reads the informations about a single SystemPropertyCategory in ServiceNow.
func DataSourceSystemPropertyCategory() *schema.Resource {
	// Copy the schema from the resource.
	resourceSchema := ResourceSystemPropertyCategory().Schema
	setOnlyRequiredSchema(resourceSchema, systemPropertyCategoryName)

	return &schema.Resource{
		Read:   readDataSourceSystemPropertyCategory,
		Schema: resourceSchema,
	}
}

func readDataSourceSystemPropertyCategory(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	systemPropertyCategory := &client.SystemPropertyCategory{}
	if err := snowClient.GetObjectByName(client.EndpointSystemPropertyCategory, data.Get(systemPropertyCategoryName).(string), systemPropertyCategory); err != nil {
		data.SetId("")
		return err
	}

	resourceFromSystemPropertyCategory(data, systemPropertyCategory)

	return nil
}
