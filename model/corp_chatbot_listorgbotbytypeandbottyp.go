package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpChatbotListorgbotbytypeandbottypeRequest() *CorpChatbotListorgbotbytypeandbottypeRequest {
	return &CorpChatbotListorgbotbytypeandbottypeRequest{}
}

type CorpChatbotListorgbotbytypeandbottypeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpChatbotListorgbotbytypeandbottypeResponse
	BotType         int64
	TopHttpMethod   string
	TopResponseType string
	Type            string
}

func (this *CorpChatbotListorgbotbytypeandbottypeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) SetBotType(botType2 int64) {
	this.BotType = botType2
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) GetBotType() int64 {
	return this.BotType
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) GetType() string {
	return this.Type
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) GetApiMethodName() string {
	return "dingtalk.corp.chatbot.listorgbotbytypeandbottype"
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["bot_type"] = this.BotType
	txtParams["type"] = this.Type
	return txtParams
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BotType, "botType"); err != nil {
		return err
	}
	return nil
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpChatbotListorgbotbytypeandbottypeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpChatbotListorgbotbytypeandbottypeResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
