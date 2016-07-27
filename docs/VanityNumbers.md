# \VanityNumbers

All URIs are relative to *https://api.dialmycalls.com/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVanityNumberById**](VanityNumbers.md#DeleteVanityNumberById) | **Delete** /vanitynumber/{VanityNumberId} | Delete Vanity Number
[**GetVanityNumberById**](VanityNumbers.md#GetVanityNumberById) | **Get** /vanitynumber/{VanityNumberId} | Get Vanity Number
[**GetVanityNumbers**](VanityNumbers.md#GetVanityNumbers) | **Get** /vanitynumbers | List Vanity Numbers
[**UpdateVanityNumberById**](VanityNumbers.md#UpdateVanityNumberById) | **Put** /vanitynumber/{VanityNumberId} | Update Vanity Number


# **DeleteVanityNumberById**
> Object DeleteVanityNumberById($vanityNumberId)

Delete Vanity Number

Delete a vanity number. <br><br> Returns the following if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X DELETE https://$API_KEY@api.dialmycalls.com/2.0/keyword/$VANITYNUMBER_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vanityNumberId** | **string**| VanityNumberId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVanityNumberById**
> Object GetVanityNumberById($vanityNumberId)

Get Vanity Number

Retrieve a vanity number. <br><br> Returns a vanitynumber object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/vanitynumber/$VANITYNUMBER_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vanityNumberId** | **string**| VanityNumberId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVanityNumbers**
> Object GetVanityNumbers($range_)

List Vanity Numbers

Retrieve a list of vanity numbers. <br><br> Returns a list of vanitynumber objects. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/vanitynumbers ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of vanitynumbers requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVanityNumberById**
> Object UpdateVanityNumberById($updateVanityNumberByIdParameters, $vanityNumberId)

Update Vanity Number

Update a vanity number. <br><br> Returns a vanitynumber object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X PUT -d "{\"ptt_number_id\":\"aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeee\",\"call_options\":{\"voicemail\":true}}" https://$API_KEY@api.dialmycalls.com/2.0/vanitynumber/$VANITYNUMBER_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateVanityNumberByIdParameters** | [**UpdateVanityNumberByIdParameters**](UpdateVanityNumberByIdParameters.md)| Request body | 
 **vanityNumberId** | **string**| VanityNumberId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

