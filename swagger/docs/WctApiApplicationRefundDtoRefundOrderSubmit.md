# WctApiApplicationRefundDtoRefundOrderSubmit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int64** | 订单Id | [optional] [default to null]
**RefundType** | [***WctEntitiesRefundType**](WCT.Entities.RefundType.md) |  | [optional] [default to null]
**OrderProductId** | [**[]WctApiApplicationRefundDtoRefundOrderProductDto**](WCT.Api.Application.Refund.Dto.RefundOrderProductDto.md) | 部分退款商品信息 | [optional] [default to null]
**Reason** | **string** | 退货原因（存储的为枚举，需要进行转换） | [optional] [default to null]
**Remark** | **string** | 备注信息 | [optional] [default to null]
**UploadImgInfo** | **string** | 上传图片信息 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

