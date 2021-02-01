package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatMemberFriendswitchUpdateRequest() *OapiChatMemberFriendswitchUpdateRequest {
	return &OapiChatMemberFriendswitchUpdateRequest{}
}

type OapiChatMemberFriendswitchUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatMemberFriendswitchUpdateResponse
	Chatid          string
	IsProhibit      bool
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatMemberFriendswitchUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatMemberFriendswitchUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatMemberFriendswitchUpdateRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatMemberFriendswitchUpdateRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatMemberFriendswitchUpdateRequest) SetIsProhibit(isProhibit2 bool) {
	this.IsProhibit = isProhibit2
}
func (this *OapiChatMemberFriendswitchUpdateRequest) GetIsProhibit() bool {
	return this.IsProhibit
}
func (this *OapiChatMemberFriendswitchUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.member.friendswitch.update"
}
func (this *OapiChatMemberFriendswitchUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatMemberFriendswitchUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatMemberFriendswitchUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatMemberFriendswitchUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatMemberFriendswitchUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	txtParams["is_prohibit"] = this.IsProhibit
	return txtParams
}
func (this *OapiChatMemberFriendswitchUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatMemberFriendswitchUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatMemberFriendswitchUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatMemberFriendswitchUpdateResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
