# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppLogisticsGetActiveCompanyListGet**](LogisticsApi.md#ApiServicesAppLogisticsGetActiveCompanyListGet) | **Get** /api/services/app/Logistics/GetActiveCompanyList | 获取开启的物流公司
[**ApiServicesAppLogisticsGetLogicsIdsGet**](LogisticsApi.md#ApiServicesAppLogisticsGetLogicsIdsGet) | **Get** /api/services/app/Logistics/GetLogicsIds | 获取已对接物流公司Id,多个以逗号隔开
[**ApiServicesAppLogisticsGetPageListGet**](LogisticsApi.md#ApiServicesAppLogisticsGetPageListGet) | **Get** /api/services/app/Logistics/GetPageList | 物流公司列表 分页
[**ApiServicesAppLogisticsLogicsInfoSearchPost**](LogisticsApi.md#ApiServicesAppLogisticsLogicsInfoSearchPost) | **Post** /api/services/app/Logistics/LogicsInfoSearch | 第三方物流跟踪
[**ApiServicesAppLogisticsPostCompanyList2Post**](LogisticsApi.md#ApiServicesAppLogisticsPostCompanyList2Post) | **Post** /api/services/app/Logistics/PostCompanyList2 | 获取物流公司(包含运费)
[**ApiServicesAppLogisticsPostCompanyListPost**](LogisticsApi.md#ApiServicesAppLogisticsPostCompanyListPost) | **Post** /api/services/app/Logistics/PostCompanyList | 获取物流公司(包含运费)
[**ApiServicesAppLogisticsPostCompanyListStockPost**](LogisticsApi.md#ApiServicesAppLogisticsPostCompanyListStockPost) | **Post** /api/services/app/Logistics/PostCompanyListStock | 获取商品详情物流公司的运费,只取第一个信息
[**ApiServicesAppLogisticsPostCompanyfreightPost**](LogisticsApi.md#ApiServicesAppLogisticsPostCompanyfreightPost) | **Post** /api/services/app/Logistics/PostCompanyfreight | 获取大物流公司运费  VANEX
[**ApiServicesAppLogisticsPostPost**](LogisticsApi.md#ApiServicesAppLogisticsPostPost) | **Post** /api/services/app/Logistics/Post | 
[**ApiServicesAppLogisticsSearchPost**](LogisticsApi.md#ApiServicesAppLogisticsSearchPost) | **Post** /api/services/app/Logistics/Search | 
[**ApiServicesAppLogisticsShipmentCreateTestPost**](LogisticsApi.md#ApiServicesAppLogisticsShipmentCreateTestPost) | **Post** /api/services/app/Logistics/ShipmentCreateTest | 
[**ApiServicesAppLogisticsShipmentTestPost**](LogisticsApi.md#ApiServicesAppLogisticsShipmentTestPost) | **Post** /api/services/app/Logistics/ShipmentTest | 
[**ApiServicesAppLogisticsShipmentTrackTestPost**](LogisticsApi.md#ApiServicesAppLogisticsShipmentTrackTestPost) | **Post** /api/services/app/Logistics/ShipmentTrackTest | 

# **ApiServicesAppLogisticsGetActiveCompanyListGet**
> []WctApiLogisticsDtoReturnLogisticsCompanyDto ApiServicesAppLogisticsGetActiveCompanyListGet(ctx, optional)
获取开启的物流公司

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsGetActiveCompanyListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsGetActiveCompanyListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **optional.Int32**| tenantId | 
 **provinceCode** | **optional.String**| provinceCode | 
 **cityCode** | **optional.String**| cityCode | 
 **notlogistics** | **optional.Bool**| notlogistics | [default to false]

### Return type

[**[]WctApiLogisticsDtoReturnLogisticsCompanyDto**](WCT.Api.Logistics.Dto.ReturnLogisticsCompanyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsGetLogicsIdsGet**
> string ApiServicesAppLogisticsGetLogicsIdsGet(ctx, )
获取已对接物流公司Id,多个以逗号隔开

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsGetPageListGet**
> AbpApplicationServicesDtoPagedResultDto1WctApiLogisticsDtoLogisticsPageDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppLogisticsGetPageListGet(ctx, optional)
物流公司列表 分页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsGetPageListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsGetPageListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| 名称 | 
 **code** | **optional.String**| 编码 | 
 **contact** | **optional.String**| 联系人 | 
 **phoneNumber** | **optional.String**| 手机号 | 
 **pickUpProvince** | **optional.String**| 取件区域-省 | 
 **pickUpProvinceCode** | **optional.String**| 取件区域-省代码 | 
 **pickUpCity** | **optional.String**| 取件区域-市 | 
 **pickUpCityCode** | **optional.String**| 取件区域-市代码 | 
 **pickUpArea** | **optional.String**| 取件区域-区 | 
 **pickUpAreaCode** | **optional.String**| 取件区域-区代码 | 
 **isActive** | **optional.Bool**| 是否启用 | 
 **creationBeginTime** | **optional.Time**| 创建开始时间 | 
 **creationEndTime** | **optional.Time**| 创建结束时间 | 
 **tenantId** | **optional.Int32**| 商户 | 
 **skipCount** | **optional.Int32**|  | 
 **maxResultCount** | **optional.Int32**|  | 

### Return type

[**AbpApplicationServicesDtoPagedResultDto1WctApiLogisticsDtoLogisticsPageDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.PagedResultDto&#x60;1[[WCT.Api.Logistics.Dto.LogisticsPageDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsLogicsInfoSearchPost**
> WctApiLogisticsDtoShipmnetIntegrationShipmentResponse ApiServicesAppLogisticsLogicsInfoSearchPost(ctx, optional)
第三方物流跟踪

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsLogicsInfoSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsLogicsInfoSearchPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiLogisticsDtoShipmnetIntegrationLogisticsInputDto**](WctApiLogisticsDtoShipmnetIntegrationLogisticsInputDto.md)|  | 

### Return type

[**WctApiLogisticsDtoShipmnetIntegrationShipmentResponse**](WCT.Api.Logistics.Dto.ShipmnetIntegration.ShipmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsPostCompanyList2Post**
> []WctApiApplicationLogisticsDtoLogisticsCompanyDto ApiServicesAppLogisticsPostCompanyList2Post(ctx, optional)
获取物流公司(包含运费)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsPostCompanyList2PostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsPostCompanyList2PostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiLogisticsDtoGetCompanyInput**](WctApiLogisticsDtoGetCompanyInput.md)|  | 

### Return type

[**[]WctApiApplicationLogisticsDtoLogisticsCompanyDto**](WCT.Api.Application.Logistics.Dto.LogisticsCompanyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsPostCompanyListPost**
> WctApiLogisticsDtoLogisticsCompanyOutput ApiServicesAppLogisticsPostCompanyListPost(ctx, optional)
获取物流公司(包含运费)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsPostCompanyListPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsPostCompanyListPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiLogisticsDtoCreateCompanyInput**](WctApiLogisticsDtoCreateCompanyInput.md)|  | 

### Return type

[**WctApiLogisticsDtoLogisticsCompanyOutput**](WCT.Api.Logistics.Dto.LogisticsCompanyOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsPostCompanyListStockPost**
> []WctApiApplicationLogisticsDtoLogisticsCompanyDto ApiServicesAppLogisticsPostCompanyListStockPost(ctx, optional)
获取商品详情物流公司的运费,只取第一个信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsPostCompanyListStockPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsPostCompanyListStockPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiLogisticsDtoGetCompanyInput**](WctApiLogisticsDtoGetCompanyInput.md)|  | 

### Return type

[**[]WctApiApplicationLogisticsDtoLogisticsCompanyDto**](WCT.Api.Application.Logistics.Dto.LogisticsCompanyDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsPostCompanyfreightPost**
> float64 ApiServicesAppLogisticsPostCompanyfreightPost(ctx, optional)
获取大物流公司运费  VANEX

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsPostCompanyfreightPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsPostCompanyfreightPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiLogisticsDtoCreateCompanyVanexInput**](WctApiLogisticsDtoCreateCompanyVanexInput.md)|  | 

### Return type

**float64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsPostPost**
> string ApiServicesAppLogisticsPostPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsPostPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SystemTextEncoding**](SystemTextEncoding.md)|  | 
 **url** | **optional.**|  | 
 **data** | **optional.**|  | 
 **type_** | **optional.**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsSearchPost**
> ApiServicesAppLogisticsSearchPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsSearchPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationLogisticsDtoLogisticsSearch**](WctApiApplicationLogisticsDtoLogisticsSearch.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsShipmentCreateTestPost**
> ApiServicesAppLogisticsShipmentCreateTestPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsShipmentCreateTestPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsShipmentCreateTestPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiLogisticsDtoShipmnetIntegrationCreateShipmentInputDto**](WctApiLogisticsDtoShipmnetIntegrationCreateShipmentInputDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsShipmentTestPost**
> ApiServicesAppLogisticsShipmentTestPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsShipmentTestPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsShipmentTestPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiLogisticsDtoShipmnetIntegrationCreateShipmentRequestDto**](WctApiLogisticsDtoShipmnetIntegrationCreateShipmentRequestDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppLogisticsShipmentTrackTestPost**
> ApiServicesAppLogisticsShipmentTrackTestPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogisticsApiApiServicesAppLogisticsShipmentTrackTestPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogisticsApiApiServicesAppLogisticsShipmentTrackTestPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiLogisticsDtoShipmnetIntegrationShippingTrackInputDto**](WctApiLogisticsDtoShipmnetIntegrationShippingTrackInputDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

