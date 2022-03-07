# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppProductGetCategoriesGet**](ProductApi.md#ApiServicesAppProductGetCategoriesGet) | **Get** /api/services/app/Product/GetCategories | 
[**ApiServicesAppProductGetDetailGet**](ProductApi.md#ApiServicesAppProductGetDetailGet) | **Get** /api/services/app/Product/GetDetail | 获取商品图文描述
[**ApiServicesAppProductGetExchangeRateGet**](ProductApi.md#ApiServicesAppProductGetExchangeRateGet) | **Get** /api/services/app/Product/GetExchangeRate | 最新汇率列表
[**ApiServicesAppProductGetGet**](ProductApi.md#ApiServicesAppProductGetGet) | **Get** /api/services/app/Product/Get | 获取商品详情
[**ApiServicesAppProductGetPicturesGet**](ProductApi.md#ApiServicesAppProductGetPicturesGet) | **Get** /api/services/app/Product/GetPictures | 获取商品滚动图
[**ApiServicesAppProductGetStockPost**](ProductApi.md#ApiServicesAppProductGetStockPost) | **Post** /api/services/app/Product/GetStock | 获取商品库存
[**ApiServicesAppProductGuessYouLikePost**](ProductApi.md#ApiServicesAppProductGuessYouLikePost) | **Post** /api/services/app/Product/GuessYouLike | 猜你喜欢
[**ApiServicesAppProductSearchProductsPost**](ProductApi.md#ApiServicesAppProductSearchProductsPost) | **Post** /api/services/app/Product/SearchProducts | 
[**ApiServicesAppProductTenantRoughInfoPost**](ProductApi.md#ApiServicesAppProductTenantRoughInfoPost) | **Post** /api/services/app/Product/TenantRoughInfo | 获取商户粗略信息  PS：商户LOGO、所有商品数量等

# **ApiServicesAppProductGetCategoriesGet**
> AbpApplicationServicesDtoListResultDto1WctApiApplicationProductsDtoCategoryDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppProductGetCategoriesGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AbpApplicationServicesDtoListResultDto1WctApiApplicationProductsDtoCategoryDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.ListResultDto&#x60;1[[WCT.Api.Application.Products.Dto.CategoryDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductGetDetailGet**
> string ApiServicesAppProductGetDetailGet(ctx, optional)
获取商品图文描述

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiApiServicesAppProductGetDetailGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiApiServicesAppProductGetDetailGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isPreview** | **optional.Bool**|  | 
 **source** | **optional.Int32**|  | 
 **id** | **optional.Int64**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductGetExchangeRateGet**
> []WctApiApplicationProductsDtoExchangeRateDto ApiServicesAppProductGetExchangeRateGet(ctx, )
最新汇率列表

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]WctApiApplicationProductsDtoExchangeRateDto**](WCT.Api.Application.Products.Dto.ExchangeRateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductGetGet**
> WctApiApplicationProductsDtoProductDto ApiServicesAppProductGetGet(ctx, optional)
获取商品详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiApiServicesAppProductGetGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiApiServicesAppProductGetGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isPreview** | **optional.Bool**|  | 
 **source** | **optional.Int32**|  | 
 **id** | **optional.Int64**|  | 

### Return type

[**WctApiApplicationProductsDtoProductDto**](WCT.Api.Application.Products.Dto.ProductDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductGetPicturesGet**
> []string ApiServicesAppProductGetPicturesGet(ctx, optional)
获取商品滚动图

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiApiServicesAppProductGetPicturesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiApiServicesAppProductGetPicturesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isPreview** | **optional.Bool**|  | 
 **source** | **optional.Int32**|  | 
 **id** | **optional.Int64**|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductGetStockPost**
> []WctApiCoreServicesProductsStockInfo ApiServicesAppProductGetStockPost(ctx, optional)
获取商品库存

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiApiServicesAppProductGetStockPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiApiServicesAppProductGetStockPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationProductsDtoGetStockInput**](WctApiApplicationProductsDtoGetStockInput.md)|  | 

### Return type

[**[]WctApiCoreServicesProductsStockInfo**](WCT.Api.Core.Services.Products.StockInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductGuessYouLikePost**
> []WctApiApplicationProductsDtoProductListItemDto ApiServicesAppProductGuessYouLikePost(ctx, optional)
猜你喜欢

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiApiServicesAppProductGuessYouLikePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiApiServicesAppProductGuessYouLikePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationProductsDtoGuessYouLikeDto**](WctApiApplicationProductsDtoGuessYouLikeDto.md)|  | 

### Return type

[**[]WctApiApplicationProductsDtoProductListItemDto**](WCT.Api.Application.Products.Dto.ProductListItemDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductSearchProductsPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppProductSearchProductsPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiApiServicesAppProductSearchProductsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiApiServicesAppProductSearchProductsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationProductsDtoSearchProductsInput**](WctApiApplicationProductsDtoSearchProductsInput.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationProductsDtoProductListItemDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Products.Dto.ProductListItemDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppProductTenantRoughInfoPost**
> WctApiApplicationProductsDtoTenantRoughDto ApiServicesAppProductTenantRoughInfoPost(ctx, optional)
获取商户粗略信息  PS：商户LOGO、所有商品数量等

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiApiServicesAppProductTenantRoughInfoPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiApiServicesAppProductTenantRoughInfoPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **optional.Int32**| 商家ID | 
 **showSize** | **optional.Int32**| 显示多少商品，默认3个 | [default to 3]

### Return type

[**WctApiApplicationProductsDtoTenantRoughDto**](WCT.Api.Application.Products.Dto.TenantRoughDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

