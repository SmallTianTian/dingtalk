package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCspaceGrantCustomSpaceRequest() *OapiCspaceGrantCustomSpaceRequest {
	return &OapiCspaceGrantCustomSpaceRequest{}
}

type OapiCspaceGrantCustomSpaceRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCspaceGrantCustomSpaceResponse
	AgentId         string
	Domain          string
	Duration        int64
	Fileids         string
	Path            string
	TopHttpMethod   string
	TopResponseType string
	Type            string
	Userid          string
}

func (this *OapiCspaceGrantCustomSpaceRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCspaceGrantCustomSpaceRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiCspaceGrantCustomSpaceRequest) SetDomain(domain2 string) {
	this.Domain = domain2
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetDomain() string {
	return this.Domain
}
func (this *OapiCspaceGrantCustomSpaceRequest) SetDuration(duration2 int64) {
	this.Duration = duration2
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetDuration() int64 {
	return this.Duration
}
func (this *OapiCspaceGrantCustomSpaceRequest) SetFileids(fileids2 string) {
	this.Fileids = fileids2
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetFileids() string {
	return this.Fileids
}
func (this *OapiCspaceGrantCustomSpaceRequest) SetPath(path2 string) {
	this.Path = path2
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetPath() string {
	return this.Path
}
func (this *OapiCspaceGrantCustomSpaceRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetType() string {
	return this.Type
}
func (this *OapiCspaceGrantCustomSpaceRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetApiMethodName() string {
	return "dingtalk.oapi.cspace.grant_custom_space"
}
func (this *OapiCspaceGrantCustomSpaceRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCspaceGrantCustomSpaceRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCspaceGrantCustomSpaceRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["domain"] = this.Domain
	txtParams["duration"] = this.Duration
	txtParams["fileids"] = this.Fileids
	txtParams["path"] = this.Path
	txtParams["type"] = this.Type
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCspaceGrantCustomSpaceRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCspaceGrantCustomSpaceRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCspaceGrantCustomSpaceResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
