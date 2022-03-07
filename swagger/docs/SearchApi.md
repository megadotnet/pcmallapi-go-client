# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppSearchGetBrandsGet**](SearchApi.md#ApiServicesAppSearchGetBrandsGet) | **Get** /api/services/app/Search/GetBrands | 
[**ApiServicesAppSearchGetHotWordsGet**](SearchApi.md#ApiServicesAppSearchGetHotWordsGet) | **Get** /api/services/app/Search/GetHotWords | 热词
[**ApiServicesAppSearchGetRecommendGet**](SearchApi.md#ApiServicesAppSearchGetRecommendGet) | **Get** /api/services/app/Search/GetRecommend | 推荐商品
[**ApiServicesAppSearchGetSuggestsGet**](SearchApi.md#ApiServicesAppSearchGetSuggestsGet) | **Get** /api/services/app/Search/GetSuggests | 搜索商品
[**ApiServicesAppSearchMoreLikeThisOldPost**](SearchApi.md#ApiServicesAppSearchMoreLikeThisOldPost) | **Post** /api/services/app/Search/MoreLikeThis_Old | 相似商品（搜索引擎）
[**ApiServicesAppSearchMoreLikeThisPost**](SearchApi.md#ApiServicesAppSearchMoreLikeThisPost) | **Post** /api/services/app/Search/MoreLikeThis | 相似商品
[**ApiServicesAppSearchSearchProductPost**](SearchApi.md#ApiServicesAppSearchSearchProductPost) | **Post** /api/services/app/Search/SearchProduct | 
[**ApiServicesAppSearchSearchProductPropsPost**](SearchApi.md#ApiServicesAppSearchSearchProductPropsPost) | **Post** /api/services/app/Search/SearchProductProps | 
[**ApiServicesAppSearchSearchTenantPost**](SearchApi.md#ApiServicesAppSearchSearchTenantPost) | **Post** /api/services/app/Search/SearchTenant | 搜索店铺，包括销量排三的商品

# **ApiServicesAppSearchGetBrandsGet**
> AbpApplicationServicesDtoListResultDto1WctApiApplicationSearchDtoBrandDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppSearchGetBrandsGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiApiServicesAppSearchGetBrandsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiApiServicesAppSearchGetBrandsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **optional.Int32**|  | 

### Return type

[**AbpApplicationServicesDtoListResultDto1WctApiApplicationSearchDtoBrandDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.ListResultDto&#x60;1[[WCT.Api.Application.Search.Dto.BrandDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppSearchGetHotWordsGet**
> []string ApiServicesAppSearchGetHotWordsGet(ctx, )
热词

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppSearchGetRecommendGet**
> AbpApplicationServicesDtoListResultDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppSearchGetRecommendGet(ctx, optional)
推荐商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiApiServicesAppSearchGetRecommendGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiApiServicesAppSearchGetRecommendGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **optional.Int32**| 商户id | 
 **skipCount** | **optional.Int32**|  | 
 **maxResultCount** | **optional.Int32**|  | 

### Return type

[**AbpApplicationServicesDtoListResultDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.ListResultDto&#x60;1[[WCT.Api.Application.Products.Dto.ProductListItemDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppSearchGetSuggestsGet**
> []string ApiServicesAppSearchGetSuggestsGet(ctx, optional)
搜索商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiApiServicesAppSearchGetSuggestsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiApiServicesAppSearchGetSuggestsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keywords** | **optional.String**|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppSearchMoreLikeThisOldPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppSearchMoreLikeThisOldPost(ctx, optional)
相似商品（搜索引擎）

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiApiServicesAppSearchMoreLikeThisOldPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiApiServicesAppSearchMoreLikeThisOldPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationSearchDtoMoreLikeThisInput**](WctApiApplicationSearchDtoMoreLikeThisInput.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Products.Dto.ProductListItemDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppSearchMoreLikeThisPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppSearchMoreLikeThisPost(ctx, optional)
相似商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiApiServicesAppSearchMoreLikeThisPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiApiServicesAppSearchMoreLikeThisPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationSearchDtoMoreLikeThisInput**](WctApiApplicationSearchDtoMoreLikeThisInput.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Products.Dto.ProductListItemDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppSearchSearchProductPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppSearchSearchProductPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiApiServicesAppSearchSearchProductPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiApiServicesAppSearchSearchProductPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationSearchDtoGetPropsInput**](WctApiApplicationSearchDtoGetPropsInput.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Products.Dto.ProductListItemDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppSearchSearchProductPropsPost**
> WctApiApplicationSearchDtoPagedSearchProductDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppSearchSearchProductPropsPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiApiServicesAppSearchSearchProductPropsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiApiServicesAppSearchSearchProductPropsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationSearchDtoGetPropsInput**](WctApiApplicationSearchDtoGetPropsInput.md)|  | 

### Return type

[**WctApiApplicationSearchDtoPagedSearchProductDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](WCT.Api.Application.Search.Dto.PagedSearchProductDto&#x60;1[[WCT.Api.Application.Products.Dto.ProductListItemDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppSearchSearchTenantPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationSearchDtoTenantListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppSearchSearchTenantPost(ctx, optional)
搜索店铺，包括销量排三的商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiApiServicesAppSearchSearchTenantPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiApiServicesAppSearchSearchTenantPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationSearchDtoSearchInput**](WctApiApplicationSearchDtoSearchInput.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationSearchDtoTenantListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Search.Dto.TenantListItemDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

