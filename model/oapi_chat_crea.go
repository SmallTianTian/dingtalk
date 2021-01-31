package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiChatCreateRequest() *OapiChatCreateRequest {
	return &OapiChatCreateRequest{}
}

type OapiChatCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                OapiChatCreateResponse
	ChatBannedType      int64
	ConversationTag     int64
	Extidlist           []string
	Icon                string
	ManagementType      int64
	MentionAllAuthority int64
	Name                string
	Owner               string
	OwnerType           string
	Searchable          int64
	ShowHistoryType     int64
	TopHttpMethod       string
	TopResponseType     string
	Useridlist          []string
	ValidationType      int64
}

func (this *OapiChatCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatCreateRequest) SetChatBannedType(chatBannedType2 int64) {
	this.ChatBannedType = chatBannedType2
}
func (this *OapiChatCreateRequest) GetChatBannedType() int64 {
	return this.ChatBannedType
}
func (this *OapiChatCreateRequest) SetConversationTag(conversationTag2 int64) {
	this.ConversationTag = conversationTag2
}
func (this *OapiChatCreateRequest) GetConversationTag() int64 {
	return this.ConversationTag
}
func (this *OapiChatCreateRequest) SetExtidlist(extidlist2 []string) {
	this.Extidlist = extidlist2
}
func (this *OapiChatCreateRequest) GetExtidlist() []string {
	return this.Extidlist
}
func (this *OapiChatCreateRequest) SetIcon(icon2 string) {
	this.Icon = icon2
}
func (this *OapiChatCreateRequest) GetIcon() string {
	return this.Icon
}
func (this *OapiChatCreateRequest) SetManagementType(managementType2 int64) {
	this.ManagementType = managementType2
}
func (this *OapiChatCreateRequest) GetManagementType() int64 {
	return this.ManagementType
}
func (this *OapiChatCreateRequest) SetMentionAllAuthority(mentionAllAuthority2 int64) {
	this.MentionAllAuthority = mentionAllAuthority2
}
func (this *OapiChatCreateRequest) GetMentionAllAuthority() int64 {
	return this.MentionAllAuthority
}
func (this *OapiChatCreateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiChatCreateRequest) GetName() string {
	return this.Name
}
func (this *OapiChatCreateRequest) SetOwner(owner2 string) {
	this.Owner = owner2
}
func (this *OapiChatCreateRequest) GetOwner() string {
	return this.Owner
}
func (this *OapiChatCreateRequest) SetOwnerType(ownerType2 string) {
	this.OwnerType = ownerType2
}
func (this *OapiChatCreateRequest) GetOwnerType() string {
	return this.OwnerType
}
func (this *OapiChatCreateRequest) SetSearchable(searchable2 int64) {
	this.Searchable = searchable2
}
func (this *OapiChatCreateRequest) GetSearchable() int64 {
	return this.Searchable
}
func (this *OapiChatCreateRequest) SetShowHistoryType(showHistoryType2 int64) {
	this.ShowHistoryType = showHistoryType2
}
func (this *OapiChatCreateRequest) GetShowHistoryType() int64 {
	return this.ShowHistoryType
}
func (this *OapiChatCreateRequest) SetUseridlist(useridlist2 []string) {
	this.Useridlist = useridlist2
}
func (this *OapiChatCreateRequest) GetUseridlist() []string {
	return this.Useridlist
}
func (this *OapiChatCreateRequest) SetValidationType(validationType2 int64) {
	this.ValidationType = validationType2
}
func (this *OapiChatCreateRequest) GetValidationType() int64 {
	return this.ValidationType
}
func (this *OapiChatCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.create"
}
func (this *OapiChatCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatBannedType"] = this.ChatBannedType
	txtParams["conversationTag"] = this.ConversationTag
	txtParams["extidlist"] = this.Extidlist
	txtParams["icon"] = this.Icon
	txtParams["managementType"] = this.ManagementType
	txtParams["mentionAllAuthority"] = this.MentionAllAuthority
	txtParams["name"] = this.Name
	txtParams["owner"] = this.Owner
	txtParams["ownerType"] = this.OwnerType
	txtParams["searchable"] = this.Searchable
	txtParams["showHistoryType"] = this.ShowHistoryType
	txtParams["useridlist"] = this.Useridlist
	txtParams["validationType"] = this.ValidationType
	return txtParams
}
func (this *OapiChatCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiChatCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatCreateResponse struct {
	taobao.TaobaoResponse
	Chatid             string `json:"chatid,omitempty"`
	ConversationTag    int64  `json:"conversationTag,omitempty"`
	Errcode            int64  `json:"errcode,omitempty"`
	Errmsg             string `json:"errmsg,omitempty"`
	OpenConversationId string `json:"openConversationId,omitempty"`
}
