package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatServicegroupMemberUpdateRequest() *OapiImChatServicegroupMemberUpdateRequest {
	return &OapiImChatServicegroupMemberUpdateRequest{}
}

type OapiImChatServicegroupMemberUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiImChatServicegroupMemberUpdateResponse
	Action            string
	ChatId            string
	MemberDingtalkIds string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiImChatServicegroupMemberUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatServicegroupMemberUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatServicegroupMemberUpdateRequest) SetAction(action2 string) {
	this.Action = action2
}
func (this *OapiImChatServicegroupMemberUpdateRequest) GetAction() string {
	return this.Action
}
func (this *OapiImChatServicegroupMemberUpdateRequest) SetChatId(chatId2 string) {
	this.ChatId = chatId2
}
func (this *OapiImChatServicegroupMemberUpdateRequest) GetChatId() string {
	return this.ChatId
}
func (this *OapiImChatServicegroupMemberUpdateRequest) SetMemberDingtalkIds(memberDingtalkIds2 string) {
	this.MemberDingtalkIds = memberDingtalkIds2
}
func (this *OapiImChatServicegroupMemberUpdateRequest) GetMemberDingtalkIds() string {
	return this.MemberDingtalkIds
}
func (this *OapiImChatServicegroupMemberUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.servicegroup.member.update"
}
func (this *OapiImChatServicegroupMemberUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatServicegroupMemberUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatServicegroupMemberUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatServicegroupMemberUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatServicegroupMemberUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["action"] = this.Action
	txtParams["chat_id"] = this.ChatId
	txtParams["member_dingtalk_ids"] = this.MemberDingtalkIds
	return txtParams
}
func (this *OapiImChatServicegroupMemberUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Action, "action"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatServicegroupMemberUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatServicegroupMemberUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatServicegroupMemberUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
