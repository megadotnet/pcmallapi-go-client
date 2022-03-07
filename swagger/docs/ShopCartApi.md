# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppShopCartDeletePost**](ShopCartApi.md#ApiServicesAppShopCartDeletePost) | **Post** /api/services/app/ShopCart/Delete | 删除购物车商品
[**ApiServicesAppShopCartGetCountGet**](ShopCartApi.md#ApiServicesAppShopCartGetCountGet) | **Get** /api/services/app/ShopCart/GetCount | 获取购物车商品全部数量
[**ApiServicesAppShopCartGetGet**](ShopCartApi.md#ApiServicesAppShopCartGetGet) | **Get** /api/services/app/ShopCart/Get | 获取购物车商品
[**ApiServicesAppShopCartUpdatePost**](ShopCartApi.md#ApiServicesAppShopCartUpdatePost) | **Post** /api/services/app/ShopCart/Update | 更新购物车商品

# **ApiServicesAppShopCartDeletePost**
> ApiServicesAppShopCartDeletePost(ctx, optional)
删除购物车商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShopCartApiApiServicesAppShopCartDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShopCartApiApiServicesAppShopCartDeletePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []WctApiApplicationShopCartDtoDeleteShopItemInput**](WCT.Api.Application.ShopCart.Dto.DeleteShopItemInput.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppShopCartGetCountGet**
> int32 ApiServicesAppShopCartGetCountGet(ctx, )
获取购物车商品全部数量

### Required Parameters
This endpoint does not need any parameter.

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppShopCartGetGet**
> []WctApiApplicationShopCartDtoShopCartDto ApiServicesAppShopCartGetGet(ctx, )
获取购物车商品

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]WctApiApplicationShopCartDtoShopCartDto**](WCT.Api.Application.ShopCart.Dto.ShopCartDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppShopCartUpdatePost**
> int32 ApiServicesAppShopCartUpdatePost(ctx, optional)
更新购物车商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShopCartApiApiServicesAppShopCartUpdatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShopCartApiApiServicesAppShopCartUpdatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationShopCartDtoUpdateShopItemInput**](WctApiApplicationShopCartDtoUpdateShopItemInput.md)|  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

