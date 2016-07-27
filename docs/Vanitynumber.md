# Vanitynumber

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**Uuid**](UUID.md) | Unique identifier for this group. | [optional] [default to null]
**Phone** | **string** | The phone number. | [optional] [default to null]
**Status** | **string** | The status of the vanity number. Options: active, onhold, billingdecline, pendingdelete | [optional] [default to null]
**MinutesUsed** | **float32** | The amount of minutes used since last billing. | [optional] [default to null]
**MinutesAllowed** | **float32** | The amount of minutes included with the vanity number before billed additionally. | [optional] [default to null]
**VoicemailsNew** | **int32** | The amount of voicemails that have not been downloaded. | [optional] [default to null]
**VoicemailsOld** | **int32** | The amount of voicemails that have been downloaded. | [optional] [default to null]
**CreatedAt** | **string** | When the keyword was created. | [optional] [default to null]
**UpdatedAt** | **string** | When the keyword was last updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


