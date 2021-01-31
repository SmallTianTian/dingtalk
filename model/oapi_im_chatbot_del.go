package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatbotDeleteRequest() *OapiImChatbotDeleteRequest {
	return &OapiImChatbotDeleteRequest{}
}

type OapiImChatbotDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiImChatbotDeleteResponse
	ChatbotUserId      string
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiImChatbotDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatbotDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatbotDeleteRequest) SetChatbotUserId(chatbotUserId2 string) {
	this.ChatbotUserId = chatbotUserId2
}
func (this *OapiImChatbotDeleteRequest) GetChatbotUserId() string {
	return this.ChatbotUserId
}
func (this *OapiImChatbotDeleteRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiImChatbotDeleteRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiImChatbotDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chatbot.delete"
}
func (this *OapiImChatbotDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatbotDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatbotDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatbotDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatbotDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_user_id"] = this.ChatbotUserId
	txtParams["open_conversation_id"] = this.OpenConversationId
	return txtParams
}
func (this *OapiImChatbotDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatbotUserId, "chatbotUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatbotDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatbotDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatbotDeleteResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
