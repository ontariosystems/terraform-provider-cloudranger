# \OrganizationApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeGet**](OrganizationApi.md#AuthorizeGet) | **Get** /authorize | 
[**AuthorizeOptions**](OrganizationApi.md#AuthorizeOptions) | **Options** /authorize | 
[**OrganizationsOrganizationIdGet**](OrganizationApi.md#OrganizationsOrganizationIdGet) | **Get** /organizations/{organization_id} | 
[**OrganizationsOrganizationIdOptions**](OrganizationApi.md#OrganizationsOrganizationIdOptions) | **Options** /organizations/{organization_id} | 
[**OrganizationsOrganizationIdPut**](OrganizationApi.md#OrganizationsOrganizationIdPut) | **Put** /organizations/{organization_id} | 


# **AuthorizeGet**
> Token AuthorizeGet(ctx, )


Use the API key to get your authenticaiton token

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Token**](Token.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthorizeOptions**
> Empty AuthorizeOptions(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Empty**](empty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdGet**
> Organization OrganizationsOrganizationIdGet(ctx, organizationId)


Returns the details of the associated organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **organizationId** | **string**| Organization ID | 

### Return type

[**Organization**](Organization.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdOptions**
> OrganizationsOrganizationIdOptions(ctx, organizationId)


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

# **OrganizationsOrganizationIdPut**
> Empty OrganizationsOrganizationIdPut(ctx, organizationId, organization)


Update organization details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **organizationId** | **string**| Organization ID | 
  **organization** | [**Organization**](Organization.md)|  | 

### Return type

[**Empty**](empty.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

