# WctApiApplicationRefundDtoRefundOrderDetail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**OrderId** | **string** |  | [optional] [default to null]
**OrderCreationTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**PayNo** | **int64** |  | [optional] [default to null]
**OrderCode** | **string** |  | [optional] [default to null]
**TenantId** | **int64** |  | [optional] [default to null]
**UserId** | **int64** |  | [optional] [default to null]
**RefundAmount** | **float64** |  | [optional] [default to null]
**OriginalRefundAmount** | **float64** |  | [optional] [default to null]
**RefundType** | [***WctEntitiesRefundType**](WCT.Entities.RefundType.md) |  | [optional] [default to null]
**AuditType** | [***WctEntitiesAuditType**](WCT.Entities.AuditType.md) |  | [optional] [default to null]
**Status** | [***WctEntitiesAuditState**](WCT.Entities.AuditState.md) |  | [optional] [default to null]
**RefundStatus** | **int32** | 退款状态 0.未退款 1.退款中 2.已退款 3.退款失败 | [optional] [default to null]
**Reason** | **string** |  | [optional] [default to null]
**Remark** | **string** |  | [optional] [default to null]
**UploadImgInfo** | **[]string** |  | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**ConfirmTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**RefundTypeName** | **string** |  | [optional] [default to null]
**StatusName** | **string** |  | [optional] [default to null]
**TenantName** | **string** |  | [optional] [default to null]
**SupplierRemark** | **string** |  | [optional] [default to null]
**PlatformRemark** | **string** |  | [optional] [default to null]
**ReturnGoodsProvince** | **string** |  | [optional] [default to null]
**ReturnGoodsCity** | **string** |  | [optional] [default to null]
**ReturnGoodsAddresses** | **string** |  | [optional] [default to null]
**ReturnGoodsConsignee** | **string** |  | [optional] [default to null]
**ReturnGoodsMoblie** | **string** |  | [optional] [default to null]
**AutoRefundTime** | **int32** | （秒）退货退款：自动退款时间；    换货：自动平台介入时间 | [optional] [default to null]
**AutoCloseTime** | **int32** | （秒）自动关闭时间 | [optional] [default to null]
**IsDelivery** | **bool** |  | [optional] [default to null]
**RefundOrderDtl** | [**[]WctEntitiesRefundOrderDtl**](WCT.Entities.RefundOrderDtl.md) |  | [optional] [default to null]
**RefundOrderLog** | [**[]WctApiApplicationRefundDtoRefundOrderLogDto**](WCT.Api.Application.Refund.Dto.RefundOrderLogDto.md) |  | [optional] [default to null]
**RefundExchangeDtl** | [**[]WctEntitiesRefundOrderDtl**](WCT.Entities.RefundOrderDtl.md) | 换货明细 | [optional] [default to null]
**RefundOrderExchangeInfo** | [***WctApiApplicationRefundDtoRefundOrderExchangeDto**](WCT.Api.Application.Refund.Dto.RefundOrderExchangeDto.md) |  | [optional] [default to null]
**RefundProgress** | [**[]WctApiApplicationRefundDtoRefundProgressModel**](WCT.Api.Application.Refund.Dto.RefundProgressModel.md) | 售后进度 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

