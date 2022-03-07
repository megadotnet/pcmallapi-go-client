# WctEntitiesOrderMst

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [optional] [default to null]
**OrderCode** | **string** |  | [optional] [default to null]
**TenantId** | **int32** |  | [optional] [default to null]
**UserId** | **int64** |  | [optional] [default to null]
**Consignee** | **string** |  | [optional] [default to null]
**Mobile** | **string** |  | [optional] [default to null]
**PayNo** | **int64** |  | [optional] [default to null]
**Status** | [***WctEntitiesOrderState**](WCT.Entities.OrderState.md) |  | [optional] [default to null]
**PayStatus** | [***WctEntitiesPayState**](WCT.Entities.PayState.md) |  | [optional] [default to null]
**RefundStatus** | [***WctEntitiesRefundState**](WCT.Entities.RefundState.md) |  | [optional] [default to null]
**OrderAmount** | **float64** |  | [optional] [default to null]
**PayTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**PaymentID** | [***WctEntitiesPayment**](WCT.Entities.Payment.md) |  | [optional] [default to null]
**PlatformOrderNo** | **string** |  | [optional] [default to null]
**DeliveryTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**CompletedTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**CreatorUserId** | **int64** |  | [optional] [default to null]
**DeletionTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**DeleterUserId** | **int64** |  | [optional] [default to null]
**IsDeleted** | **bool** |  | [optional] [default to null]
**SourceId** | **int32** |  | [optional] [default to null]
**ChannelOrderId** | **string** |  | [optional] [default to null]
**ChannelCode** | **string** |  | [optional] [default to null]
**OrderType** | [***WctEntitiesOrderType**](WCT.Entities.OrderType.md) |  | [optional] [default to null]
**GroupId** | **int64** |  | [optional] [default to null]
**TeamId** | **int64** |  | [optional] [default to null]
**Platform** | [***WctEntitiesPlatform**](WCT.Entities.Platform.md) |  | [optional] [default to null]
**IsInnerChannelOrder** | **bool** |  | [optional] [default to null]
**ShareUserId** | **int64** |  | [optional] [default to null]
**LeaveFormatData** | **string** |  | [optional] [default to null]
**OrderDtls** | [**[]WctEntitiesOrderDtl**](WCT.Entities.OrderDtl.md) |  | [optional] [default to null]
**OrderProducts** | [**[]WctEntitiesOrderProduct**](WCT.Entities.OrderProduct.md) |  | [optional] [default to null]
**OrderPayRecords** | [**[]WctEntitiesOrderPayRecord**](WCT.Entities.OrderPayRecord.md) |  | [optional] [default to null]
**OrderLogs** | [**[]WctEntitiesOrderLog**](WCT.Entities.OrderLog.md) |  | [optional] [default to null]
**OrderExpresses** | [**[]WctEntitiesOrderExpress**](WCT.Entities.OrderExpress.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

