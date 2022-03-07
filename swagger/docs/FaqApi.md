# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppFaqDetailGet**](FaqApi.md#ApiServicesAppFaqDetailGet) | **Get** /api/services/app/Faq/Detail | FAQ详情
[**ApiServicesAppFaqFaqCateogryListPost**](FaqApi.md#ApiServicesAppFaqFaqCateogryListPost) | **Post** /api/services/app/Faq/FaqCateogryList | 常见问题分类列表
[**ApiServicesAppFaqFaqListPost**](FaqApi.md#ApiServicesAppFaqFaqListPost) | **Post** /api/services/app/Faq/FaqList | 常见问题分类下的问题列表
[**ApiServicesAppFaqHotFaqListGet**](FaqApi.md#ApiServicesAppFaqHotFaqListGet) | **Get** /api/services/app/Faq/HotFaqList | 热门问题
[**ApiServicesAppFaqLatestFAQGet**](FaqApi.md#ApiServicesAppFaqLatestFAQGet) | **Get** /api/services/app/Faq/LatestFAQ | 最新FAQ

# **ApiServicesAppFaqDetailGet**
> WctApiFaqDtoFaqDetailDto ApiServicesAppFaqDetailGet(ctx, optional)
FAQ详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FaqApiApiServicesAppFaqDetailGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FaqApiApiServicesAppFaqDetailGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int32**|  | 

### Return type

[**WctApiFaqDtoFaqDetailDto**](WCT.Api.FAQ.Dto.FaqDetailDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFaqFaqCateogryListPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiFaqDtoFaqCateogryListOutputWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppFaqFaqCateogryListPost(ctx, optional)
常见问题分类列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FaqApiApiServicesAppFaqFaqCateogryListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FaqApiApiServicesAppFaqFaqCateogryListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiFaqDtoFaqCateogryListDto**](WctApiFaqDtoFaqCateogryListDto.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiFaqDtoFaqCateogryListOutputWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.FAQ.Dto.FaqCateogryListOutput, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFaqFaqListPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiFaqDtoFaqListOutputWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppFaqFaqListPost(ctx, optional)
常见问题分类下的问题列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FaqApiApiServicesAppFaqFaqListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FaqApiApiServicesAppFaqFaqListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiFaqDtoFaqListDto**](WctApiFaqDtoFaqListDto.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiFaqDtoFaqListOutputWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.FAQ.Dto.FaqListOutput, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFaqHotFaqListGet**
> []WctApiFaqDtoHotFaqDto ApiServicesAppFaqHotFaqListGet(ctx, optional)
热门问题

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FaqApiApiServicesAppFaqHotFaqListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FaqApiApiServicesAppFaqHotFaqListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hotCategoryCount** | **optional.Int32**| 显示多少个FAQ分类版块，默认6个 | [default to 6]

### Return type

[**[]WctApiFaqDtoHotFaqDto**](WCT.Api.FAQ.Dto.HotFaqDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFaqLatestFAQGet**
> []WctApiFaqDtoFaqListOutput ApiServicesAppFaqLatestFAQGet(ctx, optional)
最新FAQ

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FaqApiApiServicesAppFaqLatestFAQGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FaqApiApiServicesAppFaqLatestFAQGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**|  | [default to 5]

### Return type

[**[]WctApiFaqDtoFaqListOutput**](WCT.Api.FAQ.Dto.FaqListOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

