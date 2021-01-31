package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCspaceAddToSingleChatRequest() *OapiCspaceAddToSingleChatRequest {
	return &OapiCspaceAddToSingleChatRequest{}
}

type OapiCspaceAddToSingleChatRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCspaceAddToSingleChatResponse
	AgentId         string
	FileName        string
	MediaId         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCspaceAddToSingleChatRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCspaceAddToSingleChatRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCspaceAddToSingleChatRequest) SetAgentId(agentId2 string) {
	this.AgentId = agentId2
}
func (this *OapiCspaceAddToSingleChatRequest) GetAgentId() string {
	return this.AgentId
}
func (this *OapiCspaceAddToSingleChatRequest) SetFileName(fileName2 string) {
	this.FileName = fileName2
}
func (this *OapiCspaceAddToSingleChatRequest) GetFileName() string {
	return this.FileName
}
func (this *OapiCspaceAddToSingleChatRequest) SetMediaId(mediaId2 string) {
	this.MediaId = mediaId2
}
func (this *OapiCspaceAddToSingleChatRequest) GetMediaId() string {
	return this.MediaId
}
func (this *OapiCspaceAddToSingleChatRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCspaceAddToSingleChatRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCspaceAddToSingleChatRequest) GetApiMethodName() string {
	return "dingtalk.oapi.cspace.add_to_single_chat"
}
func (this *OapiCspaceAddToSingleChatRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCspaceAddToSingleChatRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCspaceAddToSingleChatRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCspaceAddToSingleChatRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCspaceAddToSingleChatRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["file_name"] = this.FileName
	txtParams["media_id"] = this.MediaId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCspaceAddToSingleChatRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCspaceAddToSingleChatRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCspaceAddToSingleChatRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCspaceAddToSingleChatResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
