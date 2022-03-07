# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppShareGetSharInfoGet**](ShareApi.md#ApiServicesAppShareGetSharInfoGet) | **Get** /api/services/app/Share/GetSharInfo | 
[**ApiServicesAppShareShareCallBackGet**](ShareApi.md#ApiServicesAppShareShareCallBackGet) | **Get** /api/services/app/Share/ShareCallBack | 
[**ApiServicesAppShareShareProductPost**](ShareApi.md#ApiServicesAppShareShareProductPost) | **Post** /api/services/app/Share/ShareProduct | 分享商品 触发积分赠送

# **ApiServicesAppShareGetSharInfoGet**
> WctApiApplicationShareDtoShareInfo ApiServicesAppShareGetSharInfoGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShareApiApiServicesAppShareGetSharInfoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShareApiApiServicesAppShareGetSharInfoGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shareType** | [**optional.Interface of WctApiApplicationShareDtoShareTypeEnum**](.md)|  | 

### Return type

[**WctApiApplicationShareDtoShareInfo**](WCT.Api.Application.Share.Dto.ShareInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppShareShareCallBackGet**
> ApiServicesAppShareShareCallBackGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShareApiApiServicesAppShareShareCallBackGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShareApiApiServicesAppShareShareCallBackGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**|  | 
 **gameId** | **optional.Int64**|  | 
 **gameUserId** | **optional.Int64**| 分享人 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppShareShareProductPost**
> ApiServicesAppShareShareProductPost(ctx, optional)
分享商品 触发积分赠送

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShareApiApiServicesAppShareShareProductPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShareApiApiServicesAppShareShareProductPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **optional.Int64**| 商品ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

