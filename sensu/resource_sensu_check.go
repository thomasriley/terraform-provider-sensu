package sensu

import (
  "log"
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"

  "github.com/hashicorp/terraform/helper/schema"

)

func resourceSensuCheck() *schema.Resource {
  return &schema.Resource{
    Create: resourceSensuCheckCreate,
    Read:   resourceSensuCheckRead,
    Update: resourceSensuCheckUpdate,
    Delete: resourceSensuCheckDelete,
    Importer: &schema.ResourceImporter{
      State: schema.ImportStatePassthrough,
    },
    Schema: map[string]*schema.Schema{
      "check_name": {
        Type:     schema.TypeString,
        Required: true,
      },
    },
  }
}

func resourceSensuCheckCreate(d *schema.ResourceData, meta interface{}) error {
  client := meta.(*sensu.Client)
  httpclient.Defaults(httpclient.Map {
    httpclient.OPT_USERAGENT: "my awsome httpclient",
    "Authorization": client.Token,
  }

  log.Printf("[INFO] Creating Sensu check %s", d.Get("check_name"))

  url := client.Backend_Protocol + "://" + client.Backend_Host + ":" + client.Backend_Port + "/checks"

  payload := strings.NewReader("{\n  \"check_hooks\": null,\n  \"environment\": \"default\",\n  \"handlers\": [],\n  \"interval\": 60,\n  \"name\": \"" + d.Get("check_name") + "\",\n  \"organization\": \"default\",\n  \"proxy_entity_id\": \"\",\n  \"publish\": true,\n  \"runtime_assets\": [],\n  \"subdue\": null,\n  \"subscriptions\": []\n}")

  req, _ := http.NewRequest("POST", url, payload)

  req.Header.Add("authorization", client.Token)
  req.Header.Add("cache-control", "no-cache")

  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  d.SetId(d.Get("check_name"))
  return resourcePagerDutyAddonRead(d, meta)

}

func resourceSensuCheckRead(d *schema.ResourceData, meta interface{}) error {
  client := meta.(*sensu.Client)

}

func resourceSensuCheckUpdate(d *schema.ResourceData, meta interface{}) error {
  client := meta.(*sensu.Client)

}

func resourceSensuCheckDelete(d *schema.ResourceData, meta interface{}) error {
  client := meta.(*sensu.Client)

}