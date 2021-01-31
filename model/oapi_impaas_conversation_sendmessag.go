package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImpaasConversationSendmessageRequest() *OapiImpaasConversationSendmessageRequest {
	return &OapiImpaasConversationSendmessageRequest{}
}

type OapiImpaasConversationSendmessageRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasConversationSendmessageResponse
	Chatid          string
	Content         string
	TopHttpMethod   string
	TopResponseType string
	Type            int64
}

func (this *OapiImpaasConversationSendmessageRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasConversationSendmessageRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasConversationSendmessageRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiImpaasConversationSendmessageRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiImpaasConversationSendmessageRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiImpaasConversationSendmessageRequest) SetContentString(content2 string) {
	this.Content = content2
}
func (this *OapiImpaasConversationSendmessageRequest) GetContent() string {
	return this.Content
}
func (this *OapiImpaasConversationSendmessageRequest) SetType(type2 int64) {
	this.Type = type2
}
func (this *OapiImpaasConversationSendmessageRequest) GetType() int64 {
	return this.Type
}
func (this *OapiImpaasConversationSendmessageRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.conversation.sendmessage"
}
func (this *OapiImpaasConversationSendmessageRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasConversationSendmessageRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasConversationSendmessageRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasConversationSendmessageRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasConversationSendmessageRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiImpaasConversationSendmessageRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImpaasConversationSendmessageRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasConversationSendmessageRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImpaasConversationSendmessageResponse struct {
	taobao.TaobaoResponse
}
