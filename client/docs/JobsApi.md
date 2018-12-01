# \JobsApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdJobsGet**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/jobs | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdGet**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdLogsGet**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdLogsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id}/logs | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdLogsOptions**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdLogsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id}/logs | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdOptions**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdRetryOptions**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdRetryOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id}/retry | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdRetryPost**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdRetryPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id}/retry | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksGet**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id}/tasks | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksOptions**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id}/tasks | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdLogsGet**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdLogsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id}/tasks/{task_id}/logs | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdLogsOptions**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdLogsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id}/tasks/{task_id}/logs | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdStepsStepIdRetryOptions**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdStepsStepIdRetryOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id}/tasks/{task_id}/steps/{step_id}/retry | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdStepsStepIdRetryPost**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdStepsStepIdRetryPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/jobs/{job_id}/tasks/{task_id}/steps/{step_id}/retry | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsOptions**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/jobs | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsRangeMinSeqCountGet**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsRangeMinSeqCountGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/jobs/range/{min_seq}/{count} | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsRangeMinSeqCountOptions**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsRangeMinSeqCountOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/jobs/range/{min_seq}/{count} | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsSearchOptions**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsSearchOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/jobs/search | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsSearchPost**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsSearchPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/jobs/search | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyGet**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/jobs/tags/{tag_key} | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyOptions**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/jobs/tags/{tag_key} | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyTagValueGet**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyTagValueGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/jobs/tags/{tag_key}/{tag_value} | 
[**OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyTagValueOptions**](JobsApi.md#OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyTagValueOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/jobs/tags/{tag_key}/{tag_value} | 


# **OrganizationsOrganizationIdAccountsAccountIdJobsGet**
> Job OrganizationsOrganizationIdAccountsAccountIdJobsGet(ctx, accountId, organizationId)


Returns the list of jobs associated to the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Job**](Job.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdGet**
> Job OrganizationsOrganizationIdAccountsAccountIdJobsJobIdGet(ctx, jobId, accountId, organizationId)


Returns the job associated with the given job id of the provided account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Job**](Job.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdLogsGet**
> Job OrganizationsOrganizationIdAccountsAccountIdJobsJobIdLogsGet(ctx, jobId, accountId, organizationId)


Returns the logs of the given job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Job**](Job.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdLogsOptions**
> OrganizationsOrganizationIdAccountsAccountIdJobsJobIdLogsOptions(ctx, jobId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdOptions**
> OrganizationsOrganizationIdAccountsAccountIdJobsJobIdOptions(ctx, jobId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdRetryOptions**
> OrganizationsOrganizationIdAccountsAccountIdJobsJobIdRetryOptions(ctx, jobId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdRetryPost**
> Empty OrganizationsOrganizationIdAccountsAccountIdJobsJobIdRetryPost(ctx, jobId, accountId, organizationId, empty)


Retries the job of the given id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksGet**
> JobTask OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksGet(ctx, jobId, accountId, organizationId)


Returns all the tasks associated with the given job id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**JobTask**](JobTask.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksOptions**
> OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksOptions(ctx, jobId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdLogsGet**
> JobTask OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdLogsGet(ctx, jobId, taskId, accountId, organizationId)


Returns the logs of a specific task associated with the provided job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
  **taskId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**JobTask**](JobTask.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdLogsOptions**
> OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdLogsOptions(ctx, jobId, taskId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
  **taskId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdStepsStepIdRetryOptions**
> OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdStepsStepIdRetryOptions(ctx, jobId, taskId, stepId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
  **taskId** | **string**|  | 
  **stepId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdStepsStepIdRetryPost**
> Empty OrganizationsOrganizationIdAccountsAccountIdJobsJobIdTasksTaskIdStepsStepIdRetryPost(ctx, jobId, taskId, stepId, accountId, organizationId, empty)


Retries the specific step in the log of assocaited task of the provided job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **string**|  | 
  **taskId** | **string**|  | 
  **stepId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdJobsOptions**
> OrganizationsOrganizationIdAccountsAccountIdJobsOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdJobsRangeMinSeqCountGet**
> Job OrganizationsOrganizationIdAccountsAccountIdJobsRangeMinSeqCountGet(ctx, count, minSeq, accountId, organizationId)


Returns the job list of count and range specified in the URL for the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **count** | **string**|  | 
  **minSeq** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Job**](Job.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdJobsRangeMinSeqCountOptions**
> OrganizationsOrganizationIdAccountsAccountIdJobsRangeMinSeqCountOptions(ctx, count, minSeq, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **count** | **string**|  | 
  **minSeq** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdJobsSearchOptions**
> OrganizationsOrganizationIdAccountsAccountIdJobsSearchOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdJobsSearchPost**
> Job OrganizationsOrganizationIdAccountsAccountIdJobsSearchPost(ctx, accountId, organizationId, jobSearchRequest)


Returns the existing jobs of the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **jobSearchRequest** | [**JobSearchRequest**](JobSearchRequest.md)|  | 

### Return type

[**Job**](Job.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyGet**
> TagsKey OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyGet(ctx, accountId, tagKey, organizationId)


Returns the tag key of the specific job of the provided account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **tagKey** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**TagsKey**](TagsKey.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyOptions**
> OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyOptions(ctx, accountId, tagKey, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **tagKey** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyTagValueGet**
> TagsValue OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyTagValueGet(ctx, tagValue, tagKey, accountId, organizationId)


Returns the tag value of the tag key related to specific job of the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **tagValue** | **string**|  | 
  **tagKey** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**TagsValue**](TagsValue.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyTagValueOptions**
> OrganizationsOrganizationIdAccountsAccountIdJobsTagsTagKeyTagValueOptions(ctx, tagValue, tagKey, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **tagValue** | **string**|  | 
  **tagKey** | **string**|  | 
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

