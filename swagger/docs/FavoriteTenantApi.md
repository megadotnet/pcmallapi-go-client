# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppFavoriteTenantAddPost**](FavoriteTenantApi.md#ApiServicesAppFavoriteTenantAddPost) | **Post** /api/services/app/FavoriteTenant/Add | 添加收藏
[**ApiServicesAppFavoriteTenantGetMyFavoriteCountGet**](FavoriteTenantApi.md#ApiServicesAppFavoriteTenantGetMyFavoriteCountGet) | **Get** /api/services/app/FavoriteTenant/GetMyFavoriteCount | 获取收藏商家当量
[**ApiServicesAppFavoriteTenantGetMyFavoritesGet**](FavoriteTenantApi.md#ApiServicesAppFavoriteTenantGetMyFavoritesGet) | **Get** /api/services/app/FavoriteTenant/GetMyFavorites | 收藏的店铺
[**ApiServicesAppFavoriteTenantRemovePost**](FavoriteTenantApi.md#ApiServicesAppFavoriteTenantRemovePost) | **Post** /api/services/app/FavoriteTenant/Remove | 移除收藏

# **ApiServicesAppFavoriteTenantAddPost**
> ApiServicesAppFavoriteTenantAddPost(ctx, optional)
添加收藏

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoriteTenantApiApiServicesAppFavoriteTenantAddPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoriteTenantApiApiServicesAppFavoriteTenantAddPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationTenantsDtosTenantIdDto**](WctApiApplicationTenantsDtosTenantIdDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFavoriteTenantGetMyFavoriteCountGet**
> int64 ApiServicesAppFavoriteTenantGetMyFavoriteCountGet(ctx, )
获取收藏商家当量

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

# **ApiServicesAppFavoriteTenantGetMyFavoritesGet**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationTenantsDtosTenantAndProductsDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppFavoriteTenantGetMyFavoritesGet(ctx, optional)
收藏的店铺

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoriteTenantApiApiServicesAppFavoriteTenantGetMyFavoritesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoriteTenantApiApiServicesAppFavoriteTenantGetMyFavoritesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skipCount** | **optional.Int32**|  | 
 **maxResultCount** | **optional.Int32**|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationTenantsDtosTenantAndProductsDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Tenants.Dtos.TenantAndProductsDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFavoriteTenantRemovePost**
> ApiServicesAppFavoriteTenantRemovePost(ctx, optional)
移除收藏

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoriteTenantApiApiServicesAppFavoriteTenantRemovePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoriteTenantApiApiServicesAppFavoriteTenantRemovePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **optional.Int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

