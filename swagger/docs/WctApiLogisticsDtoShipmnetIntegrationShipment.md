# WctApiLogisticsDtoShipmnetIntegrationShipment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference1** | **string** | 客户想补充的关于货物的一般细节 | [optional] [default to null]
**Reference2** | **string** | 客户想补充的关于货物的一般细节 | [optional] [default to null]
**Reference3** | **string** | 客户想补充的关于货物的一般细节 | [optional] [default to null]
**Shipper** | [***WctApiLogisticsDtoShipmnetIntegrationUserInformation**](WCT.Api.Logistics.Dto.ShipmnetIntegration.UserInformation.md) |  | [default to null]
**Consignee** | [***WctApiLogisticsDtoShipmnetIntegrationUserInformation**](WCT.Api.Logistics.Dto.ShipmnetIntegration.UserInformation.md) |  | [default to null]
**ThirdParty** | [***WctApiLogisticsDtoShipmnetIntegrationUserInformation**](WCT.Api.Logistics.Dto.ShipmnetIntegration.UserInformation.md) |  | [optional] [default to null]
**ShippingDateTime** | **string** | 发货时间-格式为：/Date(1613132715000-0500)/ | [default to null]
**DueDate** | **string** | 到货时间-格式为：/Date(1613132715000-0500)/ | [default to null]
**Comments** | **string** | 评论 | [optional] [default to null]
**PickupLocation** | **string** | 自提点地址 | [optional] [default to null]
**OperationsInstructions** | **string** | 如何处理货物的说明 | [optional] [default to null]
**AccountingInstrcutions** | **string** | 账户-付款方式说明细节 | [optional] [default to null]
**Details** | [***WctApiLogisticsDtoShipmnetIntegrationShipmentDetails**](WCT.Api.Logistics.Dto.ShipmnetIntegration.ShipmentDetails.md) |  | [default to null]
**ForeignHAWB** | **string** | 客户的装运编号（如有）。如果已填写，则此字段对于每个装运都必须是唯一的 | [optional] [default to null]
**TransportType** | [***WctApiLogisticsDtoShipmnetIntegrationTransportType**](WCT.Api.Logistics.Dto.ShipmnetIntegration.TransportType.md) |  | [optional] [default to null]
**PickupGUID** | **string** | To add Shipments to existing pickups | [optional] [default to null]
**Number** | **string** | If shipment numbers are required to be entered  manually then aramex operations will provide a stock range from which to fill this field with. Otherwise if empty a number will be assigned to the created shipment automatically and returned in the response. | [optional] [default to null]
**ScheduledDelivery** | [***WctApiLogisticsDtoShipmnetIntegrationScheduledDelivery**](WCT.Api.Logistics.Dto.ShipmnetIntegration.ScheduledDelivery.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

