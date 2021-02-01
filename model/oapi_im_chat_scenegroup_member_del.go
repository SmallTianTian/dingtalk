package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatScenegroupMemberDeleteRequest() *OapiImChatScenegroupMemberDeleteRequest {
	return &OapiImChatScenegroupMemberDeleteRequest{}
}

type OapiImChatScenegroupMemberDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiImChatScenegroupMemberDeleteResponse
	ContactStaffIds    string
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
	UserIds            string
}

func (this *OapiImChatScenegroupMemberDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatScenegroupMemberDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatScenegroupMemberDeleteRequest) SetContactStaffIds(contactStaffIds2 string) {
	this.ContactStaffIds = contactStaffIds2
}
func (this *OapiImChatScenegroupMemberDeleteRequest) GetContactStaffIds() string {
	return this.ContactStaffIds
}
func (this *OapiImChatScenegroupMemberDeleteRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiImChatScenegroupMemberDeleteRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiImChatScenegroupMemberDeleteRequest) SetUserIds(userIds2 string) {
	this.UserIds = userIds2
}
func (this *OapiImChatScenegroupMemberDeleteRequest) GetUserIds() string {
	return this.UserIds
}
func (this *OapiImChatScenegroupMemberDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.scenegroup.member.delete"
}
func (this *OapiImChatScenegroupMemberDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatScenegroupMemberDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatScenegroupMemberDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatScenegroupMemberDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatScenegroupMemberDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contact_staff_ids"] = this.ContactStaffIds
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["user_ids"] = this.UserIds
	return txtParams
}
func (this *OapiImChatScenegroupMemberDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.ContactStaffIds, 999, "contactStaffIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatScenegroupMemberDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatScenegroupMemberDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatScenegroupMemberDeleteResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
