# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppTenantGetBaseInfoGet**](TenantApi.md#ApiServicesAppTenantGetBaseInfoGet) | **Get** /api/services/app/Tenant/GetBaseInfo | 获取店铺基本信息

# **ApiServicesAppTenantGetBaseInfoGet**
> WctApiCoreTenantsDtosTenantDto ApiServicesAppTenantGetBaseInfoGet(ctx, optional)
获取店铺基本信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenantApiApiServicesAppTenantGetBaseInfoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantApiApiServicesAppTenantGetBaseInfoGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **optional.Int32**|  | 

### Return type

[**WctApiCoreTenantsDtosTenantDto**](WCT.Api.Core.Tenants.Dtos.TenantDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

