# WctApiOrderDtoBuildOrderConfirmOrderOutput

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopItems** | [**[]WctApiOrderDtoBuildOrderShopItemModel**](WCT.Api.Order.Dto.BuildOrder.ShopItemModel.md) |  | [optional] [default to null]
**TotalPostage** | **float64** | 邮费 | [optional] [default to null]
**DiscountAmount** | **float64** | 商品折扣优惠金额 | [optional] [default to null]
**CouponAmount** | **float64** | 优惠券优惠金额 | [optional] [default to null]
**FullReductionOfferAmount** | **float64** | 满减活动优惠金额 | [optional] [default to null]
**IntegralAmount** | **float64** | 积分抵扣优惠金额 | [optional] [default to null]
**TotalProductQuantity** | **int32** | 商品总数 | [optional] [default to null]
**TotalProductAmount** | **float64** | 商品金额总数 | [optional] [default to null]
**AfterDiscountAmount** | **float64** | 商品优惠后金额总数 | [optional] [default to null]
**OrderAmount** | **float64** | 订单金额 | [optional] [default to null]
**CanInvoice** | **bool** | 是否支持开票 | [optional] [default to null]
**UseIntegralInfo** | [***WctAdminIntegralServiceDtoIntegralArriveCashOutput**](WCT.Admin.Integral.Service.Dto.IntegralArriveCashOutput.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

