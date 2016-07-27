# Service

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**Uuid**](UUID.md) | Unique identifier for this service. | [optional] [default to null]
**AccessaccountId** | [**Uuid**](UUID.md) | Unique identifier for the access account which initiated this if available. | [optional] [default to null]
**Name** | **string** | The name of the service. | [optional] [default to null]
**PendingCancel** | **bool** | Whether the service has been flagged to be cancelled. | [optional] [default to null]
**CreditCost** | **float32** | The amount of credits required to schedule this service. | [optional] [default to null]
**TotalRecipients** | **float32** | The amount of recipients for this service. | [optional] [default to null]
**SendAt** | **string** | When the service will be sent. | [optional] [default to null]
**CreatedAt** | **string** | When the service was created. | [optional] [default to null]
**UpdatedAt** | **string** | When the service was last updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


