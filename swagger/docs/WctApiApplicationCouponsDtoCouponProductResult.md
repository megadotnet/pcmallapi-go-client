# WctApiApplicationCouponsDtoCouponProductResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | UUID | [optional] [default to null]
**Name** | **string** | 优惠券名称 | [optional] [default to null]
**CouponType** | [***WctAdminCouponsCouponType**](WCT.Admin.Coupons.CouponType.md) |  | [optional] [default to null]
**CouponCardType** | [***WctAdminCouponsCouponCardType**](WCT.Admin.Coupons.CouponCardType.md) |  | [optional] [default to null]
**CouponTypeValue** | **float64** | 优惠券类型值 | [optional] [default to null]
**LimitValue** | **float64** | 使用门槛 | [optional] [default to null]
**Value** | **string** | 价值 | [optional] [default to null]
**LimitCount** | **int32** | 限制数量（0 是不限制） | [optional] [default to null]
**GiveCount** | **int64** | 发放数量 | [optional] [default to null]
**BeginDateValue** | **int32** | 领券后开始时间 | [optional] [default to null]
**ResidueDays** | **int32** | 残余时间 | [optional] [default to null]
**BeginTime** | [**time.Time**](time.Time.md) | 优惠券开始时间 | [optional] [default to null]
**EndTime** | [**time.Time**](time.Time.md) | 优惠券结束时间 | [optional] [default to null]
**AlreadyCount** | **int64** | 已领取数量 | [optional] [default to null]
**AlreadyUser** | **int64** | 已领取人 | [optional] [default to null]
**UseCount** | **int64** | 使用数量 | [optional] [default to null]
**StatusName** | **string** | 状态 | [optional] [default to null]
**IsOriginalPrice** | **bool** | 非原价是否可以参与 | [optional] [default to null]
**LimitType** | [***WctAdminCouponsLimitType**](WCT.Admin.Coupons.LimitType.md) |  | [optional] [default to null]
**CanReceive** | **bool** | 是否可以领取 | [optional] [default to null]
**IdentifyingType** | [***WctAdminCouponsIdentifyingType**](WCT.Admin.Coupons.IdentifyingType.md) |  | [optional] [default to null]
**Instructions** | **string** | 说明 | [optional] [default to null]
**TenantId** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

