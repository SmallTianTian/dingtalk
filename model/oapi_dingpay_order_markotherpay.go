package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingpayOrderMarkotherpayRequest() *OapiDingpayOrderMarkotherpayRequest {
	return &OapiDingpayOrderMarkotherpayRequest{}
}

type OapiDingpayOrderMarkotherpayRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiDingpayOrderMarkotherpayResponse
	ApplyPayOperatorUserid string
	Extension              string
	OrderNos               string
	PayChannelPayerRealUid string
	TopHttpMethod          string
	TopResponseType        string
}

func (this *OapiDingpayOrderMarkotherpayRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingpayOrderMarkotherpayRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingpayOrderMarkotherpayRequest) SetApplyPayOperatorUserid(applyPayOperatorUserid2 string) {
	this.ApplyPayOperatorUserid = applyPayOperatorUserid2
}
func (this *OapiDingpayOrderMarkotherpayRequest) GetApplyPayOperatorUserid() string {
	return this.ApplyPayOperatorUserid
}
func (this *OapiDingpayOrderMarkotherpayRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayOrderMarkotherpayRequest) SetExtensionString(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayOrderMarkotherpayRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiDingpayOrderMarkotherpayRequest) SetOrderNos(orderNos2 string) {
	this.OrderNos = orderNos2
}
func (this *OapiDingpayOrderMarkotherpayRequest) GetOrderNos() string {
	return this.OrderNos
}
func (this *OapiDingpayOrderMarkotherpayRequest) SetPayChannelPayerRealUid(payChannelPayerRealUid2 string) {
	this.PayChannelPayerRealUid = payChannelPayerRealUid2
}
func (this *OapiDingpayOrderMarkotherpayRequest) GetPayChannelPayerRealUid() string {
	return this.PayChannelPayerRealUid
}
func (this *OapiDingpayOrderMarkotherpayRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingpay.order.markotherpay"
}
func (this *OapiDingpayOrderMarkotherpayRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingpayOrderMarkotherpayRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingpayOrderMarkotherpayRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingpayOrderMarkotherpayRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingpayOrderMarkotherpayRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["apply_pay_operator_userid"] = this.ApplyPayOperatorUserid
	txtParams["extension"] = this.Extension
	txtParams["order_nos"] = this.OrderNos
	txtParams["pay_channel_payer_real_uid"] = this.PayChannelPayerRealUid
	return txtParams
}
func (this *OapiDingpayOrderMarkotherpayRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.OrderNos, 20, "orderNos"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingpayOrderMarkotherpayRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingpayOrderMarkotherpayRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingpayOrderMarkotherpayResponse struct {
	taobao.TaobaoResponse
	Result  OrderMarkOtherPayOpenResponse `json:"result,omitempty"`
	Success bool                          `json:"success,omitempty"`
}
type OrderMarkOtherPayOpenResponse struct {
	ResultMap string `json:"result_map,omitempty"`
}
