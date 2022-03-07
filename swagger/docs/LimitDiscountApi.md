# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppLimitDiscountGetAllGet**](LimitDiscountApi.md#ApiServicesAppLimitDiscountGetAllGet) | **Get** /api/services/app/LimitDiscount/GetAll | 

# **ApiServicesAppLimitDiscountGetAllGet**
> AbpApplicationServicesDtoListResultDto1WctApiApplicationDtoLimitDiscountProductDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull ApiServicesAppLimitDiscountGetAllGet(ctx, productIds, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productIds** | **string**|  | 
 **optional** | ***LimitDiscountApiApiServicesAppLimitDiscountGetAllGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LimitDiscountApiApiServicesAppLimitDiscountGetAllGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activityId** | **optional.Int32**|  | 

### Return type

[**AbpApplicationServicesDtoListResultDto1WctApiApplicationDtoLimitDiscountProductDtoWctApiApplicationVersion1000CultureneutralPublicKeyTokennull**](Abp.Application.Services.Dto.ListResultDto&#x60;1[[WCT.Api.Application.Dto.LimitDiscountProductDto, WCT.Api.Application, Version&#x3D;1.0.0.0, Culture&#x3D;neutral, PublicKeyToken&#x3D;null]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

