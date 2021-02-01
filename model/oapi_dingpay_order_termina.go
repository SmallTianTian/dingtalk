package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingpayOrderTerminateRequest() *OapiDingpayOrderTerminateRequest {
	return &OapiDingpayOrderTerminateRequest{}
}

type OapiDingpayOrderTerminateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingpayOrderTerminateResponse
	Extension       string
	Operator        string
	OrderNos        string
	Reason          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingpayOrderTerminateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingpayOrderTerminateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingpayOrderTerminateRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayOrderTerminateRequest) SetExtensionString(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayOrderTerminateRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiDingpayOrderTerminateRequest) SetOperator(operator2 string) {
	this.Operator = operator2
}
func (this *OapiDingpayOrderTerminateRequest) GetOperator() string {
	return this.Operator
}
func (this *OapiDingpayOrderTerminateRequest) SetOrderNos(orderNos2 string) {
	this.OrderNos = orderNos2
}
func (this *OapiDingpayOrderTerminateRequest) GetOrderNos() string {
	return this.OrderNos
}
func (this *OapiDingpayOrderTerminateRequest) SetReason(reason2 string) {
	this.Reason = reason2
}
func (this *OapiDingpayOrderTerminateRequest) GetReason() string {
	return this.Reason
}
func (this *OapiDingpayOrderTerminateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingpay.order.terminate"
}
func (this *OapiDingpayOrderTerminateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingpayOrderTerminateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingpayOrderTerminateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingpayOrderTerminateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingpayOrderTerminateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["extension"] = this.Extension
	txtParams["operator"] = this.Operator
	txtParams["order_nos"] = this.OrderNos
	txtParams["reason"] = this.Reason
	return txtParams
}
func (this *OapiDingpayOrderTerminateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Operator, "operator"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingpayOrderTerminateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingpayOrderTerminateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingpayOrderTerminateResponse struct {
	taobao.TaobaoResponse
	Result  OrderTerminationOpenResponse `json:"result,omitempty"`
	Success bool                         `json:"success,omitempty"`
}
type OrderTerminateResultItem struct {
	OrderNo    string `json:"order_no,omitempty"`
	Terminated bool   `json:"terminated,omitempty"`
}
type OrderTerminationOpenResponse struct {
	TerminateResult []OrderTerminateResultItem `json:"terminate_result,omitempty"`
}
