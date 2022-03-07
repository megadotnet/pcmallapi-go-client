# WctApiApplicationProductCommentDtoCreateCommentDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [optional] [default to null]
**OrderId** | **string** | 订单Id | [optional] [default to null]
**ProductId** | **int64** | 商品Id | [optional] [default to null]
**SkuId** | **int64** | SKUId | [optional] [default to null]
**Content** | **string** | 评价内容 | [default to null]
**DescScore** | **int32** | 商品描述得分 1~5（对应1星~5星） | [optional] [default to null]
**ServiceScore** | **int32** | 客服服务得分 1~5（对应1星~5星） | [optional] [default to null]
**ExpressScore** | **int32** | 发货速度得分 1~5（对应1星~5星） | [optional] [default to null]
**IsAnonymous** | **bool** | 是否匿名 0-否；1-是 | [optional] [default to null]
**Type_** | [***WctAdminProductCommentCommentType**](WCT.Admin.ProductComment.CommentType.md) |  | [optional] [default to null]
**CommentImgs** | [**[]WctAdminProductCommentCommentImg**](WCT.Admin.ProductComment.CommentImg.md) | 晒图 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

