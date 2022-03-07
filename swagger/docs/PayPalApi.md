# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppPayPalCreatePaymentOrderPost**](PayPalApi.md#ApiServicesAppPayPalCreatePaymentOrderPost) | **Post** /api/services/app/PayPal/CreatePaymentOrder | 创建支付单 单位默认:USD  200:成功  则返回批准交易url,点此url登录批准交易,跳转url上取token值再调用支付接口
[**ApiServicesAppPayPalOrderPaymentPost**](PayPalApi.md#ApiServicesAppPayPalOrderPaymentPost) | **Post** /api/services/app/PayPal/OrderPayment | 支付-传支付单Id,成功返回支付单Id,非订单Id

# **ApiServicesAppPayPalCreatePaymentOrderPost**
> WctApiApplicationCommonResultDto ApiServicesAppPayPalCreatePaymentOrderPost(ctx, optional)
创建支付单 单位默认:USD  200:成功  则返回批准交易url,点此url登录批准交易,跳转url上取token值再调用支付接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PayPalApiApiServicesAppPayPalCreatePaymentOrderPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayPalApiApiServicesAppPayPalCreatePaymentOrderPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationPayPalDtoOrderRequestDto**](WctApiApplicationPayPalDtoOrderRequestDto.md)|  | 

### Return type

[**WctApiApplicationCommonResultDto**](WCT.Api.Application.Common.ResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppPayPalOrderPaymentPost**
> WctApiApplicationCommonResultDto ApiServicesAppPayPalOrderPaymentPost(ctx, optional)
支付-传支付单Id,成功返回支付单Id,非订单Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PayPalApiApiServicesAppPayPalOrderPaymentPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayPalApiApiServicesAppPayPalOrderPaymentPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **optional.String**| 支付单Id | 

### Return type

[**WctApiApplicationCommonResultDto**](WCT.Api.Application.Common.ResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

