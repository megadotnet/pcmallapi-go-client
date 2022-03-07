# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppPayGainPaymentWayPost**](PayApi.md#ApiServicesAppPayGainPaymentWayPost) | **Post** /api/services/app/Pay/GainPaymentWay | 获取支付方式

# **ApiServicesAppPayGainPaymentWayPost**
> []WctApiApplicationPayPaymentDto ApiServicesAppPayGainPaymentWayPost(ctx, optional)
获取支付方式

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PayApiApiServicesAppPayGainPaymentWayPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayApiApiServicesAppPayGainPaymentWayPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationPayPayInput**](WctApiApplicationPayPayInput.md)|  | 

### Return type

[**[]WctApiApplicationPayPaymentDto**](WCT.Api.Application.Pay.PaymentDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

