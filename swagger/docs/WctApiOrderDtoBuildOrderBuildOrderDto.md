# WctApiOrderDtoBuildOrderBuildOrderDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payment** | [***WctEntitiesPayment**](WCT.Entities.Payment.md) |  | [optional] [default to null]
**Remark** | **string** | 备注信息 | [optional] [default to null]
**Remarks** | [**[]WctApiOrderDtoBuildOrderBuildRemarks**](WCT.Api.Order.Dto.BuildOrder.BuildRemarks.md) | 按商家填写备注信息，Remarks优先级大于remark | [optional] [default to null]
**LeaveFormatData** | **string** | 自定义留言信息 | [optional] [default to null]
**AddressInfo** | [***WctApiApplicationOrderDtoAddressInfo**](WCT.Api.Application.Order.Dto.AddressInfo.md) |  | [optional] [default to null]
**ProductInfo** | [**[]WctApiOrderDtoBuildOrderBuildProductInfo**](WCT.Api.Order.Dto.BuildOrder.BuildProductInfo.md) | 商品明细 | [optional] [default to null]
**CouponInfos** | **[]string** | 优惠券信息 | [optional] [default to null]
**IntegralArriveCash** | **bool** | 是否启用积分抵现 | [optional] [default to null]
**Platform** | [***WctEntitiesPlatform**](WCT.Entities.Platform.md) |  | [optional] [default to null]
**PickUpStations** | [**[]WctApiOrderDtoBuildOrderBuildPickUpStation**](WCT.Api.Order.Dto.BuildOrder.BuildPickUpStation.md) | 自提点详情 | [optional] [default to null]
**LogisticsCompanyDto** | [**[]WctApiOrderDtoBuildOrderBuildLogisticsCompanyInput**](WCT.Api.Order.Dto.BuildOrder.BuildLogisticsCompanyInput.md) | 物流公司id | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

