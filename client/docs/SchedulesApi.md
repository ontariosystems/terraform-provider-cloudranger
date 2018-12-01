# \SchedulesApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesAverageGet**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesAverageGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/schedules/average | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesAverageOptions**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesAverageOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/schedules/average | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesCountGet**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesCountGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/schedules/count | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesCountOptions**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesCountOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/schedules/count | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesGet**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/schedules | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesOptions**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/schedules | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesPost**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/schedules | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdDelete**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdDelete) | **Delete** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdGet**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdOptions**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdPut**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdPut) | **Put** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersGet**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id}/servers | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersOptions**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id}/servers | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersServerIdDelete**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersServerIdDelete) | **Delete** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id}/servers/{server_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersServerIdOptions**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersServerIdOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id}/servers/{server_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStartOptions**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStartOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id}/start | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStartPost**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStartPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id}/start | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStopOptions**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStopOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id}/stop | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStopPost**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStopPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id}/stop | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdTimesGet**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdTimesGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id}/times | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdTimesOptions**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdTimesOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/schedules/{schedule_id}/times | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsKeysGet**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsKeysGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/schedules/tags/keys | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsKeysOptions**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsKeysOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/schedules/tags/keys | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsValuesOptions**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsValuesOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/schedules/tags/values | 
[**OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsValuesPost**](SchedulesApi.md#OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsValuesPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/schedules/tags/values | 


# **OrganizationsOrganizationIdAccountsAccountIdSchedulesAverageGet**
> ScheduleAverage OrganizationsOrganizationIdAccountsAccountIdSchedulesAverageGet(ctx, accountId, organizationId)


Returns the average of all the schedules associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**ScheduleAverage**](ScheduleAverage.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesAverageOptions**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesAverageOptions(ctx, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesCountGet**
> Count OrganizationsOrganizationIdAccountsAccountIdSchedulesCountGet(ctx, accountId, organizationId)


Returns the count of all the schedules associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Count**](Count.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesCountOptions**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesCountOptions(ctx, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesGet**
> Schedule OrganizationsOrganizationIdAccountsAccountIdSchedulesGet(ctx, accountId, organizationId)


Returns all schedules associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesOptions**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesOptions(ctx, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesPost**
> Schedule OrganizationsOrganizationIdAccountsAccountIdSchedulesPost(ctx, accountId, organizationId, schedule)


Creates the schedules associated to the specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **schedule** | [**Schedule**](Schedule.md)|  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdDelete**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdDelete(ctx, scheduleId, accountId, organizationId)


Deletes the schedule with the given id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdGet**
> Schedule OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdGet(ctx, scheduleId, accountId, organizationId)


Returns the specific schedule associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdOptions**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdOptions(ctx, scheduleId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdPut**
> Schedule OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdPut(ctx, scheduleId, accountId, organizationId, schedule)


Updates the specific schedule associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **schedule** | [**Schedule**](Schedule.md)|  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersGet**
> Server OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersGet(ctx, scheduleId, accountId, organizationId)


Returns all the servers of the specific schedule associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Server**](Server.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersOptions**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersOptions(ctx, scheduleId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersServerIdDelete**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersServerIdDelete(ctx, serverId, scheduleId, accountId, organizationId)


Deletes the specific server associated to the given schedule of an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**|  | 
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersServerIdOptions**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdServersServerIdOptions(ctx, serverId, scheduleId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**|  | 
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStartOptions**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStartOptions(ctx, scheduleId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStartPost**
> Empty OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStartPost(ctx, scheduleId, accountId, organizationId, empty)


Starts the specific schedule associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **empty** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](empty.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStopOptions**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStopOptions(ctx, scheduleId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStopPost**
> Empty OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdStopPost(ctx, scheduleId, accountId, organizationId, empty)


Stops the specific schedule associated to the account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **empty** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](empty.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdTimesGet**
> ScheduleTime OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdTimesGet(ctx, scheduleId, accountId, organizationId)


Returns the number of times specific schedule is turned on associated to the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**ScheduleTime**](ScheduleTime.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdTimesOptions**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesScheduleIdTimesOptions(ctx, scheduleId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **scheduleId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsKeysGet**
> TagsKey OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsKeysGet(ctx, accountId, organizationId)


Returns the keys of tags of the schedules associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**TagsKey**](TagsKey.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsKeysOptions**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsKeysOptions(ctx, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsValuesOptions**
> OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsValuesOptions(ctx, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsValuesPost**
> Empty OrganizationsOrganizationIdAccountsAccountIdSchedulesTagsValuesPost(ctx, accountId, organizationId, tagsValueRequest)


Creates the values of the tags associated with the schedule of an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **tagsValueRequest** | [**TagsValueRequest**](TagsValueRequest.md)|  | 

### Return type

[**Empty**](empty.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

