# \AccountsApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdDelete**](AccountsApi.md#OrganizationsOrganizationIdAccountsAccountIdDelete) | **Delete** /organizations/{organization_id}/accounts/{account_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdGet**](AccountsApi.md#OrganizationsOrganizationIdAccountsAccountIdGet) | **Get** /organizations/{organization_id}/accounts/{account_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdOptions**](AccountsApi.md#OrganizationsOrganizationIdAccountsAccountIdOptions) | **Options** /organizations/{organization_id}/accounts/{account_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdPut**](AccountsApi.md#OrganizationsOrganizationIdAccountsAccountIdPut) | **Put** /organizations/{organization_id}/accounts/{account_id} | 
[**OrganizationsOrganizationIdAccountsGet**](AccountsApi.md#OrganizationsOrganizationIdAccountsGet) | **Get** /organizations/{organization_id}/accounts | 
[**OrganizationsOrganizationIdAccountsOptions**](AccountsApi.md#OrganizationsOrganizationIdAccountsOptions) | **Options** /organizations/{organization_id}/accounts | 
[**OrganizationsOrganizationIdAccountsPost**](AccountsApi.md#OrganizationsOrganizationIdAccountsPost) | **Post** /organizations/{organization_id}/accounts | 


# **OrganizationsOrganizationIdAccountsAccountIdDelete**
> OrganizationsOrganizationIdAccountsAccountIdDelete(ctx, accountId, organizationId)


Delete account details of specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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

# **OrganizationsOrganizationIdAccountsAccountIdGet**
> Account OrganizationsOrganizationIdAccountsAccountIdGet(ctx, accountId, organizationId)


Get account details of specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Account**](Account.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdOptions**
> OrganizationsOrganizationIdAccountsAccountIdOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdPut**
> Empty OrganizationsOrganizationIdAccountsAccountIdPut(ctx, accountId, organizationId, account)


update account details of specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **account** | [**Account**](Account.md)|  | 

### Return type

[**Empty**](empty.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsGet**
> []Account OrganizationsOrganizationIdAccountsGet(ctx, organizationId)


Returns the account details of the associated organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **organizationId** | **string**| Organization ID | 

### Return type

[**[]Account**](Account.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsOptions**
> OrganizationsOrganizationIdAccountsOptions(ctx, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsPost**
> Account OrganizationsOrganizationIdAccountsPost(ctx, organizationId, organizationAccountCreate)


Creates the account ib the organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **organizationId** | **string**| Organization ID | 
  **organizationAccountCreate** | [**OrganizationAccountCreate**](OrganizationAccountCreate.md)|  | 

### Return type

[**Account**](Account.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

