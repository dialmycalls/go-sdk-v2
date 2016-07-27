# \Calls

All URIs are relative to *https://api.dialmycalls.com/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelCallById**](Calls.md#CancelCallById) | **Delete** /service/call/{CallId} | Cancel Call
[**CreateCall**](Calls.md#CreateCall) | **Post** /service/call | Create Call
[**GetCallById**](Calls.md#GetCallById) | **Get** /service/call/{CallId} | Get Call
[**GetCallRecipientsByCallId**](Calls.md#GetCallRecipientsByCallId) | **Get** /service/call/{CallId}/recipients | Get Call Recipients
[**GetCalls**](Calls.md#GetCalls) | **Get** /service/calls | List Calls


# **CancelCallById**
> Object CancelCallById($callId)

Cancel Call

Cancel an outgoing call. <br><br> Returns the following if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X DELETE https://$API_KEY@api.dialmycalls.com/2.0/service/call/$CALL_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callId** | **string**| CallId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCall**
> Object CreateCall($createCallParameters)

Create Call

Create an outgoing call broadcast. <br><br> Returns a call service object on success, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X POST -d "{\"name\": \"Test\", \"callerid_id\": \"8bc6748e-d8a0-11e4-8b2d-00163e603cea\", \"recording_id\": \"079ff28a-1b75-11e5-88eb-00163e603cea\", \"send_immediately\": true, \"use_amd\": true, \"contacts\": [{\"phone\":\"1116551235\"},{\"phone\":\"1116551234\"}]}" https://$API_KEY@api.dialmycalls.com/2.0/service/call ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCallParameters** | [**CreateCallParameters**](CreateCallParameters.md)| Request body | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCallById**
> Object GetCallById($callId)

Get Call

Retrieve a call. <br><br> Returns a call service object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/service/call/$CALL_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callId** | **string**| CallId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCallRecipientsByCallId**
> Object GetCallRecipientsByCallId($callId, $range_)

Get Call Recipients

Retrieve a list of a call's recipients. <br><br> Returns a list of call recipient objects if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/service/call/$CALL_ID/recipients ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callId** | **string**| CallId | 
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of recipients requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCalls**
> Object GetCalls($range_)

List Calls

Retrieve a list of calls. <br><br> Returns a list of call service objects. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/service/calls ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of calls requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

