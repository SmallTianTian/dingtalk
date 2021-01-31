package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAppstoreOrdersInquiryRequest() *OapiAppstoreOrdersInquiryRequest {
	return &OapiAppstoreOrdersInquiryRequest{}
}

type OapiAppstoreOrdersInquiryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAppstoreOrdersInquiryResponse
	Corpid          string
	CycNum          int64
	CycUnit         int64
	GoodsCode       string
	ItemCode        string
	Mobile          string
	Quantity        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAppstoreOrdersInquiryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAppstoreOrdersInquiryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAppstoreOrdersInquiryRequest) SetCorpid(corpid2 string) {
	this.Corpid = corpid2
}
func (this *OapiAppstoreOrdersInquiryRequest) GetCorpid() string {
	return this.Corpid
}
func (this *OapiAppstoreOrdersInquiryRequest) SetCycNum(cycNum2 int64) {
	this.CycNum = cycNum2
}
func (this *OapiAppstoreOrdersInquiryRequest) GetCycNum() int64 {
	return this.CycNum
}
func (this *OapiAppstoreOrdersInquiryRequest) SetCycUnit(cycUnit2 int64) {
	this.CycUnit = cycUnit2
}
func (this *OapiAppstoreOrdersInquiryRequest) GetCycUnit() int64 {
	return this.CycUnit
}
func (this *OapiAppstoreOrdersInquiryRequest) SetGoodsCode(goodsCode2 string) {
	this.GoodsCode = goodsCode2
}
func (this *OapiAppstoreOrdersInquiryRequest) GetGoodsCode() string {
	return this.GoodsCode
}
func (this *OapiAppstoreOrdersInquiryRequest) SetItemCode(itemCode2 string) {
	this.ItemCode = itemCode2
}
func (this *OapiAppstoreOrdersInquiryRequest) GetItemCode() string {
	return this.ItemCode
}
func (this *OapiAppstoreOrdersInquiryRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiAppstoreOrdersInquiryRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiAppstoreOrdersInquiryRequest) SetQuantity(quantity2 int64) {
	this.Quantity = quantity2
}
func (this *OapiAppstoreOrdersInquiryRequest) GetQuantity() int64 {
	return this.Quantity
}
func (this *OapiAppstoreOrdersInquiryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.appstore.orders.inquiry"
}
func (this *OapiAppstoreOrdersInquiryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAppstoreOrdersInquiryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAppstoreOrdersInquiryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAppstoreOrdersInquiryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAppstoreOrdersInquiryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corpid"] = this.Corpid
	txtParams["cyc_num"] = this.CycNum
	txtParams["cyc_unit"] = this.CycUnit
	txtParams["goods_code"] = this.GoodsCode
	txtParams["item_code"] = this.ItemCode
	txtParams["mobile"] = this.Mobile
	txtParams["quantity"] = this.Quantity
	return txtParams
}
func (this *OapiAppstoreOrdersInquiryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Corpid, "corpid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAppstoreOrdersInquiryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAppstoreOrdersInquiryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAppstoreOrdersInquiryResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	PayFee  string `json:"pay_fee,omitempty"`
}
