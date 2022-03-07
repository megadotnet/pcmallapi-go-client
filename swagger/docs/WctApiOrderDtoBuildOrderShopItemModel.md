# WctApiOrderDtoBuildOrderShopItemModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **int32** | 商家id | [optional] [default to null]
**TenantName** | **string** | 商家名称 | [optional] [default to null]
**Postage** | **float64** | 邮费 | [optional] [default to null]
**OpenPickUpStation** | **bool** | 是否开启自提点 | [optional] [default to null]
**OrderItems** | [**[]WctApiOrderDtoBuildOrderOrderItemModel**](WCT.Api.Order.Dto.BuildOrder.OrderItemModel.md) |  | [optional] [default to null]
**DiscountAmount** | **float64** | 商品折扣优惠金额 | [optional] [default to null]
**CouponAmount** | **float64** | 优惠券优惠金额 | [optional] [default to null]
**FullReductionOfferAmount** | **float64** | 满减活动优惠金额 | [optional] [default to null]
**IntegralAmount** | **float64** | 积分抵扣优惠金额 | [optional] [default to null]
**ProductQuantity** | **int32** | 商品总数 | [optional] [default to null]
**ProductAmount** | **float64** | 商品金额总数 | [optional] [default to null]
**AfterDiscountAmount** | **float64** | 商品优惠后金额总数 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

