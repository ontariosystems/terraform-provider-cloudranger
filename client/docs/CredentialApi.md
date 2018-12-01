# \CredentialApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdCredentialGet**](CredentialApi.md#OrganizationsOrganizationIdAccountsAccountIdCredentialGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/credential | 
[**OrganizationsOrganizationIdAccountsAccountIdCredentialOptions**](CredentialApi.md#OrganizationsOrganizationIdAccountsAccountIdCredentialOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/credential | 
[**OrganizationsOrganizationIdAccountsAccountIdCredentialPost**](CredentialApi.md#OrganizationsOrganizationIdAccountsAccountIdCredentialPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/credential | 
[**OrganizationsOrganizationIdAccountsAccountIdCredentialPut**](CredentialApi.md#OrganizationsOrganizationIdAccountsAccountIdCredentialPut) | **Put** /organizations/{organization_id}/accounts/{account_id}/credential | 
[**OrganizationsOrganizationIdAccountsAccountIdCredentialRefreshOptions**](CredentialApi.md#OrganizationsOrganizationIdAccountsAccountIdCredentialRefreshOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/credential/refresh | 
[**OrganizationsOrganizationIdAccountsAccountIdCredentialRefreshPost**](CredentialApi.md#OrganizationsOrganizationIdAccountsAccountIdCredentialRefreshPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/credential/refresh | 
[**OrganizationsOrganizationIdAccountsAccountIdCredentialTestOptions**](CredentialApi.md#OrganizationsOrganizationIdAccountsAccountIdCredentialTestOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/credential/test | 
[**OrganizationsOrganizationIdAccountsAccountIdCredentialTestPost**](CredentialApi.md#OrganizationsOrganizationIdAccountsAccountIdCredentialTestPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/credential/test | 


# **OrganizationsOrganizationIdAccountsAccountIdCredentialGet**
> Credential OrganizationsOrganizationIdAccountsAccountIdCredentialGet(ctx, accountId, organizationId)


Returns the credentials of the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Credential**](Credential.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdCredentialOptions**
> OrganizationsOrganizationIdAccountsAccountIdCredentialOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdCredentialPost**
> Credential OrganizationsOrganizationIdAccountsAccountIdCredentialPost(ctx, accountId, organizationId, credentialsRequest)


Updates the credentials of the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **credentialsRequest** | [**CredentialsRequest**](CredentialsRequest.md)|  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdCredentialPut**
> Credential OrganizationsOrganizationIdAccountsAccountIdCredentialPut(ctx, accountId, organizationId, credentialsRequest)


Updates the credentials of the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **credentialsRequest** | [**CredentialsRequest**](CredentialsRequest.md)|  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdCredentialRefreshOptions**
> OrganizationsOrganizationIdAccountsAccountIdCredentialRefreshOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdCredentialRefreshPost**
> Empty OrganizationsOrganizationIdAccountsAccountIdCredentialRefreshPost(ctx, accountId, organizationId, empty)


Refreshes the permissions of a given account via cloud formation script

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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

# **OrganizationsOrganizationIdAccountsAccountIdCredentialTestOptions**
> OrganizationsOrganizationIdAccountsAccountIdCredentialTestOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdCredentialTestPost**
> Credential OrganizationsOrganizationIdAccountsAccountIdCredentialTestPost(ctx, accountId, organizationId, empty)


Checks whether Account Access Role ARN is configured

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **empty** | [**Empty**](Empty.md)|  | 

### Return type

[**Credential**](Credential.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

