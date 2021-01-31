package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAppstoreOrdersSpecialCanalCreateOrderRequest() *OapiAppstoreOrdersSpecialCanalCreateOrderRequest {
	return &OapiAppstoreOrdersSpecialCanalCreateOrderRequest{}
}

type OapiAppstoreOrdersSpecialCanalCreateOrderRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAppstoreOrdersSpecialCanalCreateOrderResponse
	Corpid          string
	CycNum          string
	CycUnit         string
	GoodsCode       string
	ItemCode        string
	Mobile          string
	OrderCenterId   string
	Price           string
	Quantity        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetCorpid(corpid2 string) {
	this.Corpid = corpid2
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetCorpid() string {
	return this.Corpid
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetCycNum(cycNum2 string) {
	this.CycNum = cycNum2
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetCycNum() string {
	return this.CycNum
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetCycUnit(cycUnit2 string) {
	this.CycUnit = cycUnit2
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetCycUnit() string {
	return this.CycUnit
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetGoodsCode(goodsCode2 string) {
	this.GoodsCode = goodsCode2
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetGoodsCode() string {
	return this.GoodsCode
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetItemCode(itemCode2 string) {
	this.ItemCode = itemCode2
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetItemCode() string {
	return this.ItemCode
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetOrderCenterId(orderCenterId2 string) {
	this.OrderCenterId = orderCenterId2
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetOrderCenterId() string {
	return this.OrderCenterId
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetPrice(price2 string) {
	this.Price = price2
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetPrice() string {
	return this.Price
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetQuantity(quantity2 string) {
	this.Quantity = quantity2
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetQuantity() string {
	return this.Quantity
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetApiMethodName() string {
	return "dingtalk.oapi.appstore.orders.special_canal.create_order"
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corpid"] = this.Corpid
	txtParams["cyc_num"] = this.CycNum
	txtParams["cyc_unit"] = this.CycUnit
	txtParams["goods_code"] = this.GoodsCode
	txtParams["item_code"] = this.ItemCode
	txtParams["mobile"] = this.Mobile
	txtParams["order_center_id"] = this.OrderCenterId
	txtParams["price"] = this.Price
	txtParams["quantity"] = this.Quantity
	return txtParams
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Corpid, "corpid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAppstoreOrdersSpecialCanalCreateOrderRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAppstoreOrdersSpecialCanalCreateOrderResponse struct {
	taobao.TaobaoResponse
	DingOrderId string `json:"ding_order_id,omitempty"`
	Errcode     int64  `json:"errcode,omitempty"`
	Errmsg      string `json:"errmsg,omitempty"`
}
