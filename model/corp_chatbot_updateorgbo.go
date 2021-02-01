package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpChatbotUpdateorgbotRequest() *CorpChatbotUpdateorgbotRequest {
	return &CorpChatbotUpdateorgbotRequest{}
}

type CorpChatbotUpdateorgbotRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpChatbotUpdateorgbotResponse
	ChatbotId       int64
	Icon            string
	Name            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpChatbotUpdateorgbotRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpChatbotUpdateorgbotRequest) SetChatbotId(chatbotId2 int64) {
	this.ChatbotId = chatbotId2
}
func (this *CorpChatbotUpdateorgbotRequest) GetChatbotId() int64 {
	return this.ChatbotId
}
func (this *CorpChatbotUpdateorgbotRequest) SetIcon(icon2 string) {
	this.Icon = icon2
}
func (this *CorpChatbotUpdateorgbotRequest) GetIcon() string {
	return this.Icon
}
func (this *CorpChatbotUpdateorgbotRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *CorpChatbotUpdateorgbotRequest) GetName() string {
	return this.Name
}
func (this *CorpChatbotUpdateorgbotRequest) GetApiMethodName() string {
	return "dingtalk.corp.chatbot.updateorgbot"
}
func (this *CorpChatbotUpdateorgbotRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpChatbotUpdateorgbotRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpChatbotUpdateorgbotRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpChatbotUpdateorgbotRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpChatbotUpdateorgbotRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpChatbotUpdateorgbotRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_id"] = this.ChatbotId
	txtParams["icon"] = this.Icon
	txtParams["name"] = this.Name
	return txtParams
}
func (this *CorpChatbotUpdateorgbotRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatbotId, "chatbotId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpChatbotUpdateorgbotRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpChatbotUpdateorgbotRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpChatbotUpdateorgbotResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
