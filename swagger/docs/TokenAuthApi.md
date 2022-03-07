# {{classname}}

All URIs are relative to *https://lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTokenAuthAuthenticatePost**](TokenAuthApi.md#ApiTokenAuthAuthenticatePost) | **Post** /api/TokenAuth/Authenticate | 
[**ApiTokenAuthBindPhoneOrEmailPost**](TokenAuthApi.md#ApiTokenAuthBindPhoneOrEmailPost) | **Post** /api/TokenAuth/BindPhoneOrEmail | 
[**ApiTokenAuthChangeAccountPost**](TokenAuthApi.md#ApiTokenAuthChangeAccountPost) | **Post** /api/TokenAuth/ChangeAccount | 
[**ApiTokenAuthExternalAuthenticatePost**](TokenAuthApi.md#ApiTokenAuthExternalAuthenticatePost) | **Post** /api/TokenAuth/ExternalAuthenticate | 
[**ApiTokenAuthForgetPasswordPost**](TokenAuthApi.md#ApiTokenAuthForgetPasswordPost) | **Post** /api/TokenAuth/ForgetPassword | 
[**ApiTokenAuthLogOutGet**](TokenAuthApi.md#ApiTokenAuthLogOutGet) | **Get** /api/TokenAuth/LogOut | 
[**ApiTokenAuthManualSendNotificePost**](TokenAuthApi.md#ApiTokenAuthManualSendNotificePost) | **Post** /api/TokenAuth/ManualSendNotifice | 
[**ApiTokenAuthModifyPasswordPost**](TokenAuthApi.md#ApiTokenAuthModifyPasswordPost) | **Post** /api/TokenAuth/ModifyPassword | 
[**ApiTokenAuthOldPwdVerifyPost**](TokenAuthApi.md#ApiTokenAuthOldPwdVerifyPost) | **Post** /api/TokenAuth/OldPwdVerify | 
[**ApiTokenAuthRefreshPost**](TokenAuthApi.md#ApiTokenAuthRefreshPost) | **Post** /api/TokenAuth/Refresh | 
[**ApiTokenAuthRegisterPost**](TokenAuthApi.md#ApiTokenAuthRegisterPost) | **Post** /api/TokenAuth/Register | 

# **ApiTokenAuthAuthenticatePost**
> WctApiModelsTokenAuthAuthenticateResultModel ApiTokenAuthAuthenticatePost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenAuthApiApiTokenAuthAuthenticatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenAuthApiApiTokenAuthAuthenticatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiModelsTokenAuthAuthenticateModel**](WctApiModelsTokenAuthAuthenticateModel.md)|  | 

### Return type

[**WctApiModelsTokenAuthAuthenticateResultModel**](WCT.Api.Models.TokenAuth.AuthenticateResultModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTokenAuthBindPhoneOrEmailPost**
> ApiTokenAuthBindPhoneOrEmailPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenAuthApiApiTokenAuthBindPhoneOrEmailPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenAuthApiApiTokenAuthBindPhoneOrEmailPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationCustomersDtoBindPhoneOrEmailDto**](WctApiApplicationCustomersDtoBindPhoneOrEmailDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTokenAuthChangeAccountPost**
> bool ApiTokenAuthChangeAccountPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenAuthApiApiTokenAuthChangeAccountPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenAuthApiApiTokenAuthChangeAccountPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiModelsTokenAuthChangeAccountModel**](WctApiModelsTokenAuthChangeAccountModel.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTokenAuthExternalAuthenticatePost**
> WctApiModelsTokenAuthExternalAuthenticateResultModel ApiTokenAuthExternalAuthenticatePost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenAuthApiApiTokenAuthExternalAuthenticatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenAuthApiApiTokenAuthExternalAuthenticatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiModelsTokenAuthExternalAuthenticateModel**](WctApiModelsTokenAuthExternalAuthenticateModel.md)|  | 

### Return type

[**WctApiModelsTokenAuthExternalAuthenticateResultModel**](WCT.Api.Models.TokenAuth.ExternalAuthenticateResultModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTokenAuthForgetPasswordPost**
> bool ApiTokenAuthForgetPasswordPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenAuthApiApiTokenAuthForgetPasswordPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenAuthApiApiTokenAuthForgetPasswordPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiModelsTokenAuthForgetPasswordModel**](WctApiModelsTokenAuthForgetPasswordModel.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTokenAuthLogOutGet**
> bool ApiTokenAuthLogOutGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTokenAuthManualSendNotificePost**
> ApiTokenAuthManualSendNotificePost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenAuthApiApiTokenAuthManualSendNotificePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenAuthApiApiTokenAuthManualSendNotificePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctAdminMessagesNotificeNotificeDto**](WctAdminMessagesNotificeNotificeDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTokenAuthModifyPasswordPost**
> bool ApiTokenAuthModifyPasswordPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenAuthApiApiTokenAuthModifyPasswordPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenAuthApiApiTokenAuthModifyPasswordPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiModelsTokenAuthModifyPasswordModel**](WctApiModelsTokenAuthModifyPasswordModel.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTokenAuthOldPwdVerifyPost**
> bool ApiTokenAuthOldPwdVerifyPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenAuthApiApiTokenAuthOldPwdVerifyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenAuthApiApiTokenAuthOldPwdVerifyPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiModelsTokenAuthOldPwdVerifyModel**](WctApiModelsTokenAuthOldPwdVerifyModel.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTokenAuthRefreshPost**
> ApiTokenAuthRefreshPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenAuthApiApiTokenAuthRefreshPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenAuthApiApiTokenAuthRefreshPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**|  | 
 **refreshToken** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTokenAuthRegisterPost**
> bool ApiTokenAuthRegisterPost(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TokenAuthApiApiTokenAuthRegisterPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenAuthApiApiTokenAuthRegisterPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WctApiApplicationAuthorizationDtoRegisterModel**](WctApiApplicationAuthorizationDtoRegisterModel.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

