# WctApiApplicationOrderDtoOrderProductDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | 主键 | [optional] [default to null]
**ProductId** | **int64** | 商品ID | [optional] [default to null]
**SkuId** | **int64** | SkuID | [optional] [default to null]
**ProductName** | **string** | 商品名称 | [optional] [default to null]
**ProductImg** | **string** | 商品图片 | [optional] [default to null]
**ProductType** | [***WctAdminProductsProductType**](WCT.Admin.Products.ProductType.md) |  | [optional] [default to null]
**ProductPrice** | **float64** | 商品价格 | [optional] [default to null]
**OriginalPrice** | **float64** | 商品原售价 | [optional] [default to null]
**CouponProductPrice** | **float64** | 优惠券金额 | [optional] [default to null]
**IntegralPrice** | **float64** | 积分抵扣的金额 | [optional] [default to null]
**ProductCount** | **int32** | 商品数量 | [optional] [default to null]
**Models** | **string** | 型号 | [optional] [default to null]
**RefundOrderId** | **string** | 售后订单Id | [optional] [default to null]
**RefundStatus** | [***WctEntitiesAuditState**](WCT.Entities.AuditState.md) |  | [optional] [default to null]
**RefundStatusName** | **string** | 退款状态值 | [optional] [default to null]
**CommentStatus** | [***WctAdminProductCommentCommentStatus**](WCT.Admin.ProductComment.CommentStatus.md) |  | [optional] [default to null]
**CommentId** | **int64** | 商品评价Id | [optional] [default to null]
**FullReductionId** | **int64** | 满减优惠id | [optional] [default to null]
**FullReductionPrice** | **float64** | 满减优惠金额 | [optional] [default to null]
**IsSeven** | **bool** | 是否支持7天退款 | [optional] [default to null]
**AfterSaleButton** | **bool** | 是否显示售后按钮 | [optional] [default to null]
**IsEnableIntegralExchange** | **bool** | 是否开启积分兑换 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

