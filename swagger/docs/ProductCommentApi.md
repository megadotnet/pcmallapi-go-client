# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppProductCommentCommentLikeGet**](ProductCommentApi.md#ApiServicesAppProductCommentCommentLikeGet) | **Get** /api/services/app/ProductComment/CommentLike | 对评价点赞或者取消点赞
[**ApiServicesAppProductCommentCreateCommentPost**](ProductCommentApi.md#ApiServicesAppProductCommentCreateCommentPost) | **Post** /api/services/app/ProductComment/CreateComment | 评价（首评+追评）
[**ApiServicesAppProductCommentCreateReplyPost**](ProductCommentApi.md#ApiServicesAppProductCommentCreateReplyPost) | **Post** /api/services/app/ProductComment/CreateReply | 回复
[**ApiServicesAppProductCommentGetCommentDetailGet**](ProductCommentApi.md#ApiServicesAppProductCommentGetCommentDetailGet) | **Get** /api/services/app/ProductComment/GetCommentDetail | 评价详情
[**ApiServicesAppProductCommentGetMyCommentCountGet**](ProductCommentApi.md#ApiServicesAppProductCommentGetMyCommentCountGet) | **Get** /api/services/app/ProductComment/GetMyCommentCount | 获取评论数
[**ApiServicesAppProductCommentMyCommentPost**](ProductCommentApi.md#ApiServicesAppProductCommentMyCommentPost) | **Post** /api/services/app/ProductComment/MyComment | 我的评价
[**ApiServicesAppProductCommentProductCommentListPost**](ProductCommentApi.md#ApiServicesAppProductCommentProductCommentListPost) | **Post** /api/services/app/ProductComment/ProductCommentList | 获取商品评价列表
[**ApiServicesAppProductCommentReplyLikeGet**](ProductCommentApi.md#ApiServicesAppProductCommentReplyLikeGet) | **Get** /api/services/app/ProductComment/ReplyLike | 对回复点赞或者取消点赞
[**ApiServicesAppProductCommentReplyListPost**](ProductCommentApi.md#ApiServicesAppProductCommentReplyListPost) | **Post** /api/services/app/ProductComment/ReplyList | 回复列表

# **ApiServicesAppProductCommentCommentLikeGet**
> bool ApiServicesAppProductCommentCommentLikeGet(ctx, optional)
对评价点赞或者取消点赞

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductCommentApiApiServicesAppProductCommentCommentLikeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCommentApiApiServicesAppProductCommentCommentLikeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commentId** | **optional.Int64**| 首评Id | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductCommentCreateCommentPost**
> bool ApiServicesAppProductCommentCreateCommentPost(ctx, optional)
评价（首评+追评）

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductCommentApiApiServicesAppProductCommentCreateCommentPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCommentApiApiServicesAppProductCommentCreateCommentPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationProductCommentDtoCreateCommentDto**](WctApiApplicationProductCommentDtoCreateCommentDto.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductCommentCreateReplyPost**
> WctApiApplicationProductCommentDtoCreateReplyDto ApiServicesAppProductCommentCreateReplyPost(ctx, optional)
回复

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductCommentApiApiServicesAppProductCommentCreateReplyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCommentApiApiServicesAppProductCommentCreateReplyPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationProductCommentDtoCreateReplyDto**](WctApiApplicationProductCommentDtoCreateReplyDto.md)|  | 

### Return type

[**WctApiApplicationProductCommentDtoCreateReplyDto**](WCT.Api.Application.ProductComment.Dto.CreateReplyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductCommentGetCommentDetailGet**
> []WctApiApplicationProductCommentDtoCommentDetailDto ApiServicesAppProductCommentGetCommentDetailGet(ctx, optional)
评价详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductCommentApiApiServicesAppProductCommentGetCommentDetailGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCommentApiApiServicesAppProductCommentGetCommentDetailGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commentId** | **optional.Int64**| 首评Id | 

### Return type

[**[]WctApiApplicationProductCommentDtoCommentDetailDto**](WCT.Api.Application.ProductComment.Dto.CommentDetailDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductCommentGetMyCommentCountGet**
> int64 ApiServicesAppProductCommentGetMyCommentCountGet(ctx, )
获取评论数

### Required Parameters
This endpoint does not need any parameter.

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductCommentMyCommentPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductCommentDtoMyCommentDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppProductCommentMyCommentPost(ctx, optional)
我的评价

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductCommentApiApiServicesAppProductCommentMyCommentPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCommentApiApiServicesAppProductCommentMyCommentPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationProductCommentDtoGetMyCommentInput**](WctApiApplicationProductCommentDtoGetMyCommentInput.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductCommentDtoMyCommentDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.ProductComment.Dto.MyCommentDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductCommentProductCommentListPost**
> WctApiApplicationProductCommentDtoProductCommentDto ApiServicesAppProductCommentProductCommentListPost(ctx, optional)
获取商品评价列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductCommentApiApiServicesAppProductCommentProductCommentListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCommentApiApiServicesAppProductCommentProductCommentListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationProductCommentDtoGetProductCommentInput**](WctApiApplicationProductCommentDtoGetProductCommentInput.md)|  | 

### Return type

[**WctApiApplicationProductCommentDtoProductCommentDto**](WCT.Api.Application.ProductComment.Dto.ProductCommentDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductCommentReplyLikeGet**
> bool ApiServicesAppProductCommentReplyLikeGet(ctx, optional)
对回复点赞或者取消点赞

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductCommentApiApiServicesAppProductCommentReplyLikeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCommentApiApiServicesAppProductCommentReplyLikeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replyId** | **optional.Int64**| 回复记录Id | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductCommentReplyListPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductCommentDtoReplyDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppProductCommentReplyListPost(ctx, optional)
回复列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductCommentApiApiServicesAppProductCommentReplyListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCommentApiApiServicesAppProductCommentReplyListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationProductCommentDtoGetReplyInput**](WctApiApplicationProductCommentDtoGetReplyInput.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductCommentDtoReplyDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.ProductComment.Dto.ReplyDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

