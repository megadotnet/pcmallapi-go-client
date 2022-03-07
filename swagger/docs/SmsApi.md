# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServicesAppSmsGetValidateCodeGet**](SmsApi.md#ApiServicesAppSmsGetValidateCodeGet) | **Get** /api/services/app/Sms/GetValidateCode | 获取图形验证码
[**ApiServicesAppSmsSendSMSCodePost**](SmsApi.md#ApiServicesAppSmsSendSMSCodePost) | **Post** /api/services/app/Sms/SendSMSCode | 发送(短信or邮件)验证码
[**ApiServicesAppSmsSendSMSMobileCodePost**](SmsApi.md#ApiServicesAppSmsSendSMSMobileCodePost) | **Post** /api/services/app/Sms/SendSMSMobileCode | 发送短信验证码  商户入驻使用  有效期为15分钟
[**ApiServicesAppSmsSmsCodeValidatePost**](SmsApi.md#ApiServicesAppSmsSmsCodeValidatePost) | **Post** /api/services/app/Sms/SmsCodeValidate | 校验(短信or邮件)验证码

# **ApiServicesAppSmsGetValidateCodeGet**
> string ApiServicesAppSmsGetValidateCodeGet(ctx, optional)
获取图形验证码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmsApiApiServicesAppSmsGetValidateCodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmsApiApiServicesAppSmsGetValidateCodeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageKey** | **optional.String**| 随机字符串 [需要保存在客户端用来校对图形验证码] | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppSmsSendSMSCodePost**
> map[string]string ApiServicesAppSmsSendSMSCodePost(ctx, optional)
发送(短信or邮件)验证码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmsApiApiServicesAppSmsSendSMSCodePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmsApiApiServicesAppSmsSendSMSCodePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationSmsDtoSendSmsDto**](WctApiApplicationSmsDtoSendSmsDto.md)|  | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppSmsSendSMSMobileCodePost**
> map[string]string ApiServicesAppSmsSendSMSMobileCodePost(ctx, optional)
发送短信验证码  商户入驻使用  有效期为15分钟

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmsApiApiServicesAppSmsSendSMSMobileCodePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmsApiApiServicesAppSmsSendSMSMobileCodePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationSmsDtoSendSmsDto**](WctApiApplicationSmsDtoSendSmsDto.md)|  | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiServicesAppSmsSmsCodeValidatePost**
> bool ApiServicesAppSmsSmsCodeValidatePost(ctx, optional)
校验(短信or邮件)验证码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmsApiApiServicesAppSmsSmsCodeValidatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmsApiApiServicesAppSmsSmsCodeValidatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationSmsDtoSmsCodeValidateDto**](WctApiApplicationSmsDtoSmsCodeValidateDto.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

