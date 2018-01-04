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
        DefaultFunc: schema.EnvDefaultFunc("SENSU_BACKEND_HOST", "localhost"),
        Description: "The hostname of the Sensu backend API",
      },
      "username": {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("SENSU_BACKEND_USER", "admin"),
        Description: "The username to use for authenticating when accessing the Sensu backend API.",
      },
      "password": {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("SENSU_BACKEND_PASSWORD", "P@ssw0rd!"),
        Description: "The password to use for authenticating when accessing the Sensu backend API.",
      },
      "backend_port": {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("SENSU_BACKEND_PORT", "8080"),
        Description: "The port of the Sensu backend API",
      },
      "backend_protocol": {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("SENSU_BACKEND_PROTCOL", "http"),
        Description: "Protocol for accessing the Sensu API",
      },
      "authorization_bearer": {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("SENSU_BACKEND_AUTHORIZATION_BEARER", ""),
        Description: "Token for authenticating with the API",
      },
    },

    DataSourcesMap: map[string]*schema.Resource{

    },

    ResourcesMap: map[string]*schema.Resource{
      "sensu_check": resourceSensuCheck(),
    },

    ConfigureFunc: providerConfigure,
  }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
  config := Config{
    Token: d.Get("authorization_bearer").(string),
    Backend_Host: d.Get("backend_host").(string),
    Backend_Port: d.Get("backend_port").(string),
    Backend_Protocol: d.Get("backend_protocol").(string),
  }

  log.Println("[INFO] Initializing Sensu client")
  return config.Client()
}