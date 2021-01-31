package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImpaasConversationModifymemberRequest() *OapiImpaasConversationModifymemberRequest {
	return &OapiImpaasConversationModifymemberRequest{}
}

type OapiImpaasConversationModifymemberRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasConversationModifymemberResponse
	Channel         string
	Chatid          string
	MemberidList    string
	TopHttpMethod   string
	TopResponseType string
	Type            int64
}

func (this *OapiImpaasConversationModifymemberRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasConversationModifymemberRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasConversationModifymemberRequest) SetChannel(channel2 string) {
	this.Channel = channel2
}
func (this *OapiImpaasConversationModifymemberRequest) GetChannel() string {
	return this.Channel
}
func (this *OapiImpaasConversationModifymemberRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiImpaasConversationModifymemberRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiImpaasConversationModifymemberRequest) SetMemberidList(memberidList2 string) {
	this.MemberidList = memberidList2
}
func (this *OapiImpaasConversationModifymemberRequest) GetMemberidList() string {
	return this.MemberidList
}
func (this *OapiImpaasConversationModifymemberRequest) SetType(type2 int64) {
	this.Type = type2
}
func (this *OapiImpaasConversationModifymemberRequest) GetType() int64 {
	return this.Type
}
func (this *OapiImpaasConversationModifymemberRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.conversation.modifymember"
}
func (this *OapiImpaasConversationModifymemberRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasConversationModifymemberRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasConversationModifymemberRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasConversationModifymemberRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasConversationModifymemberRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["channel"] = this.Channel
	txtParams["chatid"] = this.Chatid
	txtParams["memberid_list"] = this.MemberidList
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiImpaasConversationModifymemberRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Channel, "channel"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImpaasConversationModifymemberRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasConversationModifymemberRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImpaasConversationModifymemberResponse struct {
	taobao.TaobaoResponse
}
