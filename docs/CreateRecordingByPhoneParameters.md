# CreateRecordingByPhoneParameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | (Required)  The name of the recording. | [optional] [default to null]
**CalleridId** | [**Uuid**](UUID.md) | The caller id that the create recording message should be sent from. | [optional] [default to null]
**Whitelabel** | **bool** | Add or remove the DialMyCalls intro message. | [optional] [default to null]
**Phone** | **string** | (Required)  The recipient&#39;s phone number who will record the message. | [optional] [default to null]
**Extension** | **string** | The recipient&#39;s phone extension up to 10 numeric digits. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


