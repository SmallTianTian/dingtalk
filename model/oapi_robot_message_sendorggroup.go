package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotMessageSendorggroupRequest() *OapiRobotMessageSendorggroupRequest {
	return &OapiRobotMessageSendorggroupRequest{}
}

type OapiRobotMessageSendorggroupRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiRobotMessageSendorggroupResponse
	ChatbotId          string
	MsgKey             string
	MsgParam           string
	OpenConversationId string
	Token              string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiRobotMessageSendorggroupRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotMessageSendorggroupRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotMessageSendorggroupRequest) SetChatbotId(chatbotId2 string) {
	this.ChatbotId = chatbotId2
}
func (this *OapiRobotMessageSendorggroupRequest) GetChatbotId() string {
	return this.ChatbotId
}
func (this *OapiRobotMessageSendorggroupRequest) SetMsgKey(msgKey2 string) {
	this.MsgKey = msgKey2
}
func (this *OapiRobotMessageSendorggroupRequest) GetMsgKey() string {
	return this.MsgKey
}
func (this *OapiRobotMessageSendorggroupRequest) SetMsgParam(msgParam2 string) {
	this.MsgParam = msgParam2
}
func (this *OapiRobotMessageSendorggroupRequest) GetMsgParam() string {
	return this.MsgParam
}
func (this *OapiRobotMessageSendorggroupRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiRobotMessageSendorggroupRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiRobotMessageSendorggroupRequest) SetToken(token2 string) {
	this.Token = token2
}
func (this *OapiRobotMessageSendorggroupRequest) GetToken() string {
	return this.Token
}
func (this *OapiRobotMessageSendorggroupRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.message.sendorggroup"
}
func (this *OapiRobotMessageSendorggroupRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotMessageSendorggroupRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotMessageSendorggroupRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotMessageSendorggroupRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotMessageSendorggroupRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_id"] = this.ChatbotId
	txtParams["msg_key"] = this.MsgKey
	txtParams["msg_param"] = this.MsgParam
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["token"] = this.Token
	return txtParams
}
func (this *OapiRobotMessageSendorggroupRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MsgKey, "msgKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotMessageSendorggroupRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotMessageSendorggroupRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotMessageSendorggroupResponse struct {
	taobao.TaobaoResponse
	Errcode int64                       `json:"errcode,omitempty"`
	Errmsg  string                      `json:"errmsg,omitempty"`
	Result  GroupMessageSendTopResponse `json:"result,omitempty"`
	Success bool                        `json:"success,omitempty"`
}
