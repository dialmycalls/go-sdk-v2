# \Groups

All URIs are relative to *https://api.dialmycalls.com/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](Groups.md#CreateGroup) | **Post** /group | Add Group
[**DeleteGroupById**](Groups.md#DeleteGroupById) | **Delete** /group/{GroupId} | Delete Group
[**GetGroupById**](Groups.md#GetGroupById) | **Get** /group/{GroupId} | Get Group
[**GetGroups**](Groups.md#GetGroups) | **Get** /groups | List Groups
[**UpdateGroupById**](Groups.md#UpdateGroupById) | **Put** /group/{GroupId} | Update Group


# **CreateGroup**
> Object CreateGroup($createGroupParameters)

Add Group

Add a contact group. <br><br> Returns a group object on success, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X POST -d "{\"name\": \"Test Group\"}" https://$API_KEY@api.dialmycalls.com/2.0/group ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupParameters** | [**CreateGroupParameters**](CreateGroupParameters.md)| Request body | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupById**
> Object DeleteGroupById($groupId)

Delete Group

Delete a contact group. <br><br> Returns the following if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X DELETE https://$API_KEY@api.dialmycalls.com/2.0/group/$GROUP_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| GroupId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupById**
> Object GetGroupById($groupId)

Get Group

Retrieve a contact group. <br><br> Returns a group object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/group/$GROUP_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string**| GroupId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroups**
> Object GetGroups($range_)

List Groups

Retrieve a list of contact groups. <br><br> Returns a list of group objects. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/groups ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of groups requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupById**
> Object UpdateGroupById($updateGroupByIdParameters, $groupId)

Update Group

Update an existing contact group. <br><br> Returns a group object if a valid identifier was provided and input validation passed, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X PUT -d "{\"name\": \"Test Group Updated\"}" https://$API_KEY@api.dialmycalls.com/2.0/group/$GROUP_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateGroupByIdParameters** | [**UpdateGroupByIdParameters**](UpdateGroupByIdParameters.md)| Request body | 
 **groupId** | **string**| GroupId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

