# WctApiLogisticsDtoShipmnetIntegrationCreateShipmentInputDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingCompany** | **string** | 物流公司  ARAMEX | [optional] [default to null]
**ShipperAddressLine1** | **string** | 发货人地址信息，如建筑编号，街区，街道名称。 | [optional] [default to null]
**ShipperAddressLine2** | **string** | 发货人其他地址信息。 | [optional] [default to null]
**ShipperAddressLine3** | **string** | 发货人其他地址信息。 可选 | [optional] [default to null]
**ShipperPhoneNumber1** | **string** | 发货人电话号码1，必须为2189XXXXXXXX格式 | [optional] [default to null]
**ShipperPhoneNumber2** | **string** | 发货人电话号码2，必须为2189XXXXXXXX格式   可选 | [optional] [default to null]
**ShipperEmailAddress** | **string** | 发货人邮件 | [optional] [default to null]
**ShipperPersonName** | **string** | 发货人姓名 | [optional] [default to null]
**ShipperCompanyName** | **string** | 发货人公司名称 | [optional] [default to null]
**ConsigneeAddressLine1** | **string** | 收货人地址信息，如建筑编号，街区，街道名称。 | [optional] [default to null]
**ConsigneeAddressLine2** | **string** | 收货人其他地址信息。 | [optional] [default to null]
**ConsigneeAddressLine3** | **string** | 收货人其他地址信息。 可选 | [optional] [default to null]
**ConsigneePhoneNumber1** | **string** | 收货人电话号码1，必须为2189XXXXXXXX格式 | [optional] [default to null]
**ConsigneePhoneNumber2** | **string** | 收货人电话号码2，必须为2189XXXXXXXX格式 可选 | [optional] [default to null]
**ConsigneeEmailAddress** | **string** | 收货人邮件 | [optional] [default to null]
**ConsigneePersonName** | **string** | 收货人姓名 | [optional] [default to null]
**ShippingDateTime** | **string** | 发货时间-Aramex收到待装运货物的日期。必须为&#x27;yyyy-M-d h:mm:ss tt&#x27;格式。示例：“2021-03-03 8:30:51 PM”不能超过未来4天。 | [optional] [default to null]
**DueDate** | **string** | 指定的交货日期。 必须为&#x27;yyyy-M-d h:mm:ss tt&#x27;格式。例如：“2021-03-08 9:39:51 AM”大于发货时间小于10天。 | [optional] [default to null]
**NumberOfPieces** | **int32** | 装运件数。必须是大于零且小于1000的数字 | [optional] [default to null]
**ActualWeight** | **float64** | 实际装运总重量（单位：KG）。必须是大于零且小于1000000的数字 | [optional] [default to null]
**DescriptionOfGoods** | **string** | 装运内容的性质。例如：衣服、电子产品…。。等 | [optional] [default to null]
**ShipmentReference** | **string** | 客户希望添加的关于装运的一般细节。 | [optional] [default to null]
**TransactionReference** | **string** | 交易关联信息 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

