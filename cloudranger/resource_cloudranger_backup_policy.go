package cloudranger

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cr "github.com/ontariosystems/terraform-provider-cloudranger/client"
)

func resourceCloudrangerBackupPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudrangerBackupPolicyCreate,
		Read:          resourceCloudrangerBackupPolicyRead,
		Update:        resourceCloudrangerBackupPolicyUpdate,
		Delete:        resourceCloudrangerBackupPolicyDelete,
		// Exists: resourceCloudrangerBackupPolicyExists,

		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDRANGER_ORGANIZATION_ID", nil),
			},
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDRANGER_ACCOUNT_ID", nil),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timezone_location": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_source": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backup_target": {
				Type:     schema.TypeString,
				Required: true,
			},
			"create_cron": {
				Type:     schema.TypeString,
				Required: true,
			},
			"retention": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"perform_reboot": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"match_all_tags": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}
func resourceCloudrangerBackupPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cli := m.(*cr.APIClient)

	policy := buildBackupPolicy(d)
	createdPolicy, _, err := cli.PoliciesApi.OrganizationsOrganizationIdAccountsAccountIdPoliciesPost(
		ctx,
		policy.AccountId,
		policy.OrganizationId,
		*policy,
	)
	if err != nil {
		return diag.Errorf("error creating policy: %s", err.Error())
	}

	log.Printf("[Debug] created policy: %v", createdPolicy)

	d.SetId(strconv.Itoa(int(createdPolicy.PolicyId)))

	return nil
}

func resourceCloudrangerBackupPolicyRead(d *schema.ResourceData, m interface{}) error {
	cli := m.(*cr.APIClient)

	policyID := d.Id()

	attr, ok := d.GetOk("account_id")
	if !ok {
		return fmt.Errorf("account_id cannot be determined in terraform state")
	}
	accountID := attr.(string)

	attr, ok = d.GetOk("organization_id")
	if !ok {
		return fmt.Errorf("organization_id cannot be determined in terraform state")
	}
	organizationID := attr.(string)

	policy, _, err := cli.PoliciesApi.OrganizationsOrganizationIdAccountsAccountIdPoliciesPolicyIdGet(
		context.Background(),
		policyID,
		accountID,
		organizationID,
	)

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] policy: %v", policy)

	if err := d.Set("account_id", policy.AccountId); err != nil {
		return err
	}
	if err := d.Set("organization_id", policy.OrganizationId); err != nil {
		return err
	}
	if err := d.Set("name", policy.Name); err != nil {
		return err
	}
	if err := d.Set("description", policy.Description); err != nil {
		return err
	}
	if err := d.Set("timezone_location", policy.TimezoneLocation); err != nil {
		return err
	}
	if err := d.Set("backup_source", policy.BackupSource); err != nil {
		return err
	}
	if err := d.Set("backup_target", policy.BackupTarget); err != nil {
		return err
	}
	if err := d.Set("create_cron", policy.CreateCron); err != nil {
		return err
	}
	if err := d.Set("retention", int(policy.RetentionInHours)); err != nil {
		return err
	}
	if err := d.Set("active", policy.IsActive); err != nil {
		return err
	}
	if err := d.Set("match_all_tags", policy.MatchAllTagsFlag); err != nil {
		return err
	}
	if err := d.Set("perform_reboot", policy.PerformReboot); err != nil {
		return err
	}

	return loadTags(policy.Tags, d)
}

func resourceCloudrangerBackupPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	cli := m.(*cr.APIClient)

	policy := buildBackupPolicy(d)
	policyID := d.Id()

	attr, ok := d.GetOk("account_id")
	if !ok {
		return fmt.Errorf("account_id cannot be determined in terraform state")
	}
	accountID := attr.(string)

	attr, ok = d.GetOk("organization_id")
	if !ok {
		return fmt.Errorf("organization_id cannot be determined in terraform state")
	}
	organizationID := attr.(string)

	updatedPolicy, resp, err := cli.PoliciesApi.OrganizationsOrganizationIdAccountsAccountIdPoliciesPolicyIdPut(
		context.Background(),
		policyID,
		accountID,
		organizationID,
		*policy,
	)
	if resp.StatusCode != 204 {
		return fmt.Errorf("error updating policy: %s, %v", err.Error(), resp)
	}

	log.Printf("[DEBUG] updated policy: %v", updatedPolicy)

	return resourceCloudrangerBackupPolicyRead(d, m)
}

func resourceCloudrangerBackupPolicyDelete(d *schema.ResourceData, m interface{}) error {
	cli := m.(*cr.APIClient)

	policyID := d.Id()

	attr, ok := d.GetOk("account_id")
	if !ok {
		return fmt.Errorf("account_id cannot be determined in terraform state")
	}
	accountID := attr.(string)

	attr, ok = d.GetOk("organization_id")
	if !ok {
		return fmt.Errorf("organization_id cannot be determined in terraform state")
	}
	organizationID := attr.(string)

	_, err := cli.PoliciesApi.OrganizationsOrganizationIdAccountsAccountIdPoliciesPolicyIdDelete(
		context.Background(),
		policyID,
		accountID,
		organizationID,
	)

	if err != nil {
		return fmt.Errorf("policy %s cannot be deleted: %v", policyID, err)
	}

	return nil
}

func buildBackupPolicy(d *schema.ResourceData) *cr.Policy {
	var policy cr.Policy

	if attr, ok := d.GetOk("organization_id"); ok {
		policy.OrganizationId = attr.(string)
	}
	if attr, ok := d.GetOk("account_id"); ok {
		policy.AccountId = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		policy.Name = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		policy.Description = attr.(string)
	}
	if attr, ok := d.GetOk("timezone_location"); ok {
		policy.TimezoneLocation = attr.(string)
	}
	if attr, ok := d.GetOk("backup_source"); ok {
		policy.BackupSource = attr.(string)
	}
	if attr, ok := d.GetOk("backup_target"); ok {
		policy.BackupTarget = attr.(string)
	}
	if attr, ok := d.GetOk("active"); ok {
		policy.IsActive = attr.(bool)
	}
	if attr, ok := d.GetOk("match_all_tags"); ok {
		policy.MatchAllTagsFlag = attr.(bool)
	}
	if attr, ok := d.GetOk("create_cron"); ok {
		policy.CreateCron = attr.(string)
	}
	if attr, ok := d.GetOk("retention"); ok {
		policy.RetentionUnit = "hour"
		policy.RetentionInHours = float32(attr.(int))
		policy.RetentionAmount = float32(attr.(int))
	}
	if attr, ok := d.GetOk("perform_reboot"); ok {
		policy.PerformReboot = attr.(bool)
	}

	policy.Tags = tagsFrom(d)
	return &policy
}

// func resourceCloudrangerBackupPolicyExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {

// 	return true, nil
// }
