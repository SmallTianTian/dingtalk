package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMicroappDeleteRequest() *OapiMicroappDeleteRequest {
	return &OapiMicroappDeleteRequest{}
}

type OapiMicroappDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappDeleteResponse
	AgentId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMicroappDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappDeleteRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappDeleteRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.delete"
}
func (this *OapiMicroappDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentId"] = this.AgentId
	return txtParams
}
func (this *OapiMicroappDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMicroappDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappDeleteResponse struct {
	taobao.TaobaoResponse
}
