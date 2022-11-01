package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sainagarjunareddy-tadiparthi/terraform-provider-servicenow/servicenow/client"
)

const scriptIncludeName = "name"
const scriptIncludeClientCallable = "client_callable"
const scriptIncludeDescription = "description"
const scriptIncludeScript = "script"
const scriptIncludeActive = "active"
const scriptIncludeAccess = "access"
const scriptIncludeAPIName = "api_name"

// ResourceScriptInclude manages a Script Include in ServiceNow.
func ResourceScriptInclude() *schema.Resource {
	return &schema.Resource{
		Create: createResourceScriptInclude,
		Read:   readResourceScriptInclude,
		Update: updateResourceScriptInclude,
		Delete: deleteResourceScriptInclude,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			scriptIncludeName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Display name of the script. Needed to have an api_name.",
			},
			scriptIncludeScript: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Javascript script to run when this Script Include is called.",
			},
			scriptIncludeDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Describe what the script does.",
			},
			scriptIncludeClientCallable: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not this script can be called from the client-side or only server-side.",
			},
			scriptIncludeActive: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not this Script Include is enabled.",
			},
			scriptIncludeAccess: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "package_private",
				Description: "Whether this Script can be accessed from only this application scope or all application scopes. Values can be 'package_private' or 'public'.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"package_private", "public"})
					return
				},
			},
			scriptIncludeAPIName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Full name of the Script Include needed to call it.",
			},
			commonProtectionPolicy: getProtectionPolicySchema(),
			commonScope:            getScopeSchema(),
		},
	}
}

func readResourceScriptInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	scriptInclude := &client.ScriptInclude{}
	if err := snowClient.GetObject(client.EndpointScriptInclude, data.Id(), scriptInclude); err != nil {
		data.SetId("")
		return err
	}

	resourceFromScriptInclude(data, scriptInclude)

	return nil
}

func createResourceScriptInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	scriptInclude := resourceToScriptInclude(data)
	if err := snowClient.CreateObject(client.EndpointScriptInclude, scriptInclude); err != nil {
		return err
	}

	resourceFromScriptInclude(data, scriptInclude)

	return readResourceScriptInclude(data, serviceNowClient)
}

func updateResourceScriptInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointScriptInclude, resourceToScriptInclude(data)); err != nil {
		return err
	}

	return readResourceScriptInclude(data, serviceNowClient)
}

func deleteResourceScriptInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointScriptInclude, data.Id())
}

func resourceFromScriptInclude(data *schema.ResourceData, scriptInclude *client.ScriptInclude) {
	data.SetId(scriptInclude.ID)
	data.Set(scriptIncludeName, scriptInclude.Name)
	data.Set(scriptIncludeClientCallable, scriptInclude.ClientCallable)
	data.Set(scriptIncludeDescription, scriptInclude.Description)
	data.Set(scriptIncludeScript, scriptInclude.Script)
	data.Set(scriptIncludeActive, scriptInclude.Active)
	data.Set(scriptIncludeAccess, scriptInclude.Access)
	data.Set(scriptIncludeAPIName, scriptInclude.APIName)
	data.Set(commonProtectionPolicy, scriptInclude.ProtectionPolicy)
	data.Set(commonScope, scriptInclude.Scope)
}

func resourceToScriptInclude(data *schema.ResourceData) *client.ScriptInclude {
	scriptInclude := client.ScriptInclude{
		Name:           data.Get(scriptIncludeName).(string),
		ClientCallable: data.Get(scriptIncludeClientCallable).(bool),
		Description:    data.Get(scriptIncludeDescription).(string),
		Script:         data.Get(scriptIncludeScript).(string),
		Active:         data.Get(scriptIncludeActive).(bool),
		Access:         data.Get(scriptIncludeAccess).(string),
	}
	scriptInclude.ID = data.Id()
	scriptInclude.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	scriptInclude.Scope = data.Get(commonScope).(string)
	return &scriptInclude
}
