/*
 * CloudRanger API
 *
 *  # Welcome to CloudRanger  Here are some instructions on  how to use the API  ## Authentication and Authorization   * Access to the API is controlled by your API key generated in the CloudRanger dashboard and token  * The token is returned by calling the /authorize endpoint and supplying the x-api-key header  * All other calls use the x-api-key header and the Authorzation header with Bearer followed by your token
 *
 * API version: 2018-05
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloudranger

type ScheduleTime struct {
	Day string `json:"day,omitempty"`

	StartHour string `json:"start_hour,omitempty"`

	StopHour string `json:"stop_hour,omitempty"`

	StartCron string `json:"start_cron,omitempty"`

	StopCron string `json:"stop_cron,omitempty"`
}