# \LogsApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdLogCountGet**](LogsApi.md#OrganizationsOrganizationIdAccountsAccountIdLogCountGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/log_count | 
[**OrganizationsOrganizationIdAccountsAccountIdLogCountOptions**](LogsApi.md#OrganizationsOrganizationIdAccountsAccountIdLogCountOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/log_count | 
[**OrganizationsOrganizationIdAccountsAccountIdLogsGet**](LogsApi.md#OrganizationsOrganizationIdAccountsAccountIdLogsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/logs | 
[**OrganizationsOrganizationIdAccountsAccountIdLogsOptions**](LogsApi.md#OrganizationsOrganizationIdAccountsAccountIdLogsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/logs | 


# **OrganizationsOrganizationIdAccountsAccountIdLogCountGet**
> Count OrganizationsOrganizationIdAccountsAccountIdLogCountGet(ctx, accountId, organizationId)


Returns the log count of the provided account

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

# **OrganizationsOrganizationIdAccountsAccountIdLogCountOptions**
> OrganizationsOrganizationIdAccountsAccountIdLogCountOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdLogsGet**
> Empty OrganizationsOrganizationIdAccountsAccountIdLogsGet(ctx, accountId, organizationId)


Returns the logs of the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Empty**](empty.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdLogsOptions**
> OrganizationsOrganizationIdAccountsAccountIdLogsOptions(ctx, accountId, organizationId)


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

