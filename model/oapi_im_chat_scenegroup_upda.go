package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatScenegroupUpdateRequest() *OapiImChatScenegroupUpdateRequest {
	return &OapiImChatScenegroupUpdateRequest{}
}

type OapiImChatScenegroupUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                OapiImChatScenegroupUpdateResponse
	ChatBannedType      int64
	Icon                string
	ManagementType      int64
	MentionAllAuthority int64
	OpenConversationId  string
	OwnerUserId         string
	Searchable          int64
	ShowHistoryType     int64
	Title               string
	TopHttpMethod       string
	TopResponseType     string
	ValidationType      int64
}

func (this *OapiImChatScenegroupUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatScenegroupUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatScenegroupUpdateRequest) SetChatBannedType(chatBannedType2 int64) {
	this.ChatBannedType = chatBannedType2
}
func (this *OapiImChatScenegroupUpdateRequest) GetChatBannedType() int64 {
	return this.ChatBannedType
}
func (this *OapiImChatScenegroupUpdateRequest) SetIcon(icon2 string) {
	this.Icon = icon2
}
func (this *OapiImChatScenegroupUpdateRequest) GetIcon() string {
	return this.Icon
}
func (this *OapiImChatScenegroupUpdateRequest) SetManagementType(managementType2 int64) {
	this.ManagementType = managementType2
}
func (this *OapiImChatScenegroupUpdateRequest) GetManagementType() int64 {
	return this.ManagementType
}
func (this *OapiImChatScenegroupUpdateRequest) SetMentionAllAuthority(mentionAllAuthority2 int64) {
	this.MentionAllAuthority = mentionAllAuthority2
}
func (this *OapiImChatScenegroupUpdateRequest) GetMentionAllAuthority() int64 {
	return this.MentionAllAuthority
}
func (this *OapiImChatScenegroupUpdateRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiImChatScenegroupUpdateRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiImChatScenegroupUpdateRequest) SetOwnerUserId(ownerUserId2 string) {
	this.OwnerUserId = ownerUserId2
}
func (this *OapiImChatScenegroupUpdateRequest) GetOwnerUserId() string {
	return this.OwnerUserId
}
func (this *OapiImChatScenegroupUpdateRequest) SetSearchable(searchable2 int64) {
	this.Searchable = searchable2
}
func (this *OapiImChatScenegroupUpdateRequest) GetSearchable() int64 {
	return this.Searchable
}
func (this *OapiImChatScenegroupUpdateRequest) SetShowHistoryType(showHistoryType2 int64) {
	this.ShowHistoryType = showHistoryType2
}
func (this *OapiImChatScenegroupUpdateRequest) GetShowHistoryType() int64 {
	return this.ShowHistoryType
}
func (this *OapiImChatScenegroupUpdateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiImChatScenegroupUpdateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiImChatScenegroupUpdateRequest) SetValidationType(validationType2 int64) {
	this.ValidationType = validationType2
}
func (this *OapiImChatScenegroupUpdateRequest) GetValidationType() int64 {
	return this.ValidationType
}
func (this *OapiImChatScenegroupUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.scenegroup.update"
}
func (this *OapiImChatScenegroupUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatScenegroupUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatScenegroupUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatScenegroupUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatScenegroupUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_banned_type"] = this.ChatBannedType
	txtParams["icon"] = this.Icon
	txtParams["management_type"] = this.ManagementType
	txtParams["mention_all_authority"] = this.MentionAllAuthority
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["owner_user_id"] = this.OwnerUserId
	txtParams["searchable"] = this.Searchable
	txtParams["show_history_type"] = this.ShowHistoryType
	txtParams["title"] = this.Title
	txtParams["validation_type"] = this.ValidationType
	return txtParams
}
func (this *OapiImChatScenegroupUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenConversationId, "openConversationId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatScenegroupUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatScenegroupUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatScenegroupUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
