package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringPersonalorderPushRequest() *OapiCateringPersonalorderPushRequest {
	return &OapiCateringPersonalorderPushRequest{}
}

type OapiCateringPersonalorderPushRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp             OapiCateringPersonalorderPushResponse
	FeeActuallyPay   int64
	FeeAfterDiscount int64
	FeeOriginal      int64
	FeeShouldPay     int64
	OrderDetails     string
	OrderId          string
	PaymentTime      time.Time
	ShopId           string
	ShopName         string
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiCateringPersonalorderPushRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringPersonalorderPushRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringPersonalorderPushRequest) SetFeeActuallyPay(feeActuallyPay2 int64) {
	this.FeeActuallyPay = feeActuallyPay2
}
func (this *OapiCateringPersonalorderPushRequest) GetFeeActuallyPay() int64 {
	return this.FeeActuallyPay
}
func (this *OapiCateringPersonalorderPushRequest) SetFeeAfterDiscount(feeAfterDiscount2 int64) {
	this.FeeAfterDiscount = feeAfterDiscount2
}
func (this *OapiCateringPersonalorderPushRequest) GetFeeAfterDiscount() int64 {
	return this.FeeAfterDiscount
}
func (this *OapiCateringPersonalorderPushRequest) SetFeeOriginal(feeOriginal2 int64) {
	this.FeeOriginal = feeOriginal2
}
func (this *OapiCateringPersonalorderPushRequest) GetFeeOriginal() int64 {
	return this.FeeOriginal
}
func (this *OapiCateringPersonalorderPushRequest) SetFeeShouldPay(feeShouldPay2 int64) {
	this.FeeShouldPay = feeShouldPay2
}
func (this *OapiCateringPersonalorderPushRequest) GetFeeShouldPay() int64 {
	return this.FeeShouldPay
}
func (this *OapiCateringPersonalorderPushRequest) SetOrderDetails(orderDetails2 string) {
	this.OrderDetails = orderDetails2
}
func (this *OapiCateringPersonalorderPushRequest) GetOrderDetails() string {
	return this.OrderDetails
}
func (this *OapiCateringPersonalorderPushRequest) SetOrderId(orderId2 string) {
	this.OrderId = orderId2
}
func (this *OapiCateringPersonalorderPushRequest) GetOrderId() string {
	return this.OrderId
}
func (this *OapiCateringPersonalorderPushRequest) SetPaymentTime(paymentTime2 time.Time) {
	this.PaymentTime = paymentTime2
}
func (this *OapiCateringPersonalorderPushRequest) GetPaymentTime() time.Time {
	return this.PaymentTime
}
func (this *OapiCateringPersonalorderPushRequest) SetShopId(shopId2 string) {
	this.ShopId = shopId2
}
func (this *OapiCateringPersonalorderPushRequest) GetShopId() string {
	return this.ShopId
}
func (this *OapiCateringPersonalorderPushRequest) SetShopName(shopName2 string) {
	this.ShopName = shopName2
}
func (this *OapiCateringPersonalorderPushRequest) GetShopName() string {
	return this.ShopName
}
func (this *OapiCateringPersonalorderPushRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.personalorder.push"
}
func (this *OapiCateringPersonalorderPushRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringPersonalorderPushRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringPersonalorderPushRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringPersonalorderPushRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringPersonalorderPushRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["fee_actually_pay"] = this.FeeActuallyPay
	txtParams["fee_after_discount"] = this.FeeAfterDiscount
	txtParams["fee_original"] = this.FeeOriginal
	txtParams["fee_should_pay"] = this.FeeShouldPay
	txtParams["order_details"] = this.OrderDetails
	txtParams["order_id"] = this.OrderId
	txtParams["payment_time"] = this.PaymentTime
	txtParams["shop_id"] = this.ShopId
	txtParams["shop_name"] = this.ShopName
	return txtParams
}
func (this *OapiCateringPersonalorderPushRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OrderDetails, "orderDetails"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringPersonalorderPushRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringPersonalorderPushRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringPersonalorderPushResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
