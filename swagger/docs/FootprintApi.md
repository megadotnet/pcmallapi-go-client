# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppFootprintAddFootprintPost**](FootprintApi.md#ApiServicesAppFootprintAddFootprintPost) | **Post** /api/services/app/Footprint/AddFootprint | 添加轨迹
[**ApiServicesAppFootprintBatchDeletePost**](FootprintApi.md#ApiServicesAppFootprintBatchDeletePost) | **Post** /api/services/app/Footprint/BatchDelete | 批量删除
[**ApiServicesAppFootprintDeleteAllPost**](FootprintApi.md#ApiServicesAppFootprintDeleteAllPost) | **Post** /api/services/app/Footprint/DeleteAll | 批量删除
[**ApiServicesAppFootprintDeletePost**](FootprintApi.md#ApiServicesAppFootprintDeletePost) | **Post** /api/services/app/Footprint/Delete | 删除
[**ApiServicesAppFootprintFootprintCountPost**](FootprintApi.md#ApiServicesAppFootprintFootprintCountPost) | **Post** /api/services/app/Footprint/FootprintCount | 足迹数量
[**ApiServicesAppFootprintFootprintListPost**](FootprintApi.md#ApiServicesAppFootprintFootprintListPost) | **Post** /api/services/app/Footprint/FootprintList | 获取轨迹信息

# **ApiServicesAppFootprintAddFootprintPost**
> string ApiServicesAppFootprintAddFootprintPost(ctx, optional)
添加轨迹

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FootprintApiApiServicesAppFootprintAddFootprintPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FootprintApiApiServicesAppFootprintAddFootprintPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationFootprintDtoFootprintDto**](WctApiApplicationFootprintDtoFootprintDto.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFootprintBatchDeletePost**
> bool ApiServicesAppFootprintBatchDeletePost(ctx, optional)
批量删除

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FootprintApiApiServicesAppFootprintBatchDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FootprintApiApiServicesAppFootprintBatchDeletePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []WctApiApplicationFootprintDtoFootprintDto**](WCT.Api.Application.Footprint.Dto.FootprintDto.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFootprintDeleteAllPost**
> bool ApiServicesAppFootprintDeleteAllPost(ctx, )
批量删除

### Required Parameters
This endpoint does not need any parameter.

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFootprintDeletePost**
> bool ApiServicesAppFootprintDeletePost(ctx, optional)
删除

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FootprintApiApiServicesAppFootprintDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FootprintApiApiServicesAppFootprintDeletePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationFootprintDtoFootprintDto**](WctApiApplicationFootprintDtoFootprintDto.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFootprintFootprintCountPost**
> int32 ApiServicesAppFootprintFootprintCountPost(ctx, )
足迹数量

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

# **ApiServicesAppFootprintFootprintListPost**
> []WctApiApplicationFootprintDtoFootprintDto ApiServicesAppFootprintFootprintListPost(ctx, )
获取轨迹信息

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]WctApiApplicationFootprintDtoFootprintDto**](WCT.Api.Application.Footprint.Dto.FootprintDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

