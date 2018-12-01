# \RestoresApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdRestoresGet**](RestoresApi.md#OrganizationsOrganizationIdAccountsAccountIdRestoresGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/restores | 
[**OrganizationsOrganizationIdAccountsAccountIdRestoresOptions**](RestoresApi.md#OrganizationsOrganizationIdAccountsAccountIdRestoresOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/restores | 
[**OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdDelete**](RestoresApi.md#OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdDelete) | **Delete** /organizations/{organization_id}/accounts/{account_id}/restores/{restore_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdGet**](RestoresApi.md#OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/restores/{restore_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdOptions**](RestoresApi.md#OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/restores/{restore_id} | 


# **OrganizationsOrganizationIdAccountsAccountIdRestoresGet**
> Restore OrganizationsOrganizationIdAccountsAccountIdRestoresGet(ctx, accountId, organizationId)


Returns all restores associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Restore**](Restore.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdRestoresOptions**
> OrganizationsOrganizationIdAccountsAccountIdRestoresOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdDelete**
> OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdDelete(ctx, restoreId, accountId, organizationId)


Deletes the specific restore associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **restoreId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdGet**
> Restore OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdGet(ctx, restoreId, accountId, organizationId)


Returns the specific restore associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **restoreId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Restore**](Restore.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdOptions**
> OrganizationsOrganizationIdAccountsAccountIdRestoresRestoreIdOptions(ctx, restoreId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **restoreId** | **string**|  | 
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

