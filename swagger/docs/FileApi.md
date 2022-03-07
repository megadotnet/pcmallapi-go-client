# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiFileDownloadInvoiceGet**](FileApi.md#ApiFileDownloadInvoiceGet) | **Get** /api/File/DownloadInvoice | 
[**ApiFileShareFacebookGet**](FileApi.md#ApiFileShareFacebookGet) | **Get** /api/File/ShareFacebook | 

# **ApiFileDownloadInvoiceGet**
> ApiFileDownloadInvoiceGet(ctx, orderId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int64**|  | 
 **optional** | ***FileApiApiFileDownloadInvoiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileApiApiFileDownloadInvoiceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiFileShareFacebookGet**
> ApiFileShareFacebookGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileApiApiFileShareFacebookGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileApiApiFileShareFacebookGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **optional.Int64**|  | 
 **skuId** | **optional.Int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

