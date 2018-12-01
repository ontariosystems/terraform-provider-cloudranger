package cloudranger

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	cr "github.com/jaxi/terraform-provider-cloudranger/client"
)

const testAccCheckCloudrangerScheduleConfig = `
resource "cloudranger_schedule" "basic_schedule"{
  name = "basic schedule"
  description = "basic schedule provisioned by terraform"
  timezone_location = "Europe/London"

  schedule {
		monday {
			start = 9
			end = 19
		}
		tuesday {
			start = 9
			end = 20
		}
  }

  active = true
  match_all_tags = false

  tags {
		key = "workload-type"
		value = "other"
	}

  tags {
		key = "Terraform"
		value = "1"
	}
}
`

const testAccCheckCloudrangerScheduleConfigUpdated = `
resource "cloudranger_schedule" "basic_schedule"{
  name = "basic schedule"
  description = "basic schedule provisioned by terraform"
  timezone_location = "Europe/London"

  schedule {
		tuesday {
			start = 8
			end = 15
		}

		tuesday {
			start = 16
			end = -1
			action_type = "Reboot"
		}

		tuesday {
			start = 18
			end = 22
		}
  }

  active = false
  match_all_tags = true

  tags {
		key = "Terraform"
		value = "1"
	}
}
`

func TestAccCloudrangerSchedule_Basic(t *testing.T) {
	resourceName := "cloudranger_schedule.basic_schedule"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccChecCloudrangerScheduleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCloudrangerScheduleConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "basic schedule"),
					resource.TestCheckResourceAttr(resourceName, "description", "basic schedule provisioned by terraform"),
					resource.TestCheckResourceAttr(resourceName, "timezone_location", "Europe/London"),
				),
			},
		},
	})
}

func TestAccCloudrangerSchedule_Update(t *testing.T) {
	resourceName := "cloudranger_schedule.basic_schedule"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccChecCloudrangerScheduleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCloudrangerScheduleConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "basic schedule"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "match_all_tags", "false"),

					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "workload-type"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", "other"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.key", "Terraform"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.value", "1"),

					resource.TestCheckResourceAttr(resourceName, "schedule.0.monday.0.start", "9"),
					resource.TestCheckResourceAttr(resourceName, "schedule.0.monday.0.end", "19"),
					resource.TestCheckResourceAttr(resourceName, "schedule.0.monday.0.action_type", "Start-Stop"),

					resource.TestCheckResourceAttr(resourceName, "schedule.0.tuesday.0.start", "9"),
					resource.TestCheckResourceAttr(resourceName, "schedule.0.tuesday.0.end", "20"),
					resource.TestCheckResourceAttr(resourceName, "schedule.0.tuesday.0.action_type", "Start-Stop"),
				),
			},
			resource.TestStep{
				Config: testAccCheckCloudrangerScheduleConfigUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "basic schedule"),
					resource.TestCheckResourceAttr(resourceName, "active", "false"),
					resource.TestCheckResourceAttr(resourceName, "match_all_tags", "true"),

					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "Terraform"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", "1"),

					resource.TestCheckResourceAttr(resourceName, "schedule.0.tuesday.0.start", "8"),
					resource.TestCheckResourceAttr(resourceName, "schedule.0.tuesday.0.end", "15"),
					resource.TestCheckResourceAttr(resourceName, "schedule.0.tuesday.0.action_type", "Start-Stop"),

					resource.TestCheckResourceAttr(resourceName, "schedule.0.tuesday.1.start", "16"),
					resource.TestCheckResourceAttr(resourceName, "schedule.0.tuesday.1.action_type", "Reboot"),

					resource.TestCheckResourceAttr(resourceName, "schedule.0.tuesday.2.start", "18"),
					resource.TestCheckResourceAttr(resourceName, "schedule.0.tuesday.2.end", "22"),
					resource.TestCheckResourceAttr(resourceName, "schedule.0.tuesday.2.action_type", "Start-Stop"),
				),
			},
		},
	})
}

func testAccChecCloudrangerScheduleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cr.APIClient)

	for i, r := range s.RootModule().Resources {
		fmt.Printf("Resource %s type = %s\n", i, r.Type)

		if r.Type != "cloudranger_schedule" {
			continue
		}

		id := r.Primary.ID
		accountID := r.Primary.Attributes["account_id"]
		organisationID := r.Primary.Attributes["organisation_id"]

		schedule, resp, err := client.SchedulesApi.OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdGet(
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
			return fmt.Errorf("Cannot retrieve schedule %s", err)
		}

		if !schedule.IsActive {
			continue
		}
		return fmt.Errorf("Schedule %s still exists", id)
	}

	return nil
}
