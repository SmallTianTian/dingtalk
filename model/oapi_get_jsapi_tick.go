package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiGetJsapiTicketRequest() *OapiGetJsapiTicketRequest {
	return &OapiGetJsapiTicketRequest{}
}

type OapiGetJsapiTicketRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiGetJsapiTicketResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiGetJsapiTicketRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiGetJsapiTicketRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiGetJsapiTicketRequest) GetApiMethodName() string {
	return "dingtalk.oapi.get_jsapi_ticket"
}
func (this *OapiGetJsapiTicketRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiGetJsapiTicketRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiGetJsapiTicketRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiGetJsapiTicketRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiGetJsapiTicketRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiGetJsapiTicketRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiGetJsapiTicketRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiGetJsapiTicketRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiGetJsapiTicketResponse struct {
	taobao.TaobaoResponse

	ExpiresIn int64  `json:"expires_in,omitempty"`
	Ticket    string `json:"ticket,omitempty"`
}
