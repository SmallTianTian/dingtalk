package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"time"
)

func NewOapiCateringOrderPushRequest() *OapiCateringOrderPushRequest {
	return &OapiCateringOrderPushRequest{}
}

type OapiCateringOrderPushRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp             OapiCateringOrderPushResponse
	FeeActuallyPay   int64
	FeeAfterDiscount int64
	FeeOriginal      int64
	FeeShouldPay     int64
	OrderDetails     string
	OrderId          string
	OrdererUserId    string
	PaymentTime      time.Time
	ShopId           string
	ShopName         string
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiCateringOrderPushRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringOrderPushRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringOrderPushRequest) SetFeeActuallyPay(feeActuallyPay2 int64) {
	this.FeeActuallyPay = feeActuallyPay2
}
func (this *OapiCateringOrderPushRequest) GetFeeActuallyPay() int64 {
	return this.FeeActuallyPay
}
func (this *OapiCateringOrderPushRequest) SetFeeAfterDiscount(feeAfterDiscount2 int64) {
	this.FeeAfterDiscount = feeAfterDiscount2
}
func (this *OapiCateringOrderPushRequest) GetFeeAfterDiscount() int64 {
	return this.FeeAfterDiscount
}
func (this *OapiCateringOrderPushRequest) SetFeeOriginal(feeOriginal2 int64) {
	this.FeeOriginal = feeOriginal2
}
func (this *OapiCateringOrderPushRequest) GetFeeOriginal() int64 {
	return this.FeeOriginal
}
func (this *OapiCateringOrderPushRequest) SetFeeShouldPay(feeShouldPay2 int64) {
	this.FeeShouldPay = feeShouldPay2
}
func (this *OapiCateringOrderPushRequest) GetFeeShouldPay() int64 {
	return this.FeeShouldPay
}
func (this *OapiCateringOrderPushRequest) SetOrderDetails(orderDetails2 string) {
	this.OrderDetails = orderDetails2
}
func (this *OapiCateringOrderPushRequest) GetOrderDetails() string {
	return this.OrderDetails
}
func (this *OapiCateringOrderPushRequest) SetOrderId(orderId2 string) {
	this.OrderId = orderId2
}
func (this *OapiCateringOrderPushRequest) GetOrderId() string {
	return this.OrderId
}
func (this *OapiCateringOrderPushRequest) SetOrdererUserId(ordererUserId2 string) {
	this.OrdererUserId = ordererUserId2
}
func (this *OapiCateringOrderPushRequest) GetOrdererUserId() string {
	return this.OrdererUserId
}
func (this *OapiCateringOrderPushRequest) SetPaymentTime(paymentTime2 time.Time) {
	this.PaymentTime = paymentTime2
}
func (this *OapiCateringOrderPushRequest) GetPaymentTime() time.Time {
	return this.PaymentTime
}
func (this *OapiCateringOrderPushRequest) SetShopId(shopId2 string) {
	this.ShopId = shopId2
}
func (this *OapiCateringOrderPushRequest) GetShopId() string {
	return this.ShopId
}
func (this *OapiCateringOrderPushRequest) SetShopName(shopName2 string) {
	this.ShopName = shopName2
}
func (this *OapiCateringOrderPushRequest) GetShopName() string {
	return this.ShopName
}
func (this *OapiCateringOrderPushRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.order.push"
}
func (this *OapiCateringOrderPushRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringOrderPushRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringOrderPushRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringOrderPushRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringOrderPushRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["fee_actually_pay"] = this.FeeActuallyPay
	txtParams["fee_after_discount"] = this.FeeAfterDiscount
	txtParams["fee_original"] = this.FeeOriginal
	txtParams["fee_should_pay"] = this.FeeShouldPay
	txtParams["order_details"] = this.OrderDetails
	txtParams["order_id"] = this.OrderId
	txtParams["orderer_user_id"] = this.OrdererUserId
	txtParams["payment_time"] = this.PaymentTime
	txtParams["shop_id"] = this.ShopId
	txtParams["shop_name"] = this.ShopName
	return txtParams
}
func (this *OapiCateringOrderPushRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCateringOrderPushRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringOrderPushRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringOrderPushResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
