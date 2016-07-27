# \Texts

All URIs are relative to *https://api.dialmycalls.com/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTextById**](Texts.md#CancelTextById) | **Delete** /service/text/{TextId} | Cancel Text
[**CreateText**](Texts.md#CreateText) | **Post** /service/text | Create Text
[**DeleteIncomingTextById**](Texts.md#DeleteIncomingTextById) | **Delete** /incoming/text/{TextId} | Delete Incoming Text
[**GetIncomingTextById**](Texts.md#GetIncomingTextById) | **Get** /incoming/text/{TextId} | Get Incoming Text
[**GetIncomingTexts**](Texts.md#GetIncomingTexts) | **Get** /incoming/texts | List Incoming Texts
[**GetShortCodes**](Texts.md#GetShortCodes) | **Get** /shortcodes | List Shortcodes
[**GetTextById**](Texts.md#GetTextById) | **Get** /service/text/{TextId} | Get Text
[**GetTextRecipientsByTextId**](Texts.md#GetTextRecipientsByTextId) | **Get** /service/text/{TextId}/recipients | Get Text Recipients
[**GetTexts**](Texts.md#GetTexts) | **Get** /service/texts | List Texts


# **CancelTextById**
> Object CancelTextById($textId)

Cancel Text

Cancel an outgoing text. <br><br> Returns the following if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X DELETE https://$API_KEY@api.dialmycalls.com/2.0/service/text/$TEXT_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textId** | **string**| TextId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateText**
> Object CreateText($createTextParameters)

Create Text

Create an outgoing text broadcast. <br><br> Returns a service object on success, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X POST -d "{\"keyword_id\": \"dfe49537-a0a8-4f4a-98a1-e03df388af11\", \"send_immediately\": true,\"messages\": [\"Testing testing\"], \"contacts\": [{\"phone\":\"1116551235\"},{\"phone\":\"1116551234\"}]}" https://$API_KEY@api.dialmycalls.com/2.0/service/text ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTextParameters** | [**CreateTextParameters**](CreateTextParameters.md)| Request body | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIncomingTextById**
> Object DeleteIncomingTextById($textId)

Delete Incoming Text

Delete a incoming text. <br><br> Returns the following if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X DELETE https://$API_KEY@api.dialmycalls.com/2.0/incoming/text/$TEXT_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textId** | **string**| TextId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIncomingTextById**
> Object GetIncomingTextById($textId)

Get Incoming Text

Retrieve a text. <br><br> Returns a Incoming Text object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/incoming/text/$TEXT_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textId** | **string**| TextId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIncomingTexts**
> Object GetIncomingTexts($range_)

List Incoming Texts

Retrieve a list of texts. <br><br> Returns a list of Incoming Text objects. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/incoming/texts ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of texts requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShortCodes**
> Object GetShortCodes($range_)

List Shortcodes

Retrieve a list of shortcodes. <br><br> Returns a list of shortcode objects. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/shortcodes ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of shortcodes requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTextById**
> Object GetTextById($textId)

Get Text

Retrieve a text. <br><br> Returns a service object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/service/text/$TEXT_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textId** | **string**| TextId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTextRecipientsByTextId**
> Object GetTextRecipientsByTextId($textId, $range_)

Get Text Recipients

Retrieve a list of a text's recipients. <br><br> Returns a list of text recipient objects if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/service/text/$TEXT_ID/recipients ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textId** | **string**| TextId | 
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of recipients requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTexts**
> Object GetTexts($range_)

List Texts

Retrieve a list of texts. <br><br> Returns a list of service objects. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/service/texts ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of texts requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

