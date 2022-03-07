# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppRefundOrderCancelPost**](RefundOrderApi.md#ApiServicesAppRefundOrderCancelPost) | **Post** /api/services/app/RefundOrder/Cancel | 取消售后
[**ApiServicesAppRefundOrderConfirmReceiptPost**](RefundOrderApi.md#ApiServicesAppRefundOrderConfirmReceiptPost) | **Post** /api/services/app/RefundOrder/ConfirmReceipt | 换货确认收货
[**ApiServicesAppRefundOrderCreatePost**](RefundOrderApi.md#ApiServicesAppRefundOrderCreatePost) | **Post** /api/services/app/RefundOrder/Create | 申请售后
[**ApiServicesAppRefundOrderDetailPost**](RefundOrderApi.md#ApiServicesAppRefundOrderDetailPost) | **Post** /api/services/app/RefundOrder/Detail | 售后详情
[**ApiServicesAppRefundOrderListPost**](RefundOrderApi.md#ApiServicesAppRefundOrderListPost) | **Post** /api/services/app/RefundOrder/List | 售后订单列表
[**ApiServicesAppRefundOrderSubmitLogisticsPost**](RefundOrderApi.md#ApiServicesAppRefundOrderSubmitLogisticsPost) | **Post** /api/services/app/RefundOrder/SubmitLogistics | 提交物流信息
[**ApiServicesAppRefundOrderSubmitPlatformPost**](RefundOrderApi.md#ApiServicesAppRefundOrderSubmitPlatformPost) | **Post** /api/services/app/RefundOrder/SubmitPlatform | 平台介入

# **ApiServicesAppRefundOrderCancelPost**
> bool ApiServicesAppRefundOrderCancelPost(ctx, optional)
取消售后

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefundOrderApiApiServicesAppRefundOrderCancelPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundOrderApiApiServicesAppRefundOrderCancelPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationRefundDtoSubmitLogisticsDto**](WctApiApplicationRefundDtoSubmitLogisticsDto.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppRefundOrderConfirmReceiptPost**
> bool ApiServicesAppRefundOrderConfirmReceiptPost(ctx, optional)
换货确认收货

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefundOrderApiApiServicesAppRefundOrderConfirmReceiptPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundOrderApiApiServicesAppRefundOrderConfirmReceiptPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationRefundDtoSubmitLogisticsDto**](WctApiApplicationRefundDtoSubmitLogisticsDto.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppRefundOrderCreatePost**
> bool ApiServicesAppRefundOrderCreatePost(ctx, optional)
申请售后

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefundOrderApiApiServicesAppRefundOrderCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundOrderApiApiServicesAppRefundOrderCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationRefundDtoRefundOrderSubmit**](WctApiApplicationRefundDtoRefundOrderSubmit.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppRefundOrderDetailPost**
> WctApiApplicationRefundDtoRefundOrderDetail ApiServicesAppRefundOrderDetailPost(ctx, optional)
售后详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefundOrderApiApiServicesAppRefundOrderDetailPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundOrderApiApiServicesAppRefundOrderDetailPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationRefundDtoSubmitLogisticsDto**](WctApiApplicationRefundDtoSubmitLogisticsDto.md)|  | 

### Return type

[**WctApiApplicationRefundDtoRefundOrderDetail**](WCT.Api.Application.Refund.Dto.RefundOrderDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppRefundOrderListPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationRefundDtoRefundOrderDetailWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppRefundOrderListPost(ctx, optional)
售后订单列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefundOrderApiApiServicesAppRefundOrderListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundOrderApiApiServicesAppRefundOrderListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AbpApplicationServicesDtoPagedResultRequestDto**](AbpApplicationServicesDtoPagedResultRequestDto.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationRefundDtoRefundOrderDetailWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Refund.Dto.RefundOrderDetail, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppRefundOrderSubmitLogisticsPost**
> bool ApiServicesAppRefundOrderSubmitLogisticsPost(ctx, optional)
提交物流信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefundOrderApiApiServicesAppRefundOrderSubmitLogisticsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundOrderApiApiServicesAppRefundOrderSubmitLogisticsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationRefundDtoSubmitLogisticsDto**](WctApiApplicationRefundDtoSubmitLogisticsDto.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppRefundOrderSubmitPlatformPost**
> bool ApiServicesAppRefundOrderSubmitPlatformPost(ctx, optional)
平台介入

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefundOrderApiApiServicesAppRefundOrderSubmitPlatformPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundOrderApiApiServicesAppRefundOrderSubmitPlatformPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationRefundDtoSubmitLogisticsDto**](WctApiApplicationRefundDtoSubmitLogisticsDto.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

