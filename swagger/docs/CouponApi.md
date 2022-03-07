# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppCouponBuildCouponGet**](CouponApi.md#ApiServicesAppCouponBuildCouponGet) | **Get** /api/services/app/Coupon/BuildCoupon | 创建优惠券信息
[**ApiServicesAppCouponCouponCountPost**](CouponApi.md#ApiServicesAppCouponCouponCountPost) | **Post** /api/services/app/Coupon/CouponCount | 可使用优惠券数量
[**ApiServicesAppCouponDetailPost**](CouponApi.md#ApiServicesAppCouponDetailPost) | **Post** /api/services/app/Coupon/Detail | 优惠券详情(前端)
[**ApiServicesAppCouponExchangePost**](CouponApi.md#ApiServicesAppCouponExchangePost) | **Post** /api/services/app/Coupon/Exchange | 优惠码兑换优惠券
[**ApiServicesAppCouponProductConsumesPost**](CouponApi.md#ApiServicesAppCouponProductConsumesPost) | **Post** /api/services/app/Coupon/ProductConsumes | 商品可消费优惠券(商品详情用)
[**ApiServicesAppCouponRamdomCouponCodeGet**](CouponApi.md#ApiServicesAppCouponRamdomCouponCodeGet) | **Get** /api/services/app/Coupon/RamdomCouponCode | 随机抽取一个未领取的优惠码信息
[**ApiServicesAppCouponSearchCouponListPost**](CouponApi.md#ApiServicesAppCouponSearchCouponListPost) | **Post** /api/services/app/Coupon/SearchCouponList | 获取优惠券列表
[**ApiServicesAppCouponSearchPost**](CouponApi.md#ApiServicesAppCouponSearchPost) | **Post** /api/services/app/Coupon/Search | 前端查询优惠券
[**ApiServicesAppCouponSearchProductPost**](CouponApi.md#ApiServicesAppCouponSearchProductPost) | **Post** /api/services/app/Coupon/SearchProduct | 搜索优惠券满足的商品
[**ApiServicesAppCouponSearchTenantCouponPost**](CouponApi.md#ApiServicesAppCouponSearchTenantCouponPost) | **Post** /api/services/app/Coupon/SearchTenantCoupon | 获取商家优惠券
[**ApiServicesAppCouponTobeConsumesPost**](CouponApi.md#ApiServicesAppCouponTobeConsumesPost) | **Post** /api/services/app/Coupon/TobeConsumes | 商品可消费优惠券(下单用)

# **ApiServicesAppCouponBuildCouponGet**
> bool ApiServicesAppCouponBuildCouponGet(ctx, optional)
创建优惠券信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CouponApiApiServicesAppCouponBuildCouponGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CouponApiApiServicesAppCouponBuildCouponGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **optional.String**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppCouponCouponCountPost**
> int32 ApiServicesAppCouponCouponCountPost(ctx, )
可使用优惠券数量

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

# **ApiServicesAppCouponDetailPost**
> WctApiApplicationCouponsDtoCouponDetailDto ApiServicesAppCouponDetailPost(ctx, optional)
优惠券详情(前端)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CouponApiApiServicesAppCouponDetailPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CouponApiApiServicesAppCouponDetailPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationCouponsDtoCouponDto**](WctApiApplicationCouponsDtoCouponDto.md)|  | 

### Return type

[**WctApiApplicationCouponsDtoCouponDetailDto**](WCT.Api.Application.Coupons.dto.CouponDetailDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppCouponExchangePost**
> bool ApiServicesAppCouponExchangePost(ctx, code)
优惠码兑换优惠券

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppCouponProductConsumesPost**
> []WctApiApplicationCouponsDtoCouponProductResult ApiServicesAppCouponProductConsumesPost(ctx, optional)
商品可消费优惠券(商品详情用)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CouponApiApiServicesAppCouponProductConsumesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CouponApiApiServicesAppCouponProductConsumesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiCouponsDtoCouponConsumeDto**](WctApiCouponsDtoCouponConsumeDto.md)|  | 

### Return type

[**[]WctApiApplicationCouponsDtoCouponProductResult**](WCT.Api.Application.Coupons.dto.CouponProductResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppCouponRamdomCouponCodeGet**
> string ApiServicesAppCouponRamdomCouponCodeGet(ctx, optional)
随机抽取一个未领取的优惠码信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CouponApiApiServicesAppCouponRamdomCouponCodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CouponApiApiServicesAppCouponRamdomCouponCodeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **optional.String**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppCouponSearchCouponListPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationCouponsDtoCouponProductResultWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppCouponSearchCouponListPost(ctx, optional)
获取优惠券列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CouponApiApiServicesAppCouponSearchCouponListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CouponApiApiServicesAppCouponSearchCouponListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationCouponsDtoGetCouponListInput**](WctApiApplicationCouponsDtoGetCouponListInput.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationCouponsDtoCouponProductResultWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Coupons.dto.CouponProductResult, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppCouponSearchPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationCouponsDtoCouponSearchResultWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppCouponSearchPost(ctx, optional)
前端查询优惠券

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CouponApiApiServicesAppCouponSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CouponApiApiServicesAppCouponSearchPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationCouponsDtoCouponSearchDto**](WctApiApplicationCouponsDtoCouponSearchDto.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationCouponsDtoCouponSearchResultWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Coupons.dto.CouponSearchResult, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppCouponSearchProductPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationSearchDtoSearchProductDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppCouponSearchProductPost(ctx, optional)
搜索优惠券满足的商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CouponApiApiServicesAppCouponSearchProductPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CouponApiApiServicesAppCouponSearchProductPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationCouponsDtoCouponExtendDto**](WctApiApplicationCouponsDtoCouponExtendDto.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationSearchDtoSearchProductDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Search.Dto.SearchProductDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppCouponSearchTenantCouponPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationCouponsDtoCouponProductResultWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppCouponSearchTenantCouponPost(ctx, optional)
获取商家优惠券

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CouponApiApiServicesAppCouponSearchTenantCouponPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CouponApiApiServicesAppCouponSearchTenantCouponPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationCouponsDtoTenantCouponSearchDto**](WctApiApplicationCouponsDtoTenantCouponSearchDto.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationCouponsDtoCouponProductResultWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Coupons.dto.CouponProductResult, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppCouponTobeConsumesPost**
> []WctApiCouponsDtoTobeConsumeOutput ApiServicesAppCouponTobeConsumesPost(ctx, optional)
商品可消费优惠券(下单用)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CouponApiApiServicesAppCouponTobeConsumesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CouponApiApiServicesAppCouponTobeConsumesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiCouponsDtoCouponConsumeDto**](WctApiCouponsDtoCouponConsumeDto.md)|  | 

### Return type

[**[]WctApiCouponsDtoTobeConsumeOutput**](WCT.Api.Coupons.dto.TobeConsumeOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

