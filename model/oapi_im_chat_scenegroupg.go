package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatScenegroupGetRequest() *OapiImChatScenegroupGetRequest {
	return &OapiImChatScenegroupGetRequest{}
}

type OapiImChatScenegroupGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiImChatScenegroupGetResponse
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiImChatScenegroupGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatScenegroupGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatScenegroupGetRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiImChatScenegroupGetRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiImChatScenegroupGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.scenegroup.get"
}
func (this *OapiImChatScenegroupGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatScenegroupGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatScenegroupGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatScenegroupGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatScenegroupGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_conversation_id"] = this.OpenConversationId
	return txtParams
}
func (this *OapiImChatScenegroupGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenConversationId, "openConversationId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatScenegroupGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatScenegroupGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatScenegroupGetResponse struct {
	taobao.TaobaoResponse
	Result  DecorationGroupQueryResponse `json:"result,omitempty"`
	Success bool                         `json:"success,omitempty"`
}
type ManagementOptions struct {
	ChatBannedType      string `json:"chat_banned_type,omitempty"`
	ManagementType      string `json:"management_type,omitempty"`
	MentionAllAuthority string `json:"mention_all_authority,omitempty"`
	Searchable          string `json:"searchable,omitempty"`
	ShowHistoryType     string `json:"show_history_type,omitempty"`
	ValidationType      string `json:"validation_type,omitempty"`
}
type DecorationGroupQueryResponse struct {
	GroupUrl           string            `json:"group_url,omitempty"`
	Icon               string            `json:"icon,omitempty"`
	ManagementOptions  ManagementOptions `json:"management_options,omitempty"`
	MemberAmount       int64             `json:"member_amount,omitempty"`
	OpenConversationId string            `json:"open_conversation_id,omitempty"`
	OwnerStaffId       string            `json:"owner_staff_id,omitempty"`
	SubAdminStaffIds   []string          `json:"sub_admin_staff_ids,omitempty"`
	TemplateId         string            `json:"template_id,omitempty"`
	Title              string            `json:"title,omitempty"`
}
