# \MembersApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdMembersGet**](MembersApi.md#OrganizationsOrganizationIdAccountsAccountIdMembersGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/members | 
[**OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdDelete**](MembersApi.md#OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdDelete) | **Delete** /organizations/{organization_id}/accounts/{account_id}/members/{member_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdOptions**](MembersApi.md#OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/members/{member_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdPut**](MembersApi.md#OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdPut) | **Put** /organizations/{organization_id}/accounts/{account_id}/members/{member_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdMembersOptions**](MembersApi.md#OrganizationsOrganizationIdAccountsAccountIdMembersOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/members | 
[**OrganizationsOrganizationIdAccountsAccountIdMembersPost**](MembersApi.md#OrganizationsOrganizationIdAccountsAccountIdMembersPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/members | 
[**OrganizationsOrganizationIdMembersGet**](MembersApi.md#OrganizationsOrganizationIdMembersGet) | **Get** /organizations/{organization_id}/members | 
[**OrganizationsOrganizationIdMembersMemberIdDelete**](MembersApi.md#OrganizationsOrganizationIdMembersMemberIdDelete) | **Delete** /organizations/{organization_id}/members/{member_id} | 
[**OrganizationsOrganizationIdMembersMemberIdOptions**](MembersApi.md#OrganizationsOrganizationIdMembersMemberIdOptions) | **Options** /organizations/{organization_id}/members/{member_id} | 
[**OrganizationsOrganizationIdMembersMemberIdPut**](MembersApi.md#OrganizationsOrganizationIdMembersMemberIdPut) | **Put** /organizations/{organization_id}/members/{member_id} | 
[**OrganizationsOrganizationIdMembersOptions**](MembersApi.md#OrganizationsOrganizationIdMembersOptions) | **Options** /organizations/{organization_id}/members | 
[**OrganizationsOrganizationIdMembersPost**](MembersApi.md#OrganizationsOrganizationIdMembersPost) | **Post** /organizations/{organization_id}/members | 


# **OrganizationsOrganizationIdAccountsAccountIdMembersGet**
> Member OrganizationsOrganizationIdAccountsAccountIdMembersGet(ctx, accountId, organizationId)


Returns the members of the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Member**](Member.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdDelete**
> OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdDelete(ctx, memberId, accountId, organizationId)


Deletes the specific member of the account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **memberId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdOptions**
> OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdOptions(ctx, memberId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **memberId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdPut**
> Member OrganizationsOrganizationIdAccountsAccountIdMembersMemberIdPut(ctx, memberId, accountId, organizationId, accountMemberRequest)


Updates the specific member of the account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **memberId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **accountMemberRequest** | [**AccountMemberRequest**](AccountMemberRequest.md)|  | 

### Return type

[**Member**](Member.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdMembersOptions**
> OrganizationsOrganizationIdAccountsAccountIdMembersOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdMembersPost**
> Member OrganizationsOrganizationIdAccountsAccountIdMembersPost(ctx, accountId, organizationId, accountMemberRequest)


Creates the members of the given account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **accountMemberRequest** | [**AccountMemberRequest**](AccountMemberRequest.md)|  | 

### Return type

[**Member**](Member.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdMembersGet**
> []Member OrganizationsOrganizationIdMembersGet(ctx, organizationId)


Returns the members of associated organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **organizationId** | **string**| Organization ID | 

### Return type

[**[]Member**](Member.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdMembersMemberIdDelete**
> OrganizationsOrganizationIdMembersMemberIdDelete(ctx, memberId, organizationId)


Deletes the specific member associated to an organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **memberId** | **string**|  | 
  **organizationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdMembersMemberIdOptions**
> OrganizationsOrganizationIdMembersMemberIdOptions(ctx, memberId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **memberId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdMembersMemberIdPut**
> Member OrganizationsOrganizationIdMembersMemberIdPut(ctx, memberId, organizationId, member)


Updates the specific member associated to an organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **memberId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **member** | [**Member**](Member.md)|  | 

### Return type

[**Member**](Member.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdMembersOptions**
> OrganizationsOrganizationIdMembersOptions(ctx, organizationId)


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

# **OrganizationsOrganizationIdMembersPost**
> Member OrganizationsOrganizationIdMembersPost(ctx, organizationId, member)


Creates the members of associated organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **organizationId** | **string**| Organization ID | 
  **member** | [**Member**](Member.md)|  | 

### Return type

[**Member**](Member.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

