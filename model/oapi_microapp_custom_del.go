package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMicroappCustomDeleteRequest() *OapiMicroappCustomDeleteRequest {
	return &OapiMicroappCustomDeleteRequest{}
}

type OapiMicroappCustomDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappCustomDeleteResponse
	AgentId         int64
	AppCorpId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMicroappCustomDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappCustomDeleteRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappCustomDeleteRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappCustomDeleteRequest) SetAppCorpId(appCorpId2 string) {
	this.AppCorpId = appCorpId2
}
func (this *OapiMicroappCustomDeleteRequest) GetAppCorpId() string {
	return this.AppCorpId
}
func (this *OapiMicroappCustomDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.custom.delete"
}
func (this *OapiMicroappCustomDeleteRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *OapiMicroappCustomDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappCustomDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappCustomDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappCustomDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappCustomDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["app_corp_id"] = this.AppCorpId
	return txtParams
}
func (this *OapiMicroappCustomDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMicroappCustomDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappCustomDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappCustomDeleteResponse struct {
	taobao.TaobaoResponse
	Errmsg  string `json:"errmsg,omitempty"`
	Errocde int64  `json:"errocde,omitempty"`
}
