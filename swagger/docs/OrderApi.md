# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppOrderCheckPayStatusPost**](OrderApi.md#ApiServicesAppOrderCheckPayStatusPost) | **Post** /api/services/app/Order/CheckPayStatus | 校验支付状态
[**ApiServicesAppOrderCloseOrderPost**](OrderApi.md#ApiServicesAppOrderCloseOrderPost) | **Post** /api/services/app/Order/CloseOrder | 关闭订单
[**ApiServicesAppOrderCompletedActionPost**](OrderApi.md#ApiServicesAppOrderCompletedActionPost) | **Post** /api/services/app/Order/CompletedAction | 订单完成统一处理
[**ApiServicesAppOrderCompletedOrderPost**](OrderApi.md#ApiServicesAppOrderCompletedOrderPost) | **Post** /api/services/app/Order/CompletedOrder | 订单收货
[**ApiServicesAppOrderCompletedPost**](OrderApi.md#ApiServicesAppOrderCompletedPost) | **Post** /api/services/app/Order/Completed | 
[**ApiServicesAppOrderConfirmOrderPost**](OrderApi.md#ApiServicesAppOrderConfirmOrderPost) | **Post** /api/services/app/Order/ConfirmOrder | 确认订单
[**ApiServicesAppOrderCreateOrderPost**](OrderApi.md#ApiServicesAppOrderCreateOrderPost) | **Post** /api/services/app/Order/CreateOrder | 创建订单
[**ApiServicesAppOrderCreatePaymentOrderPost**](OrderApi.md#ApiServicesAppOrderCreatePaymentOrderPost) | **Post** /api/services/app/Order/CreatePaymentOrder | 创建支付单
[**ApiServicesAppOrderDeletePost**](OrderApi.md#ApiServicesAppOrderDeletePost) | **Post** /api/services/app/Order/Delete | 删除订单
[**ApiServicesAppOrderDeliveryReminderGet**](OrderApi.md#ApiServicesAppOrderDeliveryReminderGet) | **Get** /api/services/app/Order/DeliveryReminder | 提醒卖家发货
[**ApiServicesAppOrderDetailPost**](OrderApi.md#ApiServicesAppOrderDetailPost) | **Post** /api/services/app/Order/Detail | 订单详情
[**ApiServicesAppOrderGainOrderAmountByUserIdPost**](OrderApi.md#ApiServicesAppOrderGainOrderAmountByUserIdPost) | **Post** /api/services/app/Order/GainOrderAmountByUserId | 
[**ApiServicesAppOrderGetFreightFeeAsyncPost**](OrderApi.md#ApiServicesAppOrderGetFreightFeeAsyncPost) | **Post** /api/services/app/Order/GetFreightFeeAsync | 
[**ApiServicesAppOrderGetOrderCountGet**](OrderApi.md#ApiServicesAppOrderGetOrderCountGet) | **Get** /api/services/app/Order/GetOrderCount | 获取用户订单数 （新 返回为对象）
[**ApiServicesAppOrderGetOrderPayRecordGet**](OrderApi.md#ApiServicesAppOrderGetOrderPayRecordGet) | **Get** /api/services/app/Order/GetOrderPayRecord | 获取订单商户支付信息
[**ApiServicesAppOrderGetSadadUserMobileGet**](OrderApi.md#ApiServicesAppOrderGetSadadUserMobileGet) | **Get** /api/services/app/Order/GetSadadUserMobile | 获取SADAD/Edfali用户手机号
[**ApiServicesAppOrderInternalPaymentCallbackPost**](OrderApi.md#ApiServicesAppOrderInternalPaymentCallbackPost) | **Post** /api/services/app/Order/InternalPaymentCallback | 内部支付-确认
[**ApiServicesAppOrderLogisticsSearchPost**](OrderApi.md#ApiServicesAppOrderLogisticsSearchPost) | **Post** /api/services/app/Order/LogisticsSearch | 查询物流信息
[**ApiServicesAppOrderModifyDeliveryAddressPost**](OrderApi.md#ApiServicesAppOrderModifyDeliveryAddressPost) | **Post** /api/services/app/Order/ModifyDeliveryAddress | 修改收货地址（待付款订单）
[**ApiServicesAppOrderNewCustomerPost**](OrderApi.md#ApiServicesAppOrderNewCustomerPost) | **Post** /api/services/app/Order/NewCustomer | 判断是否新用户（不存在订单信息）
[**ApiServicesAppOrderOrderCountPost**](OrderApi.md#ApiServicesAppOrderOrderCountPost) | **Post** /api/services/app/Order/OrderCount | 获取用户订单数
[**ApiServicesAppOrderPayOrderPost**](OrderApi.md#ApiServicesAppOrderPayOrderPost) | **Post** /api/services/app/Order/PayOrder | 待付款-创建支付单
[**ApiServicesAppOrderPendingDetailPost**](OrderApi.md#ApiServicesAppOrderPendingDetailPost) | **Post** /api/services/app/Order/PendingDetail | 订单详情(待支付)
[**ApiServicesAppOrderReceiptPreviewGet**](OrderApi.md#ApiServicesAppOrderReceiptPreviewGet) | **Get** /api/services/app/Order/ReceiptPreview | 小票预览接口
[**ApiServicesAppOrderSearchPendingOrderPost**](OrderApi.md#ApiServicesAppOrderSearchPendingOrderPost) | **Post** /api/services/app/Order/SearchPendingOrder | 获取待支付订单信息
[**ApiServicesAppOrderSearchPost**](OrderApi.md#ApiServicesAppOrderSearchPost) | **Post** /api/services/app/Order/Search | 查询订单信息
[**ApiServicesAppOrderSelectPaymentCallbackPost**](OrderApi.md#ApiServicesAppOrderSelectPaymentCallbackPost) | **Post** /api/services/app/Order/SelectPaymentCallback | SelectSystem支付回调操作
[**ApiServicesAppOrderTriggerIntegralPost**](OrderApi.md#ApiServicesAppOrderTriggerIntegralPost) | **Post** /api/services/app/Order/TriggerIntegral | 触发积分

# **ApiServicesAppOrderCheckPayStatusPost**
> bool ApiServicesAppOrderCheckPayStatusPost(ctx, optional)
校验支付状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderCheckPayStatusPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderCheckPayStatusPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e**](AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderCloseOrderPost**
> bool ApiServicesAppOrderCloseOrderPost(ctx, optional)
关闭订单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderCloseOrderPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderCloseOrderPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e**](AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderCompletedActionPost**
> bool ApiServicesAppOrderCompletedActionPost(ctx, optional)
订单完成统一处理

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderCompletedActionPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderCompletedActionPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctEntitiesOrderMst**](WctEntitiesOrderMst.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderCompletedOrderPost**
> SystemValueTuple2SystemBooleanSystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798eSystemInt32SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e ApiServicesAppOrderCompletedOrderPost(ctx, optional)
订单收货

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderCompletedOrderPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderCompletedOrderPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiOrderDtoCompletedOrderDto**](WctApiOrderDtoCompletedOrderDto.md)|  | 

### Return type

[**SystemValueTuple2SystemBooleanSystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798eSystemInt32SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e**](System.ValueTuple&#x60;2[[System.Boolean, System.Private.CoreLib, Version&#x3D;4.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;7cec85d7bea7798e],[System.Int32, System.Private.CoreLib, Version&#x3D;4.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;7cec85d7bea7798e]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderCompletedPost**
> bool ApiServicesAppOrderCompletedPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderCompletedPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderCompletedPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctEntitiesOrderMst**](WctEntitiesOrderMst.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderConfirmOrderPost**
> WctApiOrderDtoBuildOrderConfirmOrderOutput ApiServicesAppOrderConfirmOrderPost(ctx, optional)
确认订单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderConfirmOrderPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderConfirmOrderPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiOrderDtoBuildOrderConfirmOrderDto**](WctApiOrderDtoBuildOrderConfirmOrderDto.md)|  | 

### Return type

[**WctApiOrderDtoBuildOrderConfirmOrderOutput**](WCT.Api.Order.Dto.BuildOrder.ConfirmOrderOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderCreateOrderPost**
> WctApiOrderDtoBuildOrderCreateOrderOutput ApiServicesAppOrderCreateOrderPost(ctx, optional)
创建订单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderCreateOrderPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderCreateOrderPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiOrderDtoBuildOrderBuildOrderDto**](WctApiOrderDtoBuildOrderBuildOrderDto.md)|  | 

### Return type

[**WctApiOrderDtoBuildOrderCreateOrderOutput**](WCT.Api.Order.Dto.BuildOrder.CreateOrderOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderCreatePaymentOrderPost**
> ApiServicesAppOrderCreatePaymentOrderPost(ctx, optional)
创建支付单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderCreatePaymentOrderPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderCreatePaymentOrderPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationOrderDtoPaymentOrderDto**](WctApiApplicationOrderDtoPaymentOrderDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderDeletePost**
> bool ApiServicesAppOrderDeletePost(ctx, optional)
删除订单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderDeletePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e**](AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderDeliveryReminderGet**
> bool ApiServicesAppOrderDeliveryReminderGet(ctx, optional)
提醒卖家发货

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderDeliveryReminderGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderDeliveryReminderGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **optional.Int64**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderDetailPost**
> WctApiApplicationOrderDtoOrderPortalDetailDto ApiServicesAppOrderDetailPost(ctx, optional)
订单详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderDetailPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderDetailPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e**](AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e.md)|  | 

### Return type

[**WctApiApplicationOrderDtoOrderPortalDetailDto**](WCT.Api.Application.Order.Dto.OrderPortalDetailDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderGainOrderAmountByUserIdPost**
> []WctApiApplicationOrderDtoOrderDto ApiServicesAppOrderGainOrderAmountByUserIdPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderGainOrderAmountByUserIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderGainOrderAmountByUserIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationOrderDtoTimeInput**](WctApiApplicationOrderDtoTimeInput.md)|  | 

### Return type

[**[]WctApiApplicationOrderDtoOrderDto**](WCT.Api.Application.Order.Dto.OrderDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderGetFreightFeeAsyncPost**
> WctApiApplicationOrderDtoProductFreightFee ApiServicesAppOrderGetFreightFeeAsyncPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderGetFreightFeeAsyncPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderGetFreightFeeAsyncPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationOrderDtoProductDetailFrightFeeInput**](WctApiApplicationOrderDtoProductDetailFrightFeeInput.md)|  | 

### Return type

[**WctApiApplicationOrderDtoProductFreightFee**](WCT.Api.Application.Order.Dto.ProductFreightFee.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderGetOrderCountGet**
> WctApiApplicationOrderDtoOrderCountOutput ApiServicesAppOrderGetOrderCountGet(ctx, )
获取用户订单数 （新 返回为对象）

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WctApiApplicationOrderDtoOrderCountOutput**](WCT.Api.Application.Order.Dto.OrderCountOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderGetOrderPayRecordGet**
> []WctApiOrderDtoTriggerIntegralDto ApiServicesAppOrderGetOrderPayRecordGet(ctx, optional)
获取订单商户支付信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderGetOrderPayRecordGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderGetOrderPayRecordGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderList** | [**optional.Interface of []WctEntitiesOrderMst**](WctEntitiesOrderMst.md)|  | 

### Return type

[**[]WctApiOrderDtoTriggerIntegralDto**](WCT.Api.Order.Dto.TriggerIntegralDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderGetSadadUserMobileGet**
> string ApiServicesAppOrderGetSadadUserMobileGet(ctx, optional)
获取SADAD/Edfali用户手机号

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderGetSadadUserMobileGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderGetSadadUserMobileGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payment** | [**optional.Interface of WctEntitiesPayment**](.md)| 判断记录支付类型  SADAD 为120  Edfali 为130 | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderInternalPaymentCallbackPost**
> ApiServicesAppOrderInternalPaymentCallbackPost(ctx, optional)
内部支付-确认

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderInternalPaymentCallbackPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderInternalPaymentCallbackPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiOrderDtoInternalPaymentConfirmTransactionInputDto**](WctApiOrderDtoInternalPaymentConfirmTransactionInputDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderLogisticsSearchPost**
> []WctApiApplicationOrderDtoOrderExpressDto ApiServicesAppOrderLogisticsSearchPost(ctx, optional)
查询物流信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderLogisticsSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderLogisticsSearchPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e**](AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e.md)|  | 

### Return type

[**[]WctApiApplicationOrderDtoOrderExpressDto**](WCT.Api.Application.Order.Dto.OrderExpressDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderModifyDeliveryAddressPost**
> bool ApiServicesAppOrderModifyDeliveryAddressPost(ctx, optional)
修改收货地址（待付款订单）

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderModifyDeliveryAddressPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderModifyDeliveryAddressPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiOrderDtoUpdateOrderAddressInput**](WctApiOrderDtoUpdateOrderAddressInput.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderNewCustomerPost**
> bool ApiServicesAppOrderNewCustomerPost(ctx, )
判断是否新用户（不存在订单信息）

### Required Parameters
This endpoint does not need any parameter.

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderOrderCountPost**
> []int32 ApiServicesAppOrderOrderCountPost(ctx, )
获取用户订单数

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderPayOrderPost**
> ApiServicesAppOrderPayOrderPost(ctx, optional)
待付款-创建支付单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderPayOrderPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderPayOrderPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationOrderDtoPayDto**](WctApiApplicationOrderDtoPayDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderPendingDetailPost**
> []WctApiApplicationOrderDtoOrderPortalDetailDto ApiServicesAppOrderPendingDetailPost(ctx, optional)
订单详情(待支付)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderPendingDetailPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderPendingDetailPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e**](AbpApplicationServicesDtoEntityDto1SystemInt64SystemPrivateCoreLibVersion4000CultureneutralPublicKeyToken7cec85d7bea7798e.md)|  | 

### Return type

[**[]WctApiApplicationOrderDtoOrderPortalDetailDto**](WCT.Api.Application.Order.Dto.OrderPortalDetailDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderReceiptPreviewGet**
> WctApiOrderDtoReceiptPreviewDto ApiServicesAppOrderReceiptPreviewGet(ctx, optional)
小票预览接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderReceiptPreviewGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderReceiptPreviewGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **optional.Int64**|  | 

### Return type

[**WctApiOrderDtoReceiptPreviewDto**](WCT.Api.Order.Dto.ReceiptPreviewDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderSearchPendingOrderPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationOrderDtoOrderPortalDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppOrderSearchPendingOrderPost(ctx, optional)
获取待支付订单信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderSearchPendingOrderPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderSearchPendingOrderPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationOrderDtoOrderPortalSearch**](WctApiApplicationOrderDtoOrderPortalSearch.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationOrderDtoOrderPortalDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Order.Dto.OrderPortalDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderSearchPost**
> AbpApplicationServicesDtoPagedResultDto1WctApiApplicationOrderDtoOrderPortalDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppOrderSearchPost(ctx, optional)
查询订单信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderSearchPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationOrderDtoOrderPortalSearch**](WctApiApplicationOrderDtoOrderPortalSearch.md)|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiApplicationOrderDtoOrderPortalDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Application.Order.Dto.OrderPortalDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderSelectPaymentCallbackPost**
> WctApiOrderDtoSelectSystemSelectSystemPayCallbackReturn ApiServicesAppOrderSelectPaymentCallbackPost(ctx, optional)
SelectSystem支付回调操作

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderSelectPaymentCallbackPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderSelectPaymentCallbackPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiOrderDtoSelectSystemSelectSystemPayCallbackInfo**](WctApiOrderDtoSelectSystemSelectSystemPayCallbackInfo.md)|  | 

### Return type

[**WctApiOrderDtoSelectSystemSelectSystemPayCallbackReturn**](WCT.Api.Order.Dto.SelectSystem.SelectSystemPayCallbackReturn.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppOrderTriggerIntegralPost**
> ApiServicesAppOrderTriggerIntegralPost(ctx, optional)
触发积分

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiApiServicesAppOrderTriggerIntegralPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiApiServicesAppOrderTriggerIntegralPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []WctApiOrderDtoTriggerIntegralDto**](WCT.Api.Order.Dto.TriggerIntegralDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

