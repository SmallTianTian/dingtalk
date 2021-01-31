package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingpayOrderApplypayRequest() *OapiDingpayOrderApplypayRequest {
	return &OapiDingpayOrderApplypayRequest{}
}

type OapiDingpayOrderApplypayRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiDingpayOrderApplypayResponse
	ApplyPayOperatorUserid string
	Extension              string
	OrderNos               string
	PayChannel             string
	PayChannelPayerRealUid string
	TopHttpMethod          string
	TopResponseType        string
}

func (this *OapiDingpayOrderApplypayRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingpayOrderApplypayRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingpayOrderApplypayRequest) SetApplyPayOperatorUserid(applyPayOperatorUserid2 string) {
	this.ApplyPayOperatorUserid = applyPayOperatorUserid2
}
func (this *OapiDingpayOrderApplypayRequest) GetApplyPayOperatorUserid() string {
	return this.ApplyPayOperatorUserid
}
func (this *OapiDingpayOrderApplypayRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayOrderApplypayRequest) SetExtensionString(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayOrderApplypayRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiDingpayOrderApplypayRequest) SetOrderNos(orderNos2 string) {
	this.OrderNos = orderNos2
}
func (this *OapiDingpayOrderApplypayRequest) GetOrderNos() string {
	return this.OrderNos
}
func (this *OapiDingpayOrderApplypayRequest) SetPayChannel(payChannel2 string) {
	this.PayChannel = payChannel2
}
func (this *OapiDingpayOrderApplypayRequest) GetPayChannel() string {
	return this.PayChannel
}
func (this *OapiDingpayOrderApplypayRequest) SetPayChannelPayerRealUid(payChannelPayerRealUid2 string) {
	this.PayChannelPayerRealUid = payChannelPayerRealUid2
}
func (this *OapiDingpayOrderApplypayRequest) GetPayChannelPayerRealUid() string {
	return this.PayChannelPayerRealUid
}
func (this *OapiDingpayOrderApplypayRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingpay.order.applypay"
}
func (this *OapiDingpayOrderApplypayRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingpayOrderApplypayRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingpayOrderApplypayRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingpayOrderApplypayRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingpayOrderApplypayRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["apply_pay_operator_userid"] = this.ApplyPayOperatorUserid
	txtParams["extension"] = this.Extension
	txtParams["order_nos"] = this.OrderNos
	txtParams["pay_channel"] = this.PayChannel
	txtParams["pay_channel_payer_real_uid"] = this.PayChannelPayerRealUid
	return txtParams
}
func (this *OapiDingpayOrderApplypayRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ApplyPayOperatorUserid, "applyPayOperatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingpayOrderApplypayRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingpayOrderApplypayRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingpayOrderApplypayResponse struct {
	taobao.TaobaoResponse
	Result  OrderApplyPayOpenResponse `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
type OrderApplyPayOpenResponse struct {
	OrderStr string `json:"orderStr,omitempty"`
}
