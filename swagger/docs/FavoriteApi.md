# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppFavoriteAddPost**](FavoriteApi.md#ApiServicesAppFavoriteAddPost) | **Post** /api/services/app/Favorite/Add | 新增收藏
[**ApiServicesAppFavoriteBacthAddPost**](FavoriteApi.md#ApiServicesAppFavoriteBacthAddPost) | **Post** /api/services/app/Favorite/BacthAdd | 批量新增收藏
[**ApiServicesAppFavoriteBacthDeletePost**](FavoriteApi.md#ApiServicesAppFavoriteBacthDeletePost) | **Post** /api/services/app/Favorite/BacthDelete | 删除收藏
[**ApiServicesAppFavoriteBrushAddPost**](FavoriteApi.md#ApiServicesAppFavoriteBrushAddPost) | **Post** /api/services/app/Favorite/BrushAdd | 新增刷收藏
[**ApiServicesAppFavoriteDeleteByProductIdPost**](FavoriteApi.md#ApiServicesAppFavoriteDeleteByProductIdPost) | **Post** /api/services/app/Favorite/DeleteByProductId | 删除收藏
[**ApiServicesAppFavoriteDeletePost**](FavoriteApi.md#ApiServicesAppFavoriteDeletePost) | **Post** /api/services/app/Favorite/Delete | 删除收藏
[**ApiServicesAppFavoriteFavoriteCountPost**](FavoriteApi.md#ApiServicesAppFavoriteFavoriteCountPost) | **Post** /api/services/app/Favorite/FavoriteCount | 收藏数量
[**ApiServicesAppFavoriteJudgePost**](FavoriteApi.md#ApiServicesAppFavoriteJudgePost) | **Post** /api/services/app/Favorite/Judge | 判断是否收藏
[**ApiServicesAppFavoriteSearchPost**](FavoriteApi.md#ApiServicesAppFavoriteSearchPost) | **Post** /api/services/app/Favorite/Search | 查询收藏列表

# **ApiServicesAppFavoriteAddPost**
> bool ApiServicesAppFavoriteAddPost(ctx, optional)
新增收藏

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoriteApiApiServicesAppFavoriteAddPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoriteApiApiServicesAppFavoriteAddPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctAdminFavoriteFavoriteMst**](WctAdminFavoriteFavoriteMst.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFavoriteBacthAddPost**
> bool ApiServicesAppFavoriteBacthAddPost(ctx, optional)
批量新增收藏

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoriteApiApiServicesAppFavoriteBacthAddPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoriteApiApiServicesAppFavoriteBacthAddPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []WctAdminFavoriteFavoriteMst**](WCT.Admin.Favorite.FavoriteMst.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFavoriteBacthDeletePost**
> bool ApiServicesAppFavoriteBacthDeletePost(ctx, optional)
删除收藏

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoriteApiApiServicesAppFavoriteBacthDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoriteApiApiServicesAppFavoriteBacthDeletePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []int64**](int64.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFavoriteBrushAddPost**
> bool ApiServicesAppFavoriteBrushAddPost(ctx, optional)
新增刷收藏

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoriteApiApiServicesAppFavoriteBrushAddPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoriteApiApiServicesAppFavoriteBrushAddPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationFavoriteDtoBrushFavoriteDto**](WctApiApplicationFavoriteDtoBrushFavoriteDto.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFavoriteDeleteByProductIdPost**
> bool ApiServicesAppFavoriteDeleteByProductIdPost(ctx, optional)
删除收藏

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoriteApiApiServicesAppFavoriteDeleteByProductIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoriteApiApiServicesAppFavoriteDeleteByProductIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctAdminFavoriteFavoriteMst**](WctAdminFavoriteFavoriteMst.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFavoriteDeletePost**
> bool ApiServicesAppFavoriteDeletePost(ctx, optional)
删除收藏

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoriteApiApiServicesAppFavoriteDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoriteApiApiServicesAppFavoriteDeletePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFavoriteFavoriteCountPost**
> int32 ApiServicesAppFavoriteFavoriteCountPost(ctx, )
收藏数量

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

# **ApiServicesAppFavoriteJudgePost**
> int64 ApiServicesAppFavoriteJudgePost(ctx, optional)
判断是否收藏

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoriteApiApiServicesAppFavoriteJudgePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoriteApiApiServicesAppFavoriteJudgePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctAdminFavoriteFavoriteMst**](WctAdminFavoriteFavoriteMst.md)|  | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFavoriteSearchPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationFavoriteDtoFavoriteDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppFavoriteSearchPost(ctx, optional)
查询收藏列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FavoriteApiApiServicesAppFavoriteSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FavoriteApiApiServicesAppFavoriteSearchPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationFavoriteDtoFavoriteSearch**](WctApiApplicationFavoriteDtoFavoriteSearch.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationFavoriteDtoFavoriteDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Favorite.Dto.FavoriteDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

