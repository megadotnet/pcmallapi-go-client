# WctApiApplicationPayPalDtoApplicationContext

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandName** | **string** | 覆盖PayPal站点上PayPal帐户中的业务名称的标签 | [optional] [default to null]
**CancelUrl** | **string** | 客户取消付款后重定向到的网址。 | [default to null]
**LandingPage** | **string** | 要显示在PayPal站点上供客户结帐的登陆页面的类型。  LOGIN：当客户点击PayPal检查，客户将被重定向到一个页面，以登录到PayPal并批准付款。  BILLING：当客户点击PayPal检查，将客户重定向到页面，输入信用卡或借记卡以及完成购买所需的其他相关记帐信息。  NO_PREFERENCE：当客户点击PayPal检查，客户将被重定向到登录到PayPal并批准付款的页面，或者重定向到页面输入完成购买所需的信用卡或借记卡以及其他相关的账单信息，这取决于他们以前与PayPal的交互。 | [optional] [default to null]
**ReturnUrl** | **string** | 在客户批准付款后，客户被重定向的URL。 | [default to null]
**UserAction** | **string** | 配置“继续”或“立即支付”的结账流程。  CONTINUE 或 PAY_NOW | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

