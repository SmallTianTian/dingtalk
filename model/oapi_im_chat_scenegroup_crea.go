package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatScenegroupCreateRequest() *OapiImChatScenegroupCreateRequest {
	return &OapiImChatScenegroupCreateRequest{}
}

type OapiImChatScenegroupCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                OapiImChatScenegroupCreateResponse
	ChatBannedType      int64
	Icon                string
	ManagementType      int64
	MentionAllAuthority int64
	OwnerUserId         string
	Searchable          int64
	ShowHistoryType     int64
	TemplateId          string
	Title               string
	TopHttpMethod       string
	TopResponseType     string
	UserIds             string
	Uuid                string
	ValidationType      int64
}

func (this *OapiImChatScenegroupCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatScenegroupCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatScenegroupCreateRequest) SetChatBannedType(chatBannedType2 int64) {
	this.ChatBannedType = chatBannedType2
}
func (this *OapiImChatScenegroupCreateRequest) GetChatBannedType() int64 {
	return this.ChatBannedType
}
func (this *OapiImChatScenegroupCreateRequest) SetIcon(icon2 string) {
	this.Icon = icon2
}
func (this *OapiImChatScenegroupCreateRequest) GetIcon() string {
	return this.Icon
}
func (this *OapiImChatScenegroupCreateRequest) SetManagementType(managementType2 int64) {
	this.ManagementType = managementType2
}
func (this *OapiImChatScenegroupCreateRequest) GetManagementType() int64 {
	return this.ManagementType
}
func (this *OapiImChatScenegroupCreateRequest) SetMentionAllAuthority(mentionAllAuthority2 int64) {
	this.MentionAllAuthority = mentionAllAuthority2
}
func (this *OapiImChatScenegroupCreateRequest) GetMentionAllAuthority() int64 {
	return this.MentionAllAuthority
}
func (this *OapiImChatScenegroupCreateRequest) SetOwnerUserId(ownerUserId2 string) {
	this.OwnerUserId = ownerUserId2
}
func (this *OapiImChatScenegroupCreateRequest) GetOwnerUserId() string {
	return this.OwnerUserId
}
func (this *OapiImChatScenegroupCreateRequest) SetSearchable(searchable2 int64) {
	this.Searchable = searchable2
}
func (this *OapiImChatScenegroupCreateRequest) GetSearchable() int64 {
	return this.Searchable
}
func (this *OapiImChatScenegroupCreateRequest) SetShowHistoryType(showHistoryType2 int64) {
	this.ShowHistoryType = showHistoryType2
}
func (this *OapiImChatScenegroupCreateRequest) GetShowHistoryType() int64 {
	return this.ShowHistoryType
}
func (this *OapiImChatScenegroupCreateRequest) SetTemplateId(templateId2 string) {
	this.TemplateId = templateId2
}
func (this *OapiImChatScenegroupCreateRequest) GetTemplateId() string {
	return this.TemplateId
}
func (this *OapiImChatScenegroupCreateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiImChatScenegroupCreateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiImChatScenegroupCreateRequest) SetUserIds(userIds2 string) {
	this.UserIds = userIds2
}
func (this *OapiImChatScenegroupCreateRequest) GetUserIds() string {
	return this.UserIds
}
func (this *OapiImChatScenegroupCreateRequest) SetUuid(uuid2 string) {
	this.Uuid = uuid2
}
func (this *OapiImChatScenegroupCreateRequest) GetUuid() string {
	return this.Uuid
}
func (this *OapiImChatScenegroupCreateRequest) SetValidationType(validationType2 int64) {
	this.ValidationType = validationType2
}
func (this *OapiImChatScenegroupCreateRequest) GetValidationType() int64 {
	return this.ValidationType
}
func (this *OapiImChatScenegroupCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.scenegroup.create"
}
func (this *OapiImChatScenegroupCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatScenegroupCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatScenegroupCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatScenegroupCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatScenegroupCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_banned_type"] = this.ChatBannedType
	txtParams["icon"] = this.Icon
	txtParams["management_type"] = this.ManagementType
	txtParams["mention_all_authority"] = this.MentionAllAuthority
	txtParams["owner_user_id"] = this.OwnerUserId
	txtParams["searchable"] = this.Searchable
	txtParams["show_history_type"] = this.ShowHistoryType
	txtParams["template_id"] = this.TemplateId
	txtParams["title"] = this.Title
	txtParams["user_ids"] = this.UserIds
	txtParams["uuid"] = this.Uuid
	txtParams["validation_type"] = this.ValidationType
	return txtParams
}
func (this *OapiImChatScenegroupCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OwnerUserId, "ownerUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatScenegroupCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatScenegroupCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatScenegroupCreateResponse struct {
	taobao.TaobaoResponse
	Result  OpenSceneGroupCreateResponse `json:"result,omitempty"`
	Success bool                         `json:"success,omitempty"`
}
type OpenSceneGroupCreateResponse struct {
	ChatId             string `json:"chat_id,omitempty"`
	OpenConversationId string `json:"open_conversation_id,omitempty"`
}
