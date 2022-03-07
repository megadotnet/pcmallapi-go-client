# WctApiLogisticsDtoShipmnetIntegrationShipmentDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | [***WctApiLogisticsDtoShipmnetIntegrationDimensions**](WCT.Api.Logistics.Dto.ShipmnetIntegration.Dimensions.md) |  | [optional] [default to null]
**NumberOfPieces** | **int64** | 装运件数 | [default to null]
**ActualWeight** | [***WctApiLogisticsDtoShipmnetIntegrationWeight**](WCT.Api.Logistics.Dto.ShipmnetIntegration.Weight.md) |  | [default to null]
**ProductGroup** | **string** | EXP&#x3D;快递 DOM&#x3D;国内 | [default to null]
**ProductType** | **string** | 官方文档中无CDS  CDS   产品类型包括关于交货的某些特点 产品，例如：优先级、时间敏感性和无论是文件还是非文件。有关产品类型及其产品组的列表，请参阅附录A | [default to null]
**PaymentType** | **string** | P  付款方式。参考更多详情请参见附录B | [default to null]
**PaymentOptions** | **string** | 必填-基于付款类型“C”  ASCC&#x3D;需要输入发货人帐号 填满。ARCC&#x3D;需要收货人帐户 要填写的数字。  可选-基于付款类型“P”，则可选择填写。现金&#x3D;现金账户&#x3D;账户Aramex 国际运输信息42 Aramex装运API嵌入指南 PPST &#x3D; 预付股CRDT &#x3D; 信贷 | [optional] [default to null]
**Services** | **string** | CODS | [optional] [default to null]
**DescriptionOfGoods** | **string** | 货物描述 | [default to null]
**GoodsOriginCountry** | **string** | 货物的原产地来自。请参阅附录D以获取完整的 国家代码列表。 | [default to null]
**CustomsValueAmount** | [***WctApiLogisticsDtoShipmnetIntegrationMoney**](WCT.Api.Logistics.Dto.ShipmnetIntegration.Money.md) |  | [optional] [default to null]
**CashOnDelivery** | [***WctApiLogisticsDtoShipmnetIntegrationMoney**](WCT.Api.Logistics.Dto.ShipmnetIntegration.Money.md) |  | [optional] [default to null]
**CashOnDeliveryAmount** | [***WctApiLogisticsDtoShipmnetIntegrationMoney**](WCT.Api.Logistics.Dto.ShipmnetIntegration.Money.md) |  | [optional] [default to null]
**InsuranceAmount** | [***WctApiLogisticsDtoShipmnetIntegrationMoney**](WCT.Api.Logistics.Dto.ShipmnetIntegration.Money.md) |  | [optional] [default to null]
**CashAdditionalAmount** | [***WctApiLogisticsDtoShipmnetIntegrationMoney**](WCT.Api.Logistics.Dto.ShipmnetIntegration.Money.md) |  | [optional] [default to null]
**CashAdditionalDescription** | **string** | 有条件-基于付款类型“3”，并填写现金附加金额 | [optional] [default to null]
**CollectAmount** | [***WctApiLogisticsDtoShipmnetIntegrationMoney**](WCT.Api.Logistics.Dto.ShipmnetIntegration.Money.md) |  | [optional] [default to null]
**Items** | [**[]WctApiLogisticsDtoShipmnetIntegrationShipment**](WCT.Api.Logistics.Dto.ShipmnetIntegration.Shipment.md) | null | [optional] [default to null]
**ContainsDangerousGoods** | **int32** | 危险物品 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

