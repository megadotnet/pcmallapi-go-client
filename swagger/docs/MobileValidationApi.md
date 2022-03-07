# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppMobileValidationValidationPost**](MobileValidationApi.md#ApiServicesAppMobileValidationValidationPost) | **Post** /api/services/app/MobileValidation/Validation | 校验手机号

# **ApiServicesAppMobileValidationValidationPost**
> WctAdminValidationMobileAscription ApiServicesAppMobileValidationValidationPost(ctx, optional)
校验手机号

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MobileValidationApiApiServicesAppMobileValidationValidationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MobileValidationApiApiServicesAppMobileValidationValidationPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationValidationDtoMobileValidationDto**](WctApiApplicationValidationDtoMobileValidationDto.md)|  | 

### Return type

[**WctAdminValidationMobileAscription**](WCT.Admin.Validation.MobileAscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

