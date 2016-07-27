# \Accounts

All URIs are relative to *https://api.dialmycalls.com/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessAccount**](Accounts.md#CreateAccessAccount) | **Post** /accessaccount | Add Access Account
[**DeleteAccessAccountById**](Accounts.md#DeleteAccessAccountById) | **Delete** /accessaccount/{AccessAccountId} | Delete Access Account
[**GetAccessAccountById**](Accounts.md#GetAccessAccountById) | **Get** /accessaccount/{AccessAccountId} | Get Access Account
[**GetAccessAccounts**](Accounts.md#GetAccessAccounts) | **Get** /accessaccounts | List Access Accounts
[**GetAccount**](Accounts.md#GetAccount) | **Get** /account | Get Account
[**UpdateAccessAccountById**](Accounts.md#UpdateAccessAccountById) | **Put** /accessaccount/{AccessAccountId} | Update Access Account


# **CreateAccessAccount**
> Object CreateAccessAccount($createAccessAccountParameters)

Add Access Account

Add a access account. <br><br> Returns a access account object on success, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X POST -d "{\"email\": \"test@test.com\", \"password\": \"1234A5678\", \"name\": \"John Doe\"}" https://$API_KEY@api.dialmycalls.com/2.0/accessaccount ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccessAccountParameters** | [**CreateAccessAccountParameters**](CreateAccessAccountParameters.md)| Request body | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccessAccountById**
> Object DeleteAccessAccountById($accessAccountId)

Delete Access Account

Delete a access account. <br><br> Returns the following if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X DELETE https://$API_KEY@api.dialmycalls.com/2.0/accessaccount/$ACCESSACCOUNT_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessAccountId** | **string**| AccessAccountId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessAccountById**
> Object GetAccessAccountById($accessAccountId)

Get Access Account

Retrieve an access account. <br><br> Returns a access account object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/accessaccount/$ACCESSACCOUNT_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessAccountId** | **string**| AccessAccountId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessAccounts**
> Object GetAccessAccounts($range_)

List Access Accounts

Retrieve a list of access accounts. <br><br> Returns a list of access account objects. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/accessaccounts ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string**| Range (ie \&quot;records&#x3D;201-300\&quot;) of accessaccounts requested | [optional] 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccount**
> Object GetAccount()

Get Account

Retrieve account details. <br><br> Returns a account object if a valid identifier was provided, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X GET https://$API_KEY@api.dialmycalls.com/2.0/account ```


### Parameters
This endpoint does not need any parameter.

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccessAccountById**
> Object UpdateAccessAccountById($updateAccessAccountByIdParameters, $accessAccountId)

Update Access Account

Update an existing access account. <br><br> Returns a access account object if a valid identifier was provided and input validation passed, and returns an error otherwise. <br><br> ``` curl -i -H "Content-Type: application/json" -X PUT -d "{\"name\": \"New Name\"}" https://$API_KEY@api.dialmycalls.com/2.0/accessaccount/$ACCESSACCOUNT_ID ```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAccessAccountByIdParameters** | [**UpdateAccessAccountByIdParameters**](UpdateAccessAccountByIdParameters.md)| Request body | 
 **accessAccountId** | **string**| AccessAccountId | 

### Return type

[**Object**](object.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

