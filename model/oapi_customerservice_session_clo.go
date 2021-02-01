package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCustomerserviceSessionCloseRequest() *OapiCustomerserviceSessionCloseRequest {
	return &OapiCustomerserviceSessionCloseRequest{}
}

type OapiCustomerserviceSessionCloseRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCustomerserviceSessionCloseResponse
	CloseSession    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCustomerserviceSessionCloseRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomerserviceSessionCloseRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomerserviceSessionCloseRequest) SetCloseSession(closeSession2 string) {
	this.CloseSession = closeSession2
}
func (this *OapiCustomerserviceSessionCloseRequest) GetCloseSession() string {
	return this.CloseSession
}
func (this *OapiCustomerserviceSessionCloseRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customerservice.session.close"
}
func (this *OapiCustomerserviceSessionCloseRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomerserviceSessionCloseRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomerserviceSessionCloseRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomerserviceSessionCloseRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomerserviceSessionCloseRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["close_session"] = this.CloseSession
	return txtParams
}
func (this *OapiCustomerserviceSessionCloseRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCustomerserviceSessionCloseRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomerserviceSessionCloseRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CloseSessionDTO struct {
	BuId          string `json:"bu_id,omitempty"`
	OperatorId    string `json:"operator_id,omitempty"`
	OperatorType  string `json:"operator_type,omitempty"`
	Reason        string `json:"reason,omitempty"`
	ServiceId     string `json:"service_id,omitempty"`
	SessionSource string `json:"session_source,omitempty"`
	Sid           string `json:"sid,omitempty"`
}
type OapiCustomerserviceSessionCloseResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
