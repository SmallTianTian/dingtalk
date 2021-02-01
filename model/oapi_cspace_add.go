package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCspaceAddRequest() *OapiCspaceAddRequest {
	return &OapiCspaceAddRequest{}
}

type OapiCspaceAddRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCspaceAddResponse
	AgentId         string
	Code            string
	FolderId        string
	MediaId         string
	Name            string
	Overwrite       bool
	SpaceId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCspaceAddRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiCspaceAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCspaceAddRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiCspaceAddRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiCspaceAddRequest) SetCode(code2 string) {
	this.Code = code2
}
func (this *OapiCspaceAddRequest) GetCode() string {
	return this.Code
}
func (this *OapiCspaceAddRequest) SetFolderId(folderId2 string) {
	this.FolderId = folderId2
}
func (this *OapiCspaceAddRequest) GetFolderId() string {
	return this.FolderId
}
func (this *OapiCspaceAddRequest) SetMediaId(mediaId2 string) {
	this.MediaId = mediaId2
}
func (this *OapiCspaceAddRequest) GetMediaId() string {
	return this.MediaId
}
func (this *OapiCspaceAddRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiCspaceAddRequest) GetName() string {
	return this.Name
}
func (this *OapiCspaceAddRequest) SetOverwrite(overwrite2 bool) {
	this.Overwrite = overwrite2
}
func (this *OapiCspaceAddRequest) GetOverwrite() bool {
	return this.Overwrite
}
func (this *OapiCspaceAddRequest) SetSpaceId(spaceId2 string) {
	this.SpaceId = spaceId2
}
func (this *OapiCspaceAddRequest) GetSpaceId() string {
	return this.SpaceId
}
func (this *OapiCspaceAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.cspace.add"
}
func (this *OapiCspaceAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCspaceAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCspaceAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCspaceAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCspaceAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams[taobao.ERROR_CODE] = this.Code
	txtParams["folder_id"] = this.FolderId
	txtParams["media_id"] = this.MediaId
	txtParams["name"] = this.Name
	txtParams["overwrite"] = this.Overwrite
	txtParams["space_id"] = this.SpaceId
	return txtParams
}
func (this *OapiCspaceAddRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCspaceAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCspaceAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCspaceAddResponse struct {
	taobao.TaobaoResponse
	Dentry string `json:"dentry,omitempty"`
}
