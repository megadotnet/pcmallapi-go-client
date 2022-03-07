# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppAppUpgradeGetUpgradeInfoGet**](AppUpgradeApi.md#ApiServicesAppAppUpgradeGetUpgradeInfoGet) | **Get** /api/services/app/AppUpgrade/GetUpgradeInfo | 获取APP升级信息

# **ApiServicesAppAppUpgradeGetUpgradeInfoGet**
> WctApiApplicationAppUpgradesDtoAppUpgradeDto ApiServicesAppAppUpgradeGetUpgradeInfoGet(ctx, optional)
获取APP升级信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppUpgradeApiApiServicesAppAppUpgradeGetUpgradeInfoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppUpgradeApiApiServicesAppAppUpgradeGetUpgradeInfoGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **optional.Int32**| 平台类型 1-IOS 2-Android | 
 **version** | **optional.Int32**| 版本号，数值型 5.0.0&gt;500 | 

### Return type

[**WctApiApplicationAppUpgradesDtoAppUpgradeDto**](WCT.Api.Application.AppUpgrades.Dto.AppUpgradeDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

