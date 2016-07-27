# \CallerIds

All URIs are relative to *https://api.dialmycalls.com/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCallerId**](CallerIds.md#CreateCallerId) | **Post** /callerid | Add Caller ID
[**CreateUnverifiedCallerId**](CallerIds.md#CreateUnverifiedCallerId) | **Post** /verify/callerid | Add Caller ID (Unverified)
[**DeleteCallerIdById**](CallerIds.md#DeleteCallerIdById) | **Delete** /callerid/{CalleridId} | Delete Caller ID
[**GetCallerIdById**](CallerIds.md#GetCallerIdById) | **Get** /callerid/{CalleridId} | Get Caller ID
[**GetCallerIds**](CallerIds.md#GetCallerIds) | **Get** /callerids | List Caller IDs
[**UpdateCallerIdById**](CallerIds.md#UpdateCallerIdById) | **Put** /callerid/{CalleridId} | Update Caller ID
[**VerifyCallerIdById**](CallerIds.md#VerifyCallerIdById) | **Put** /verify/callerid/{CalleridId} | Verify Caller ID


# **CreateCallerId**
> Object CreateCallerId($createCallerIdParameters)

Add Caller ID

Add a caller ID. <br><br> Returns a caller ID object on success, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X POST -d "{\"phone\": \"5555555555\", \"name\": \"New Number\"}" https://$API_KEY@api.dialmycalls.com/2.0/callerid ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCallerIdParameters** | [**CreateCallerIdParameters**](CreateCallerIdParameters.md)| Request body | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUnverifiedCallerId**
> Object CreateUnverifiedCallerId($createUnverifiedCallerIdParameters)

Add Caller ID (Unverified)

Initiate adding a new caller ID when you need to go through the verification process. You will receive a phone call at the phone number provided and will need to make a subsequent API call with the PIN code that you are provided. <br><br> Returns a caller ID object on success, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X POST -d "{\"phone\": \"5555555555\", \"name\": \"New Number\"}" https://$API_KEY@api.dialmycalls.com/2.0/verify/callerid ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUnverifiedCallerIdParameters** | [**CreateUnverifiedCallerIdParameters**](CreateUnverifiedCallerIdParameters.md)| Request body | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCallerIdById**
> Object DeleteCallerIdById($calleridId)

Delete Caller ID

Delete a caller ID. <br><br> Returns the following if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X DELETE https://$API_KEY@api.dialmycalls.com/2.0/callerid/$CALLERID_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calleridId** | **string**| CalleridId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCallerIdById**
> Object GetCallerIdById($calleridId)

Get Caller ID

Retrieve a caller ID. <br><br> Returns a caller ID object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/callerid/$CALLERID_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calleridId** | **string**| CalleridId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCallerIds**
> Object GetCallerIds($range_)

List Caller IDs

Retrieve a list of caller IDs. <br><br> Returns a list of caller ID objects. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/callerids ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of callerids requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCallerIdById**
> Object UpdateCallerIdById($updateCallerIdByIdParameters, $calleridId)

Update Caller ID

Update an existing caller ID. <br><br> Returns a caller ID object if a valid identifier was provided and input validation passed, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X PUT -d "{\"name\": \"New Number Updated\"}" https://$API_KEY@api.dialmycalls.com/2.0/callerid/$CALLERID_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCallerIdByIdParameters** | [**UpdateCallerIdByIdParameters**](UpdateCallerIdByIdParameters.md)| Request body | 
 **calleridId** | **string**| CalleridId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyCallerIdById**
> Object VerifyCallerIdById($verifyCallerIdByIdParameters, $calleridId)

Verify Caller ID

Verify a new caller ID. <br><br> Returns a caller ID object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X POST -d "{\"phone\": \"5555555555\", \"name\": \"New Number\"}" https://$API_KEY@api.dialmycalls.com/2.0/verify/callerid/$CALLERID_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyCallerIdByIdParameters** | [**VerifyCallerIdByIdParameters**](VerifyCallerIdByIdParameters.md)| Request body | 
 **calleridId** | **string**| CalleridId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

