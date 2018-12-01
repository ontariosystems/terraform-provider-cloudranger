# \BackupsApi

All URIs are relative to *https://api.cloudranger.com/201805*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAmirestoreOptions**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAmirestoreOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id}/amirestore | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAmirestorePost**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAmirestorePost) | **Post** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id}/amirestore | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAttachrestoreOptions**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAttachrestoreOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id}/attachrestore | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAttachrestorePost**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAttachrestorePost) | **Post** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id}/attachrestore | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdDelete**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdDelete) | **Delete** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdFilerestoreOptions**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdFilerestoreOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id}/filerestore | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdFilerestorePost**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdFilerestorePost) | **Post** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id}/filerestore | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdGet**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdOptions**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdPut**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdPut) | **Put** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id} | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdVolumerestoreOptions**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdVolumerestoreOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id}/volumerestore | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdVolumerestorePost**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdVolumerestorePost) | **Post** /organizations/{organization_id}/accounts/{account_id}/backups/{backup_id}/volumerestore | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsCountGet**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsCountGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/backups/count | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsCountOptions**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsCountOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/backups/count | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsCountTagOptions**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsCountTagOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/backups/count/tag | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsCountTagPost**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsCountTagPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/backups/count/tag | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsGet**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/backups | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsOptions**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/backups | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsTagsKeysGet**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsTagsKeysGet) | **Get** /organizations/{organization_id}/accounts/{account_id}/backups/tags/keys | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsTagsKeysOptions**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsTagsKeysOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/backups/tags/keys | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsTagsValuesOptions**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsTagsValuesOptions) | **Options** /organizations/{organization_id}/accounts/{account_id}/backups/tags/values | 
[**OrganizationsOrganizationIdAccountsAccountIdBackupsTagsValuesPost**](BackupsApi.md#OrganizationsOrganizationIdAccountsAccountIdBackupsTagsValuesPost) | **Post** /organizations/{organization_id}/accounts/{account_id}/backups/tags/values | 


# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAmirestoreOptions**
> OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAmirestoreOptions(ctx, backupId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAmirestorePost**
> AmiRestore OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAmirestorePost(ctx, backupId, accountId, organizationId, amiRestoreRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **amiRestoreRequest** | [**AmiRestoreRequest**](AmiRestoreRequest.md)|  | 

### Return type

[**AmiRestore**](AmiRestore.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAttachrestoreOptions**
> OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAttachrestoreOptions(ctx, backupId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAttachrestorePost**
> Empty OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdAttachrestorePost(ctx, backupId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdDelete**
> Empty OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdDelete(ctx, backupId, accountId, organizationId)


Deletes the backup with the given id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**|  | 

### Return type

[**Empty**](empty.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdFilerestoreOptions**
> OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdFilerestoreOptions(ctx, backupId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdFilerestorePost**
> Restore OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdFilerestorePost(ctx, backupId, accountId, organizationId, fileRestore)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **fileRestore** | [**FileRestore**](FileRestore.md)|  | 

### Return type

[**Restore**](Restore.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdGet**
> BackupDetailOne OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdGet(ctx, backupId, accountId, organizationId)


Returns the backup assciated with the given backup id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**BackupDetailOne**](BackupDetailOne.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdOptions**
> OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdOptions(ctx, backupId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdPut**
> OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdPut(ctx, backupId, accountId, organizationId, backupDetailOne)


Updates the backup attributes of the given backup id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **backupDetailOne** | [**BackupDetailOne**](BackupDetailOne.md)|  | 

### Return type

 (empty response body)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdVolumerestoreOptions**
> OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdVolumerestoreOptions(ctx, backupId, accountId, organizationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdVolumerestorePost**
> Restore OrganizationsOrganizationIdAccountsAccountIdBackupsBackupIdVolumerestorePost(ctx, backupId, accountId, organizationId, volumeRestore)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **backupId** | **string**|  | 
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **volumeRestore** | [**VolumeRestore**](VolumeRestore.md)|  | 

### Return type

[**Restore**](Restore.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdBackupsCountGet**
> BackupsCount OrganizationsOrganizationIdAccountsAccountIdBackupsCountGet(ctx, accountId, organizationId)


get backups count details of specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**BackupsCount**](BackupsCount.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdBackupsCountOptions**
> OrganizationsOrganizationIdAccountsAccountIdBackupsCountOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsCountTagOptions**
> OrganizationsOrganizationIdAccountsAccountIdBackupsCountTagOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsCountTagPost**
> CountTag OrganizationsOrganizationIdAccountsAccountIdBackupsCountTagPost(ctx, accountId, organizationId, backupsCountsTagRequest)


Get total count and already gone count for backups tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **backupsCountsTagRequest** | [**BackupsCountsTagRequest**](BackupsCountsTagRequest.md)|  | 

### Return type

[**CountTag**](CountTag.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdBackupsGet**
> Backup OrganizationsOrganizationIdAccountsAccountIdBackupsGet(ctx, accountId, organizationId)


Get backups details of specific account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 

### Return type

[**Backup**](Backup.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrganizationsOrganizationIdAccountsAccountIdBackupsOptions**
> OrganizationsOrganizationIdAccountsAccountIdBackupsOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsTagsKeysGet**
> TagsKey OrganizationsOrganizationIdAccountsAccountIdBackupsTagsKeysGet(ctx, accountId, organizationId)


Returns the tag keys for backups created by the given account of the specific organization

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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsTagsKeysOptions**
> OrganizationsOrganizationIdAccountsAccountIdBackupsTagsKeysOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsTagsValuesOptions**
> OrganizationsOrganizationIdAccountsAccountIdBackupsTagsValuesOptions(ctx, accountId, organizationId)


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

# **OrganizationsOrganizationIdAccountsAccountIdBackupsTagsValuesPost**
> TagsValue OrganizationsOrganizationIdAccountsAccountIdBackupsTagsValuesPost(ctx, accountId, organizationId, backupsTagsValuesRequest)


Get tags keys value details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**|  | 
  **organizationId** | **string**| Organization ID | 
  **backupsTagsValuesRequest** | [**BackupsTagsValuesRequest**](BackupsTagsValuesRequest.md)|  | 

### Return type

[**TagsValue**](TagsValue.md)

### Authorization

[API_Authorizer](../README.md#API_Authorizer), [api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

