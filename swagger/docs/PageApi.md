# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppPageFormatProductPost**](PageApi.md#ApiServicesAppPageFormatProductPost) | **Post** /api/services/app/Page/FormatProduct | 
[**ApiServicesAppPageGetPCHomePageGet**](PageApi.md#ApiServicesAppPageGetPCHomePageGet) | **Get** /api/services/app/Page/GetPCHomePage | 
[**ApiServicesAppPageGetPCHomePageItemsGet**](PageApi.md#ApiServicesAppPageGetPCHomePageItemsGet) | **Get** /api/services/app/Page/GetPCHomePageItems | 
[**ApiServicesAppPageGetPageGet**](PageApi.md#ApiServicesAppPageGetPageGet) | **Get** /api/services/app/Page/GetPage | 

# **ApiServicesAppPageFormatProductPost**
> string ApiServicesAppPageFormatProductPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PageApiApiServicesAppPageFormatProductPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PageApiApiServicesAppPageFormatProductPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageId** | **optional.Int32**|  | 
 **platform** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **tenantId** | **optional.Int32**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppPageGetPCHomePageGet**
> WctShowCasePageCacheItem ApiServicesAppPageGetPCHomePageGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WctShowCasePageCacheItem**](WCT.ShowCase.PageCacheItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppPageGetPCHomePageItemsGet**
> ApiServicesAppPageGetPCHomePageItemsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PageApiApiServicesAppPageGetPCHomePageItemsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PageApiApiServicesAppPageGetPCHomePageItemsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppPageGetPageGet**
> WctShowCasePageCacheItem ApiServicesAppPageGetPageGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PageApiApiServicesAppPageGetPageGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PageApiApiServicesAppPageGetPageGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageId** | **optional.Int32**|  | 
 **platform** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **tenantId** | **optional.Int32**|  | 

### Return type

[**WctShowCasePageCacheItem**](WCT.ShowCase.PageCacheItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

