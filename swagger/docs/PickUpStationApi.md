# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppPickUpStationCanConfirmOrderProductIdsGet**](PickUpStationApi.md#ApiServicesAppPickUpStationCanConfirmOrderProductIdsGet) | **Get** /api/services/app/PickUpStation/CanConfirmOrderProductIds | 根据商品id获取可以自提的店铺列表和下单状态
[**ApiServicesAppPickUpStationCheckPickUpStationMemberPost**](PickUpStationApi.md#ApiServicesAppPickUpStationCheckPickUpStationMemberPost) | **Post** /api/services/app/PickUpStation/CheckPickUpStationMember | 校验当前用户是否核销人员
[**ApiServicesAppPickUpStationGetPickUpStationListGet**](PickUpStationApi.md#ApiServicesAppPickUpStationGetPickUpStationListGet) | **Get** /api/services/app/PickUpStation/GetPickUpStationList | 获取自提点信息
[**ApiServicesAppPickUpStationGetSelfPickupByProductIdGet**](PickUpStationApi.md#ApiServicesAppPickUpStationGetSelfPickupByProductIdGet) | **Get** /api/services/app/PickUpStation/GetSelfPickupByProductId | 根据商品Id获取自提地址
[**ApiServicesAppPickUpStationGetSelfPickupByProductidsGet**](PickUpStationApi.md#ApiServicesAppPickUpStationGetSelfPickupByProductidsGet) | **Get** /api/services/app/PickUpStation/GetSelfPickupByProductids | 根据商品id获取可以自提的店铺列表
[**ApiServicesAppPickUpStationPickUpPost**](PickUpStationApi.md#ApiServicesAppPickUpStationPickUpPost) | **Post** /api/services/app/PickUpStation/PickUp | 自提操作
[**ApiServicesAppPickUpStationPostPickUpStationListPost**](PickUpStationApi.md#ApiServicesAppPickUpStationPostPickUpStationListPost) | **Post** /api/services/app/PickUpStation/PostPickUpStationList | 获取自提点信息(多商户)

# **ApiServicesAppPickUpStationCanConfirmOrderProductIdsGet**
> WctApiApplicationPickUpStationDtoPickUpStationExtDto ApiServicesAppPickUpStationCanConfirmOrderProductIdsGet(ctx, optional)
根据商品id获取可以自提的店铺列表和下单状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PickUpStationApiApiServicesAppPickUpStationCanConfirmOrderProductIdsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PickUpStationApiApiServicesAppPickUpStationCanConfirmOrderProductIdsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productIds** | **optional.String**|  | 

### Return type

[**WctApiApplicationPickUpStationDtoPickUpStationExtDto**](WCT.Api.Application.PickUpStation.Dto.PickUpStationExtDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppPickUpStationCheckPickUpStationMemberPost**
> bool ApiServicesAppPickUpStationCheckPickUpStationMemberPost(ctx, optional)
校验当前用户是否核销人员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PickUpStationApiApiServicesAppPickUpStationCheckPickUpStationMemberPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PickUpStationApiApiServicesAppPickUpStationCheckPickUpStationMemberPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationPickUpStationDtoPickUpStationDto**](WctApiApplicationPickUpStationDtoPickUpStationDto.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppPickUpStationGetPickUpStationListGet**
> []WctAdminPickUpStationPickUpStations ApiServicesAppPickUpStationGetPickUpStationListGet(ctx, optional)
获取自提点信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PickUpStationApiApiServicesAppPickUpStationGetPickUpStationListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PickUpStationApiApiServicesAppPickUpStationGetPickUpStationListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **optional.Int32**| 供应商Id | 
 **orderId** | **optional.Int64**| 订单Id | 

### Return type

[**[]WctAdminPickUpStationPickUpStations**](WCT.Admin.PickUpStation.PickUpStations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppPickUpStationGetSelfPickupByProductIdGet**
> []WctApiApplicationPickUpStationDtoPickUpStationsOutPut ApiServicesAppPickUpStationGetSelfPickupByProductIdGet(ctx, optional)
根据商品Id获取自提地址

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PickUpStationApiApiServicesAppPickUpStationGetSelfPickupByProductIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PickUpStationApiApiServicesAppPickUpStationGetSelfPickupByProductIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **optional.Int64**|  | 

### Return type

[**[]WctApiApplicationPickUpStationDtoPickUpStationsOutPut**](WCT.Api.Application.PickUpStation.Dto.PickUpStationsOutPut.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppPickUpStationGetSelfPickupByProductidsGet**
> []WctAdminPickUpStationPickUpStations ApiServicesAppPickUpStationGetSelfPickupByProductidsGet(ctx, optional)
根据商品id获取可以自提的店铺列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PickUpStationApiApiServicesAppPickUpStationGetSelfPickupByProductidsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PickUpStationApiApiServicesAppPickUpStationGetSelfPickupByProductidsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productIds** | **optional.String**|  | 

### Return type

[**[]WctAdminPickUpStationPickUpStations**](WCT.Admin.PickUpStation.PickUpStations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppPickUpStationPickUpPost**
> WctApiApplicationPickUpStationDtoPickUpResultDto ApiServicesAppPickUpStationPickUpPost(ctx, optional)
自提操作

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PickUpStationApiApiServicesAppPickUpStationPickUpPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PickUpStationApiApiServicesAppPickUpStationPickUpPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationPickUpStationDtoPickUpStationDto**](WctApiApplicationPickUpStationDtoPickUpStationDto.md)|  | 

### Return type

[**WctApiApplicationPickUpStationDtoPickUpResultDto**](WCT.Api.Application.PickUpStation.Dto.PickUpResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppPickUpStationPostPickUpStationListPost**
> []WctApiPickUpStationDtoPostPickUpStationsOutput ApiServicesAppPickUpStationPostPickUpStationListPost(ctx, optional)
获取自提点信息(多商户)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PickUpStationApiApiServicesAppPickUpStationPostPickUpStationListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PickUpStationApiApiServicesAppPickUpStationPostPickUpStationListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiPickUpStationDtoGetPickUpStationsInput**](WctApiPickUpStationDtoGetPickUpStationsInput.md)|  | 

### Return type

[**[]WctApiPickUpStationDtoPostPickUpStationsOutput**](WCT.Api.PickUpStation.Dto.PostPickUpStationsOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

