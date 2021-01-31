package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpChatbotListorgbotRequest() *CorpChatbotListorgbotRequest {
	return &CorpChatbotListorgbotRequest{}
}

type CorpChatbotListorgbotRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpChatbotListorgbotResponse
	AgentId         int64
	TopHttpMethod   string
	TopResponseType string
	Type            string
}

func (this *CorpChatbotListorgbotRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpChatbotListorgbotRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *CorpChatbotListorgbotRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *CorpChatbotListorgbotRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *CorpChatbotListorgbotRequest) GetType() string {
	return this.Type
}
func (this *CorpChatbotListorgbotRequest) GetApiMethodName() string {
	return "dingtalk.corp.chatbot.listorgbot"
}
func (this *CorpChatbotListorgbotRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpChatbotListorgbotRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpChatbotListorgbotRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpChatbotListorgbotRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpChatbotListorgbotRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpChatbotListorgbotRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["type"] = this.Type
	return txtParams
}
func (this *CorpChatbotListorgbotRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpChatbotListorgbotRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpChatbotListorgbotRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpChatbotListorgbotResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type ChatbotModel struct {
	BotType         int64  `json:"bot_type,omitempty"`
	Breif           string `json:"breif,omitempty"`
	ChatbotId       int64  `json:"chatbot_id,omitempty"`
	CorpId          string `json:"corp_id,omitempty"`
	Description     string `json:"description,omitempty"`
	Icon            string `json:"icon,omitempty"`
	MicroappAgentId int64  `json:"microapp_agent_id,omitempty"`
	Name            string `json:"name,omitempty"`
	OutgoingToken   string `json:"outgoing_token,omitempty"`
	OutgoingUrl     string `json:"outgoing_url,omitempty"`
	Type            string `json:"type,omitempty"`
}
