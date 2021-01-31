package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpMessageCorpconversationSendmockRequest() *CorpMessageCorpconversationSendmockRequest {
	return &CorpMessageCorpconversationSendmockRequest{}
}

type CorpMessageCorpconversationSendmockRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpMessageCorpconversationSendmockResponse
	Message         string
	MessageType     string
	MicroappAgentId int64
	ToParty         string
	ToUser          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpMessageCorpconversationSendmockRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpMessageCorpconversationSendmockRequest) SetMessage(message2 string) {
	this.Message = message2
}
func (this *CorpMessageCorpconversationSendmockRequest) SetMessageString(message2 string) {
	this.Message = message2
}
func (this *CorpMessageCorpconversationSendmockRequest) GetMessage() string {
	return this.Message
}
func (this *CorpMessageCorpconversationSendmockRequest) SetMessageType(messageType2 string) {
	this.MessageType = messageType2
}
func (this *CorpMessageCorpconversationSendmockRequest) GetMessageType() string {
	return this.MessageType
}
func (this *CorpMessageCorpconversationSendmockRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *CorpMessageCorpconversationSendmockRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *CorpMessageCorpconversationSendmockRequest) SetToParty(toParty2 string) {
	this.ToParty = toParty2
}
func (this *CorpMessageCorpconversationSendmockRequest) GetToParty() string {
	return this.ToParty
}
func (this *CorpMessageCorpconversationSendmockRequest) SetToUser(toUser2 string) {
	this.ToUser = toUser2
}
func (this *CorpMessageCorpconversationSendmockRequest) GetToUser() string {
	return this.ToUser
}
func (this *CorpMessageCorpconversationSendmockRequest) GetApiMethodName() string {
	return "dingtalk.corp.message.corpconversation.sendmock"
}
func (this *CorpMessageCorpconversationSendmockRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpMessageCorpconversationSendmockRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpMessageCorpconversationSendmockRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpMessageCorpconversationSendmockRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpMessageCorpconversationSendmockRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpMessageCorpconversationSendmockRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["message"] = this.Message
	txtParams["message_type"] = this.MessageType
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["to_party"] = this.ToParty
	txtParams["to_user"] = this.ToUser
	return txtParams
}
func (this *CorpMessageCorpconversationSendmockRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Message, "message"); err != nil {
		return err
	}
	return nil
}
func (this *CorpMessageCorpconversationSendmockRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpMessageCorpconversationSendmockRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpMessageCorpconversationSendmockResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
