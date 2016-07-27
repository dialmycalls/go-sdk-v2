# CreateTextParameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | (Required)  Name the broadcast. | [optional] [default to null]
**KeywordId** | [**Uuid**](UUID.md) | (Required)  The keyword id that should be associated with this broadcast. | [optional] [default to null]
**Messages** | **[]string** | (Required)  List of messages to send (up to 10). | [optional] [default to null]
**SendAt** | **string** | When the broadcast should be sent. | [optional] [default to null]
**SendImmediately** | **bool** | Should the broadcast go out immediately? | [optional] [default to null]
**SendEmail** | **bool** | Also send an email to the contacts? | [optional] [default to null]
**AccessaccountId** | [**Uuid**](UUID.md) | Schedule this broadcast as an access account. | [optional] [default to null]
**ShortcodeId** | [**Uuid**](UUID.md) | The shortcode id that the broadcast will be sent from. | [optional] [default to null]
**ConcatenateSms** | **bool** | Combine all SMS messages into 1 message on the end users device. | [optional] [default to null]
**Contacts** | [**[]ContactAttributes**](ContactAttributes.md) | (Required)  List of contact information that should be sent the broadcast. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


