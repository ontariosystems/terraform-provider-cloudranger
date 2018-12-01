package cloudranger

import (
	"context"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	cr "github.com/jaxi/terraform-provider-cloudranger/client"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDRANGER_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudranger_backup_policy": resourceCloudrangerBackupPolicy(),
			"cloudranger_schedule":      resourceCloudRangerSchedule(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	cfg := cr.NewConfiguration()
	cfg.AddDefaultHeader("x-api-key", d.Get("api_key").(string))
	cli := cr.NewAPIClient(cfg)

	token, _, err := cli.OrganizationApi.AuthorizeGet(context.Background())
	log.Println("[INFO] Cloudranger API authorized, adding bearer token...")

	if err != nil {
		log.Printf("[ERROR] Cloudranger Client authorization error: %v", err)
		return nil, err
	}

	cfg.AddDefaultHeader("Authorization", "Bearer "+token.Token)
	log.Printf("[INFO] Cloudranger Client added bearer token successfully.")

	return cli, nil
}
