# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppFullReductionGetRuleSettingByIdGet**](FullReductionApi.md#ApiServicesAppFullReductionGetRuleSettingByIdGet) | **Get** /api/services/app/FullReduction/GetRuleSettingById | 根据活动ID获取满减活动商品列表优惠信息内容
[**ApiServicesAppFullReductionUpdateFullReductionCachePost**](FullReductionApi.md#ApiServicesAppFullReductionUpdateFullReductionCachePost) | **Post** /api/services/app/FullReduction/UpdateFullReductionCache | 更新缓存（手动调用）

# **ApiServicesAppFullReductionGetRuleSettingByIdGet**
> WctApiApplicationFullReductionsDtoGetRuleSettingByIdDto ApiServicesAppFullReductionGetRuleSettingByIdGet(ctx, optional)
根据活动ID获取满减活动商品列表优惠信息内容

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FullReductionApiApiServicesAppFullReductionGetRuleSettingByIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FullReductionApiApiServicesAppFullReductionGetRuleSettingByIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullReductionId** | **optional.Int64**|  | 

### Return type

[**WctApiApplicationFullReductionsDtoGetRuleSettingByIdDto**](WCT.Api.Application.FullReductions.Dto.GetRuleSettingByIdDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppFullReductionUpdateFullReductionCachePost**
> ApiServicesAppFullReductionUpdateFullReductionCachePost(ctx, )
更新缓存（手动调用）

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

