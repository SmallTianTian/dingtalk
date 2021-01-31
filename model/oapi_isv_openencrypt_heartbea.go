package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiIsvOpenencryptHeartbeatRequest() *OapiIsvOpenencryptHeartbeatRequest {
	return &OapiIsvOpenencryptHeartbeatRequest{}
}

type OapiIsvOpenencryptHeartbeatRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiIsvOpenencryptHeartbeatResponse
	TopHttpMethod   string
	TopKmsHeartbeat string
	TopResponseType string
}

func (this *OapiIsvOpenencryptHeartbeatRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiIsvOpenencryptHeartbeatRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiIsvOpenencryptHeartbeatRequest) SetTopKmsHeartbeat(topKmsHeartbeat2 string) {
	this.TopKmsHeartbeat = topKmsHeartbeat2
}
func (this *OapiIsvOpenencryptHeartbeatRequest) GetTopKmsHeartbeat() string {
	return this.TopKmsHeartbeat
}
func (this *OapiIsvOpenencryptHeartbeatRequest) GetApiMethodName() string {
	return "dingtalk.oapi.isv.openencrypt.heartbeat"
}
func (this *OapiIsvOpenencryptHeartbeatRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiIsvOpenencryptHeartbeatRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiIsvOpenencryptHeartbeatRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiIsvOpenencryptHeartbeatRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiIsvOpenencryptHeartbeatRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["top_kms_heartbeat"] = this.TopKmsHeartbeat
	return txtParams
}
func (this *OapiIsvOpenencryptHeartbeatRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiIsvOpenencryptHeartbeatRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiIsvOpenencryptHeartbeatRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopKMSHeartbeat struct {
	Appid     int64  `json:"appid,omitempty"`
	Payload   string `json:"payload,omitempty"`
	Requestid string `json:"requestid,omitempty"`
}
type OapiIsvOpenencryptHeartbeatResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
