package cloudranger

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cr "github.com/ontariosystems/terraform-provider-cloudranger/client"
)

func resourceCloudRangerSchedule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudrangerScheduleCreate,
		Read:          resourceCloudrangerScheduleRead,
		Update:        resourceCloudrangerScheduleUpdate,
		Delete:        resourceCloudrangerScheduleDelete,

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
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "server",
			},
			"schedule": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monday":    scheduleElem(),
						"tuesday":   scheduleElem(),
						"wednesday": scheduleElem(),
						"thursday":  scheduleElem(),
						"friday":    scheduleElem(),
						"saturday":  scheduleElem(),
						"sunday":    scheduleElem(),
					},
				},
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

func scheduleElem() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"start": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"end": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"action_type": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  "Start-Stop",
				},
			},
		},
	}
}

func resourceCloudrangerScheduleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cli := m.(*cr.APIClient)

	schedule, err := buildSchedule(d)
	if err != nil {
		return diag.Errorf("error creating schedule: %s", err.Error())
	}

	createdSchedule, _, err := cli.SchedulesApi.OrganizationsOrganizationIdAccountsAccountIdSchedulesPost(
		ctx,
		schedule.AccountId,
		schedule.OrganizationId,
		*schedule,
	)
	if err != nil {
		return diag.Errorf("error creating schedule: %s", err.Error())
	}

	log.Printf("[Debug] created schedule: %v", createdSchedule)

	d.SetId(strconv.Itoa(int(createdSchedule.Id)))

	return nil
}

func resourceCloudrangerScheduleRead(d *schema.ResourceData, m interface{}) error {
	cli := m.(*cr.APIClient)

	scheduleID := d.Id()

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

	schedule, _, err := cli.SchedulesApi.OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdGet(
		context.Background(),
		scheduleID,
		accountID,
		organizationID,
	)

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] schedule: %v", schedule)

	if err := d.Set("account_id", schedule.AccountId); err != nil {
		return err
	}
	if err := d.Set("organization_id", schedule.OrganizationId); err != nil {
		return err
	}
	if err := d.Set("name", schedule.Name); err != nil {
		return err
	}
	if err := d.Set("description", schedule.Description); err != nil {
		return err
	}
	if err := d.Set("timezone_location", schedule.TimezoneLocation); err != nil {
		return err
	}
	if err := d.Set("active", schedule.IsActive); err != nil {
		return err
	}
	if err := d.Set("type", schedule.Type_); err != nil {
		return err
	}
	if err := d.Set("match_all_tags", schedule.MatchAllTagsFlag); err != nil {
		return err
	}
	if err := loadTags(schedule.Tags, d); err != nil {
		return err
	}

	var scheduleData crSchedule
	if err := json.Unmarshal([]byte(schedule.ScheduleJson), &scheduleData); err != nil {
		return err
	}

	log.Printf("[DEBUG] schedule data: %v", scheduleData)
	return scheduleLoad(scheduleData, d)
}

func resourceCloudrangerScheduleUpdate(d *schema.ResourceData, m interface{}) error {
	cli := m.(*cr.APIClient)

	schedule, err := buildSchedule(d)
	if err != nil {
		return fmt.Errorf("error creating schedule: %s", err.Error())
	}

	scheduleID := d.Id()

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

	updatedSchedule, resp, err := cli.SchedulesApi.OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdPut(
		context.Background(),
		scheduleID,
		accountID,
		organizationID,
		*schedule,
	)

	if resp.StatusCode != 204 {
		return fmt.Errorf("error updating schedule: %s, %v", err.Error(), resp)
	}

	log.Printf("[DEBUG] updated schedule: %v", updatedSchedule)
	return resourceCloudrangerScheduleRead(d, m)
}

func resourceCloudrangerScheduleDelete(d *schema.ResourceData, m interface{}) error {
	cli := m.(*cr.APIClient)

	scheduleID := d.Id()

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

	_, err := cli.SchedulesApi.OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdDelete(
		context.Background(),
		scheduleID,
		accountID,
		organizationID,
	)

	if err != nil {
		return fmt.Errorf("schedule %s cannot be deleted: %v", scheduleID, err)
	}

	return nil
}

func buildSchedule(d *schema.ResourceData) (*cr.Schedule, error) {
	schedule := cr.Schedule{}

	if attr, ok := d.GetOk("organization_id"); ok {
		schedule.OrganizationId = attr.(string)
	}
	if attr, ok := d.GetOk("account_id"); ok {
		schedule.AccountId = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		schedule.Name = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		schedule.Description = attr.(string)
	}
	if attr, ok := d.GetOk("timezone_location"); ok {
		schedule.TimezoneLocation = attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		schedule.Type_ = attr.(string)
	}
	if attr, ok := d.GetOk("active"); ok {
		schedule.IsActive = attr.(bool)
	}
	if attr, ok := d.GetOk("match_all_tags"); ok {
		schedule.MatchAllTagsFlag = attr.(bool)
	}
	schedule.Tags = tagsFrom(d)

	b, err := json.Marshal(scheduleFrom(d))

	if err != nil {
		return nil, err
	}

	schedule.ScheduleJson = string(b)
	return &schedule, nil
}

type crScheduleEntity struct {
	Start      int    `json:"start"`
	End        int    `json:"end"`
	ActionType string `json:"action_type"`
}

type crDailySchedule []crScheduleEntity
type crSchedule map[string]crDailySchedule

func scheduleFrom(d *schema.ResourceData) *crSchedule {
	schedulePayload := &crSchedule{}

	attr, ok := d.GetOk("schedule")
	if !ok {
		return schedulePayload
	}
	scheduleItems := attr.([]interface{})

	if len(scheduleItems) == 0 {
		return schedulePayload
	}

	scheduleItem := scheduleItems[0].(map[string]interface{})

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for _, day := range days {
		var dailySchedule crDailySchedule
		if schedules, ok := scheduleItem[strings.ToLower(day)]; ok {
			dailySchedule = make(crDailySchedule, len(schedules.([]interface{})))
			for i, schedule := range schedules.([]interface{}) {
				scheduleMap := schedule.(map[string]interface{})
				dailySchedule[i] = crScheduleEntity{
					Start:      scheduleMap["start"].(int),
					End:        scheduleMap["end"].(int),
					ActionType: scheduleMap["action_type"].(string),
				}
			}
		}

		(*schedulePayload)[day] = dailySchedule
	}
	return schedulePayload
}

func scheduleLoad(schedule crSchedule, d *schema.ResourceData) error {
	scheduleData := map[string]interface{}{}

	for day, dailyschedules := range schedule {
		scheduleDaily := make([]map[string]interface{}, len(dailyschedules))
		for i, scheduleEntity := range dailyschedules {
			scheduleDaily[i] = map[string]interface{}{
				"start":       scheduleEntity.Start,
				"end":         scheduleEntity.End,
				"action_type": scheduleEntity.ActionType,
			}
		}
		scheduleData[strings.ToLower(day)] = scheduleDaily
	}

	return d.Set("schedule", []map[string]interface{}{scheduleData})
}
