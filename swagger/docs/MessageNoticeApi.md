# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppMessageNoticeGetTotalUnreadGet**](MessageNoticeApi.md#ApiServicesAppMessageNoticeGetTotalUnreadGet) | **Get** /api/services/app/MessageNotice/GetTotalUnread | 获取消息未读总数
[**ApiServicesAppMessageNoticeQueryServiceMessagesPost**](MessageNoticeApi.md#ApiServicesAppMessageNoticeQueryServiceMessagesPost) | **Post** /api/services/app/MessageNotice/QueryServiceMessages | 查询消息通知
[**ApiServicesAppMessageNoticeSaveMessageStatusPost**](MessageNoticeApi.md#ApiServicesAppMessageNoticeSaveMessageStatusPost) | **Post** /api/services/app/MessageNotice/SaveMessageStatus | 消息已读（所有or单条）

# **ApiServicesAppMessageNoticeGetTotalUnreadGet**
> int32 ApiServicesAppMessageNoticeGetTotalUnreadGet(ctx, )
获取消息未读总数

### Required Parameters
This endpoint does not need any parameter.

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppMessageNoticeQueryServiceMessagesPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationServiceMessageDtoMessageOutputDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppMessageNoticeQueryServiceMessagesPost(ctx, optional)
查询消息通知

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessageNoticeApiApiServicesAppMessageNoticeQueryServiceMessagesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessageNoticeApiApiServicesAppMessageNoticeQueryServiceMessagesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationServiceMessageDtoMessageInputDto**](WctApiApplicationServiceMessageDtoMessageInputDto.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationServiceMessageDtoMessageOutputDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.ServiceMessage.Dto.MessageOutputDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppMessageNoticeSaveMessageStatusPost**
> bool ApiServicesAppMessageNoticeSaveMessageStatusPost(ctx, optional)
消息已读（所有or单条）

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessageNoticeApiApiServicesAppMessageNoticeSaveMessageStatusPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessageNoticeApiApiServicesAppMessageNoticeSaveMessageStatusPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationServiceMessageDtoMessageStatusInputDto**](WctApiApplicationServiceMessageDtoMessageStatusInputDto.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

