package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiFinanceAlipayAccountGetbyuidRequest() *OapiFinanceAlipayAccountGetbyuidRequest {
	return &OapiFinanceAlipayAccountGetbyuidRequest{}
}

type OapiFinanceAlipayAccountGetbyuidRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceAlipayAccountGetbyuidResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiFinanceAlipayAccountGetbyuidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceAlipayAccountGetbyuidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceAlipayAccountGetbyuidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.alipay.account.getbyuid"
}
func (this *OapiFinanceAlipayAccountGetbyuidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceAlipayAccountGetbyuidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceAlipayAccountGetbyuidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceAlipayAccountGetbyuidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceAlipayAccountGetbyuidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiFinanceAlipayAccountGetbyuidRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiFinanceAlipayAccountGetbyuidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceAlipayAccountGetbyuidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceAlipayAccountGetbyuidResponse struct {
	taobao.TaobaoResponse
	Result AlipayUserVo `json:"result,omitempty"`
}
type AlipayUserVo struct {
	AlipayUserId string `json:"alipay_user_id,omitempty"`
}
