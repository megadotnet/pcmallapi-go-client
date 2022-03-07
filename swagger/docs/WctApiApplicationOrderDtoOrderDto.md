# WctApiApplicationOrderDtoOrderDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [optional] [default to null]
**OrderCode** | **string** | 订单号 | [optional] [default to null]
**TenantID** | **int64** | 供应商ID | [optional] [default to null]
**TenantName** | **string** | 供应商名 | [optional] [default to null]
**UserID** | **int64** | 用户ID | [optional] [default to null]
**UserMobile** | **int64** | 用户手机号码 | [optional] [default to null]
**Mobile** | **string** | 手机号码 | [optional] [default to null]
**Status** | [***WctEntitiesOrderState**](WCT.Entities.OrderState.md) |  | [optional] [default to null]
**StatusName** | **string** | 订单状态名称 | [optional] [default to null]
**PayStatus** | [***WctEntitiesPayState**](WCT.Entities.PayState.md) |  | [optional] [default to null]
**PayStatusName** | **string** | 支付状态名 | [optional] [default to null]
**PayNo** | **int64** | 支付单号 | [optional] [default to null]
**RefundStatus** | [***WctEntitiesRefundState**](WCT.Entities.RefundState.md) |  | [optional] [default to null]
**RefundStatusName** | **string** | 售后状态名 | [optional] [default to null]
**OrderAmount** | **float64** | 订单金额 | [optional] [default to null]
**DeliveryTime** | [**time.Time**](time.Time.md) | 发货时间 | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | 创建时间 | [optional] [default to null]
**PayTime** | [**time.Time**](time.Time.md) | 支付时间 | [optional] [default to null]
**CreatorUserId** | **int64** | 创建人 | [optional] [default to null]
**CreatorUserName** | **string** | 创建人名 | [optional] [default to null]
**CouponAmount** | **float64** | 优惠券金额 | [optional] [default to null]
**CouponType** | **string** | 优惠券类型 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

