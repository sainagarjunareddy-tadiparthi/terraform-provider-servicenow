package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sainagarjunareddy-tadiparthi/terraform-provider-servicenow/servicenow/client"
)

const systemPropertySuffix = "suffix"
const systemPropertyType = "type"
const systemPropertyChoices = "choices"
const systemPropertyIsPrivate = "is_private"
const systemPropertyIgnoreCache = "ignore_cache"
const systemPropertyDescription = "description"
const systemPropertyWriteRoles = "write_roles"
const systemPropertyReadRoles = "read_roles"
const systemPropertyName = "name"

// ResourceSystemProperty manages a System Property in ServiceNow.
func ResourceSystemProperty() *schema.Resource {
	return &schema.Resource{
		Create: createResourceSystemProperty,
		Read:   readResourceSystemProperty,
		Update: updateResourceSystemProperty,
		Delete: deleteResourceSystemProperty,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			systemPropertySuffix: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Internal suffix for the property used to create the name.",
			},
			systemPropertyType: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "string",
				Description: "Type of the property. This is used when displaying the input used in the UI to change the property.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{
						"string",
						"integer",
						"boolean",
						"choicelist",
						"color",
						"date_format",
						"image",
						"password",
						"password2",
						"short_string",
						"time_format",
						"timezone",
						"uploaded_image"})
					return
				},
			},
			systemPropertyChoices: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Comma-separated list of possible choices when the type is set to 'choicelist'. The values can be in format 'label=value' for alternate display labels.",
			},
			systemPropertyIsPrivate: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If set to 'true', this property will not move from one site to another.",
			},
			systemPropertyIgnoreCache: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If set to 'true', changing this property will not flush the cache.",
			},
			systemPropertyDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Short description of the property that will be displayed above it in the UI.",
			},
			systemPropertyWriteRoles: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Security roles required to modify this property.",
			},
			systemPropertyReadRoles: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Security roles required to read this property.",
			},
			systemPropertyName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Internal name of the property used to access it in scripts.",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceSystemProperty(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	systemProperty := &client.SystemProperty{}
	if err := snowClient.GetObject(client.EndpointSystemProperty, data.Id(), systemProperty); err != nil {
		data.SetId("")
		return err
	}

	resourceFromSystemProperty(data, systemProperty)

	return nil
}

func createResourceSystemProperty(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	systemProperty := resourceToSystemProperty(data)
	if err := snowClient.CreateObject(client.EndpointSystemProperty, systemProperty); err != nil {
		return err
	}

	resourceFromSystemProperty(data, systemProperty)

	return readResourceSystemProperty(data, serviceNowClient)
}

func updateResourceSystemProperty(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointSystemProperty, resourceToSystemProperty(data)); err != nil {
		return err
	}

	return readResourceSystemProperty(data, serviceNowClient)
}

func deleteResourceSystemProperty(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointSystemProperty, data.Id())
}

func resourceFromSystemProperty(data *schema.ResourceData, systemProperty *client.SystemProperty) {
	data.SetId(systemProperty.ID)
	data.Set(systemPropertySuffix, systemProperty.Suffix)
	data.Set(systemPropertyType, systemProperty.Type)
	data.Set(systemPropertyChoices, systemProperty.Choices)
	data.Set(systemPropertyIsPrivate, systemProperty.IsPrivate)
	data.Set(systemPropertyIgnoreCache, systemProperty.IgnoreCache)
	data.Set(systemPropertyDescription, systemProperty.Description)
	data.Set(systemPropertyWriteRoles, systemProperty.WriteRoles)
	data.Set(systemPropertyReadRoles, systemProperty.ReadRoles)
	data.Set(systemPropertyName, systemProperty.Name)
	data.Set(commonScope, systemProperty.Scope)
}

func resourceToSystemProperty(data *schema.ResourceData) *client.SystemProperty {
	systemProperty := client.SystemProperty{
		Suffix:      data.Get(systemPropertySuffix).(string),
		Type:        data.Get(systemPropertyType).(string),
		Choices:     data.Get(systemPropertyChoices).(string),
		IsPrivate:   data.Get(systemPropertyIsPrivate).(bool),
		IgnoreCache: data.Get(systemPropertyIgnoreCache).(bool),
		Description: data.Get(systemPropertyDescription).(string),
		WriteRoles:  data.Get(systemPropertyWriteRoles).(string),
		ReadRoles:   data.Get(systemPropertyReadRoles).(string),
	}
	systemProperty.ID = data.Id()
	systemProperty.Scope = data.Get(commonScope).(string)
	return &systemProperty
}
