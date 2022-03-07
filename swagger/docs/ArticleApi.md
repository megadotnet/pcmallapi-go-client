# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppArticleGetAllGet**](ArticleApi.md#ApiServicesAppArticleGetAllGet) | **Get** /api/services/app/Article/GetAll | 
[**ApiServicesAppArticleGetArticlesPost**](ArticleApi.md#ApiServicesAppArticleGetArticlesPost) | **Post** /api/services/app/Article/GetArticles | 
[**ApiServicesAppArticleGetGet**](ArticleApi.md#ApiServicesAppArticleGetGet) | **Get** /api/services/app/Article/Get | 

# **ApiServicesAppArticleGetAllGet**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationNewsDtoArticleListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppArticleGetAllGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ArticleApiApiServicesAppArticleGetAllGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticleApiApiServicesAppArticleGetAllGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **optional.Int32**|  | 
 **skipCount** | **optional.Int32**|  | 
 **maxResultCount** | **optional.Int32**|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationNewsDtoArticleListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.News.Dto.ArticleListItemDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppArticleGetArticlesPost**
> []WctApiApplicationNewsDtoArticleListItemDto ApiServicesAppArticleGetArticlesPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ArticleApiApiServicesAppArticleGetArticlesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticleApiApiServicesAppArticleGetArticlesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationNewsDtoGetArticlesInput**](WctApiApplicationNewsDtoGetArticlesInput.md)|  | 

### Return type

[**[]WctApiApplicationNewsDtoArticleListItemDto**](WCT.Api.Application.News.Dto.ArticleListItemDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppArticleGetGet**
> WctEntitiesNewsCacheArticleCacheItem ApiServicesAppArticleGetGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ArticleApiApiServicesAppArticleGetGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticleApiApiServicesAppArticleGetGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int32**|  | 

### Return type

[**WctEntitiesNewsCacheArticleCacheItem**](WCT.Entities.News.Cache.ArticleCacheItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

