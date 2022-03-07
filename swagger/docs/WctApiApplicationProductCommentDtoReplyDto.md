# WctApiApplicationProductCommentDtoReplyDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [optional] [default to null]
**CommentId** | **int64** | 当前回复对应的评论Id | [optional] [default to null]
**ReplyType** | [***WctAdminProductCommentReplyType**](WCT.Admin.ProductComment.ReplyType.md) |  | [optional] [default to null]
**FromCustomId** | **int64** | 回复人Id | [optional] [default to null]
**FromNickname** | **string** | 回复人昵称 | [optional] [default to null]
**FromHeadimg** | **string** | 回复人头像 | [optional] [default to null]
**ToCustomId** | **int64** | 回复对象Id | [optional] [default to null]
**ToNickname** | **string** | 回复对象昵称 | [optional] [default to null]
**ToHeadimg** | **string** | 回复对象头像 | [optional] [default to null]
**ReplyContent** | **string** | 回复内容 | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | 回复时间 | [optional] [default to null]
**LikeNum** | **int32** | 点赞次数 | [optional] [default to null]
**LikeStatus** | **int32** | 当前登录用户对当前回复是否点赞。1-已赞；0-未赞 | [optional] [default to null]
**LikeId** | **int64** | 点赞记录Id | [optional] [default to null]
**SubReply** | [**[]WctApiApplicationProductCommentDtoReplyDto**](WCT.Api.Application.ProductComment.Dto.ReplyDto.md) | 当前回复记录下的所有回复 | [optional] [default to null]
**Approve** | [***WctAdminProductCommentCommentApprove**](WCT.Admin.ProductComment.CommentApprove.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

