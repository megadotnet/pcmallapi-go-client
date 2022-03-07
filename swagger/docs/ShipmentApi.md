# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppShipmentCalculateExpressPost**](ShipmentApi.md#ApiServicesAppShipmentCalculateExpressPost) | **Post** /api/services/app/Shipment/CalculateExpress | 计算运费
[**ApiServicesAppShipmentTrackExpressPost**](ShipmentApi.md#ApiServicesAppShipmentTrackExpressPost) | **Post** /api/services/app/Shipment/TrackExpress | 获取物流查询状态

# **ApiServicesAppShipmentCalculateExpressPost**
> WctApiShipmentsDtoResponseCalculateData ApiServicesAppShipmentCalculateExpressPost(ctx, optional)
计算运费

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShipmentApiApiServicesAppShipmentCalculateExpressPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShipmentApiApiServicesAppShipmentCalculateExpressPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiShipmentsDtoCalculateInput**](WctApiShipmentsDtoCalculateInput.md)|  | 

### Return type

[**WctApiShipmentsDtoResponseCalculateData**](WCT.Api.Shipments.Dto.ResponseCalculateData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppShipmentTrackExpressPost**
> string ApiServicesAppShipmentTrackExpressPost(ctx, optional)
获取物流查询状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShipmentApiApiServicesAppShipmentTrackExpressPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShipmentApiApiServicesAppShipmentTrackExpressPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiShipmentsDtoRequestTrackDto**](WctApiShipmentsDtoRequestTrackDto.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

