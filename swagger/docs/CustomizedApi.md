# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppCustomizedHandlePost**](CustomizedApi.md#ApiServicesAppCustomizedHandlePost) | **Post** /api/services/app/Customized/Handle | 

# **ApiServicesAppCustomizedHandlePost**
> WctApiApplicationCustomizedDtoCustomizedResult ApiServicesAppCustomizedHandlePost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomizedApiApiServicesAppCustomizedHandlePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomizedApiApiServicesAppCustomizedHandlePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationCustomizedDtoCustomizedParam**](WctApiApplicationCustomizedDtoCustomizedParam.md)|  | 

### Return type

[**WctApiApplicationCustomizedDtoCustomizedResult**](WCT.Api.Application.Customized.Dto.CustomizedResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

