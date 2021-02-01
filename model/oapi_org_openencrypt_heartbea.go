package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiOrgOpenencryptHeartbeatRequest() *OapiOrgOpenencryptHeartbeatRequest {
	return &OapiOrgOpenencryptHeartbeatRequest{}
}

type OapiOrgOpenencryptHeartbeatRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiOrgOpenencryptHeartbeatResponse
	TopHttpMethod   string
	TopKmsHeartbeat string
	TopResponseType string
}

func (this *OapiOrgOpenencryptHeartbeatRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOrgOpenencryptHeartbeatRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOrgOpenencryptHeartbeatRequest) SetTopKmsHeartbeat(topKmsHeartbeat2 string) {
	this.TopKmsHeartbeat = topKmsHeartbeat2
}
func (this *OapiOrgOpenencryptHeartbeatRequest) GetTopKmsHeartbeat() string {
	return this.TopKmsHeartbeat
}
func (this *OapiOrgOpenencryptHeartbeatRequest) GetApiMethodName() string {
	return "dingtalk.oapi.org.openencrypt.heartbeat"
}
func (this *OapiOrgOpenencryptHeartbeatRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOrgOpenencryptHeartbeatRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOrgOpenencryptHeartbeatRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOrgOpenencryptHeartbeatRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOrgOpenencryptHeartbeatRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["top_kms_heartbeat"] = this.TopKmsHeartbeat
	return txtParams
}
func (this *OapiOrgOpenencryptHeartbeatRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiOrgOpenencryptHeartbeatRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOrgOpenencryptHeartbeatRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOrgOpenencryptHeartbeatResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
