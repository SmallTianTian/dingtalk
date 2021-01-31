package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpChatbotCreateorgbotRequest() *CorpChatbotCreateorgbotRequest {
	return &CorpChatbotCreateorgbotRequest{}
}

type CorpChatbotCreateorgbotRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               CorpChatbotCreateorgbotResponse
	CreateChatBotModel string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *CorpChatbotCreateorgbotRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpChatbotCreateorgbotRequest) SetCreateChatBotModel(createChatBotModel2 string) {
	this.CreateChatBotModel = createChatBotModel2
}
func (this *CorpChatbotCreateorgbotRequest) GetCreateChatBotModel() string {
	return this.CreateChatBotModel
}
func (this *CorpChatbotCreateorgbotRequest) GetApiMethodName() string {
	return "dingtalk.corp.chatbot.createorgbot"
}
func (this *CorpChatbotCreateorgbotRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpChatbotCreateorgbotRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpChatbotCreateorgbotRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpChatbotCreateorgbotRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpChatbotCreateorgbotRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpChatbotCreateorgbotRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["create_chat_bot_model"] = this.CreateChatBotModel
	return txtParams
}
func (this *CorpChatbotCreateorgbotRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpChatbotCreateorgbotRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpChatbotCreateorgbotRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CreateChatBotModel struct {
	BotType         int64  `json:"bot_type,omitempty"`
	Breif           string `json:"breif,omitempty"`
	CorpId          string `json:"corp_id,omitempty"`
	Description     string `json:"description,omitempty"`
	Icon            string `json:"icon,omitempty"`
	MicroappAgentId int64  `json:"microapp_agent_id,omitempty"`
	Name            string `json:"name,omitempty"`
	OutgoingToken   string `json:"outgoing_token,omitempty"`
	OutgoingUrl     string `json:"outgoing_url,omitempty"`
	Type            string `json:"type,omitempty"`
}
type CorpChatbotCreateorgbotResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
