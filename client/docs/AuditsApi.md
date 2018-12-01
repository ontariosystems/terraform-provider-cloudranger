# \AuditsApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdAuditsCountGet**](AuditsApi.md#OrganizationsOrganizationIdAccountsAccountIdAuditsCountGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/audits/count | 
[**OrganizationsOrganizationIdAccountsAccountIdAuditsCountOptions**](AuditsApi.md#OrganizationsOrganizationIdAccountsAccountIdAuditsCountOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/audits/count | 
[**OrganizationsOrganizationIdAccountsAccountIdAuditsGet**](AuditsApi.md#OrganizationsOrganizationIdAccountsAccountIdAuditsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/audits | 
[**OrganizationsOrganizationIdAccountsAccountIdAuditsOptions**](AuditsApi.md#OrganizationsOrganizationIdAccountsAccountIdAuditsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/audits | 
[**OrganizationsOrganizationIdAuditsCountGet**](AuditsApi.md#OrganizationsOrganizationIdAuditsCountGet) | **Get** /organizations/{organization_id}/audits/count | 
[**OrganizationsOrganizationIdAuditsCountOptions**](AuditsApi.md#OrganizationsOrganizationIdAuditsCountOptions) | **Options** /organizations/{organization_id}/audits/count | 
[**OrganizationsOrganizationIdAuditsGet**](AuditsApi.md#OrganizationsOrganizationIdAuditsGet) | **Get** /organizations/{organization_id}/audits | 
[**OrganizationsOrganizationIdAuditsOptions**](AuditsApi.md#OrganizationsOrganizationIdAuditsOptions) | **Options** /organizations/{organization_id}/audits | 


# **OrganizationsOrganizationIdAccountsAccountIdAuditsCountGet**
> Count OrganizationsOrganizationIdAccountsAccountIdAuditsCountGet(ctx, accountId, organizationId)


Get audit count

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

# **OrganizationsOrganizationIdAccountsAccountIdAuditsCountOptions**
> OrganizationsOrganizationIdAccountsAccountIdAuditsCountOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdAuditsGet**
> Audit OrganizationsOrganizationIdAccountsAccountIdAuditsGet(ctx, accountId, organizationId)


Get audit details of specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Audit**](Audit.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdAuditsOptions**
> OrganizationsOrganizationIdAccountsAccountIdAuditsOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAuditsCountGet**
> Count OrganizationsOrganizationIdAuditsCountGet(ctx, organizationId)


Returns the count of audits related to accounts of associated organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **organizationId** | **string**| Organization ID | 

### Return type

[**Count**](Count.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAuditsCountOptions**
> OrganizationsOrganizationIdAuditsCountOptions(ctx, organizationId)


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

# **OrganizationsOrganizationIdAuditsGet**
> Audit OrganizationsOrganizationIdAuditsGet(ctx, organizationId)


Returns the audits of accounts associated to organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **organizationId** | **string**| Organization ID | 

### Return type

[**Audit**](Audit.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAuditsOptions**
> OrganizationsOrganizationIdAuditsOptions(ctx, organizationId)


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

