package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatScenegroupMemberAddRequest() *OapiImChatScenegroupMemberAddRequest {
	return &OapiImChatScenegroupMemberAddRequest{}
}

type OapiImChatScenegroupMemberAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiImChatScenegroupMemberAddResponse
	ContactStaffIds    string
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
	UserIds            string
}

func (this *OapiImChatScenegroupMemberAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatScenegroupMemberAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatScenegroupMemberAddRequest) SetContactStaffIds(contactStaffIds2 string) {
	this.ContactStaffIds = contactStaffIds2
}
func (this *OapiImChatScenegroupMemberAddRequest) GetContactStaffIds() string {
	return this.ContactStaffIds
}
func (this *OapiImChatScenegroupMemberAddRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiImChatScenegroupMemberAddRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiImChatScenegroupMemberAddRequest) SetUserIds(userIds2 string) {
	this.UserIds = userIds2
}
func (this *OapiImChatScenegroupMemberAddRequest) GetUserIds() string {
	return this.UserIds
}
func (this *OapiImChatScenegroupMemberAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.scenegroup.member.add"
}
func (this *OapiImChatScenegroupMemberAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatScenegroupMemberAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatScenegroupMemberAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatScenegroupMemberAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatScenegroupMemberAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contact_staff_ids"] = this.ContactStaffIds
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["user_ids"] = this.UserIds
	return txtParams
}
func (this *OapiImChatScenegroupMemberAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.ContactStaffIds, 999, "contactStaffIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatScenegroupMemberAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatScenegroupMemberAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatScenegroupMemberAddResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
