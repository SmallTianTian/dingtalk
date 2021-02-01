package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpChatbotAddchatbotinstanceRequest() *CorpChatbotAddchatbotinstanceRequest {
	return &CorpChatbotAddchatbotinstanceRequest{}
}

type CorpChatbotAddchatbotinstanceRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               CorpChatbotAddchatbotinstanceResponse
	ChatbotId          string
	IconMediaId        string
	Name               string
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *CorpChatbotAddchatbotinstanceRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpChatbotAddchatbotinstanceRequest) SetChatbotId(chatbotId2 string) {
	this.ChatbotId = chatbotId2
}
func (this *CorpChatbotAddchatbotinstanceRequest) GetChatbotId() string {
	return this.ChatbotId
}
func (this *CorpChatbotAddchatbotinstanceRequest) SetIconMediaId(iconMediaId2 string) {
	this.IconMediaId = iconMediaId2
}
func (this *CorpChatbotAddchatbotinstanceRequest) GetIconMediaId() string {
	return this.IconMediaId
}
func (this *CorpChatbotAddchatbotinstanceRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *CorpChatbotAddchatbotinstanceRequest) GetName() string {
	return this.Name
}
func (this *CorpChatbotAddchatbotinstanceRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *CorpChatbotAddchatbotinstanceRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *CorpChatbotAddchatbotinstanceRequest) GetApiMethodName() string {
	return "dingtalk.corp.chatbot.addchatbotinstance"
}
func (this *CorpChatbotAddchatbotinstanceRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpChatbotAddchatbotinstanceRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpChatbotAddchatbotinstanceRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpChatbotAddchatbotinstanceRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpChatbotAddchatbotinstanceRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpChatbotAddchatbotinstanceRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_id"] = this.ChatbotId
	txtParams["icon_media_id"] = this.IconMediaId
	txtParams["name"] = this.Name
	txtParams["open_conversation_id"] = this.OpenConversationId
	return txtParams
}
func (this *CorpChatbotAddchatbotinstanceRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatbotId, "chatbotId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpChatbotAddchatbotinstanceRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpChatbotAddchatbotinstanceRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpChatbotAddchatbotinstanceResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type ChatbotInstanceVo struct {
	ChatbotUserId string `json:"chatbot_user_id,omitempty"`
	Webhook       string `json:"webhook,omitempty"`
}
