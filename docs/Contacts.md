# \Contacts

All URIs are relative to *https://api.dialmycalls.com/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContact**](Contacts.md#CreateContact) | **Post** /contact | Add Contact
[**DeleteContactById**](Contacts.md#DeleteContactById) | **Delete** /contact/{ContactId} | Delete Contact
[**GetContactById**](Contacts.md#GetContactById) | **Get** /contact/{ContactId} | Get Contact
[**GetContacts**](Contacts.md#GetContacts) | **Get** /contacts | List Contacts
[**GetContactsByGroupId**](Contacts.md#GetContactsByGroupId) | **Get** /contacts/{GroupId} | List Contacts in Group
[**UpdateContactById**](Contacts.md#UpdateContactById) | **Put** /contact/{ContactId} | Update Contact


# **CreateContact**
> Object CreateContact($createContactParameters)

Add Contact

Add a contact to your contact list. <br><br> Returns a contact object on success, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X POST -d "{\"phone\": \"5555555555\"}" https://$API_KEY@api.dialmycalls.com/2.0/contact ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContactParameters** | [**CreateContactParameters**](CreateContactParameters.md)| Request body | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContactById**
> Object DeleteContactById($contactId)

Delete Contact

Delete a contact from your contact list. <br><br> Returns the following if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X DELETE https://$API_KEY@api.dialmycalls.com/2.0/contact/$CONTACT_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **string**| ContactId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactById**
> Object GetContactById($contactId)

Get Contact

Retrieve a contact to your contact list. <br><br> Returns a contact object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/contact/$CONTACT_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **string**| ContactId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContacts**
> Object GetContacts($range_)

List Contacts

Retrieve a list of contacts. <br><br> Returns a list of contact objects. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/contacts ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of contacts requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactsByGroupId**
> Object GetContactsByGroupId($groupId)

List Contacts in Group

Retrieve a list of contacts in a contact group. <br><br> Returns a list of contact objects. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/contacts/$GROUP_ID ```


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

# **UpdateContactById**
> Object UpdateContactById($updateContactByIdParameters, $contactId)

Update Contact

Update an existing contact in your contact list. <br><br> Returns a contact object if a valid identifier was provided and input validation passed, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X PUT -d "{\"phone\": \"5555555555\"}" https://$API_KEY@api.dialmycalls.com/2.0/contact/$CONTACT_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateContactByIdParameters** | [**UpdateContactByIdParameters**](UpdateContactByIdParameters.md)| Request body | 
 **contactId** | **string**| ContactId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

