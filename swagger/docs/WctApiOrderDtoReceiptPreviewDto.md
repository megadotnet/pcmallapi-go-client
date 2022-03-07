# WctApiOrderDtoReceiptPreviewDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNo** | **string** | 发票编号 | [optional] [default to null]
**OrderId** | **string** | 订单ID | [optional] [default to null]
**TenantName** | **string** | 店铺名称 | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | 创建时间 | [optional] [default to null]
**Products** | [**[]WctApiOrderDtoReceiptPreviewProductDto**](WCT.Api.Order.Dto.ReceiptPreviewProductDto.md) | 商品信息 | [optional] [default to null]
**ProductQuantityTotal** | **int32** | 数量合计 | [optional] [default to null]
**ProductAmountTotal** | **float64** | 商品金额合计 | [optional] [default to null]
**ShipmentFee** | **float64** | 运费 | [optional] [default to null]
**DiscountAmount** | **float64** | 优惠金额 | [optional] [default to null]
**TotalPayment** | **float64** | 支付合计 | [optional] [default to null]
**BuyerUser** | **string** | 买家姓名 | [optional] [default to null]
**BuyerMobile** | **string** | 买家手机号码 | [optional] [default to null]
**BuyerEmail** | **string** | 买家邮箱账号 | [optional] [default to null]
**LogisticsCompanyName** | **string** | 物流公司名称 | [optional] [default to null]
**ReceiptProvince** | **string** | 收货地址省 | [optional] [default to null]
**ReceiptCity** | **string** | 收货市 | [optional] [default to null]
**ReceiptAddress** | **string** | 收货地址 | [optional] [default to null]
**DispathTime** | [**time.Time**](time.Time.md) | 发货时间 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

