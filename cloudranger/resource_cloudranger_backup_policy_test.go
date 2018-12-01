package cloudranger

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	cr "github.com/jaxi/terraform-provider-cloudranger/client"
)

const testAccCheckCloudrangerBackupPolicyConfig = `
resource "cloudranger_backup_policy" "basic" {
  name = "basic backup policy"
  description = "basic backup policy - managed by cloudranger!"
  timezone_location = "Europe/London"
  backup_source = "instances"
  backup_target = "snapshots"
  retention = 24
  create_cron = "50 18 * * *"

	tags {
		key = "Terraform"
		value = "1"
	}

	tags {
		key = "Environment"
		value = "staging"
  }
}
`

const testAccCheckCloudrangerBackupPolicyUpdated = `
resource "cloudranger_backup_policy" "basic" {
  name = "basic backup policy"
  description = "basic backup policy - managed by cloudranger!"
  timezone_location = "Europe/London"
  backup_source = "instances"
  backup_target = "snapshots"
  retention = 36
  create_cron = "0 21 * * *"

	match_all_tags = true
	active = false
	perform_reboot = true

	tags {
		key = "Environment"
		value = "staging"
  }
}
`

func TestAccCloudrangerBackupPolicy_Basic(t *testing.T) {
	resourceName := "cloudranger_backup_policy.basic"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccChecCloudrangerBackupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCloudrangerBackupPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "basic backup policy"),
					resource.TestCheckResourceAttr(resourceName, "description", "basic backup policy - managed by cloudranger!"),
					resource.TestCheckResourceAttr(resourceName, "timezone_location", "Europe/London"),
					resource.TestCheckResourceAttr(resourceName, "backup_source", "instances"),
					resource.TestCheckResourceAttr(resourceName, "backup_target", "snapshots"),
				),
			},
		},
	})
}

func TestAccCloudrangerBackupPolicy_Updated(t *testing.T) {
	resourceName := "cloudranger_backup_policy.basic"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccChecCloudrangerBackupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCloudrangerBackupPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "basic backup policy"),
					resource.TestCheckResourceAttr(resourceName, "retention", "24"),
					resource.TestCheckResourceAttr(resourceName, "create_cron", "50 18 * * *"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "perform_reboot", "false"),
					resource.TestCheckResourceAttr(resourceName, "match_all_tags", "false"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "Terraform"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.key", "Environment"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.value", "staging"),
				),
			},
			resource.TestStep{
				Config: testAccCheckCloudrangerBackupPolicyUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "basic backup policy"),
					resource.TestCheckResourceAttr(resourceName, "retention", "36"),
					resource.TestCheckResourceAttr(resourceName, "create_cron", "0 21 * * *"),
					resource.TestCheckResourceAttr(resourceName, "active", "false"),
					resource.TestCheckResourceAttr(resourceName, "perform_reboot", "true"),
					resource.TestCheckResourceAttr(resourceName, "match_all_tags", "true"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "Environment"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", "staging"),
				),
			},
		},
	})
}

func testAccChecCloudrangerBackupPolicyDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cr.APIClient)

	for i, r := range s.RootModule().Resources {
		fmt.Printf("Resource %s type = %s\n", i, r.Type)

		if r.Type != "cloudranger_backup_policy" {
			continue
		}

		id := r.Primary.ID
		accountID := r.Primary.Attributes["account_id"]
		organisationID := r.Primary.Attributes["organisation_id"]

		policy, resp, err := client.PoliciesApi.OrganizationsOrganizationIdAccountsAccountIdPoliciesPolicyIdGet(
			context.Background(),
			id,
			accountID,
			organisationID,
		)

		if err != nil {
			// 403??? yeah weird I know
			if resp.StatusCode == 404 || resp.StatusCode == 403 {
				continue
			}
			return fmt.Errorf("Cannot retrieve policy %s", err)
		}

		if !policy.IsActive {
			continue
		}
		return fmt.Errorf("Backup policy %s still exists", id)
	}

	return nil
}
