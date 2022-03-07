# WctApiApplicationProductCommentDtoAppendCommentDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [optional] [default to null]
**AppendDiff** | **int32** | 追评距首评的时间差（天） | [optional] [default to null]
**ProductId** | **int64** | 商品Id | [optional] [default to null]
**FromCustomId** | **int64** | 买家Id，评价者Id | [optional] [default to null]
**FromNickname** | **string** | 买家名称，评价者昵称 | [optional] [default to null]
**FromHeadimg** | **string** | 评价者头像 | [optional] [default to null]
**Content** | **string** | 评论内容 | [optional] [default to null]
**LikeNum** | **int32** | 点赞数 | [optional] [default to null]
**ReplyCount** | **int32** | 回复数 | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | 评价时间 | [optional] [default to null]
**CommentImg** | [**[]WctAdminProductCommentCommentImg**](WCT.Admin.ProductComment.CommentImg.md) | 晒图 | [optional] [default to null]
**TenantReply** | [***WctAdminProductCommentReply**](WCT.Admin.ProductComment.Reply.md) |  | [optional] [default to null]
**Approve** | [***WctAdminProductCommentCommentApprove**](WCT.Admin.ProductComment.CommentApprove.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

