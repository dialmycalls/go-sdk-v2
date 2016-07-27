# \Recordings

All URIs are relative to *https://api.dialmycalls.com/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecording**](Recordings.md#CreateRecording) | **Post** /recording/tts | Create Recording (Text-to-Speech)
[**CreateRecordingByPhone**](Recordings.md#CreateRecordingByPhone) | **Post** /recording/phone | Create Recording (Phone)
[**CreateRecordingByUrl**](Recordings.md#CreateRecordingByUrl) | **Post** /recording/url | Create Recording (URL)
[**DeleteRecordingById**](Recordings.md#DeleteRecordingById) | **Delete** /recording/{RecordingId} | Delete Recording
[**GetRecordingById**](Recordings.md#GetRecordingById) | **Get** /recording/{RecordingId} | Get Recording
[**GetRecordings**](Recordings.md#GetRecordings) | **Get** /recordings | List Recordings
[**UpdateRecordingById**](Recordings.md#UpdateRecordingById) | **Put** /recording/{RecordingId} | Update Recording


# **CreateRecording**
> Object CreateRecording($createRecordingParameters)

Create Recording (Text-to-Speech)

Create a new recording using text-to-speech. <br><br> Returns a recording object on success, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X POST -d "{\"name\": \"Test Recording\", \"gender\": \"M\", \"language\": \"en\", \"text\": \"This is just a test.\"}" https://$API_KEY@api.dialmycalls.com/2.0/recording/tts ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRecordingParameters** | [**CreateRecordingParameters**](CreateRecordingParameters.md)| Request body | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRecordingByPhone**
> Object CreateRecordingByPhone($createRecordingByPhoneParameters)

Create Recording (Phone)

Create a new recording by phone. <br><br> Returns a recording object on success, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X POST -d "{\"name\": \"Test Recording\", \"phone\": \"5551234567\", \"callerid_id\": \"$CALLERID_ID\"}" https://$API_KEY@api.dialmycalls.com/2.0/recording/phone ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRecordingByPhoneParameters** | [**CreateRecordingByPhoneParameters**](CreateRecordingByPhoneParameters.md)| Request body | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRecordingByUrl**
> Object CreateRecordingByUrl($createRecordingByUrlParameters)

Create Recording (URL)

Create a new recording from a URL. <br><br> Returns a recording object on success, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X POST -d "{\"name\": \"Test Recording\", \"url\": \"https://ia700200.us.archive.org/1/items/testmp3testfile/mpthreetest.mp3\"}" https://$API_KEY@api.dialmycalls.com/2.0/recording/url ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRecordingByUrlParameters** | [**CreateRecordingByUrlParameters**](CreateRecordingByUrlParameters.md)| Request body | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRecordingById**
> Object DeleteRecordingById($recordingId)

Delete Recording

Delete a recording. <br><br> Returns the following if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X DELETE https://$API_KEY@api.dialmycalls.com/2.0/recording/$RECORDING_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordingId** | **string**| RecordingId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecordingById**
> Object GetRecordingById($recordingId)

Get Recording

Retrieve a recording. <br><br> Returns a recording object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/recording/$RECORDING_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordingId** | **string**| RecordingId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecordings**
> Object GetRecordings($range_)

List Recordings

Retrieve a list of recordings. <br><br> Returns a list of recording objects. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/recordings ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of recordings requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRecordingById**
> Object UpdateRecordingById($updateRecordingByIdParameters, $recordingId)

Update Recording

Update an existing recording. <br><br> Returns a recording object if a valid identifier was provided and input validation passed, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X PUT -d "{\"name\": \"Test Recording Updated\"}" https://$API_KEY@api.dialmycalls.com/2.0/recording/$RECORDING_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRecordingByIdParameters** | [**UpdateRecordingByIdParameters**](UpdateRecordingByIdParameters.md)| Request body | 
 **recordingId** | **string**| RecordingId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

