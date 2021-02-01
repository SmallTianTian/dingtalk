package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotMessageSendotoRequest() *OapiRobotMessageSendotoRequest {
	return &OapiRobotMessageSendotoRequest{}
}

type OapiRobotMessageSendotoRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRobotMessageSendotoResponse
	ChatbotId       string
	MsgKey          string
	MsgParam        string
	StaffId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRobotMessageSendotoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotMessageSendotoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotMessageSendotoRequest) SetChatbotId(chatbotId2 string) {
	this.ChatbotId = chatbotId2
}
func (this *OapiRobotMessageSendotoRequest) GetChatbotId() string {
	return this.ChatbotId
}
func (this *OapiRobotMessageSendotoRequest) SetMsgKey(msgKey2 string) {
	this.MsgKey = msgKey2
}
func (this *OapiRobotMessageSendotoRequest) GetMsgKey() string {
	return this.MsgKey
}
func (this *OapiRobotMessageSendotoRequest) SetMsgParam(msgParam2 string) {
	this.MsgParam = msgParam2
}
func (this *OapiRobotMessageSendotoRequest) GetMsgParam() string {
	return this.MsgParam
}
func (this *OapiRobotMessageSendotoRequest) SetStaffId(staffId2 string) {
	this.StaffId = staffId2
}
func (this *OapiRobotMessageSendotoRequest) GetStaffId() string {
	return this.StaffId
}
func (this *OapiRobotMessageSendotoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.message.sendoto"
}
func (this *OapiRobotMessageSendotoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotMessageSendotoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotMessageSendotoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotMessageSendotoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotMessageSendotoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_id"] = this.ChatbotId
	txtParams["msg_key"] = this.MsgKey
	txtParams["msg_param"] = this.MsgParam
	txtParams["staff_id"] = this.StaffId
	return txtParams
}
func (this *OapiRobotMessageSendotoRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatbotId, "chatbotId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotMessageSendotoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotMessageSendotoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotMessageSendotoResponse struct {
	taobao.TaobaoResponse
	Result  OtoMessageSendTopResponse `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
