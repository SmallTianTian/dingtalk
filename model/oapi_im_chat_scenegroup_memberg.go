package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatScenegroupMemberGetRequest() *OapiImChatScenegroupMemberGetRequest {
	return &OapiImChatScenegroupMemberGetRequest{}
}

type OapiImChatScenegroupMemberGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiImChatScenegroupMemberGetResponse
	Cursor             string
	OpenConversationId string
	Size               int64
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiImChatScenegroupMemberGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatScenegroupMemberGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatScenegroupMemberGetRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiImChatScenegroupMemberGetRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiImChatScenegroupMemberGetRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiImChatScenegroupMemberGetRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiImChatScenegroupMemberGetRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiImChatScenegroupMemberGetRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiImChatScenegroupMemberGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.scenegroup.member.get"
}
func (this *OapiImChatScenegroupMemberGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatScenegroupMemberGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatScenegroupMemberGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatScenegroupMemberGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatScenegroupMemberGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiImChatScenegroupMemberGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatScenegroupMemberGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatScenegroupMemberGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatScenegroupMemberGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                             `json:"errcode,omitempty"`
	Errmsg  string                            `json:"errmsg,omitempty"`
	Result  OpenSceneGroupMemberQueryResponse `json:"result,omitempty"`
	Success bool                              `json:"success,omitempty"`
}
type OpenSceneGroupMemberQueryResponse struct {
	HasMore               bool     `json:"has_more,omitempty"`
	MemberContactStaffIds []string `json:"member_contact_staff_ids,omitempty"`
	MemberUserIds         []string `json:"member_user_ids,omitempty"`
	NextCursor            string   `json:"next_cursor,omitempty"`
	UnionIds              []string `json:"union_ids,omitempty"`
}
