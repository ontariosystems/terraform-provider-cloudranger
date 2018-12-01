# \GlobalsApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdGlobalsSearchQueryGet**](GlobalsApi.md#OrganizationsOrganizationIdAccountsAccountIdGlobalsSearchQueryGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/globals/search/{query} | 
[**OrganizationsOrganizationIdAccountsAccountIdGlobalsSearchQueryOptions**](GlobalsApi.md#OrganizationsOrganizationIdAccountsAccountIdGlobalsSearchQueryOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/globals/search/{query} | 


# **OrganizationsOrganizationIdAccountsAccountIdGlobalsSearchQueryGet**
> SearchQuery OrganizationsOrganizationIdAccountsAccountIdGlobalsSearchQueryGet(ctx, query, accountId, organizationId)


Returns the result of the query after running the query

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **query** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**SearchQuery**](SearchQuery.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdGlobalsSearchQueryOptions**
> OrganizationsOrganizationIdAccountsAccountIdGlobalsSearchQueryOptions(ctx, query, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **query** | **string**|  | 
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

