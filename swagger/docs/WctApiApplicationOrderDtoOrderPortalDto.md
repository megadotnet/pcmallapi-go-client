# WctApiApplicationOrderDtoOrderPortalDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**OrderCode** | **string** | 订单号 | [optional] [default to null]
**TenantId** | **int32** | 店铺id | [optional] [default to null]
**TenantName** | **string** | 店铺名称 | [optional] [default to null]
**TenantLogo** | **string** | 店铺logo | [optional] [default to null]
**OrderState** | [***WctEntitiesOrderState**](WCT.Entities.OrderState.md) |  | [optional] [default to null]
**StatusName** | **string** | 订单状态名称 | [optional] [default to null]
**ProductCount** | **int32** | 商品数量 | [optional] [default to null]
**OrderAmount** | **float64** | 订单金额 | [optional] [default to null]
**Express** | **int32** | 快递公司 | [optional] [default to null]
**ExpressCode** | **string** | 快递单号 | [optional] [default to null]
**ExpressName** | **string** | 快递公司名 | [optional] [default to null]
**ExpressInfo** | [**[]WctApiApplicationOrderDtoExpressInfo**](WCT.Api.Application.Order.Dto.ExpressInfo.md) | 快递信息 | [optional] [default to null]
**RefundState** | [***WctEntitiesRefundState**](WCT.Entities.RefundState.md) |  | [optional] [default to null]
**RefundStateName** | **string** | 售后状态名 | [optional] [default to null]
**CompletedTime** | [**time.Time**](time.Time.md) | 收货时间 | [optional] [default to null]
**OrderType** | [***WctEntitiesOrderType**](WCT.Entities.OrderType.md) |  | [optional] [default to null]
**ProductList** | [**[]WctApiApplicationOrderDtoOrderProductDto**](WCT.Api.Application.Order.Dto.OrderProductDto.md) | 订单商品列表 | [optional] [default to null]
**GroupId** | **int64** | 拼团Id | [optional] [default to null]
**TeamId** | **int64** | 拼团组Id | [optional] [default to null]
**Freight** | **float64** | 运费 | [optional] [default to null]
**AutoCloseTime** | **int32** | 订单自动关闭时间 | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | 创建时间 | [optional] [default to null]
**DeliveryReminder** | **int32** | 提醒卖家发货次数 | [optional] [default to null]
**CanModifyAddress** | **bool** | 是否可以修改地址 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

