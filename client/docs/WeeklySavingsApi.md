# \WeeklySavingsApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdWeeklySavingsGet**](WeeklySavingsApi.md#OrganizationsOrganizationIdAccountsAccountIdWeeklySavingsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/weekly_savings | 
[**OrganizationsOrganizationIdAccountsAccountIdWeeklySavingsOptions**](WeeklySavingsApi.md#OrganizationsOrganizationIdAccountsAccountIdWeeklySavingsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/weekly_savings | 


# **OrganizationsOrganizationIdAccountsAccountIdWeeklySavingsGet**
> WeeklySaving OrganizationsOrganizationIdAccountsAccountIdWeeklySavingsGet(ctx, accountId, organizationId)


Returns the weekly savings of an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**WeeklySaving**](WeeklySaving.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdWeeklySavingsOptions**
> OrganizationsOrganizationIdAccountsAccountIdWeeklySavingsOptions(ctx, accountId, organizationId)


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

