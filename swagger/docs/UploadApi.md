# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppUploadChatUploadAttachmentPost**](UploadApi.md#ApiServicesAppUploadChatUploadAttachmentPost) | **Post** /api/services/app/Upload/ChatUploadAttachment | 客服聊天 发送文件
[**ApiServicesAppUploadUploadFilePost**](UploadApi.md#ApiServicesAppUploadUploadFilePost) | **Post** /api/services/app/Upload/UploadFile | 图片上传
[**ApiServicesAppUploadUploadImagePost**](UploadApi.md#ApiServicesAppUploadUploadImagePost) | **Post** /api/services/app/Upload/UploadImage | base64格式图片上传

# **ApiServicesAppUploadChatUploadAttachmentPost**
> WctApiApplicationUploadFilesDtoUploadDto ApiServicesAppUploadChatUploadAttachmentPost(ctx, optional)
客服聊天 发送文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UploadApiApiServicesAppUploadChatUploadAttachmentPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadApiApiServicesAppUploadChatUploadAttachmentPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **chatUploadType** | **optional.**|  | 

### Return type

[**WctApiApplicationUploadFilesDtoUploadDto**](WCT.Api.Application.UploadFiles.Dto.UploadDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppUploadUploadFilePost**
> WctApiApplicationUploadFilesDtoUploadDto ApiServicesAppUploadUploadFilePost(ctx, optional)
图片上传

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UploadApiApiServicesAppUploadUploadFilePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadApiApiServicesAppUploadUploadFilePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**WctApiApplicationUploadFilesDtoUploadDto**](WCT.Api.Application.UploadFiles.Dto.UploadDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppUploadUploadImagePost**
> WctApiApplicationUploadFilesDtoUploadDto ApiServicesAppUploadUploadImagePost(ctx, optional)
base64格式图片上传

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UploadApiApiServicesAppUploadUploadImagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadApiApiServicesAppUploadUploadImagePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of string**](string.md)|  | 

### Return type

[**WctApiApplicationUploadFilesDtoUploadDto**](WCT.Api.Application.UploadFiles.Dto.UploadDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

