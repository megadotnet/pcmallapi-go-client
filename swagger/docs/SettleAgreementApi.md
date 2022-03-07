# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppSettleAgreementQuerySettleAgreesPost**](SettleAgreementApi.md#ApiServicesAppSettleAgreementQuerySettleAgreesPost) | **Post** /api/services/app/SettleAgreement/QuerySettleAgrees | 获取协议地址

# **ApiServicesAppSettleAgreementQuerySettleAgreesPost**
> string ApiServicesAppSettleAgreementQuerySettleAgreesPost(ctx, optional)
获取协议地址

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettleAgreementApiApiServicesAppSettleAgreementQuerySettleAgreesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettleAgreementApiApiServicesAppSettleAgreementQuerySettleAgreesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiMerchantsEnteringDtoSettleAgreeInputDto**](WctApiMerchantsEnteringDtoSettleAgreeInputDto.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

