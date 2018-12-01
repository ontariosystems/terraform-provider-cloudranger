# \VolumesApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdVolumesGet**](VolumesApi.md#OrganizationsOrganizationIdAccountsAccountIdVolumesGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/volumes | 
[**OrganizationsOrganizationIdAccountsAccountIdVolumesIdsGet**](VolumesApi.md#OrganizationsOrganizationIdAccountsAccountIdVolumesIdsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/volumes/ids | 
[**OrganizationsOrganizationIdAccountsAccountIdVolumesIdsOptions**](VolumesApi.md#OrganizationsOrganizationIdAccountsAccountIdVolumesIdsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/volumes/ids | 
[**OrganizationsOrganizationIdAccountsAccountIdVolumesOptions**](VolumesApi.md#OrganizationsOrganizationIdAccountsAccountIdVolumesOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/volumes | 
[**OrganizationsOrganizationIdAccountsAccountIdVolumesTagsKeysGet**](VolumesApi.md#OrganizationsOrganizationIdAccountsAccountIdVolumesTagsKeysGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/volumes/tags/keys | 
[**OrganizationsOrganizationIdAccountsAccountIdVolumesTagsKeysOptions**](VolumesApi.md#OrganizationsOrganizationIdAccountsAccountIdVolumesTagsKeysOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/volumes/tags/keys | 
[**OrganizationsOrganizationIdAccountsAccountIdVolumesTagsValuesOptions**](VolumesApi.md#OrganizationsOrganizationIdAccountsAccountIdVolumesTagsValuesOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/volumes/tags/values | 
[**OrganizationsOrganizationIdAccountsAccountIdVolumesTagsValuesPost**](VolumesApi.md#OrganizationsOrganizationIdAccountsAccountIdVolumesTagsValuesPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/volumes/tags/values | 
[**OrganizationsOrganizationIdAccountsAccountIdVolumesVolumeIdGet**](VolumesApi.md#OrganizationsOrganizationIdAccountsAccountIdVolumesVolumeIdGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/volumes/{volume_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdVolumesVolumeIdOptions**](VolumesApi.md#OrganizationsOrganizationIdAccountsAccountIdVolumesVolumeIdOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/volumes/{volume_id} | 


# **OrganizationsOrganizationIdAccountsAccountIdVolumesGet**
> Volume OrganizationsOrganizationIdAccountsAccountIdVolumesGet(ctx, accountId, organizationId)


Returns the volumes of the servers associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Volume**](Volume.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdVolumesIdsGet**
> VolumeDetail OrganizationsOrganizationIdAccountsAccountIdVolumesIdsGet(ctx, accountId, organizationId)


Returns the ids of the volumes associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**VolumeDetail**](VolumeDetail.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdVolumesIdsOptions**
> OrganizationsOrganizationIdAccountsAccountIdVolumesIdsOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdVolumesOptions**
> OrganizationsOrganizationIdAccountsAccountIdVolumesOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdVolumesTagsKeysGet**
> TagsKey OrganizationsOrganizationIdAccountsAccountIdVolumesTagsKeysGet(ctx, accountId, organizationId)


Returns the tags keys of the volumes associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**TagsKey**](TagsKey.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdVolumesTagsKeysOptions**
> OrganizationsOrganizationIdAccountsAccountIdVolumesTagsKeysOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdVolumesTagsValuesOptions**
> OrganizationsOrganizationIdAccountsAccountIdVolumesTagsValuesOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdVolumesTagsValuesPost**
> TagsValue OrganizationsOrganizationIdAccountsAccountIdVolumesTagsValuesPost(ctx, accountId, organizationId, tagsValueRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **tagsValueRequest** | [**TagsValueRequest**](TagsValueRequest.md)|  | 

### Return type

[**TagsValue**](TagsValue.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdVolumesVolumeIdGet**
> VolumeDetail OrganizationsOrganizationIdAccountsAccountIdVolumesVolumeIdGet(ctx, volumeId, accountId, organizationId)


Returns the specific volume associated to an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **volumeId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**VolumeDetail**](VolumeDetail.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdVolumesVolumeIdOptions**
> OrganizationsOrganizationIdAccountsAccountIdVolumesVolumeIdOptions(ctx, volumeId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **volumeId** | **string**|  | 
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

