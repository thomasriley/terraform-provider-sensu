package sensu

import (
        "bytes"
        "fmt"
        "log"
        "os"

        "github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
  return &schema.Provider{
    Schema: map[string]*schema.Schema{
      "backend_host": {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("SENSU_BACKEND_HOST", ""),
        Description: "The hostname of the Sensu backend API",
      },
      "username": {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("SENSU_BACKEND_USER", ""),
        Description: "The username to use for authenticating when accessing the Sensu backend API.",
      },
      "password": {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("SENSU_BACKEND_PASSWORD", ""),
        Description: "The password to use for authenticating when accessing the Sensu backend API.",
      },
      "backend_port": {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("SENSU_BACKEND_PORT", false),
        Description: "The port of the Sensu backend API",
      },
    },

    DataSourcesMap: map[string]*schema.Resource{
    },

    ResourcesMap: map[string]*schema.Resource{
    },
    ConfigureFunc: providerConfigure,
  }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

}

func tryLoadingConfigFile(d *schema.ResourceData) (*restclient.Config, error) {

}