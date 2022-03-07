# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppSingleTranslationTranslatePost**](SingleTranslationApi.md#ApiServicesAppSingleTranslationTranslatePost) | **Post** /api/services/app/SingleTranslation/Translate | 翻译

# **ApiServicesAppSingleTranslationTranslatePost**
> WctApiApplicationCommonResultDto ApiServicesAppSingleTranslationTranslatePost(ctx, optional)
翻译

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SingleTranslationApiApiServicesAppSingleTranslationTranslatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SingleTranslationApiApiServicesAppSingleTranslationTranslatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | **optional.String**| 需翻译的内容 | 
 **from** | **optional.String**| 源,如：zh,en,ar(阿拉伯语) | 
 **to** | **optional.String**| 目标,如：zh,en,ar(阿拉伯语) | 

### Return type

[**WctApiApplicationCommonResultDto**](WCT.Api.Application.Common.ResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

