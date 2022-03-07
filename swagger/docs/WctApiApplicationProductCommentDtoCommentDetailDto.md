# WctApiApplicationProductCommentDtoCommentDetailDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [optional] [default to null]
**ProductId** | **int64** | 商品Id | [optional] [default to null]
**ProductName** | **string** | 商品名称 | [optional] [default to null]
**ProductImg** | **string** | 商品图片 | [optional] [default to null]
**SKUId** | **int64** | SKUId | [optional] [default to null]
**Models** | **string** | 型号，sku名称 | [optional] [default to null]
**OrderId** | **string** | 订单Id | [optional] [default to null]
**FromCustomId** | **int64** | 买家Id | [optional] [default to null]
**FromNickname** | **string** | 买家名称，评价者昵称 | [optional] [default to null]
**FromHeadimg** | **string** | 评价者头像 | [optional] [default to null]
**Content** | **string** | 评论内容 | [optional] [default to null]
**DescScore** | **int32** | 商品描述得分 1~5（对应1星~5星） | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | 评价时间 | [optional] [default to null]
**Type_** | [***WctAdminProductCommentCommentType**](WCT.Admin.ProductComment.CommentType.md) |  | [optional] [default to null]
**AppendDiff** | **int32** | 追评距首评的时间差（天） | [optional] [default to null]
**IsAnonymous** | **bool** | 是否匿名 0-否；1-是 | [optional] [default to null]
**LikeStatus** | **int32** | 当前登录用户对当前评价是否点赞。1-已赞；0-未赞 | [optional] [default to null]
**LikeId** | **int64** | 点赞记录Id | [optional] [default to null]
**LikeNum** | **int32** | 点赞数 | [optional] [default to null]
**Approve** | [***WctAdminProductCommentCommentApprove**](WCT.Admin.ProductComment.CommentApprove.md) |  | [optional] [default to null]
**CommentImg** | [**[]WctAdminProductCommentCommentImg**](WCT.Admin.ProductComment.CommentImg.md) | 晒图 | [optional] [default to null]
**TenantReply** | [***WctAdminProductCommentReply**](WCT.Admin.ProductComment.Reply.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

