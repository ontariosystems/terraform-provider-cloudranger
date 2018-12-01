# \SynchronizationApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdSynchronizationOptions**](SynchronizationApi.md#OrganizationsOrganizationIdAccountsAccountIdSynchronizationOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/synchronization | 
[**OrganizationsOrganizationIdAccountsAccountIdSynchronizationPost**](SynchronizationApi.md#OrganizationsOrganizationIdAccountsAccountIdSynchronizationPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/synchronization | 
[**OrganizationsOrganizationIdAccountsAccountIdSynchronizationSynchronizationIdGet**](SynchronizationApi.md#OrganizationsOrganizationIdAccountsAccountIdSynchronizationSynchronizationIdGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/synchronization/{synchronization_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdSynchronizationSynchronizationIdOptions**](SynchronizationApi.md#OrganizationsOrganizationIdAccountsAccountIdSynchronizationSynchronizationIdOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/synchronization/{synchronization_id} | 


# **OrganizationsOrganizationIdAccountsAccountIdSynchronizationOptions**
> OrganizationsOrganizationIdAccountsAccountIdSynchronizationOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdSynchronizationPost**
> Synchronization OrganizationsOrganizationIdAccountsAccountIdSynchronizationPost(ctx, accountId, organizationId, empty)


Synchronizes the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **empty** | [**Empty**](Empty.md)|  | 

### Return type

[**Synchronization**](Synchronization.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSynchronizationSynchronizationIdGet**
> Synchronization OrganizationsOrganizationIdAccountsAccountIdSynchronizationSynchronizationIdGet(ctx, synchronizationId, accountId, organizationId)


Returns the specific synchronization detail of an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **synchronizationId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Synchronization**](Synchronization.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdSynchronizationSynchronizationIdOptions**
> OrganizationsOrganizationIdAccountsAccountIdSynchronizationSynchronizationIdOptions(ctx, synchronizationId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **synchronizationId** | **string**|  | 
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

