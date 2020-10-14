package cloudranger

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cr "github.com/ontariosystems/terraform-provider-cloudranger/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDRANGER_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudranger_backup_policy": resourceCloudrangerBackupPolicy(),
			"cloudranger_schedule":      resourceCloudRangerSchedule(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	cfg := cr.NewConfiguration()
	cfg.AddDefaultHeader("x-api-key", d.Get("api_key").(string))
	cli := cr.NewAPIClient(cfg)

	token, _, err := cli.OrganizationApi.AuthorizeGet(context.Background())
	log.Println("[INFO] Cloudranger API authorized, adding bearer token...")

	if err != nil {
		log.Printf("[ERROR] Cloudranger Client authorization error: %v", err)
		return nil, diag.FromErr(err)
	}

	cfg.AddDefaultHeader("Authorization", "Bearer "+token.Token)
	log.Printf("[INFO] Cloudranger Client added bearer token successfully.")

	return cli, nil
}
