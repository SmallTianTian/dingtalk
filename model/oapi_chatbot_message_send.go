package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatbotMessageSendRequest() *OapiChatbotMessageSendRequest {
	return &OapiChatbotMessageSendRequest{}
}

type OapiChatbotMessageSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatbotMessageSendResponse
	ChatbotId       string
	Message         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiChatbotMessageSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatbotMessageSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatbotMessageSendRequest) SetChatbotId(chatbotId2 string) {
	this.ChatbotId = chatbotId2
}
func (this *OapiChatbotMessageSendRequest) GetChatbotId() string {
	return this.ChatbotId
}
func (this *OapiChatbotMessageSendRequest) SetMessage(message2 string) {
	this.Message = message2
}
func (this *OapiChatbotMessageSendRequest) GetMessage() string {
	return this.Message
}
func (this *OapiChatbotMessageSendRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiChatbotMessageSendRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiChatbotMessageSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chatbot.message.send"
}
func (this *OapiChatbotMessageSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatbotMessageSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatbotMessageSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatbotMessageSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatbotMessageSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_id"] = this.ChatbotId
	txtParams["message"] = this.Message
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiChatbotMessageSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatbotId, "chatbotId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatbotMessageSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatbotMessageSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatbotMessageSendResponse struct {
	taobao.TaobaoResponse
	Result  OtoMessageResponeModel `json:"result,omitempty"`
	Success bool                   `json:"success,omitempty"`
}
type OtoMessageResponeModel struct {
	MessageId string `json:"message_id,omitempty"`
}
