# \AWSApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdAwsRegionsGet**](AWSApi.md#OrganizationsOrganizationIdAccountsAccountIdAwsRegionsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/aws/regions | 
[**OrganizationsOrganizationIdAccountsAccountIdAwsRegionsOptions**](AWSApi.md#OrganizationsOrganizationIdAccountsAccountIdAwsRegionsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/aws/regions | 
[**OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionAzsGet**](AWSApi.md#OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionAzsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/aws/regions/{region}/azs | 
[**OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionAzsOptions**](AWSApi.md#OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionAzsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/aws/regions/{region}/azs | 
[**OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSecuritygroupsGet**](AWSApi.md#OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSecuritygroupsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/aws/regions/{region}/securitygroups | 
[**OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSecuritygroupsOptions**](AWSApi.md#OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSecuritygroupsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/aws/regions/{region}/securitygroups | 
[**OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSshkeysGet**](AWSApi.md#OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSshkeysGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/aws/regions/{region}/sshkeys | 
[**OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSshkeysOptions**](AWSApi.md#OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSshkeysOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/aws/regions/{region}/sshkeys | 
[**OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSubnetsGet**](AWSApi.md#OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSubnetsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/aws/regions/{region}/subnets | 
[**OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSubnetsOptions**](AWSApi.md#OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSubnetsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/aws/regions/{region}/subnets | 


# **OrganizationsOrganizationIdAccountsAccountIdAwsRegionsGet**
> Region OrganizationsOrganizationIdAccountsAccountIdAwsRegionsGet(ctx, accountId, organizationId)


Get AWS region details of specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Region**](Region.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdAwsRegionsOptions**
> OrganizationsOrganizationIdAccountsAccountIdAwsRegionsOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionAzsGet**
> Empty OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionAzsGet(ctx, region, accountId, organizationId)


Get amazon availibility zone details of specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **region** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionAzsOptions**
> OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionAzsOptions(ctx, region, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **region** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSecuritygroupsGet**
> Empty OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSecuritygroupsGet(ctx, region, accountId, organizationId)


Get amazon secutiy groups details of specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **region** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSecuritygroupsOptions**
> OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSecuritygroupsOptions(ctx, region, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **region** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSshkeysGet**
> SshKey OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSshkeysGet(ctx, region, accountId, organizationId)


Get amazon ssh keys details of specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **region** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**SshKey**](SSHKey.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSshkeysOptions**
> OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSshkeysOptions(ctx, region, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **region** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSubnetsGet**
> Empty OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSubnetsGet(ctx, region, accountId, organizationId)


get subnets details of specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **region** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSubnetsOptions**
> OrganizationsOrganizationIdAccountsAccountIdAwsRegionsRegionSubnetsOptions(ctx, region, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **region** | **string**|  | 
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

