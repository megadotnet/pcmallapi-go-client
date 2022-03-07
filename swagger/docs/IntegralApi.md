# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppIntegralGetIntegralExchangeProductsGet**](IntegralApi.md#ApiServicesAppIntegralGetIntegralExchangeProductsGet) | **Get** /api/services/app/Integral/GetIntegralExchangeProducts | 获得积分兑换商品
[**ApiServicesAppIntegralSearchIntegralArriveCashPost**](IntegralApi.md#ApiServicesAppIntegralSearchIntegralArriveCashPost) | **Post** /api/services/app/Integral/SearchIntegralArriveCash | 获取积分抵扣
[**ApiServicesAppIntegralSearchMyIntegralDtlPost**](IntegralApi.md#ApiServicesAppIntegralSearchMyIntegralDtlPost) | **Post** /api/services/app/Integral/SearchMyIntegralDtl | 获取我的积分明细
[**ApiServicesAppIntegralSearchMyIntegralTaskPost**](IntegralApi.md#ApiServicesAppIntegralSearchMyIntegralTaskPost) | **Post** /api/services/app/Integral/SearchMyIntegralTask | 获取我的任务

# **ApiServicesAppIntegralGetIntegralExchangeProductsGet**
> []WctApiApplicationIntegralDtoIntegralExchangeProduct ApiServicesAppIntegralGetIntegralExchangeProductsGet(ctx, optional)
获得积分兑换商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IntegralApiApiServicesAppIntegralGetIntegralExchangeProductsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegralApiApiServicesAppIntegralGetIntegralExchangeProductsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **optional.Int32**|  | 

### Return type

[**[]WctApiApplicationIntegralDtoIntegralExchangeProduct**](WCT.Api.Application.Integral.Dto.IntegralExchangeProduct.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppIntegralSearchIntegralArriveCashPost**
> WctAdminIntegralServiceDtoIntegralArriveCashOutput ApiServicesAppIntegralSearchIntegralArriveCashPost(ctx, optional)
获取积分抵扣

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IntegralApiApiServicesAppIntegralSearchIntegralArriveCashPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegralApiApiServicesAppIntegralSearchIntegralArriveCashPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctAdminIntegralServiceDtoIntegralArriveCashInput**](WctAdminIntegralServiceDtoIntegralArriveCashInput.md)|  | 

### Return type

[**WctAdminIntegralServiceDtoIntegralArriveCashOutput**](WCT.Admin.Integral.Service.Dto.IntegralArriveCashOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppIntegralSearchMyIntegralDtlPost**
> WctApiApplicationIntegralDtoIntegralDtlOutput ApiServicesAppIntegralSearchMyIntegralDtlPost(ctx, optional)
获取我的积分明细

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IntegralApiApiServicesAppIntegralSearchMyIntegralDtlPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegralApiApiServicesAppIntegralSearchMyIntegralDtlPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiIntegralDtoGetIntegralDtlInput**](WctApiIntegralDtoGetIntegralDtlInput.md)|  | 

### Return type

[**WctApiApplicationIntegralDtoIntegralDtlOutput**](WCT.Api.Application.Integral.Dto.IntegralDtlOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppIntegralSearchMyIntegralTaskPost**
> WctApiApplicationIntegralDtoMyIntegralTaskOutput ApiServicesAppIntegralSearchMyIntegralTaskPost(ctx, optional)
获取我的任务

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IntegralApiApiServicesAppIntegralSearchMyIntegralTaskPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegralApiApiServicesAppIntegralSearchMyIntegralTaskPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **optional.Int32**|  | 

### Return type

[**WctApiApplicationIntegralDtoMyIntegralTaskOutput**](WCT.Api.Application.Integral.Dto.MyIntegralTaskOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

