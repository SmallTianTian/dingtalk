package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiChatGetRequest() *OapiChatGetRequest {
	return &OapiChatGetRequest{}
}

type OapiChatGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatGetResponse
	Chatid          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatGetRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiChatGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatGetRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatGetRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.get"
}
func (this *OapiChatGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	return txtParams
}
func (this *OapiChatGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiChatGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatGetResponse struct {
	taobao.TaobaoResponse
	ChatInfo ChatInfo `json:"chat_info,omitempty"`
}
type ChatInfo struct {
	Agentidlist         []string `json:"agentidlist,omitempty"`
	ChatBannedType      int64    `json:"chatBannedType,omitempty"`
	ConversationTag     int64    `json:"conversationTag,omitempty"`
	Extidlist           []string `json:"extidlist,omitempty"`
	Icon                string   `json:"icon,omitempty"`
	ManagementType      int64    `json:"managementType,omitempty"`
	MentionAllAuthority int64    `json:"mentionAllAuthority,omitempty"`
	Name                string   `json:"name,omitempty"`
	Owner               string   `json:"owner,omitempty"`
	Searchable          int64    `json:"searchable,omitempty"`
	ShowHistoryType     int64    `json:"showHistoryType,omitempty"`
	Status              int64    `json:"status,omitempty"`
	Useridlist          []string `json:"useridlist,omitempty"`
	ValidationType      int64    `json:"validationType,omitempty"`
}
